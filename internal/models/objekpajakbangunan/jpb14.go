package objekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb14 struct {
	nop.NopDetail
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
	gh.DateModel
}

type Jpb14CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
}

type Jpb14UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
}

type Jpb14FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}
