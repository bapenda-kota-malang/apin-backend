package objekpajakbumi

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	d "github.com/bapenda-kota-malang/apin-backend/internal/models/dokument"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	msppt "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "objekpajakbumi"

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.ObjekPajakBumi

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data objekpajakbumi", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.ObjekPajakBumi
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.ObjekPajakBumi{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data objekpajakbumi", data)
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
	var data *m.ObjekPajakBumi

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data objekpajakbumi", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.ObjekPajakBumi
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data objekpajakbumi", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func UpdateBulk(input m.UpdateBulkDto) (any, error) {
	var datas []m.ObjekPajakBumi
	var dok d.Dokument

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		resDok := tx.Where("NoDokumen", &input.NomerDokumen).First(&dok)
		if resDok.RowsAffected > 0 {
			_, errTemp := sh.SetError("request", "get-data-dok", source, "failed", "no dokumen telah digunakan", input)
			return errTemp
		}

		errDok := tx.Create(&d.Dokument{
			Kanwil_Kd:      input.Kanwil_Id,
			Kppbb_Kd:       input.Kpbb_Id,
			JenisDokumen:   input.JenisDokumen,
			NoDokumen:      &input.NomerDokumen,
			TglPendataan:   input.TglPendataan,
			NipPendataan:   input.NipPendataan,
			TglPemeriksaan: input.TglPemeriksaan,
			NipPemeriksaan: input.NipPemeriksaan,
		}).Error
		if errDok != nil {
			return errDok
		}

		for _, v := range input.Datas {
			var data = new(m.ObjekPajakBumi)
			checkExist := tx.
				Where("Provinsi_Kode", input.Provinsi_Kode).
				Where("Daerah_Kode", input.Daerah_Kode).
				Where("Kecamatan_Kode", input.Kecamatan_Kode).
				Where("Kelurahan_Kode", input.Kelurahan_Kode).
				Where("Blok_Kode", input.Blok_Kode).
				Where("NoUrut", v.NoUrut).
				Where("JnsOp", v.JenisOp).
				Where("KodeZNT", input.KodeZNT).
				First(&data)
			if checkExist.Error != nil || checkExist.RowsAffected == 0 {
				return checkExist.Error
			} else {
				var cdto m.UpdateDto
				if err := sc.Copy(&cdto, &data); err != nil {
					return err
				}
				cdto.KodeZNT = input.KodeZNTBaru
				respData, err := Update(v.Id, cdto, tx)
				if err != nil {
					return err
				}
				datas = append(datas, respData.(rp.OK).Data.(m.ObjekPajakBumi))
			}
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal create bulk: "+err.Error(), datas)
	}

	return rp.OKSimple{Data: datas}, nil
}

func Delete(id int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.ObjekPajakBumi
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

// func PerubahanZnt(nop , kodeZnt string) (any, error) {

// }

func PenilaianSppt(input []msppt.Sppt, tx *gorm.DB) (any, error) {
	var data *m.ObjekPajakBumi
	for _, v := range input {
		condition := nopSearcher(v)
		result := tx.Where(condition).First(&data)
		// if result.Error != nil {
		// 	return sh.SetError("request", "penilaian-data", source, "failed", "gagal mengambil data", data)
		// }
		if result.RowsAffected != 0 {
			// TODO: nilai sistem
			if resultSave := tx.Save(&data); resultSave.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data objek pajak pbb", data)
			}
		}

	}
	return rp.OKSimple{Data: data}, nil
}
