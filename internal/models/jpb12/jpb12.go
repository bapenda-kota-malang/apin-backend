package jpb12

import (
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb12 struct {
	opp.DetailLspop
	NoBangunan   int          `json:"noBangunan"`
	TipeBangunan TipeBangunan `json:"tipeBangunan" gorm:"type:char(1)"`
	gh.DateModel
}
