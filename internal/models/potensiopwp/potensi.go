package potensiopwp

import (
	"time"

	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Status int16

const (
	StatusBaru Status = 0 //baru
	StatusLama Status = 1 //lama
)

type PotensiOp struct {
	Id          uint       `json:"id" gorm:"primaryKey"`
	Golongan    t.Golongan `json:"golongan"`
	Rekening_Id uint       `json:"rekening_id"`
	// Rekening      cm.Rekening          `gorm:"foreignKey:Rekening_Id"`
	Status      t.Status   `json:"status"`
	ClosingDate *time.Time `json:"closingDate"`
	OpeningDate *time.Time `json:"openingDate"`
	User_Id     uint       `json:"user_id"`
	// User          um.User              `gorm:"foreignKey:User_Id"`
	LuasBangunan  *string `json:"luasBangunan" gorm:"size:50"`
	JamBuka       *string `json:"jamBuka" gorm:"size:50"`
	JamTutup      *string `json:"jamTutup" gorm:"size:50"`
	Visitors      string  `json:"visitors" gorm:"size:50"`
	OmsetOp       string  `json:"omsetOp" gorm:"size:50"`
	Genset        string  `json:"genset" gorm:"size:10"`
	AirTanah      string  `json:"airTanah" gorm:"size:10"`
	VendorEtax_Id *uint   `json:"vendorEtax_id"`
	// VendorEtax    cm.VendorEtax        `gorm:"foreignKey:VendorEtax_Id"`
	PotensiPemilikWp       *PotensiPemilikWp       `gorm:"foreignKey:Potensiop_Id"`
	PotensiNarahubung      *PotensiNarahubung      `gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiOp        *DetailPotensiOp        `gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiAirTanahs []DetailPotensiAirTanah `gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiHiburans  []DetailPotensiHiburan  `gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiHotels    []DetailPotensiHotel    `gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiPPJs      []DetailPotensiPPJ      `gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiParkirs   []DetailPotensiParkir   `gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiReklames  []DetailPotensiReklame  `gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiRestos    []DetailPotensiResto    `gorm:"foreignKey:Potensiop_Id"`
	gormhelper.DateModel
}

type Potensi struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	Potensiop_Id uint   `json:"Potensiop_id"`
	Nama         string `json:"nama" gorm:"size:50"`
	Alamat       string `json:"alamat" gorm:"size:50"`
	RtRw         string `json:"rtRw" gorm:"size:10"`
	Kecamatan_Id uint   `json:"kecamatan_id"`
	// Kecamatan    adm.Kecamatan       `gorm:"foreignKey:Kecamatan_Id"`
	Kelurahan_Id uint `json:"kelurahan_id"`
	// Kelurahan    adm.Kelurahan       `gorm:"foreignKey:Kelurahan_Id"`
	Telp   string `json:"telp" gorm:"size:20"`
	Status Status `json:"status"`
	gormhelper.DateModel
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
