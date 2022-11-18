package konfigurasipajak

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
)

type KonfigurasiPajak struct {
	Id          uint64             `json:"id" gorm:"primaryKey"`
	Rekening_Id *uint64            `json:"rekening_id"`
	Rekening    *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	Oa          *string            `json:"oa"`
	Sa          *string            `json:"sa"`
	UrlDaftarOa *string            `json:"urlDaftarOa" gorm:"type:varchar(100)"`
	UrlDaftarSa *string            `json:"urlDaftarSa" gorm:"type:varchar(100)"`
}

type CreateDto struct {
	Rekening_Id *uint64 `json:"rekening_id"`
	Oa          *string `json:"oa"`
	Sa          *string `json:"sa"`
	UrlDaftarOa *string `json:"urlDaftarOa"`
	UrlDaftarSa *string `json:"urlDaftarSa"`
}

type UpdateDto struct {
	Id          *uint64 `json:"id"`
	Rekening_Id *uint64 `json:"rekening_id"`
	Oa          *string `json:"oa"`
	Sa          *string `json:"sa"`
	UrlDaftarOa *string `json:"urlDaftarOa"`
	UrlDaftarSa *string `json:"urlDaftarSa"`
}

type FilterDto struct {
	Rekening_Id *uint64 `json:"rekening_id"`
	Oa          *string `json:"oa"`
	Sa          *string `json:"sa"`
	UrlDaftarOa *string `json:"urlDaftarOa"`
	UrlDaftarSa *string `json:"urlDaftarSa"`
	Page        int     `json:"page"`
	PageSize    int     `json:"page_size"`
}
