package pengurangan

import (
	"time"

	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type Pengurangan struct {
	Id                uuid.UUID        `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Spt_Id            uuid.UUID        `json:"spt_id" gorm:"type:uuid;not null"`
	NamaPemohon       string           `json:"namaPemohon" gorm:"not null"`
	AlamatPemohon     string           `json:"alamatPemohon" gorm:"not null"`
	TelpPemohon       *string          `json:"telpPemohon"`
	JenisPengurangan  JenisPengurangan `json:"jenisPengurangan" gorm:"not null"`
	AlasanPengurangan *string          `json:"alasanPengurangan"`
	SuratPermohonan   string           `json:"suratPermohonan" gorm:"size:2048;not null"`
	FotoKtp           string           `json:"fotoKtp" gorm:"size:2048;not null"`
	LaporanKeuangan   string           `json:"laporanKeuangan" gorm:"size:2048;not null"`
	Lhp               *string          `json:"lhp" gorm:"size:2048"`
	TelaahStaff       *string          `json:"telaahStaff" gorm:"size:2048"`
	DokumenLainnya    *string          `json:"dokumenLainnya" gorm:"size:2048"`
	TanggalPengajuan  time.Time        `json:"tanggalPengajuan"`
	Posisi            PosisiVerifikasi `json:"posisi"`
	Status            Status           `json:"status"`
	PctPermohonan     float64          `json:"pctPermohonan"`
	// verifikasi berjenjang
	VerifyPetugas_User_Id *uint64           `json:"verifPetugas_user_id"`
	KeteranganPetugas     *string           `json:"keteranganPetugas"`
	TanggalVerifPetugas   *time.Time        `json:"tanggalVerifPetugas"`
	VerifyAnalis_User_Id  *uint64           `json:"verifyAnalis_user_id"`
	StatusAnalis          *StatusVerifikasi `json:"statusAnalis"`
	PersentaseAnalis      *int8             `json:"persentaseAnalis"`
	KeteranganAnalis      *string           `json:"keteranganAnalis"`
	AlasanDitolakAnalis   *string           `json:"alasanDitolakAnalis"`
	TanggalVerifAnalis    *time.Time        `json:"tanggalVerifAnalis"`
	VerifyKasubid_User_Id *uint64           `json:"verifyKasubid_user_id"`
	StatusKasubid         *StatusVerifikasi `json:"statusKasubid"`
	PersentaseKasubid     *int8             `json:"persentaseKasubid"`
	KeteranganKasubid     *string           `json:"keteranganKasubid"`
	AlasanDitolakKasubid  *string           `json:"alasanDitolakKasubid"`
	TanggalVerifKasubid   *time.Time        `json:"tanggalVerifKasubid"`
	VerifyKabid_User_Id   *uint64           `json:"verifyKabid_user_id"`
	StatusKabid           *StatusVerifikasi `json:"statusKabid"`
	PersentaseKabid       *int8             `json:"persentaseKabid"`
	KeteranganKabid       *string           `json:"keteranganKabid"`
	AlasanDitolakKabid    *string           `json:"alasanDitolakKabid"`
	TanggalVerifKabid     *time.Time        `json:"tanggalVerifKabid"`
	VerifySekban_User_Id  *uint64           `json:"verifySekban_user_id"`
	StatusSekban          *StatusVerifikasi `json:"statusSekban"`
	PersentaseSekban      *int8             `json:"persentaseSekban"`
	KeteranganSekban      *string           `json:"keteranganSekban"`
	AlasanDitolakSekban   *string           `json:"alasanDitolakSekban"`
	TanggalVerifSekban    *time.Time        `json:"tanggalVerifSekban"`
	VerifyKaban_User_Id   *uint64           `json:"verifyKaban_user_id"`
	StatusKaban           *StatusVerifikasi `json:"statusKaban"`
	PersentaseKaban       *int8             `json:"persentaseKaban"`
	KeteranganKaban       *string           `json:"keteranganKaban"`
	AlasanDitolakKaban    *string           `json:"alasanDitolakKaban"`
	TanggalVerifKaban     *time.Time        `json:"tanggalVerifKaban"`
	gormhelper.DateModel
	Spt     *ms.Spt  `json:"spt,omitempty" gorm:"foreignKey:Spt_Id"`
	Petugas *mu.User `json:"petugas,omitempty" gorm:"foreignKey:VerifyPetugas_User_Id"`
	Analis  *mu.User `json:"analis,omitempty" gorm:"foreignKey:VerifyAnalis_User_Id"`
	Kasubid *mu.User `json:"kasubid,omitempty" gorm:"foreignKey:VerifyKasubid_User_Id"`
	Kabid   *mu.User `json:"kabid,omitempty" gorm:"foreignKey:VerifyKabid_User_Id"`
	Sekban  *mu.User `json:"sekban,omitempty" gorm:"foreignKey:VerifySekban_User_Id"`
	Kaban   *mu.User `json:"kaban,omitempty" gorm:"foreignKey:VerifyKaban_User_Id"`
}

// initial create by pegawai / staff
type CreateDto struct {
	Spt_Id            uuid.UUID        `json:"spt_id" validate:"required"`
	NamaPemohon       string           `json:"namaPemohon" validate:"required"`
	AlamatPemohon     string           `json:"alamatPemohon" validate:"required"`
	TelpPemohon       *string          `json:"telpPemohon" validate:"nohp"`
	JenisPengurangan  JenisPengurangan `json:"jenisPengurangan"`
	AlasanPengurangan *string          `json:"alasanPengurangan" validate:"required"`
	TanggalPengajuan  *time.Time       `json:"tanggalPengajuan"`
	PctPermohonan     *float64         `json:"pctPermohonan" validate:"required"`
	LaporanKeuangan   string           `json:"laporanKeuangan" validate:"required;base64=pdf;b64size=1025"`
	SuratPermohonan   string           `json:"suratPermohonan" validate:"required;base64=pdf;b64size=1025"`
	FotoKtp           string           `json:"fotoKtp" validate:"required;base64=image;b64size=1025"`
	DokumenLainnya    *string          `json:"dokumenLainnya" validate:"base64=pdf;b64size=1025"`
	KeteranganPetugas *string          `json:"keteranganPetugas"`
}

// create pengurangan data from ref pengurangan by pegawai / staff
type CreateFromRefDto struct {
	RefPengurangan_Id  *uuid.UUID        `json:"refPengurangan_id" validate:"required"`
	JenisPengurangan   *JenisPengurangan `json:"jenisPengurangan" validate:"required"`
	KeteranganPetugas  *string           `json:"keteranganPetugas"`
	AlasanDitolakStaff *string           `json:"alasanDitolakStaff"`
	Status             *Status           `json:"status" validate:"required"`
}

type UpdateDto struct {
	NamaPemohon       *string           `json:"namaPemohon"`
	AlamatPemohon     *string           `json:"alamatPemohon"`
	TelpPemohon       *string           `json:"telpPemohon"`
	JenisPengurangan  *JenisPengurangan `json:"jenisPengurangan"`
	AlasanPengurangan *string           `json:"alasanPengurangan"`
	TanggalPengajuan  *time.Time        `json:"tanggalPengajuan"`
	Posisi            *PosisiVerifikasi `json:"posisi"`
	Status            *Status           `json:"status"`
	PctPermohonan     *float64          `json:"pctPermohonan"`
	LaporanKeuangan   *string           `json:"laporanKeuangan"`
	SuratPermohonan   *string           `json:"suratPermohonan"`
	FotoKtp           *string           `json:"fotoKtp"`
	DokumenLainnya    *string           `json:"dokumenLainnya"`
	KeteranganPetugas *string           `json:"keteranganPetugas"`
}

type FilterDto struct {
	Spt_Id                          *uuid.UUID       `json:"spt_id"`
	NamaPemohon                     string           `json:"namaPemohon"`
	AlamatPemohon                   string           `json:"alamatPemohon"`
	TelpPemohon                     *string          `json:"telpPemohon"`
	JenisPengurangan                JenisPengurangan `json:"jenisPengurangan"`
	AlasanPengurangan               *string          `json:"alasanPengurangan"`
	LaporanKeuangan                 *string          `json:"laporanKeuangan"`
	DokumenLainnya                  *string          `json:"dokumenLainnya"`
	SuratPermohonan                 *string          `json:"suratPermohonan"`
	FotoKtp                         *string          `json:"fotoKtp"`
	Lhp                             *string          `json:"lhp"`
	TelaahStaff                     *string          `json:"telaahStaff"`
	TanggalPengajuan                *time.Time       `json:"tanggalPengajuan"`
	Posisi                          PosisiVerifikasi `json:"posisi"`
	Status                          Status           `json:"status"`
	PersentasePermohonanPengurangan *float64         `json:"persentasePermohonanPengurangan"`
	VerifyPetugas_User_Id           *uint64          `json:"verifPetugas_user_id"`
	KeteranganPetugas               *string          `json:"keteranganPetugas"`
	TanggalVerifPetugas             *time.Time       `json:"tanggalVerifPetugas"`
	VerifyAnalis_User_Id            *uint64          `json:"verifyAnalis_user_id"`
	StatusAnalis                    StatusVerifikasi `json:"statusAnalis"`
	PersentaseAnalis                *int8            `json:"persentaseAnalis"`
	KeteranganAnalis                *string          `json:"keteranganAnalis"`
	AlasanDitolakAnalis             *string          `json:"alasanDitolakAnalis"`
	TanggalVerifAnalis              *time.Time       `json:"tanggalVerifAnalis"`
	VerifyKasubid_User_Id           *uint64          `json:"verifyKasubid_user_id"`
	StatusKasubid                   StatusVerifikasi `json:"statusKasubid"`
	PersentaseKasubid               *int8            `json:"persentaseKasubid"`
	KeteranganKasubid               *string          `json:"keteranganKasubid"`
	AlasanDitolakKasubid            *string          `json:"alasanDitolakKasubid"`
	TanggalVerifKasubid             *time.Time       `json:"tanggalVerifKasubid"`
	VerifyKabid_User_Id             *uint64          `json:"verifyKabid_user_id"`
	StatusKabid                     StatusVerifikasi `json:"statusKabid"`
	PersentaseKabid                 *int8            `json:"persentaseKabid"`
	KeteranganKabid                 *string          `json:"keteranganKabid"`
	AlasanDitolakKabid              *string          `json:"alasanDitolakKabid"`
	TanggalVerifKabid               *time.Time       `json:"tanggalVerifKabid"`
	VerifySekban_User_Id            *uint64          `json:"verifySekban_user_id"`
	StatusSekban                    StatusVerifikasi `json:"statusSekban"`
	PersentaseSekban                *int8            `json:"persentaseSekban"`
	KeteranganSekban                *string          `json:"keteranganSekban"`
	AlasanDitolakSekban             *string          `json:"alasanDitolakSekban"`
	TanggalVerifSekban              *time.Time       `json:"tanggalVerifSekban"`
	VerifyKaban_User_Id             *uint64          `json:"verifyKaban_user_id"`
	StatusKaban                     StatusVerifikasi `json:"statusKaban"`
	PersentaseKaban                 *int8            `json:"persentaseKaban"`
	KeteranganKaban                 *string          `json:"keteranganKaban"`
	AlasanDitolakKaban              *string          `json:"alasanDitolakKaban"`
	TanggalVerifKaban               *time.Time       `json:"tanggalVerifKaban"`
	Page                            int              `json:"page"`
	PageSize                        int              `json:"page_size"`
}

type VerifyDto struct {
	Lhp              *string           `json:"lhp" validate:"base64=pdf;b64size=1025"`
	TelaahStaff      *string           `json:"telaahStaff" validate:"base64=pdf;b64size=1025"`
	StatusVerifikasi *StatusVerifikasi `json:"statusVerifikasi" validate:"required"`
	Persentase       *int8             `json:"persentase"`
	Keterangan       *string           `json:"keterangan"`
	AlasanDitolak    *string           `json:"alasanDitolak"`
}
