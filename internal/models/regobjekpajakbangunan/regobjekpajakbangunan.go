package regobjekpajakbangunan

import (
	"time"

	mf "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	j "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	p "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

// lspop
type RegObjekPajakBangunan struct {
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
	Status                VerifyStatus      `json:"status"`
	TanggalPendataan      *time.Time        `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string           `json:"pendata_pegawai_nip" gorm:"type:char(9)"`
	Pendata_Pegawai       *p.Pegawai        `json:"pendata_pegawai,omitempty" gorm:"foreignKey:Pendata_Pegawai_Nip;references:Nip"`
	TanggalPemeriksaan    *time.Time        `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string           `json:"pemeriksa_pegawai_nip" gorm:"type:char(9)"`
	Pemeriksa_Pegawai     *p.Pegawai        `json:"pemeriksa_pegawai,omitempty" gorm:"foreignKey:Pemeriksa_Pegawai_Nip;references:Nip"`
	TanggalPerekaman      *time.Time        `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string           `json:"perekam_pegawai_nip" gorm:"type:char(9)"`
	Perekam_Pegawai       *p.Pegawai        `json:"perekam_pegawai,omitempty" gorm:"foreignKey:Perekam_Pegawai_Nip;references:Nip"`
	// jumlah lantai
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
	RegFasilitasBangunans *mf.CreateDto     `json:"fasilitasBangunan"`
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

type VerifyDto struct {
	Status VerifyStatus `json:"status" validate:"min=1;max=2"`
}

type OPBngFasilitasBangunan struct {
	FBJumlahACSplit        *int `json:"fbJumlahACSplit"`
	FBJumlahACWindow       *int `json:"fbJumlahACWindow"`
	FBIsACCentral          *int `json:"fbJumlahAC"`
	FBLuasKolamRenang      *int `json:"fbLuasKolamRenang"`
	FBTipeLapisanKolam     *int `json:"fbTipeLapisanKolam"`
	FBHalamanBerat         *int `json:"fbHalamanBerat"`
	FBHalamanSendang       *int `json:"fbHalamanSendang"`
	FBHalamanRingan        *int `json:"fbHalamanRingan"`
	FBHalamanLantai        *int `json:"fbHalamanLantai"`
	FBTenisLampuBeton      *int `json:"fbTenisLampuBeton"`
	FBTenisTanpaLampuBeton *int `json:"fbTenisTanpaLampuBeton"`
	FBTenisAspal1          *int `json:"fbTenisAspal1"`
	FBTenisAspal2          *int `json:"fbTenisAspal2"`
	FBTenisLiatRumput1     *int `json:"fbTenisLiatRumput1"`
	FBTenisLiatRumput2     *int `json:"fbTenisLiatRumput2"`
	FBLiftPenumpang        *int `json:"fbLiftPenumpang"`
	FBLiftKapsul           *int `json:"fbLiftKapsul"`
	FBLiftBarang           *int `json:"fbLiftBarang"`
	FBTangga80             *int `json:"fbTangga80"`
	FBTangga81             *int `json:"fbTangga81"`
	FBPagarPanjang         *int `json:"fbPagarPanjang"`
	FBPagarBahan           *int `json:"fbPagarBahan"`
	FBPKHydrant            *int `json:"fbPKHydrant"`
	FBPKSplinkler          *int `json:"fbPKSplinkler"`
	FBPKFireAI             *int `json:"fbPKFireAI"`
	FBPABX                 *int `json:"fbPABX"`
	FBSumur                *int `json:"fbSumur"`

	JpbKlinikACCentralKamar    *int    `json:"jpbKlinikACCentralKamar"`
	JpbKlinikACCentralRuang    *int    `json:"jpbKlinikACCentralRuang"`
	JpbHotelJenis              *int    `json:"jpbHotelJenis"`
	JpbHotelBintang            *int    `json:"jpbHotelBintang"`
	JpbHotelJmlKamar           *int    `json:"jpbHotelJmlKamar"`
	JpbHotelACCentralKamar     *int    `json:"jpbHotelACCentralKamar"`
	JpbHotelACCentralRuang     *int    `json:"jpbHotelACCentralRuang"`
	JpbApartemenJumlah         *int    `json:"jpbApartemenJumlah"`
	JpbApartemenACCentralKamar *int    `json:"jpbApartemenACCentralKamar"`
	JpbApartemenACCentralLain  *int    `json:"jpbApartemenACCentralLain"`
	JpbTankiKapasitas          *int    `json:"jpbTankiKapasitas"`
	JpbTankiLetak              *string `json:"jpbTankiLetak"`
	JpbProdTinggi              *int    `json:"jpbProdTinggi"`
	JpbProdLebar               *int    `json:"jpbProdLebar"`
	JpbProdDaya                *int    `json:"jpbProdDaya"`
	JpbProdKeliling            *int    `json:"jpbProdKeliling"`
	JpbProdLuas                *int    `json:"jpbProdLuas"`
}
