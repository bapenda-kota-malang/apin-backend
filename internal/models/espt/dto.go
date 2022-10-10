package espt

import (
	"gorm.io/datatypes"
)

type FilterDto struct {
	// fixed
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

type CreateDto struct {
	// esptd
	Npwpd_Id        uint           `json:"npwpd_id" validate:"min=1"`
	ObjekPajak_Id   uint           `json:"objekPajak_id" validate:"min=1"`
	Rekening_Id     uint           `json:"rekening_id" validate:"min=1"`
	Location        string         `json:"location" validate:"required"`
	PeriodeAwal     datatypes.Date `json:"periodeAwal" validate:"required"`
	PeriodeAkhir    datatypes.Date `json:"periodeAkhir" validate:"required"`
	TarifPajak_id   uint           `json:"tarifPajak_id" validate:"min=1"`
	EtaxSubTotal    string         `json:"etaxSubTotal" validate:"required"`
	EtaxTotal       string         `json:"etaxTotal" validate:"required"`
	EtaxJumlahPajak string         `json:"etaxJumlahPajak" validate:"required"`
	Omset           float64        `json:"omset"`
	JumlahPajak     float32        `json:"jumlahPajak" validate:"min=1"`
	Attachment      string         `json:"attachment" validate:"required"`
	JatuhTempo      datatypes.Date `json:"jatuhTempo" validate:"required"`
	Sunset          datatypes.Date `json:"sunset" validate:"required"`
}

type UpdateDto struct {
	Npwpd_Id        uint           `json:"npwpd_id" validate:"min=1"`
	ObjekPajak_Id   uint           `json:"objekPajak_id" validate:"min=1"`
	Rekening_Id     uint           `json:"rekening_id" validate:"min=1"`
	Location        string         `json:"location" validate:"required"`
	TarifPajak_id   uint           `json:"tarifPajak_id" validate:"min=1"`
	EtaxSubTotal    string         `json:"etaxSubTotal" validate:"required"`
	EtaxTotal       string         `json:"etaxTotal" validate:"required"`
	EtaxJumlahPajak string         `json:"etaxJumlahPajak" validate:"required"`
	Omset           float64        `json:"omset" validate:"min=1"`
	JumlahPajak     float32        `json:"jumlahPajak" validate:"min=1"`
	Attachment      string         `json:"attachment" validate:"required"`
	Sunset          datatypes.Date `json:"sunset" validate:"required"`
}

type VerifyDto struct {
	Description  string         `json:"description"`
	PeriodeAwal  datatypes.Date `json:"periodeAwal"`
	PeriodeAkhir datatypes.Date `json:"periodeAkhir"`
	JatuhTempo   datatypes.Date `json:"jatuhTempo"`
	VerifyStatus string         `json:"verifyStatus" validate:"required"`
}
