package espt

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptresto"
	"github.com/google/uuid"
)

type CreateDetailRestoDto struct {
	CreateDetailBaseDto
	DataDetails detailesptresto.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailRestoDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailRestoDto) LenDetails() int {
	newEmpty := detailesptresto.CreateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}

func (input *CreateDetailRestoDto) ReplaceEsptId(id uuid.UUID) {
	input.DataDetails.Espt_Id = id
}

type UpdateDetailRestoDto struct {
	UpdateDetailBaseDto
	DataDetails detailesptresto.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailRestoDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailRestoDto) LenDetails() int {
	newEmpty := detailesptresto.UpdateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}
