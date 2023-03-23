package bidangkerja

import t "github.com/bapenda-kota-malang/apin-backend/internal/models/types"

type BidangKerja struct {
	Id        uint64         `json:"id" gorm:"primaryKey"`
	Level     t.LevelJabatan `json:"level"`
	Parent_Id uint64         `json:"parent_id"`
	Kode      *string        `json:"kode" gorm:"type:varchar(10)"`
	Nama      *string        `json:"nama" gorm:"type:varchar(100)"`
}

type CreateDto struct {
	Level     t.LevelJabatan `json:"level" validate:"required"`
	Parent_Id uint64         `json:"parent_id"`
	Kode      string         `json:"kode" validate:"required"`
	Nama      string         `json:"nama" validate:"required"`
}

type UpdateDto struct {
	Level     t.LevelJabatan `json:"level" validate:"required"`
	Parent_Id uint64         `json:"parent_id"`
	Kode      string         `json:"kode" validate:"required"`
	Nama      string         `json:"nama" validate:"required"`
}

type FilterDto struct {
	Kode         *string        `json:"kode"`
	Level        t.LevelJabatan `json:"level"`
	Parent_Id    uint64         `json:"parent_id"`
	Nama         *string        `json:"nama"`
	Page         int            `json:"page"`
	PageSize     int            `json:"page_size"`
	NoPagination bool           `json:"no_pagination"`
}
