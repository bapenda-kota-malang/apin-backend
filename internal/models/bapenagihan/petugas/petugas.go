package petugas

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type BaPenagihanPetugas struct {
	Id             uint64    `json:"id" gorm:"primaryKey"`
	BaPenagihan_Id uuid.UUID `json:"baPenagihan_id" gorm:"type:uuid"`
	Petugas_Id     uint64    `json:"petugas_id"`
	gormhelper.DateModel
	Petugas *pegawai.Pegawai `json:"petugas,omitempty" gorm:"foreignKey:Petugas_Id"`
}

type CreateDto struct {
	BaPenagihan_Id uuid.UUID `json:"-"`
	Petugas_Id     uint64    `json:"petugas_id"`
}

type UpdateDto struct {
	BaPenagihan_Id uuid.UUID `json:"-"`
	Petugas_Id     uint64    `json:"petugas_id" validate:"required"`
	Deleted        bool      `json:"deleted"`
}
