package sts

import (
	"fmt"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sts"
	sd "github.com/bapenda-kota-malang/apin-backend/internal/services/sspd/sspddetail"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "sts detail"

func Create(input []m.StsDetailCreateDto, sts_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var dataResult []m.StsDetail
	for _, v := range input {

		var data m.StsDetail
		//  copy input (payload) ke struct data jika tidak ada akan error
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), data)
		}

		// static add value to field
		data.Sts_Id = &sts_id

		if result := sd.UpdateFromSts(v.SspdDetail_Id, sts_id, tx); result != nil {
			return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal menambahkan data sts id ke sspd detail %s", source), data)
		}

		// Transaction save to db
		// simpan data ke db satu if karena result dipakai sekali, +error
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal menyimpan data %s", source), data)
		}
		dataResult = append(dataResult, data)
	}

	return rp.OKSimple{Data: dataResult}, nil
}
