package jpb14

import (
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb14 struct {
	opp.DetailLspop
	NoBangunan int `json:"noBangunan"`
	LuasKanopi int `json:"luasKanopi"`
	gh.DateModel
}
