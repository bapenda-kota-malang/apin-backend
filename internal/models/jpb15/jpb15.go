package jpb15

import (
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb15 struct {
	opp.NopDetail
	NoBangunan     int        `json:"noBangunan"`
	LetakTanki     LetakTanki `json:"letakTanki" gorm:"type:char(1)"`
	KapasitasTanki int        `json:"kapasitasTanki"`
	gh.DateModel
}
