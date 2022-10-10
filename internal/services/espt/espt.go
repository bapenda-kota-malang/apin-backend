package espt

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

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"

	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "espt"

func GetList(input m.FilterDto) (any, error) {
	var data []m.Espt
	var count int64

	var pagination gh.Pagination
	result := a.DB.Model(&m.Espt{}).
		// Joins("Npwpd").
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

	result := a.DB.Preload(clause.Associations).First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}
	if len(*data.DetailEsptAir) == 0 {
		data.DetailEsptAir = nil
	}
	if len(*data.DetailEsptHotel) == 0 {
		data.DetailEsptHotel = nil
	}
	if len(*data.DetailEsptParkir) == 0 {
		data.DetailEsptParkir = nil
	}
	if len(*data.DetailEsptReklame) == 0 {
		data.DetailEsptReklame = nil
	}
	if len(*data.DetailEsptResto) == 0 {
		data.DetailEsptResto = nil
	}
	return rp.OKSimple{
		Data: data,
	}, nil
}

// Service insert data for table espt:
//
//	Create(input, tx) // for using transaction
//	Create(input, nil) // for not using transaction
func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
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
	data.VerifyStatus = string(m.StatusBaru)

	// save to db
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

// Service update data for table espt
func Update(id int, input any, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.Espt

	// validate data exist and copy input (payload) ke struct data jika tidak ada akan error
	dataRow := tx.First(&data, id).RowsAffected
	if dataRow == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	// copy to model struct
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	// check relasi id tabel ke tabel relasi
	// potensiop table
	// if result := tx.First(&mr.Rekening{}, dataPotensiOp.Rekening_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }

	if _, ok := input.(m.VerifyDto); ok {
		// TODO: add verify by user id
		// data.VerifyBy_User_Id = 0
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(dataRow)),
		},
		Data: data,
	}, nil
}

// Service soft delete data for table espt
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
