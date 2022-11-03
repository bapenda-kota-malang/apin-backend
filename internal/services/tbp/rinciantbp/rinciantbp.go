package tbp

import (
	"errors"
	"fmt"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/tbp"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "rincian tbp"

func Create(input m.RincianTbpCreateDto, tbp_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.RincianTbp
	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload rincian tbp", data)
	}

	// static add value to field
	data.Tbp_Id = &tbp_id

	// Transaction save to db
	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data %s", source), data)
	}
	return rp.OKSimple{Data: data}, nil
}

func Delete(tbp_id uint64, tx *gorm.DB) error {
	var data *m.RincianTbp
	result := tx.Where(m.RincianTbp{Tbp_Id: &tbp_id}).First(&data)
	if result.RowsAffected == 0 {
		return errors.New("data rincian tbp tidak dapat ditemukan")
	}

	result = tx.Delete(&data)
	if result.RowsAffected == 0 {
		return errors.New("tidak dapat menghapus data rincian tbp")
	}

	return nil
}

func Update(input m.RincianTbpUpdateDto, tbp_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	var data m.RincianTbp

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), data)
	}

	var dataFromDb *m.RincianTbp
	result := tx.Where(m.RincianTbp{Tbp_Id: &tbp_id}).First(&dataFromDb)
	if result.RowsAffected == 0 {
		return nil, errors.New("data reg pemilik tidak dapat ditemukan")
	}

	err := sc.Copy(&dataFromDb, &data)
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload rincian tbp", data)
	}

	dataFromDb.Tbp_Id = &tbp_id
	if result := tx.Save(&dataFromDb); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data rincian tbp", dataFromDb)
	}
	return rp.OKSimple{Data: data}, nil
}
