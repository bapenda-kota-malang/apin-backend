package spt

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	mrek "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mnomertracker "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/sptnomertracker"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"

	srek "github.com/bapenda-kota-malang/apin-backend/internal/services/configuration/rekening"
	snomertracker "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/sptnomertracker"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/base64helper"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const source = "spt"

func Create(input m.CreateDto, user_Id uint, newFile bool, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.Spt

	respRek, err := srek.GetDetail(int(*input.Rekening_Id))
	if err != nil {
		return nil, err
	}

	dataRekening := respRek.(rp.OKSimple).Data.(*mrek.Rekening)

	kodeJenisPajak := dataRekening.KodeJenisPajak
	if kodeJenisPajak == nil {
		parentId, err := strconv.Atoi(*dataRekening.Parent_Id)
		if err != nil {
			return nil, err
		}
		respRek, err := srek.GetDetail(parentId)
		if err != nil {
			return nil, err
		}
		dataRekening = respRek.(rp.OKSimple).Data.(*mrek.Rekening)
		kodeJenisPajak = dataRekening.KodeJenisPajak
	}

	if newFile {
		var errChan = make(chan error)
		extFile, err := base64helper.GetExtensionBase64(input.Lampiran)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		path := sh.GetResourcesPath()
		switch extFile {
		case "pdf":
			path = sh.GetPdfPath()
		case "png", "jpeg":
			path = sh.GetImgPath()
		case "xlsx", "xls":
			path = sh.GetExcelPath()
		}
		id, err := sh.GetUuidv4()
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal generate id", data)
		}
		fileName := sh.GenerateFilename("LampiranSptpd", id, user_Id, extFile)
		go sh.SaveFile(data.Lampiran, fileName, path, extFile, errChan)
		if err := <-errChan; err != nil {
			return sh.SetError("request", "create-data", source, "failed", "failed save pdf", data)
		}
		input.Lampiran = fileName
	}

	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	data.TanggalSpt = time.Now()
	yearNow := uint(data.TanggalSpt.Year())

	nomerUrut, err := snomertracker.TrxNewNumber(mnomertracker.Dto{Tahun: &yearNow, KodeJenisPajak: kodeJenisPajak}, tx)
	if err != nil {
		return nil, err
	}
	kode := *kodeJenisPajak
	yearTwoDigit := yearNow % 1e2
	nomerSpt := fmt.Sprintf("%d%s", yearTwoDigit, nomerUrut)
	data.NomorSpt = fmt.Sprintf("%c-%s", kode[0], nomerSpt)
	data.KodeBilling = fmt.Sprintf("%s%s", *dataRekening.KodeBilling, nomerSpt)
	switch data.Type {
	case mtypes.JenisPajakOA, mtypes.JenisPajakSA:
	default:
		data.Type = mtypes.JenisPajakSA
	}
	data.Status = m.StatusBelumLunas
	if data.CreateBy_User_Id == 0 {
		data.CreateBy_User_Id = user_Id
	}
	if data.PeriodeAwal == datatypes.Date(time.Time{}) {
		prevMonth := sh.BeginningOfPreviosMonth()
		data.PeriodeAwal = datatypes.Date(prevMonth)
		data.PeriodeAkhir = datatypes.Date(sh.EndOfMonth(prevMonth))
		data.JatuhTempo = datatypes.Date(sh.EndOfMonth(time.Now()))
	}

	err = tx.Create(&data).Error
	if err != nil {
		return nil, err
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.Spt
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Spt{}).
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
	var data *m.Spt

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

func Update(id int, input m.UpdateDto) (any, error) {
	var data *m.Spt
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
	//data spt
	var data *m.Spt
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	status := "deleted"

	result = a.DB.Delete(&data, id)
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
