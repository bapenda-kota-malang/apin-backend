package regobjekpajakpbb

import (
	"errors"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	maop "github.com/bapenda-kota-malang/apin-backend/internal/models/anggotaobjekpajak"
	mkk "github.com/bapenda-kota-malang/apin-backend/internal/models/kunjungankembali"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	mopb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	mopp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	mraop "github.com/bapenda-kota-malang/apin-backend/internal/models/reganggotaobjekpajak"
	mrkk "github.com/bapenda-kota-malang/apin-backend/internal/models/regkunjungankembali"
	mropb "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbumi"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakpbb"
	mrwp "github.com/bapenda-kota-malang/apin-backend/internal/models/regwajibpajakpbb"
	mwpp "github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajakpbb"
	saop "github.com/bapenda-kota-malang/apin-backend/internal/services/anggotaobjekpajak"
	skk "github.com/bapenda-kota-malang/apin-backend/internal/services/kunjungankembali"
	sopb "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakbumi"
	sraop "github.com/bapenda-kota-malang/apin-backend/internal/services/reganggotaobjekpajak"
	srkk "github.com/bapenda-kota-malang/apin-backend/internal/services/regkunjungankembali"
	sropb "github.com/bapenda-kota-malang/apin-backend/internal/services/regobjekpajakbumi"
	srwp "github.com/bapenda-kota-malang/apin-backend/internal/services/regwajibpajakpbb"
	swp "github.com/bapenda-kota-malang/apin-backend/internal/services/wajibpajakpbb"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
)

const source = "reg objekpajakpbb"

