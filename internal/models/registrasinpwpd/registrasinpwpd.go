package registrasinpwpd

import (
	"time"

	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/skpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegistrasiNpwpd struct {
	Id       uint64     `json:"id" gorm:"primarykey"`
	Golongan t.Golongan `json:"golongan"`
	Nomor    string     `json:"nomor" gorm:"type:varchar(10)"`
	Npwp     *string    `json:"npwp" gorm:"size:50"`
	// TanggalPengukuhan *time.Time               `json:"tanggalPengukuhan"`
	// TanggalNpwpd      *time.Time               `json:"tanggalNpwpd"`
	// Npwpd             *string                  `json:"npwpd" gorm:"size:22"`
	Status            t.Status           `json:"status"`
	TanggalPenutupan  *time.Time         `json:"tanggalPenutupan"`
	TanggalBuka       *time.Time         `json:"tanggalBuka"`
	JenisPajak        t.JenisPajak       `json:"jenisPajak" gorm:"size:2"`
	Skpd_Id           *uint64            `json:"skpd_id"`
	Skpd              *skpd.Skpd         `json:"skpd,omitempty" gorm:"foreignKey:Skpd_Id"`
	Rekening_Id       *uint64            `json:"rekening_id"`
	Rekening          *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	User_Name         *string            `json:"user_name" gorm:"size:20"`
	User_Id           uint64             `json:"user_id"`
	User              *user.User         `gorm:"foreignKey:User_Id;references:Id"`
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
	VerifyStatus   *VerifyStatus `json:"verifyStatus"`
	VerifiedAt     *time.Time    `json:"verifiedAt"`
	FotoKtp        string        `json:"fotoKtp" gorm:"size:50"`
	SuratIzinUsaha string        `json:"suratIzinUsaha" gorm:"size:2048"`
	LainLain       string        `json:"lainLain" gorm:"size:2048"`
	FotoObjek      string        `json:"fotoObjek" gorm:"size:2048"`
	// ModeRegistrasi    npwpd.Mode               `json:"modeRegistrasi"`
	// VendorEtax         *configurationmodel.VendorEtax `gorm:"foreignKey:VendorEtaxID"`

	RegObjekPajaks []*RegObjekPajak `json:"regobjekPajak,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	RegPemilikWps  []*RegPemilikWp  `json:"regpemilik,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	RegNarahubungs []*RegNarahubung `json:"regnarahubung,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`

	DetailRegOpAirTanah []*DetailRegOpAirTanah `json:"detail_reg_op_air_tanah,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	DetailRegOpHiburan  []*DetailRegOpHiburan  `json:"detail_reg_op_hiburan,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	DetailRegOpHotel    []*DetailRegOpHotel    `json:"detail_reg_op_hotel,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	DetailRegOpParkir   []*DetailRegOpParkir   `json:"detail_reg_op_parkir,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	DetailRegOpPpj      []*DetailRegOpPpj      `json:"detail_reg_op_ppj,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	DetailRegOpReklame  []*DetailRegOpReklame  `json:"detail_reg_op_reklame,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	DetailRegOpResto    []*DetailRegOpResto    `json:"detail_reg_op_resto,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
}

type CreateDto struct {
	JenisPajak t.JenisPajak `json:"jenisPajak" validate:"required"`
	Golongan   t.Golongan   `json:"golongan" validate:"required"`
	Npwp       *string      `json:"npwp" validate:"required"`

	Nomor                 string `json:"nomor"`
	IsNomorRegistrasiAuto bool   `json:"isNomorRegistrasiAuto"`

	TanggalPenutupan *string `json:"tanggalPenutupan"`
	TanggalBuka      *string `json:"tanggalBuka"`
	Skpd_Id          *uint64 `json:"skpd_id"`
	Rekening_Id      *uint64 `json:"rekening_id"`
	User_Name        *string `json:"user_name"`

	TanggalMulaiUsaha *string  `json:"tanggalMulaiUsaha"`
	LuasBangunan      *string  `json:"luasBangunan"`
	JamBukaUsaha      *string  `json:"jamBukaUsaha"`
	JamTutupUsaha     *string  `json:"jamTutupUsaha"`
	Pengunjung        *string  `json:"pengunjung"`
	OmsetOp           *string  `json:"omsetOp"`
	FotoKtp           string   `json:"fotoKtp"`
	SuratIzinUsaha    []string `json:"suratIzinUsaha"`
	LainLain          []string `json:"lainLain"`
	FotoObjek         []string `json:"fotoObjek"`

	Genset   bool `json:"genset"`
	AirTanah bool `json:"airTanah"`

	DetailRegOp *[]DetailRegOp `json:"detail_reg_op"`

	RegObjekPajak *RegObjekPajak   `json:"regObjekPajak"`
	RegPemilik    *[]RegPemilikWp  `json:"regPemilik"`
	RegNarahubung *[]RegNarahubung `json:"regNarahubung"`
}

type UpdateDto struct {
	TanggalPenutupan *string `json:"tanggalPenutupan"`
	TanggalBuka      *string `json:"tanggalBuka"`
	Skpd_Id          *uint64 `json:"skpd_id"`
	Rekening_Id      *uint64 `json:"rekening_id"`
	User_Name        *string `json:"user_name"`

	TanggalMulaiUsaha *string  `json:"tanggalMulaiUsaha"`
	LuasBangunan      *string  `json:"luasBangunan"`
	JamBukaUsaha      *string  `json:"jamBukaUsaha"`
	JamTutupUsaha     *string  `json:"jamTutupUsaha"`
	Pengunjung        *string  `json:"pengunjung"`
	OmsetOp           *string  `json:"omsetOp"`
	FotoKtp           string   `json:"fotoKtp"`
	SuratIzinUsaha    []string `json:"suratIzinUsaha"`
	LainLain          []string `json:"lainLain"`
	FotoObjek         []string `json:"fotoObjek"`

	Genset   bool `json:"genset"`
	AirTanah bool `json:"airTanah"`

	DetailRegOp []DetailRegOp `json:"detail_reg_op"`

	RegObjekPajak RegObjekPajak   `json:"regObjekPajak"`
	RegPemilik    []RegPemilikWp  `json:"regPemilik"`
	RegNarahubung []RegNarahubung `json:"regNarahubung"`
}

type VerifikasiDto struct {
	VerifyStatus int16 `json:"verifyStatus"`
}

type FilterDto struct {
	User_Id  *uint64 `json:"user_id"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}
