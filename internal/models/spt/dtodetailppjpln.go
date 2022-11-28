package spt

import (
	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mdppjpln "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptppjpln"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type CreateDetailPpjPlnDto struct {
	CreateDetailBaseDto
	DataDetails []mdppjpln.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailPpjPlnDto) CalculateTax(taxPercentage *float64) {
	tax := float64(0)
	for v := range input.DataDetails {
		jumlah := float64(input.DataDetails[v].JumlahRekening)
		if input.DataDetails[v].JenisPPJ.Jenis != "LAIN LAIN" {
			jumlah = jumlah * (float64(input.DataDetails[v].JenisPPJ.TarifPersen) / 100)
		}
		tax += jumlah
	}
	input.Spt.JumlahPajak = tax
}

func (input *CreateDetailPpjPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailPpjPlnDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input *CreateDetailPpjPlnDto) ReplaceSptId(id uuid.UUID) {
	for i := range input.DataDetails {
		input.DataDetails[i].Spt_Id = id
	}
}

func (input *CreateDetailPpjPlnDto) ChangeDetails(newData interface{}) {
	// Do nothing
}

func (input *CreateDetailPpjPlnDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	if err := input.CreateDetailBaseDto.DuplicateEspt(esptDetail); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &esptDetail); err != nil {
		return err
	}
	return nil
}

func (input *CreateDetailPpjPlnDto) SkpdkbDuplicate(sptDetail *Spt, skpdkb *SkpdkbExisting) error {
	if err := input.CreateDetailBaseDto.SkpdkbDuplicate(sptDetail, skpdkb); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &sptDetail); err != nil {
		return err
	}
	return nil
}

type UpdateDetailPpjPlnDto struct {
	UpdateDetailBaseDto
	DataDetails []mdppjpln.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input UpdateDetailPpjPlnDto) CalculateTax(taxPercentage *float64) {
	tax := float64(0)
	for v := range input.DataDetails {
		jumlah := float64(*input.DataDetails[v].JumlahRekening)
		if input.DataDetails[v].JenisPPJ.Jenis != "LAIN LAIN" {
			jumlah = jumlah * (float64(input.DataDetails[v].JenisPPJ.TarifPersen) / 100)
		}
		tax += jumlah
	}
	input.Spt.JumlahPajak = &tax
}

func (input *UpdateDetailPpjPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailPpjPlnDto) LenDetails() int {
	return len(input.DataDetails)
}
