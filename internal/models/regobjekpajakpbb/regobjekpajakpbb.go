package regobjekpajakpbb

import (
	"time"

	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	p "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	maop "github.com/bapenda-kota-malang/apin-backend/internal/models/reganggotaobjekpajak"
	miop "github.com/bapenda-kota-malang/apin-backend/internal/models/regindukobjekpajak"
	mkk "github.com/bapenda-kota-malang/apin-backend/internal/models/regkunjungankembali"
	mopb "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbumi"
	ot "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakpbb/types"
	mwp "github.com/bapenda-kota-malang/apin-backend/internal/models/regwajibpajakpbb"
	"gorm.io/gorm"
)

// spop
type RegObjekPajakPbb struct {
	nop.NopDetail
	RegWajibPajakPbb_Id   *uint64               `json:"regWajibPajakPbb_id"`
	RegWajibPajakPbb      *mwp.RegWajibPajakPbb `json:"regWajibPajakPbb,omitempty" gorm:"foreignKey:RegWajibPajakPbb_Id"`
	NoFormulirSpop        *string               `json:"noFormulirSpop" gorm:"type:char(11)"`
	NoPersil              *string               `json:"noPersil" gorm:"type:varchar(5)"`
	Jalan                 *string               `json:"jalan" gorm:"type:varchar(30)"`
	BlokKavNo             *string               `json:"blokKavNo" gorm:"type:varchar(15)"`
	Rw                    *string               `json:"rw" gorm:"type:char(2)"`
	Rt                    *string               `json:"rt" gorm:"type:char(3)"`
	StatusCabang          *int                  `json:"statusCabang"`
	StatusWp              ot.StatusWp           `json:"statusWp" gorm:"type:char(1)"`
	TotalLuasBumi         *int                  `json:"totalLuasBumi"`
	TotalLuasBangunan     *int                  `json:"totalLuasBangunan"`
	NjopBumi              *int                  `json:"njopBumi"`
	NjopBangunan          *int                  `json:"njopBangunan"`
	StatusPeta            *int                  `json:"statusPeta"`
	Status                VerifyStatus          `json:"status"`
	NopAsal               *string               `json:"nopAsal"`
	JenisTransaksi        ot.JenisTransaksi     `json:"jenisTransaksi" gorm:"type:varchar(2)"`
	TanggalPendataan      *time.Time            `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string               `json:"pendata_pegawai_nip" gorm:"type:char(9)"`
	PegawaiPendata        *p.Pegawai            `json:"pegawaiPendata,omitempty" gorm:"foreignKey:Pendata_Pegawai_Nip;references:Nip"`
	TanggalPemeriksaan    *time.Time            `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string               `json:"pemeriksa_pegawai_nip" gorm:"type:char(9)"`
	PegawaiPemeriksa      *p.Pegawai            `json:"pegawaiPemeriksa,omitempty" gorm:"foreignKey:Pemeriksa_Pegawai_Nip;references:Nip"`
	TanggalPerekaman      *time.Time            `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string               `json:"perekam_pegawai_nip" gorm:"type:char(9)"`
	PegawaiPerekam        *p.Pegawai            `json:"pegawaiPerekam,omitempty" gorm:"foreignKey:Perekam_Pegawai_Nip;references:Nip"`
	UpdatedAt             *time.Time            `json:"updatedAt"`
	DeletedAt             gorm.DeletedAt        `json:"deletedAt" gorm:"index"`
}

type CreateDto struct {
	nop.NopDetailCreateDto
	RegWajibPajakPbb_Id   *uint64           `json:"regWajibPajakPbb_id"`
	NoFormulirSpop        *string           `json:"noFormulirSpop"`
	NoPersil              *string           `json:"noPersil"`
	Jalan                 *string           `json:"jalan"`
	BlokKavNo             *string           `json:"blokKavNo"`
	Rw                    *string           `json:"rw"`
	Rt                    *string           `json:"rt"`
	StatusCabang          *int              `json:"statusCabang"`
	StatusWp              ot.StatusWp       `json:"statusWp"`
	TotalLuasBumi         *int              `json:"totalLuasBumi"`
	TotalLuasBangunan     *int              `json:"totalLuasBangunan"`
	StatusPeta            *int              `json:"statusPeta"`
	JenisTransaksi        ot.JenisTransaksi `json:"jenisTransaksi"`
	TanggalPendataan      *string           `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string           `json:"pendata_pegawai_nip"`
	TanggalPemeriksaan    *string           `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string           `json:"pemeriksa_pegawai_nip"`
	TanggalPerekaman      *string           `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string           `json:"perekam_pegawai_nip"`
	//data wajibpajakpbb
	RegWajibPajakPbbs     mwp.CreateDto   `json:"regWajibPajakPbb"`
	RegObjekPajakBumis    mopb.CreateDto  `json:"regObjekPajakBumi"`
	RegAnggotaObjekPajaks *maop.CreateDto `json:"regAnggotaObjekPajak"`
	RegIndukObjekPajaks   *miop.CreateDto `json:"regIndukObjekPajak"`
	RegKunjunganKembalis  *mkk.CreateDto  `json:"regKunjunganKembali"`

	//nop
	Nop        *string `json:"nop"`
	NopBersama *string `json:"nopBersama"`
	NopAsal    *string `json:"nopAsal"`
}

