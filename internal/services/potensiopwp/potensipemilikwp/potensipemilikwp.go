package potensipemilikwp

import (
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"

	mareadivision "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensipemilikwp"

	sdaerah "github.com/bapenda-kota-malang/apin-backend/internal/services/daerah"
	skecamatan "github.com/bapenda-kota-malang/apin-backend/internal/services/kecamatan"
)

const source = "potensipemilikwp"

func Create(input []m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.PotensiPemilikWp

	for i := range input {
		// check linked area
		respDaerah, err := sdaerah.GetDetail(int(input[i].Daerah_Id))
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data kecamatan", data)
		}
		respKecamatan, err := skecamatan.GetDetail(int(input[i].Kecamatan_Id))
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data kecamatan", data)
		}
		daerah := respDaerah.(rp.OKSimple).Data.(*mareadivision.Daerah)
		kecamatan := respKecamatan.(rp.OKSimple).Data.(*mareadivision.Kecamatan)
		if kecamatan.Daerah_Kode != daerah.Kode {
			return sh.SetError("request", "create-data", source, "failed", "kecamatan tidak berada pada kota terkait", data)
		}
	}

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

func Update(input []m.PotensiPemilikWp, tx *gorm.DB) (any, error) {
	// var data *m.PotensiPemilikWp
	// result := a.DB.First(&data, id)
	// if result.RowsAffected == 0 {
	// 	return nil, nil
	// }

	// if err := sc.Copy(&data, &input); err != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	// }

	// if result := a.DB.Save(&data); result.Error != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	// }

	// return rp.OK{
	// 	Meta: t.IS{
	// 		"affected": strconv.Itoa(int(result.RowsAffected)),
	// 	},
	// 	Data: data,
	// }, nil
	return nil, nil
}
