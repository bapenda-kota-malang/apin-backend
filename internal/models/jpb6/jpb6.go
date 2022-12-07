package jpb6

import (
	jt "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb2/types"
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb6 struct {
	opp.NopDetail
	NoBangunan     int              `json:"noBangunan"`
	KelasBangunan6 jt.KelasBangunan `json:"kelasBangunan6" gorm:"type:char(1)"`
	gh.DateModel
}
