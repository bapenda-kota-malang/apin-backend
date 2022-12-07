package jpb12

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb12 struct {
	nop.NopDetail
	NoBangunan   int          `json:"noBangunan"`
	TipeBangunan TipeBangunan `json:"tipeBangunan" gorm:"type:char(1)"`
	gh.DateModel
}

type CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan   *int         `json:"noBangunan"`
	TipeBangunan TipeBangunan `json:"tipeBangunan"`
}

type UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan   *int         `json:"noBangunan"`
	TipeBangunan TipeBangunan `json:"tipeBangunan"`
}

type FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan   *int         `json:"noBangunan"`
	TipeBangunan TipeBangunan `json:"tipeBangunan"`
	Page         int          `json:"page"`
	PageSize     int          `json:"page_size"`
}
