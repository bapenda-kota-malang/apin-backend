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
	Npwpd_Id         *uint64         `json:"npwpd_Id" validate:"required"`
	ObjekPajak_Id    uint64          `json:"objekPajak_Id"`
	Rekening_Id      *uint64         `json:"rekening_Id" validate:"required"`
	LuasLokasi       *string         `json:"luasLokasi"`
	Description      *string         `json:"description"`
	Omset            float64         `json:"omset"`
	JumlahPajak      float32         `json:"-"`
	Lampiran         string          `json:"lampiran" validate:"required"`
	Sunset           *datatypes.Date `json:"sunset,omitempty"`
	JatuhTempo       datatypes.Date  `json:"-"`
	Type             mt.JenisPajak   `json:"-"`
	PeriodeAkhir     *datatypes.Date `json:"-"`
	TarifPajak_Id    uint64          `json:"-"`
	PeriodeAwal      *datatypes.Date `json:"-"`
	CreateBy_User_Id uint            `json:"-"`
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
	JenisKetetapan   *string         `json:"jenisKetetapan"`
	Kenaikan         *float64        `json:"kenaikan"`
	Pengurangan      *float64        `json:"pengurangan"`
	Total            *float64        `json:"total"`
	Ref_Spt_Id       *uint64         `json:"ref_spt_id"`
	BillingPenetapan *string         `json:"billingPenetapan"`
	Teguran_Id       *uint64         `json:"teguran_id"`
	IsTeguran        bool            `json:"isTeguran"`
}

type UpdateDto struct {
	Npwpd_Id         uint64          `json:"-"`
	ObjekPajak_Id    uint64          `json:"objekPajak_Id"`
	Rekening_Id      *uint64         `json:"rekening_Id"`
	LuasLokasi       *string         `json:"luasLokasi"`
	Description      *string         `json:"description"`
	PeriodeAwal      *datatypes.Date `json:"-"`
	PeriodeAkhir     *datatypes.Date `json:"-"`
	TarifPajak_Id    *uint64         `json:"-"`
	Omset            float64         `json:"omset"`
	JumlahPajak      float32         `json:"-"`
	Lampiran         *string         `json:"lampiran"`
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
	JenisKetetapan   *string         `json:"jenisKetetapan"`
	Kenaikan         *float64        `json:"kenaikan"`
	Bunga            *float64        `json:"bunga"`
	Denda            *float64        `json:"denda"`
	Pengurangan      *float64        `json:"pengurangan"`
	Total            *float64        `json:"total"`
	Ref_Spt_Id       *uint64         `json:"ref_spt_id"`
	BillingPenetapan *string         `json:"billingPenetapan"`
	Teguran_Id       *uint64         `json:"teguran_id"`
	IsTeguran        bool            `json:"isTeguran"`
}

type VerifyDto struct {
	KeteranganPenetapan *string         `json:"keteranganPenetapan"`
	StatusPenetapan     StatusPenetapan `json:"statusPenetapan"`
	Kasubid_User_Id     *string         `json:"kasubid_user_id"`
	Kabid_User_Id       *string         `json:"kabid_user_id"`
}

type FilterDto struct {
	Npwpd_Id        *uint64          `json:"npwpd_Id"`
	TbpStatus       *TbpStatusFilter `json:"tbpStatus"`
	Type            mt.JenisPajak    `json:"-"`
	JatuhTempo      *datatypes.Date  `json:"jatuhTempo"`
	Kasubid_User_Id *string          `json:"kasubid_user_id"`
	Kabid_User_Id   *string          `json:"kabid_user_id"`
	Page            int              `json:"page"`
	PageSize        int              `json:"page_size"`
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
	input.Spt.TarifPajak_Id = uint64(id)
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
	input.Spt.JumlahPajak = float32(input.Spt.Omset * (*taxPercentage / 100))
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
