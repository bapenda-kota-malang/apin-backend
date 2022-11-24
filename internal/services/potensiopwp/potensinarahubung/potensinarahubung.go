package potensinarahubung

import (
	"strconv"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/google/uuid"
	sc "github.com/jinzhu/copier"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensinarahubung"

	skelurahan "github.com/bapenda-kota-malang/apin-backend/internal/services/kelurahan"

	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "potensinarahubung"

func Create(input []m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.PotensiNarahubung

	for i := range input {
		// check linked area
		if input[i].Kelurahan_Id != nil {
			_, err := skelurahan.CrossCheckDaerah(input[i].Daerah_Id, *input[i].Kelurahan_Id)
			if err != nil {
				return sh.SetError("request", "create-data", source, "failed", "data kelurahan bukan di kota terkait", data)
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

func Update(potensiOp_Id uuid.UUID, input []m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.PotensiNarahubung
	rowAffected := 0
	for _, v := range input {
		var item m.PotensiNarahubung
		if v.Id != nil {
			result := tx.First(&item, v.Id)
			if result.RowsAffected == 0 {
				return sh.SetError("request", "update-data", source, "failed", "data not found", data)
			}
			if item.Potensiop_Id != potensiOp_Id {
				return sh.SetError("request", "update-data", source, "failed", "tidak bisa mengubah data ini", v)
			}
		} else {
			item.Potensiop_Id = potensiOp_Id
		}
		if err := sc.Copy(&item, &v); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
		}

		if v.Delete != nil && *v.Delete {
			if result := tx.Delete(&item, v.Id); result.RowsAffected == 0 {
				return sh.SetError("request", "update-data", source, "failed", "gagal menghapus data menyimpan data", data)
			}
		} else {
			if result := tx.Save(&item); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
			}
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
