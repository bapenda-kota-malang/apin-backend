package jenisusaha

import (
	"errors"
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/jenisusaha"
	sjud "github.com/bapenda-kota-malang/apin-backend/internal/services/jenisusaha/jenisusahadetail"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "jenis usaha"

func Create(input m.CreateDto) (any, error) {
	var data m.JenisUsaha
	var dataDetail m.JenisUsahaDetailCreateDto
	var respDataDetail interface{}
	var resp t.II

	// data jenis usaha
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jenis usaha", data)
	}

	// data jenis usaha detail
	if err := sc.Copy(&dataDetail, &input.JenisUsahaDetail); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jenis usaha detail", dataDetail)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {

		// create data jenis usaha
		err := tx.Create(&data).Error
		if err != nil {
			return err
		}

		// create data jenis usaha detail
		resultDetail, err := sjud.Create(dataDetail, data.Id, tx)
		if err != nil {
			return err
		}
		respDataDetail = resultDetail

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	resp = t.II{
		"jenisUsaha":       data,
		"jenisUsahaDetail": respDataDetail.(rp.OKSimple).Data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.JenisUsaha
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.JenisUsaha{}).
		Preload(clause.Associations).
		Preload("JenisUsahaDetail.Rekening").
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data jenis usaha", data)
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
	var data *m.JenisUsaha

	result := a.DB.Model(m.JenisUsaha{}).
		Preload(clause.Associations).
		Preload("JenisUsahaDetail.Rekening").
		First(&data, id)

	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data jenis usaha", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Delete(jenisUsaha_id uint64) (any, error) {
	var status string = "deleted"
	var data *m.JenisUsaha

	// cek data jenis usaha
	result := a.DB.First(&data, jenisUsaha_id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data jenis usaha tidak dapat ditemukan")
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// delete jenis usaha detail
		err := sjud.Delete(data.Id, tx)
		if err != nil {
			status = "no deletion"
			return err
		}

		// delete jenis usaha
		result = tx.Delete(&data, jenisUsaha_id)
		if result.RowsAffected == 0 {
			return errors.New("tidak dapat menghapus data jenis usaha: " + result.Error.Error())
		}
		return nil

	})
	if err != nil {
		return sh.SetError("request", "delete-data", source, "failed", "gagal menghapus data: "+err.Error(), data)
	}

	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto) (any, error) {
	var data *m.JenisUsaha
	var dataDetail m.JenisUsahaDetailUpdateDto
	var respDataDetail interface{}
	var resp t.II

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data jenis usaha tidak dapat ditemukan")
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload jenis usaha", input)
	}

	if err := sc.Copy(&dataDetail, &input.JenisUsahaDetail); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload jenis usaha detail", input.JenisUsahaDetail)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// update sspd
		if result := tx.Save(&data); result.Error != nil {
			return errors.New("gagal menyimpan data jenis usaha")
		}

		resultDetail, err := sjud.Update(dataDetail, data.Id, tx)
		if err != nil {
			return err
		}
		respDataDetail = resultDetail

		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	resp = t.II{
		"affected":         strconv.Itoa(int(result.RowsAffected)),
		"jenisUsahaDetail": respDataDetail.(rp.OKSimple).Data,
		"jenisUsaha":       data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}
