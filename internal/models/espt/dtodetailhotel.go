package espt

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailespthotel"
	"github.com/google/uuid"
)

type CreateDetailHotelDto struct {
	CreateDetailBaseDto
	DataDetails detailespthotel.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailHotelDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailHotelDto) LenDetails() int {
	newEmpty := detailespthotel.CreateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}

func (input *CreateDetailHotelDto) ReplaceEsptId(id uuid.UUID) {
	input.DataDetails.Espt_Id = id
}

type UpdateDetailHotelDto struct {
	UpdateDetailBaseDto
	DataDetails detailespthotel.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailHotelDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailHotelDto) LenDetails() int {
	newEmpty := detailespthotel.UpdateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}
