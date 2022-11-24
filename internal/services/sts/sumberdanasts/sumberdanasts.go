package sts

import (
	"errors"
	"fmt"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sts"
)

const source = "sumber dana sts"

func Create(input []m.SumberDanaStsCreateDto, sts_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.SumberDanaSts

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// static add value to field
	for k := range data {
		data[k].Sts_Id = sts_id
	}

	// Transaction save to db
	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data %s", source), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(input []m.SumberDanaStsUpdateDto, sts_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	var data []m.SumberDanaSts

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), data)
	}

	for _, v := range data {
		var dataSDS *m.SumberDanaSts
		result := tx.First(&dataSDS, v.Id)
		if result.RowsAffected == 0 {
			return nil, errors.New("data sumber dana sts tidak dapat ditemukan")
		}
		if err := sc.Copy(&dataSDS, &v); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload sumber dana sts", dataSDS)
		}

		dataSDS.Sts_Id = sts_id
		if result := tx.Save(&dataSDS); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data sumber dana sts", dataSDS)
		}
	}
	for k := range data {
		data[k].Sts_Id = sts_id
	}
	return rp.OKSimple{Data: data}, nil
}

func Delete(sts_Id uint64, tx *gorm.DB) error {
	var dataSDS []*m.SumberDanaSts
	result := tx.Where(m.SumberDanaSts{Sts_Id: sts_Id}).Find(&dataSDS)
	if result.RowsAffected == 0 {
		return errors.New("data sumber dana sts tidak dapat ditemukan")
	}

	for _, v := range dataSDS {
		result = tx.Delete(&v)
		if result.RowsAffected == 0 {
			return errors.New("tidak dapat menghapus data sumber dana sts")
		}
	}
	return nil
}
