package referensibank

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type ReferensiBank struct {
	Id uint64 `json:"id" gorm:"primaryKey"`
	gh.DateModel
	Nama  *string `json:"nama"`
	Nomor *string `json:"nomor"`
}

type CreateDto struct {
	Nama  *string `json:"nama"`
	Nomor *string `json:"nomor"`
}

type UpdateDto struct {
	Id    *uint64 `json:"id"`
	Nama  *string `json:"nama"`
	Nomor *string `json:"nomor"`
}

type FilterDto struct {
	Nama     *string `json:"nama"`
	Nomor    *string `json:"nomor"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}
