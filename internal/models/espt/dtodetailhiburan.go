package espt

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailespthiburan"
	"github.com/google/uuid"
)

// Create
type CreateDetailHiburanDto struct {
	CreateDetailBaseDto
	DataDetails detailespthiburan.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailHiburanDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailHiburanDto) LenDetails() int {
	newEmpty := detailespthiburan.CreateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}

func (input *CreateDetailHiburanDto) ReplaceEsptId(id uuid.UUID) {
	input.DataDetails.Espt_Id = id
}

// Update
type UpdateDetailHiburanDto struct {
	UpdateDetailBaseDto
	DataDetails detailespthiburan.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailHiburanDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailHiburanDto) LenDetails() int {
	newEmpty := detailespthiburan.UpdateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}
