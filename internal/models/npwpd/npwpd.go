package npwpd

import (
	"time"

	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/skpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Npwpd struct {
	Id                uint64             `json:"id" gorm:"primarykey"`
	Golongan          t.Golongan         `json:"golongan"`
	Nomor             *uint64            `json:"nomor"`
	Npwp              *string            `json:"npwp" gorm:"size:50"`
	TanggalPengukuhan *time.Time         `json:"tanggalPengukuhan"`
	TanggalNpwpd      *time.Time         `json:"tanggalNpwpd"`
	Npwpd             *string            `json:"npwpd" gorm:"size:22"`
	Status            t.Status           `json:"status"`
	TanggalTutup      *time.Time         `json:"tanggalTutup"`
	TanggalBuka       *time.Time         `json:"tanggalBuka"`
	JenisPajak        t.JenisPajak       `json:"jenisPajak" gorm:"size:2"`
	Skpd_Id           *uint64            `json:"skpd_id"`
	Skpd              *skpd.Skpd         `json:"skpd,omitempty" gorm:"foreignKey:Skpd_Id"`
	Rekening_Id       *uint64            `json:"rekening_id"`
	Rekening          *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
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
	FotoKtp           *string            `json:"fotoKtp" gorm:"size:50"`
	SuratIzinUsaha    *string            `json:"suratIzinUsaha" gorm:"size:50"`
	LainLain          *string            `json:"lainLain" gorm:"size:50"`
	FotoObjek         *string            `json:"fotoObjek" gorm:"size:50"`
	JalurRegistrasi   t.JalurRegistrasi  `json:"modeRegistrasi"`
	// VerifyStatus      *VerifyPendaftaran `json:"verifyStatus"`
	VerifiedAt    *time.Time `json:"verifiedAt"`
	VendorEtax_Id *string    `json:"vendorEtax_id"`
	// VendorEtax         *configurationmodel.VendorEtax `gorm:"foreignKey:VendorEtaxID"`

	ObjekPajaks []*ObjekPajak `json:"objekPajak,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	PemilikWps  []*PemilikWp  `json:"pemilik,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	Narahubungs []*Narahubung `json:"narahubung,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`

	DetailOpAirTanah []*DetailOpAirTanah `json:"detail_op_air_tanah,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	DetailOpHiburan  []*DetailOpHiburan  `json:"detail_op_hiburan,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	DetailOpHotel    []*DetailOpHotel    `json:"detail_op_hotel,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	DetailOpParkir   []*DetailOpParkir   `json:"detail_op_parkir,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	DetailOpPpj      []*DetailOpPpj      `json:"detail_op_ppj,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	DetailOpReklame  []*DetailOpReklame  `json:"detail_op_reklame,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`
	DetailOpResto    []*DetailOpResto    `json:"detail_op_resto,omitempty" gorm:"foreignKey:Npwpd_Id;references:Id"`

	gormhelper.DateModel
}

type CreateDto struct {
	JenisPajak t.JenisPajak `json:"jenisPajak" validate:"required"`
	Golongan   t.Golongan   `json:"golongan" validate:"required"`
	Npwp       *string      `json:"npwp" validate:"required"`

	Nomor                 *uint64 `json:"nomor"`
	IsNomorRegistrasiAuto bool    `json:"isNomorRegistrasiAuto"`

	Npwpd             *string `json:"npwpd"`
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

	DetailOp *[]DetailOp `json:"detail_op"`

	ObjekPajak *ObjekPajak   `json:"objekPajak"`
	Pemilik    *[]PemilikWp  `json:"pemilik"`
	Narahubung *[]Narahubung `json:"narahubung"`
}

type UpdateDto struct {
	JenisPajak t.JenisPajak `json:"jenisPajak"`
	Golongan   t.Golongan   `json:"golongan"`
	Npwp       *string      `json:"npwp"`

	NomorRegistrasi       *string `json:"nomorRegistrasi"`
	IsNomorRegistrasiAuto bool    `json:"isNomorRegistrasiAuto"`

	Npwpd             *string `json:"npwpd"`
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

	// DetailOp *[]DetailOp `json:"detail_op"`
	DetailOp *DetailOpUpdateDto `json:"detail_op"`

	ObjekPajak *[]ObjekPajak `json:"objekPajak"`
	Pemilik    *[]PemilikWp  `json:"pemilik"`
	// Narahubung *[]Narahubung `json:"narahubung"`
	Narahubung *NarahubungUpdateDto `json:"narahubung"`
}

type FilterDto struct {
	User_Id  *uint64 `json:"user_id"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}
