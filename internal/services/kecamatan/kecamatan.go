package kecamatan

import (
	"errors"
	"fmt"
	"strconv"

	sc "github.com/jinzhu/copier"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "kecamatan"

func Create(input m.KecamatanCreateDto) (any, error) {
	var data m.Kecamatan
	var dataD m.Daerah
	checkRegencyCode := a.DB.Where(&m.Daerah{Kode: input.Daerah_Kode}).First(&dataD)
	if checkRegencyCode.RowsAffected == 0 {
		return nil, errors.New("daerah dengan kode tersebut belum terdaftar")
	}
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	result := a.DB.Create(&data)
	if result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.KecamatanFilterDto) (interface{}, error) {
	var data []m.Kecamatan
	var count int64
	var pagination gh.Pagination

	query := a.DB.
		Model(&m.Kecamatan{}).
		Preload("Daerah.Provinsi").
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination))
	result := query.Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
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

func GetDetail(id int) (interface{}, error) {
	var data *m.Kecamatan
	result := a.DB.Model(&m.Kecamatan{}).
		Preload("Daerah.Provinsi").
		First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func GetDetailByCode(id int) (interface{}, error) {
	var data *m.Kecamatan

	result := a.DB.Where("Nop", fmt.Sprint(id)).First(&data)
	if result.RowsAffected == 0 {
		result := a.DB.Where("Kode", fmt.Sprint(id)).First(&data)
		if result.RowsAffected == 0 {
			return nil, nil
		} else if result.Error != nil {
			return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
		}
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.KecamatanUpdateDto) (interface{}, error) {
	var data *m.Kecamatan

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (interface{}, error) {
	var data *m.Kecamatan
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
