package njoptkpflag

import (
	"errors"
	"fmt"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/njoptkpflag"
)

const source = "njoptkp flag detail"

func Create(input []m.NjoptkpFlagDetailCreateDto, njoptkpFlag_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.NjoptkpFlagDetail

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// static add value to field
	for k := range data {
		data[k].NjoptkpFlag_Id = njoptkpFlag_id
	}

	// Transaction save to db
	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data %s", source), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(input []m.NjoptkpFlagDetailUpdateDto, njoptkpFlag_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	var data []m.NjoptkpFlagDetail

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), data)
	}

	for _, v := range data {
		var dataN *m.NjoptkpFlagDetail
		result := tx.First(&dataN, v.Id)
		if result.RowsAffected == 0 {
			return nil, errors.New("data reg narahubung tidak dapat ditemukan")
		}
		if err := sc.Copy(&dataN, &v); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload njoptkp flag detail", dataN)
		}

		dataN.NjoptkpFlag_Id = njoptkpFlag_id
		if result := tx.Save(&dataN); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data reg narahubung", dataN)
		}
	}
	for k := range data {
		data[k].NjoptkpFlag_Id = njoptkpFlag_id
	}
	return rp.OKSimple{Data: data}, nil
}

func Delete(njoptkpFlag_id uint64, tx *gorm.DB) error {
	var dataN []*m.NjoptkpFlagDetail
	result := tx.Where(m.NjoptkpFlagDetail{NjoptkpFlag_Id: njoptkpFlag_id}).Find(&dataN)
	if result.RowsAffected == 0 {
		return errors.New("data njoptkp flag detail tidak dapat ditemukan")
	}

	for _, v := range dataN {
		result = tx.Delete(&v)
		if result.RowsAffected == 0 {
			return errors.New("tidak dapat menghapus data njoptkp flag detail")
		}
	}
	return nil
}
