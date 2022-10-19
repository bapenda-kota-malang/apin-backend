package detailesptppjnonpln

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DetailEsptPpjNonPln struct {
	detailesptair.DetailEspt
	JenisMesinPenggerak string  `json:"jenisMesinPenggerak"`
	TahunMesin          string  `json:"tahunMesin"`
	DayaMesin           string  `json:"dayaMesin"`
	BebanMesin          *string `json:"bebanMesin"`
	JumlahJam           *uint8  `json:"jumlahJam"`
	JumlahHari          *uint8  `json:"jumlahHari"`
	ListrikPLN          bool    `json:"listrikPln"`
	gormhelper.DateModel
}

type CreateDto struct {
	Espt_Id             uint
	JenisMesinPenggerak string  `json:"jenisMesinPenggerak"`
	TahunMesin          string  `json:"tahunMesin"`
	DayaMesin           string  `json:"dayaMesin"`
	BebanMesin          *string `json:"bebanMesin"`
	JumlahJam           *uint8  `json:"jumlahJam"`
	JumlahHari          *uint8  `json:"jumlahHari"`
	ListrikPLN          bool    `json:"listrikPln"`
}

type UpdateDto struct {
	Id                  uint    `json:"id"`
	JenisMesinPenggerak *string `json:"jenisMesinPenggerak"`
	TahunMesin          *string `json:"tahunMesin"`
	DayaMesin           *string `json:"dayaMesin"`
	BebanMesin          *string `json:"bebanMesin"`
	JumlahJam           *uint8  `json:"jumlahJam"`
	JumlahHari          *uint8  `json:"jumlahHari"`
	ListrikPLN          *bool   `json:"listrikPln"`
}
