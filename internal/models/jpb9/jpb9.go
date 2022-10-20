package jpb9

import (
	jt "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb2/types"
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb9 struct {
	opp.DetailLspop
	NoBangunan     int              `json:"noBangunan"`
	KelasBangunan9 jt.KelasBangunan `json:"kelasBangunan9" gorm:"type:char(1)"`
	gh.DateModel
}