func Create(input m.CreateDto) (any, error) {
	var data m.RegObjekPajakPbb
	var dataWajibPajakPbb mrwp.CreateDto
	var dataAnggotaObjekPajak mraop.CreateDto
	var dataObjekPajakBumi mropb.CreateDto
	var dataKunjunganKembali mrkk.CreateDto
	var respDataWajibPajakPbb interface{}
	var respDataAnggotaObjekPajak interface{}
	var respDataObjekPajakBumi interface{}
	var respDataKunjunganKembali interface{}
	var resp t.II

	// parsing nop
	resultNop, kode := sh.NopParser(*input.Nop)

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// copy data wajibpajakPbb
	if err := sc.Copy(&dataWajibPajakPbb, input.RegWajibPajakPbbs); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg wajibpajakpbb", input.RegWajibPajakPbbs)
	}

	// copy data objek pajak bumi
	if err := sc.Copy(&dataObjekPajakBumi, input.RegObjekPajakBumis); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg objek pajak bumi", input.RegObjekPajakBumis)
	}

	if input.NopBersama != nil {
		// copy data anggota objek pajak
		if err := sc.Copy(&dataAnggotaObjekPajak, input.RegAnggotaObjekPajaks); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg anggota objek pajak", input.RegAnggotaObjekPajaks)
		}

		// data induk
		resultNopBersama, kodeBersama := sh.NopParser(*input.NopBersama)
		dataAnggotaObjekPajak.Induk_Provinsi_Kode = &resultNopBersama[0]
		dataAnggotaObjekPajak.Induk_Daerah_Kode = &resultNopBersama[1]
		dataAnggotaObjekPajak.Induk_Kecamatan_Kode = &resultNopBersama[2]
		dataAnggotaObjekPajak.Induk_Kelurahan_Kode = &resultNopBersama[3]
		dataAnggotaObjekPajak.Induk_Blok_Kode = &resultNopBersama[4]
		dataAnggotaObjekPajak.Induk_NoUrut = &resultNopBersama[5]
		dataAnggotaObjekPajak.Induk_JenisOp = &resultNopBersama[6]
		dataAnggotaObjekPajak.Induk_Area_Kode = &kodeBersama

		dataAnggotaObjekPajak.Provinsi_Kode = &resultNop[0]
		dataAnggotaObjekPajak.Daerah_Kode = &resultNop[1]
		dataAnggotaObjekPajak.Kecamatan_Kode = &resultNop[2]
		dataAnggotaObjekPajak.Kelurahan_Kode = &resultNop[3]
		dataAnggotaObjekPajak.Blok_Kode = &resultNop[4]
		dataAnggotaObjekPajak.NoUrut = &resultNop[5]
		dataAnggotaObjekPajak.JenisOp = &resultNop[6]
		dataAnggotaObjekPajak.Area_Kode = &kode
	}

	if input.RegKunjunganKembalis != nil {
		// copy data kunjungan kembali
		if err := sc.Copy(&dataKunjunganKembali, input.RegKunjunganKembalis); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg kunjungan kembali", input.RegKunjunganKembalis)
		}
		resultNopKK, kode := sh.NopParser(*input.Nop)
		dataKunjunganKembali.NopDetailCreateDto.Provinsi_Kode = &resultNopKK[0]
		dataKunjunganKembali.NopDetailCreateDto.Daerah_Kode = &resultNopKK[1]
		dataKunjunganKembali.NopDetailCreateDto.Kecamatan_Kode = &resultNopKK[2]
		dataKunjunganKembali.NopDetailCreateDto.Kelurahan_Kode = &resultNopKK[3]
		dataKunjunganKembali.NopDetailCreateDto.Blok_Kode = &resultNopKK[4]
		dataKunjunganKembali.NopDetailCreateDto.NoUrut = &resultNopKK[5]
		dataKunjunganKembali.NopDetailCreateDto.JenisOp = &resultNopKK[6]
		dataKunjunganKembali.Area_Kode = &kode
	}

	// add static field for reg objek pajak bumi
	dataObjekPajakBumi.Provinsi_Kode = &resultNop[0]
	dataObjekPajakBumi.Daerah_Kode = &resultNop[1]
	dataObjekPajakBumi.Kecamatan_Kode = &resultNop[2]
	dataObjekPajakBumi.Kelurahan_Kode = &resultNop[3]
	dataObjekPajakBumi.Blok_Kode = &resultNop[4]
	dataObjekPajakBumi.NoUrut = &resultNop[5]
	dataObjekPajakBumi.JenisOp = &resultNop[6]
	dataObjekPajakBumi.Area_Kode = &kode

	// add static field for objek pajak pbb
	data.NopDetail.Provinsi_Kode = &resultNop[0]
	data.NopDetail.Daerah_Kode = &resultNop[1]
	data.NopDetail.Kecamatan_Kode = &resultNop[2]
	data.NopDetail.Kelurahan_Kode = &resultNop[3]
	data.NopDetail.Blok_Kode = &resultNop[4]
	data.NopDetail.NoUrut = &resultNop[5]
	data.NopDetail.JenisOp = &resultNop[6]
	data.NopDetail.Area_Kode = &kode

	if input.TanggalPemeriksaan != nil {

		data.TanggalPemeriksaan = th.ParseTime(input.TanggalPemeriksaan)
	}
	if input.TanggalPendataan != nil {

		data.TanggalPendataan = th.ParseTime(input.TanggalPendataan)
	}
	if input.TanggalPerekaman != nil {

		data.TanggalPerekaman = th.ParseTime(input.TanggalPerekaman)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// create data wajibpajakpbb
		resultWajibPajakPbb, err := srwp.Create(dataWajibPajakPbb, tx)
		if err != nil {
			return err
		}
		respDataWajibPajakPbb = resultWajibPajakPbb
		resultCastWajibPajakPbb := resultWajibPajakPbb.(rp.OKSimple).Data.(mrwp.RegWajibPajakPbb)
		// add static value
		data.RegWajibPajakPbb_Id = &resultCastWajibPajakPbb.Id

		// create data objekpajakpbb
		err = tx.Create(&data).Error
		if err != nil {
			return err
		}

		// create data objek pajak bumi
		resultObjekPajakBumi, err := sropb.Create(dataObjekPajakBumi, tx)
		if err != nil {
			return err
		}
		respDataObjekPajakBumi = resultObjekPajakBumi

		if input.NopBersama != nil {
			// create data anggota objek pajak
			resultAnggotaObjekPajak, err := sraop.Create(dataAnggotaObjekPajak, tx)
			if err != nil {
				return err
			}
			respDataAnggotaObjekPajak = resultAnggotaObjekPajak
		}

		if input.RegKunjunganKembalis != nil {
			// create data kunjungan kembali
			resultKunjunganKembali, err := srkk.Create(dataKunjunganKembali, tx)
			if err != nil {
				return err
			}
			respDataKunjunganKembali = resultKunjunganKembali
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	if input.NopBersama != nil && input.RegKunjunganKembalis != nil {
		resp = t.II{
			"regObjekPajakPbb":     data,
			"regWajibPajakPbb":     respDataWajibPajakPbb.(rp.OKSimple).Data,
			"regObjekPajakBumi":    respDataObjekPajakBumi.(rp.OKSimple).Data,
			"regAnggotaObjekPajak": respDataAnggotaObjekPajak.(rp.OKSimple).Data,
			"regKunjunganKembali":  respDataKunjunganKembali.(rp.OKSimple).Data,
		}
	} else if input.NopBersama != nil && input.RegKunjunganKembalis == nil {
		resp = t.II{
			"regObjekPajakPbb":     data,
			"regWajibPajakPbb":     respDataWajibPajakPbb.(rp.OKSimple).Data,
			"regObjekPajakBumi":    respDataObjekPajakBumi.(rp.OKSimple).Data,
			"regAnggotaObjekPajak": respDataAnggotaObjekPajak.(rp.OKSimple).Data,
		}
	} else if input.NopBersama == nil && input.RegKunjunganKembalis != nil {
		resp = t.II{
			"regObjekPajakPbb":    data,
			"regWajibPajakPbb":    respDataWajibPajakPbb.(rp.OKSimple).Data,
			"regObjekPajakBumi":   respDataObjekPajakBumi.(rp.OKSimple).Data,
			"regKunjunganKembali": respDataKunjunganKembali.(rp.OKSimple).Data,
		}
	} else if input.NopBersama == nil && input.RegKunjunganKembalis == nil {
		resp = t.II{
			"regObjekPajakPbb":  data,
			"regWajibPajakPbb":  respDataWajibPajakPbb.(rp.OKSimple).Data,
			"regObjekPajakBumi": respDataObjekPajakBumi.(rp.OKSimple).Data,
		}
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.RegObjekPajakPbb
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.RegObjekPajakPbb{}).
		Preload(clause.Associations).
		Preload("Kelurahan.Kecamatan.Daerah.Provinsi").
		Preload("RegWajibPajakPbb.Kelurahan").
		Preload("RegWajibPajakPbb.Daerah").
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

func GetDetail(id int) (any, error) {
	var data *m.RegObjekPajakPbb

	result := a.DB.
		Model(&m.RegObjekPajakPbb{}).
		Preload(clause.Associations).
		Preload("Kelurahan.Kecamatan.Daerah.Provinsi").
		Preload("RegWajibPajakPbb.Kelurahan").
		Preload("RegWajibPajakPbb.Daerah").
		First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

// func GetDetailbyField(field string, value string) (any, error) {
// 	var data *m.ObjekPajakPbb

// 	result := a.DB.Where(field, value).First(&data)
// 	if result.RowsAffected == 0 {
// 		return nil, nil
// 	} else if result.Error != nil {
// 		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
// 	}

// 	return rp.OKSimple{
// 		Data: data,
// 	}, nil
// }

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.RegObjekPajakPbb
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.RegObjekPajakPbb
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	result = tx.Delete(&data, id)
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

func VerifySpop(id int, input m.VerifyDto) (any, error) {
	if input.Status > 2 {
		return nil, errors.New("status tidak diketahui")
	}

	var data *m.RegObjekPajakPbb
	var dataObjekPajakPbb mopp.ObjekPajakPbb
	var dataWajibPajakPbb mwpp.CreateDto
	var dataObjekPajakBumi mopb.CreateDto
	var dataKunjunganKembali mkk.CreateDto
	var dataAnggotaObjekPajak maop.CreateDto
	var dataOpb *mopb.ObjekPajakBumi
	// var respDataWajibPajakPbb interface{}
	var respDataKunjunganKembali interface{}
	var respDataAnggotaObjekPajak interface{}
	var respDataObjekPajakBumi interface{}
	var resp t.II

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data registrasi spop tidak dapat ditemukan")
	}
	if data.Status != m.StatusBaru {
		if data.Status == m.StatusDisetujui {
			return nil, errors.New("data sudah disetujui")
		} else if data.Status == m.StatusDitolak {
			return nil, errors.New("data sudah ditolak")
		}
	}

	if input.Status == m.StatusDitolak {
		data.Status = m.StatusDitolak
		if result := a.DB.Save(&data); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data registrasi spop", data)
		}
		return rp.OK{
			Meta: t.IS{
				"affected": strconv.Itoa(int(result.RowsAffected)),
			},
			Data: data,
		}, nil
	}

	// copy data reg objek pajak pbb
	if err := sc.Copy(&dataObjekPajakPbb, &data); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data reg objek pajak pbb: "+err.Error(), data)
	}

	// find data reg wajibpajakpbb
	var dataRegWp *mrwp.RegWajibPajakPbb
	err := a.DB.Where(mrwp.RegWajibPajakPbb{Id: *data.RegWajibPajakPbb_Id}).First(&dataRegWp).Error
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data reg wajib pajak pbb", dataRegWp)
	}

	// check wp pbb exist or not
	var tmpWpPbb mwpp.WajibPajakPbb
	resultWp := a.DB.Where(mwpp.WajibPajakPbb{Nik: dataRegWp.Nik}).First(&tmpWpPbb)
	if resultWp.RowsAffected == 0 {
		// copy data reg wajib pajak
		if err := sc.Copy(&dataWajibPajakPbb, &dataRegWp); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data reg wajib pajak pbb", dataRegWp)
		}
	}

	// reg objek pajak bumi condition for query
	conditionROpb := mropb.RegObjekPajakBumi{
		NopDetail: nop.NopDetail{
			Provinsi_Kode:  data.Provinsi_Kode,
			Daerah_Kode:    data.Daerah_Kode,
			Kecamatan_Kode: data.Kecamatan_Kode,
			Kelurahan_Kode: data.Kelurahan_Kode,
			Blok_Kode:      data.Blok_Kode,
			NoUrut:         data.NoUrut,
			JenisOp:        data.JenisOp,
		}}
	// find data reg objek pajak bumi
	var dataRegOpb *mropb.RegObjekPajakBumi
	resultRopb := a.DB.Where(conditionROpb).First(&dataRegOpb)

	// if resultRopb.Error != nil {
	// 	return sh.SetError("request", "create-data", source, "failed", "gagal menyalin data reg objek pajak bumi", dataRegWp)
	// }
	if resultRopb.RowsAffected != 0 {
		// copy data reg objek pajak bumi
		if err := sc.Copy(&dataObjekPajakBumi, &dataRegOpb); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal menyalin data reg objek pajak bumi", dataRegWp)
		}

	}

	// reg kunjungan kembali condition for query
	conditionRKk := mrkk.RegKunjunganKembali{
		NopDetail: nop.NopDetail{
			Provinsi_Kode:  data.Provinsi_Kode,
			Daerah_Kode:    data.Daerah_Kode,
			Kecamatan_Kode: data.Kecamatan_Kode,
			Kelurahan_Kode: data.Kelurahan_Kode,
			Blok_Kode:      data.Blok_Kode,
			NoUrut:         data.NoUrut,
			JenisOp:        data.JenisOp,
		}}

	// find data reg kunjungan kembali
	var dataRegKk *mrkk.RegKunjunganKembali
	resultKk := a.DB.Where(conditionRKk).First(&dataRegKk)
	// if resultKk.Error != nil {
	// 	return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data reg kunjungan kembali", dataRegKk)
	// }

	if resultKk.RowsAffected != 0 {
		// copy data reg kunjungan kembali
		if err := sc.Copy(&dataKunjunganKembali, &dataRegKk); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data reg kunjungan kembali", dataRegKk)
		}
	}

	// reg anggota objek pajak condition for query
	conditionRAop := mraop.RegAnggotaObjekPajak{
		Provinsi_Kode:  data.Provinsi_Kode,
		Daerah_Kode:    data.Daerah_Kode,
		Kecamatan_Kode: data.Kecamatan_Kode,
		Kelurahan_Kode: data.Kelurahan_Kode,
		Blok_Kode:      data.Blok_Kode,
		NoUrut:         data.NoUrut,
		JenisOp:        data.JenisOp,
	}

	// find data anggota objek pajak
	var dataRegAop *mraop.RegAnggotaObjekPajak
	resultAop := a.DB.Where(conditionRAop).First(&dataRegAop)
	// if resultAop.Error != nil {
	// 	return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data reg anggota objek pajak", dataRegAop)
	// }

	if resultAop.RowsAffected != 0 {
		// copy data reg anggota objek pajak
		if err := sc.Copy(&dataAnggotaObjekPajak, &dataRegKk); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data reg anggota objek pajak", dataRegAop)
		}
	}

	err = a.DB.Transaction(func(tx *gorm.DB) error {

		if input.Status == m.StatusDisetujui {
			data.Status = m.StatusDisetujui
			if result := tx.Save(&data); result.Error != nil {
				return result.Error
			}
		}
		if tmpWpPbb.Id == 0 {
			// create data wajibpajakpbb
			resultWajibPajakPbb, err := swp.Create(dataWajibPajakPbb, tx)
			if err != nil {
				return err
			}
			// respDataWajibPajakPbb = resultWajibPajakPbb
			resultCastWajibPajakPbb := resultWajibPajakPbb.(rp.OKSimple).Data.(mwpp.WajibPajakPbb)
			// add static value
			dataObjekPajakPbb.WajibPajakPbb_Id = &resultCastWajibPajakPbb.Id
		} else {
			dataObjekPajakPbb.WajibPajakPbb_Id = &tmpWpPbb.Id
		}

		// create data objekpajakpbb
		dataObjekPajakPbb.Id = 0
		err = tx.Create(&dataObjekPajakPbb).Error
		if err != nil {
			return err
		}

		// create data kunjungan kembali
		if resultKk.RowsAffected != 0 {
			// dataKunjunganKembali.Id = 0
			resultKunjunganKembali, err := skk.Create(dataKunjunganKembali, tx)
			if err != nil {
				return err
			}
			respDataKunjunganKembali = resultKunjunganKembali
		}

		// create data anggota objek pajak
		if resultAop.RowsAffected != 0 {
			// dataAnggotaObjekPajak.Id = 0
			resultAnggotaObjekPajak, err := saop.Create(dataAnggotaObjekPajak, tx)
			if err != nil {
				return err
			}
			respDataAnggotaObjekPajak = resultAnggotaObjekPajak
		}

		// cek data objek pajak bumi jika exist dgn nop condition maka luas bumi exist dikurangi luas bumi input
		if data.NopAsal != nil {
			resultNopAsal, _ := sh.NopParser(*data.NopAsal)
			conditionOpb := mopb.ObjekPajakBumi{
				NopDetail: nop.NopDetail{
					Provinsi_Kode:  &resultNopAsal[0],
					Daerah_Kode:    &resultNopAsal[1],
					Kecamatan_Kode: &resultNopAsal[2],
					Kelurahan_Kode: &resultNopAsal[3],
					Blok_Kode:      &resultNopAsal[4],
					NoUrut:         &resultNopAsal[5],
					JenisOp:        &resultNopAsal[6],
				}}

			resultOpb := a.DB.Where(conditionOpb).First(&dataOpb)
			// if resultOpb.Error != nil {
			// 	return resultOpb.Error
			// }

			if resultOpb.RowsAffected != 0 {
				// copy data objek pajak bumi
				dataOpb.LuasBumi -= dataRegOpb.LuasBumi
				if result := a.DB.Save(&dataOpb); result.Error != nil {
					return result.Error
				}
			} else if resultOpb.RowsAffected == 0 {
				// create data objek pajak bumi
				if resultOpb.RowsAffected != 0 {
					// dataObjekPajakBumi.Id = 0
					resultObjekPajakBumi, err := sopb.Create(dataObjekPajakBumi, tx)
					if err != nil {
						return err
					}
					respDataObjekPajakBumi = resultObjekPajakBumi
				}
			}
		} else {
			// create data objek pajak bumi
			// dataObjekPajakBumi.Id = 0
			resultObjekPajakBumi, err := sopb.Create(dataObjekPajakBumi, tx)
			if err != nil {
				return err
			}
			respDataObjekPajakBumi = resultObjekPajakBumi

		}

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	if resultKk.RowsAffected != 0 && resultAop.RowsAffected != 0 && data.NopAsal == nil {
		resp = t.II{
			"objekPajakPbb": dataObjekPajakPbb,
			// "wajibPajakPbb":     respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":    respDataObjekPajakBumi.(rp.OKSimple).Data,
			"kunjunganKembali":  respDataKunjunganKembali.(rp.OKSimple).Data,
			"anggotaObjekPajak": respDataAnggotaObjekPajak.(rp.OKSimple).Data,
		}
	} else if resultKk.RowsAffected == 0 && resultAop.RowsAffected != 0 && data.NopAsal == nil {
		resp = t.II{
			"objekPajakPbb": dataObjekPajakPbb,
			// "wajibPajakPbb":     respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":    respDataObjekPajakBumi.(rp.OKSimple).Data,
			"anggotaObjekPajak": respDataAnggotaObjekPajak.(rp.OKSimple).Data,
		}
	} else if resultKk.RowsAffected != 0 && resultAop.RowsAffected == 0 && data.NopAsal == nil {
		resp = t.II{
			"objekPajakPbb": dataObjekPajakPbb,
			// "wajibPajakPbb":    respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":   respDataObjekPajakBumi.(rp.OKSimple).Data,
			"kunjunganKembali": respDataKunjunganKembali.(rp.OKSimple).Data,
		}
	} else if resultKk.RowsAffected == 0 && resultAop.RowsAffected == 0 && data.NopAsal == nil {
		resp = t.II{
			"objekPajakPbb": dataObjekPajakPbb,
			// "wajibPajakPbb":  respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi": respDataObjekPajakBumi.(rp.OKSimple).Data,
		}
	} else if resultKk.RowsAffected != 0 && resultAop.RowsAffected != 0 && data.NopAsal != nil {
		resp = t.II{
			"objekPajakPbb": dataObjekPajakPbb,
			// "wajibPajakPbb":     respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":    dataOpb,
			"kunjunganKembali":  respDataKunjunganKembali.(rp.OKSimple).Data,
			"anggotaObjekPajak": respDataAnggotaObjekPajak.(rp.OKSimple).Data,
		}
	} else if resultKk.RowsAffected == 0 && resultAop.RowsAffected != 0 && data.NopAsal != nil {
		resp = t.II{
			"objekPajakPbb": dataObjekPajakPbb,
			// "wajibPajakPbb":     respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":    dataOpb,
			"anggotaObjekPajak": respDataAnggotaObjekPajak.(rp.OKSimple).Data,
		}
	} else if resultKk.RowsAffected != 0 && resultAop.RowsAffected == 0 && data.NopAsal != nil {
		resp = t.II{
			"objekPajakPbb": dataObjekPajakPbb,
			// "wajibPajakPbb":    respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":   dataOpb,
			"kunjunganKembali": respDataKunjunganKembali.(rp.OKSimple).Data,
		}
	} else if resultKk.RowsAffected == 0 && resultAop.RowsAffected == 0 && data.NopAsal != nil {
		resp = t.II{
			"objekPajakPbb": dataObjekPajakPbb,
			// "wajibPajakPbb":  respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi": dataOpb,
		}
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}
