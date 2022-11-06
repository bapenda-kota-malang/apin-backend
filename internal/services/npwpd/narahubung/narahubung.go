package npwpd

import (
	"errors"
	"fmt"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "regpemilik"

func Create(input []m.NarahubungCreateDto, npwpd_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.Narahubung

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// static add value to field
	for k := range data {
		data[k].Npwpd_Id = npwpd_id
	}

	// Transaction save to db
	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data %s", source), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(input []m.NarahubungUpdateDto, npwpd_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	var data []m.Narahubung

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), data)
	}

	for _, v := range data {
		var dataN *m.Narahubung
		result := tx.First(&dataN, v.Id)
		if result.RowsAffected == 0 {
			return nil, errors.New("data reg narahubung tidak dapat ditemukan")
		}
		if err := sc.Copy(&dataN, &v); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload narahubung", dataN)
		}

		dataN.Npwpd_Id = npwpd_id
		if result := tx.Save(&dataN); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data narahubung", dataN)
		}
	}
	for k := range data {
		data[k].Npwpd_Id = npwpd_id
	}
	return rp.OKSimple{Data: data}, nil
}

func Delete(npwpd_Id uint64, tx *gorm.DB) error {
	var dataNarahubung []*m.Narahubung
	result := tx.Where(m.Narahubung{Npwpd_Id: npwpd_Id}).Find(&dataNarahubung)
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
