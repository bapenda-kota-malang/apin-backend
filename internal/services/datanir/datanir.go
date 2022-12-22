package datanir

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/datanir"
	d "github.com/bapenda-kota-malang/apin-backend/internal/models/dokument"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "datanir"

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.DataNir

	condition := m.DataNir{
		Provinsi_Kode:  input.Provinsi_Kode,
		Daerah_Kode:    input.Daerah_Kode,
		Kecamatan_Kode: input.Kecamatan_Kode,
		Kelurahan_Kode: input.Kelurahan_Kode,
		Tahun:          input.Tahun,
		NomerDokumen:   input.NomerDokumen,
		Znt_Kode:       input.Znt_Kode,
	}

	if err := tx.Where(condition).First(&data).Error; err != nil && err != gorm.ErrRecordNotFound {
		return sh.SetError("request", "create-data", source, "failed", "gagal mencari data existing: "+err.Error(), data)
	}

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload: "+err.Error(), data)
	}

	// if err := adh.ValidateAreaDivisionKode(data.Provinsi_Kode, data.Daerah_Kode, data.Kecamatan_Kode, data.Kelurahan_Kode); err != nil {
	// 	return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
	// }

	// simpan data ke db satu if karena result dipakai sekali, +error
	var err error
	if data.Id == 0 {
		err = tx.Create(&data).Error
	} else {
		err = tx.Save(&data).Error
	}
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data datanir: "+err.Error(), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func CreateBulk(input m.CreateBulkDto) (any, error) {
	var datas []m.DataNir
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
			if v.Znt_Kode == "XX" {
				s := "XS"
				v.Znt_Kode = s
			}
			var data = new(m.DataNir)
			checkExist := tx.First(&data, v.Id)
			if checkExist.Error != nil || checkExist.RowsAffected == 0 {
				respData, err := Create(m.CreateDto{
					Provinsi_Kode:  input.Provinsi_Kode,
					Daerah_Kode:    input.Daerah_Kode,
					Kecamatan_Kode: input.Kecamatan_Kode,
					Kelurahan_Kode: input.Kelurahan_Kode,
					Tahun:          input.Tahun,
					NomerDokumen:   input.NomerDokumen,
					Kpbb_Id:        input.Kpbb_Id,
					Kanwil_Id:      input.Kanwil_Id,
					JenisDokumen:   input.JenisDokumen,
					Znt_Kode:       v.Znt_Kode,
					Nir:            v.Nir,
				}, tx)
				if err != nil {
					return err
				}
				datas = append(datas, respData.(rp.OKSimple).Data.(m.DataNir))
			} else {
				respData, err := Update(v.Id, m.CreateDto{
					Provinsi_Kode:  input.Provinsi_Kode,
					Daerah_Kode:    input.Daerah_Kode,
					Kecamatan_Kode: input.Kecamatan_Kode,
					Kelurahan_Kode: input.Kelurahan_Kode,
					Tahun:          input.Tahun,
					NomerDokumen:   input.NomerDokumen,
					Kpbb_Id:        input.Kpbb_Id,
					Kanwil_Id:      input.Kanwil_Id,
					JenisDokumen:   input.JenisDokumen,
					Znt_Kode:       v.Znt_Kode,
					Nir:            v.Nir,
				}, tx)
				if err != nil {
					return err
				}
				datas = append(datas, respData.(rp.OK).Data.(m.DataNir))
			}
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal create bulk: "+err.Error(), datas)
	}

	return rp.OKSimple{Data: datas}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.DataNir
	var count int64
	var result *gorm.DB
	input.NomerDokumen = nil

	var pagination gh.Pagination
	result = a.DB.
		Model(&m.DataNir{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.RowsAffected == 0 {
		tempTahun, _ := strconv.Atoi(*input.Tahun)
		tempTahunS := strconv.Itoa(tempTahun - 1)
		input.Tahun = &tempTahunS
		result = a.DB.
			Model(&m.DataNir{}).
			Scopes(gh.Filter(input)).
			Count(&count).
			Scopes(gh.Paginate(input, &pagination)).
			Find(&data)
	}
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data datanir", data)
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
	var data *m.DataNir

	result := a.DB.Preload(clause.Associations).First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data datanir", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func GetDokByNoDok(no string) (any, error) {
	var data *d.Dokument

	result := a.DB.Where("NoDokumen", no).First(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data datanir", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.DataNir
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data datanir", data)
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
	var data *m.DataNir
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
