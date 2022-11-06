package npwpd

import (
	"time"

	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	op "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/skpd"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Npwpd struct {
	Id                uint64             `json:"id" gorm:"primarykey"`
	ObjekPajak_Id     uint64             `json:"objekPajak_id"`
	ObjekPajak        *op.ObjekPajak     `json:"objekPajak" gorm:"foreignKey:ObjekPajak_Id"`
	Golongan          t.Golongan         `json:"golongan"`
	Nomor             int                `json:"nomor"`
	TanggalPengukuhan *time.Time         `json:"tanggalPengukuhan"`
	TanggalNpwpd      *time.Time         `json:"tanggalNpwpd"`
	Npwpd             *string            `json:"npwpd" gorm:"size:22"`
	Status            t.Status           `json:"status"`
	TanggalTutup      *time.Time         `json:"tanggalTutup"`
	TanggalBuka       *time.Time         `json:"tanggalBuka"`
	JenisPajak        mt.JenisPajak      `json:"jenisPajak" gorm:"size:2"`
	Skpd_Id           *uint64            `json:"skpd_id"`
	Skpd              *skpd.Skpd         `json:"skpd,omitempty" gorm:"foreignKey:Skpd_Id"`
	Rekening_Id       *uint64            `json:"rekening_id"`
	Rekening          *rekening.Rekening `json:"rekening" gorm:"foreignKey:Rekening_Id"`
	Nama              *string            `json:"nama" gorm:"size:20"`
	User_Id           *uint64            `json:"user_id"`
	User              *user.User         `gorm:"foreignKey:User_Id"`
	OmsetOp           *string            `json:"omsetOp" gorm:"size:50"`
	Genset            *bool              `json:"genset"`
	AirTanah          *bool              `json:"airTanah"`
	TanggalMulaiUsaha *time.Time         `json:"tanggalMulaiUsaha"`
	LuasBangunan      *string            `json:"luasBangunan" gorm:"size:50"`
	JamBukaUsaha      *string            `json:"jamBukaUsaha" gorm:"size:50"`
	JamTutupUsaha     *string            `json:"jamTutupUsaha" gorm:"size:50"`
	Pengunjung        *string            `json:"pengunjung" gorm:"size:50"`
	FotoKtp           string             `json:"fotoKtp" gorm:"size:2048"`
	SuratIzinUsaha    string             `json:"suratIzinUsaha" gorm:"size:2048"`
	LainLain          string             `json:"lainLain" gorm:"size:2048"`
	FotoObjek         string             `json:"fotoObjek" gorm:"size:2048"`
	JalurRegistrasi   t.JalurRegistrasi  `json:"modeRegistrasi"`
	// VerifyStatus      *VerifyPendaftaran `json:"verifyStatus"`
	VerifiedAt    *time.Time `json:"verifiedAt"`
	VendorEtax_Id *string    `json:"vendorEtax_id"`
	// VendorEtax         *configurationmodel.VendorEtax `gorm:"foreignKey:VendorEtaxID"`

	PemilikWps  []*PemilikWp  `json:"pemilik,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	Narahubungs []*Narahubung `json:"narahubung,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`

	DetailOpAirTanah []*DetailObjekPajakAirTanah `json:"detailObjekPajakAirTanah,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	DetailOpHiburan  []*DetailObjekPajakHiburan  `json:"detailObjekPajakHiburan,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	DetailOpHotel    []*DetailObjekPajakHotel    `json:"detailObjekPajakHotel,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	DetailOpParkir   []*DetailObjekPajakParkir   `json:"detailObjekPajakParkir,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	DetailOpPpj      []*DetailObjekPajakPpj      `json:"detailObjekPajakPpj,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	DetailOpReklame  []*DetailObjekPajakReklame  `json:"detailObjekPajakReklame,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	DetailOpResto    []*DetailObjekPajakResto    `json:"detailObjekPajakResto,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	gormhelper.DateModel
}

type CreateDto struct {
	JenisPajak            mt.JenisPajak `json:"jenisPajak" validate:"required"`
	Golongan              t.Golongan    `json:"golongan" validate:"required"`
	Nomor                 int           `json:"nomor"`
	IsNomorRegistrasiAuto bool          `json:"isNomorRegistrasiAuto"`
	Npwpd                 *string       `json:"npwpd"`
	TanggalPengukuhan     *string       `json:"tanggalPengukuhan"`
	TanggalNpwpd          *string       `json:"tanggalNpwpd"`
	TanggalMulaiUsaha     *string       `json:"tanggalMulaiUsaha"`
	LuasBangunan          *string       `json:"luasBangunan"`
	JamBukaUsaha          *string       `json:"jamBukaUsaha"`
	JamTutupUsaha         *string       `json:"jamTutupUsaha"`
	Pengunjung            *string       `json:"pengunjung"`
	OmsetOp               *string       `json:"omsetOp"`
	Rekening_Id           uint64        `json:"rekening_id" validate:"required"`
	Genset                bool          `json:"genset" validate:"required"`
	AirTanah              bool          `json:"airTanah" validate:"required"`

	DetailObjekPajak *[]DetailObjekPajakCreateDto `json:"detailObjekPajak"`
	ObjekPajak       *op.ObjekPajakCreateDto      `json:"objekPajak"`
	Pemilik          *[]PemilikWpCreateDto        `json:"pemilik"`
	Narahubung       *[]NarahubungCreateDto       `json:"narahubung"`
}

type UpdateDto struct {
	// JenisPajak t.JenisPajak `json:"jenisPajak"`
	// Golongan   t.Golongan   `json:"golongan"`
	// Npwp       *string      `json:"npwp"`

	// NomorRegistrasi       *string `json:"nomorRegistrasi"`
	// IsNomorRegistrasiAuto bool    `json:"isNomorRegistrasiAuto"`

	// Npwpd             *string `json:"npwpd"`
	TanggalPengukuhan *string `json:"tanggalPengukuhan"`
	TanggalNpwpd      *string `json:"tanggalNpwpd"`

	TanggalMulaiUsaha *string `json:"tanggalMulaiUsaha"`
	LuasBangunan      *string `json:"luasBangunan"`
	JamBukaUsaha      *string `json:"jamBukaUsaha"`
	JamTutupUsaha     *string `json:"jamTutupUsaha"`
	Pengunjung        *string `json:"pengunjung"`
	OmsetOp           *string `json:"omsetOp"`

	Rekening_Id uint64 `json:"rekening_id"`

	Genset   bool `json:"genset"`
	AirTanah bool `json:"airTanah"`

	DetailObjekPajak []DetailObjekPajakUpdateDto `json:"detailObjekPajak"`

	ObjekPajak op.ObjekPajakUpdateDto `json:"objekPajak"`
	Pemilik    []PemilikWpUpdateDto   `json:"pemilik"`
	Narahubung []NarahubungUpdateDto  `json:"narahubung"`
}

type FilterDto struct {
	User_Id           *uint64        `json:"user_id"`
	JenisPajak        *mt.JenisPajak `json:"jenisPajak"`
	Golongan          *t.Golongan    `json:"golongan"`
	Nomor             *int           `json:"nomor"`
	Npwpd             *string        `json:"npwpd"`
	TanggalPengukuhan *string        `json:"tanggalPengukuhan"`
	TanggalNpwpd      *string        `json:"tanggalNpwpd"`
	TanggalMulaiUsaha *string        `json:"tanggalMulaiUsaha"`
	LuasBangunan      *string        `json:"luasBangunan"`
	JamBukaUsaha      *string        `json:"jamBukaUsaha"`
	JamTutupUsaha     *string        `json:"jamTutupUsaha"`
	Pengunjung        *string        `json:"pengunjung"`
	OmsetOp           *string        `json:"omsetOp"`
	Rekening_Id       *uint64        `json:"rekening_id"`
	Genset            *bool          `json:"genset"`
	AirTanah          *bool          `json:"airTanah"`
	// SuratIzinUsaha    *[]string     `json:"suratIzinUsaha"`
	// LainLain          *[]string     `json:"lainLain"`
	// FotoObjek         *[]string     `json:"fotoObjek"`
	//fixed
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PhotoUpdate struct {
	FotoKtp        string   `json:"fotoKtp"`
	SuratIzinUsaha []string `json:"suratIzinUsaha"`
	LainLain       []string `json:"lainLain"`
	FotoObjek      []string `json:"fotoObjek"`
}
