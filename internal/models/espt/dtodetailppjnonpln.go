package espt

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptppjnonpln"
	"github.com/google/uuid"
)

type CreateDetailPpjNonPlnDto struct {
	CreateDetailBaseDto
	DataDetails detailesptppjnonpln.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailPpjNonPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailPpjNonPlnDto) LenDetails() int {
	newEmpty := detailesptppjnonpln.CreateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}

func (input *CreateDetailPpjNonPlnDto) ReplaceEsptId(id uuid.UUID) {
	input.DataDetails.Espt_Id = id
}

type UpdateDetailPpjNonPlnDto struct {
	UpdateDetailBaseDto
	DataDetails detailesptppjnonpln.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailPpjNonPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailPpjNonPlnDto) LenDetails() int {
	newEmpty := detailesptppjnonpln.UpdateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}
