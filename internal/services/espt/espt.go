package espt

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	sc "github.com/jinzhu/copier"
	"github.com/xuri/excelize/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/base64helper"
	"github.com/bapenda-kota-malang/apin-backend/pkg/excelhelper"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	sspt "github.com/bapenda-kota-malang/apin-backend/internal/services/spt"

	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "espt"

func filePreProcess(b64String string, userId uint, oldId uuid.UUID) (fileName, path, extFile string, id uuid.UUID, err error) {
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
	fileName = sh.GenerateFilename("AttachmentEsptpd", id, userId, extFile)
	return
}

func GetList(input m.FilterDto, user_Id uint) (any, error) {
	var data []m.Espt
	var count int64

	var pagination gh.Pagination
	baseQuery := a.DB.Model(&m.Espt{}).
		Joins("Rekening")
	if user_Id != 0 {
		baseQuery = baseQuery.Where(&m.Espt{LaporBy_User_Id: user_Id})
	}

	result := baseQuery.
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Joins("Npwpd").
		// Joins("ObjekPajak").
		Preload("LaporUser", func(db *gorm.DB) *gorm.DB {
			return db.Omit("Password")
		}).
		// Joins("Rekening").
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

func DownloadExcelList(input m.FilterDto, user_Id uint) (*excelize.File, error) {
	var data []m.Espt

	baseQuery := a.DB.Model(&m.Espt{})
	if user_Id != 0 {
		baseQuery = baseQuery.Where(&m.Espt{LaporBy_User_Id: user_Id})
	}

	result := baseQuery.
		Scopes(gh.Filter(input)).
		Joins("Npwpd").
		// Joins("ObjekPajak").
		Preload("LaporUser", func(db *gorm.DB) *gorm.DB {
			return db.Omit("Password")
		}).
		Joins("Rekening").
		Find(&data)
	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
		return nil, err
	}

	var excelData = func() []interface{} {
		var tmp []interface{}
		tmp = append(tmp, map[string]interface{}{
			"A": "No",
			"B": "Nomor",
			"C": "Tanggal",
			"D": "Masa Pajak",
			"E": "Jatuh Tempo",
			"F": "Pajak/Retribusi",
			"G": "NPWPD",
			"H": "Nama WP",
			"I": "Jumlah Pajak",
			"J": "Status",
		})
		for i, v := range data {
			n := i + 1
			tmp = append(tmp, map[string]interface{}{
				"A": n,
				"B": func() string {
					s := strings.Split(v.Id.String(), "-")
					return s[4]
				}(),
				"C": func() string {
					return v.CreatedAt.Format("2006-01-02")
				}(),
				"D": func() string {
					// start, _ := time.Parse("2006-01-02", v.PeriodeAwal.GormDataType())
					return v.PeriodeAwal.GormDataType() + " s/d " + v.PeriodeAkhir.GormDataType()
				}(),
				"E": func() string {
					return v.JatuhTempo.GormDataType()
				}(),
				"F": "",
				"G": *v.Npwpd.Npwpd,
				"H": v.LaporUser.Name,
				"I": v.TarifPajak_Id,
				"J": v.VerifyStatus,
			})
		}
		return tmp
	}()
	return excelhelper.ExportList(excelData, "List")
}

func GetDetail(id uuid.UUID, user_Id uint) (any, error) {
	var data *m.Espt

	baseQuery := a.DB.
		Preload("ObjekPajak.Kecamatan").
		Preload("ObjekPajak.Kelurahan").
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		})
	if user_Id != 0 {
		baseQuery = baseQuery.Where(&m.Espt{LaporBy_User_Id: user_Id})
	}

	result := baseQuery.First(&data, "\"Id\" = ?", id.String())
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

