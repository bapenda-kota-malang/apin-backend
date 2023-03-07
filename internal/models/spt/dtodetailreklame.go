package spt

import (
	"math"

	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mdsrek "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/datatypes"
)

type CreateJambongDto struct {
	Spt_Id         uuid.UUID       `json:"-"`
	Tanggal        *datatypes.Date `json:"tanggal" validate:"required"`
	TipeReklame    *uint8          `json:"tipeReklame"`
	Nominal        *float64        `json:"nominal"`
	TanggalBatas   *datatypes.Date `json:"TanggalBatas"`
	BiayaPemutusan *float64        `json:"biayaPemutusan"`
	DetailsJambong *[]int64        `json:"detailJambong" validate:"required"`
}

type CreateDetailReklameDto struct {
	CreateDetailBaseDto
	DataDetails    []mdsrek.CreateDto `json:"dataDetails" validate:"required"`
	JaminanBongkar *CreateJambongDto  `json:"jaminanBongkar"`
}

func (input *CreateDetailReklameDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailReklameDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input *CreateDetailReklameDto) ReplaceSptId(id uuid.UUID) {
	for i := range input.DataDetails {
		input.DataDetails[i].Spt_Id = id
	}
}

func (input *CreateDetailReklameDto) CalculateTax(taxPercentage *float64) {
	tax := float64(0)
	for v := range input.DataDetails {
		// safe dereference for jumlah tahun, bulan, minggu, hari
		tahun := float64(0)
		bulan := float64(0)
		minggu := float64(0)
		hari := float64(0)
		if input.Spt.JumlahTahun != nil {
			tahun = *input.Spt.JumlahTahun
		}
		if input.Spt.JumlahBulan != nil {
			bulan = *input.Spt.JumlahBulan
		}
		if input.Spt.JumlahMinggu != nil {
			minggu = *input.Spt.JumlahMinggu
		}
		if input.Spt.JumlahHari != nil {
			hari = *input.Spt.JumlahHari
		}
		// calculate base on jenis masa pajak
		switch input.DataDetails[v].TarifReklame.JenisMasa {
		case mtypes.MasaPajakTahun:
			input.DataDetails[v].TarifTahun = tahun * *input.DataDetails[v].TarifReklame.Tarif
			input.DataDetails[v].TarifBulan = bulan * (*input.DataDetails[v].TarifReklame.Tarif / 12)
		case mtypes.MasaPajakBulan:
			input.DataDetails[v].TarifBulan = bulan * *input.DataDetails[v].TarifReklame.Tarif
		case mtypes.MasaPajakHari:
			input.DataDetails[v].TarifBulan *= bulan
			input.DataDetails[v].TarifMinggu *= minggu
			input.DataDetails[v].TarifHari *= hari
		case mtypes.MasaPajakPenyelenggara:
			input.DataDetails[v].TarifHari = *input.DataDetails[v].TarifReklame.Tarif
		}
		// basic calculate tax
		basicTax := (input.DataDetails[v].TarifHari +
			input.DataDetails[v].TarifMinggu +
			input.DataDetails[v].TarifBulan +
			input.DataDetails[v].TarifTahun) * float64(input.DataDetails[v].Jumlah)
		// if dasar pengenaan luas then add to calculation
		if *input.DataDetails[v].TarifReklame.DasarPengenaan == "Luas" {
			sisi := float64(1)
			if input.DataDetails[v].Sisi > 1 {
				sisi = float64(input.DataDetails[v].Sisi)
			}
			input.DataDetails[v].Sisi = uint64(sisi)
			luas := float64(0)
			switch input.DataDetails[v].JenisDimensi {
			case mdsrek.JenisDimensiPersegiPanjang:
				luas = float64(input.DataDetails[v].Panjang) * float64(input.DataDetails[v].Lebar)
			case mdsrek.JenisDimensiPersegiLingkaran:
				r := float64(*input.DataDetails[v].Diameter) / 2
				luas = math.Pi * math.Pow(r, 2)
			}
			basicTax *= (luas * sisi)
		}
		input.DataDetails[v].JumlahRp = basicTax
		tax += basicTax
	}
	input.Spt.JumlahPajak = tax
}

func (input *CreateDetailReklameDto) DuplicateEspt(sptDetail *mespt.Espt) error {
	// do nothing because spt don't have reklame
	return nil
}

func (input *CreateDetailReklameDto) SkpdkbDuplicate(sptDetail *Spt, skpdkb *SkpdkbExisting) error {
	if err := input.CreateDetailBaseDto.SkpdkbDuplicate(sptDetail, skpdkb); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &sptDetail); err != nil {
		return err
	}
	return nil
}

type UpdateDetailReklameDto struct {
	UpdateDetailBaseDto
	DataDetails []mdsrek.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailReklameDto) CalculateTax(taxPercentage *float64) {
	// TODO: CHANGE THIS CALCULATION PROCESS
	calc := *input.Spt.Omset * *taxPercentage / 100
	input.Spt.JumlahPajak = &calc
}

func (input *UpdateDetailReklameDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailReklameDto) LenDetails() int {
	return len(input.DataDetails)
}
