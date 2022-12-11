package objekpajakbangunan

import (
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
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	sfb "github.com/bapenda-kota-malang/apin-backend/internal/services/fasilitasbangunan"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "objekpajakbangunan"

func Create(input m.OpbJpb2CreateDto) (any, error) {

	var data m.ObjekPajakBangunan
	var dataFasilitasBangunan mfb.CreateDto
	var respDataFasilitasBangunan interface{}
	var respDataJpb interface{}
	var resp t.II

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	if input.FasilitasBangunans != nil {
		// copy data fasilitas bangunan
		if err := sc.Copy(&dataFasilitasBangunan, input.FasilitasBangunans); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload fasilitas bangunan", input.FasilitasBangunans)
		}
		resultNop, kode := sh.NopParser(*input.Nop)
		dataFasilitasBangunan.NopDetailCreateDto.Provinsi_Kode = &resultNop[0]
		dataFasilitasBangunan.NopDetailCreateDto.Daerah_Kode = &resultNop[1]
		dataFasilitasBangunan.NopDetailCreateDto.Kecamatan_Kode = &resultNop[2]
		dataFasilitasBangunan.NopDetailCreateDto.Kelurahan_Kode = &resultNop[3]
		dataFasilitasBangunan.NopDetailCreateDto.Blok_Kode = &resultNop[4]
		dataFasilitasBangunan.NopDetailCreateDto.NoUrut = &resultNop[5]
		dataFasilitasBangunan.NopDetailCreateDto.JenisOp = &resultNop[6]
		dataFasilitasBangunan.NopDetailCreateDto.Area_Kode = &kode

	}

	// add static field for objek pajak bangunan
	resultNop, kode := sh.NopParser(*input.Nop)
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
		respDataJpb = resultJpb

		if input.FasilitasBangunans != nil {
			// create data fasilitas bangunan
			resultFasilitasBangunan, err := sfb.Create(dataFasilitasBangunan, tx)
			if err != nil {
				return err
			}
			respDataFasilitasBangunan = resultFasilitasBangunan
		}

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}
	if input.FasilitasBangunans != nil {
		resp = t.II{
			"objekPajakBangunan": data,
			"fasilitasBangunan":  respDataFasilitasBangunan.(rp.OKSimple).Data,
			"jpb":                respDataJpb.(rp.OKSimple).Data,
		}
	} else {
		resp = t.II{
			"objekPajakBangunan": data,
			"jpb":                respDataJpb.(rp.OKSimple).Data,
		}
	}

	return rp.OKSimple{
		Data: resp,
	}, nil

}

func GetList(input m.FilterDto) (any, error) {
	var data []m.ObjekPajakBangunan
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.ObjekPajakBangunan{}).
		Preload(clause.Associations).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data objekpajakbangunan", data)
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
	var data *m.ObjekPajakBangunan

	result := a.DB.
		Model(&m.ObjekPajakBangunan{}).
		Preload(clause.Associations).
		First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data objekpajakbangunan", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.ObjekPajakBangunan
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data objekpajakbangunan", data)
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
	var data *m.ObjekPajakBangunan
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
