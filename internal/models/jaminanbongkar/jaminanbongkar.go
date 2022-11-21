package jaminanbongkar

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifjambong"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifjambongrek"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type JaminanBongkar struct {
	Id                 uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id             uuid.UUID       `json:"spt_Id" gorm:"type:uuid"`
	Nomor              string          `json:"nomor" gorm:"type:varchar(20)"`
	Tanggal            *datatypes.Date `json:"tanggal"`
	JenisReklame       JenisReklame    `json:"jenisReklame" gorm:"type:smallint"`
	TipeReklame        *uint8          `json:"tipeReklame" gorm:"type:smallint"`
	TarifJambong_Id    *uint64         `json:"tarifJambong_id"`
	TarifJambongRek_Id *uint64         `json:"tarifJambongRek_id"`
	Nominal            *float64        `json:"nominal" gorm:"type:decimal"`
	Batal              bool            `json:"batal"`
	TanggalBatas       *datatypes.Date `json:"TanggalBatas"`
	BiayaPemutusan     *float64        `json:"biayaPemutusan" gorm:"type:decimal"`
	Sisi               *uint64         `json:"sisi"`
	CreateBy_User_Id   uint            `json:"createBy_user_id"`
	gormhelper.DateModel
	TarifJambong    *tarifjambong.TarifJambong       `json:"tarifJambong,omitempty" gorm:"foreignKey:TarifJambong_Id"`
	TarifJambongRek *tarifjambongrek.TarifJambongRek `json:"tarifJambongRek,omitempty" gorm:"foreignKey:TarifJambongRek_Id"`
}

type CreateDto struct {
	Spt_Id             uuid.UUID       `json:"spt_Id"`
	Nomor              string          `json:"nomor"`
	Tanggal            *datatypes.Date `json:"tanggal"`
	JenisReklame       JenisReklame    `json:"jenisReklame"`
	TipeReklame        *uint8          `json:"tipeReklame"`
	TarifJambong_Id    *uint64         `json:"tarifJambong_Id"`
	TarifJambongRek_Id *uint64         `json:"tarifJambongRek_Id"`
	Nominal            *float64        `json:"nominal"`
	TanggalBatas       *datatypes.Date `json:"TanggalBatas"`
	BiayaPemutusan     *float64        `json:"biayaPemutusan"`
	Sisi               *uint64         `json:"sisi"`
}

type UpdateDto struct {
	Spt_Id             *uuid.UUID      `json:"spt_Id"`
	Nomor              string          `json:"nomor"`
	Tanggal            *datatypes.Date `json:"tanggal"`
	JenisReklame       *JenisReklame   `json:"jenisReklame"`
	TipeReklame        *uint8          `json:"tipeReklame"`
	TarifJambong_Id    *uint64         `json:"tarifJambong_Id"`
	TarifJambongRek_Id *uint64         `json:"tarifJambongRek_Id"`
	Nominal            *float64        `json:"nominal"`
	TanggalBatas       *datatypes.Date `json:"TanggalBatas"`
	BiayaPemutusan     *float64        `json:"biayaPemutusan"`
	Sisi               *uint64         `json:"sisi"`
}
