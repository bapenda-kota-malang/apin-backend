package fasilitasbangunan

import (
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type FasilitasBangunan struct {
	opp.DetailLspop
	NoBangunan    int    `json:"noBangunan"`
	KodeFasilitas string `json:"kodeFasilitas" gorm:"type:char(2)"`
	JumlahSatuan  int    `json:"jumlahSatuan"`
	gh.DateModel
}
