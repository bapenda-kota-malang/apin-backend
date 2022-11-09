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
	"github.com/google/uuid"

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
	"gorm.io/gorm/clause"
)

const source = "spt"

func filePreProcess(b64String, docsname string, userId uint, oldId uuid.UUID) (fileName, path, extFile string, id uuid.UUID, err error) {
	extFile, err = base64helper.GetExtensionBase64(b64String)
	if err != nil {
		return
	}
	path = sh.GetResourcesPath()
	switch extFile {
	case "pdf":
		path = sh.GetPdfPath()
	case "png", "jpeg":
		path = sh.GetImgPath()
	case "xlsx", "xls":
		path = sh.GetExcelPath()
	default:
		err = errors.New("file tidak diketahui")
		return
	}
	if oldId == uuid.Nil {
		id, err = sh.GetUuidv4()
		if err != nil {
			err = errors.New("gagal generate uuid")
			return
		}
	} else {
		id = oldId
	}
	fileName = sh.GenerateFilename(docsname, id, userId, extFile)
	return
}

func Create(input m.CreateDto, opts map[string]interface{}, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.Spt

	// get data rekening
	respRek, err := srek.GetDetail(int(*input.Rekening_Id))
	if err != nil {
		return nil, err
	}
	dataRekening := respRek.(rp.OKSimple).Data.(*mrek.Rekening)

	// if kodeJenisPajak nil then search data rekening from parent id
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

	// if mark as new file not clone from esptpd then save new file
	if opts["newFile"].(bool) {
		var errChan = make(chan error)
		fileName, path, extFile, id, err := filePreProcess(input.Lampiran, "Lampiran"+opts["baseUri"].(string), opts["userId"].(uint), uuid.Nil)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		go sh.SaveFile(input.Lampiran, fileName, path, extFile, errChan)
		if err := <-errChan; err != nil {
			return sh.SetError("request", "create-data", source, "failed", "failed save pdf", data)
		}
		input.Lampiran = fileName
		data.Id = id
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
		data.CreateBy_User_Id = opts["userId"].(uint)
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

func GetList(input m.FilterDto, userId uint) (any, error) {
	var data []m.Spt
	var count int64

	var pagination gh.Pagination
	baseQuery := a.DB.Model(&m.Spt{})
	if userId != 0 {
		baseQuery.Joins("JOIN \"Npwpd\" ON \"Spt\".\"Npwpd_Id\" = \"Npwpd\".\"Id\" AND \"Npwpd\".\"User_Id\" = ?", userId)
	}
	if input.TbpStatus == nil {
		// baseQuery.Joins("JOIN \"Tbp\" ON \"Spt\".\"Id\"")
	} else {
		tbpStatus := *input.TbpStatus
		switch tbpStatus {
		case m.TbpStatusFilterBaru:
		case m.TbpStatusFilterLunas:
		case m.TbpStatusFilterJatuhTempo:
		}
	}
	input.TbpStatus = nil
	if input.JatuhTempo != nil {
		val, _ := input.JatuhTempo.Value()
		endOfMonthInput := datatypes.Date(sh.EndOfMonth(val.(time.Time)))
		input.JatuhTempo = &endOfMonthInput
	}
	result := baseQuery.
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

func GetDetail(id uuid.UUID, userId uint) (any, error) {
	var data *m.Spt

	baseQuery := a.DB.Model(&m.Spt{})
	if userId != 0 {
		baseQuery.Joins("JOIN \"Npwpd\" ON \"Spt\".\"Npwpd_Id\" = \"Npwpd\".\"Id\" AND \"Npwpd\".\"User_Id\" = ?", userId)
	}
	result := baseQuery.
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		First(&data, "\"Id\" = ?", id.String())
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id uuid.UUID, input any, opts map[string]interface{}, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.Spt
	// validate data exist and copy input (payload) ke struct data jika tidak ada akan error
	dataRow := tx.First(&data, "\"Id\" = ?", id.String()).RowsAffected
	if dataRow == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	// copy to model struct
	if inputData, ok := input.(m.VerifyDto); ok {
		// if data.StatusPenetapan == m.StatusPenetapanDisetujuiKabid {
		// 	return sh.SetError("request", "update-data", source, "failed", "data telah disetujui", data)
		// }
		if err := sc.Copy(&data, &inputData); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
		}
		switch data.StatusPenetapan {
		case m.StatusPenetapanBaru,
			m.StatusPenetapanDisetujuiKabid,
			m.StatusPenetapanDisetujuiKasubid,
			m.StatusPenetapanDitolakKabid,
			m.StatusPenetapanDitolakKasubid:
			// do nothing
		default:
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
		}
		// data.VerifyBy_User_Id = &userId
	} else {
		inputData := input.(m.UpdateDto)
		if inputData.Lampiran != nil {
			var errChan = make(chan error)
			fileName, path, extFile, _, err := filePreProcess(*inputData.Lampiran, "Lampiran"+opts["baseUri"].(string), opts["userId"].(uint), data.Id)
			if err != nil {
				return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
			}
			go sh.ReplaceFile(data.Lampiran, *inputData.Lampiran, fileName, path, extFile, errChan)
			if err := <-errChan; err != nil {
				return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("failed save file: %s", err), data)
			}
			inputData.Lampiran = &fileName
		}
		if err := sc.Copy(&data, &inputData); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
		}
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(dataRow)),
		},
		Data: data,
	}, nil
}

func Delete(id uuid.UUID) (any, error) {
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
