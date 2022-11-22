package pengurangan

import (
	"time"

	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pengurangan struct {
	Id                    uint64          `json:"id" gorm:"primaryKey"`
	Spt_Id                *uuid.UUID      `json:"spt_id" gorm:"type:uuid"`
	Pemohon_User_Id       *uint64         `json:"pemohon_user_id"`
	NominalKetetapan      *float64        `json:"nominalKetetapan"`
	AlasanPengurangan     *string         `json:"alasanPengurangan" gorm:"type:varchar(255)"`
	SuratPermohonan       *string         `json:"suratPermohonan" gorm:"size:2048"`
	FotoKtp               *string         `json:"fotoKtp" gorm:"size:2048"`
	LaporanKeuangan       *string         `json:"laporanKeuangan" gorm:"size:2048"`
	LaporanPengeluaran    *string         `json:"laporanPengeluaran" gorm:"size:2048"`
	DokumenLainnya        *string         `json:"dokumenLainnya" gorm:"size:2048"`
	PersentasePengurangan *float64        `json:"persentasePengurangan"`
	Status                *int            `json:"status"`
	VerifKasubid_User_Id  *uint64         `json:"verifKasubid_user_id"`
	VerifKabid_User_Id    *uint64         `json:"verifKabid_user_id"`
	VerifKaban_User_Id    *uint64         `json:"verifKaban_user_id"`
	VerifPetugas_User_Id  *uint64         `json:"verifPetugas_user_id"`
	TanggalPengajuan      *time.Time      `json:"tanggalPengajuan"`
	UpdatedAt             *time.Time      `json:"updatedAt"`
	DeletedAt             *gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Spt                   *ms.Spt         `json:"spt,omitempty" gorm:"foreignKey:Spt_Id"`
	Pemohon               *mu.User        `json:"pemohon" gorm:"foreignKey:Pemohon_User_Id"`
	KasubidUser           *mu.User        `json:"kasubidUser" gorm:"foreignKey:VerifKasubid_User_Id"`
	KabidUser             *mu.User        `json:"kabidUser" gorm:"foreignKey:VerifKabid_User_Id"`
	KabanUser             *mu.User        `json:"kabanUser" gorm:"foreignKey:VerifKaban_User_Id"`
	PetugasUser           *mu.User        `json:"petugasUser" gorm:"foreignKey:VerifPetugas_User_Id"`
}

type PenguranganCreateDto struct {
	Spt_Id                *uuid.UUID      `json:"spt_id"`
	Pemohon_User_Id       *uint64         `json:"pemohon_user_id"`
	NominalKetetapan      *float64        `json:"nominalKetetapan"`
	AlasanPengurangan     *string         `json:"alasanPengurangan"`
	SuratPermohonan       *string         `json:"suratPermohonan"`
	FotoKtp               *string         `json:"fotoKtp"`
	LaporanKeuangan       *string         `json:"laporanKeuangan"`
	LaporanPengeluaran    *string         `json:"laporanPengeluaran"`
	DokumenLainnya        *string         `json:"dokumenLainnya"`
	PersentasePengurangan *float64        `json:"persentasePengurangan"`
	Status                *int            `json:"status"`
	VerifKasubid_User_Id  *uint64         `json:"verifKasubid_user_id"`
	VerifKabid_User_Id    *uint64         `json:"verifKabid_user_id"`
	VerifKaban_User_Id    *uint64         `json:"verifKaban_user_id"`
	VerifPetugas_User_Id  *uint64         `json:"verifPetugas_user_id"`
	TanggalPengajuan      *time.Time      `json:"tanggalPengajuan"`
	UpdatedAt             *time.Time      `json:"updatedAt"`
	DeletedAt             *gorm.DeletedAt `json:"deletedAt"`
}
