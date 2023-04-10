package pelaporanppat

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	sc "github.com/jinzhu/copier"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/bapenda-kota-malang/apin-backend/pkg/pdf"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	msptpd "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pelaporanppat"
	mppat "github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "pelaporan ppat"

var months = []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.PelaporanPpat

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data pelaporan ppat", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.PelaporanPpat
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.PelaporanPpat{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data pelaporan ppat", data)
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
	var data *m.PelaporanPpat

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data pelaporan ppat", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func GetDetailbyField(field string, value string) (any, error) {
	var data *[]m.PelaporanPpat

	result := a.DB.Where(field, value).Find(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func GetListTransaksiPPAT(input msptpd.FilterPPATDto) (any, error) {
	var data []m.ResponsePelaporanPpat
	var count int64

	queryBase := a.DB.Model(&m.PelaporanPpat{})
	queryBase = queryBase.
		Select("DISTINCT ON (\"PelaporanPpat\".\"Ppat_Id\") \"PelaporanPpat\".\"Ppat_Id\"", "(\"Ppat\".\"Nama\") \"Ppat_Name\"", "\"PelaporanPpat\".\"TglLapor\"", "count(\"PelaporanPpat\".\"Sptpd_Id\") \"Sptpd_Id\"", "sum(\"BphtbSptpd\".\"NilaiOp\") \"NilaiOp\"", "sum(\"BphtbSptpd\".\"JumlahSetor\") \"JumlahSetor\"", "\"BphtbSptpd\".\"Status\"").
		Joins("LEFT JOIN \"BphtbSptpd\" ON \"BphtbSptpd\".\"Sptpd_Id\" = \"PelaporanPpat\".\"Sptpd_Id\"").
		Joins("LEFT JOIN \"Ppat\" ON \"Ppat\".\"Id\" = CAST(\"PelaporanPpat\".\"Ppat_Id\" AS INTEGER) ")

	if input.Ppat_Id != nil {
		queryBase = queryBase.Where("\"PelaporanPpat\".\"Ppat_Id\"", input.Ppat_Id)
	}

	if input.Bulan != nil {
		queryBase = queryBase.Where("EXTRACT('month' from \"TglLapor\") = ?", input.Bulan)
	}

	if input.Tahun != nil {
		queryBase = queryBase.Where("EXTRACT('year' from \"TglLapor\") = ?", input.Tahun)
	}

	result := queryBase.
		Order("\"PelaporanPpat\".\"Ppat_Id\"").
		Group("\"PelaporanPpat\".\"Ppat_Id\", \"Ppat\".\"Nama\", \"PelaporanPpat\".\"TglLapor\", \"BphtbSptpd\".\"Status\"").
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func GetDetailTransaksiPPAT(input msptpd.FilterPPATDto) (any, error) {
	var data []m.ResponseDetilPelaporanPpat
	var dataPPAT mppat.Ppat
	var count int64
	var dateString string
	var namePPAT string

	queryBase := a.DB.Model(&m.PelaporanPpat{})
	queryBase = queryBase.
		Select("\"PelaporanPpat\".\"Sptpd_Id\"", "(\"Ppat\".\"Nama\") \"Ppat_Name\"", "\"PelaporanPpat\".\"TglLapor\"", "\"PelaporanPpat\".\"TglSSP\"", "\"PelaporanPpat\".\"NominalSSP\"", "\"PelaporanPpat\".\"PihakYgMengalihkan\"", "\"BphtbSptpd\".*", "(\"BphtbSptpd\".\"Id\") \"SptpdUuid\"").
		Joins("LEFT JOIN \"BphtbSptpd\" ON \"BphtbSptpd\".\"Sptpd_Id\" = \"PelaporanPpat\".\"Sptpd_Id\"").
		Joins("LEFT JOIN \"Ppat\" ON \"Ppat\".\"Id\" = CAST(\"PelaporanPpat\".\"Ppat_Id\" AS INTEGER) ")

	if input.Ppat_Id != nil {
		queryBase = queryBase.Where("\"PelaporanPpat\".\"Ppat_Id\"", input.Ppat_Id)
	}

	if input.Bulan != nil {
		queryBase = queryBase.Where("EXTRACT('month' from \"VerifikasiPpatAt\") = ?", input.Bulan)
	}

	if input.Tahun != nil {
		queryBase = queryBase.Where("EXTRACT('year' from \"VerifikasiPpatAt\") = ?", input.Tahun)
	}

	queryBase = queryBase.Order("\"TahunPajakSppt\"")
	result := queryBase.
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	tempPPATID, _ := strconv.Atoi(*input.Ppat_Id)
	resPPAT := a.DB.First(dataPPAT, tempPPATID)
	if resPPAT.Error != nil {
		namePPAT = "PPAT" + *input.Ppat_Id
	} else {
		namePPAT = dataPPAT.Nama
	}

	if input.Bulan != nil && input.Tahun != nil {
		dateString = strconv.Itoa(*input.Tahun) + "-" + strconv.Itoa(*input.Bulan) + "-01"
	} else if input.Bulan != nil {
		dateString = "2022-" + strconv.Itoa(*input.Bulan) + "-01"
	} else if input.Tahun != nil {
		dateString = strconv.Itoa(*input.Tahun) + "-12-01"
	} else {
		if len(data) > 0 && data[0].TglLapor != nil {
			dateString = *data[0].TglLapor + "-12-01"
		} else {
			dateString = "2022-12-01"
		}
	}
	resT, errPer := time.Parse("2006-01-02", dateString)
	if errPer != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "periode tidak ditemukan", data)
	}

	startPeriode := time.Date(resT.Year(), resT.Month(), 1, 0, 0, 0, 0, time.UTC)
	endPeriode := startPeriode.AddDate(0, 1, 0).Add(-time.Nanosecond)
	if input.Bulan == nil {
		startPeriode = time.Date(resT.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"start":        startPeriode.String(),
			"end":          endPeriode.String(),
			"name":         namePPAT,
		},
		Data: data,
	}, nil
}

func LastDayOfMonth(t time.Time) int {
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, 0).Add(-time.Nanosecond)
	return lastDay.Day()
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.PelaporanPpat
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data pelaporan ppat", data)
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
	var data *m.PelaporanPpat
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

func getData(input msptpd.FilterPPATDto) (*[]m.ResponsePelaporanPpat, error) {
	var data []m.ResponsePelaporanPpat

	queryBase := a.DB.Model(&m.PelaporanPpat{})
	queryBase = queryBase.
		Select("DISTINCT ON (\"PelaporanPpat\".\"Ppat_Id\") \"PelaporanPpat\".\"Ppat_Id\"", "(\"Ppat\".\"Nama\") \"Ppat_Name\"", "\"PelaporanPpat\".\"TglLapor\"", "count(\"PelaporanPpat\".\"Sptpd_Id\") \"Sptpd_Id\"", "sum(\"BphtbSptpd\".\"NilaiOp\") \"NilaiOp\"", "sum(\"BphtbSptpd\".\"JumlahSetor\") \"JumlahSetor\"", "\"BphtbSptpd\".\"Status\"").
		Joins("LEFT JOIN \"BphtbSptpd\" ON \"BphtbSptpd\".\"Sptpd_Id\" = \"PelaporanPpat\".\"Sptpd_Id\"").
		Joins("LEFT JOIN \"Ppat\" ON \"Ppat\".\"Id\" = CAST(\"PelaporanPpat\".\"Ppat_Id\" AS INTEGER) ")

	if input.Ppat_Id != nil {
		queryBase = queryBase.Where("\"PelaporanPpat\".\"Ppat_Id\"", input.Ppat_Id)
	}

	if input.Bulan != nil {
		queryBase = queryBase.Where("EXTRACT('month' from \"TglLapor\") = ?", input.Bulan)
	}

	if input.Tahun != nil {
		queryBase = queryBase.Where("EXTRACT('year' from \"TglLapor\") = ?", input.Tahun)
	}

	result := queryBase.
		Order("\"PelaporanPpat\".\"Ppat_Id\"").
		Group("\"PelaporanPpat\".\"Ppat_Id\", \"Ppat\".\"Nama\", \"PelaporanPpat\".\"TglLapor\", \"BphtbSptpd\".\"Status\"").
		Find(&data)
	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data laporan ppat", data)
		return nil, err
	}

	return &data, nil
}

