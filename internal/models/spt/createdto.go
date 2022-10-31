package spt

import (
	mdsa "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptair"
	mdsh "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthotel"
	mdsp "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	mdsrek "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	mdsres "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"
	mdsjbr "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/jaminanbongkarreklame"
)

type CreateDetailAirDto struct {
	Spt         CreateDto      `json:"spt" validate:"required"`
	DataDetails mdsa.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailAirDto) GetSpt() CreateDto {
	input.Spt.LuasLokasi = nil
	return input.Spt
}

func (input CreateDetailAirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailAirDto) LenDetails() int {
	return 1
}

func (input *CreateDetailAirDto) ReplaceSptId(id uint) {
	input.DataDetails.Spt_Id = id
}

type CreateDetailHotelDto struct {
	Spt         CreateDto        `json:"spt" validate:"required"`
	DataDetails []mdsh.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailHotelDto) GetSpt() CreateDto {
	input.Spt.LuasLokasi = nil
	return input.Spt
}

func (input CreateDetailHotelDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailHotelDto) LenDetails() int {
	return 1
}

func (input *CreateDetailHotelDto) ReplaceSptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Spt_Id = id
	}
}

type CreateDetailParkirDto struct {
	Spt         CreateDto        `json:"spt" validate:"required"`
	DataDetails []mdsp.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailParkirDto) GetSpt() CreateDto {
	input.Spt.LuasLokasi = nil
	return input.Spt
}

func (input CreateDetailParkirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailParkirDto) LenDetails() int {
	return 1
}

func (input *CreateDetailParkirDto) ReplaceSptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Spt_Id = id
	}
}

type CreateDetailReklameDto struct {
	Spt                   CreateDto        `json:"spt" validate:"required"`
	DataDetails           []mdsh.CreateDto `json:"dataDetails" validate:"required"`
	DetailSptReklame      mdsrek.CreateDto `json:"detailSptReklame"`
	JaminanBongkarReklame mdsjbr.CreateDto `json:"jaminanBongkarReklame"`
}

func (input CreateDetailReklameDto) GetSpt() CreateDto {
	input.Spt.LuasLokasi = nil
	return input.Spt
}

func (input CreateDetailReklameDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailReklameDto) LenDetails() int {
	return 1
}

func (input *CreateDetailReklameDto) ReplaceSptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Spt_Id = id
	}
}

type CreateDetailRestoDto struct {
	Spt         CreateDto        `json:"spt" validate:"required"`
	DataDetails mdsres.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailRestoDto) GetSpt() CreateDto {
	input.Spt.LuasLokasi = nil
	return input.Spt
}

func (input CreateDetailRestoDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailRestoDto) LenDetails() int {
	return 1
}

func (input *CreateDetailRestoDto) ReplaceSptId(id uint) {
	input.DataDetails.Spt_Id = id
}
