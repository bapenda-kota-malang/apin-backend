package regobjekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegJpb2 struct {
	nop.NopDetail
	NoBangunan     *int          `json:"noBangunan"`
	KelasBangunan2 KelasBangunan `json:"kelasBangunan2" gorm:"type:char(1)"`
	gh.DateModel
}

type RegJpb2CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan     *int          `json:"noBangunan"`
	KelasBangunan2 KelasBangunan `json:"kelasBangunan2"`
}

type RegJpb2UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan     *int          `json:"noBangunan"`
	KelasBangunan2 KelasBangunan `json:"kelasBangunan2"`
}

type RegJpb2FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}
