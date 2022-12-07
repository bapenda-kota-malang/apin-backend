package espt

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptair"
	"github.com/google/uuid"
)

// Create
type CreateDetailAirDto struct {
	CreateDetailBaseDto
	DataDetails detailesptair.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailAirDto) CalculateTax(taxPercentage *float64) {
	input.Espt.JumlahPajak = input.Espt.Omset * (*taxPercentage / 100) * float64(input.DataDetails.Pengenaan)
}

func (input *CreateDetailAirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailAirDto) LenDetails() int {
	newEmpty := detailesptair.CreateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}

func (input *CreateDetailAirDto) ChangeDetails(newData interface{}) {
	input.DataDetails.Pengenaan = newData.(detailesptair.CreateDto).Pengenaan
	input.DataDetails.JenisAbt = newData.(detailesptair.CreateDto).JenisAbt
}

func (input *CreateDetailAirDto) ReplaceEsptId(id uuid.UUID) {
	input.DataDetails.Espt_Id = id
}

// Update
type UpdateDetailAirDto struct {
	UpdateDetailBaseDto
	DataDetails detailesptair.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailAirDto) CalculateTax(taxPercentage *float64) {
	calc := *input.Espt.Omset * *taxPercentage / 100 * float64(input.DataDetails.Pengenaan)
	input.Espt.JumlahPajak = &calc
}

func (input *UpdateDetailAirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailAirDto) LenDetails() int {
	newEmpty := detailesptair.UpdateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}