// GetReportExcelTransaksiPPAT func
func GetReportExcelTransaksiPPAT(input msptpd.FilterPPATDto) (*excelize.File, error) {
	data, err := getData(input)
	if err != nil {
		return nil, err
	}

	// header
	xlsx := excelize.NewFile()
	xlsx.SetSheetName(xlsx.GetSheetName(0), "List")
	start := 4
	xlsx.SetCellValue("List", fmt.Sprintf(`A%d`, start), "No")
	xlsx.SetCellValue("List", fmt.Sprintf(`B%d`, start), "Nama PPAT")
	xlsx.SetCellValue("List", fmt.Sprintf(`C%d`, start), "Tanggal Laporan")
	xlsx.SetCellValue("List", fmt.Sprintf(`D%d`, start), "Jumlah Transaksi")
	xlsx.SetCellValue("List", fmt.Sprintf(`E%d`, start), "Nominal Transaksi (Rp.)")
	xlsx.SetCellValue("List", fmt.Sprintf(`F%d`, start), "Nominal BPHTB (Rp.)")
	xlsx.SetCellValue("List", fmt.Sprintf(`G%d`, start), "Status")

	for i, v := range *data {
		n := start + i + 1
		xlsx.SetCellValue("List", fmt.Sprintf(`A%d`, n), i+1)
		xlsx.SetCellValue("List", fmt.Sprintf(`B%d`, n), func() string {
			if v.Ppat_Name != nil {
				return *v.Ppat_Name
			}
			return ""
		}())
		xlsx.SetCellValue("List", fmt.Sprintf(`C%d`, n), func() string {
			if v.TglLapor != nil {
				return *v.TglLapor
			}
			return ""
		}())
		xlsx.SetCellValue("List", fmt.Sprintf(`D%d`, n), func() string {
			if v.Sptpd_Id != nil {
				return *v.Sptpd_Id
			}
			return ""
		}())
		xlsx.SetCellValue("List", fmt.Sprintf(`E%d`, n), func() string {
			if v.NilaiOp != nil {
				return *v.NilaiOp
			}
			return ""
		}())
		xlsx.SetCellValue("List", fmt.Sprintf(`F%d`, n), func() string {
			if v.JumlahSetor != nil {
				return *v.JumlahSetor
			}
			return ""
		}())
		xlsx.SetCellValue("List", fmt.Sprintf(`G%d`, n), func() string {
			if v.Status != nil {
				return *v.Status
			}
			return ""
		}())
	}

	xlsx.SetCellValue("List", "A1", "Transaksi PPAT")
	xlsx.MergeCell("List", "A1", "G1")

	xlsx.SetCellValue("List", "A2", fmt.Sprintf(`Periode: %s %d`, months[*input.Bulan-1], *input.Tahun))
	xlsx.MergeCell("List", "A2", "G2")
	return xlsx, nil
}

