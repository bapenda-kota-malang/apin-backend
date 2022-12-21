package regkunjungankembali

import (
	"time"

	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

// untuk spop
type RegKunjunganKembali struct {
	nop.NopDetail
	NoBangunan              *int       `json:"noBangunan"`
	TanggalKunjunganKembali *time.Time `json:"tanggalKunjunganKembali"`
	Status                  *int       `json:"number"`
	gh.DateModel
}

type CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan              *int    `json:"noBangunan"`
	TanggalKunjunganKembali *string `json:"tanggalKunjunganKembali"`
	Status                  *int    `json:"status"`
}

type UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan              *int    `json:"noBangunan"`
	TanggalKunjunganKembali *string `json:"tanggalKunjunganKembali"`
	Status                  *int    `json:"status"`
}

type FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan              *int    `json:"noBangunan"`
	TanggalKunjunganKembali *string `json:"tanggalKunjunganKembali"`
	Status                  *int    `json:"status"`
	Page                    int     `json:"page"`
	PageSize                int     `json:"page_size"`
}
