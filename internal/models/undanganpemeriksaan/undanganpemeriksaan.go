package undanganpemeriksaan

import (
	"time"

	mn "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/undanganpemeriksaan/pemeriksa"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type UndanganPemeriksaan struct {
	Id               uint64          `json:"id" gorm:"primaryKey"`
	JenisPajak       int             `json:"jenisPajak"`
	Tanggal          *time.Time      `json:"tanggal"`
	PukulAwal        *datatypes.Time `json:"pukulAwal"`
	PukulAkhir       *datatypes.Time `json:"pukulAkhir"`
	Tempat           *string         `json:"tempat"`
	Keperluan        *string         `json:"keperluan"`
	Petugas_User_Id  *uint64         `json:"petugas_user_id"`
	Petugas          *mu.User        `json:"petugas,omitempty" gorm:"foreignKey:Petugas_User_Id"`
	Status           Status          `json:"status"`
	StatusKirim      StatusKirim     `json:"statusKirim"`
	NoSuratUndangan  *string         `json:"noSuratUndangan"`
	Npwpd_Id         uint64          `json:"npwpd_id"`
	Npwpd            *mn.Npwpd       `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	ObjekPajakPbb_Id uint64          `json:"objekPajakPbb_id"`
	gh.DateModel
	Pemeriksa *[]pemeriksa.UndanganPemeriksaanPemeriksa `json:"pemeriksa,omitempty" gorm:"foreignKey:UndanganPemeriksaan_Id"`
}

type CreateDto struct {
	JenisPajak       int         `json:"jenisPajak" validate:"required"`
	Tanggal          *string     `json:"tanggal" validate:"required"`
	PukulAwal        string      `json:"pukulAwal" validate:"required"`
	PukulAkhir       string      `json:"pukulAkhir" validate:"required"`
	Tempat           *string     `json:"tempat" validate:"required"`
	Keperluan        *string     `json:"keperluan" validate:"required"`
	Petugas_User_Id  *uint64     `json:"petugas_user_id"`
	Status           Status      `json:"status"`
	StatusKirim      StatusKirim `json:"statusKirim"`
	NoSuratUndangan  *string     `json:"noSuratUndangan"`
	Npwpd_Id         uint64      `json:"npwpd_id" validate:"required"`
	ObjekPajakPbb_Id uint64      `json:"objekPajakPbb_id" validate:"min=1"`
	PemeriksaDetail  []uint64    `json:"pemeriksaDetail" validate:"required;min=1"`
}

type UpdateDto struct {
	JenisPajak       *int                  `json:"jenisPajak"`
	Tanggal          *string               `json:"tanggal"`
	PukulAwal        *string               `json:"pukulAwal"`
	PukulAkhir       *string               `json:"pukulAkhir"`
	Tempat           *string               `json:"tempat"`
	Keperluan        *string               `json:"keperluan"`
	Petugas_User_Id  *uint64               `json:"petugas_user_id"`
	Status           *Status               `json:"status"`
	StatusKirim      *StatusKirim          `json:"statusKirim"`
	NoSuratUndangan  *string               `json:"noSuratUndangan"`
	Npwpd_Id         *uint64               `json:"npwpd_id"`
	ObjekPajakPbb_Id *uint64               `json:"objekPajakPbb_id"`
	PemeriksaDetail  []pemeriksa.UpdateDto `json:"pemeriksaDetail"`
}

type FilterDto struct {
	JenisPajak      *int            `json:"jenisPajak"`
	Tanggal         *time.Time      `json:"tanggal"`
	PukulAwal       *datatypes.Time `json:"pukulAwal"`
	PukulAkhir      *datatypes.Time `json:"pukulAkhir"`
	Tempat          *string         `json:"tempat"`
	Keperluan       *string         `json:"keperluan"`
	Petugas_User_Id *uint64         `json:"petugas_user_id"`
	Status          *Status         `json:"status"`
	StatusKirim     *StatusKirim    `json:"statusKirim"`
	NoSuratUndangan *string         `json:"noSuratUndangan"`
	Npwpd_Id        uint64          `json:"npwpd_id"`
	Page            int             `json:"page"`
	PageSize        int             `json:"page_size"`
}

type UpdateStatusTerbitDto struct {
	Id []uint64 `json:"id"`
}
