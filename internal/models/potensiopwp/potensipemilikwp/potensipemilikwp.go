package potensipemilikwp

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type PotensiPemilikWp struct {
	Id           uint                    `json:"id" gorm:"primaryKey"`
	Potensiop_Id uint                    `json:"potensiop_id"`
	Nama         string                  `json:"nama" gorm:"size:50"`
	Alamat       string                  `json:"alamat" gorm:"size:50"`
	RtRw         string                  `json:"rtRw" gorm:"size:10"`
	Daerah_Id    uint                    `json:"daerah_id"`
	Kecamatan_Id *uint                   `json:"kecamatan_id"`
	Telp         *string                 `json:"telp" gorm:"size:20"`
	Status       t.StatusBL              `json:"status"`
	Nik          string                  `json:"nik" gorm:"size:20"`
	NoIdPemilik  *string                 `json:"noIdPemilik" gorm:"size:20"`
	Daerah       *areadivision.Daerah    `json:"daerah,omitempty" gorm:"foreignKey:Daerah_Id"`
	Kecamatan    *areadivision.Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Id"`
	gormhelper.DateModel
}

type CreateDto struct {
	Potensiop_Id uint       `json:"-"`
	Nama         string     `json:"nama" validate:"required"`
	Alamat       string     `json:"alamat" validate:"required"`
	RtRw         string     `json:"rtRw"`
	Daerah_Id    uint       `json:"daerah_id" validate:"required;min=1"`
	Kecamatan_Id *uint      `json:"kecamatan_id" validate:"min=1"`
	Telp         string     `json:"telp" validate:"nohp"`
	Status       t.StatusBL `json:"status" validate:"required"`
	Nik          *string    `json:"nik" validate:"nik"`
	NoIdPemilik  *string    `json:"noIdPemilik"`
}

type UpdateDto struct {
	Nama         *string     `json:"nama"`
	Alamat       *string     `json:"alamat"`
	RtRw         *string     `json:"rtRw"`
	Daerah_Id    *uint       `json:"daerah_id" validate:"min=1"`
	Kecamatan_Id *uint       `json:"kecamatan_id" validate:"min=1"`
	Telp         *string     `json:"telp" validate:"nohp"`
	Status       *t.StatusBL `json:"status"`
	Nik          *string     `json:"nik" validate:"nik"`
	NoIdPemilik  *string     `json:"noIdPemilik"`
}
