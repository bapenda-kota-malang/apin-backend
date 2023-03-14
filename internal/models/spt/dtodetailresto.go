package spt

import (
	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mdsres "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type CreateDetailRestoDto struct {
	CreateDetailBaseDto
	DataDetails mdsres.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailRestoDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailRestoDto) LenDetails() int {
	newEmpty := mdsres.CreateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}

func (input *CreateDetailRestoDto) ReplaceSptId(id uuid.UUID) {
	input.DataDetails.Spt_Id = id
}

func (input *CreateDetailRestoDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	if err := input.CreateDetailBaseDto.DuplicateEspt(esptDetail); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &esptDetail.DetailEsptResto); err != nil {
		return err
	}
	return nil
}

func (input *CreateDetailRestoDto) SkpdkbDuplicate(sptDetail *Spt, skpdkb *SkpdkbExisting) error {
	if err := input.CreateDetailBaseDto.SkpdkbDuplicate(sptDetail, skpdkb); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &sptDetail.DetailSptResto); err != nil {
		return err
	}
	return nil
}

type UpdateDetailRestoDto struct {
	UpdateDetailBaseDto
	DataDetails mdsres.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailRestoDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailRestoDto) LenDetails() int {
	newEmpty := mdsres.UpdateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}
