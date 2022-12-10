package regobjekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegJpb14 struct {
	nop.NopDetail
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
	gh.DateModel
}

type RegJpb14CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
}

type RegJpb14UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
}

type RegJpb14FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}
