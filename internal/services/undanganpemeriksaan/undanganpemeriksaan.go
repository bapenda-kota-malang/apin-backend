package undanganpemeriksaan

import (
	"errors"
	"fmt"
	"strconv"

	sc "github.com/jinzhu/copier"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/excelhelper"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/undanganpemeriksaan"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
)

const source = "undangan pemeriksaan"

func Create(input m.CreateDto, user_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.UndanganPemeriksaan

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}
	data.Tanggal = th.ParseTime(input.Tanggal)
	data.Pukul = parseCurrentTime(input.Pukul)
	data.Petugas_User_Id = &user_id
	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data undangan pemeriksaan", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.UndanganPemeriksaan
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.UndanganPemeriksaan{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Npwpd.ObjekPajak").
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data undangan pemeriksaan", data)
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

func DownloadExcelList(input m.FilterDto) (*excelize.File, error) {
	var data []m.UndanganPemeriksaan
	result := a.DB.
		Model(&m.UndanganPemeriksaan{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Npwpd.ObjekPajak").
		Scopes(gh.Filter(input)).
		Find(&data)
	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data undangan pemeriksaan", data)
		return nil, err
	}

	var excelData = func() []interface{} {
		var tmp []interface{}
		tmp = append(tmp, map[string]interface{}{
			"A": "No",
			"B": "NPWD",
			"C": "Nama Usaha",
			"D": "No Surat Pemberitahuan",
			"E": "Tanggal Pemeriksaan",
			"F": "Status",
		})
		for i, v := range data {
			n := i + 1
			tmp = append(tmp, map[string]interface{}{
				"A": n,
				"B": func() string {
					if v.Npwpd != nil && v.Npwpd.Npwpd != nil {
						return *v.Npwpd.Npwpd
					}
					return ""
				}(),
				"C": func() string {
					if v.Npwpd != nil && v.Npwpd.ObjekPajak != nil {
						return *v.Npwpd.ObjekPajak.Nama
					}
					return ""
				}(),
				"D": *v.NoSuratUndangan,
				"E": func() string {
					if v.Tanggal != nil {
						return v.Tanggal.Format("2006-01-02")
					}
					return ""
				}(),
				"F": int(v.Status),
			})
		}
		return tmp
	}()
	return excelhelper.ExportList(excelData, "List")
}

func GetDetail(id int) (any, error) {
	var data *m.UndanganPemeriksaan

	result := a.DB.Model(&m.UndanganPemeriksaan{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Npwpd.ObjekPajak").
		First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data undangan pemeriksaan", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, user_id uint64, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.UndanganPemeriksaan
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}
	data.Tanggal = th.ParseTime(input.Tanggal)
	data.Pukul = parseCurrentTime(input.Pukul)
	data.Petugas_User_Id = &user_id
	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data undangan pemeriksaan", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.UndanganPemeriksaan
	result := tx.First(&data, id)
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

func UpdateStatusTerbit(input m.UpdateStatusTerbitDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var dataResult []m.UndanganPemeriksaan
	for _, v := range input.Id {
		var data m.UndanganPemeriksaan
		result := tx.First(&data, v)
		if result.RowsAffected == 0 {
			return nil, fmt.Errorf("data undangan pemeriksaan dengan nomor id %d tidak ditemukan", v)
		}
		data.Status = m.StatusTerbit
		if result := tx.Save(&data); result.Error != nil {
			return nil, errors.New("gagal menyimpan data undangan pemeriksaan: " + result.Error.Error())
		}
		dataResult = append(dataResult, data)
	}
	return rp.OKSimple{Data: dataResult}, nil
}
