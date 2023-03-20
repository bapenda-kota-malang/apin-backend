package spt

import (
	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mdsp "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type CreateDetailParkirDto struct {
	CreateDetailBaseDto
	DataDetails []mdsp.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailParkirDto) GetSpt(baseUri string) interface{} {
	typeSpt := input.Spt.Type
	if typeSpt == "" {
		typeSpt = mt.JenisPajakSA
	}
	if baseUri == "skpd" {
		typeSpt = mt.JenisPajakOA
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
	if err := copier.Copy(&input.DataDetails, &esptDetail.DetailEsptParkir); err != nil {
		return err
	}
	return nil
}

func (input *CreateDetailParkirDto) SkpdkbDuplicate(sptDetail *Spt, skpdkb *SkpdkbExisting) error {
	if err := input.CreateDetailBaseDto.SkpdkbDuplicate(sptDetail, skpdkb); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &sptDetail.DetailSptParkir); err != nil {
		return err
	}
	return nil
}

type UpdateDetailParkirDto struct {
	UpdateDetailBaseDto
	DataDetails []mdsp.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailParkirDto) GetSpt(baseUri string) interface{} {
	return input.Spt
}

func (input *UpdateDetailParkirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailParkirDto) LenDetails() int {
	return len(input.DataDetails)
}
