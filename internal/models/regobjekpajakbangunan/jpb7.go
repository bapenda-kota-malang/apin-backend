package regobjekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegJpb7 struct {
	nop.NopDetail
	NoBangunan             *int          `json:"noBangunan"`
	JenisHotel             JenisHotel    `json:"jenisHotel" gorm:"type:char(1)"`
	JumlahBintang          JumlahBintang `json:"jumlahBintang" gorm:"type:char(1)"`
	JumlahKamar            *int          `json:"jumlahKamar"`
	LuasKamarAcCentral     *int          `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
	gh.DateModel
}

type RegJpb7CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan             *int          `json:"noBangunan"`
	JenisHotel             JenisHotel    `json:"jenisHotel"`
	JumlahBintang          JumlahBintang `json:"jumlahBintang"`
	JumlahKamar            *int          `json:"jumlahKamar"`
	LuasKamarAcCentral     *int          `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
}

type RegJpb7UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan             *int          `json:"noBangunan"`
	JenisHotel             JenisHotel    `json:"jenisHotel"`
	JumlahBintang          JumlahBintang `json:"jumlahBintang"`
	JumlahKamar            *int          `json:"jumlahKamar"`
	LuasKamarAcCentral     *int          `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
}

type RegJpb7FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan             *int `json:"noBangunan"`
	JumlahKamar            *int `json:"jumlahKamar"`
	LuasKamarAcCentral     *int `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int `json:"luasRuangLainAcCentral"`
	Page                   int  `json:"page"`
	PageSize               int  `json:"page_size"`
}
