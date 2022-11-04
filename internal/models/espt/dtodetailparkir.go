package espt

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptparkir"
	"github.com/google/uuid"
)

type CreateDetailParkirDto struct {
	CreateDetailBaseDto
	DataDetails []detailesptparkir.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailParkirDto) GetEspt() interface{} {
	return input.Espt
}

func (input *CreateDetailParkirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailParkirDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input *CreateDetailParkirDto) ReplaceEsptId(id uuid.UUID) {
	for i := range input.DataDetails {
		input.DataDetails[i].Espt_Id = id
	}
}

type UpdateDetailParkirDto struct {
	UpdateDetailBaseDto
	DataDetails []detailesptparkir.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailParkirDto) GetEspt() interface{} {
	return input.Espt
}

func (input *UpdateDetailParkirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailParkirDto) LenDetails() int {
	return len(input.DataDetails)
}
