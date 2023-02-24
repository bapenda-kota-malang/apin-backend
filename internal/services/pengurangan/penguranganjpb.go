package pengurangan

import (
	"errors"
	"fmt"
	"strconv"

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

const sourceJPB = "penguranganJPB"

func CreateJPB(input m.CreateDtoJPB) (any, error) {
	var data m.PenguranganJPB
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", sourceJPB, "failed", "gagal mengambil data payload", data)
	}

	txDb := a.DB

	err := txDb.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&data)
		if result.Error != nil {
			return result.Error
		}
		_, err := ssppt.UpdatePenguranganByNop(msppt.NopDto{
			Propinsi_Id:        data.Provinsi_Kode_Pemohon,
			Dati2_Id:           data.Daerah_Kode_Pemohon,
			Kecamatan_Id:       data.Kecamatan_Kode_Pemohon,
			Keluarahan_Id:      data.Kelurahan_Kode_Pemohon,
			Blok_Id:            data.Blok_Kode_Pemohon,
			NoUrut:             data.NoUrutPelayanan,
			JenisOP_Id:         data.JenisOp_Pemohon,
			TahunPajakskp_sppt: data.ThnPengenaan,
		}, *data.PCTPenguranganJPB, tx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", sourceJPB, "failed", fmt.Sprintf("gagal mengambil menyimpan data: %s", err), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetListJPB(input m.FilterDtoJPB) (any, error) {
	var data []m.PenguranganJPB
	var count int64
	var pagination gh.Pagination

	result := a.DB.
		Model(&m.PenguranganJPB{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", sourceJPB, "failed", "gagal mengambil data", data)
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

func GetDetailJPB(id int) (interface{}, error) {
	var data *m.PenguranganJPB
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", sourceJPB, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func UpdateJPB(id int, input m.UpdateDtoJPB) (interface{}, error) {
	var data *m.PenguranganJPB

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.CopyWithOption(&data, input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "update-data", sourceJPB, "failed", "gagal mengambil data payload", data)
	}
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Save(&data); result.Error != nil {
			return result.Error
		}
		if input.PCTPenguranganJPB != nil {
			_, err := ssppt.UpdatePenguranganByNop(msppt.NopDto{
				Propinsi_Id:        data.Provinsi_Kode_Pemohon,
				Dati2_Id:           data.Daerah_Kode_Pemohon,
				Kecamatan_Id:       data.Kecamatan_Kode_Pemohon,
				Keluarahan_Id:      data.Kelurahan_Kode_Pemohon,
				Blok_Id:            data.Blok_Kode_Pemohon,
				NoUrut:             data.NoUrut_Pemohon,
				JenisOP_Id:         data.JenisOp_Pemohon,
				TahunPajakskp_sppt: data.ThnPengenaan,
			}, *data.PCTPenguranganJPB, tx)
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

func DeleteJPB(id int) (interface{}, error) {
	var data *m.PenguranganJPB
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
