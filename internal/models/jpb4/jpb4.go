package jpb4

import (
	jt "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb2/types"
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb4 struct {
	opp.DetailLspop
	NoBangunan     int              `json:"noBangunan"`
	KelasBangunan4 jt.KelasBangunan `json:"kelasBangunan4" gorm:"type:char(1)"`
	gh.DateModel
}
