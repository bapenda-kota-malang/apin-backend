package spt

import (
	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mdhiburan "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthiburan"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type CreateDetailHiburanDto struct {
	CreateDetailBaseDto
	DataDetails mdhiburan.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailHiburanDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailHiburanDto) LenDetails() int {
	newEmpty := mdhiburan.CreateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}

func (input *CreateDetailHiburanDto) ReplaceSptId(id uuid.UUID) {
	input.DataDetails.Spt_Id = id
}

func (input *CreateDetailHiburanDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	if err := input.CreateDetailBaseDto.DuplicateEspt(esptDetail); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &esptDetail); err != nil {
		return err
	}
	return nil
}

func (input *CreateDetailHiburanDto) SkpdkbDuplicate(sptDetail *Spt, skpdkb *SkpdkbExisting) error {
	if err := input.CreateDetailBaseDto.SkpdkbDuplicate(sptDetail, skpdkb); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &sptDetail); err != nil {
		return err
	}
	return nil
}

// Update
type UpdateDetailHiburanDto struct {
	UpdateDetailBaseDto
	DataDetails mdhiburan.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailHiburanDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailHiburanDto) LenDetails() int {
	newEmpty := mdhiburan.UpdateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}
