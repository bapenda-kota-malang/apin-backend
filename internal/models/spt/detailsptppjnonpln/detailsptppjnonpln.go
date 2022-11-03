package detailsptppjnonpln

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailSptPpjNonPln struct {
	Id                  uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id              uuid.UUID `json:"spt_id" gorm:"type:uuid"`
	JenisMesinPenggerak string    `json:"jenisMesinPenggerak"`
	TahunMesin          string    `json:"tahunMesin"`
	DayaMesin           string    `json:"dayaMesin"`
	BebanMesin          *string   `json:"bebanMesin"`
	JumlahJam           *uint8    `json:"jumlahJam"`
	JumlahHari          *uint8    `json:"jumlahHari"`
	ListrikPLN          bool      `json:"listrikPln"`
	gormhelper.DateModel
}

type CreateDto struct {
	Spt_Id              uuid.UUID
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
