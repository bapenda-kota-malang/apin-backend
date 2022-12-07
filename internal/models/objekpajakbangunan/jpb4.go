package objekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb4 struct {
	nop.NopDetail
	NoBangunan     *int          `json:"noBangunan"`
	KelasBangunan4 KelasBangunan `json:"kelasBangunan4" gorm:"type:char(1)"`
	gh.DateModel
}

type Jpb4CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan     *int          `json:"noBangunan"`
	KelasBangunan4 KelasBangunan `json:"kelasBangunan4"`
}

type Jpb4UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan     *int          `json:"noBangunan"`
	KelasBangunan4 KelasBangunan `json:"kelasBangunan4"`
}

type Jpb4FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}
