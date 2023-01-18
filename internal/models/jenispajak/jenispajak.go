package jenispajak

import "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"

type JenisPajak struct {
	Id          uint64             `json:"id" gorm:"primaryKey"`
	Kode        *string            `json:"kode" gorm:"type:varchar(5);unique"`
	Uraian      *string            `json:"uraian" gorm:"type:varchar(100)"`
	Rekening_Id *uint64            `json:"rekening_id"`
	Rekening    *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	KodeBilling *string            `json:"kodeBilling" gorm:"type:varchar(20)"`
	KodeVAJatim *string            `json:"kodeVAJatim" gorm:"type:varchar(100)"`
}

type CreateDto struct {
	Kode        *string `json:"kode" gorm:"type:varchar(5)"`
	Uraian      *string `json:"uraian" gorm:"type:varchar(100)"`
	Rekening_Id *uint64 `json:"rekening_id"`
	KodeBilling *string `json:"kodeBilling" gorm:"type:varchar(20)"`
	KodeVAJatim *string `json:"kodeVAJatim" gorm:"type:varchar(100)"`
}

type UpdateDto struct {
	Id          uint64  `json:"id"`
	Kode        *string `json:"kode"`
	Uraian      *string `json:"uraian"`
	Rekening_Id *uint64 `json:"rekening_id"`
	KodeBilling *string `json:"kodeBilling"`
	KodeVAJatim *string `json:"kodeVAJatim"`
}

type FilterDto struct {
	Kode        *string `json:"kode"`
	Uraian      *string `json:"uraian"`
	Rekening_Id *uint64 `json:"rekening_id"`
	KodeBilling *string `json:"kodeBilling"`
	KodeVAJatim *string `json:"kodeVAJatim"`
	Page        int     `json:"page"`
	PageSize    int     `json:"page_size"`
}
