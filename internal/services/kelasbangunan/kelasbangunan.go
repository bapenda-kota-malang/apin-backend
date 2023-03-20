package kelasbangunan

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/excelhelper"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/kelasbangunan"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "kelas bangunan"

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.KelasBangunan

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data kelas bangunan", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.KelasBangunan
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.KelasBangunan{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data kelas bangunan", data)
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
	var data *m.KelasBangunan

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data kelas bangunan", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func GetDetailByCode(kd string) (interface{}, error) {
	var data *m.KelasBangunan
	result := a.DB.Where(&m.KelasBangunan{
		KdBangunan: kd,
	}).Order("\"Id\" desc").First(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.KelasBangunan
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data kelas bangunan", data)
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
	var data *m.KelasBangunan
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

func DownloadExcelList(input m.FilterDto) (*excelize.File, error) {
	var data []m.KelasBangunan

	result := a.DB.
		Model(&m.KelasBangunan{}).
		Scopes(gh.Filter(input)).
		Find(&data)
	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data kelas bangunan", data)
		return nil, err
	}

	var excelData = func() []interface{} {
		var tmp []interface{}
		tmp = append(tmp, map[string]interface{}{
			"A": "No",
			"B": "Kode",
			"C": "Tahun Awal",
			"D": "Tahun Akhir",
			"E": "Nilai Min",
			"F": "Nilai Max",
			"G": "Nilai /m2",
		})
		for i, v := range data {
			n := i + 1
			tmp = append(tmp, map[string]interface{}{
				"A": n,
				"B": v.KdBangunan,
				"C": func() string {
					if v.TahunAwalKelasBangunan != nil {
						return *v.TahunAwalKelasBangunan
					}
					return ""
				}(),
				"D": func() string {
					if v.TahunAkhirKelasBangunan != nil {
						return *v.TahunAkhirKelasBangunan
					}
					return ""
				}(),
				"E": func() string {
					if v.NilaiMinBangunan != nil {
						return strconv.FormatFloat(*v.NilaiMinBangunan, 'f', 0, 64)
					}
					return ""
				}(),
				"F": func() string {
					if v.NilaiMaxBangunan != nil {
						return strconv.FormatFloat(*v.NilaiMaxBangunan, 'f', 0, 64)
					}
					return ""
				}(),
				"G": func() string {
					if v.NilaiPerM2Bangunan != nil {
						return strconv.FormatFloat(*v.NilaiPerM2Bangunan, 'f', 0, 64)
					}
					return ""
				}(),
			})
		}
		return tmp
	}()
	return excelhelper.ExportList(excelData, "List")
}
