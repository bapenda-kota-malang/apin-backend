package jpb2

import (
	jt "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb2/types"
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb2 struct {
	opp.NopDetail
	NoBangunan     int              `json:"noBangunan"`
	KelasBangunan2 jt.KelasBangunan `json:"kelasBangunan2" gorm:"type:char(1)"`
	gh.DateModel
}
