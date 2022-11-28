package spt

import (
	"time"

	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/datatypes"
)

type CreateDto struct {
	Npwpd_Id         *uint64         `json:"npwpd_id" validate:"required;min=1"`
	ObjekPajak_Id    *uint64         `json:"objekPajak_id" validate:"required;min=1"`
	Rekening_Id      *uint64         `json:"rekening_id" validate:"required;min=1"`
	LuasLokasi       *string         `json:"luasLokasi"`
	Description      *string         `json:"description"`
	Omset            float64         `json:"omset"`
	JumlahPajak      float64         `json:"-"`
	Lampiran         string          `json:"lampiran" validate:"required;base64=pdf,image,excel;b64size=1024"`
	JatuhTempo       *datatypes.Date `json:"jatuhTempo"`
	PeriodeAkhir     *datatypes.Date `json:"periodeAkhir"`
	PeriodeAwal      *datatypes.Date `json:"periodeAwal"`
	TarifPajak_Id    uint64          `json:"-"`
	CreateBy_User_Id uint            `json:"-"`
	Type             mt.JenisPajak   `json:"-"`
	IdBT             *uint64         `json:"idBT"`
	VaJatim          *string         `json:"-"`
	// for reklame
	JumlahTahun     *float64 `json:"jumlahTahun"`
	JumlahBulan     *float64 `json:"jumlahBulan"`
	JumlahMinggu    *float64 `json:"jumlahMinggu"`
	JumlahHari      *float64 `json:"jumlahHari"`
	Gambar          *string  `json:"gambar" validate:"base64=image;b64size=1024"`
	KeteranganPajak *string  `json:"keteranganPajak"`
	KoefisienPajak  *uint64  `json:"koefisienPajak"`
	NamaProduk      *string  `json:"productName"`
	NomorRegister   *string  `json:"registerNumber"`
	// skdpdkb
	JenisKetetapan   *JenisKetetapan `json:"jenisKetetapan"`
	DasarPengenaan   *string         `json:"dasarPengenaan"`
	Kenaikan         *float64        `json:"-"`
	Bunga            *float64        `json:"-"`
	Denda            *float64        `json:"denda"`
	Pengurangan      *float64        `json:"pengurangan"`
	Total            *float64        `json:"-"`
	Ref_Spt_Id       *uuid.UUID      `json:"-"`
	BillingPenetapan *string         `json:"-"`
	Teguran_Id       *uint64         `json:"teguran_id"`
	IsTeguran        bool            `json:"isTeguran"`
}

type UpdateDto struct {
	Npwpd_Id         uint64          `json:"-"`
	ObjekPajak_Id    uint64          `json:"objekPajak_id"`
	Rekening_Id      *uint64         `json:"rekening_id"`
	LuasLokasi       *string         `json:"luasLokasi"`
	Description      *string         `json:"description"`
	PeriodeAwal      *datatypes.Date `json:"-"`
	PeriodeAkhir     *datatypes.Date `json:"-"`
	TarifPajak_Id    *uint64         `json:"-"`
	Omset            *float64        `json:"omset"`
	JumlahPajak      *float64        `json:"-"`
	Lampiran         *string         `json:"lampiran" validate:"base64=pdf,image,excel;b64size=1024"`
	JatuhTempo       *datatypes.Date `json:"-"`
	Sunset           *datatypes.Date `json:"sunset,omitempty"`
	KodeBilling      *string         `json:"kodeBilling"`
	Status           *SptStatus      `json:"-"`
	TanggalLunas     *time.Time      `json:"tanggalLunas"`
	BatalPenetapan   *string         `json:"batalPenetapan"`
	IdBT             *uint64         `json:"idBT"`
	JumlahTahun      *float64        `json:"jumlahTahun"`
	JumlahBulan      *float64        `json:"jumlahBulan"`
	JumlahMinggu     *float64        `json:"jumlahMinggu"`
	JumlahHari       *float64        `json:"jumlahHari"`
	Gambar           *string         `json:"gambar"`
	KeteranganPajak  *string         `json:"keteranganPajak"`
	KoefisienPajak   *uint64         `json:"koefisienPajak"`
	NamaProduk       *string         `json:"productName"`
	NomorRegister    *string         `json:"registerNumber"`
	VaJatim          *string         `json:"vaJatim"`
	JenisKetetapan   *JenisKetetapan `json:"jenisKetetapan"`
	Kenaikan         *float64        `json:"kenaikan"`
	Bunga            *float64        `json:"bunga"`
	Denda            *float64        `json:"denda"`
	Pengurangan      *float64        `json:"pengurangan"`
	Total            *float64        `json:"total"`
	Ref_Spt_Id       *uint64         `json:"ref_spt_id"`
	BillingPenetapan *string         `json:"billingPenetapan"`
	Teguran_Id       *uint64         `json:"teguran_id"`
	IsTeguran        *bool           `json:"isTeguran"`
}

type VerifyDto struct {
	KeteranganPenetapan *string `json:"keteranganPenetapan"`
	StatusPenetapan     string  `json:"statusPenetapan" validate:"required"`
}

