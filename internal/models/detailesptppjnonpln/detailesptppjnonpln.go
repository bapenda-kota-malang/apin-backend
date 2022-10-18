package detailesptppjnonpln

import "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"

type DetailEsptPpjNonPln struct {
	detailesptair.DetailEspt
	JenisMesinPenggerak string  `json:"jenisMesinPenggerak" gorm:""`
	TahunMesin          string  `json:"tahunMesin"`
	DayaMesin           string  `json:"dayaMesin"`
	BebanMesin          *string `json:"bebanMesin"`
	JumlahJam           *uint8  `json:"jumlahJam"`
	JumlahHari          *uint8  `json:"jumlahHari"`
	ListrikPLN          bool    `json:"listrikPLN"`
}
