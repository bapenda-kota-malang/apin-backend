package potensiopwp

import (
	"time"

	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
)

type CreatePotensiOp struct {
	Golongan      rm.Golongan `json:"golongan" validate:"required,oneof=badan orang_pribadi"`
	Rekening_Id   uint        `json:"rekening_id" validate:"required,gt=0"`
	ClosingDate   *time.Time  `json:"closingDate"`
	OpeningDate   *time.Time  `json:"openingDate"`
	User_Id       uint        `json:"user_id" validate:"required,gt=0"`
	LuasBangunan  *string     `json:"luasBangunan"`
	JamBuka       *string     `json:"jamBuka"`
	JamTutup      *string     `json:"jamTutup"`
	Visitors      string      `json:"visitors" validate:"required"`
	OmsetOp       string      `json:"omsetOp" validate:"required"`
	Genset        string      `json:"genset" validate:"required"`
	AirTanah      string      `json:"airTanah" validate:"required"`
	VendorEtax_Id *uint       `json:"vendorEtax_id"`
}
type CreateDetailPotensiOp struct {
	Nama         string `json:"nama" validate:"required"`
	Alamat       string `json:"alamat" validate:"required"`
	RtRw         string `json:"rtRw" validate:"required"`
	Kecamatan_Id uint   `json:"kecamatan_id" validate:"required"`
	Kelurahan_Id uint   `json:"kelurahan_id" validate:"required"`
	Telp         string `json:"telp" validate:"required"`
	Status       string `json:"status" validate:"required"`
	IsNpwpd      uint8  `json:"isNpwpd" validate:"required"`
}
type CreatePotensiPemilikWp struct {
	Nama         string `json:"nama" validate:"required"`
	Alamat       string `json:"alamat" validate:"required"`
	RtRw         string `json:"rtRw" validate:"required"`
	Kecamatan_Id uint   `json:"kecamatan_id" validate:"required"`
	Kelurahan_Id uint   `json:"kelurahan_id" validate:"required"`
	Telp         string `json:"telp" validate:"required"`
	Status       string `json:"status" validate:"required"`
	Nik          string `json:"nik" validate:"required"`
	NoIdPemilik  string `json:"noIdPemilik" validate:"required"`
}

type CreatePotensiNarahubung struct {
	Nama         string `json:"nama" validate:"required"`
	Alamat       string `json:"alamat" validate:"required"`
	RtRw         string `json:"rtRw" validate:"required"`
	Kecamatan_Id uint   `json:"kecamatan_id" validate:"required"`
	Kelurahan_Id uint   `json:"kelurahan_id" validate:"required"`
	Telp         string `json:"telp" validate:"required"`
	Status       string `json:"status" validate:"required"`
	Nik          string `json:"nik"`
}

type DetailPajakDto struct {
	Id           *uint  `json:"id"`
	Potensiop_Id uint   `json:"potensiop_id"`
	JenisOp      string `json:"jenisOp" validate:"oneof='Resto' 'Reklame' 'PPJ' 'Parkir' 'Hotel' 'Hiburan' 'Air Tanah'"`
	JumlahOp     string `json:"jumlahOp"`
	TarifOp      string `json:"tarifOp"`
	UnitOp       string `json:"unitOp"`
	Notes        string `json:"notes"`
}

type CreateDto struct {
	CreatePotensiOp         `json:"potensiOp"`
	CreateDetailPotensiOp   `json:"detailPotensiOp"`
	CreatePotensiPemilikWp  `json:"potensiPemilikWp"`
	CreatePotensiNarahubung *CreatePotensiNarahubung `json:"potensiNarahubung"`
	DetailPajakDtos         []DetailPajakDto         `json:"detailPajaks"`
}

type FilterDto struct {
	// dynamic from column table field
	// potensiop
	Golongan     *rm.Golongan `json:"golongan"`
	Status       *rm.Status   `json:"status"`
	ClosingDate  *time.Time   `json:"closingDate"`
	OpeningDate  *time.Time   `json:"openingDate"`
	LuasBangunan *string      `json:"luasBangunan"`
	JamBuka      *string      `json:"jamBuka"`
	JamTutup     *string      `json:"jamTutup"`
	Visitors     *string      `json:"visitors"`
	OmsetOp      *string      `json:"omsetOp"`
	Genset       *string      `json:"genset"`
	AirTanah     *string      `json:"airTanah"`
	// detail potensiop
	JenisOp  *string `json:"jenisOp"`
	JumlahOp *string `json:"jumlahOp"`
	TarifOp  *string `json:"tarifOp"`
	UnitOp   *string `json:"unitOp"`
	Notes    *string `json:"notes"`
	// fixed
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}
