package objekpajakbangunan

import (
	"time"

	mf "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	j "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	p "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

// lspop
type ObjekPajakBangunan struct {
	nop.NopDetail
	NoBangunan            *int              `json:"noBangunan"`
	Jpb_Kode              string            `json:"jpb_kode" gorm:"type:char(2)"`
	Jpb                   *j.Jpb            `json:"jpb,omitempty" gorm:"foreignKey:Jpb_Kode;references:Kode"`
	NoFormulirSpop        *string           `json:"noFormulirSpop" gorm:"type:char(11)"`
	TahunDibangun         *string           `json:"tahunDibangun" gorm:"type:char(4)"`
	TahunRenovasi         *string           `json:"tahunRenovasi" gorm:"type:char(4)"`
	LuasBangunan          *int              `json:"luasBangunan"`
	Kondisi               Kondisi           `json:"kondisi" gorm:"type:char(1)"`
	JenisKonstruksi       JenisKonstruksi   `json:"jenisKonstruksi" gorm:"type:char(1)"`
	JenisAtap             JenisAtap         `json:"jenisAtap" gorm:"type:char(1)"`
	KodeDinding           JenisDinding      `json:"kodeDinding" gorm:"type:char(1)"`
	KodeLantai            JenisLantai       `json:"kodeLantai" gorm:"type:char(1)"`
	KodeLangitLangit      JenisLangitLangit `json:"kodeLangitLangit" gorm:"type:char(1)"`
	NilaiSistem           *int              `json:"nilaiSistem"`
	JenisTransaksi        JenisTransaksi    `json:"jenisTransaksi" gorm:"type:varchar(2)"`
	TanggalPendataan      *time.Time        `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string           `json:"pendata_pegawai_nip" gorm:"type:char(9)"`
	Pendata_Pegawai       *p.Pegawai        `json:"pendata_pegawai,omitempty" gorm:"foreignKey:Pendata_Pegawai_Nip;references:Nip"`
	TanggalPemeriksaan    *time.Time        `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string           `json:"pemeriksa_pegawai_nip" gorm:"type:char(9)"`
	Pemeriksa_Pegawai     *p.Pegawai        `json:"pemeriksa_pegawai,omitempty" gorm:"foreignKey:Pemeriksa_Pegawai_Nip;references:Nip"`
	TanggalPerekaman      *time.Time        `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string           `json:"perekam_pegawai_nip" gorm:"type:char(9)"`
	Perekam_Pegawai       *p.Pegawai        `json:"perekam_pegawai,omitempty" gorm:"foreignKey:Perekam_Pegawai_Nip;references:Nip"`
	JmlLantaiBng          *int              `json:"jmlLantaiBng"`
	gh.DateModel
}

type CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan            *int              `json:"noBangunan"`
	Jpb_Kode              string            `json:"jpb_kode"`
	NoFormulirSpop        *string           `json:"noFormulirSpop"`
	TahunDibangun         *string           `json:"tahunDibangun"`
	TahunRenovasi         *string           `json:"tahunRenovasi"`
	LuasBangunan          *int              `json:"luasBangunan"`
	Kondisi               Kondisi           `json:"kondisi"`
	JenisKonstruksi       JenisKonstruksi   `json:"jenisKonstruksi"`
	JenisAtap             JenisAtap         `json:"jenisAtap"`
	KodeDinding           JenisDinding      `json:"kodeDinding"`
	KodeLantai            JenisLantai       `json:"kodeLantai"`
	KodeLangitLangit      JenisLangitLangit `json:"kodeLangitLangit"`
	NilaiSistem           *int              `json:"nilaiSistem"`
	JenisTransaksi        JenisTransaksi    `json:"jenisTransaksi"`
	TanggalPendataan      *string           `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string           `json:"pendata_pegawai_nip"`
	TanggalPemeriksaan    *string           `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string           `json:"pemeriksa_pegawai_nip"`
	TanggalPerekaman      *string           `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string           `json:"perekam_pegawai_nip"`
	Nop                   *string           `json:"nop"`
	FasilitasBangunans    *mf.CreateDto     `json:"fasilitasBangunan"`
}

type UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan            *int              `json:"noBangunan"`
	Jpb_Kode              *string           `json:"jpb_kode"`
	NoFormulirSpop        *string           `json:"noFormulirSpop"`
	TahunDibangun         *string           `json:"tahunDibangun"`
	TahunRenovasi         *string           `json:"tahunRenovasi"`
	LuasBangunan          *int              `json:"luasBangunan"`
	Kondisi               Kondisi           `json:"kondisi"`
	JenisKonstruksi       JenisKonstruksi   `json:"jenisKonstruksi"`
	JenisAtap             JenisAtap         `json:"jenisAtap"`
	KodeDinding           JenisDinding      `json:"kodeDinding"`
	KodeLantai            JenisLantai       `json:"kodeLantai"`
	KodeLangitLangit      JenisLangitLangit `json:"kodeLangitLangit"`
	NilaiSistem           *int              `json:"nilaiSistem"`
	JenisTransaksi        JenisTransaksi    `json:"jenisTransaksi"`
	TanggalPendataan      *string           `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string           `json:"pendata_pegawai_nip"`
	TanggalPemeriksaan    *string           `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string           `json:"pemeriksa_pegawai_nip"`
	TanggalPerekaman      *string           `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string           `json:"perekam_pegawai_nip"`
}
type FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan            *int    `json:"noBangunan"`
	Jpb_Kode              *string `json:"jpb_kode"`
	NoFormulirSpop        *string `json:"noFormulirSpop"`
	TahunDibangun         *string `json:"tahunDibangun"`
	TahunRenovasi         *string `json:"tahunRenovasi"`
	LuasBangunan          *int    `json:"luasBangunan"`
	NilaiSistem           *int    `json:"nilaiSistem"`
	TanggalPendataan      *string `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string `json:"pendata_pegawai_nip"`
	TanggalPemeriksaan    *string `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string `json:"pemeriksa_pegawai_nip"`
	TanggalPerekaman      *string `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string `json:"perekam_pegawai_nip"`
	Page                  int     `json:"page"`
	PageSize              int     `json:"page_size"`
}
