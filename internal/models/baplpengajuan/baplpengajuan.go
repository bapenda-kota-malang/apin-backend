package baplpengajuan

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"

	mk "github.com/bapenda-kota-malang/apin-backend/internal/models/keberatan"
	mn "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	mp "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	mpengurangan "github.com/bapenda-kota-malang/apin-backend/internal/models/pengurangan"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type PengajuanBapl struct {
	Id                         uint64         `json:"id" gorm:"primaryKey"`
	Keberatan_Id               *uint64        `json:"keberatan_id"`
	TanggalKunjungan           *time.Time     `json:"tanggalKunjungan"`
	Npwpd_Id                   *uint64        `json:"npwpd_id"`
	Hasil                      *string        `json:"hasil" gorm:"type:varchar(255)"`
	PetugasLapangan_Pegawai_Id *pq.Int64Array `json:"petugasLapangan_pegawai_id" gorm:"type:integer[]"`
	Pegawai                    []mp.Pegawai   `json:"pegawai,omitempty" gorm:"-"`
	Status                     int            `json:"status"`
	VerifKasubid_User_Id       *uint64        `json:"verifKasubid_user_id"`
	EntryBy_User_Id            *uint64        `json:"entryBy_user_id"`
	Dokumentasi                string         `json:"dokumentasi"`
	DokumenLainnya             *string        `json:"dokumenLainnya"`
	JenisTransaksi             int            `json:"jenisTransaksi"`
	Pengurangan_Id             *uuid.UUID     `json:"pengurangan_id" gorm:"type:uuid"`
	gh.DateModel
	Keberatan   *mk.Keberatan             `json:"keberatan,omitempty" gorm:"foreignKey:Keberatan_Id"`
	Npwpd       *mn.Npwpd                 `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	KasubidUser *mu.User                  `json:"kasubidUser,omitempty" gorm:"foreignKey:VerifKasubid_User_Id"`
	User        *mu.User                  `json:"user,omitempty" gorm:"foreignKey:EntryBy_User_Id"`
	Pengurangan *mpengurangan.Pengurangan `json:"pengurangan,omitempty" gorm:"foreignKey:Pengurangan_Id"`
}

type CreateDto struct {
	Keberatan_Id               *uint64   `json:"keberatan_id"`
	TanggalKunjungan           *string   `json:"tanggalKunjungan"`
	Npwpd_Id                   *uint64   `json:"npwpd_id"`
	Hasil                      *string   `json:"hasil" gorm:"type:varchar(255)"`
	PetugasLapangan_Pegawai_Id []*uint64 `json:"petugasLapangan_pegawai_id"`
	Dokumentasi                string    `json:"dokumentasi"`
	DokumenLainnya             *string   `json:"dokumenLainnya"`
	Pengurangan_Id             *uint64   `json:"pengurangan_id"`
}

type UpdateDto struct {
	Keberatan_Id               *uint64   `json:"keberatan_id"`
	TanggalKunjungan           *string   `json:"tanggalKunjungan"`
	Npwpd_Id                   *uint64   `json:"npwpd_id"`
	Hasil                      *string   `json:"hasil" gorm:"type:varchar(255)"`
	PetugasLapangan_Pegawai_Id []*uint64 `json:"petugasLapangan_pegawai_id"`
	VerifKasubid_User_Id       *uint64   `json:"verifKasubid_user_id"`
	EntryBy_User_Id            *uint64   `json:"entryBy_user_id"`
	Pengurangan_Id             *uint64   `json:"pengurangan_id"`
}

type FilterDto struct {
	Keberatan_Id               *uint64 `json:"keberatan_id"`
	TanggalKunjungan           *string `json:"tanggalKunjungan"`
	Npwpd_Id                   *uint64 `json:"npwpd_id"`
	Hasil                      *string `json:"hasil" gorm:"type:varchar(255)"`
	PetugasLapangan_Pegawai_Id *uint64 `json:"petugasLapangan_pegawai_id"`
	VerifKasubid_User_Id       *uint64 `json:"verifKasubid_user_id"`
	EntryBy_User_Id            *uint64 `json:"entryBy_user_id"`
	Dokumentasi                *string `json:"dokumentasi"`
	DokumenLainnya             *string `json:"dokumenLainnya"`
	JenisTransaksi             *int    `json:"jenisTransaksi"`
	Pengurangan_Id             *uint64 `json:"pengurangan_id"`
	Page                       int     `json:"page"`
	PageSize                   int     `json:"page_size"`
}

type VerifyDto struct {
	Status int `json:"status" validate:"min=1;max=2"`
}
