package sinkronisasi

import (
	"errors"
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sinkronisasi"
	"github.com/bapenda-kota-malang/apin-backend/pkg/base64helper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/xuri/excelize/v2"
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

func stringToFloat64(input string) *float64 {
	s, _ := strconv.ParseFloat(input, 64)
	return &s
}

func readExcelFilePdl(filename string) ([]m.ExcelPdl, error) {
	var excelPdl []m.ExcelPdl
	filepath := sh.GetExcelPath() + "/" + filename

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return nil, err
	}

	sheet := "Sheet1"

	rowLen, _ := f.GetRows(sheet)
	for i := 0; i < len(rowLen)+1; i++ {
		if i > 1 {
			n := strconv.Itoa(i)

			cNoRekening, _ := f.GetCellValue(sheet, "A"+n)
			cKeterangan, _ := f.GetCellValue(sheet, "B"+n)
			cNominal, _ := f.GetCellValue(sheet, "C"+n)
			cTanggal, _ := f.GetCellValue(sheet, "D"+n)
			cNamaRekening, _ := f.GetCellValue(sheet, "E"+n)

			excelPdl = append(excelPdl, m.ExcelPdl{
				NoRekening:   &cNoRekening,
				Keterangan:   &cKeterangan,
				Nominal:      stringToFloat64(cNominal),
				Tanggal:      &cTanggal,
				NamaRekening: &cNamaRekening,
			})
		}
	}
	return excelPdl, nil
}

func readExcelFilePbbBillAgregat(filename string) ([]m.ExcelPbbBillerAgregat, error) {

	var excelPbbBillAgregat []m.ExcelPbbBillerAgregat
	filepath := sh.GetExcelPath() + "/" + filename

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return nil, err
	}

	sheet := "Biller Agregat"

	rowLen, _ := f.GetRows(sheet)
	for i := 0; i < len(rowLen)+1; i++ {
		if i > 1 {
			n := strconv.Itoa(i)

			cNop, _ := f.GetCellValue(sheet, "C"+n)
			cNama, _ := f.GetCellValue(sheet, "D"+n)
			cTahun, _ := f.GetCellValue(sheet, "E"+n)
			cPokok, _ := f.GetCellValue(sheet, "F"+n)
			cDenda, _ := f.GetCellValue(sheet, "G"+n)
			cNominal, _ := f.GetCellValue(sheet, "H"+n)
			cTanggal, _ := f.GetCellValue(sheet, "I"+n)
			cJam, _ := f.GetCellValue(sheet, "J"+n)
			cBiller, _ := f.GetCellValue(sheet, "K"+n)

			excelPbbBillAgregat = append(excelPbbBillAgregat, m.ExcelPbbBillerAgregat{
				Nop:     &cNop,
				Nama:    &cNama,
				Tahun:   &cTahun,
				Pokok:   stringToFloat64(cPokok),
				Denda:   stringToFloat64(cDenda),
				Nominal: stringToFloat64(cNominal),
				Tanggal: &cTanggal,
				Jam:     &cJam,
				Biller:  &cBiller,
			})
		}
	}
	return excelPbbBillAgregat, nil
}
