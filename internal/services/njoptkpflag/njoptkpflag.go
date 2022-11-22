package njoptkpflag

import (
	"errors"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/njoptkpflag"
	snfd "github.com/bapenda-kota-malang/apin-backend/internal/services/njoptkpflag/njoptkpdetail"
)

const source = "njoptkp flag"

func Create(input m.CreateDto) (interface{}, error) {
	var data m.NjoptkpFlag
	var dataDetail []m.NjoptkpFlagDetailCreateDto
	var respDataDetail interface{}
	var resp t.II

	// data njoptkp flag
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload njoptkp flag", data)
	}

	// data njoptkp flag detail
	if err := sc.Copy(&dataDetail, &input.NjoptkpFlagDetail); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload njoptkp flag detail", dataDetail)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {

		err := tx.Create(&data).Error
		if err != nil {
			return err
		}

		resultDetail, err := snfd.Create(dataDetail, data.Id, tx)
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
		"njoptkpFlagDetail": respDataDetail.(rp.OKSimple).Data,
		"njoptkpFlag":       data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.NjoptkpFlag
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.NjoptkpFlag{}).
		Preload(clause.Associations).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data njoptkp flag", data)
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
	var data *m.NjoptkpFlag

	result := a.DB.Model(m.NjoptkpFlag{}).
		Preload(clause.Associations).
		First(&data, id)

	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data njoptkp flag", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Delete(njoptkpFlag_id uint64) (any, error) {
	var status string = "deleted"
	var data *m.NjoptkpFlag

	// cek data sspd
	result := a.DB.First(&data, njoptkpFlag_id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data njoptkp flag tidak dapat ditemukan")
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// delete njoptkp flag detail
		err := snfd.Delete(data.Id, tx)
		if err != nil {
			status = "no deletion"
			return err
		}

		// delete sspd
		result = tx.Delete(&data, njoptkpFlag_id)
		if result.RowsAffected == 0 {
			return errors.New("tidak dapat menghapus data njoptkp flag: " + result.Error.Error())
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
	var data *m.NjoptkpFlag
	var dataNjoptkpFlagDetail []m.NjoptkpFlagDetailUpdateDto
	var respDataNjoptkpFlagDetail interface{}
	var resp t.II

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data njoptkp flag tidak dapat ditemukan")
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload njoptkp flag", input)
	}

	if err := sc.Copy(&dataNjoptkpFlagDetail, &input.NjoptkpFlagDetail); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload njoptkp flag detail", input.NjoptkpFlagDetail)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// update njoptkp flag
		if result := tx.Save(&data); result.Error != nil {
			return errors.New("gagal menyimpan data njoptkp flag")
		}

		resultNjoptkpFlagDetail, err := snfd.Update(dataNjoptkpFlagDetail, data.Id, tx)
		if err != nil {
			return err
		}
		respDataNjoptkpFlagDetail = resultNjoptkpFlagDetail

		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	resp = t.II{
		"affected":          strconv.Itoa(int(result.RowsAffected)),
		"njoptkpFlagDetail": respDataNjoptkpFlagDetail.(rp.OKSimple).Data,
		"njoptkpFlag":       data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}
