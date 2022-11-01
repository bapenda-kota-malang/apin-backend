package registrasinpwpd

import (
	"time"

	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	rop "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/skpd"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegistrasiNpwpd struct {
	Id                uint64             `json:"id" gorm:"primarykey"`
	RegObjekPajak_Id  uint64             `json:"regObjekPajak_id"`
	RegObjekPajak     *rop.RegObjekPajak `json:"regObjekPajak,omitempty" gorm:"foreignKey:RegObjekPajak_Id"`
	Golongan          t.Golongan         `json:"golongan"`
	Nomor             int                `json:"nomor"`
	Npwp              *string            `json:"npwp" gorm:"size:50"`
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
	FotoKtp        string       `json:"fotoKtp" gorm:"size:50"`
	SuratIzinUsaha string       `json:"suratIzinUsaha" gorm:"size:2048"`
	LainLain       string       `json:"lainLain" gorm:"size:2048"`
	FotoObjek      string       `json:"fotoObjek" gorm:"size:2048"`
	// ModeRegistrasi    npwpd.Mode               `json:"modeRegistrasi"`
	// VendorEtax         *configurationmodel.VendorEtax `gorm:"foreignKey:VendorEtaxID"`

	RegPemilikWps  []*RegPemilikWp  `json:"regPemilik,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	RegNarahubungs []*RegNarahubung `json:"regNarahubung,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`

	DetailRegOpAirTanah []*DetailRegObjekPajakAirTanah `json:"detail_reg_op_air_tanah,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	DetailRegOpHiburan  []*DetailRegObjekPajakHiburan  `json:"detail_reg_op_hiburan,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	DetailRegOpHotel    []*DetailRegObjekPajakHotel    `json:"detail_reg_op_hotel,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	DetailRegOpParkir   []*DetailRegObjekPajakParkir   `json:"detail_reg_op_parkir,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	DetailRegOpPpj      []*DetailRegObjekPajakPpj      `json:"detail_reg_op_ppj,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	DetailRegOpReklame  []*DetailRegObjekPajakReklame  `json:"detail_reg_op_reklame,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	DetailRegOpResto    []*DetailRegObjekPajakResto    `json:"detail_reg_op_resto,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
}

type CreateDto struct {
	// JenisPajak t.JenisPajak `json:"jenisPajak" validate:"required"`
	Golongan t.Golongan `json:"golongan" validate:"required"`
	Npwp     *string    `json:"npwp"`

	// Nomor                 int  `json:"nomor"`
	// IsNomorRegistrasiAuto bool `json:"isNomorRegistrasiAuto"`

	TanggalPenutupan *string `json:"tanggalPenutupan"`
	TanggalBuka      *string `json:"tanggalBuka"`
	Skpd_Id          *uint64 `json:"skpd_id"`
	Rekening_Id      *uint64 `json:"rekening_id" validate:"required"`
	User_Name        *string `json:"user_name"`

	TanggalMulaiUsaha *string  `json:"tanggalMulaiUsaha"`
	LuasBangunan      *string  `json:"luasBangunan"`
	JamBukaUsaha      *string  `json:"jamBukaUsaha"`
	JamTutupUsaha     *string  `json:"jamTutupUsaha"`
	Pengunjung        *string  `json:"pengunjung"`
	OmsetOp           *string  `json:"omsetOp"`
	FotoKtp           string   `json:"fotoKtp" validate:"required"`
	SuratIzinUsaha    []string `json:"suratIzinUsaha" validate:"required"`
	LainLain          []string `json:"lainLain"`
	FotoObjek         []string `json:"fotoObjek" validate:"required"`

	Genset   bool `json:"genset" validate:"required"`
	AirTanah bool `json:"airTanah" validate:"required"`

	DetailRegOp *[]DetailRegObjekPajak `json:"detailRegObjekPajak"`

	RegObjekPajak *rop.RegObjekPajakCreate `json:"regObjekPajak"`
	RegPemilik    *[]RegPemilikWpCreate    `json:"regPemilik"`
	RegNarahubung *[]RegNarahubungCreate   `json:"regNarahubung"`
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

	DetailRegObjekPajak []DetailRegObjekPajakUpdate `json:"detailRegObjekPajak"`

	RegObjekPajak rop.RegObjekPajakUpdate `json:"regObjekPajak"`
	RegPemilik    []RegPemilikWpUpdate    `json:"regPemilik"`
	RegNarahubung []RegNarahubungUpdate   `json:"regNarahubung"`
}

type VerifikasiDto struct {
	VerifyStatus int16 `json:"verifyStatus"`
}

type FilterDto struct {
	User_Id  *uint64 `json:"user_id"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}
