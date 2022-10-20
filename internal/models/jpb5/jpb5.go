package jpb5

import (
	jt "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb2/types"
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb5 struct {
	opp.DetailLspop
	NoBangunan             int              `json:"noBangunan"`
	KelasBangunan5         jt.KelasBangunan `json:"kelasBangunan5" gorm:"type:char(1)"`
	LuasKamarAcCentral     int              `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral int              `json:"luasRuangLainAcCentral"`
	gh.DateModel
}
