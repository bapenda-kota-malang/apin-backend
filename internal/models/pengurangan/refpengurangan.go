package pengurangan

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type RefPengurangan struct {
	Id                    uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Spt_Id                uuid.UUID  `json:"spt_id" gorm:"type:uuid;not null"`
	Pengurangan_Id        *uuid.UUID `json:"-" gorm:"type:uuid"`
	Pemohon_User_Id       uint64     `json:"pemohon_user_id"`
	NominalKetetapan_Spt  float64    `json:"nominalKetetapan_spt" gorm:"not null"`
	PctPermohonan         float64    `json:"pctPermohonan"`
	AlasanPengurangan     *string    `json:"alasanPengurangan"`
	SuratPermohonan       string     `json:"suratPermohonan" gorm:"size:2048"`
	FotoKtp               string     `json:"fotoKtp" gorm:"size:2048"`
	LaporanKeuangan       string     `json:"laporanKeuangan" gorm:"size:2048"`
	DokumenLainnya        *string    `json:"dokumenLainnya" gorm:"size:2048"`
	TanggalPengajuan      time.Time  `json:"tanggalPengajuan"`
	Status                Status     `json:"status" gorm:"comment:0: Proses; 1:Diterima; 2:Ditolak"`
	KeteranganStaff       *string    `json:"KeteranganStaff"`
	AlasanPenolakanStaff  *string    `json:"alasanPenolakanStaff"`
	VerifyPetugas_User_Id *uint64    `json:"verifPetugas_user_id"`
	TanggalVerifPetugas   *time.Time `json:"tanggalVerifPetugas"`
	VerifyAnalis_User_Id  *uint64    `json:"verifyAnalis_user_id"`
	TanggalVerifAnalis    *time.Time `json:"tanggalVerifAnalis"`
	VerifyKasubid_User_Id *uint64    `json:"verifyKasubid_user_id"`
	TanggalVerifKasubid   *time.Time `json:"tanggalVerifKasubid"`
	VerifyKabid_User_Id   *uint64    `json:"verifyKabid_user_id"`
	TanggalVerifKabid     *time.Time `json:"tanggalVerifKabid"`
	VerifySekban_User_Id  *uint64    `json:"verifySekban_user_id"`
	TanggalVerifSekban    *time.Time `json:"tanggalVerifSekban"`
	VerifyKaban_User_Id   *uint64    `json:"verifyKaban_user_id"`
	TanggalVerifKaban     *time.Time `json:"tanggalVerifKaban"`
	gormhelper.DateModel
	Spt         *spt.Spt     `json:"spt,omitempty" gorm:"foreignKey:Spt_Id"`
	Pengurangan *Pengurangan `json:"pengurangan,omitempty" gorm:"foreignKey:Pengurangan_Id"`
	Pemohon     *user.User   `json:"pemohon,omitempty" gorm:"foreignKey:Pemohon_User_Id"`
	Petugas     *user.User   `json:"petugas,omitempty" gorm:"foreignKey:VerifyPetugas_User_Id"`
	Analis      *user.User   `json:"analis,omitempty" gorm:"foreignKey:VerifyAnalis_User_Id"`
	Kasubid     *user.User   `json:"kasubid,omitempty" gorm:"foreignKey:VerifyKasubid_User_Id"`
	Kabid       *user.User   `json:"kabid,omitempty" gorm:"foreignKey:VerifyKabid_User_Id"`
	Sekban      *user.User   `json:"sekban,omitempty" gorm:"foreignKey:VerifySekban_User_Id"`
	Kaban       *user.User   `json:"kaban,omitempty" gorm:"foreignKey:VerifyKaban_User_Id"`
}

type CreateDtoRefPengurangan struct {
	Spt_Id            uuid.UUID `json:"spt_id" validate:"required"`
	AlasanPengurangan *string   `json:"alasanPengurangan"`
	PctPermohonan     *float64  `json:"pctPermohonan" validate:"required"`
	SuratPermohonan   string    `json:"suratPermohonan" validate:"required;base64=pdf;b64size=1025"`
	FotoKtp           string    `json:"fotoKtp" validate:"required;base64=image;b64size=1025"`
	LaporanKeuangan   string    `json:"laporanKeuangan" validate:"required;base64=pdf;b64size=1025"`
	DokumenLainnya    *string   `json:"dokumenLainnya" validate:"base64=pdf;b64size=1025"`
}

type UpdateDtoRefPengurangan struct {
	AlasanPengurangan *string `json:"alasanPengurangan"`
	SuratPermohonan   *string `json:"suratPermohonan"`
	FotoKtp           *string `json:"fotoKtp"`
	LaporanKeuangan   *string `json:"laporanKeuangan"`
	DokumenLainnya    *string `json:"dokumenLainnya"`
}

type VerifyDtoRefPengurangan struct {
	JenisPengurangan   *JenisPengurangan `json:"jenisPengurangan" validate:"required"`
	KeteranganPetugas  *string           `json:"keteranganPetugas"`
	AlasanDitolakStaff *string           `json:"alasanDitolakStaff"`
	StatusVerifikasi   *StatusVerifikasi `json:"statusVerifikasi" validate:"required"`
}

type FilterDtoRefPengurangan struct {
	Pemohon_User_Id       *uint64    `json:"pemohon_user_id"`
	NominalKetetapan_Spt  *float64   `json:"nominalKetetapan_spt"`
	AlasanPengurangan     *string    `json:"alasanPengurangan"`
	TanggalPengajuan      *time.Time `json:"tanggalPengajuan"`
	Status                *Status    `json:"status"`
	KeteranganStaff       *string    `json:"KeteranganStaff"`
	AlasanPenolakanStaff  *string    `json:"alasanPenolakanStaff"`
	VerifyPetugas_User_Id *uint64    `json:"verifPetugas_user_id"`
	TanggalVerifPetugas   *time.Time `json:"tanggalVerifPetugas"`
	VerifyAnalis_User_Id  *uint64    `json:"verifyAnalis_user_id"`
	TanggalVerifAnalis    *time.Time `json:"tanggalVerifAnalis"`
	VerifyKasubid_User_Id *uint64    `json:"verifyKasubid_user_id"`
	TanggalVerifKasubid   *time.Time `json:"tanggalVerifKasubid"`
	VerifyKabid_User_Id   *uint64    `json:"verifyKabid_user_id"`
	TanggalVerifKabid     *time.Time `json:"tanggalVerifKabid"`
	VerifySekban_User_Id  *uint64    `json:"verifySekban_user_id"`
	TanggalVerifSekban    *time.Time `json:"tanggalVerifSekban"`
	VerifyKaban_User_Id   *uint64    `json:"verifyKaban_user_id"`
	TanggalVerifKaban     *time.Time `json:"tanggalVerifKaban"`
	Page                  int        `json:"page"`
	PageSize              int        `json:"page_size"`
}
