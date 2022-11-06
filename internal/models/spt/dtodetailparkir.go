package spt

import (
	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mdsp "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type CreateDetailParkirDto struct {
	CreateDetailBaseDto
	DataDetails []mdsp.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailParkirDto) GetSpt(baseUri string) interface{} {
	typeSpt := mtypes.JenisPajakSA
	if baseUri == "skpd" {
		typeSpt = mtypes.JenisPajakOA
	}
	input.Spt.Type = typeSpt
	return input.Spt
}

func (input *CreateDetailParkirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailParkirDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input *CreateDetailParkirDto) ReplaceSptId(id uuid.UUID) {
	for i := range input.DataDetails {
		input.DataDetails[i].Spt_Id = id
	}
}

func (input *CreateDetailParkirDto) ChangeDetails(newData interface{}) {
	// Do nothing
}

func (input *CreateDetailParkirDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	if err := input.CreateDetailBaseDto.DuplicateEspt(esptDetail); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &esptDetail); err != nil {
		return err
	}
	return nil
}
