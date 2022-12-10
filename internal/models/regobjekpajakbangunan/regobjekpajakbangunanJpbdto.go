package regobjekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	mrf "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
)

type OpbJpb2CreateDto struct {
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
	RegFasilitasBangunans *mrf.CreateDto    `json:"regFasilitasBangunan"`
	RegJpbs               *RegJpb2CreateDto `json:"regJpb"`
}

type OpbJpb3CreateDto struct {
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
	RegFasilitasBangunans *mrf.CreateDto    `json:"regFasilitasBangunan"`
	RegJpbs               *RegJpb3CreateDto `json:"regJpb"`
}

type OpbJpb4CreateDto struct {
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
	RegFasilitasBangunans *mrf.CreateDto    `json:"regFasilitasBangunan"`
	RegJpbs               *RegJpb4CreateDto `json:"regJpb"`
}

type OpbJpb5CreateDto struct {
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
	RegFasilitasBangunans *mrf.CreateDto    `json:"regFasilitasBangunan"`
	RegJpbs               *RegJpb5CreateDto `json:"regJpb"`
}

type OpbJpb6CreateDto struct {
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
	RegFasilitasBangunans *mrf.CreateDto    `json:"regFasilitasBangunan"`
	RegJpbs               *RegJpb6CreateDto `json:"regJpb"`
}

type OpbJpb7CreateDto struct {
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
	RegFasilitasBangunans *mrf.CreateDto    `json:"regFasilitasBangunan"`
	RegJpbs               *RegJpb7CreateDto `json:"regJpb"`
}

type OpbJpb8CreateDto struct {
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
	RegFasilitasBangunans *mrf.CreateDto    `json:"regFasilitasBangunan"`
	RegJpbs               *RegJpb8CreateDto `json:"regJpb"`
}

type OpbJpb9CreateDto struct {
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
	RegFasilitasBangunans *mrf.CreateDto    `json:"regFasilitasBangunan"`
	RegJpbs               *RegJpb9CreateDto `json:"regJpb"`
}

type OpbJpb12CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan            *int               `json:"noBangunan"`
	Jpb_Kode              string             `json:"jpb_kode"`
	NoFormulirSpop        *string            `json:"noFormulirSpop"`
	TahunDibangun         *string            `json:"tahunDibangun"`
	TahunRenovasi         *string            `json:"tahunRenovasi"`
	LuasBangunan          *int               `json:"luasBangunan"`
	Kondisi               Kondisi            `json:"kondisi"`
	JenisKonstruksi       JenisKonstruksi    `json:"jenisKonstruksi"`
	JenisAtap             JenisAtap          `json:"jenisAtap"`
	KodeDinding           JenisDinding       `json:"kodeDinding"`
	KodeLantai            JenisLantai        `json:"kodeLantai"`
	KodeLangitLangit      JenisLangitLangit  `json:"kodeLangitLangit"`
	NilaiSistem           *int               `json:"nilaiSistem"`
	JenisTransaksi        JenisTransaksi     `json:"jenisTransaksi"`
	TanggalPendataan      *string            `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string            `json:"pendata_pegawai_nip"`
	TanggalPemeriksaan    *string            `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string            `json:"pemeriksa_pegawai_nip"`
	TanggalPerekaman      *string            `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string            `json:"perekam_pegawai_nip"`
	Nop                   *string            `json:"nop"`
	RegFasilitasBangunans *mrf.CreateDto     `json:"regFasilitasBangunan"`
	RegJpbs               *RegJpb12CreateDto `json:"regJpb"`
}

type OpbJpb13CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan            *int               `json:"noBangunan"`
	Jpb_Kode              string             `json:"jpb_kode"`
	NoFormulirSpop        *string            `json:"noFormulirSpop"`
	TahunDibangun         *string            `json:"tahunDibangun"`
	TahunRenovasi         *string            `json:"tahunRenovasi"`
	LuasBangunan          *int               `json:"luasBangunan"`
	Kondisi               Kondisi            `json:"kondisi"`
	JenisKonstruksi       JenisKonstruksi    `json:"jenisKonstruksi"`
	JenisAtap             JenisAtap          `json:"jenisAtap"`
	KodeDinding           JenisDinding       `json:"kodeDinding"`
	KodeLantai            JenisLantai        `json:"kodeLantai"`
	KodeLangitLangit      JenisLangitLangit  `json:"kodeLangitLangit"`
	NilaiSistem           *int               `json:"nilaiSistem"`
	JenisTransaksi        JenisTransaksi     `json:"jenisTransaksi"`
	TanggalPendataan      *string            `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string            `json:"pendata_pegawai_nip"`
	TanggalPemeriksaan    *string            `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string            `json:"pemeriksa_pegawai_nip"`
	TanggalPerekaman      *string            `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string            `json:"perekam_pegawai_nip"`
	Nop                   *string            `json:"nop"`
	RegFasilitasBangunans *mrf.CreateDto     `json:"regFasilitasBangunan"`
	RegJpbs               *RegJpb13CreateDto `json:"regJpb"`
}

