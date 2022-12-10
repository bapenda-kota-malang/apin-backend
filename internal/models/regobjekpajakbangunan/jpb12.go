package regobjekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegJpb12 struct {
	nop.NopDetail
	NoBangunan   int          `json:"noBangunan"`
	TipeBangunan TipeBangunan `json:"tipeBangunan" gorm:"type:char(1)"`
	gh.DateModel
}

type RegJpb12CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan   *int         `json:"noBangunan"`
	TipeBangunan TipeBangunan `json:"tipeBangunan"`
}

type RegJpb12UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan   *int         `json:"noBangunan"`
	TipeBangunan TipeBangunan `json:"tipeBangunan"`
}

type RegJpb12FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan   *int         `json:"noBangunan"`
	TipeBangunan TipeBangunan `json:"tipeBangunan"`
	Page         int          `json:"page"`
	PageSize     int          `json:"page_size"`
}
