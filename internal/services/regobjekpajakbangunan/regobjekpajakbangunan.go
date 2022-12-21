package regobjekpajakbangunan

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
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"

	mfb "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	mopb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbangunan"
	sfb "github.com/bapenda-kota-malang/apin-backend/internal/services/fasilitasbangunan"
	srfb "github.com/bapenda-kota-malang/apin-backend/internal/services/regfasilitasbangunan"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "reg objekpajakbangunan"

func Create(input m.RegInput) (any, error) {

	var data m.RegObjekPajakBangunan
	var dataRegFasilitasBangunan mrfb.CreateDto
	var respDataRegFasilitasBangunan interface{}
	var respDataRegJpb interface{}
	var resp t.II

	resultGetFasilitasBangunan := input.GetFasilitasBangunan()
	resultGetNop := input.GetNop()
	resultGetPemeriksaan := input.GetTanggalPemeriksaan()
	resultGetPendataan := input.GetTanggalPendataan()
	resultGetPerekaman := input.GetTanggalPerekaman()
	resultGetObjekPajakBangun, errGet := input.GetObjekPajakBangunan()
	if errGet != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}
	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &resultGetObjekPajakBangun); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	if resultGetFasilitasBangunan != nil {
		// copy data fasilitas bangunan
		if err := sc.Copy(&dataRegFasilitasBangunan, resultGetFasilitasBangunan); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload fasilitas bangunan", resultGetFasilitasBangunan)
		}
		resultNop, kode := sh.NopParser(*resultGetNop)
		dataRegFasilitasBangunan.NopDetailCreateDto.Provinsi_Kode = &resultNop[0]
		dataRegFasilitasBangunan.NopDetailCreateDto.Daerah_Kode = &resultNop[1]
		dataRegFasilitasBangunan.NopDetailCreateDto.Kecamatan_Kode = &resultNop[2]
		dataRegFasilitasBangunan.NopDetailCreateDto.Kelurahan_Kode = &resultNop[3]
		dataRegFasilitasBangunan.NopDetailCreateDto.Blok_Kode = &resultNop[4]
		dataRegFasilitasBangunan.NopDetailCreateDto.NoUrut = &resultNop[5]
		dataRegFasilitasBangunan.NopDetailCreateDto.JenisOp = &resultNop[6]
		dataRegFasilitasBangunan.NopDetailCreateDto.Area_Kode = &kode

	}

	// add static field for objek pajak bangunan
	resultNop, kode := sh.NopParser(*resultGetNop)
	data.NopDetail.Provinsi_Kode = &resultNop[0]
	data.NopDetail.Daerah_Kode = &resultNop[1]
	data.NopDetail.Kecamatan_Kode = &resultNop[2]
	data.NopDetail.Kelurahan_Kode = &resultNop[3]
	data.NopDetail.Blok_Kode = &resultNop[4]
	data.NopDetail.NoUrut = &resultNop[5]
	data.NopDetail.JenisOp = &resultNop[6]
	data.NopDetail.Area_Kode = &kode

	if resultGetPemeriksaan != nil {

		data.TanggalPemeriksaan = th.ParseTime(resultGetPemeriksaan)
	}
	if resultGetPendataan != nil {

		data.TanggalPendataan = th.ParseTime(resultGetPendataan)
	}
	if resultGetPerekaman != nil {

		data.TanggalPerekaman = th.ParseTime(resultGetPerekaman)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// create data objek pajak bangunan
		err := tx.Create(&data).Error
		if err != nil {
			return err
		}

		// copy data jpb
		resultJpb, err := jpbCopier(input, resultNop, kode, tx)
		if err != nil {
			return err
		}
		respDataRegJpb = resultJpb

		if resultGetFasilitasBangunan != nil {
			// create data fasilitas bangunan
			resultFasilitasBangunan, err := srfb.Create(dataRegFasilitasBangunan, tx)
			if err != nil {
				return err
			}
			respDataRegFasilitasBangunan = resultFasilitasBangunan
		}

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}
	if resultGetFasilitasBangunan != nil {
		resp = t.II{
			"regObjekPajakBangunan": data,
			"fasilitasBangunan":     respDataRegFasilitasBangunan.(rp.OKSimple).Data,
			"regJpb":                respDataRegJpb.(rp.OKSimple).Data,
		}
	} else {
		resp = t.II{
			"regObjekPajakBangunan": data,
			"regJpb":                respDataRegJpb.(rp.OKSimple).Data,
		}
	}

	return rp.OKSimple{
		Data: resp,
	}, nil

}

