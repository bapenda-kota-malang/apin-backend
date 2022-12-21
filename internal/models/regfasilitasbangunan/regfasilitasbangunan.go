package regfasilitasbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegFasilitasBangunan struct {
	nop.NopDetail
	NoBangunan    *int    `json:"noBangunan"`
	KodeFasilitas *string `json:"kodeFasilitas" gorm:"type:char(2)"`
	JumlahSatuan  *int    `json:"jumlahSatuan"`
	gh.DateModel
}

type CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan    *int    `json:"noBangunan"`
	KodeFasilitas *string `json:"kodeFasilitas"`
	JumlahSatuan  *int    `json:"jumlahSatuan"`
}

type UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan    *int    `json:"noBangunan"`
	KodeFasilitas *string `json:"kodeFasilitas"`
	JumlahSatuan  *int    `json:"jumlahSatuan"`
}

type FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan    *int    `json:"noBangunan"`
	KodeFasilitas *string `json:"kodeFasilitas"`
	JumlahSatuan  *int    `json:"jumlahSatuan"`
	Page          int     `json:"page"`
	PageSize      int     `json:"page_size"`
}
