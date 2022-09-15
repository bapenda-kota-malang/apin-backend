package potensiop

import (
	"errors"
	"net/url"
	"strconv"

	sc "github.com/jinzhu/copier"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	registration "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "potensiop"

func GetList(query url.Values, pagination gh.Pagination) (any, error) {
	var data []m.PotensiOp
	var count int64

	filtered := a.DB.Table("PotensiOp").Scopes(gh.Filter(query, m.PotensiOp{}))
	filtered.Count(&count)

	result := filtered.Scopes(gh.Paginate(&pagination)).Find(&data)
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
	var data *m.PotensiOp

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

func Create(input m.CreatePotensiOp) (any, error) {
	var data m.PotensiOp

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// TODO: check rekening_id exist in table rekening
	// TODO: check user_id exist in table user

	data.Status = registration.StatusAktif

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := a.DB.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(id int, input m.CreatePotensiOp) (any, error) {
	var data *m.PotensiOp
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
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
	var data *m.PotensiOp
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
