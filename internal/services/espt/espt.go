package espt

import (
	"errors"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"

	sair "github.com/bapenda-kota-malang/apin-backend/internal/services/detailesptair"

	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "espt"

func GetList(input m.FilterDto) (any, error) {
	var data []m.Espt
	var count int64

	var pagination gh.Pagination
	result := a.DB.Model(&m.Espt{}).
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
	var data *m.Espt

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

func Create(input m.CreateDto, inputDetail []interface{}, category string) (any, error) {
	var data m.Espt

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// check relasi id tabel ke tabel relasi
	// espt table
	// if result := a.DB.First(&mr.Rekening{}, dataPotensiOp.Rekening_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }

	// static add value to field

	// Transaction save to db
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data ke db satu if karena result dipakai sekali, +error
		if result := tx.Create(&data); result.Error != nil {
			return result.Error
		}
		switch category {
		case "air":
			var listDataDetail []detailesptair.CreateDetailAirDto
			for _, v := range inputDetail {
				dataDetail := v.(detailesptair.CreateDetailAirDto)
				dataDetail.Espt_Id = data.Id
				listDataDetail = append(listDataDetail, dataDetail)
			}
			_, err := sair.Create(listDataDetail, tx)
			if err != nil {
				return err
			}
		case "hotel":

		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(id int, input m.UpdateDto) (any, error) {
	var data m.Espt

	// validate data exist and copy input (payload) ke struct data jika tidak ada akan error
	dataRow := a.DB.First(&data, id).RowsAffected
	if dataRow == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	// check relasi id tabel ke tabel relasi
	// potensiop table
	// if result := a.DB.First(&mr.Rekening{}, dataPotensiOp.Rekening_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(dataRow)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (any, error) {
	var data *m.Espt
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
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
