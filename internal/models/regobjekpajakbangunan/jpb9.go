package regobjekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegJpb9 struct {
	nop.NopDetail
	NoBangunan     *int          `json:"noBangunan"`
	KelasBangunan9 KelasBangunan `json:"kelasBangunan9" gorm:"type:char(1)"`
	gh.DateModel
}

type RegJpb9CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan     *int          `json:"noBangunan"`
	KelasBangunan9 KelasBangunan `json:"kelasBangunan9"`
}

type RegJpb9UpdateDto struct {
	nop.NopDetailCreateDto
	NoBangunan     *int          `json:"noBangunan"`
	KelasBangunan9 KelasBangunan `json:"kelasBangunan9"`
}

type RegJpb9FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}
