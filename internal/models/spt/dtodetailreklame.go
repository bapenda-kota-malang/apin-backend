package spt

import (
	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mdsrek "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	mdsjbr "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/jaminanbongkarreklame"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

// TODO: REKLAME GES
type CreateDetailReklameDto struct {
	CreateDetailBaseDto
	DataDetails           []mdsrek.CreateDto `json:"dataDetails" validate:"required"`
	JaminanBongkarReklame mdsjbr.CreateDto   `json:"jaminanBongkarReklame"`
}

func (input *CreateDetailReklameDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *CreateDetailReklameDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input *CreateDetailReklameDto) ReplaceSptId(id uuid.UUID) {
	for i := range input.DataDetails {
		input.DataDetails[i].Spt_Id = id
	}
}

func (input *CreateDetailReklameDto) CalculateTax(taxPercentage *float64) {
	// TODO: CHANGE THIS CALCULATION PROCESS
	input.Spt.JumlahPajak = input.Spt.Omset * (*taxPercentage / 100)
}

func (input *CreateDetailReklameDto) DuplicateEspt(sptDetail *mespt.Espt) error {
	// do nothing because spt don't have reklame
	return nil
}

func (input *CreateDetailReklameDto) SkpdkbDuplicate(sptDetail *Spt, skpdkb *SkpdkbExisting) error {
	if err := input.CreateDetailBaseDto.SkpdkbDuplicate(sptDetail, skpdkb); err != nil {
		return err
	}
	if err := copier.Copy(&input.DataDetails, &sptDetail); err != nil {
		return err
	}
	return nil
}

type UpdateDetailReklameDto struct {
	UpdateDetailBaseDto
	DataDetails mdsrek.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailReklameDto) CalculateTax(taxPercentage *float64) {
	// TODO: CHANGE THIS CALCULATION PROCESS
	calc := *input.Spt.Omset * *taxPercentage / 100
	input.Spt.JumlahPajak = &calc
}

func (input *UpdateDetailReklameDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input *UpdateDetailReklameDto) LenDetails() int {
	newEmpty := mdsrek.UpdateDto{}
	lenData := 1
	if input.DataDetails == newEmpty {
		lenData = 0
	}
	return lenData
}
