package espt

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptppjpln"
	"github.com/google/uuid"
)

type CreateDetailPpjPlnDto struct {
	CreateDetailBaseDto
	DataDetails []detailesptppjpln.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailPpjPlnDto) CalculateTax(taxPercentage *float64) {
	tax := float64(0)
	for v := range input.DataDetails {
		if input.DataDetails[v].JenisPPJ.Jenis == "LAIN LAIN" {
			continue
		}
		jumlah := float64(input.DataDetails[v].JumlahRekening)
		jumlah = jumlah * (float64(input.DataDetails[v].JenisPPJ.TarifPersen) / 100)
		tax += jumlah
	}
	input.Espt.JumlahPajak = tax + input.Espt.Omset
}

func (input *CreateDetailPpjPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailPpjPlnDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input *CreateDetailPpjPlnDto) ReplaceEsptId(id uuid.UUID) {
	for i := range input.DataDetails {
		input.DataDetails[i].Espt_Id = id
	}
}

type UpdateDetailPpjPlnDto struct {
	UpdateDetailBaseDto
	DataDetails []detailesptppjpln.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input UpdateDetailPpjPlnDto) CalculateTax(taxPercentage *float64) {
	tax := float64(0)
	for v := range input.DataDetails {
		if input.DataDetails[v].JenisPPJ.Jenis == "LAIN LAIN" {
			continue
		}
		jumlah := float64(*input.DataDetails[v].JumlahRekening)
		jumlah = jumlah * (float64(input.DataDetails[v].JenisPPJ.TarifPersen) / 100)
		tax += jumlah
	}
	tax += *input.Espt.Omset
	input.Espt.JumlahPajak = &tax
}

func (input *UpdateDetailPpjPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailPpjPlnDto) LenDetails() int {
	return len(input.DataDetails)
}
