package potensiopwp

import (
	"time"

	cm "github.com/bapenda-kota-malang/apin-backend/internal/models/configurationmodel"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"gorm.io/gorm"
)

type PotensiOp struct {
	ID           uint                 `json:"id" gorm:"primaryKey"`
	Golongan     rm.Golongan          `json:"golongan"`
	RekeningID   uint                 `json:"rekening_id"`
	Status       rm.StatusPendaftaran `json:"status"`
	ClosingDate  time.Time            `json:"closingDate"`
	OpeningDate  time.Time            `json:"openingDate"`
	UserId       uint                 `json:"user_id"`
	LuasBangunan string               `json:"luasBangunan" gorm:"size:50"`
	JamBuka      string               `json:"jamBuka" gorm:"size:50"`
	JamTutup     string               `json:"jamTutup" gorm:"size:50"`
	Visitors     string               `json:"visitors" gorm:"size:50"`
	OmsetOp      string               `json:"omsetOp" gorm:"size:50"`
	Genset       string               `json:"genset" gorm:"size:10"`
	AirTanah     string               `json:"airTanah" gorm:"size:10"`
	VendorEtaxID uint                 `json:"vendorEtax_id"`
	CreatedAt    time.Time            `json:"createdAt"`
	UpdatedAt    time.Time            `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt       `json:"deletedAt" gorm:"index"`
	Rekening     cm.Rekening
	User         um.User
	// VendorEtax   cm.VendorEtax
}

type Potensi struct {
	ID          uint                `json:"id" gorm:"primaryKey"`
	PotensiopID uint                `json:"Potensiop_id"`
	Nama        string              `json:"nama" gorm:"size:50"`
	Alamat      string              `json:"alamat" gorm:"size:50"`
	RtRw        string              `json:"rt_rw" gorm:"size:10"`
	KecamatanID uint                `json:"kecamatan_id"`
	KelurahanID uint                `json:"kelurahan_id"`
	Telp        string              `json:"telp" gorm:"size:20"`
	Status      rm.StatusObjekPajak `json:"status"`
	CreatedAt   time.Time           `json:"createdAt"`
	UpdatedAt   time.Time           `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt      `json:"deletedAt" gorm:"index"`
	Potensiop   PotensiOp
	Kecamatan   cm.Kecamatan
	Kelurahan   cm.Kelurahan
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
