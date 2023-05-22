package npwpd

import (
	"time"

	mnt "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	op "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	skpd "github.com/bapenda-kota-malang/apin-backend/internal/models/satuankerja"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Npwpd struct {
	Id                uint64              `json:"id" gorm:"primarykey"`
	ObjekPajak_Id     uint64              `json:"objekPajak_id"`
	ObjekPajak        *op.ObjekPajak      `json:"objekPajak" gorm:"foreignKey:ObjekPajak_Id"`
	Golongan          t.Golongan          `json:"golongan"`
	Nomor             int                 `json:"nomor"`
	TanggalPengukuhan *time.Time          `json:"tanggalPengukuhan"`
	TanggalNpwpd      *time.Time          `json:"tanggalNpwpd"`
	Npwpd             *string             `json:"npwpd" gorm:"unique;size:22;not null"`
	Status            t.Status            `json:"status"`
	TanggalTutup      *time.Time          `json:"tanggalTutup"`
	TanggalBuka       *time.Time          `json:"tanggalBuka"`
	JenisPajak        mt.JenisPajak       `json:"jenisPajak" gorm:"size:2"`
	SatuanKerja_Id    *uint64             `json:"skpd_id"`
	SatuanKerja       *skpd.SatuanKerja   `json:"skpd,omitempty" gorm:"foreignKey:SatuanKerja_Id"`
	Rekening_Id       *uint64             `json:"rekening_id"`
	Rekening          *rekening.Rekening  `json:"rekening" gorm:"foreignKey:Rekening_Id"`
	Nama              *string             `json:"nama" gorm:"size:20;not null"`
	User_Id           *uint64             `json:"user_id"`
	User              *user.User          `gorm:"foreignKey:User_Id"`
	OmsetOp           *string             `json:"omsetOp" gorm:"size:50"`
	Genset            *bool               `json:"genset"`
	AirTanah          *bool               `json:"airTanah"`
	TanggalMulaiUsaha *time.Time          `json:"tanggalMulaiUsaha"`
	LuasBangunan      *string             `json:"luasBangunan" gorm:"size:50"`
	JamBukaUsaha      *string             `json:"jamBukaUsaha" gorm:"size:50"`
	JamTutupUsaha     *string             `json:"jamTutupUsaha" gorm:"size:50"`
	Pengunjung        *string             `json:"pengunjung" gorm:"size:50"`
	FotoKtp           string              `json:"fotoKtp" gorm:"size:2048"`
	SuratIzinUsaha    string              `json:"suratIzinUsaha" gorm:"size:2048"`
	LainLain          string              `json:"lainLain" gorm:"size:2048"`
	FotoObjek         string              `json:"fotoObjek" gorm:"size:2048"`
	JalurRegistrasi   mnt.JalurRegistrasi `json:"modeRegistrasi"`
	VerifiedAt        *time.Time          `json:"verifiedAt"`
	VendorEtax_Id     *string             `json:"vendorEtax_id"`

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
	JenisPajak            *mt.JenisPajak `json:"jenisPajak" validate:"required"`
	Golongan              *t.Golongan    `json:"golongan" validate:"required"`
	Nomor                 int            `json:"nomor"`
	IsNomorRegistrasiAuto bool           `json:"isNomorRegistrasiAuto"`
	Npwpd                 *string        `json:"npwpd"`
	Rekening_Id           *uint64        `json:"rekening_id" validate:"required"`
	TanggalPengukuhan     *string        `json:"tanggalPengukuhan"`
	TanggalNpwpd          *string        `json:"tanggalNpwpd"`
	TanggalMulaiUsaha     *string        `json:"tanggalMulaiUsaha"`
	LuasBangunan          *string        `json:"luasBangunan"`
	JamBukaUsaha          *string        `json:"jamBukaUsaha"`
	JamTutupUsaha         *string        `json:"jamTutupUsaha"`
	Pengunjung            *string        `json:"pengunjung"`
	OmsetOp               *string        `json:"omsetOp"`
	Genset                *bool          `json:"genset" validate:"required"`
	AirTanah              *bool          `json:"airTanah" validate:"required"`

	ObjekPajak       op.ObjekPajakCreateDto      `json:"objekPajak" validate:"required"`
	DetailObjekPajak []DetailObjekPajakCreateDto `json:"detailObjekPajak"`
	Pemilik          []PemilikWpCreateDto        `json:"pemilik" validate:"required"`
	Narahubung       []NarahubungCreateDto       `json:"narahubung" validate:"required"`
}

type UpdateDto struct {
	// JenisPajak t.JenisPajak `json:"jenisPajak"`
	// Golongan   t.Golongan   `json:"golongan"`
	// Npwp       *string      `json:"npwp"`

	// NomorRegistrasi       *string `json:"nomorRegistrasi"`
	// IsNomorRegistrasiAuto bool    `json:"isNomorRegistrasiAuto"`

	// Npwpd             *string `json:"npwpd"`
	Rekening_Id       *uint64 `json:"rekening_id"`
	TanggalPengukuhan *string `json:"tanggalPengukuhan"`
	TanggalNpwpd      *string `json:"tanggalNpwpd"`
	TanggalMulaiUsaha *string `json:"tanggalMulaiUsaha"`
	LuasBangunan      *string `json:"luasBangunan"`
	JamBukaUsaha      *string `json:"jamBukaUsaha"`
	JamTutupUsaha     *string `json:"jamTutupUsaha"`
	Pengunjung        *string `json:"pengunjung"`
	OmsetOp           *string `json:"omsetOp"`
	Genset            *bool   `json:"genset"`
	AirTanah          *bool   `json:"airTanah"`

	ObjekPajak       op.ObjekPajakUpdateDto      `json:"objekPajak"`
	DetailObjekPajak []DetailObjekPajakUpdateDto `json:"detailObjekPajak"`
	Pemilik          []PemilikWpUpdateDto        `json:"pemilik"`
	Narahubung       []NarahubungUpdateDto       `json:"narahubung"`
}

type FilterDto struct {
	User_Id             *uint64   `json:"user_id"`
	JenisPajak          *string   `json:"jenisPajak"`
	Golongan            *int16    `json:"golongan"`
	Nomor               *int      `json:"nomor"`
	ObjekPajak_Nama     *string   `json:"objekPajak_nama" refsource:"ObjekPajak.Nama" reffunc:"LOWER"`
	ObjekPajak_Nama_Opt *string   `json:"objekPajak_nama_opt"`
	Npwpd               *string   `json:"npwpd"`
	TanggalPengukuhan   *string   `json:"tanggalPengukuhan"`
	TanggalNpwpd        *string   `json:"tanggalNpwpd"`
	TanggalMulaiUsaha   *string   `json:"tanggalMulaiUsaha"`
	LuasBangunan        *string   `json:"luasBangunan"`
	JamBukaUsaha        *string   `json:"jamBukaUsaha"`
	JamTutupUsaha       *string   `json:"jamTutupUsaha"`
	Pengunjung          *string   `json:"pengunjung"`
	OmsetOp             *string   `json:"omsetOp"`
	Rekening_Id         *uint64   `json:"rekening_id"`
	Rekening_Objek      *string   `json:"rekening_objek" refsource:"Rekening.Objek"`
	Rekening_ObjekS     *[]string `json:"rekening_objeks" refsource:"Rekening.Objek"`
	Genset              *bool     `json:"genset"`
	AirTanah            *bool     `json:"airTanah"`
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
