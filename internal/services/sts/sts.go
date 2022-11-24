package sts

import (
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
		Preload("StsDetail.Rekening").
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
		Preload("StsDetail.Rekening").
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
