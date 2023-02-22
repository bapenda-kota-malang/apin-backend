package tandaterima

import (
	"fmt"
	"strconv"
	"time"

	sc "github.com/jinzhu/copier"

	mpegawai "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt/tandaterima"
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

const source = "sppttandaterima"

func Create(input m.CreateDto, userId int) (any, error) {
	var (
		data      m.SpptTandaTerima
		countData int64
	)
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// get data sppt
	_, err := ssppt.GetByNop(
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

	// check exist, if exist reject create data
	resSpptTandaTerimaId := a.DB.
		Model(&m.SpptTandaTerima{}).
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
		Count(&countData)
	if resSpptTandaTerimaId.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data log sppt tanda terima: "+resSpptTandaTerimaId.Error.Error(), data)
	}
	if countData > 0 {
		return sh.SetError("request", "create-data", source, "failed", "gagal membuat data sppt tanda terima: data sudah pernah dibuat", data)
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
	dataPegawai := resPegawai.(rp.OKSimple).Data.(*mpegawai.OutputDto)

	// set data
	data.TglRekamTtrSppt = time.Now()
	data.NipRekamTtrSppt = dataPegawai.Nip

	// save data to db
	if result := a.DB.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+result.Error.Error(), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.SpptTandaTerima
	var count int64
	var pagination gh.Pagination

	result := a.DB.
		Model(&m.SpptTandaTerima{}).
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
	var data *m.SpptTandaTerima
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

func Update(id int, input m.UpdateDto, userId int) (interface{}, error) {
	var data *m.SpptTandaTerima

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("data tidak dapat ditemukan")
	}
	if err := sc.CopyWithOption(&data, input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
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
	dataPegawai := resPegawai.(rp.OKSimple).Data.(*mpegawai.OutputDto)

	// set data
	data.TglRekamTtrSppt = time.Now()
	data.NipRekamTtrSppt = dataPegawai.Nip

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("gagal mengubah data: %s", result.Error), data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (interface{}, error) {
	var data *m.SpptTandaTerima
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
