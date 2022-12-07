package npwpd

import (
	"errors"
	"fmt"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
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
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal menyimpan data %s", source), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(input []m.PemilikWpUpdateDto, npwpd_id uint64, golongan t.Golongan, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	var dataUpdate []m.PemilikWp
	var dataDelete []m.PemilikWp
	var tmpDataUpdate m.PemilikWp
	var tmpDataDelete m.PemilikWp

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
		var dataPU *m.PemilikWp
		result := tx.First(&dataPU, v.Id)
		if result.RowsAffected == 0 {
			return nil, errors.New("data pemilik tidak dapat ditemukan")
		}
		if err := sc.Copy(&dataPU, &v); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload pemilik", dataPU)
		}
		// set directur value to null if golongan orang pribadi
		if golongan == 1 {
			dataPU.Direktur_Nama = nil
			dataPU.Direktur_Alamat = nil
			dataPU.Direktur_Daerah_Id = nil
			dataPU.Direktur_Kelurahan_Id = nil
			dataPU.Direktur_Nik = nil
			dataPU.Direktur_Telp = nil
		}
		dataPU.Npwpd_Id = npwpd_id
		if result := tx.Save(&dataPU); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data pemilik", dataPU)
		}
	}
	for k := range dataUpdate {
		dataUpdate[k].Npwpd_Id = npwpd_id
	}

	// delete data
	for _, v := range dataDelete {
		var dataPD *m.PemilikWp
		result := tx.First(&dataPD, v.Id)
		if result.RowsAffected == 0 {
			return nil, errors.New("data pemilik tidak dapat ditemukan")
		}

		result = tx.Delete(&dataPD)
		if result.RowsAffected == 0 {
			return nil, errors.New("tidak dapat menghapus data reg pemilik")
		}
	}
	return rp.OKSimple{Data: dataUpdate}, nil
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
