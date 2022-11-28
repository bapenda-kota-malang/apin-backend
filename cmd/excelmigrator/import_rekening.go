package main

import (
	"log"
	"strconv"

	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/xuri/excelize/v2"
)

func ImportRekening() {
	var rekening []rm.Rekening
	result := DB.First(&rekening)
	if result.RowsAffected > 0 {
		log.Fatal("Rekening not empty")
		return
	}

	f, err := excelize.OpenFile("assets/dataDummyRekening.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	sheet := "Data Rekening"

	rowLen, _ := f.GetRows(sheet)
	for i := 0; i < len(rowLen)+1; i++ {
		if i > 1 {
			n := strconv.Itoa(i)

			// cParentId := f.GetCellValue(sheet, "B"+n)
			cTipe, _ := f.GetCellValue(sheet, "C"+n)
			cKelompok, _ := f.GetCellValue(sheet, "D"+n)
			cJenis, _ := f.GetCellValue(sheet, "E"+n)
			cObjek, _ := f.GetCellValue(sheet, "F"+n)
			cRincian, _ := f.GetCellValue(sheet, "G"+n)
			cSub1, _ := f.GetCellValue(sheet, "H"+n)
			cSub2, _ := f.GetCellValue(sheet, "I"+n)
			cSub3, _ := f.GetCellValue(sheet, "J"+n)
			cKode, _ := f.GetCellValue(sheet, "K"+n)
			cNama, _ := f.GetCellValue(sheet, "L"+n)
			cKodeJenisPajak, _ := f.GetCellValue(sheet, "O"+n)
			cKodeVaJatim, _ := f.GetCellValue(sheet, "P"+n)
			cKodeBilling, _ := f.GetCellValue(sheet, "Q"+n)
			cKodeJenisUsaha, _ := f.GetCellValue(sheet, "R"+n)
			cJenisUsaha, _ := f.GetCellValue(sheet, "S"+n)

			rekening = append(rekening, rm.Rekening{
				Tipe:           &cTipe,
				Kelompok:       &cKelompok,
				Jenis:          &cJenis,
				Objek:          &cObjek,
				Rincian:        &cRincian,
				Sub1:           &cSub1,
				Sub2:           &cSub2,
				Sub3:           &cSub3,
				Kode:           &cKode,
				Nama:           &cNama,
				KodeJenisPajak: &cKodeJenisPajak,
				KodeVaJatim:    &cKodeVaJatim,
				KodeBilling:    &cKodeBilling,
				KodeJenisUsaha: &cKodeJenisUsaha,
				JenisUsaha:     &cJenisUsaha,
			})
		}
	}

	err = DB.Create(&rekening).Error
	if err != nil {
		log.Fatal("Failed to create")
	}

	log.Println("Success to migrate rekening")
}
