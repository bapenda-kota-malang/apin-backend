package datapetaznt

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/datapetaznt"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "datapetaznt"

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.DataPetaZnt

	condition := m.DataPetaZnt{
		Provinsi_Kode:  input.Provinsi_Kode,
		Daerah_Kode:    input.Daerah_Kode,
		Kecamatan_Kode: input.Kecamatan_Kode,
		Kelurahan_Kode: input.Kelurahan_Kode,
		Blok_Kode:      input.Blok_Kode,
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
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data DataPetaZnt: "+err.Error(), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func CreateBulk(input m.CreateBulkDto) (any, error) {
	var datas []m.DataPetaZnt

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range input.Datas {
			if v.Id != nil {
				var data = new(m.DataPetaZnt)
				checkExist := tx.
					Where("Id", v.Id).
					First(&data)
				if checkExist.RowsAffected > 0 {
					respData, err := Update(*v.Id, m.CreateDto{
						Provinsi_Kode:  input.Provinsi_Kode,
						Daerah_Kode:    input.Daerah_Kode,
						Kecamatan_Kode: input.Kecamatan_Kode,
						Kelurahan_Kode: input.Kelurahan_Kode,
						Blok_Kode:      v.Blok_Kode,
						Znt_Kode:       v.Znt_Kode,
					}, tx)
					if err != nil {
						return err
					}
					datas = append(datas, respData.(rp.OK).Data.(m.DataPetaZnt))
				}
			} else {
				respData, err := Create(m.CreateDto{
					Provinsi_Kode:  input.Provinsi_Kode,
					Daerah_Kode:    input.Daerah_Kode,
					Kecamatan_Kode: input.Kecamatan_Kode,
					Kelurahan_Kode: input.Kelurahan_Kode,
					Blok_Kode:      v.Blok_Kode,
					Znt_Kode:       v.Znt_Kode,
				}, tx)
				if err != nil {
					return err
				}
				datas = append(datas, respData.(rp.OKSimple).Data.(m.DataPetaZnt))
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
	var data []m.DataPetaZnt
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.DataPetaZnt{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data DataPetaZnt", data)
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
	var data *m.DataPetaZnt

	result := a.DB.Preload(clause.Associations).First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data DataPetaZnt", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.DataPetaZnt
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data DataPetaZnt", data)
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
	var data *m.DataPetaZnt
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
