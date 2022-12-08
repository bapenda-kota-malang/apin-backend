package objekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb15 struct {
	nop.NopDetail
	NoBangunan     *int       `json:"noBangunan"`
	LetakTanki     LetakTanki `json:"letakTanki" gorm:"type:char(1)"`
	KapasitasTanki *int       `json:"kapasitasTanki"`
	gh.DateModel
}

type Jpb15CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan     *int       `json:"noBangunan"`
	LetakTanki     LetakTanki `json:"letakTanki"`
	KapasitasTanki *int       `json:"kapasitasTanki"`
}

type Jpb15UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan     *int       `json:"noBangunan"`
	LetakTanki     LetakTanki `json:"letakTanki"`
	KapasitasTanki *int       `json:"kapasitasTanki"`
}

type Jpb15FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan     *int `json:"noBangunan"`
	KapasitasTanki *int `json:"kapasitasTanki"`
	Page           int  `json:"page"`
	PageSize       int  `json:"page_size"`
}
