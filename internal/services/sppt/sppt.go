package skpd

import (
	"strconv"

	sc "github.com/jinzhu/copier"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "skpd"

func Create(input m.RequestDto) (any, error) {
	var data m.Sppt

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := a.DB.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.Sppt
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Sppt{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
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

func GetDetail(id int) (any, error) {
	var data *m.Sppt

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func GetByNop(provinsiKode, daerahKode, kecamatanKode, kelurahanKode, blokKode, noUrut, jenisOp string) (any, error) {
	var data m.Sppt
	result := a.DB.Where(&m.Sppt{
		Propinsi_Id:   &provinsiKode,
		Dati2_Id:      &daerahKode,
		Kecamatan_Id:  &kecamatanKode,
		Keluarahan_Id: &kelurahanKode,
		Blok_Id:       &blokKode,
		NoUrut:        &noUrut,
		JenisOP_Id:    &jenisOp,
	}).First(&data)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "get-data-by-nop", source, "failed", "data tidak ada", data)
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-by-nop", source, "failed", "gagal mendapatkan data: "+result.Error.Error(), data)
	}
	return rp.OKSimple{Data: data}, nil
}

func Update(id int, input m.RequestDto) (any, error) {
	var data *m.Sppt
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	if err := sc.Copy(&data, &input); err != nil {
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

func Delete(id int) (any, error) {
	var data *m.Sppt
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