func GetList(input m.FilterDto) (any, error) {
	var data []m.RegObjekPajakBangunan
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.RegObjekPajakBangunan{}).
		Preload(clause.Associations).
		Preload("Kelurahan.Kecamatan.Daerah.Provinsi").
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data reg objekpajakbangunan", data)
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
	var data *m.RegObjekPajakBangunan

	result := a.DB.
		Model(&m.RegObjekPajakBangunan{}).
		Preload(clause.Associations).
		Preload("Kelurahan.Kecamatan.Daerah.Provinsi").
		First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data reg objekpajakbangunan", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.RegObjekPajakBangunan
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data reg objekpajakbangunan", data)
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
	var data *m.RegObjekPajakBangunan
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

func VerifyLspop(id int, input m.VerifyDto) (any, error) {
	if input.Status > 2 {
		return nil, errors.New("status tidak diketahui")
	}

	var data m.RegObjekPajakBangunan
	var dataObjekPajakBangunan mopb.ObjekPajakBangunan
	var dataFasilitasBangunan mfb.CreateDto
	var respFasilitasBangunan interface{}
	var respJpb interface{}
	var resp t.II

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data registrasi lspop tidak dapat ditemukan")
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
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data registrasi lspop", data)
		}
		return rp.OK{
			Meta: t.IS{
				"affected": strconv.Itoa(int(result.RowsAffected)),
			},
			Data: data,
		}, nil
	}

	// copy data reg objek pajak bangunan
	if err := sc.Copy(&dataObjekPajakBangunan, &data); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data reg objek pajak bangunan: "+err.Error(), data)
	}

	// reg fasilitas bangunan condition for query
	conditionRFB := mrfb.RegFasilitasBangunan{
		NopDetail: nop.NopDetail{
			Provinsi_Kode:  data.Provinsi_Kode,
			Daerah_Kode:    data.Daerah_Kode,
			Kecamatan_Kode: data.Kecamatan_Kode,
			Kelurahan_Kode: data.Kelurahan_Kode,
			Blok_Kode:      data.Blok_Kode,
			NoUrut:         data.NoUrut,
			JenisOp:        data.JenisOp,
		}}
	// find data reg fasilitas bangunan
	var dataRegFB *mrfb.RegFasilitasBangunan
	resultRFB := a.DB.Where(conditionRFB).First(&dataRegFB)

	if resultRFB.RowsAffected != 0 {
		// copy data reg fasilitas bumi
		if err := sc.Copy(&dataFasilitasBangunan, &dataRegFB); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal menyalin data reg fasilitas bangunan", dataRegFB)
		}

	}

	conditionJpb := nop.NopDetail{
		Provinsi_Kode:  data.Provinsi_Kode,
		Daerah_Kode:    data.Daerah_Kode,
		Kecamatan_Kode: data.Kecamatan_Kode,
		Kelurahan_Kode: data.Kelurahan_Kode,
		Blok_Kode:      data.Blok_Kode,
		NoUrut:         data.NoUrut,
		JenisOp:        data.JenisOp,
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {

		if input.Status == m.StatusDisetujui {
			data.Status = m.StatusDisetujui
			if result := tx.Save(&data); result.Error != nil {
				return result.Error
			}
		}

		// create data objek pajak bangunan
		dataObjekPajakBangunan.Id = 0
		err := tx.Create(&dataObjekPajakBangunan).Error
		if err != nil {
			return err
		}

		if dataRegFB.Provinsi_Kode != nil {
			resultFasilitasBangunan, errFB := sfb.Create(dataFasilitasBangunan, tx)
			if errFB != nil {
				return errFB
			}
			respFasilitasBangunan = resultFasilitasBangunan
		}

		resultJpb, errJpb := verifyJpb(data.Jpb_Kode, conditionJpb, tx)
		if errJpb != nil {
			return errJpb
		}
		respJpb = resultJpb
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	if dataRegFB.Provinsi_Kode != nil {
		resp = t.II{
			"objekPajakBangunan": dataObjekPajakBangunan,
			"fasilitasBangunan":  respFasilitasBangunan.(rp.OKSimple).Data,
			"jpb":                respJpb.(rp.OKSimple).Data,
		}
	} else {
		resp = t.II{
			"objekPajakBangunan": dataObjekPajakBangunan,
			"jpb":                respJpb.(rp.OKSimple).Data,
		}
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}
