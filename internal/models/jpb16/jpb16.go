package jpb16

import (
	jt "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb2/types"
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb16 struct {
	opp.DetailLspop
	NoBangunan      int              `json:"noBangunan"`
	KelasBangunan16 jt.KelasBangunan `json:"kelasBangunan16" gorm:"type:char(1)"`
	gh.DateModel
}
