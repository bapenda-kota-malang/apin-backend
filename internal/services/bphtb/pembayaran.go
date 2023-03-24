package sptpd

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	// "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

func CreatePembayaran(input m.RequestPembayaranBphtb, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.PembayaranBphtb

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data ", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetListPembayaran(input m.RequestPembayaranBphtb) (any, error) {
	var (
		data  []m.PembayaranBphtb
		count int64
	)

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.PembayaranBphtb{}).Preload(clause.Associations).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data ", data)
	}

	meta := t.IS{
		"totalCount":   strconv.Itoa(int(count)),
		"currentCount": strconv.Itoa(int(result.RowsAffected)),
		"page":         strconv.Itoa(pagination.Page),
		"pageSize":     strconv.Itoa(pagination.PageSize),
	}

	if len(data) == 0 {
		return rp.OK{Meta: meta, Data: []string{}}, nil
	}

	return rp.OK{
		Meta: meta,
		Data: data,
	}, nil
}

func GetDetailPembayaran(id int) (any, error) {
	var data *m.PembayaranBphtb

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data ", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func UpdatePembayaran(id int, input m.RequestPembayaranBphtb, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.PembayaranBphtb
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data ", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func DeletePembayaran(id int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.PembayaranBphtb
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	result = tx.Delete(&data, id)
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
