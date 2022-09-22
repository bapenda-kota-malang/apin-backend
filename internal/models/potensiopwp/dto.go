package potensiopwp

import (
	"time"

	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
)

type CreatePotensiOp struct {
	Golongan      rm.Golongan `json:"golongan" validate:"required,oneof=badan orang_pribadi"`
	Rekening_Id   uint        `json:"rekeningId" validate:"required,gt=0"`
	ClosingDate   *time.Time  `json:"closingDate"`
	OpeningDate   *time.Time  `json:"openingDate"`
	User_Id       uint        `json:"userId" validate:"required,gt=0"`
	LuasBangunan  *string     `json:"luasBangunan"`
	JamBuka       *string     `json:"jamBuka"`
	JamTutup      *string     `json:"jamTutup"`
	Visitors      string      `json:"visitors" validate:"required"`
	OmsetOp       string      `json:"omsetOp" validate:"required"`
	Genset        string      `json:"genset" validate:"required"`
	AirTanah      string      `json:"airTanah" validate:"required"`
	VendorEtax_Id *uint       `json:"vendorEtaxId"`
}
type CreateDetailPotensiOp struct {
	Potensiop_Id uint   `json:"potensiopId" validate:"required"`
	Nama         string `json:"nama" validate:"required"`
	Alamat       string `json:"alamat" validate:"required"`
	RtRw         string `json:"rtRw" validate:"required"`
	Kecamatan_Id uint   `json:"kecamatanId" validate:"required"`
	Kelurahan_Id uint   `json:"kelurahanId" validate:"required"`
	Telp         string `json:"telp" validate:"required"`
	Status       string `json:"status" validate:"required"`
	IsNpwpd      uint8  `json:"isNpwpd" validate:"required"`
}
type CreatePotensiPemilikWp struct {
	Potensiop_Id uint   `json:"potensiopId" validate:"required,gt=0"`
	Nama         string `json:"nama" validate:"required"`
	Alamat       string `json:"alamat" validate:"required"`
	RtRw         string `json:"rtRw" validate:"required"`
	Kecamatan_Id uint   `json:"kecamatanId" validate:"required,gt=0"`
	Kelurahan_Id uint   `json:"kelurahanId" validate:"required,gt=0"`
	Telp         string `json:"telp" validate:"required"`
	Status       string `json:"status" validate:"required"`
	NoIdPemilik  string `json:"noIdPemilik" validate:"required"`
	Nik          string `json:"nik" validate:"required"`
}

type CreatePotensiNarahubung struct {
	Potensiop_Id uint   `json:"potensiopId"`
	Nama         string `json:"nama"`
	Alamat       string `json:"alamat"`
	RtRw         string `json:"rtRw"`
	Kecamatan_Id uint   `json:"kecamatanId"`
	Kelurahan_Id uint   `json:"kelurahanId"`
	Telp         string `json:"telp"`
	Status       string `json:"status"`
	Nik          string `json:"nik"`
}

type DetailPajakDto struct {
	Potensiop_Id uint   `json:"potensiopId"`
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
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}
