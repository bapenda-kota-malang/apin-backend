package main

import (
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
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

	rowLen := len(f.GetRows(sheet))
	for i := 0; i < rowLen+1; i++ {
		if i > 1 {
			n := strconv.Itoa(i)

			// cParentId := f.GetCellValue(sheet, "B"+n)
			cTipe := f.GetCellValue(sheet, "C"+n)
			cKelompok := f.GetCellValue(sheet, "D"+n)
			cJenis := f.GetCellValue(sheet, "E"+n)
			cObjek := f.GetCellValue(sheet, "F"+n)
			cRincian := f.GetCellValue(sheet, "G"+n)
			cSub1 := f.GetCellValue(sheet, "H"+n)
			cSub2 := f.GetCellValue(sheet, "I"+n)
			cSub3 := f.GetCellValue(sheet, "J"+n)
			cKode := f.GetCellValue(sheet, "K"+n)
			cNama := f.GetCellValue(sheet, "L"+n)
			cKodeJenisPajak := f.GetCellValue(sheet, "O"+n)
			cKodeVaJatim := f.GetCellValue(sheet, "P"+n)
			cKodeBilling := f.GetCellValue(sheet, "Q"+n)
			cKodeJenisUsaha := f.GetCellValue(sheet, "R"+n)
			cJenisUsaha := f.GetCellValue(sheet, "S"+n)

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
