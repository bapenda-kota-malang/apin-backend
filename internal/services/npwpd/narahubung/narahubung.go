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

	var dataUpdate []m.Narahubung
	var dataDelete []m.Narahubung
	var tmpDataUpdate m.Narahubung
	var tmpDataDelete m.Narahubung

	for _, v := range input {
		if v.Id != 0 && v.IsDeleted {
			if err := sc.Copy(&tmpDataDelete, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), dataDelete)
			}
			dataDelete = append(dataDelete, tmpDataDelete)
		} else {
			if err := sc.Copy(&tmpDataUpdate, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), dataUpdate)
			}
			dataUpdate = append(dataUpdate, tmpDataUpdate)
		}
	}

	// update data
	for _, v := range dataUpdate {
		var dataNU *m.Narahubung
		result := tx.First(&dataNU, v.Id)
		if result.RowsAffected == 0 {
			return nil, errors.New("data pemilik tidak dapat ditemukan")
		}
		if err := sc.Copy(&dataNU, &v); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload narahubung", dataNU)
		}
		// set directur value to null if golongan orang pribadi
		dataNU.Npwpd_Id = npwpd_id
		if result := tx.Save(&dataNU); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data narahubung", dataNU)
		}
	}
	for k := range dataUpdate {
		dataUpdate[k].Npwpd_Id = npwpd_id
	}

	// delete data
	for _, v := range dataDelete {
		var dataND *m.Narahubung
		result := tx.First(&dataND, v.Id)
		if result.RowsAffected == 0 {
			return nil, errors.New("data pemilik tidak dapat ditemukan")
		}

		result = tx.Delete(&dataND)
		if result.RowsAffected == 0 {
			return nil, errors.New("tidak dapat menghapus data reg narahubung")
		}
	}
	return rp.OKSimple{Data: dataUpdate}, nil
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
