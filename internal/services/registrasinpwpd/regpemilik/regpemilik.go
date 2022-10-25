package registrasinpwpd

import (
	"fmt"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/registrasinpwpd"

	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
)

const source = "regpemilik"

func Create(input []m.RegPemilikWpCreate, regis_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.RegPemilikWp
	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// static add value to field
	for k := range data {
		data[k].RegistrasiNpwpd_Id = regis_id
		data[k].Status = nt.StatusBaru
	}

	// Transaction save to db
	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data %s", source), data)
	}

	return rp.OKSimple{Data: data}, nil
}
