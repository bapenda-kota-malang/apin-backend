package fasilitasbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type FasilitasBangunan struct {
	nop.NopDetail
	NoBangunan    int    `json:"noBangunan"`
	KodeFasilitas string `json:"kodeFasilitas" gorm:"type:char(2)"`
	JumlahSatuan  int    `json:"jumlahSatuan"`
	gh.DateModel
}
