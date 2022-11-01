package spt

import (
	"time"

	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"gorm.io/datatypes"
)

type CreateDto struct {
	Npwpd_Id            uint64          `json:"npwpd_Id"`
	ObjekPajak_Id       uint64          `json:"objekPajak_Id"`
	Rekening_Id         *uint64         `json:"rekening_Id"`
	LuasLokasi          *string         `json:"luasLokasi"`
	Description         *string         `json:"description"`
	PeriodeAwal         *datatypes.Date `json:"periodeAwal,omitempty"`
	PeriodeAkhir        *datatypes.Date `json:"periodeAkhir,omitempty"`
	TarifPajak_Id       uint            `json:"tarifPajak_id"`
	EtaxSubTotal        *string         `json:"etaxSubTotal,omitempty"`
	EtaxTotal           *string         `json:"etaxTotal,omitempty"`
	EtaxJumlahPajak     *string         `json:"etaxJumlahPajak,omitempty"`
	Omset               float64         `json:"omset"`
	JumlahPajak         float32         `json:"jumlahPajak"`
	Lampiran            string          `json:"lampiran"`
	JatuhTempo          *datatypes.Date `json:"jatuhTempo"`
	Sunset              *datatypes.Date `json:"sunset,omitempty"`
	CreateBy_User_Id    uint            `json:"createBy_user_id"`
	TanggalSpt          *time.Time      `json:"tanggalSpt"`
	NomorSpt            *string         `json:"NomorSpt"`
	KodeBilling         *string         `json:"kodeBilling"`
	Type                *mt.JenisPajak  `json:"type"`
	Status              *SptStatus      `json:"status"`
	TanggalLunas        *time.Time      `json:"tanggalLunas"`
	BatalPenetapan      *string         `json:"batalPenetapan"`
	IdBT                *uint64         `json:"idBT"`
	JumlahTahun         *float64        `json:"jumlahTahun"`
	JumlahBulan         *float64        `json:"jumlahBulan"`
	JumlahMinggu        *float64        `json:"jumlahMinggu"`
	JumlahHari          *float64        `json:"jumlahHari"`
	Gambar              *string         `json:"gambar"`
	Keterangan          *string         `json:"keterangan"`
	KoefisienPajak      *uint64         `json:"koefisienPajak"`
	NamaProduk          *string         `json:"productName"`
	NomorRegister       *string         `json:"registerNumber"`
	VaJatim             *string         `json:"vaJatim"`
	JenisKetetapan      *string         `json:"jenisKetetapan"`
	Kenaikan            *float64        `json:"kenaikan"`
	Bunga               *float64        `json:"bunga"`
	Denda               *float64        `json:"denda"`
	Pengurangan         *float64        `json:"pengurangan"`
	Total               *float64        `json:"total"`
	Ref_Spt_Id          *uint64         `json:"ref_spt_id"`
	BillingPenetapan    *string         `json:"billingPenetapan"`
	Teguran_Id          *uint64         `json:"teguran_id"`
	IsTeguran           bool            `json:"isTeguran"`
	KeteranganPenetapan *string         `json:"keteranganPenetapan"`
	Kasubid_User_Id     *string         `json:"kasubid_user_id"`
	Kabid_User_Id       *string         `json:"kabid_user_id"`
}

type UpdateDto struct {
	Npwpd_Id            uint64          `json:"npwpd_Id"`
	ObjekPajak_Id       uint64          `json:"objekPajak_Id"`
	Rekening_Id         *uint64         `json:"rekening_Id"`
	LuasLokasi          *string         `json:"luasLokasi"`
	Description         *string         `json:"description"`
	PeriodeAwal         *datatypes.Date `json:"periodeAwal,omitempty"`
	PeriodeAkhir        *datatypes.Date `json:"periodeAkhir,omitempty"`
	TarifPajak_Id       uint            `json:"tarifPajak_id"`
	EtaxSubTotal        *string         `json:"etaxSubTotal,omitempty"`
	EtaxTotal           *string         `json:"etaxTotal,omitempty"`
	EtaxJumlahPajak     *string         `json:"etaxJumlahPajak,omitempty"`
	Omset               float64         `json:"omset"`
	JumlahPajak         float32         `json:"jumlahPajak"`
	Lampiran            string          `json:"lampiran"`
	JatuhTempo          *datatypes.Date `json:"jatuhTempo"`
	Sunset              *datatypes.Date `json:"sunset,omitempty"`
	CreateBy_User_Id    uint            `json:"createBy_user_id"`
	TanggalSpt          *time.Time      `json:"tanggalSpt"`
	NomorSpt            *string         `json:"NomorSpt"`
	KodeBilling         *string         `json:"kodeBilling"`
	Type                *mt.JenisPajak  `json:"type"`
	Status              *SptStatus      `json:"status"`
	TanggalLunas        *time.Time      `json:"tanggalLunas"`
	BatalPenetapan      *string         `json:"batalPenetapan"`
	IdBT                *uint64         `json:"idBT"`
	JumlahTahun         *float64        `json:"jumlahTahun"`
	JumlahBulan         *float64        `json:"jumlahBulan"`
	JumlahMinggu        *float64        `json:"jumlahMinggu"`
	JumlahHari          *float64        `json:"jumlahHari"`
	Gambar              *string         `json:"gambar"`
	Keterangan          *string         `json:"keterangan"`
	KoefisienPajak      *uint64         `json:"koefisienPajak"`
	NamaProduk          *string         `json:"productName"`
	NomorRegister       *string         `json:"registerNumber"`
	VaJatim             *string         `json:"vaJatim"`
	JenisKetetapan      *string         `json:"jenisKetetapan"`
	Kenaikan            *float64        `json:"kenaikan"`
	Bunga               *float64        `json:"bunga"`
	Denda               *float64        `json:"denda"`
	Pengurangan         *float64        `json:"pengurangan"`
	Total               *float64        `json:"total"`
	Ref_Spt_Id          *uint64         `json:"ref_spt_id"`
	BillingPenetapan    *string         `json:"billingPenetapan"`
	Teguran_Id          *uint64         `json:"teguran_id"`
	IsTeguran           bool            `json:"isTeguran"`
	KeteranganPenetapan *string         `json:"keteranganPenetapan"`
	Kasubid_User_Id     *string         `json:"kasubid_user_id"`
	Kabid_User_Id       *string         `json:"kabid_user_id"`
}

type FilterDto struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
