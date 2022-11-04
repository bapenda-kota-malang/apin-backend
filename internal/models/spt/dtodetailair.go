package spt

import (
	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mdsa "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptair"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type CreateDetailAirDto struct {
	CreateDetailBaseDto
	DataDetails mdsa.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailAirDto) CalculateTax(taxPercentage *float64) {
	input.Spt.JumlahPajak = float32(input.Spt.Omset * (*taxPercentage / 100) * float64(input.DataDetails.Pengenaan))
}

func (input *CreateDetailAirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailAirDto) LenDetails() int {
	newEmpty := mdsa.CreateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}

func (input *CreateDetailAirDto) ReplaceSptId(id uuid.UUID) {
	input.DataDetails.Spt_Id = id
}

func (input *CreateDetailAirDto) ChangeDetails(newData interface{}) {
	input.DataDetails.Pengenaan = newData.(mdsa.CreateDto).Pengenaan
	input.DataDetails.JenisAbt = newData.(mdsa.CreateDto).JenisAbt
}

func (input *CreateDetailAirDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	if err := input.CreateDetailBaseDto.DuplicateEspt(esptDetail); err != nil {
		return err
	}
	if err := copier.Copy(input.DataDetails, esptDetail.DetailEsptAir); err != nil {
		return err
	}
	return nil
}
