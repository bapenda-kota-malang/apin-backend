package jpb3

import (
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb3 struct {
	opp.DetailLspop
	NoBangunan       int    `json:"noBangunan"`
	TipeKonstruksi   string `json:"tipeKonstruksi" gorm:"type:char(1)"`
	TinggiKolom3     int    `json:"tinggiKolom3"`
	LebarBentang3    int    `json:"lebarBentang3"`
	LuasMezzanine3   int    `json:"luasMezzanine3"`
	KelilingDinding3 int    `json:"kelilingDinding3"`
	DayaDukungLantai int    `json:"dayaDukungLantai"`
	gh.DateModel
}
