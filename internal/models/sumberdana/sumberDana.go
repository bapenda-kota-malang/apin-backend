package sumberdana

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
)

type SumberDana struct {
	Id             uint64             `json:"id" gorm:"primaryKey"`
	Rekening_Id    *uint64            `json:"rekening_id"`
	Rekening       *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	Nama           *string            `json:"nama" gorm:"type:varchar(100)"`
	Bank           *string            `json:"bank" gorm:"type:varchar(100)"`
	NoRekeningBank *string            `json:"noRekeningBank" gorm:"type:varchar(100)"`
	UpdatedAt      time.Time          `json:"updatedAt"`
}

type CreateDto struct {
	Rekening_Id    uint64 `json:"rekening_id" validate:"required"`
	Nama           string `json:"nama" validate:"required"`
	Bank           string `json:"bank" validate:"required"`
	NoRekeningBank string `json:"noRekeningBank" validate:"required"`
}

type UpdateDto struct {
	Rekening_Id    uint64 `json:"rekening_id"`
	Nama           string `json:"nama"`
	Bank           string `json:"bank"`
	NoRekeningBank string `json:"noRekeningBank"`
}

type FilterDto struct {
	Rekening_Id    *uint64 `json:"rekening_id"`
	Nama           *string `json:"nama"`
	Bank           *string `json:"bank"`
	NoRekeningBank *string `json:"noRekeningBank"`
	Page           int     `json:"page"`
	PageSize       int     `json:"page_size"`
}
