package regnpwpd

import (
	"errors"
	"fmt"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/regnpwpd"

	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
)

const source = "regpemilik"

func Create(input []m.RegNarahubungCreateDto, regis_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.RegNarahubung

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// static add value to field
	for k := range data {
		data[k].RegNpwpd_Id = regis_id
		data[k].Status = nt.StatusBaru
	}

	// Transaction save to db
	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data %s", source), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(input []m.RegNarahubungUpdateDto, regNpwpd_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	var data []m.RegNarahubung

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), data)
	}

	for _, v := range data {
		var dataN *m.RegNarahubung
		result := tx.First(&dataN, v.Id)
		if result.RowsAffected == 0 {
			return nil, errors.New("data reg narahubung tidak dapat ditemukan")
		}
		if err := sc.Copy(&dataN, &v); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload reg narahubung", dataN)
		}

		dataN.RegNpwpd_Id = regNpwpd_id
		if result := tx.Save(&dataN); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data reg narahubung", dataN)
		}
	}
	for k := range data {
		data[k].RegNpwpd_Id = regNpwpd_id
	}
	return rp.OKSimple{Data: data}, nil
}

func Delete(regNpwpd_Id uint64, tx *gorm.DB) error {
	var dataNarahubung []*m.RegNarahubung
	result := tx.Where(m.RegNarahubung{RegNpwpd_Id: regNpwpd_Id}).Find(&dataNarahubung)
	if result.RowsAffected == 0 {
		return errors.New("data reg narahubung tidak dapat ditemukan")
	}

	for _, v := range dataNarahubung {
		result = tx.Delete(&v)
		if result.RowsAffected == 0 {
			return errors.New("tidak dapat menghapus data reg narahubung")
		}
	}
	return nil
}
