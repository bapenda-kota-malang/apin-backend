package potensiopwp

import (
	"time"

	cm "github.com/bapenda-kota-malang/apin-backend/internal/models/configurationmodel"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"gorm.io/gorm"
)

type PotensiOp struct {
	Id            uint                 `json:"id" gorm:"primaryKey"`
	Golongan      rm.Golongan          `json:"golongan" validate:"required"`
	Rekening_Id   uint                 `json:"rekening_id" validate:"required"`
	Status        rm.StatusPendaftaran `json:"status" validate:"required"`
	ClosingDate   time.Time            `json:"closingDate"`
	OpeningDate   time.Time            `json:"openingDate"`
	User_Id       uint                 `json:"user_id" validate:"required"`
	LuasBangunan  string               `json:"luasBangunan" gorm:"size:50"`
	JamBuka       string               `json:"jamBuka" gorm:"size:50"`
	JamTutup      string               `json:"jamTutup" gorm:"size:50"`
	Visitors      string               `json:"visitors" gorm:"size:50" validate:"required"`
	OmsetOp       string               `json:"omsetOp" gorm:"size:50" validate:"required"`
	Genset        string               `json:"genset" gorm:"size:10" validate:"required"`
	AirTanah      string               `json:"airTanah" gorm:"size:10" validate:"required"`
	VendorEtax_Id uint                 `json:"vendorEtax_id"`
	CreatedAt     time.Time            `json:"createdAt"`
	UpdatedAt     time.Time            `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt       `json:"deletedAt" gorm:"index"`
	Rekening      cm.Rekening          `gorm:"foreignKey:Rekening_Id"`
	User          um.User              `gorm:"foreignKey:User_Id"`
	// VendorEtax    cm.VendorEtax        `gorm:"foreignKey:VendorEtax_Id"`
}

type Potensi struct {
	Id           uint                `json:"id" gorm:"primaryKey"`
	Potensiop_Id uint                `json:"Potensiop_id"`
	Nama         string              `json:"nama" gorm:"size:50"`
	Alamat       string              `json:"alamat" gorm:"size:50"`
	RtRw         string              `json:"rt_rw" gorm:"size:10"`
	Kecamatan_Id uint                `json:"kecamatan_id"`
	Kelurahan_Id uint                `json:"kelurahan_id"`
	Telp         string              `json:"telp" gorm:"size:20"`
	Status       rm.StatusObjekPajak `json:"status"`
	CreatedAt    time.Time           `json:"createdAt"`
	UpdatedAt    time.Time           `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt      `json:"deletedAt" gorm:"index"`
	Potensiop    PotensiOp           `gorm:"foreignKey:Potensiop_Id"`
	Kecamatan    cm.Kecamatan        `gorm:"foreignKey:Kecamatan_Id"`
	Kelurahan    cm.Kelurahan        `gorm:"foreignKey:Kelurahan_Id"`
}

type PotensiPemilikWp struct {
	Potensi
	NoIdPemilik string `json:"noIdPemilik" gorm:"size:20"`
	Nik         string `json:"nik" gorm:"size:20"`
}

type PotensiNarahubung struct {
	Potensi
	Nik string `json:"nik" gorm:"size:20"`
}
