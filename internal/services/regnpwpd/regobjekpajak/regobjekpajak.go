package regnpwpd

import (
	"errors"
	"fmt"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	mop "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajak"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajak"

	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
)

const source = "regobjekpajak"

func Create(input m.RegObjekPajakCreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.RegObjekPajak

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// static add value to field
	data.Status = nt.StatusBaru

	// Transaction save to db
	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data %s", source), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(input m.RegObjekPajakUpdateDto, regObjekPajak_Id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	var data m.RegObjekPajak

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), data)
	}
	var dataOp *m.RegObjekPajak
	result := tx.First(&dataOp, regObjekPajak_Id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data reg objek pajak tidak dapat ditemukan")
	}
	if err := sc.Copy(&dataOp, &data); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload reg objek pajak", data)
	}

	dataOp.Id = regObjekPajak_Id
	if result := tx.Save(&dataOp); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data reg objek pajak", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func Delete(regObjekPajak_Id uint64, tx *gorm.DB) error {
	var dataRegObjekPajak *m.RegObjekPajak
	result := tx.Where(m.RegObjekPajak{Id: uint64(regObjekPajak_Id)}).First(&dataRegObjekPajak)
	if result.RowsAffected == 0 {
		return errors.New("data reg objek pajak tidak dapat ditemukan")
	}
	result = tx.Delete(&dataRegObjekPajak)
	if result.RowsAffected == 0 {
		return errors.New("tidak dapat menghapus data reg objek pajak")
	}
	return nil
}

func Verify(regObjekPajak_Id uint64, tx *gorm.DB) (any, error) {
	var dataRegObjekPajak *m.RegObjekPajak
	result := tx.Where(m.RegObjekPajak{Id: regObjekPajak_Id}).First(&dataRegObjekPajak)
	if result.RowsAffected == 0 {
		return nil, errors.New("data reg objek pajak tidak dapat ditemukan")
	}

	var dataObjekPajak mop.ObjekPajak
	if err := sc.Copy(&dataObjekPajak, dataRegObjekPajak); err != nil {
		return nil, errors.New("gagal menyalin data reg objek pajak ke objek pajak")
	}

	dataObjekPajak.Id = 0
	err := tx.Create(&dataObjekPajak).Error
	if err != nil {
		return nil, errors.New("gagal memindahkan data objek pajak")
	}
	return rp.OKSimple{Data: dataObjekPajak}, nil

}
