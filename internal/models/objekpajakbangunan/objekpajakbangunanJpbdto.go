package objekpajakbangunan

import (
	mf "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
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
	FasilitasBangunans    *mf.CreateDto     `json:"fasilitasBangunan"`
	Jpb2s                 *Jpb2CreateDto    `json:"jpb2"`
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
	FasilitasBangunans    *mf.CreateDto     `json:"fasilitasBangunan"`
	Jpb3s                 *Jpb3CreateDto    `json:"jpb3"`
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
	FasilitasBangunans    *mf.CreateDto     `json:"fasilitasBangunan"`
	Jpb4s                 *Jpb4CreateDto    `json:"jpb4"`
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
	FasilitasBangunans    *mf.CreateDto     `json:"fasilitasBangunan"`
	Jpb5s                 *Jpb5CreateDto    `json:"jpb5"`
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
	FasilitasBangunans    *mf.CreateDto     `json:"fasilitasBangunan"`
	Jpb6s                 *Jpb6CreateDto    `json:"jpb6"`
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
	FasilitasBangunans    *mf.CreateDto     `json:"fasilitasBangunan"`
	Jpb7s                 *Jpb7CreateDto    `json:"jpb7"`
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
	FasilitasBangunans    *mf.CreateDto     `json:"fasilitasBangunan"`
	Jpb8s                 *Jpb8CreateDto    `json:"jpb8"`
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
	FasilitasBangunans    *mf.CreateDto     `json:"fasilitasBangunan"`
	Jpb9s                 *Jpb9CreateDto    `json:"jpb9"`
}

type OpbJpb12CreateDto struct {
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
	Jpb12s                *Jpb12CreateDto   `json:"jpb12"`
}

type OpbJpb13CreateDto struct {
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
	Jpb13s                *Jpb13CreateDto   `json:"jpb13"`
}

type OpbJpb14CreateDto struct {
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
	Jpb14s                *Jpb14CreateDto   `json:"jpb14"`
}

type OpbJpb15CreateDto struct {
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
	Jpb15s                *Jpb15CreateDto   `json:"jpb15"`
}

type OpbJpb16CreateDto struct {
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
	Jpb16s                *Jpb16CreateDto   `json:"jpb16"`
}