// ResponseFile struct
type ResponseFile struct {
	FileName string `json:"fileName"`
}

// GetReportPDFTransaksiPPAT func
func GetReportPDFTransaksiPPAT(input msptpd.FilterPPATDto) (interface{}, error) {
	data, err := getData(input)
	if err != nil {
		return nil, err
	}

	uuid, err := sh.GetUuidv4()
	if err != nil {
		return nil, err
	}
	fileName := sh.GenerateFilename("laporan-ppat", uuid, 0, "pdf")

	outFile := filepath.Join(sh.GetPdfPath(), fileName)

	// get template
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	exPath := filepath.Dir(ex)
	path := filepath.Join(exPath, "../../", "internal", "models", "pelaporanppat", "templates", "pelaporanppat.html")

	// generate pdf
	pdfg, err := pdf.NewPdfg(wkhtmltopdf.PageSizeA4, wkhtmltopdf.OrientationPortrait, false)
	if err != nil {
		return nil, err
	}

	var tmpData []any
	// generate data
	for i, v := range *data {
		n := i + 1

		tmp := []any{
			n,
			func() string {
				if v.Ppat_Name != nil {
					return *v.Ppat_Name
				}
				return ""
			}(),
			func() string {
				if v.TglLapor != nil {
					return *v.TglLapor
				}
				return ""
			}(),
			func() string {
				if v.Sptpd_Id != nil {
					return *v.Sptpd_Id
				}
				return ""
			}(),
			func() string {
				if v.NilaiOp != nil {
					return *v.NilaiOp
				}
				return ""
			}(),
			func() string {
				if v.JumlahSetor != nil {
					return *v.JumlahSetor
				}
				return ""
			}(),
			func() string {
				if v.Status != nil {
					return *v.Status
				}
				return ""
			}(),
		}
		tmpData = append(tmpData, tmp)
	}
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, map[string]interface{}{
		"title":    "Transaksi PPAT",
		"subtitle": fmt.Sprintf(`Periode: %s %d`, months[*input.Bulan-1], *input.Tahun),
		"data":     tmpData,
	})
	if err != nil {
		return nil, err
	}

	if err := pdfg.GeneratePdfFromReader(buf, outFile); err != nil {
		return nil, err
	}

	_r := &ResponseFile{
		FileName: outFile,
	}
	return rp.OKSimple{Data: _r}, nil
}
