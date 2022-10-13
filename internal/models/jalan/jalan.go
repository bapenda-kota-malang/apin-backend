package jalan

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/klasifikasijalan"
)

type Jalan struct {
	Id                  uint64                             `json:"id" gorm:"primaryKey"`
	KlasifikasiJalan_Id *string                            `json:"klasifikasiJalan_id" gorm:"type:varchar(3)"`
	KlasifikasiJalan    *klasifikasijalan.KlasifikasiJalan `json:"rekening,omitempty" gorm:"foreignKey:KlasifikasiJalan_Id"`
	Nama                *string                            `json:"nama" gorm:"type:varchar(100)"`
}

type CreateDto struct {
	KlasifikasiJalan_Id string `json:"klasifikasiJalan_id" validate:"required"`
	Nama                string `json:"nama" validate:"required"`
}

type UpdateDto struct {
	KlasifikasiJalan_Id string `json:"klasifikasiJalan_id"`
	Nama                string `json:"nama"`
}

type FilterDto struct {
	KlasifikasiJalan_Id *string `json:"klasifikasiJalan_id"`
	Nama                *string `json:"nama"`
	Page                int     `json:"page"`
	PageSize            int     `json:"page_size"`
}
