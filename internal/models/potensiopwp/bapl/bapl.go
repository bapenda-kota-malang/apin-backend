package bapl

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type PotensiBapl struct {
	Id                     uint                   `json:"id" gorm:"primaryKey"`
	TanggalPeninjauan      time.Time              `json:"tanggalPeninjauan"`
	Potensiop_Id           uuid.UUID              `json:"potensiop_id" gorm:"type:uuid"`
	CreateBy_Pegawai_Id    uint                   `json:"createBy_pegawai_id"`
	Koordinator_Pegawai_Id uint                   `json:"koordinator_pegawai_id"`
	Petugas_Pegawai_Id     *pq.Int64Array         `json:"petugas_pegawai_id" gorm:"type:integer[]"`
	Status                 types.StatusVerifikasi `json:"status"`
	Kasubid_Pegawai_Id     *int                   `json:"kasubid_pegawai_id,omitempty"`
	Kabid_Pegawai_Id       *int                   `json:"kabid_pegawai_id,omitempty"`
	gormhelper.DateModel
	CreateBy    *pegawai.Pegawai   `json:"createBy,omitempty" gorm:"foreignKey:CreateBy_Pegawai_Id"`
	Koordinator *pegawai.Pegawai   `json:"koordinator,omitempty" gorm:"foreignKey:Koordinator_Pegawai_Id"`
	Kasubid     *pegawai.Pegawai   `json:"kasubid,omitempty" gorm:"foreignKey:Kasubid_Pegawai_Id"`
	Kabid       *pegawai.Pegawai   `json:"kabid,omitempty" gorm:"foreignKey:Kabid_Pegawai_Id"`
	Petugas     *[]pegawai.Pegawai `json:"petugas,omitempty" gorm:"-"`
}

type CreateDto struct {
	TanggalPeninjauan      *time.Time             `json:"tanggalPeninjauan"`
	Potensiop_Id           uuid.UUID              `json:"-"`
	CreateBy_Pegawai_Id    uint                   `json:"-"`
	Koordinator_Pegawai_Id uint                   `json:"koordinator_pegawai_id"`
	Petugas_Pegawai_Id     *pq.Int64Array         `json:"petugas_pegawai_id"`
	Status                 types.StatusVerifikasi `json:"-"`
}

type UpdateDto struct {
	TanggalPeninjauan      *time.Time     `json:"tanggalPeninjauan"`
	Koordinator_Pegawai_Id *uint          `json:"koordinator_pegawai_id"`
	Petugas_Pegawai_Id     *pq.Int64Array `json:"petugas_pegawai_id"`
}

type VerifyDto struct {
	Status             string `json:"status" validate:"required"`
	Kasubid_Pegawai_Id *uint  `json:"-"`
	Kabid_Pegawai_Id   *uint  `json:"-"`
}
