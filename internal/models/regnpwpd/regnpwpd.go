package regnpwpd

import (
	"time"

	rop "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/skpd"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegNpwpd struct {
	Id                uint64             `json:"id" gorm:"primarykey"`
	RegObjekPajak_Id  uint64             `json:"regObjekPajak_id"`
	RegObjekPajak     *rop.RegObjekPajak `json:"regObjekPajak,omitempty" gorm:"foreignKey:RegObjekPajak_Id"`
	Golongan          t.Golongan         `json:"golongan"`
	Nomor             int                `json:"nomor"`
	Status            t.Status           `json:"status"`
	TanggalPenutupan  *time.Time         `json:"tanggalPenutupan"`
	TanggalBuka       *time.Time         `json:"tanggalBuka"`
	JenisPajak        mt.JenisPajak      `json:"jenisPajak" gorm:"size:2"`
	Skpd_Id           *uint64            `json:"skpd_id"`
	Skpd              *skpd.Skpd         `json:"skpd,omitempty" gorm:"foreignKey:Skpd_Id"`
	Rekening_Id       *uint64            `json:"rekening_id"`
	Rekening          *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	User_Name         *string            `json:"user_name" gorm:"size:20"`
	User_Id           uint64             `json:"user_id"`
	User              *user.User         `json:"user" gorm:"foreignKey:User_Id;references:Id"`
	OmsetOp           *string            `json:"omsetOp" gorm:"size:50"`
	Genset            *bool              `json:"genset"`
	AirTanah          *bool              `json:"airTanah"`
	TanggalMulaiUsaha *time.Time         `json:"tanggalMulaiUsaha"`
	LuasBangunan      *string            `json:"luasBangunan" gorm:"size:50"`
	JamBukaUsaha      *string            `json:"jamBukaUsaha" gorm:"size:50"`
	JamTutupUsaha     *string            `json:"jamTutupUsaha" gorm:"size:50"`
	Pengunjung        *string            `json:"pengunjung" gorm:"size:50"`
	VendorEtax_Id     *string            `json:"vendorEtax_id"`
	gormhelper.DateModel
	VerifyStatus   VerifyStatus `json:"verifyStatus"`
	VerifiedAt     *time.Time   `json:"verifiedAt"`
	FotoKtp        string       `json:"fotoKtp" gorm:"size:2048"`
	SuratIzinUsaha string       `json:"suratIzinUsaha" gorm:"size:2048"`
	LainLain       string       `json:"lainLain" gorm:"size:2048"`
	FotoObjek      string       `json:"fotoObjek" gorm:"size:2048"`
	// ModeRegistrasi    npwpd.Mode               `json:"modeRegistrasi"`
	// VendorEtax         *configurationmodel.VendorEtax `gorm:"foreignKey:VendorEtaxID"`

	RegPemilikWps  []*RegPemilikWp  `json:"regPemilik,omitempty" gorm:"foreignKey:RegNpwpd_Id;references:Id"`
	RegNarahubungs []*RegNarahubung `json:"regNarahubung,omitempty" gorm:"foreignKey:RegNpwpd_Id;references:Id"`

	DetailRegOpAirTanah []*DetailRegObjekPajakAirTanah `json:"detailRegObjekPajakAirTanah,omitempty" gorm:"foreignKey:RegNpwpd_Id;references:Id"`
	DetailRegOpHiburan  []*DetailRegObjekPajakHiburan  `json:"detailRegObjekPajakHiburan,omitempty" gorm:"foreignKey:RegNpwpd_Id;references:Id"`
	DetailRegOpHotel    []*DetailRegObjekPajakHotel    `json:"detailRegObjekPajakHotel,omitempty" gorm:"foreignKey:RegNpwpd_Id;references:Id"`
	DetailRegOpParkir   []*DetailRegObjekPajakParkir   `json:"detailRegObjekPajakParkir,omitempty" gorm:"foreignKey:RegNpwpd_Id;references:Id"`
	DetailRegOpPpj      []*DetailRegObjekPajakPpj      `json:"detailRegObjekPajakPpj,omitempty" gorm:"foreignKey:RegNpwpd_Id;references:Id"`
	DetailRegOpReklame  []*DetailRegObjekPajakReklame  `json:"detailRegObjekPajakReklame,omitempty" gorm:"foreignKey:RegNpwpd_Id;references:Id"`
	DetailRegOpResto    []*DetailRegObjekPajakResto    `json:"detailRegObjekPajakResto,omitempty" gorm:"foreignKey:RegNpwpd_Id;references:Id"`
}

type CreateDto struct {
	Golongan          t.Golongan `json:"golongan" validate:"required"`
	TanggalPenutupan  *string    `json:"tanggalPenutupan"`
	TanggalBuka       *string    `json:"tanggalBuka"`
	Skpd_Id           *uint64    `json:"skpd_id"`
	Rekening_Id       *uint64    `json:"rekening_id" validate:"required"`
	User_Name         *string    `json:"user_name"`
	TanggalMulaiUsaha *string    `json:"tanggalMulaiUsaha"`
	LuasBangunan      *string    `json:"luasBangunan"`
	JamBukaUsaha      *string    `json:"jamBukaUsaha"`
	JamTutupUsaha     *string    `json:"jamTutupUsaha"`
	Pengunjung        *string    `json:"pengunjung"`
	OmsetOp           *string    `json:"omsetOp"`
	FotoKtp           string     `json:"fotoKtp" validate:"required;base64"`
	SuratIzinUsaha    []string   `json:"suratIzinUsaha" validate:"required"`
	LainLain          []string   `json:"lainLain"`
	FotoObjek         []string   `json:"fotoObjek" validate:"required"`

	Genset   bool `json:"genset" validate:"required"`
	AirTanah bool `json:"airTanah" validate:"required"`

	DetailRegOp *[]DetailRegObjekPajakCreateDto `json:"detailRegObjekPajak"`

	RegObjekPajak rop.RegObjekPajakCreateDto `json:"regObjekPajak" validate:"required"`
	RegPemilik    []RegPemilikWpCreateDto    `json:"regPemilik" validate:"required"`
	RegNarahubung []RegNarahubungCreateDto   `json:"regNarahubung" validate:"required"`
}

type UpdateDto struct {
	TanggalPenutupan *string `json:"tanggalPenutupan"`
	TanggalBuka      *string `json:"tanggalBuka"`
	Skpd_Id          *uint64 `json:"skpd_id"`
	Rekening_Id      *uint64 `json:"rekening_id"`
	User_Name        *string `json:"user_name"`

	TanggalMulaiUsaha *string `json:"tanggalMulaiUsaha"`
	LuasBangunan      *string `json:"luasBangunan"`
	JamBukaUsaha      *string `json:"jamBukaUsaha"`
	JamTutupUsaha     *string `json:"jamTutupUsaha"`
	Pengunjung        *string `json:"pengunjung"`
	OmsetOp           *string `json:"omsetOp"`

	Genset   bool `json:"genset"`
	AirTanah bool `json:"airTanah"`

	DetailRegObjekPajak []DetailRegObjekPajakUpdateDto `json:"detailRegObjekPajak"`

	RegObjekPajak rop.RegObjekPajakUpdateDto `json:"regObjekPajak"`
	RegPemilik    []RegPemilikWpUpdateDto    `json:"regPemilik"`
	RegNarahubung []RegNarahubungUpdateDto   `json:"regNarahubung"`
}

type VerifikasiDto struct {
	VerifyStatus int16 `json:"verifyStatus"`
}

type FilterDto struct {
	User_Id           *uint64     `json:"user_id"`
	Golongan          *t.Golongan `json:"golongan"`
	TanggalPenutupan  *string     `json:"tanggalPenutupan"`
	TanggalBuka       *string     `json:"tanggalBuka"`
	Skpd_Id           *uint64     `json:"skpd_id"`
	Rekening_Id       *uint64     `json:"rekening_id"`
	User_Name         *string     `json:"user_name"`
	TanggalMulaiUsaha *string     `json:"tanggalMulaiUsaha"`
	LuasBangunan      *string     `json:"luasBangunan"`
	JamBukaUsaha      *string     `json:"jamBukaUsaha"`
	JamTutupUsaha     *string     `json:"jamTutupUsaha"`
	Pengunjung        *string     `json:"pengunjung"`
	OmsetOp           *string     `json:"omsetOp"`
	FotoKtp           string      `json:"fotoKtp"`
	// SuratIzinUsaha    *[]string   `json:"suratIzinUsaha"`
	// LainLain          *[]string   `json:"lainLain"`
	// FotoObjek         *[]string   `json:"fotoObjek"`
	Genset   *bool `json:"genset"`
	AirTanah *bool `json:"airTanah"`
	//fixed
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
