package dbkbjpb16

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

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb16"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "dbkbjpb16"

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.DbkbJpb16

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data dbkbjpb16", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.DbkbJpb16
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.DbkbJpb16{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data dbkbjpb16", data)
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
	var data *m.DbkbJpb16

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data dbkbjpb16", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.DbkbJpb16
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data dbkbjpb16", data)
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
	var data *m.DbkbJpb16
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
	var data []m.DbkbJpb16

	result := a.DB.
		Model(&m.DbkbJpb16{}).
		Scopes(gh.Filter(input)).
		Find(&data)
	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data dbkbjpb16", data)
		return nil, err
	}

	var excelData = func() []interface{} {
		var tmp []interface{}
		tmp = append(tmp, map[string]interface{}{
			"A": "No",
			"B": "Kode Prov.",
			"C": "Kode Kota",
			"D": "Tahun",
			"E": "Kelas",
			"F": "Min",
			"G": "Max",
			"H": "Nilai",
		})
		for i, v := range data {
			n := i + 1
			tmp = append(tmp, map[string]interface{}{
				"A": n,
				"B": *v.Provinsi_Kode,
				"C": *v.Daerah_Kode,
				"D": func() string {
					if v.TahunDbkbJpb16 != nil {
						return *v.TahunDbkbJpb16
					}
					return ""
				}(),
				"E": func() string {
					if v.KelasDbkbJpb16 != nil {
						return *v.KelasDbkbJpb16
					}
					return ""
				}(),
				"F": func() string {
					if v.LantaiMinJpb16 != nil {
						return strconv.Itoa(*v.LantaiMinJpb16)
					}
					return ""
				}(),
				"G": func() string {
					if v.LantaiMaxJpb16 != nil {
						return strconv.Itoa(*v.LantaiMaxJpb16)
					}
					return ""
				}(),
				"H": func() string {
					if v.NilaiDbkbJpb16 != nil {
						s := strconv.FormatFloat(*v.NilaiDbkbJpb16, 'f', 0, 64)
						return s
					}
					return ""
				}(),
			})
		}
		return tmp
	}()
	return excelhelper.ExportList(excelData, "List")
}
