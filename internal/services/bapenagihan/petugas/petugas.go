package petugas

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bapenagihan/petugas"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "bapenagihanpetugas"

func Create(input []m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.BaPenagihanPetugas

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", result.Error.Error(), data)
	}

	return rp.OKSimple{Data: data}, nil
}

// not true update but this process check if data is flaggged as delete, if flagged as delete then delete else is create data
func Update(input []m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var datas []m.BaPenagihanPetugas
	rowsAffected := 0
	for _, v := range input {
		var data m.BaPenagihanPetugas
		result := tx.Where(&m.BaPenagihanPetugas{BaPenagihan_Id: v.BaPenagihan_Id, Petugas_Id: v.Petugas_Id}).First(&data)
		rowsAffected += int(result.RowsAffected)
		if v.Deleted && result.RowsAffected != 0 {
			if result := tx.Delete(&data); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal menghapus data", data)
			}
		} else {
			// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
			if err := sc.Copy(&data, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
			}

			if result.RowsAffected == 0 {
				if result := tx.Create(&data); result.Error != nil {
					return sh.SetError("request", "create-data", source, "failed", result.Error.Error(), data)
				}
			} else {
				if result := tx.Save(&data); result.Error != nil {
					return sh.SetError("request", "create-data", source, "failed", result.Error.Error(), data)
				}
			}
			rowsAffected += 1
		}
		datas = append(datas, data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(rowsAffected),
		},
		Data: datas,
	}, nil
}

func Delete(id int) (any, error) {
	var data *m.BaPenagihanPetugas
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	result = a.DB.Delete(&data, id)
	status := "deleted"
	if result.RowsAffected == 0 {
		data = nil
		status = "no deletion"
	}

	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil
}
