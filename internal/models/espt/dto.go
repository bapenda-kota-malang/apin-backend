package espt

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type FilterDto struct {
	Npwpd_Id     *uint64         `json:"npwpd_Id"`
	LuasLokasi   *uint           `json:"luasLokasi"`
	Omset        *float64        `json:"omset"`
	JumlahPajak  *float64        `json:"jumlahPajak"`
	JatuhTempo   *datatypes.Date `json:"jatuhTempo"`
	PeriodeAwal  *datatypes.Date `json:"periodeAwal"`
	PeriodeAkhir *datatypes.Date `json:"periodeAkhir"`
	VerifyStatus *string         `json:"verifyStatus"`
	CreatedAt    *time.Time      `json:"createdAt"`
	// fixed
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

type CreateDto struct {
	// esptd
	Npwpd_Id      uint    `json:"npwpd_id" validate:"required;min=1"`
	ObjekPajak_Id uint    `json:"objekPajak_id" validate:"required;min=1"`
	Rekening_Id   uint    `json:"rekening_id" validate:"required;min=1"`
	LuasLokasi    *uint   `json:"luasLokasi" validate:"min=1"`
	TarifPajak_Id uint    `json:"-"`
	Omset         float64 `json:"omset"`
	JumlahPajak   float64 `json:"-"`
	Attachment    string  `json:"attachment" validate:"base64=pdf,image,excel;b64size=1024"`
}

type UpdateDto struct {
	Npwpd_Id      *uint    `json:"npwpd_id" validate:"min=1"`
	ObjekPajak_Id *uint    `json:"objekPajak_id" validate:"min=1"`
	Rekening_Id   *uint    `json:"rekening_id" validate:"min=1"`
	LuasLokasi    *uint    `json:"luasLokasi" validate:"min=1"`
	Omset         *float64 `json:"omset"`
	JumlahPajak   *float64 `json:"-"`
	Attachment    *string  `json:"attachment" validate:"base64=pdf,image,excel;b64size=1024"`
}

type VerifyDto struct {
	Description  *string         `json:"description"`
	PeriodeAwal  *datatypes.Date `json:"periodeAwal"`
	PeriodeAkhir *datatypes.Date `json:"periodeAkhir"`
	JatuhTempo   *datatypes.Date `json:"jatuhTempo"`
	VerifyStatus string          `json:"verifyStatus" validate:"required"`
}

type CreateDetailBaseDto struct {
	Espt CreateDto `json:"espt" validate:"required"`
}

func (input *CreateDetailBaseDto) GetEspt() interface{} {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input *CreateDetailBaseDto) CalculateTax(taxPercentage *float64) {
	input.Espt.JumlahPajak = input.Espt.Omset * *taxPercentage / 100
}

func (input *CreateDetailBaseDto) ReplaceTarifPajakId(id uint) {
	input.Espt.TarifPajak_Id = id
}

func (input *CreateDetailBaseDto) GetDetails() interface{} {
	return nil
}

func (input *CreateDetailBaseDto) LenDetails() int {
	return 0
}

func (input *CreateDetailBaseDto) ChangeDetails(newData interface{}) {
	// do nothing
}

func (input *CreateDetailBaseDto) ReplaceEsptId(id uuid.UUID) {
	// do nothing
}

type UpdateDetailBaseDto struct {
	Espt UpdateDto `json:"espt" validate:"required"`
}

func (input *UpdateDetailBaseDto) GetEspt() interface{} {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input *UpdateDetailBaseDto) CalculateTax(taxPercentage *float64) {
	calc := *input.Espt.Omset * *taxPercentage / 100
	input.Espt.JumlahPajak = &calc
}

func (input *UpdateDetailBaseDto) ReplaceTarifPajakId(id uint) {
	// do nothing
}

func (input *UpdateDetailBaseDto) GetDetails() interface{} {
	return nil
}

func (input *UpdateDetailBaseDto) LenDetails() int {
	return 0
}

func (input *UpdateDetailBaseDto) ChangeDetails(newData interface{}) {
	// do nothing
}

func (input *UpdateDetailBaseDto) ReplaceEsptId(id uuid.UUID) {
	// do nothing
}
