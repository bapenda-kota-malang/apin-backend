package jpb8

import (
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb8 struct {
	opp.DetailLspop
	NoBangunan       int    `json:"noBangunan"`
	TipeKonstruksi   string `json:"tipeKonstruksi" gorm:"type:char(1)"`
	TinggiKolom8     int    `json:"tinggiKolom8"`
	LebarBentang8    int    `json:"lebarBentang8"`
	LuasMezzanine8   int    `json:"luasMezzanine8"`
	KelilingDinding8 int    `json:"kelilingDinding8"`
	DayaDukungLantai int    `json:"dayaDukungLantai"`
	gh.DateModel
}
