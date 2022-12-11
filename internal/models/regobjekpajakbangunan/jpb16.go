package regobjekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegJpb16 struct {
	nop.NopDetail
	NoBangunan      *int          `json:"noBangunan"`
	KelasBangunan16 KelasBangunan `json:"kelasBangunan16" gorm:"type:char(1)"`
	gh.DateModel
}

type RegJpb16CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan      *int          `json:"noBangunan"`
	KelasBangunan16 KelasBangunan `json:"kelasBangunan16"`
}

type RegJpb16UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan      *int          `json:"noBangunan"`
	KelasBangunan16 KelasBangunan `json:"kelasBangunan16"`
}

type RegJpb16FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}
