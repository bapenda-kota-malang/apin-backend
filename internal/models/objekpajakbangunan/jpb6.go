package objekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb6 struct {
	nop.NopDetail
	NoBangunan   *int          `json:"noBangunan"`
	KelasBanguna KelasBangunan `json:"kelasBangunan6" gorm:"type:char(1)"`
	gh.DateModel
}

type Jpb6CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan     *int          `json:"noBangunan"`
	KelasBangunan6 KelasBangunan `json:"kelasBangunan6"`
}

type Jpb6UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan     *int          `json:"noBangunan"`
	KelasBangunan6 KelasBangunan `json:"kelasBangunan6"`
}

type Jpb6FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}