type SkpdkbExisting struct {
	Spt_Id           uuid.UUID      `json:"spt_id" validate:"required"`
	JenisKetetapan   JenisKetetapan `json:"jenisKetetapan"`
	DasarPengenaan   *string        `json:"dasarPengenaan" validate:"required"`
	Kenaikan         *float64       `json:"-"`
	Bunga            *float64       `json:"-"`
	Denda            *float64       `json:"denda"`
	Pengurangan      *float64       `json:"pengurangan"`
	Total            *float64       `json:"-"`
	Ref_Spt_Id       *uuid.UUID     `json:"-"`
	BillingPenetapan *string        `json:"-"`
	Teguran_Id       *uint64        `json:"teguran_id"`
	IsTeguran        bool           `json:"isTeguran"`
}

type FilterDto struct {
	Npwpd_Id        *uint64         `json:"npwpd_id"`
	StatusData      *uint8          `json:"statusData"`
	Type            mt.JenisPajak   `json:"-"`
	JatuhTempo      *datatypes.Date `json:"jatuhTempo"`
	JatuhTempo_Opt  *string         `json:"jatuhTempo_opt"`
	Kasubid_User_Id *string         `json:"kasubid_user_id"`
	Kabid_User_Id   *string         `json:"kabid_user_id"`
	Page            int             `json:"page"`
	PageSize        int             `json:"page_size"`
}

type ListDataDto struct {
	Spt
	NominalBayar *float64 `json:"-"`
	StatusFinal  *string  `json:"statusFinal"`
}

type CreateDetailBaseDto struct {
	Spt CreateDto `json:"spt" validate:"required"`
}

func (input *CreateDetailBaseDto) GetSpt(baseUri string) interface{} {
	typeSpt := input.Spt.Type
	if typeSpt == "" {
		typeSpt = mt.JenisPajakSA
	}
	if baseUri == "skpd" {
		typeSpt = mt.JenisPajakOA
		jenisKetetapan := JenisKetetapanSkpd
		input.Spt.JenisKetetapan = &jenisKetetapan
	}
	input.Spt.Type = typeSpt
	input.Spt.LuasLokasi = nil
	return input.Spt
}

func (input *CreateDetailBaseDto) ReplaceTarifPajakId(id uint) {
	input.Spt.TarifPajak_Id = uint64(id)
}

func (input *CreateDetailBaseDto) CalculateTax(taxPercentage *float64) {
	input.Spt.JumlahPajak = input.Spt.Omset * (*taxPercentage / 100)
}

func (input *CreateDetailBaseDto) GetDetails() interface{} {
	return nil
}

func (input *CreateDetailBaseDto) LenDetails() int {
	return 0
}

func (input *CreateDetailBaseDto) ReplaceSptId(id uuid.UUID) {
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

func (input *CreateDetailBaseDto) SkpdkbDuplicate(sptDetail *Spt, skpdkb *SkpdkbExisting) error {
	if err := copier.Copy(&input.Spt, sptDetail); err != nil {
		return err
	}
	if err := copier.Copy(&input.Spt, skpdkb); err != nil {
		return err
	}
	input.Spt.JatuhTempo = &datatypes.Date{}
	input.Spt.Ref_Spt_Id = &sptDetail.Id
	return nil
}

// calculate skpdkb process
//
// kenaikan = jumlah pajak * 0.25 (25%)
//
// bunga = jumlah pajak * 0.02 (2%)
//
// total = jumlah pajak + kenaikan + bunga + denda - pengurangan
func (input *CreateDetailBaseDto) CalculateSkpdkb() {
	kenaikan := input.Spt.JumlahPajak * 0.25
	bunga := input.Spt.JumlahPajak * 0.02
	denda := float64(0)
	pengurangan := float64(0)
	if input.Spt.Denda != nil {
		denda = *input.Spt.Denda
	}
	if input.Spt.Pengurangan != nil {
		pengurangan = *input.Spt.Pengurangan
	}
	total := input.Spt.JumlahPajak + kenaikan + bunga + denda - pengurangan
	input.Spt.Kenaikan = &kenaikan
	input.Spt.Bunga = &bunga
	input.Spt.Denda = &denda
	input.Spt.Pengurangan = &pengurangan
	input.Spt.Total = &total
}

type UpdateDetailBaseDto struct {
	Spt UpdateDto `json:"spt" validate:"required"`
}

func (input *UpdateDetailBaseDto) GetSpt(baseUri string) interface{} {
	// typeSpt := mt.JenisPajakSA
	// if baseUri == "skpd" {
	// 	typeSpt = mt.JenisPajakOA
	// }
	// input.Spt.Type = typeSpt
	input.Spt.LuasLokasi = nil
	return input.Spt
}

func (input *UpdateDetailBaseDto) ReplaceTarifPajakId(id uint) {
	// input.Spt.TarifPajak_Id = &id
}

func (input *UpdateDetailBaseDto) CalculateTax(taxPercentage *float64) {
	calc := *input.Spt.Omset * *taxPercentage / 100
	input.Spt.JumlahPajak = &calc
}

func (input *UpdateDetailBaseDto) GetDetails() interface{} {
	return nil
}

func (input *UpdateDetailBaseDto) LenDetails() int {
	return 0
}

func (input *UpdateDetailBaseDto) ReplaceSptId(id uuid.UUID) {
	// do nothing
}

func (input *UpdateDetailBaseDto) ChangeDetails(newData interface{}) {
	// do nothing
}

func (input *UpdateDetailBaseDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	// do nothing
	return nil
}

func (input *UpdateDetailBaseDto) SkpdkbDuplicate(sptDetail *Spt, skpdkb *SkpdkbExisting) error {
	// do nothiing
	return nil
}

func (input *UpdateDetailBaseDto) CalculateSkpdkb() {
	// do nothing
}
