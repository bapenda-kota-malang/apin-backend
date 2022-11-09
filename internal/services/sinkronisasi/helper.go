package sinkronisasi

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sinkronisasi"
	"github.com/bapenda-kota-malang/apin-backend/pkg/base64helper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

func filePreProcess(b64String string, userId uint, docsName string) (fileName, path, extFile string, err error) {
	extFile, err = base64helper.GetExtensionBase64(b64String)
	if err != nil {
		return
	}
	path = sh.GetResourcesPath()
	switch extFile {
	case "xlsx", "xls":
		path = sh.GetExcelPath()
	default:
		err = errors.New("file bukan excel")
		return
	}
	id, err := sh.GetUuidv4()
	if err != nil {
		err = errors.New("gagal generate uuid")
		return
	}
	fileName = sh.GenerateFilename(docsName, id, userId, extFile)
	return
}

func stringToFloat64(input string) float64 {
	s, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("gagal parse float")
	}
	return s
}

func readExcelFile(filename string) []m.ExcelPdl {
	var excelPdl []m.ExcelPdl
	filepath := sh.GetExcelPath() + "/" + filename

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	sheet := "Sheet1"

	rowLen := len(f.GetRows(sheet))
	for i := 0; i < rowLen+1; i++ {
		if i > 1 {
			n := strconv.Itoa(i)

			cNoRekening := f.GetCellValue(sheet, "A"+n)
			cKeterangan := f.GetCellValue(sheet, "B"+n)
			cNominal := f.GetCellValue(sheet, "C"+n)
			cTanggal := f.GetCellValue(sheet, "D"+n)
			cNamaRekening := f.GetCellValue(sheet, "E"+n)

			excelPdl = append(excelPdl, m.ExcelPdl{
				NoRekening:   cNoRekening,
				Keterangan:   cKeterangan,
				Nominal:      stringToFloat64(cNominal),
				Tanggal:      cTanggal,
				NamaRekening: cNamaRekening,
			})
		}
	}
	fmt.Println("hasil excelPdl: ", excelPdl)
	return excelPdl
}
