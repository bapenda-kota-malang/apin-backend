package sspd

import (
	"errors"
	"fmt"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "detail sspd"

func Create(input m.SspdDetailCreateDto, sspd_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.SspdDetail
	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload sspd detail", data)
	}

	// static add value to field
	data.Sspd_Id = &sspd_id

	// Transaction save to db
	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data %s", source), data)
	}
	return rp.OKSimple{Data: data}, nil
}

func Delete(sspd_id uint64, tx *gorm.DB) error {
	var data *m.SspdDetail
	result := tx.Where(m.SspdDetail{Sspd_Id: &sspd_id}).First(&data)
	if result.RowsAffected == 0 {
		return errors.New("data sspd detail tidak dapat ditemukan")
	}

	result = tx.Delete(&data)
	if result.RowsAffected == 0 {
		return errors.New("tidak dapat menghapus data sspd detail")
	}

	return nil
}

func Update(input m.SspdDetailUpdateDto, sspd_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	var data m.SspdDetail

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), data)
	}

	var dataFromDb *m.SspdDetail
	result := tx.Where(m.SspdDetail{Sspd_Id: &sspd_id}).First(&dataFromDb)
	if result.RowsAffected == 0 {
		return nil, errors.New("data sspd detail tidak dapat ditemukan")
	}

	err := sc.Copy(&dataFromDb, &data)
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload sspd detail", data)
	}

	dataFromDb.Sspd_Id = &sspd_id
	if result := tx.Save(&dataFromDb); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data sspd detail", dataFromDb)
	}
	return rp.OKSimple{Data: data}, nil
}
