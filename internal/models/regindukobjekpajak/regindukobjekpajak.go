package regindukobjekpajak

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegIndukObjekPajak struct {
	nop.NopDetail
	gh.DateModel
}

type CreateDto struct {
	nop.NopDetailCreateDto
}

type UpdateDto struct {
	nop.NopDetailUpdateDto
}

type FilterDto struct {
	nop.NopDetailCreateDto
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
