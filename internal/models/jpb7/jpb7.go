package jpb7

import (
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb7 struct {
	opp.NopDetail
	NoBangunan             int           `json:"noBangunan"`
	JenisHotel             JenisHotel    `json:"jenisHotel" gorm:"type:char(1)"`
	JumlahBintang          JumlahBintang `json:"jumlahBintang" gorm:"type:char(1)"`
	JumlahKamar            int           `json:"jumlahKamar"`
	LuasKamarAcCentral     int           `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral int           `json:"luasRuangLainAcCentral"`
	gh.DateModel
}
