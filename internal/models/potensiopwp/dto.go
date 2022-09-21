package potensiopwp

import (
	"time"

	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
)

type CreatePotensiOp struct {
	Golongan      rm.Golongan `json:"golongan" validate:"required"`
	Rekening_Id   uint        `json:"rekening_id" validate:"required,gt=0"`
	ClosingDate   *time.Time  `json:"closingDate"`
	OpeningDate   *time.Time  `json:"openingDate"`
	User_Id       uint        `json:"user_id" validate:"required,gt=0"`
	LuasBangunan  *string     `json:"luasBangunan"`
	JamBuka       *string     `json:"jamBuka"`
	JamTutup      *string     `json:"jamTutup"`
	Visitors      string      `json:"visitors" validate:"required"`
	OmsetOp       string      `json:"omsetOp" validate:"required"`
	Genset        string      `json:"genset" validate:"required"`
	AirTanah      string      `json:"airTanah" validate:"required"`
	VendorEtax_Id *uint       `json:"vendorEtax_id"`
}

type FilterDto struct {
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}
