package spt

import (
	"time"

	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/jinzhu/copier"
	"gorm.io/datatypes"
)

type CreateDto struct {
	Npwpd_Id            *uint64         `json:"npwpd_Id" validate:"required"`
	ObjekPajak_Id       uint64          `json:"objekPajak_Id"`
	Rekening_Id         *uint64         `json:"rekening_Id"`
	LuasLokasi          *string         `json:"luasLokasi"`
	Description         *string         `json:"description"`
	TarifPajak_Id       uint            `json:"tarifPajak_id"`
	EtaxSubTotal        *string         `json:"etaxSubTotal,omitempty"`
	EtaxTotal           *string         `json:"etaxTotal,omitempty"`
	EtaxJumlahPajak     *string         `json:"etaxJumlahPajak,omitempty"`
	Omset               float64         `json:"omset"`
	JumlahPajak         float32         `json:"jumlahPajak"`
	Lampiran            string          `json:"lampiran" validate:"required"`
	Sunset              *datatypes.Date `json:"sunset,omitempty"`
	JatuhTempo          datatypes.Date
	Type                mt.JenisPajak
	PeriodeAkhir        *datatypes.Date
	PeriodeAwal         *datatypes.Date
	CreateBy_User_Id    uint
	IdBT                *uint64  `json:"idBT"`
	JumlahTahun         *float64 `json:"jumlahTahun"`
	JumlahBulan         *float64 `json:"jumlahBulan"`
	JumlahMinggu        *float64 `json:"jumlahMinggu"`
	JumlahHari          *float64 `json:"jumlahHari"`
	Gambar              *string  `json:"gambar"`
	Keterangan          *string  `json:"keterangan"`
	KoefisienPajak      *uint64  `json:"koefisienPajak"`
	NamaProduk          *string  `json:"productName"`
	NomorRegister       *string  `json:"registerNumber"`
	VaJatim             *string  `json:"vaJatim"`
	JenisKetetapan      *string  `json:"jenisKetetapan"`
	Kenaikan            *float64 `json:"kenaikan"`
	Pengurangan         *float64 `json:"pengurangan"`
	Total               *float64 `json:"total"`
	Ref_Spt_Id          *uint64  `json:"ref_spt_id"`
	BillingPenetapan    *string  `json:"billingPenetapan"`
	Teguran_Id          *uint64  `json:"teguran_id"`
	IsTeguran           bool     `json:"isTeguran"`
	KeteranganPenetapan *string  `json:"keteranganPenetapan"`
	Kasubid_User_Id     *string  `json:"kasubid_user_id"`
	Kabid_User_Id       *string  `json:"kabid_user_id"`
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
	Npwpd_Id uint64    `json:"npwpd_Id"`
	Status   SptStatus `json:"status"`
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
}

type CreateDetailBaseDto struct {
	Spt CreateDto `json:"spt" validate:"required"`
}

func (input *CreateDetailBaseDto) GetSpt(baseUri string) interface{} {
	typeSpt := mt.JenisPajakSA
	if baseUri == "skpd" {
		typeSpt = mt.JenisPajakOA
	}
	input.Spt.Type = typeSpt
	input.Spt.LuasLokasi = nil
	return input.Spt
}

func (input *CreateDetailBaseDto) ReplaceTarifPajakId(id uint) {
	input.Spt.TarifPajak_Id = id
}

func (input *CreateDetailBaseDto) CalculateTax(taxPercentage *float64) {
	input.Spt.JumlahPajak = float32(input.Spt.Omset * (*taxPercentage / 100))
}

func (input *CreateDetailBaseDto) GetDetails() interface{} {
	return nil
}

func (input *CreateDetailBaseDto) LenDetails() int {
	return 0
}

func (input *CreateDetailBaseDto) ReplaceSptId(id uint) {
	// do nothing
}

func (input *CreateDetailBaseDto) ChangeDetails(newData interface{}) {
	// do nothing
}

func (input *CreateDetailBaseDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	if err := copier.Copy(&input.Spt, &esptDetail); err != nil {
		return err
	}
	input.Spt.Lampiran = esptDetail.Attachment
	input.Spt.CreateBy_User_Id = *esptDetail.VerifyBy_User_Id
	return nil
}
