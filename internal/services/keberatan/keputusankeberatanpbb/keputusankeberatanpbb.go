package keputusankeberatanpbb

import (
	"errors"
	"fmt"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/keberatan"
	msksk "github.com/bapenda-kota-malang/apin-backend/internal/models/sksk"
	ssksk "github.com/bapenda-kota-malang/apin-backend/internal/services/sksk"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "KeputusanKeberatanPbb"

func Create(input m.CreateDtoKepKebPbb, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.KeputusanKeberatanPbb
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	err := tx.Transaction(func(txChild *gorm.DB) error {
		result := txChild.Create(&data)
		if result.Error != nil {
			return result.Error
		}
		var createDtoSksk msksk.CreateDto
		if err := sc.Copy(&createDtoSksk, data); err != nil {
			return err
		}
		if _, err := ssksk.Create(createDtoSksk, tx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data: %s", err), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDtoKepKebPbb) (interface{}, error) {
	var data []m.KeputusanKeberatanPbb
	var count int64
	var pagination gh.Pagination

	query := a.DB.
		Model(&m.KeputusanKeberatanPbb{}).
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
	var data *m.KeputusanKeberatanPbb
	result := a.DB.Model(&m.KeputusanKeberatanPbb{}).
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

func Update(id int, input m.UpdateDtoKepKebPbb, tx *gorm.DB) (interface{}, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.KeputusanKeberatanPbb

	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.CopyWithOption(&data, input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}
	err := tx.Transaction(func(txChild *gorm.DB) error {
		if result := txChild.Save(&data); result.Error != nil {
			return result.Error
		}
		if *data.JnsKeputusan == m.JnsKeputusanDiterima {
			// TODO: Call SP sppt recalculate penilaian dan penetapan, ask luke
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("gagal memperbarui data: %s", err), data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (interface{}, error) {
	var data *m.KeputusanKeberatanPbb
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
