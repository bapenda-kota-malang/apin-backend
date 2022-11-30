package suratpemberitahuan

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/suratpemberitahuan"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "suratpemberitahuan"

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.SuratPemberitahuan

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// default value
	data.Status = 0

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.SuratPemberitahuan
	var count int64

	var pagination gh.Pagination
	result := tx.
		Model(&m.SuratPemberitahuan{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Preload("Npwpd").
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

func GetDetail(id int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.SuratPemberitahuan

	result := tx.Preload(clause.Associations).Preload("SuratPemberitahuanDetail.Spt").First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func UpdateBulk(input m.UpdateBulkDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var datas []m.SuratPemberitahuan
	resultRow := 0
	for _, v := range input.Datas {
		respUpdateSingle, err := Update(int(*v.Id), m.UpdateDto{Tanggal: v.Tanggal, Status: v.Status}, tx)
		if err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", datas)
		}
		datas = append(datas, respUpdateSingle.(rp.OK).Data.(m.SuratPemberitahuan))
		result, err := strconv.Atoi(respUpdateSingle.(rp.OK).Meta["affected"])
		if err != nil {
			return sh.SetError("request", "update-data", source, "failed", "meta affected bukan angka", datas)
		}
		resultRow += result
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(resultRow),
		},
		Data: datas,
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	data := m.SuratPemberitahuan{}
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	if err := sc.CopyWithOption(&data, &input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}
	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.SuratPemberitahuan
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
