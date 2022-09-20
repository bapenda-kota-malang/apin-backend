package main

import (
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/configurationmodel"
)

func ImportRekening() {
	var rekening []configurationmodel.Rekening
	result := DB.First(&rekening)
	if result.RowsAffected > 0 {
		log.Fatal("Rekening not empty")
		return
	}

	f, err := excelize.OpenFile("assets/rekening.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	sheet := "Data Rekening"

	rowLen := len(f.GetRows(sheet))
	for i := 0; i < rowLen+1; i++ {
		if i > 1 {
			n := strconv.Itoa(i)

			cTipe := f.GetCellValue(sheet, "B"+n)
			cKelompok := f.GetCellValue(sheet, "C"+n)
			cJenis := f.GetCellValue(sheet, "D"+n)
			cObjek := f.GetCellValue(sheet, "E"+n)
			cRincian := f.GetCellValue(sheet, "F"+n)
			cSub1 := f.GetCellValue(sheet, "G"+n)
			cSub2 := f.GetCellValue(sheet, "H"+n)
			cSub3 := f.GetCellValue(sheet, "I"+n)
			cKode := f.GetCellValue(sheet, "J"+n)
			cNama := f.GetCellValue(sheet, "K"+n)

			rekening = append(rekening, configurationmodel.Rekening{
				Tipe:     &cTipe,
				Kelompok: &cKelompok,
				Jenis:    &cJenis,
				Objek:    &cObjek,
				Rincian:  &cRincian,
				Sub1:     &cSub1,
				Sub2:     &cSub2,
				Sub3:     &cSub3,
				Kode:     &cKode,
				Nama:     &cNama,
			})
		}
	}

	err = DB.Create(&rekening).Error
	if err != nil {
		log.Fatal("Failed to create")
	}

	log.Println("Success to migrate rekening")
}
