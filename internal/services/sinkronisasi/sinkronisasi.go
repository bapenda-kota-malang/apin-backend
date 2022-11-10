package sinkronisasi

import (
	"fmt"
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sinkronisasi"
	srs "github.com/bapenda-kota-malang/apin-backend/internal/services/sinkronisasi/rinciansinkronisasi"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "sinkronisasi"

func GetList(input m.FilterDto) (any, error) {
	var data []m.Sinkronisasi
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Sinkronisasi{}).
		Preload(clause.Associations).
		Preload("User").
		Preload("Tbp").
		Preload("RincianSinkronisasi.RincianTbp").
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

func GetDetail(tbp_id int) (any, error) {
	var data *m.Sinkronisasi
	err := a.DB.
		Model(&m.Sinkronisasi{}).
		Preload(clause.Associations).
		Preload("User").
		Preload("Tbp").
		Preload("RincianSinkronisasi.RincianTbp").
		First(&data, tbp_id).Error
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

func Create(input m.CreateDto, user_Id uint64) (any, error) {
	var dataSinkronisasi m.Sinkronisasi
	var dataRincianSinkronisasi m.RincianSinkronisasiCreateDto
	var respDataRincianSinkronisasi interface{}
	var resp t.II
	var errChan = make(chan error)
	baseDocsName := "sinkronisasi"

	fileName, path, extFile, err := filePreProcess(input.File, uint(user_Id), baseDocsName+"File")
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}

	go sh.SaveFile(input.File, fileName, path, extFile, errChan)
	if err := <-errChan; err != nil {
		return sh.SetError("request", "create-data", source, "failed", "excel unsupported", input)
	}

	// data sinkronisasi
	if err := sc.Copy(&dataSinkronisasi, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload sinkronisasi", dataSinkronisasi)
	}

	// data rincian sinkronisasi
	if err := sc.Copy(&dataRincianSinkronisasi, &input.RincianSinkronisasi); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload rincian sinkronisasi", dataRincianSinkronisasi)
	}

	dataSinkronisasi.File = fileName
	dataSinkronisasi.User_Id = &user_Id

	if dataSinkronisasi.JenisPajak == "pdl" {
		dataExcel := readExcelFile(fileName)
		fmt.Println("data excel: ", dataExcel)
	}

	err = a.DB.Transaction(func(tx *gorm.DB) error {

		// create data sinkronisasi
		err := tx.Create(&dataSinkronisasi).Error
		if err != nil {
			return err
		}

		// create data rincian sinkronisasi
		resultRincianSinkronisasi, err := srs.Create(dataRincianSinkronisasi, dataSinkronisasi.Id, tx)
		if err != nil {
			return err
		}
		respDataRincianSinkronisasi = resultRincianSinkronisasi

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataSinkronisasi)
	}

	resp = t.II{
		"sinkronisasi":        dataSinkronisasi,
		"rincianSinkronisasi": respDataRincianSinkronisasi.(rp.OKSimple).Data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}
