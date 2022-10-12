package espt

import (
	"time"

	"gorm.io/datatypes"
)

type FilterDto struct {
	Omset        *float64        `json:"omset"`
	JumlahPajak  *float32        `json:"jumlahPajak"`
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
	Npwpd_Id      uint    `json:"npwpd_id" validate:"min=1"`
	ObjekPajak_Id uint    `json:"objekPajak_id" validate:"min=1"`
	Rekening_Id   uint    `json:"rekening_id" validate:"min=1"`
	Location      *string `json:"location"`
	TarifPajak_id uint    `json:"tarifPajak_id" validate:"min=1"`
	Omset         float64 `json:"omset"`
	JumlahPajak   float32 `json:"jumlahPajak" validate:"min=1"`
	Attachment    string  `json:"attachment" validate:"base64"`
}

type UpdateDto struct {
	Npwpd_Id      *uint    `json:"npwpd_id"`
	ObjekPajak_Id *uint    `json:"objekPajak_id"`
	Rekening_Id   *uint    `json:"rekening_id"`
	Location      *string  `json:"location"`
	TarifPajak_id *uint    `json:"tarifPajak_id"`
	Omset         *float64 `json:"omset"`
	JumlahPajak   *float32 `json:"jumlahPajak"`
	Attachment    *string  `json:"attachment" validate:"base64"`
}

type VerifyDto struct {
	Description  *string         `json:"description"`
	PeriodeAwal  *datatypes.Date `json:"periodeAwal"`
	PeriodeAkhir *datatypes.Date `json:"periodeAkhir"`
	JatuhTempo   *datatypes.Date `json:"jatuhTempo"`
	VerifyStatus string          `json:"verifyStatus" validate:"required"`
}