type OpbJpb14CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan            *int               `json:"noBangunan"`
	Jpb_Kode              string             `json:"jpb_kode"`
	NoFormulirSpop        *string            `json:"noFormulirSpop"`
	TahunDibangun         *string            `json:"tahunDibangun"`
	TahunRenovasi         *string            `json:"tahunRenovasi"`
	LuasBangunan          *int               `json:"luasBangunan"`
	Kondisi               Kondisi            `json:"kondisi"`
	JenisKonstruksi       JenisKonstruksi    `json:"jenisKonstruksi"`
	JenisAtap             JenisAtap          `json:"jenisAtap"`
	KodeDinding           JenisDinding       `json:"kodeDinding"`
	KodeLantai            JenisLantai        `json:"kodeLantai"`
	KodeLangitLangit      JenisLangitLangit  `json:"kodeLangitLangit"`
	NilaiSistem           *int               `json:"nilaiSistem"`
	JenisTransaksi        JenisTransaksi     `json:"jenisTransaksi"`
	TanggalPendataan      *string            `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string            `json:"pendata_pegawai_nip"`
	TanggalPemeriksaan    *string            `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string            `json:"pemeriksa_pegawai_nip"`
	TanggalPerekaman      *string            `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string            `json:"perekam_pegawai_nip"`
	Nop                   *string            `json:"nop"`
	RegFasilitasBangunans *mrf.CreateDto     `json:"regFasilitasBangunan"`
	RegJpbs               *RegJpb14CreateDto `json:"regJpb"`
}

type OpbJpb15CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan            *int               `json:"noBangunan"`
	Jpb_Kode              string             `json:"jpb_kode"`
	NoFormulirSpop        *string            `json:"noFormulirSpop"`
	TahunDibangun         *string            `json:"tahunDibangun"`
	TahunRenovasi         *string            `json:"tahunRenovasi"`
	LuasBangunan          *int               `json:"luasBangunan"`
	Kondisi               Kondisi            `json:"kondisi"`
	JenisKonstruksi       JenisKonstruksi    `json:"jenisKonstruksi"`
	JenisAtap             JenisAtap          `json:"jenisAtap"`
	KodeDinding           JenisDinding       `json:"kodeDinding"`
	KodeLantai            JenisLantai        `json:"kodeLantai"`
	KodeLangitLangit      JenisLangitLangit  `json:"kodeLangitLangit"`
	NilaiSistem           *int               `json:"nilaiSistem"`
	JenisTransaksi        JenisTransaksi     `json:"jenisTransaksi"`
	TanggalPendataan      *string            `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string            `json:"pendata_pegawai_nip"`
	TanggalPemeriksaan    *string            `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string            `json:"pemeriksa_pegawai_nip"`
	TanggalPerekaman      *string            `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string            `json:"perekam_pegawai_nip"`
	Nop                   *string            `json:"nop"`
	RegFasilitasBangunans *mrf.CreateDto     `json:"regFasilitasBangunan"`
	RegJpbs               *RegJpb15CreateDto `json:"regJpb"`
}

type OpbJpb16CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan            *int               `json:"noBangunan"`
	Jpb_Kode              string             `json:"jpb_kode"`
	NoFormulirSpop        *string            `json:"noFormulirSpop"`
	TahunDibangun         *string            `json:"tahunDibangun"`
	TahunRenovasi         *string            `json:"tahunRenovasi"`
	LuasBangunan          *int               `json:"luasBangunan"`
	Kondisi               Kondisi            `json:"kondisi"`
	JenisKonstruksi       JenisKonstruksi    `json:"jenisKonstruksi"`
	JenisAtap             JenisAtap          `json:"jenisAtap"`
	KodeDinding           JenisDinding       `json:"kodeDinding"`
	KodeLantai            JenisLantai        `json:"kodeLantai"`
	KodeLangitLangit      JenisLangitLangit  `json:"kodeLangitLangit"`
	NilaiSistem           *int               `json:"nilaiSistem"`
	JenisTransaksi        JenisTransaksi     `json:"jenisTransaksi"`
	TanggalPendataan      *string            `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string            `json:"pendata_pegawai_nip"`
	TanggalPemeriksaan    *string            `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string            `json:"pemeriksa_pegawai_nip"`
	TanggalPerekaman      *string            `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string            `json:"perekam_pegawai_nip"`
	Nop                   *string            `json:"nop"`
	RegFasilitasBangunans *mrf.CreateDto     `json:"regFasilitasBangunan"`
	RegJpbs               *RegJpb16CreateDto `json:"regJpb"`
}
