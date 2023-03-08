// service
package ppat

import (
	"errors"
	"strconv"

	sc "github.com/jinzhu/copier"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/tempatpembayaranspptmasal"
)

// /// Private funcs start here
const source = "tempatpembayaranspptmasal"

///// Exported funcs start here

func Create(input m.CreateDto) (any, error) {
	var data m.TempatPembayaranSPPTMasal

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := a.DB.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+result.Error.Error(), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.TempatPembayaranSPPTMasal
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Debug().
		Model(&m.TempatPembayaranSPPTMasal{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Preload("Kelurahan").
		Preload("TempatPembayaran").
		Preload("BankTunggal").
		Preload("BankPersepsi").
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	meta := t.IS{
		"totalCount":   strconv.Itoa(int(count)),
		"currentCount": strconv.Itoa(int(result.RowsAffected)),
		"page":         strconv.Itoa(pagination.Page),
		"pageSize":     strconv.Itoa(pagination.PageSize),
	}

	if len(data) == 0 {
		return rp.OK{
			Meta: meta,
			Data: []string{},
		}, nil
	}

	return rp.OK{
		Meta: meta,
		Data: data,
	}, nil
}

func GetDetail(id int) (interface{}, error) {
	var data *m.TempatPembayaranSPPTMasal
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(id int, input m.UpdateDto) (interface{}, error) {
	var data *m.TempatPembayaranSPPTMasal

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "data ppat tidak dapat ditemukan", data)
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	result = a.DB.Save(&data)
	if result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data ppat: "+result.Error.Error(), data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (interface{}, error) {
	var data *m.TempatPembayaranSPPTMasal
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
