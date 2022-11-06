package spt

import (
	mdsa "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptair"
	mdsh "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthotel"
	mdsp "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	mdsrek "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	mdsres "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"
	mdsjbr "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/jaminanbongkarreklame"
)

type UpdateAirDto struct {
	Spt         UpdateDto      `json:"spt" validate:"required"`
	DataDetails mdsa.UpdateDto `json:"dataDetails" validate:"required"`
}

type UpdateHotelDto struct {
	Spt         UpdateDto        `json:"spt" validate:"required"`
	DataDetails []mdsh.UpdateDto `json:"dataDetails" validate:"required"`
}

type UpdateParkirDto struct {
	Spt         UpdateDto        `json:"spt" validate:"required"`
	DataDetails []mdsp.UpdateDto `json:"dataDetails" validate:"required"`
}

type UpdateReklameDto struct {
	Spt                   UpdateDto        `json:"spt" validate:"required"`
	DataDetails           []mdsh.UpdateDto `json:"dataDetails" validate:"required"`
	DetailSptReklame      mdsrek.UpdateDto `json:"detailSptReklame"`
	JaminanBongkarReklame mdsjbr.UpdateDto `json:"jaminanBongkarReklame"`
}

type UpdateRestoDto struct {
	Spt         UpdateDto        `json:"spt" validate:"required"`
	DataDetails mdsres.UpdateDto `json:"dataDetails" validate:"required"`
}
