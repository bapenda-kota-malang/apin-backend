package pembayaran

import (
	"fmt"
	"strconv"
	"time"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	mpegawai "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	msppt "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt/pembayaran"
	muser "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	spegawai "github.com/bapenda-kota-malang/apin-backend/internal/services/pegawai"
	ssppt "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt"
	userservice "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "spptpembayaran"

func Create(input m.CreateDto, userId int) (any, error) {
	var (
		data                   m.SpptPembayaran
		dataLogSpptPembayaran  []m.SpptPembayaran
		countLogSpptPembayaran int64
	)
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// get data sppt
	resSppt, err := ssppt.GetByNop(
		&data.Provinsi_Kd,
		&data.Daerah_Kd,
		&data.Kecamatan_Kd,
		&data.Kelurahan_Kd,
		&data.Blok_Kd,
		&data.NoUrut_Kd,
		&data.JenisOp_Kd,
		&data.TahunPajakSppt,
	)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data sppt: "+err.Error(), data)
	}
	dataSppt := resSppt.(rp.OKSimple).Data.(msppt.Sppt)
	if dataSppt.StatusPembayaran_sppt == nil {
		return sh.SetError("request", "create-data", source, "failed", "sppt status pembayaran tidak valid", data)
	}
	if *dataSppt.StatusPembayaran_sppt == "1" {
		return sh.SetError("request", "create-data", source, "failed", "sppt sudah lunas", data)
	}
	if dataSppt.PBBygHarusDibayar_sppt == nil {
		return sh.SetError("request", "create-data", source, "failed", "sppt pbb yang harus dibayar tidak valid", data)
	}

	// get count log sppt pembayaran from data nop + tahun pajak
	resLogSpptPembayaran := a.DB.
		Model(&m.SpptPembayaran{}).
		Select("JumlahSpptYgDibayar").
		Scopes(gh.Filter(m.FilterDto{
			Provinsi_Kd:    &data.Provinsi_Kd,
			Daerah_Kd:      &data.Daerah_Kd,
			Kecamatan_Kd:   &data.Kecamatan_Kd,
			Kelurahan_Kd:   &data.Kelurahan_Kd,
			Blok_Kd:        &data.Blok_Kd,
			NoUrut_Kd:      &data.NoUrut_Kd,
			JenisOp_Kd:     &data.JenisOp_Kd,
			TahunPajakSppt: &data.TahunPajakSppt,
		})).
		Count(&countLogSpptPembayaran).
		Find(&dataLogSpptPembayaran)
	if resLogSpptPembayaran.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data log sppt pembayaran: "+resLogSpptPembayaran.Error.Error(), data)
	}

	// get data user from auth user id
	resUser, err := userservice.GetDetail(userId)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data user: "+err.Error(), data)
	}
	if resUser == nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data user: data tidak ditemukan", data)
	}
	dataUser := resUser.(rp.OKSimple).Data.(*muser.User)
	if dataUser.Position != 1 {
		return sh.SetError("request", "create-data", source, "failed", "position user tidak valid", data)
	}

	// get data pegawai from user ref id
	resPegawai, err := spegawai.GetDetail(dataUser.Ref_Id)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data pegawai: "+err.Error(), data)
	}
	if resPegawai == nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data pegawai: data tidak ditemukan", data)
	}
	dataPegawai := resPegawai.(rp.OKSimple).Data.(*mpegawai.Pegawai)

	// set data
	data.KanwilBank_Kd = dataSppt.KanwilBank_Id
	data.KppbbBank_Kd = dataSppt.KPPBBbank_Id
	data.BankPersepsi_Kd = dataSppt.BankPersepsi_Id
	data.BankTunggal_Kd = dataSppt.BankTunggal_Id
	data.TP_Kd = dataSppt.TP_Id
	data.PembayaranSpptKe = uint8(countLogSpptPembayaran + 1)
	data.TglRekamBayarSppt = time.Now()
	data.NipRekamBayarSppt = dataPegawai.Nip

	err = a.DB.Transaction(func(tx *gorm.DB) error {
		// save data to db
		if result := tx.Create(&data); result.Error != nil {
			return result.Error
		}
		// sum all sppt dibayar
		sumSpptDibayar := data.JumlahSpptYgDibayar
		for _, v := range dataLogSpptPembayaran {
			sumSpptDibayar += v.JumlahSpptYgDibayar
		}
		if sumSpptDibayar < int64(*dataSppt.PBBygHarusDibayar_sppt) {
			return nil
		}
		// if sum >= sppt pbb yang harus dibayar update status to "1"
		statusPembayaranDone := "1"
		dataSppt.StatusPembayaran_sppt = &statusPembayaranDone
		if result := tx.Save(&dataSppt); result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data: "+err.Error(), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.SpptPembayaran
	var count int64
	var pagination gh.Pagination

	result := a.DB.
		Model(&m.SpptPembayaran{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: data,
	}, nil
}

func GetDetail(id int) (interface{}, error) {
	var data *m.SpptPembayaran
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto) (interface{}, error) {
	// var data *m.SpptPembayaran

	// result := a.DB.First(&data, id)
	// if result.RowsAffected == 0 {
	// 	return nil, fmt.Errorf("data tidak dapat ditemukan")
	// }
	// if err := sc.CopyWithOption(&data, input, sc.Option{IgnoreEmpty: true}); err != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	// }
	// err := a.DB.Transaction(func(tx *gorm.DB) error {
	// 	if result := tx.Save(&data); result.Error != nil {
	// 		return result.Error
	// 	}
	// 	return nil
	// })
	// if err != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data: %s", err), data)
	// }

	// return rp.OK{
	// 	Meta: t.IS{
	// 		"affected": strconv.Itoa(int(result.RowsAffected)),
	// 	},
	// 	Data: data,
	// }, nil
	return nil, nil
}

func Delete(id int) (interface{}, error) {
	var data *m.SpptPembayaran
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("data tidak dapat ditemukan")
	}

	result = a.DB.Delete(&data, id)
	status := "deleted"
	if result.RowsAffected == 0 {
		data = nil
		status = "no deletion"
	}

	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil
}
