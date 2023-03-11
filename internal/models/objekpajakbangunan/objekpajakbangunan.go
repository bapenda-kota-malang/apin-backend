package objekpajakbangunan

import (
	"time"

	mf "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	j "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	p "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	mopbng "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

// lspop
type ObjekPajakBangunan struct {
	nop.NopDetail
	NoBangunan            *int              `json:"noBangunan"`
	Jpb_Kode              string            `json:"jpb_kode" gorm:"type:char(2)"`
	Jpb                   *j.Jpb            `json:"jpb,omitempty" gorm:"foreignKey:Jpb;references:Kode"`
	NoFormulirLspop       *string           `json:"noFormulirLspop" gorm:"type:char(11)"` // prev: NoFormulirSpop, and some others below
	TahunDibangun         *string           `json:"tahunDibangun" gorm:"type:char(4)"`
	TahunRenovasi         *string           `json:"tahunRenovasi" gorm:"type:char(4)"`
	LuasBangunan          *int              `json:"luasBangunan"`
	JmlLantaiBangunan     *int              `json:"jmlLantaiBangunan"` // prev: jmlLantaiBng, and some others below
	Kondisi               Kondisi           `json:"kondisi" gorm:"type:char(1)"`
	JenisKonstruksi       JenisKonstruksi   `json:"konstruksi" gorm:"type:char(1)"`
	JenisAtap             JenisAtap         `json:"jenisAtap" gorm:"type:char(1)"`
	JenisDinding          JenisDinding      `json:"jenisDinding" gorm:"type:char(1)"` // prev: KodeDinding, and some others below
	JenisLantai           JenisLantai       `json:"jenisLantai" gorm:"type:char(1)"`  // prev: KodeLantai, and some others below
	JenisLangit           JenisLangitLangit `json:"jenisLangit" gorm:"type:char(1)"`  // prev: KodeLangitLangit, and some others below
	NilaiSistem           *int              `json:"nilaiSistem"`
	TanggalPendataan      *time.Time        `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string           `json:"pendata_pegawai_nip" gorm:"type:char(9)"`
	Pendata_Pegawai       *p.Pegawai        `json:"pendata_pegawai,omitempty" gorm:"foreignKey:Pendata_Pegawai_Nip;references:Nip"`
	TanggalPemeriksaan    *time.Time        `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string           `json:"pemeriksa_pegawai_nip" gorm:"type:char(9)"`
	Pemeriksa_Pegawai     *p.Pegawai        `json:"pemeriksa_pegawai,omitempty" gorm:"foreignKey:Pemeriksa_Pegawai_Nip;references:Nip"`
	TanggalPerekaman      *time.Time        `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string           `json:"perekam_pegawai_nip" gorm:"type:char(9)"`
	Perekam_Pegawai       *p.Pegawai        `json:"perekam_pegawai,omitempty" gorm:"foreignKey:Perekam_Pegawai_Nip;references:Nip"`
	JenisTransaksi        JenisTransaksi    `json:"jenisTransaksi" gorm:"type:varchar(2)"`
	gh.DateModel
}

type CreateDto struct {
	nop.NopDetailCreateRequiredDto
	NoBangunan            *int                           `json:"noBangunan" validate:"required"`
	Jpb_Kode              string                         `json:"jpb_kode" validate:"required"`
	NoFormulirLspop       *string                        `json:"noFormulirSpop" validate:"required"`
	TahunDibangun         *string                        `json:"tahunDibangun" validate:"required"`
	TahunRenovasi         *string                        `json:"tahunRenovasi"`
	LuasBangunan          *int                           `json:"luasBangunan" validate:"required"`
	JmlLantaiBangunan     *int                           `json:"jmlLantaiBangunan" validate:"required"`
	Kondisi               Kondisi                        `json:"kondisi" validate:"required"`
	JenisKonstruksi       JenisKonstruksi                `json:"jenisKonstruksi"`
	JenisAtap             JenisAtap                      `json:"jenisAtap"`
	JenisDinding          JenisDinding                   `json:"jenisDinding"`
	JenisLantai           JenisLantai                    `json:"jenisLantai"`
	JenisLangit           JenisLangitLangit              `json:"jenisLangit"`
	NilaiSistem           *int                           `json:"nilaiSistem"`
	TanggalPerekaman      *string                        `json:"tanggalPerekaman" validate:"required"`
	Perekam_Pegawai_Nip   *string                        `json:"perekam_pegawai_nip"`
	TanggalPendataan      *string                        `json:"tanggalPendataan" validate:"required"`
	Pendata_Pegawai_Nip   *string                        `json:"pendata_pegawai_nip" validate:"required"`
	TanggalPemeriksaan    *string                        `json:"tanggalPemeriksaan" validate:"required"`
	Pemeriksa_Pegawai_Nip *string                        `json:"pemeriksa_pegawai_nip" validate:"required"`
	JenisTransaksi        JenisTransaksi                 `json:"jenisTransaksi" validate:"required"`
	Nop                   *string                        `json:"nop"`
	FasilitasBangunans    *mf.CreateDto                  `json:"fasilitasBangunan"`
	RegFasBangunan        *mopbng.OPBngFasilitasBangunan `json:"regFasBangunan"`
}

type UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan            *int              `json:"noBangunan" validate:"required"`
	Jpb_Kode              string            `json:"jpb_kode" validate:"required"`
	NoFormulirLspop       *string           `json:"noFormulirSpop" validate:"required"`
	TahunDibangun         *string           `json:"tahunDibangun" validate:"required"`
	TahunRenovasi         *string           `json:"tahunRenovasi"`
	LuasBangunan          *int              `json:"luasBangunan" validate:"required"`
	JmlLantaiBangunan     *int              `json:"jmlLantaiBangunan" validate:"required"`
	Kondisi               Kondisi           `json:"kondisi" validate:"required"`
	JenisKonstruksi       JenisKonstruksi   `json:"jenisKonstruksi"`
	JenisAtap             JenisAtap         `json:"jenisAtap"`
	JenisDinding          JenisDinding      `json:"jenisDinding"`
	JenisLantai           JenisLantai       `json:"jenisLantai"`
	JenisLangit           JenisLangitLangit `json:"jenisLangit"`
	NilaiSistem           *int              `json:"nilaiSistem"`
	TanggalPerekaman      *string           `json:"tanggalPerekaman" validate:"required"`
	Perekam_Pegawai_Nip   *string           `json:"perekam_pegawai_nip"`
	TanggalPendataan      *string           `json:"tanggalPendataan" validate:"required"`
	Pendata_Pegawai_Nip   *string           `json:"pendata_pegawai_nip" validate:"required"`
	TanggalPemeriksaan    *string           `json:"tanggalPemeriksaan" validate:"required"`
	Pemeriksa_Pegawai_Nip *string           `json:"pemeriksa_pegawai_nip" validate:"required"`
	JenisTransaksi        JenisTransaksi    `json:"jenisTransaksi" validate:"required"`
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
