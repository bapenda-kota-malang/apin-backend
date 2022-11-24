package sts

import (
	"errors"
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sts"
	ssd "github.com/bapenda-kota-malang/apin-backend/internal/services/sts/stsdetail"
	sss "github.com/bapenda-kota-malang/apin-backend/internal/services/sts/sumberdanasts"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "sts"

func Create(input m.CreateDto) (any, error) {
	var dataSts m.Sts
	var dataStsDetail []m.StsDetailCreateDto
	var dataSumberDanaSts []m.SumberDanaStsCreateDto
	var respDataStsDetail interface{}
	var respDataSumberDanaSts interface{}
	var resp t.II

	// data sts
	if err := sc.Copy(&dataSts, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload sts", dataSts)
	}

	// data sts detail
	if err := sc.Copy(&dataStsDetail, &input.StsDetail); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload sts detail", dataStsDetail)
	}

	// data sumber dana sts
	if err := sc.Copy(&dataSumberDanaSts, &input.SumberDanaSts); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload sts detail", dataSumberDanaSts)
	}

	// static value
	dataSts.NomorUrut = generateNomorUrut()
	dataSts.NomorOutput = generateNomorOutput(dataSts.NomorUrut)
	if input.TanggalSetor == nil {
		dataSts.TanggalSetor = th.TimeNow()
	} else {
		dataSts.TanggalSetor = th.ParseTime(input.TanggalSetor)
	}

	if input.TanggalSts == nil {
		dataSts.TanggalSts = th.TimeNow()
	} else {
		dataSts.TanggalSts = th.ParseTime(input.TanggalSts)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {

		// create data sts
		err := tx.Create(&dataSts).Error
		if err != nil {
			return err
		}

		// create data sts detail
		resultStsDetail, err := ssd.Create(dataStsDetail, dataSts.Id, tx)
		if err != nil {
			return err
		}
		respDataStsDetail = resultStsDetail

		// create data sts detail
		resultSumberDanaSts, err := sss.Create(dataSumberDanaSts, dataSts.Id, tx)
		if err != nil {
			return err
		}
		respDataSumberDanaSts = resultSumberDanaSts

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataSts)
	}

	resp = t.II{
		"sts":           dataSts,
		"stsDetail":     respDataStsDetail.(rp.OKSimple).Data,
		"sumberDanaSts": respDataSumberDanaSts.(rp.OKSimple).Data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.Sts
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Sts{}).
		Preload(clause.Associations).
		Preload("StsDetails.Rekening").
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data sts", data)
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

func GetDetail(sts_id int) (any, error) {
	var data *m.Sts
	err := a.DB.
		Model(&m.Sts{}).
		Preload(clause.Associations).
		Preload("StsDetails.Rekening").
		First(&data, sts_id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return rp.OKSimple{
		Data: data,
	}, err
}

func Update(id int, input m.UpdateDto) (any, error) {
	var dataSts *m.Sts
	var dataStsDetail []m.StsDetailUpdateDto
	var dataSumberDanaSts []m.SumberDanaStsUpdateDto
	var respDataStsDetail interface{}
	var respDataSumberDanaSts interface{}
	var resp t.II

	result := a.DB.First(&dataSts, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data sts tidak dapat ditemukan")
	}

	// data sts
	if err := sc.Copy(&dataSts, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload sts", input)
	}

	// data sts detail
	if err := sc.Copy(&dataStsDetail, &input.StsDetail); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload sts detail", input.StsDetail)
	}

	// data sumber dana sts
	if err := sc.Copy(&dataSumberDanaSts, &input.SumberDanaSts); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload sumber dana sts", input.SumberDanaSts)
	}

	if input.TanggalSetor == nil {
		dataSts.TanggalSetor = th.TimeNow()
	} else {
		dataSts.TanggalSetor = th.ParseTime(input.TanggalSetor)
	}

	if input.TanggalSts == nil {
		dataSts.TanggalSts = th.TimeNow()
	} else {
		dataSts.TanggalSts = th.ParseTime(input.TanggalSts)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// update sts
		if result := tx.Save(&dataSts); result.Error != nil {
			return errors.New("gagal menyimpan data sts")
		}

		// update sts detail
		resultStsDetail, err := ssd.Update(dataStsDetail, dataSts.Id, tx)
		if err != nil {
			return err
		}
		respDataStsDetail = resultStsDetail

		// update sumber dana sts
		resultSumberDanaSts, err := sss.Update(dataSumberDanaSts, dataSts.Id, tx)
		if err != nil {
			return err
		}
		respDataSumberDanaSts = resultSumberDanaSts

		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataSts)
	}

	resp = t.II{
		"affected":      strconv.Itoa(int(result.RowsAffected)),
		"stsDetail":     respDataStsDetail.(rp.OKSimple).Data,
		"sumberDanaSts": respDataSumberDanaSts.(rp.OKSimple).Data,
		"sts":           dataSts,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func Delete(sts_id uint64) (any, error) {
	var status string = "deleted"
	var data *m.Sts

	// cek data sspd
	result := a.DB.First(&data, sts_id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data sts tidak dapat ditemukan")
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// delete sts detail
		err := ssd.Delete(data.Id, tx)
		if err != nil {
			status = "no deletion"
			return err
		}

		// delete sumber dana sts
		err = sss.Delete(data.Id, tx)
		if err != nil {
			status = "no deletion"
			return err
		}

		// delete sts
		result = tx.Delete(&data, sts_id)
		if result.RowsAffected == 0 {
			return errors.New("tidak dapat menghapus data sts: " + result.Error.Error())
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
