package jaminanbongkarreklame

import (
	"time"

	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/types"
	mtj "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifjambong"
	mtjr "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifjambongrek"
)

type JaminanBongkarReklame struct {
	Id                    uint64               `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id                uint64               `json:"spt_id"`
	Nomor                 *string              `json:"nomor" gorm:"type:varchar(20)"`
	Date                  *time.Time           `json:"date"`
	JenisReklame          mt.JenisReklame      `json:"jenisReklame" gorm:"type:smallint"`
	TipeReklame           *uint64              `json:"tipeReklame" gorm:"type:smallint"`
	TarifJambong_Id       *uint64              `json:"tarifJambong_id"`
	TarifJambong          mtj.TarifJambong     `json:"tarifJambong,omitempty" gorm:"foreignKey:TarifJambong_Id"`
	TarifJambongRek_Id    *uint64              `json:"tarifJambongRek_id"`
	TarifJambongRek       mtjr.TarifJambongRek `json:"tarifJambongRek,omitempty" gorm:"foreignKey:TarifJambongRek_Id"`
	Nominal               *float64             `json:"nominal" gorm:"type:decimal"`
	Cancel                *uint64              `json:"cancel"`
	DueDate               *time.Time           `json:"dueDate"`
	BiayaPemutusanListrik *float64             `json:"biayaPemutusanListrik" gorm:"type:decimal"`
	Sisi                  *uint64              `json:"sisi"`
}

type CreateDto struct {
	Nomor                 string          `json:"nomor"`
	Date                  string          `json:"date"`
	JenisReklame          mt.JenisReklame `json:"jenisReklame"`
	TipeReklame           uint64          `json:"tipeReklame"`
	TarifJambong_Id       uint64          `json:"tarifJambong_id"`
	TarifJambongRek_Id    uint64          `json:"tarifJambongRek_id"`
	Nominal               float64         `json:"nominal"`
	Cancel                uint64          `json:"cancel"`
	DueDate               string          `json:"dueDate"`
	BiayaPemutusanListrik float64         `json:"biayaPemutusanListrik"`
	Sisi                  uint64          `json:"sisi"`
}

type UpdateDto struct {
	Nomor                 string          `json:"nomor"`
	Date                  string          `json:"date"`
	JenisReklame          mt.JenisReklame `json:"jenisReklame"`
	TipeReklame           uint64          `json:"tipeReklame"`
	TarifJambong_Id       uint64          `json:"tarifJambong_id"`
	TarifJambongRek_Id    uint64          `json:"tarifJambongRek_id"`
	Nominal               float64         `json:"nominal"`
	Cancel                uint64          `json:"cancel"`
	DueDate               string          `json:"dueDate"`
	BiayaPemutusanListrik float64         `json:"biayaPemutusanListrik"`
	Sisi                  uint64          `json:"sisi"`
}
