package espt

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthotel"
	"gorm.io/datatypes"
)

type FilterDto struct {
	// fixed
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

type CreateDto struct {
	// esptd
	Npwpd_Id        uint           `json:"npwp_id" validate:"required"`
	ObjekPajak_Id   uint           `json:"objekPajak_id"  validate:"required"`
	Rekening_Id     uint           `json:"rekening_id"  validate:"required"`
	Location        string         `json:"location"  validate:"required"`
	Description     string         `json:"description"  validate:"required"`
	PeriodeAwal     datatypes.Date `json:"periodeAwal"  validate:"required"`
	PeriodeAkhir    datatypes.Date `json:"periodeAkhir"  validate:"required"`
	TarifPajak_id   uint           `json:"tarifPajak_id"  validate:"required"`
	EtaxSubTotal    string         `json:"etaxSubTotal"  validate:"required"`
	EtaxTotal       string         `json:"etaxTotal"  validate:"required"`
	EtaxJumlahPajak string         `json:"etaxJumlahPajak"  validate:"required"`
	Omset           float64        `json:"omset"  validate:"required"`
	JumlahPajak     float32        `json:"jumlahPajak"  validate:"required"`
	Attachment      string         `json:"attachment"  validate:"required"`
	JatuhTempo      datatypes.Date `json:"jatuhTempo"  validate:"required"`
	Sunset          datatypes.Date `json:"sunset"  validate:"required"`
}

type UpdateDto struct {
	CreateDto
}

type CreateDetailAirDto struct {
	CreateDto
	DataDetails []detailesptair.CreateDto `json:"dataDetails" validate:"required"`
}

type CreateDetailHotelDto struct {
	CreateDto
	DataDetails []detailespthotel.CreateDto `json:"dataDetails" validate:"required"`
}
