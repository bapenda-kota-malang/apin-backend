package detail

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/suratpemberitahuan/detail"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "suratpemberitahuandetail"

func Create(input []m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.SuratPemberitahuanDetail

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

// func Update(id int, input m.UpdateDto) (any, error) {
// 	var data *m.SuratPemberitahuanDetail
// 	result := a.DB.First(&data, id)
// 	if result.RowsAffected == 0 {
// 		return nil, nil
// 	}

// 	if err := sc.Copy(&data, &input); err != nil {
// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
// 	}

// 	if result := a.DB.Save(&data); result.Error != nil {
// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
// 	}

// 	return rp.OK{
// 		Meta: t.IS{
// 			"affected": strconv.Itoa(int(result.RowsAffected)),
// 		},
// 		Data: data,
// 	}, nil
// }

func Delete(id int) (any, error) {
	var data *m.SuratPemberitahuanDetail
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
