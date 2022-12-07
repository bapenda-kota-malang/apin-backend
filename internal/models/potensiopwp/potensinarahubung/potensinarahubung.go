package potensinarahubung

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type PotensiNarahubung struct {
	Id           uint                    `json:"id" gorm:"primaryKey"`
	Potensiop_Id uuid.UUID               `json:"potensiop_id" gorm:"type:uuid"`
	Nama         string                  `json:"nama" gorm:"size:50"`
	Alamat       string                  `json:"alamat" gorm:"size:50"`
	RtRw         string                  `json:"rtRw" gorm:"size:10"`
	Daerah_Id    uint                    `json:"daerah_id"`
	Kelurahan_Id *uint                   `json:"kelurahan_id"`
	Telp         *string                 `json:"telp" gorm:"size:20"`
	Status       t.StatusBL              `json:"status"`
	Nik          string                  `json:"nik" gorm:"size:20"`
	Email        *string                 `json:"email"`
	Daerah       *areadivision.Daerah    `json:"daerah,omitempty" gorm:"foreignKey:Daerah_Id"`
	Kelurahan    *areadivision.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Id"`
	gormhelper.DateModel
}

type CreateDto struct {
	Potensiop_Id uuid.UUID  `json:"-"`
	Nama         string     `json:"nama" validate:"required"`
	Alamat       string     `json:"alamat" validate:"required"`
	RtRw         string     `json:"rtRw"`
	Daerah_Id    uint       `json:"daerah_id" validate:"required;min=1"`
	Kelurahan_Id *uint      `json:"kelurahan_id" validate:"min=1"`
	Telp         string     `json:"telp" validate:"required"`
	Status       t.StatusBL `json:"status" validate:"required"`
	Nik          string     `json:"nik" validate:"nik"`
	Email        *string    `json:"email" validate:"email"`
}

type UpdateDto struct {
	Id           *uint       `json:"id"`
	Nama         *string     `json:"nama"`
	Alamat       *string     `json:"alamat"`
	RtRw         *string     `json:"rtRw"`
	Daerah_Id    *uint       `json:"daerah_id" validate:"min=1"`
	Kelurahan_Id *uint       `json:"kelurahan_id" validate:"min=1"`
	Telp         *string     `json:"telp"`
	Status       *t.StatusBL `json:"status"`
	Nik          *string     `json:"nik" validate:"nik"`
	Email        *string     `json:"email" validate:"email"`
	Delete       *bool       `json:"delete"`
}
