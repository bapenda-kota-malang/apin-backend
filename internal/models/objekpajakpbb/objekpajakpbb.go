package objekpajakpbb

import (
	"time"

	ad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	p "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	"gorm.io/gorm"
)

type ObjekPajakPbb struct {
	Id           uint64        `json:"id" gorm:"primarykey"`
	Provinsi_Id  string        `json:"provinsi_id" gorm:"type:char(2)"`
	Provinsi     *ad.Provinsi  `json:"provinsi,omitempty" gorm:"foreignKey:Provinsi_Id;references:Id"`
	Kota_Id      string        `json:"kota_id" gorm:"type:char(2)"`
	Kota         *ad.Daerah    `json:"kota,omitempty" gorm:"foreignKey:Kota_Id;references:Id"`
	Kecamatan_Id string        `json:"kecamatan_id" gorm:"type:char(3)"`
	Kecamatan    *ad.Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Id;references:Id"`
	Kelurahan_Id string        `json:"kelurahan_id" gorm:"type:char(3)"`
	Kelurahan    *ad.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Id;references:Id"`
	// Blok_Id      string        `json:"blok_id" gorm:"type:char(3)"`
	// Blok         *ad.Blok      `json:"blok,omitempty" gorm:"foreignKey:Blok_Id;references:Id"`
	NoUrut         string `json:"noUrut" gorm:"type:char(4)"`
	JenisOp_Id     string `json:"jenisOp_id" gorm:"type:char(1)"`
	Nik            string `json:"nik" gorm:"type:char(30)"`
	NoFormulirSpop string `json:"noFormulirSpop" gorm:"type:char(11)"`
	NoPersil       string `json:"noPersil" gorm:"type:varchar(5)"`
	Jalan          string `json:"jalan" gorm:"type:varchar(30)"`
	BlokKavNo      string `json:"blokKavNo" gorm:"type:varchar(15)"`
	Rw             string `json:"rw" gorm:"type:char(2)"`
	Rt             string `json:"rt" gorm:"type:char(3)"`
	// StatusCabang_Id int `json:"statusCabang_id"`
	// StatusCabang StatusCabang `json:"statusCabang,omitempty" gorm:"foreignKey:StatusCabang_Id;references:Id"`
	// StatusWp_Id string `json:"statusWp_id" gorm:"type:char(1)"`
	// StatusWp StatusWp `json:"statusWp,omitempty" gorm:"foreignKey:StatusWp_Id;references:Id"`
	TotalLuasBumi         int            `json:"totalLuasBumi"`
	TotalLuasBangunan     int            `json:"totalLuasBangunan"`
	NjopBumi              int            `json:"njopBumi"`
	NjopBangunan          int            `json:"njopBangunan"`
	StatusPeta            int            `json:"statusPeta"`
	JenisTransaksi        JenisTransaksi `json:"jenisTransaksi" gorm:"type:char(1)"`
	TanggalPendataan      *time.Time     `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   string         `json:"pendata_pegawai_nip" gorm:"type:char(9)"`
	Pegawai               *p.Pegawai     `json:"pegawai,omitempty" gorm:"foreignKey:Pendata_Pegawai_Nip;references:Nip"`
	TanggalPemeriksaan    *time.Time     `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip string         `json:"pemeriksa_pegawai_nip" gorm:"type:char(9)"`
	TanggalPerekaman      *time.Time     `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   string         `json:"perekam_pegawai_nip" gorm:"type:char(9)"`
	UpdatedAt             *time.Time     `json:"updatedAt"`
	DeletedAt             gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
