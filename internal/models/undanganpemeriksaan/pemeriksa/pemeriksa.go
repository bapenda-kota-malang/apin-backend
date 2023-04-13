package pemeriksa

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type UndanganPemeriksaanPemeriksa struct {
	Id                     uint64 `json:"id" gorm:"primaryKey"`
	UndanganPemeriksaan_Id uint64 `json:"undanganPemeriksaan_id" gorm:"index"`
	Petugas_Id             uint64 `json:"petugas_id"`
	gormhelper.DateModel
	Petugas *pegawai.Pegawai `json:"petugas,omitempty" gorm:"foreignKey:Petugas_Id"`
}

type CreateDto struct {
	UndanganPemeriksaan_Id uint64 `json:"-"`
	Petugas_Id             uint64 `json:"petugas_id"`
}

type UpdateDto struct {
	UndanganPemeriksaan_Id uint64 `json:"-"`
	Petugas_Id             uint64 `json:"petugas_id" validate:"required"`
	Deleted                bool   `json:"deleted"`
}
