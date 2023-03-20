package spt

import (
	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mdppjnonpln "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptppjnonpln"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type CreateDetailPpjNonPlnDto struct {
	CreateDetailBaseDto
	DataDetails mdppjnonpln.CreateDto `json:"dataDetails" validate:"required"`
}

func (input *CreateDetailPpjNonPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailPpjNonPlnDto) LenDetails() int {
	newEmpty := mdppjnonpln.CreateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}

func (input *CreateDetailPpjNonPlnDto) ReplaceSptId(id uuid.UUID) {
	input.DataDetails.Spt_Id = id
}

func (input *CreateDetailPpjNonPlnDto) ChangeDetails(newData interface{}) {
	// Do nothing
}

func (input *CreateDetailPpjNonPlnDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	if err := input.CreateDetailBaseDto.DuplicateEspt(esptDetail); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &esptDetail.DetailEsptPpjNonPln); err != nil {
		return err
	}
	return nil
}

func (input *CreateDetailPpjNonPlnDto) SkpdkbDuplicate(sptDetail *Spt, skpdkb *SkpdkbExisting) error {
	if err := input.CreateDetailBaseDto.SkpdkbDuplicate(sptDetail, skpdkb); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &sptDetail.DetailSptNonPln); err != nil {
		return err
	}
	return nil
}

type UpdateDetailPpjNonPlnDto struct {
	UpdateDetailBaseDto
	DataDetails mdppjnonpln.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailPpjNonPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailPpjNonPlnDto) LenDetails() int {
	newEmpty := mdppjnonpln.UpdateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}