// Service insert data for table espt:
//
//	Create(input, tx) // for using transaction
//	Create(input, nil) // for not using transaction
func Create(input m.CreateDto, user_Id uint, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.Espt
	var errChan = make(chan error)

	fileName, path, extFile, id, err := filePreProcess(input.Attachment, user_Id, uuid.Nil)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
	}

	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// check relasi id tabel ke tabel relasi
	// espt table
	// if result := a.DB.First(&mr.Rekening{}, dataPotensiOp.Rekening_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }

	// static add value to field
	data.Id = id
	go sh.SaveFile(data.Attachment, fileName, path, extFile, errChan)
	prevMonth := sh.BeginningOfPreviosMonth()
	data.PeriodeAwal = datatypes.Date(prevMonth)
	data.PeriodeAkhir = datatypes.Date(sh.EndOfMonth(prevMonth))
	data.JatuhTempo = datatypes.Date(sh.EndOfMonth(time.Now()))
	data.VerifyStatus = m.StatusBaru
	data.LaporBy_User_Id = user_Id
	data.Attachment = fileName

	// save to db
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	if err := <-errChan; err != nil {
		return sh.SetError("request", "create-data", source, "failed", "failed save file", data)
	}

	return rp.OKSimple{Data: data}, nil
}

// Service update data for table espt
func Update(ctx context.Context, id uuid.UUID, input any, userId uint, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.Espt

	// validate data exist and copy input (payload) ke struct data jika tidak ada akan error
	dataRow := tx.
		Preload("DetailEsptAir").
		Preload("DetailEsptHiburan").
		Preload("DetailEsptHotel").
		Preload("DetailEsptParkir").
		Preload("DetailEsptPpjNonPln").
		Preload("DetailEsptPpjPln").
		Preload("DetailEsptResto").
		First(&data, "\"Id\" = ?", id.String()).RowsAffected
	if dataRow == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	// copy to model struct
	if inputData, ok := input.(m.VerifyDto); ok {
		if data.VerifyStatus == m.StatusDisetujui {
			return sh.SetError("request", "update-data", source, "failed", "data telah disetujui", data)
		}
		if err := sc.Copy(&data, &inputData); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
		}
		data.VerifyBy_User_Id = &userId
	} else {
		inputData := input.(m.UpdateDto)
		if inputData.Attachment != nil {
			var errChan = make(chan error)
			fileName, path, extFile, _, err := filePreProcess(*inputData.Attachment, userId, data.Id)
			if err != nil {
				return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
			}
			go sh.ReplaceFile(data.Attachment, *inputData.Attachment, fileName, path, extFile, errChan)
			if err := <-errChan; err != nil {
				return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("failed save file: %s", err), data)
			}
			inputData.Attachment = &fileName
		}
		if err := sc.Copy(&data, &inputData); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
		}
	}

	// check relasi id tabel ke tabel relasi
	// potensiop table
	// if result := tx.First(&mr.Rekening{}, dataPotensiOp.Rekening_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }

	switch data.VerifyStatus {
	case m.StatusBaru, m.StatusDisetujui, m.StatusDitolak:
		// do nothing
	default:
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	err := tx.Transaction(func(tx *gorm.DB) error {
		if result := tx.Save(&data); result.Error != nil {
			return result.Error
		}

		// if verify, clone data to spt table with detail
		if _, ok := input.(m.VerifyDto); ok && data.VerifyStatus == m.StatusDisetujui {
			inputSpt, err := sspt.TransformEspt(&data)
			if err != nil {
				return err
			}

			opts := make(map[string]interface{})
			opts["userId"] = uint(userId)
			opts["newFile"] = false
			opts["baseUri"] = "sptpd"

			_, err = sspt.CreateDetail(ctx, inputSpt, opts, tx)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal memperbarui data: "+err.Error(), data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(dataRow)),
		},
		Data: data,
	}, nil
}

// Service soft delete data for table espt
func Delete(id uuid.UUID) (any, error) {
	var data *m.Espt
	result := a.DB.First(&data, "\"Id\" = ?", id.String())
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	result = a.DB.Where("\"Id\" = ?", id.String()).Delete(&data)
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
