package penerimaaan

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sts"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/pdf"
	"github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	strh "github.com/bapenda-kota-malang/apin-backend/pkg/stringhelper"
)

// Private funcs start here
const source = "penerimaan"

var title1 = "Rekapitulasi Penerimaan"
var title2 = ""
var months = []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

func getData(input sspd.FilterReport) (*[]sspd.Sspd, []sts.Sts, error) {
	var data []sspd.Sspd
	var stsData []sts.Sts
	db := a.DB
	now := time.Now()

	override := func(d *sspd.Sspd) {
		if d.SspdDetails[0].Sts_Id == nil {
			return
		}
		var tmp sts.Sts
		db.Model(&sts.Sts{}).First(&tmp, db.Where(sts.Sts{
			Id: *d.SspdDetails[0].Sts_Id,
		}))

		stsData = append(stsData, tmp)
	}

	result := db.Model(&data)
	// filter conditions
	if input.StartMonth != nil && input.StopMonth != nil { // bulanan
		result = result.Where(`EXTRACT('Month' FROM "Sspd"."CreatedAt") >= ? AND EXTRACT('Month' FROM "Sspd"."CreatedAt") <= ?`, input.StartMonth, input.StopMonth)

		title2 = fmt.Sprintf(`Bulan %s - %s`, months[*input.StartMonth+1], months[*input.StopMonth+1])
		if input.Year != nil {
			result = result.Where(`EXTRACT('Year' FROM "Sspd"."CreatedAt") = ?`, input.Year)

			title2 = fmt.Sprintf(`%s %s`, title2, *input.Year)
		} else {
			result = result.Where(`EXTRACT('Year' FROM "Sspd"."CreatedAt") = EXTRACT('Year' FROM CURRENT_DATE)`)

			title2 = fmt.Sprintf(`%s %d`, title2, now.Year())
		}
	} else if input.Year != nil { // tahunan
		result = result.Where(`EXTRACT('Year' FROM "Sspd"."CreatedAt") = ?`, input.Year)

		title2 = fmt.Sprintf(`Tahun %s`, *input.Year)
	} else if input.StartDate != nil && input.StopDate != nil { // harian
		result = result.Where(`"Sspd"."CreatedAt" BETWEEN TO_DATE(?, 'YYYY-MM-DD') - INTERVAL '1 day' AND TO_DATE(?, 'YYYY-MM-DD') + INTERVAL '1 day'`, input.StartDate, input.StopDate)

		title2 = fmt.Sprintf(`%s %s`, title2, *input.Year)
	}

	result = result.
		Joins(`ObjekPajak`).
		Joins(`Rekening`).
		Joins(`Npwpd`).
		Preload("SspdDetails.Spt").
		Find(&data)

	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data penerimaan", data)
		return nil, nil, err
	}

	for index := range data {
		override(&data[index])
	}

	return &data, stsData, nil
}

// GetList func
func GetList(input sspd.FilterReport) (any, error) {
	data, _, err := getData(input)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetReportExcelPenerimaan func
func GetReportExcelPenerimaan(input sspd.FilterReport) (*excelize.File, error) {
	data, stsData, err := getData(input)
	if err != nil {
		return nil, err
	}

	xlsx := excelize.NewFile()

	penerimaanReport := title1
	start := 5
	xlsx.SetSheetName(xlsx.GetSheetName(1), penerimaanReport)

	// create header
	xlsx.SetCellValue(penerimaanReport, "A1", fmt.Sprintf(`%s `, title1))
	xlsx.MergeCell(penerimaanReport, "A1", "N1")
	xlsx.SetCellValue(penerimaanReport, "A2", fmt.Sprintf(`%s `, title2))
	xlsx.MergeCell(penerimaanReport, "A2", "N2")

	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("A%d", start), "No")
	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("B%d", start), "No SSPD")
	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("C%d", start), "No SKPD")
	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("D%d", start), "Tanggal TBP")
	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("E%d", start), "Tanggal STS")
	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("F%d", start), "Nama")
	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("G%d", start), "Alamat")
	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("H%d", start), "NPWPD")
	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("I%d", start), "Bulan")
	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("J%d", start), "Tahun")
	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("K%d", start), "Rekening")
	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("L%d", start), "Jenis Pajak")
	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("M%d", start), "Jumlah")
	xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("N%d", start), "Keterangan")

	// set value
	for i, v := range *data {
		n := start + i + 1

		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("A%d", n), i+1)
		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("B%d", n), v.NomorOutput)
		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("C%d", n), strh.NullToAny(&v.SspdDetails[0].Spt.NomorSpt, ""))
		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("D%d", n), v.TanggalBayar.Format("2006-01-02"))
		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("E%d", n), func() string {
			if len(stsData) > 0 && stsData[i].TanggalSts != nil {
				return stsData[i].TanggalSts.Format("2006-01-02")
			}

			return ""
		}())
		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("F%d", n), *v.PenyetorName)
		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("G%d", n), *v.PenyetorAddress)
		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("H%d", n), *v.Npwpd_Npwpd)
		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("I%d", n), months[v.CreatedAt.Month()-1])
		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("J%d", n), v.CreatedAt.Year())
		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("K%d", n), *v.Rekening.Kode)
		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("L%d", n), *v.Rekening.Nama)
		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("M%d", n), servicehelper.FormatCurrency(*v.Total))
		xlsx.SetCellValue(penerimaanReport, fmt.Sprintf("N%d", n), *v.Note)
	}

	style, _ := xlsx.NewStyle(`{"alignment":{"horizontal": "center","vertical": "center"}}`)
	xlsx.SetCellStyle(penerimaanReport, "A1", fmt.Sprintf(`N%d`, len(*data)+start), style)
	return xlsx, nil
}

// ResponseFile struct
type ResponseFile struct {
	FileName string `json:"fileName"`
}

// GetReportPDFPenerimaan func
func GetReportPDFPenerimaan(input sspd.FilterReport) (interface{}, error) {
	data, stsData, err := getData(input)
	if err != nil {
		return nil, err
	}

	uuid, err := sh.GetUuidv4()
	if err != nil {
		return nil, err
	}
	fileName := sh.GenerateFilename("rekapitulasi-penerimaan", uuid, 0, "pdf")

	outFile := filepath.Join(sh.GetPdfPath(), fileName)

	// get template
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	exPath := filepath.Dir(ex)
	path := filepath.Join(exPath, "../../", "internal", "models", "penerimaan", "templates", "penerimaan.html")

	// generate pdf
	pdfg, err := pdf.NewPdfg(wkhtmltopdf.PageSizeLegal, wkhtmltopdf.OrientationLandscape, false)
	if err != nil {
		return nil, err
	}

	var tmpData []any
	// generate data
	for i, v := range *data {
		n := i + 1

		tmp := []any{n, v.NomorOutput, strh.NullToAny(&v.SspdDetails[0].Spt.NomorSpt, ""), v.TanggalBayar.Format("2006-01-02"), func() string {
			if len(stsData) > 0 && stsData[i].TanggalSts != nil {
				return stsData[i].TanggalSts.Format("2006-01-02")
			}

			return ""
		}(), *v.PenyetorName, *v.PenyetorAddress, *v.Npwpd_Npwpd, months[v.CreatedAt.Month()-1], v.CreatedAt.Year(), *v.Rekening.Kode, *v.Rekening.Nama, servicehelper.FormatCurrency(*v.Total), *v.Note}
		tmpData = append(tmpData, tmp)
	}
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, map[string]interface{}{
		"title":    title1,
		"subtitle": title2,
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
