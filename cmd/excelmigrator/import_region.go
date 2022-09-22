package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	adm "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"gorm.io/gorm"
)

func migrateProvinsi() {
	var provinsi []adm.Provinsi
	result := DB.First(&provinsi)
	if result.RowsAffected > 0 {
		fmt.Println("Provinsi not empty")
		return
	}

	filePath := "assets/provinces.csv"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	for _, r := range records {
		provinsi = append(provinsi, adm.Provinsi{
			Kode: r[0],
			Nama: r[2],
		})
	}

	err = DB.Create(&provinsi).Error
	if err != nil {
		log.Fatal("Failed to create")
	}
	log.Println("Success to migrate provinsi")
}

func migrateDaerah() {
	var daerah []adm.Daerah
	result := DB.First(&daerah)
	if result.RowsAffected > 0 {
		fmt.Println("Daerah not empty")
		return
	}

	filePath := "assets/regencies.csv"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	for _, r := range records {
		daerah = append(daerah, adm.Daerah{
			Kode:         r[0],
			ProvinsiKode: r[1],
			Nama:         r[2],
		})
	}

	err = DB.Create(&daerah).Error
	if err != nil {
		log.Fatal("Failed to create : ", err)
	}
	log.Println("Success to migrate daerah")
}

func migrateKecamatan() {
	var kecamatan []adm.Kecamatan
	result := DB.First(&kecamatan)
	if result.RowsAffected > 0 {
		fmt.Println("Kecamatan not empty")
		return
	}

	filePath := "assets/districts.csv"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	err = DB.Transaction(func(tx *gorm.DB) error {
		for _, r := range records {
			if err := tx.Create(&adm.Kecamatan{
				DaerahKode: r[1],
				Kode:       r[0],
				Nama:       r[2],
			}).Error; err != nil {
				fmt.Println("error kode : ", r[0])
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal("Failed to create : ", err)
	}
	log.Println("Success to migrate kecamatan")
}

func migrateKelurahan() {
	var kelurahan []adm.Kelurahan
	result := DB.First(&kelurahan)
	if result.RowsAffected > 0 {
		fmt.Println("Kelurahan not empty")
		return
	}

	filePath := "assets/villages.csv"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	err = DB.Transaction(func(tx *gorm.DB) error {
		for _, r := range records {
			if err := tx.Create(&adm.Kelurahan{
				KecamatanKode: r[1],
				Kode:          r[0],
				Nama:          r[2],
			}).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal("Failed to create :", err)
	}
	log.Println("Success to migrate kelurahan")
}

func ImportRegion() {
	migrateProvinsi()
	migrateDaerah()
	migrateKecamatan()
	migrateKelurahan()
}
