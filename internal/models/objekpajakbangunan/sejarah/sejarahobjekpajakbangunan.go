package sejarah

import (
	j "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	opbgn "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	p "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"time"
)

type SejarahObjekPajakBangunan struct {
	nop.NopDetail
	ObjekPajakBangunan_Id *uint64                 `json:"objekPajakBangunan_Id"`
	NoBangunan            *int                    `json:"noBangunan"`
	Jpb_Kode              string                  `json:"jpb_kode" gorm:"type:char(2)"`
	Jpb                   *j.Jpb                  `json:"jpb,omitempty" gorm:"foreignKey:Jpb_Kode;references:Kode"`
	NoFormulirSpop        *string                 `json:"noFormulirSpop" gorm:"type:char(11)"`
	TahunDibangun         *string                 `json:"tahunDibangun" gorm:"type:char(4)"`
	TahunRenovasi         *string                 `json:"tahunRenovasi" gorm:"type:char(4)"`
	LuasBangunan          *int                    `json:"luasBangunan"`
	Kondisi               opbgn.Kondisi           `json:"kondisi" gorm:"type:char(1)"`
	JenisKonstruksi       opbgn.JenisKonstruksi   `json:"jenisKonstruksi" gorm:"type:char(1)"`
	JenisAtap             opbgn.JenisAtap         `json:"jenisAtap" gorm:"type:char(1)"`
	KodeDinding           opbgn.JenisDinding      `json:"kodeDinding" gorm:"type:char(1)"`
	KodeLantai            opbgn.JenisLantai       `json:"kodeLantai" gorm:"type:char(1)"`
	KodeLangitLangit      opbgn.JenisLangitLangit `json:"kodeLangitLangit" gorm:"type:char(1)"`
	NilaiSistem           *int                    `json:"nilaiSistem"`
	JenisTransaksi        opbgn.JenisTransaksi    `json:"jenisTransaksi" gorm:"type:varchar(2)"`
	TanggalPendataan      *time.Time              `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string                 `json:"pendata_pegawai_nip" gorm:"type:char(9)"`
	Pendata_Pegawai       *p.Pegawai              `json:"pendata_pegawai,omitempty" gorm:"foreignKey:Pendata_Pegawai_Nip;references:Nip"`
	TanggalPemeriksaan    *time.Time              `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string                 `json:"pemeriksa_pegawai_nip" gorm:"type:char(9)"`
	Pemeriksa_Pegawai     *p.Pegawai              `json:"pemeriksa_pegawai,omitempty" gorm:"foreignKey:Pemeriksa_Pegawai_Nip;references:Nip"`
	TanggalPerekaman      *time.Time              `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string                 `json:"perekam_pegawai_nip" gorm:"type:char(9)"`
	Perekam_Pegawai       *p.Pegawai              `json:"perekam_pegawai,omitempty" gorm:"foreignKey:Perekam_Pegawai_Nip;references:Nip"`
	JmlLantaiBng          *int                    `json:"jmlLantaiBng"`
	gh.DateModel
}

type ListSejarahOpBangunan struct {
	Jpb_Kode         string     `json:"jpb_kode"`
	NoFormulirSpop   *string    `json:"noFormulirSpop"`
	LuasBangunan     *int       `json:"luasBangunan"`
	NilaiSistem      *int       `json:"nilaiSistem"`
	TanggalPerekaman *time.Time `json:"tanggalPerekaman"`
	NamaPerekam      *string    `json:"namaPerekam"`
}
