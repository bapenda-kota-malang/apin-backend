package jenisperolehan

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type JenisPerolehan struct {
	Id uint64 `json:"id" gorm:"primaryKey"`
	gh.DateModel
	KodePerolehan *string `json:"kodePerolehan"`
	Nama          *string `json:"nama"`
}

type CreateDto struct {
	KodePerolehan *string `json:"kodePerolehan"`
	Nama          *string `json:"nama"`
}

type UpdateDto struct {
	Id            *uint64 `json:"id"`
	KodePerolehan *string `json:"kodePerolehan"`
	Nama          *string `json:"nama"`
}

type FilterDto struct {
	KodePerolehan *string `json:"kodePerolehan"`
	Nama          *string `json:"nama"`
	Page          int     `json:"page"`
	PageSize      int     `json:"page_size"`
}