type UpdateDto struct {
	nop.NopDetailUpdateDto
	RegWajibPajakPbb_Id   *uint64           `json:"regWajibPajakPbb_id"`
	NoFormulirSpop        *string           `json:"noFormulirSpop"`
	NoPersil              *string           `json:"noPersil"`
	Jalan                 *string           `json:"jalan"`
	BlokKavNo             *string           `json:"blokKavNo"`
	Rw                    *string           `json:"rw"`
	Rt                    *string           `json:"rt"`
	StatusCabang          *int              `json:"statusCabang"`
	StatusWp              ot.StatusWp       `json:"statusWp"`
	TotalLuasBumi         *int              `json:"totalLuasBumi"`
	TotalLuasBangunan     *int              `json:"totalLuasBangunan"`
	NjopBumi              *int              `json:"njopBumi"`
	NjopBangunan          *int              `json:"njopBangunan"`
	StatusPeta            *int              `json:"statusPeta"`
	JenisTransaksi        ot.JenisTransaksi `json:"jenisTransaksi"`
	TanggalPendataan      *string           `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string           `json:"pendata_pegawai_nip"`
	TanggalPemeriksaan    *string           `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string           `json:"pemeriksa_pegawai_nip"`
	TanggalPerekaman      *string           `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string           `json:"perekam_pegawai_nip"`
	//data wajibpajakpbb
	RegWajibPajakPbbs mwp.CreateDto `json:"wajibPajakPbb"`
}

type FilterDto struct {
	Nop                   *string            `json:"nop"`
	Provinsi_Kode         *string            `json:"provinsi_kode" gorm:"type:char(2)"`
	Kota_Kode             *string            `json:"kota_kode" gorm:"type:char(2)"`
	Kecamatan_Kode        *string            `json:"kecamatan_kode" gorm:"type:char(3)"`
	Kelurahan_Kode        *string            `json:"kelurahan_kode" gorm:"type:char(3)"`
	Blok_Id               *string            `json:"blok_id" gorm:"type:char(3)"`
	NoUrut                *string            `json:"noUrut" gorm:"type:char(4)"`
	JenisOp               *string            `json:"jenisOp" gorm:"type:char(1)"`
	WajibPajakPbb_Id      *uint64            `json:"WajibPajakPbb_id"`
	NoFormulirSpop        *string            `json:"noFormulirSpop"`
	NoPersil              *string            `json:"noPersil"`
	Jalan                 *string            `json:"jalan"`
	BlokKavNo             *string            `json:"blokKavNo"`
	Rw                    *string            `json:"rw"`
	Rt                    *string            `json:"rt"`
	StatusCabang          *int               `json:"statusCabang"`
	StatusWp              *ot.StatusWp       `json:"statusWp"`
	TotalLuasBumi         *int               `json:"totalLuasBumi"`
	TotalLuasBangunan     *int               `json:"totalLuasBangunan"`
	NjopBumi              *int               `json:"njopBumi"`
	NjopBangunan          *int               `json:"njopBangunan"`
	StatusPeta            *int               `json:"statusPeta"`
	JenisTransaksi        *ot.JenisTransaksi `json:"jenisTransaksi"`
	TanggalPendataan      *time.Time         `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string            `json:"pendata_pegawai_nip"`
	TanggalPemeriksaan    *time.Time         `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string            `json:"pemeriksa_pegawai_nip"`
	TanggalPerekaman      *time.Time         `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string            `json:"perekam_pegawai_nip"`
	Page                  int                `json:"page"`
	PageSize              int                `json:"page_size"`
}

type VerifyDto struct {
	Status VerifyStatus `json:"status" validate:"min=1;max=2"`
}
