package pengurangan

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pengurangan"
	msppt "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	ssppt "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const sourcePermanen = "penguranganPermanen"

func CreatePermanen(input m.CreateDtoPermanen) (any, error) {
	var data m.PenguranganPermanen
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", sourcePermanen, "failed", "gagal mengambil data payload", data)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&data)
		if result.Error != nil {
			return result.Error
		}

		yearNow := time.Now().Year()
		tahunAwal, _ := strconv.Atoi(*data.TahunPenguranganPermanenAwal)
		tahunAkhir, _ := strconv.Atoi(*data.TahunPenguranganPermanenAkhir)
		if yearNow < tahunAwal || yearNow > tahunAkhir {
			return fmt.Errorf("tahun tidak dalam range tahun pengurangan permanen")
		}

		_, err := ssppt.UpdatePenguranganByNop(msppt.NopDto{
			Propinsi_Id:        data.ProvinsiPemohon_Kd,
			Dati2_Id:           data.DaerahPemohon_Kd,
			Kecamatan_Id:       data.KecamatanPemohon_Kd,
			Keluarahan_Id:      data.KelurahanPemohon_Kd,
			Blok_Id:            data.BlokPemohon_Kd,
			NoUrut:             data.NoUrutPemohon_Kd,
			JenisOP_Id:         data.JenisOpPemohon_Kd,
			TahunPajakskp_sppt: data.TahunPenguranganPermanenAwal,
		}, *data.PctPenguranganPermanen, tx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", sourcePermanen, "failed", fmt.Sprintf("gagal mengambil menyimpan data: %s", err), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetListPermanen(input m.FilterDtoPermanen) (any, error) {
	var data []m.PenguranganPermanen
	var count int64
	var pagination gh.Pagination

	result := a.DB.
		Model(&m.PenguranganPermanen{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", sourcePermanen, "failed", "gagal mengambil data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: data,
	}, nil
}

func GetDetailPermanen(id int) (interface{}, error) {
	var data *m.PenguranganPermanen
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", sourcePermanen, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func UpdatePermanen(id int, input m.UpdateDtoPermanen) (interface{}, error) {
	var data *m.PenguranganPermanen

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.CopyWithOption(&data, input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "update-data", sourcePermanen, "failed", "gagal mengambil data payload", data)
	}
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Save(&data); result.Error != nil {
			return result.Error
		}
		if input.PctPenguranganPermanen != nil {
			yearNow := time.Now().Year()
			tahunAwal, _ := strconv.Atoi(*data.TahunPenguranganPermanenAwal)
			tahunAkhir, _ := strconv.Atoi(*data.TahunPenguranganPermanenAkhir)
			if yearNow < tahunAwal || yearNow > tahunAkhir {
				return fmt.Errorf("tahun tidak dalam range tahun pengurangan permanen")
			}
			_, err := ssppt.UpdatePenguranganByNop(msppt.NopDto{
				Propinsi_Id:        data.ProvinsiPemohon_Kd,
				Dati2_Id:           data.DaerahPemohon_Kd,
				Kecamatan_Id:       data.KecamatanPemohon_Kd,
				Keluarahan_Id:      data.KelurahanPemohon_Kd,
				Blok_Id:            data.BlokPemohon_Kd,
				NoUrut:             data.NoUrutPemohon_Kd,
				JenisOP_Id:         data.JenisOpPemohon_Kd,
				TahunPajakskp_sppt: data.TahunPenguranganPermanenAwal,
			}, *data.PctPenguranganPermanen, tx)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", sourceDendaADM, "failed", fmt.Sprintf("gagal mengambil menyimpan data: %s", err), data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func DeletePermanen(id int) (interface{}, error) {
	var data *m.PenguranganPermanen
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
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
