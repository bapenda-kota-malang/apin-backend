package jenisperolehan

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type JenisPerolehan struct {
	Id uint64 `json:"id" gorm:"primaryKey"`
	gh.DateModel
	Kode *string `json:"kode"`
	Nama *string `json:"nama"`
}

type CreateDto struct {
	Kode *string `json:"kode"`
	Nama *string `json:"nama"`
}

type UpdateDto struct {
	Id   *uint64 `json:"id"`
	Kode *string `json:"kode"`
	Nama *string `json:"nama"`
}

type FilterDto struct {
	Kode     *string `json:"kode"`
	Nama     *string `json:"nama" reffunc:"LOWER"`
	Nama_Opt *string `json:"nama_opt"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}
