package jpb6

import (
	jt "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb2/types"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb6 struct {
	nop.NopDetail
	NoBangunan     *int              `json:"noBangunan"`
	KelasBangunan6 *jt.KelasBangunan `json:"kelasBangunan6" gorm:"type:char(1)"`
	gh.DateModel
}

type CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan     *int              `json:"noBangunan"`
	KelasBangunan6 *jt.KelasBangunan `json:"kelasBangunan6"`
}

type UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan     *int              `json:"noBangunan"`
	KelasBangunan6 *jt.KelasBangunan `json:"kelasBangunan6"`
}

type FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}
