package jpb13

import (
	jt "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb2/types"
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb13 struct {
	opp.NopDetail
	NoBangunan             int              `json:"noBangunan"`
	KelasBangunan13        jt.KelasBangunan `json:"kelasBangunan13" gorm:"type:char(1)"`
	JumlahApartment        int              `json:"jumlahApartment"`
	LuasApartAcCentral     int              `json:"luasApartAcCentral"`
	LuasRuangLainAcCentral int              `json:"luasRuangLainAcCentral"`
	gh.DateModel
}
