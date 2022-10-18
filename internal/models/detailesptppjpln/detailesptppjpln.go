package detailesptppjpln

import "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"

type DetailEsptPpjNonPln struct {
	detailesptair.DetailEspt
	JumlahPelanggan string  `json:"jenisMesinPenggerak" gorm:""`
	JumlahRekening  string  `json:"tahunMesin"`
	DayaMesin       string  `json:"dayaMesin"`
	BebanMesin      *string `json:"bebanMesin"`
	JumlahJam       *uint8  `json:"jumlahJam"`
	JumlahHari      *uint8  `json:"jumlahHari"`
	ListrikPLN      bool    `json:"listrikPLN"`
}
