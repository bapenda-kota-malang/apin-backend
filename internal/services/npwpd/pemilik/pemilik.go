package npwpd

import (
	"errors"
	"fmt"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "pemilik"

func Create(input []m.PemilikWpCreateDto, npwpd_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.PemilikWp
	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), data)
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

func Update(input []m.PemilikWpUpdateDto, npwpd_id uint64, golongan nt.Golongan, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	var data []m.PemilikWp

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), data)
	}

	for _, v := range data {
		var dataP *m.PemilikWp
		result := tx.First(&dataP, v.Id)
		if result.RowsAffected == 0 {
			return nil, errors.New("data pemilik tidak dapat ditemukan")
		}
		if err := sc.Copy(&dataP, &v); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload pemilik", dataP)
		}
		// set directur value to null if golongan orang pribadi
		if golongan == 1 {
			dataP.Direktur_Nama = nil
			dataP.Direktur_Alamat = nil
			dataP.Direktur_Daerah_Id = nil
			dataP.Direktur_Kelurahan_Id = nil
			dataP.Direktur_Nik = nil
			dataP.Direktur_Telp = nil
		}
		dataP.Npwpd_Id = npwpd_id
		if result := tx.Save(&dataP); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data pemilik", dataP)
		}
	}
	for k := range data {
		data[k].Npwpd_Id = npwpd_id
	}
	return rp.OKSimple{Data: data}, nil
}

func Delete(npwpd_Id uint64, tx *gorm.DB) error {
	if tx == nil {
		tx = a.DB
	}
	var dataPemilik []*m.PemilikWp
	result := tx.Where(m.PemilikWp{Npwpd_Id: npwpd_Id}).Find(&dataPemilik)
	if result.RowsAffected == 0 {
		return errors.New("data reg pemilik tidak dapat ditemukan")
	}

	for _, v := range dataPemilik {
		result = tx.Delete(&v)
		if result.RowsAffected == 0 {
			return errors.New("tidak dapat menghapus data reg pemilik")
		}
	}
	return nil
}
