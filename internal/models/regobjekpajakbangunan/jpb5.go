package regobjekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegJpb5 struct {
	nop.NopDetail
	NoBangunan             *int          `json:"noBangunan"`
	KelasBangunan5         KelasBangunan `json:"kelasBangunan5" gorm:"type:char(1)"`
	LuasKamarAcCentral     *int          `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
	gh.DateModel
}

type RegJpb5CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan             *int          `json:"noBangunan"`
	KelasBangunan5         KelasBangunan `json:"kelasBangunan5"`
	LuasKamarAcCentral     *int          `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
}

type RegJpb5UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan             *int          `json:"noBangunan"`
	KelasBangunan5         KelasBangunan `json:"kelasBangunan5"`
	LuasKamarAcCentral     *int          `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
}

type RegJpb5FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan             *int `json:"noBangunan"`
	LuasKamarAcCentral     *int `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int `json:"luasRuangLainAcCentral"`
	Page                   int  `json:"page"`
	PageSize               int  `json:"page_size"`
}
