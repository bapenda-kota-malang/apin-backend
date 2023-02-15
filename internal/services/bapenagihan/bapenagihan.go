package bapenagihan

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
	sc "github.com/jinzhu/copier"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/base64helper"
	"github.com/bapenda-kota-malang/apin-backend/pkg/excelhelper"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bapenagihan"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "bapenagihan"

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

func Create(input m.CreateDto, userId uint, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.BaPenagihan

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	id, err := sh.GetUuidv4()
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal membuat uuid", data)
	}
	data.Id = id

	// dokumentasi files
	dokumentasiFilename := []string{}
	for i, v := range input.Dokumentasi {
		var errChan = make(chan error)
		fileName, path, extFile, _, err := filePreProcess(v, "BaPenagihanDokumentasi"+strconv.Itoa(i+1), userId, data.Id)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		go sh.SaveFile(v, fileName, path, extFile, errChan)
		if err := <-errChan; err != nil {
			return sh.SetError("request", "create-data", source, "failed", "failed save file: "+err.Error(), data)
		}
		dokumentasiFilename = append(dokumentasiFilename, fileName)
	}
	byteDokumentasi, err := json.Marshal(dokumentasiFilename)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal membuat json dokumentasi", data)
	}
	data.Dokumentasi = byteDokumentasi

	// dokumen lain lain files
	dokumenLainLainFilename := []string{}
	for i, v := range input.DokumenLainLain {
		var errChan = make(chan error)
		fileName, path, extFile, _, err := filePreProcess(v, "BaPenagihanDokumenLainLain"+strconv.Itoa(i+1), userId, data.Id)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		go sh.SaveFile(v, fileName, path, extFile, errChan)
		if err := <-errChan; err != nil {
			return sh.SetError("request", "create-data", source, "failed", "failed save file: "+err.Error(), data)
		}
		dokumenLainLainFilename = append(dokumenLainLainFilename, fileName)
	}
	byteDokumenLainLain, err := json.Marshal(dokumenLainLainFilename)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal membuat json dokumen lain lain", data)
	}
	data.DokumenLainLain = byteDokumenLainLain

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
	var data []m.BaPenagihan
	var count int64

	var pagination gh.Pagination
	result := tx.
		Model(&m.BaPenagihan{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Preload("Undangan.Npwpd.ObjekPajak").
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

func DownloadExcelList(input m.FilterDto, tx *gorm.DB) (*excelize.File, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.BaPenagihan
	result := tx.
		Model(&m.BaPenagihan{}).
		Scopes(gh.Filter(input)).
		Preload("Undangan.Npwpd.ObjekPajak").
		Find(&data)
	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
		return nil, err
	}

	var excelData = func() []interface{} {
		var tmp []interface{}
		tmp = append(tmp, map[string]interface{}{
			"A": "No",
			"B": "NPWD",
			"C": "Nama Usaha",
			"D": "No Surat Undangan",
			"E": "Tanggal Pemeriksaan",
			"F": "Verifikasi Kasubid",
		})
		for i, _ := range data {
			n := i + 1
			tmp = append(tmp, map[string]interface{}{
				"A": n,
				"B": "",
				"C": "",
				"D": "",
				"E": "",
				"F": "",
			})
		}
		return tmp
	}()
	return excelhelper.ExportList(excelData, "List")
}

func GetDetail(id uuid.UUID, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.BaPenagihan

	result := tx.
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("BaPenagihanPetugas.Petugas", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Undangan.Npwpd.ObjekPajak").
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

func Update(id uuid.UUID, input m.UpdateDto, userId uint, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	data := m.BaPenagihan{}
	result := tx.First(&data, "\"Id\" = ?", id.String())
	if result.RowsAffected == 0 {
		return nil, nil
	}

	if err := sc.CopyWithOption(&data, &input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	// dokumentasi files
	if len(input.Dokumentasi) > 0 {
		dokumentasiFilename := []string{}
		lenData := 0
		// check from data existing and get last number file
		if data.Dokumentasi != nil {
			err := json.Unmarshal(data.Dokumentasi, &dokumentasiFilename)
			if err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal unmarshal json", data)
			}
			re := regexp.MustCompile(`^BaPenagihanDokumentasi(\d*)`)
			if match := re.FindStringSubmatch(dokumentasiFilename[len(dokumentasiFilename)-1]); len(match) > 0 {
				lenData, _ = strconv.Atoi(match[1])
			}
		}
		// sort delete data last
		for i := 0; i < len(input.Dokumentasi); i++ {
			if input.Dokumentasi[i].Deleted && (i+1 == len(input.Dokumentasi)) {
				break
			}
			tmp := input.Dokumentasi[i]
			input.Dokumentasi = append(input.Dokumentasi[:i], input.Dokumentasi[i+1:]...)
			if tmp.Deleted {
				input.Dokumentasi = append(input.Dokumentasi, tmp)
			} else {
				input.Dokumentasi = append([]m.FilesUpdateDokumentasi{tmp}, input.Dokumentasi...)
			}
		}
		for i, v := range input.Dokumentasi {
			// if deleted flag then delete from slice
			if v.Deleted {
				if v.OldFile == nil {
					return sh.SetError("request", "update-data", source, "failed", "old file empty when try to delete", data)
				}
				for i := range dokumentasiFilename {
					dokFile := dokumentasiFilename[i]
					oldFile := *v.OldFile
					fmt.Println(dokFile == oldFile)
					if dokumentasiFilename[i] == *v.OldFile {
						dokumentasiFilename = append(dokumentasiFilename[:i], dokumentasiFilename[i+1:]...)
						path := sh.GetPathByFilename(*v.OldFile)
						if err := sh.RemoveFile(path, *v.OldFile); err != nil {
							return sh.SetError("request", "update-data", source, "failed", "failed remove file", v)
						}
						break
					}
				}
				continue
			}
			if v.NewFile == nil {
				return sh.SetError("request", "update-data", source, "failed", "new file empty", data)
			}
			// flag not delete assume append new file
			errChan := make(chan error)
			fileName, path, extFile, _, err := filePreProcess(*v.NewFile, "BaPenagihanDokumentasi"+strconv.Itoa(i+lenData+1), userId, data.Id)
			if err != nil {
				return sh.SetError("request", "update-data", source, "failed", err.Error(), data)
			}
			go sh.SaveFile(*v.NewFile, fileName, path, extFile, errChan)
			if err := <-errChan; err != nil {
				return sh.SetError("request", "update-data", source, "failed", "failed save file: "+err.Error(), data)
			}
			dokumentasiFilename = append(dokumentasiFilename, fileName)
		}
		byteDokumentasi, err := json.Marshal(dokumentasiFilename)
		if err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal membuat json dokumentasi", data)
		}
		data.Dokumentasi = byteDokumentasi
	}

	// dokumen lain lain files
	if len(input.DokumenLainLain) > 0 {
		dokumenLainLainFilename := []string{}
		lenData := 0
		// check from data existing and get last number file
		if data.DokumenLainLain != nil {
			err := json.Unmarshal(data.DokumenLainLain, &dokumenLainLainFilename)
			if err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal unmarshal json", data)
			}
			re := regexp.MustCompile(`^BaPenagihanDokumenLainLain(\d*)`)
			if match := re.FindStringSubmatch(dokumenLainLainFilename[len(dokumenLainLainFilename)-1]); len(match) > 0 {
				lenData, _ = strconv.Atoi(match[1])
			}
		}
		// sort delete data last
		for i := 0; i < len(input.DokumenLainLain); i++ {
			if input.DokumenLainLain[i].Deleted && (i+1 == len(input.DokumenLainLain)) {
				break
			}
			tmp := input.DokumenLainLain[i]
			input.DokumenLainLain = append(input.DokumenLainLain[:i], input.DokumenLainLain[i+1:]...)
			if tmp.Deleted {
				input.DokumenLainLain = append(input.DokumenLainLain, tmp)
			} else {
				input.DokumenLainLain = append([]m.FilesUpdateDokumenLainLain{tmp}, input.DokumenLainLain...)
			}
		}
		for i, v := range input.DokumenLainLain {
			// if deleted flag then delete from slice
			if v.Deleted {
				if v.OldFile == nil {
					return sh.SetError("request", "update-data", source, "failed", "old file empty when try to delete", data)
				}
				for i := range dokumenLainLainFilename {
					if dokumenLainLainFilename[i] == *v.OldFile {
						dokumenLainLainFilename = append(dokumenLainLainFilename[:i], dokumenLainLainFilename[i+1:]...)
						path := sh.GetPathByFilename(*v.OldFile)
						if err := sh.RemoveFile(path, *v.OldFile); err != nil {
							return sh.SetError("request", "update-data", source, "failed", "failed remove file", v)
						}
						break
					}
				}
				continue
			}
			if v.NewFile == nil {
				return sh.SetError("request", "update-data", source, "failed", "new file empty", data)
			}
			errChan := make(chan error)
			fileName, path, extFile, _, err := filePreProcess(*v.NewFile, "BaPenagihanDokumenLainLain"+strconv.Itoa(i+lenData+1), userId, data.Id)
			if err != nil {
				return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
			}
			go sh.SaveFile(*v.NewFile, fileName, path, extFile, errChan)
			if err := <-errChan; err != nil {
				return sh.SetError("request", "create-data", source, "failed", "failed save file: "+err.Error(), data)
			}
			dokumenLainLainFilename = append(dokumenLainLainFilename, fileName)
		}
		byteDokumenLainLain, err := json.Marshal(dokumenLainLainFilename)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal membuat json dokumen lain lain", data)
		}
		data.DokumenLainLain = byteDokumenLainLain
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

func Verify(id uuid.UUID, input m.VerifikasiDto, userId uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	data := m.BaPenagihan{}
	// cek user id apakah pegawai kasubid
	// resp, err := suser.GetJabatanPegawai(uint(userId))
	// if err != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal data pegawai: "+err.Error(), nil)
	// }
	// if kasubid := strings.Contains(strings.ToUpper(resp.(string)), "KEPALA SUB BIDANG"); !kasubid {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal data pegawai: pegawai bukan kasubid", nil)
	// }
	result := tx.First(&data, "\"Id\" = ?", id.String())
	if result.RowsAffected == 0 {
		return nil, nil
	}

	if err := sc.CopyWithOption(&data, &input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}
	data.VerifyBy_User_Id = &userId
	now := time.Now()
	data.VerifyAt = &now

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

func Delete(id uuid.UUID, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.BaPenagihan
	result := tx.First(&data, "\"Id\" = ?", id.String())
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
