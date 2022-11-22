package jenisusaha

import (
	"errors"
	"fmt"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/jenisusaha"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "jenis usaha detail"

func Create(input m.JenisUsahaDetailCreateDto, jenisUsaha_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.JenisUsahaDetail
	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jenis usaha detail", data)
	}

	// static add value to field
	data.JenisUsaha_Id = jenisUsaha_id

	// Transaction save to db
	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data %s", source), data)
	}
	return rp.OKSimple{Data: data}, nil
}

func Delete(jenisUsaha_id uint64, tx *gorm.DB) error {
	var data *m.JenisUsahaDetail
	result := tx.Where(m.JenisUsahaDetail{JenisUsaha_Id: jenisUsaha_id}).First(&data)
	if result.RowsAffected == 0 {
		return errors.New("data jenis usaha detail tidak dapat ditemukan")
	}

	result = tx.Delete(&data)
	if result.RowsAffected == 0 {
		return errors.New("tidak dapat menghapus data jenis usaha detail")
	}

	return nil
}

func Update(input m.JenisUsahaDetailUpdateDto, jenisUsaha_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	var data m.JenisUsahaDetail

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), data)
	}

	var dataFromDb *m.JenisUsahaDetail
	result := tx.Where(m.JenisUsahaDetail{JenisUsaha_Id: jenisUsaha_id}).First(&dataFromDb)
	if result.RowsAffected == 0 {
		return nil, errors.New("data jenis usaha detail tidak dapat ditemukan")
	}

	err := sc.Copy(&dataFromDb, &data)
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload jenis usaha detail", data)
	}

	dataFromDb.JenisUsaha_Id = jenisUsaha_id
	if result := tx.Save(&dataFromDb); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data jenis usaha detail", dataFromDb)
	}
	return rp.OKSimple{Data: dataFromDb}, nil
}
