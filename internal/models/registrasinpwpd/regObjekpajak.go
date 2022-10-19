package registrasinpwpd

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegObjekPajak struct {
	Id                 uint64                  `json:"id" gorm:"primarykey"`
	RegistrasiNpwpd_Id uint64                  `json:"npwpd_id"`
	RegistrasiNpwpd    *RegistrasiNpwpd        `json:"npwpd,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id"`
	Nama               *string                 `json:"nama" gorm:"size:200"`
	Nop                *string                 `json:"nop" gorm:"size:50"`
	Alamat             *string                 `json:"alamat" gorm:"size:200"`
	RtRw               *string                 `json:"rtRw" gorm:"size:10"`
	Kecamatan_Id       *uint64                 `json:"kecamatan_id"`
	Kecamatan          *areadivision.Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Id"`
	Kelurahan_Id       *uint64                 `json:"kelurahan_id"`
	Kelurahan          *areadivision.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Id"`
	Telp               *string                 `json:"telp" gorm:"size:20"`
	Status             t.StatusBL              `json:"status"`
	IsNpwpd            uint64                  `json:"isNpwpd"`
	gormhelper.DateModel
}

type DetailRegOp struct {
	Id                 uint64           `json:"id" gorm:"primaryKey"`
	RegistrasiNpwpd_Id uint64           `json:"registrasiNpwpd_id"`
	RegistrasiNpwpd    *RegistrasiNpwpd `json:"npwpd,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	JenisOp            *string          `json:"jenisOp" gorm:"size:200"`
	JumlahOp           *string          `json:"jumlahOp" gorm:"size:200"`
	TarifOp            *string          `json:"tarifOp" gorm:"size:200"`
	UnitOp             *string          `json:"unitOp" gorm:"size:50"`
	Notes              *string          `json:"notes" gorm:"size:200"`
}

type DetailRegOpAirTanah struct {
	DetailRegOp
}

type DetailRegOpHiburan struct {
	DetailRegOp
}

type DetailRegOpHotel struct {
	DetailRegOp
}

type DetailRegOpParkir struct {
	DetailRegOp
}

type DetailRegOpPpj struct {
	DetailRegOp
}

type DetailRegOpReklame struct {
	DetailRegOp
}

type DetailRegOpResto struct {
	DetailRegOp
}

type DetailRegOpCreate struct {
	JumlahOp *string `json:"jumlahOp" gorm:"size:200"`
	TarifOp  *string `json:"tarifOp" gorm:"size:200"`
	UnitOp   *string `json:"unitOp" gorm:"size:50"`
	Notes    *string `json:"notes" gorm:"size:200"`
}

type DetailRegOpUpdate struct {
	Id       uint64  `json:"id" gorm:"primaryKey"`
	JumlahOp *string `json:"jumlahOp" gorm:"size:200"`
	TarifOp  *string `json:"tarifOp" gorm:"size:200"`
	UnitOp   *string `json:"unitOp" gorm:"size:50"`
	Notes    *string `json:"notes" gorm:"size:200"`
}

type RegObjekPajakCreate struct {
	Nama         *string `json:"nama" gorm:"size:200"`
	Nop          *string `json:"nop" gorm:"size:50"`
	Alamat       *string `json:"alamat" gorm:"size:200"`
	RtRw         *string `json:"rtRw" gorm:"size:10"`
	Kecamatan_Id *uint64 `json:"kecamatan_id"`
	Kelurahan_Id *uint64 `json:"kelurahan_id"`
	Telp         *string `json:"telp" gorm:"size:20"`
}

type RegObjekPajakUpdate struct {
	Id           uint64  `json:"id" gorm:"primaryKey"`
	Nama         *string `json:"nama" gorm:"size:200"`
	Nop          *string `json:"nop" gorm:"size:50"`
	Alamat       *string `json:"alamat" gorm:"size:200"`
	RtRw         *string `json:"rtRw" gorm:"size:10"`
	Kecamatan_Id *uint64 `json:"kecamatan_id"`
	Kelurahan_Id *uint64 `json:"kelurahan_id"`
	Telp         *string `json:"telp" gorm:"size:20"`
}
