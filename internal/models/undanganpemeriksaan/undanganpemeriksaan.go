package undanganpemeriksaan

import (
	"time"

	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type UndanganPemeriksaan struct {
	Id                    uint64  `json:"id" gorm:"primaryKey"`
	JenisPajak            int     `json:"jenisPajak"`
	SuratPemberitahuan_Id *uint64 `json:"suratPemberitahuan_id"`
	// SuratPembertahuan msp.SuratPemberitahuan `json:"suratPemberitahuan,omitempty" gorm:"foreignKey:SuratPemberitahuan_Id"`
	SuratPemberitahuanPbb_Id *uint64     `json:"suratPemberitahuanPbb_id"`
	Tanggal                  *time.Time  `json:"tanggal"`
	Keperluan                *string     `json:"keperluan"`
	Petugas_User_Id          *uint64     `json:"petugas_user_id"`
	Petugas                  *mu.User    `json:"petugas,omitempty" gorm:"foreignKey:Petugas_User_Id"`
	Status                   Status      `json:"status"`
	StatusKirim              StatusKirim `json:"statusKirim"`
	NoSuratUndangan          *string     `json:"noSuratUndangan"`
	gh.DateModel
}

type CreateDto struct {
	JenisPajak               int         `json:"jenisPajak"`
	SuratPemberitahuan_Id    *uint64     `json:"suratPemberitahuan_id"`
	SuratPemberitahuanPbb_Id *uint64     `json:"suratPemberitahuanPbb_id"`
	Tanggal                  *string     `json:"tanggal"`
	Keperluan                *string     `json:"keperluan"`
	Petugas_User_Id          *uint64     `json:"petugas_user_id"`
	Status                   Status      `json:"status"`
	StatusKirim              StatusKirim `json:"statusKirim"`
	NoSuratUndangan          *string     `json:"noSuratUndangan"`
}

type UpdateDto struct {
	JenisPajak               int         `json:"jenisPajak"`
	SuratPemberitahuan_Id    *uint64     `json:"suratPemberitahuan_id"`
	SuratPemberitahuanPbb_Id *uint64     `json:"suratPemberitahuanPbb_id"`
	Tanggal                  *string     `json:"tanggal"`
	Keperluan                *string     `json:"keperluan"`
	Petugas_User_Id          *uint64     `json:"petugas_user_id"`
	Status                   Status      `json:"status"`
	StatusKirim              StatusKirim `json:"statusKirim"`
	NoSuratUndangan          *string     `json:"noSuratUndangan"`
}

type FilterDto struct {
	JenisPajak               int         `json:"jenisPajak"`
	SuratPemberitahuan_Id    *uint64     `json:"suratPemberitahuan_id"`
	SuratPemberitahuanPbb_Id *uint64     `json:"suratPemberitahuanPbb_id"`
	Tanggal                  *string     `json:"tanggal"`
	Keperluan                *string     `json:"keperluan"`
	Petugas_User_Id          *uint64     `json:"petugas_user_id"`
	Status                   Status      `json:"status"`
	StatusKirim              StatusKirim `json:"statusKirim"`
	NoSuratUndangan          *string     `json:"noSuratUndangan"`
	Page                     int         `json:"page"`
	PageSize                 int         `json:"page_size"`
}

type UpdateStatusTerbitDto struct {
	Id []uint64 `json:"id"`
}
