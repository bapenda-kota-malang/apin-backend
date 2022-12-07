package dbkbfasum

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbFasum struct {
	Id             uint64 `json:"id" gorm:"primaryKey"`
	Kode           string `json:"kode" gorm:"unique;type:char(2)"`
	Nama           string `json:"nama" gorm:"size:50"`
	Satuan         string `json:"satuan" gorm:"size:10"`
	Status         string `json:"status" gorm:"type:char(1)"`
	Ketergantungan string `json:"ketergantungan" gorm:"type:char(1)"`
	gormhelper.DateModel
}

type CreateDto struct {
	Kode           string `json:"kode" validate:"required;minLength=2;maxLength=2"`
	Nama           string `json:"nama" validate:"required;maxLength=50"`
	Satuan         string `json:"satuan" validate:"required;maxLength=10"`
	Status         string `json:"status" validate:"required;minLength=1;maxLength=1"`
	Ketergantungan string `json:"ketergantungan" validate:"required;minLength=1;maxLength=1"`
}

type FilterDto struct {
	Kode           *string `json:"kode"`
	Nama           *string `json:"nama"`
	Satuan         *string `json:"satuan"`
	Status         *string `json:"status"`
	Ketergantungan *string `json:"ketergantungan"`
	Page           int     `json:"page"`
	PageSize       int     `json:"page_size"`
	NoPagination   bool    `json:"no_pagination"`
}
