package refumum

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RefUmum struct {
	Id uint64 `json:"id" gorm:"primaryKey"`
	gh.DateModel
	Kode       int    `json:"kode" gorm:"type:integer"`
	Keterangan string `json:"keterangan" gorm:"type:text"`
	Nama       string `json:"nama" gorm:"type:varchar(255)"`
}

type CreateDto struct {
	Kode       int    `json:"kode"`
	Keterangan string `json:"keterangan"`
	Nama       string `json:"nama"`
}

type UpdateDto struct {
	Id         *uint64 `json:"id"`
	Kode       int     `json:"kode"`
	Keterangan string  `json:"keterangan"`
	Nama       string  `json:"nama"`
}

type FilterDto struct {
	Kode       int    `json:"kode"`
	Keterangan string `json:"keterangan"`
	Nama       string `json:"nama"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
}
