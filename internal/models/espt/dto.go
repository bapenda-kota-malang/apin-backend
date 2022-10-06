package espt

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthotel"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptparkir"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptreklame"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptresto"
	"gorm.io/datatypes"
)

type FilterDto struct {
	// fixed
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

type CreateDto struct {
	// esptd
	Npwpd_Id        uint           `json:"npwp_id" validate:"min=1"`
	ObjekPajak_Id   uint           `json:"objekPajak_id"  validate:"min=1"`
	Rekening_Id     uint           `json:"rekening_id"  validate:"min=1"`
	Location        string         `json:"location"  validate:"required"`
	Description     string         `json:"description"  validate:"required"`
	PeriodeAwal     datatypes.Date `json:"periodeAwal"  validate:"required"`
	PeriodeAkhir    datatypes.Date `json:"periodeAkhir"  validate:"required"`
	TarifPajak_id   uint           `json:"tarifPajak_id"  validate:"min=1"`
	EtaxSubTotal    string         `json:"etaxSubTotal"  validate:"required"`
	EtaxTotal       string         `json:"etaxTotal"  validate:"required"`
	EtaxJumlahPajak string         `json:"etaxJumlahPajak"  validate:"required"`
	Omset           float64        `json:"omset"  validate:"min=1"`
	JumlahPajak     float32        `json:"jumlahPajak"  validate:"min=1"`
	Attachment      string         `json:"attachment"  validate:"required"`
	JatuhTempo      datatypes.Date `json:"jatuhTempo"  validate:"required"`
	Sunset          datatypes.Date `json:"sunset"  validate:"required"`
}

type UpdateDto struct {
	CreateDto
}

type CreateDetailAirDto struct {
	Espt        CreateDto                 `json:"espt" validate:"required"`
	DataDetails []detailesptair.CreateDto `json:"dataDetails" validate:"required"`
}

type CreateDetailHotelDto struct {
	Espt        CreateDto                   `json:"espt" validate:"required"`
	DataDetails []detailespthotel.CreateDto `json:"dataDetails" validate:"required"`
}
type CreateDetailParkirDto struct {
	Espt        CreateDto                    `json:"espt" validate:"required"`
	DataDetails []detailesptparkir.CreateDto `json:"dataDetails" validate:"required"`
}

type CreateDetailReklameDto struct {
	Espt        CreateDto                     `json:"espt" validate:"required"`
	DataDetails []detailesptreklame.CreateDto `json:"dataDetails" validate:"required"`
}

type CreateDetailRestoDto struct {
	Espt        CreateDto                   `json:"espt" validate:"required"`
	DataDetails []detailesptresto.CreateDto `json:"dataDetails" validate:"required"`
}
