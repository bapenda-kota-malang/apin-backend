package spt

import (
	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mdsh "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthotel"
	mdsrek "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	mdsjbr "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/jaminanbongkarreklame"
	"github.com/google/uuid"
)

// TODO: REKLAME GES
type CreateDetailReklameDto struct {
	CreateDetailBaseDto
	DataDetails           []mdsh.CreateDto `json:"dataDetails" validate:"required"`
	DetailSptReklame      mdsrek.CreateDto `json:"detailSptReklame"`
	JaminanBongkarReklame mdsjbr.CreateDto `json:"jaminanBongkarReklame"`
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
	input.Spt.JumlahPajak = float32(input.Spt.Omset * (*taxPercentage / 100))
}

func (input *CreateDetailReklameDto) DuplicateEspt(sptDetail *mespt.Espt) error {
	// do nothing because spt don't have reklame
	return nil
}

type UpdateDetailReklameDto struct {
	UpdateDetailBaseDto
	DataDetails mdsrek.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input *UpdateDetailReklameDto) CalculateTax(taxPercentage *float64) {
	// TODO: CHANGE THIS CALCULATION PROCESS
	input.Spt.JumlahPajak = float32(input.Spt.Omset * (*taxPercentage / 100))
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
