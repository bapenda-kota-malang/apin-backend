package potensipemilikwp

import (
	"strconv"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"

	mareadivision "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensipemilikwp"

	sdaerah "github.com/bapenda-kota-malang/apin-backend/internal/services/daerah"
	skecamatan "github.com/bapenda-kota-malang/apin-backend/internal/services/kecamatan"

	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "potensipemilikwp"

func Create(input []m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.PotensiPemilikWp

	for i := range input {
		// check linked area
		if input[i].Kecamatan_Id != nil {
			respDaerah, err := sdaerah.GetDetail(int(input[i].Daerah_Id))
			if err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data kecamatan", data)
			}
			respKecamatan, err := skecamatan.GetDetail(int(*input[i].Kecamatan_Id))
			if err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data kecamatan", data)
			}
			daerah := respDaerah.(rp.OKSimple).Data.(*mareadivision.Daerah)
			kecamatan := respKecamatan.(rp.OKSimple).Data.(*mareadivision.Kecamatan)
			if kecamatan.Daerah_Kode != daerah.Kode {
				return sh.SetError("request", "create-data", source, "failed", "kecamatan tidak berada pada kota terkait", data)
			}
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

func Update(potensiOp_Id int, input []m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.PotensiPemilikWp
	rowAffected := 0
	for _, v := range input {
		var item m.PotensiPemilikWp
		if v.Id != nil {
			result := tx.First(&item, v.Id)
			if result.RowsAffected == 0 {
				return sh.SetError("request", "update-data", source, "failed", "data not found", data)
			}
			if item.Potensiop_Id != uint(potensiOp_Id) {
				return sh.SetError("request", "update-data", source, "failed", "tidak bisa mengubah data ini", v)
			}
		} else {
			item.Potensiop_Id = uint(potensiOp_Id)
		}
		if err := sc.Copy(&item, &v); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
		}
		if result := tx.Save(&item); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
		}
		data = append(data, item)
		rowAffected++
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(rowAffected),
		},
		Data: data,
	}, nil
}
