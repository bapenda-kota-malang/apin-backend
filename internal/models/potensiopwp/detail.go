package potensiopwp

import (
	"time"

	"gorm.io/gorm"
)

type DetailPotensiOp struct {
	Potensi
	IsNpwpd uint8 `json:"isNpwpd"`
}

type Detail struct {
	Id           uint           `json:"id" gorm:"primaryKey"`
	Potensiop_Id uint           `json:"potensiop_id"`
	JenisOp      string         `json:"jenisOp"`
	JumlahOp     string         `json:"jumlahOp"`
	TarifOp      string         `json:"tarifOp"`
	UnitOp       string         `json:"unitOp"`
	Notes        string         `json:"notes"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Potensiop    PotensiOp      `gorm:"foreignKey:Potensi_Id"`
}

type DetailPotensiAirTanah struct {
	Detail
}

type DetailPotensiHiburan struct {
	Detail
}

type DetailPotensiHotel struct {
	Detail
}

type DetailPotensiParkir struct {
	Detail
}

type DetailPotensiPPJ struct {
	Detail
}

type DetailPotensiReklame struct {
	Detail
}

type DetailPotensiResto struct {
	Detail
}
