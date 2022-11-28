package spt

import (
	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mdsh "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthotel"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type CreateDetailHotelDto struct {
	CreateDetailBaseDto
	DataDetails []mdsh.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailHotelDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailHotelDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input *CreateDetailHotelDto) ReplaceSptId(id uuid.UUID) {
	for i := range input.DataDetails {
		input.DataDetails[i].Spt_Id = id
	}
}

func (input *CreateDetailHotelDto) ChangeDetails(newData interface{}) {
	// Do nothing
}

func (input *CreateDetailHotelDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	if err := input.CreateDetailBaseDto.DuplicateEspt(esptDetail); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &esptDetail); err != nil {
		return err
	}
	return nil
}

func (input *CreateDetailHotelDto) SkpdkbDuplicate(sptDetail *Spt, skpdkb *SkpdkbExisting) error {
	if err := input.CreateDetailBaseDto.SkpdkbDuplicate(sptDetail, skpdkb); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &sptDetail); err != nil {
		return err
	}
	return nil
}

type UpdateDetailHotelDto struct {
	UpdateDetailBaseDto
	DataDetails []mdsh.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailHotelDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailHotelDto) LenDetails() int {
	return len(input.DataDetails)
}
