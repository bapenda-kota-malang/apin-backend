package npwpd

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type ObjekPajak struct {
	Id           uint64                  `json:"id" gorm:"primarykey"`
	Npwpd_Id     uint64                  `json:"npwpd_id"`
	Npwpd        *Npwpd                  `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	Nama         *string                 `json:"nama" gorm:"size:200"`
	Nop          *string                 `json:"nop" gorm:"size:50"`
	Alamat       *string                 `json:"alamat" gorm:"size:200"`
	RtRw         *string                 `json:"rtRw" gorm:"size:10"`
	Kecamatan_Id *uint64                 `json:"kecamatan_id"`
	Kecamatan    *areadivision.Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Id"`
	Kelurahan_Id *uint64                 `json:"kelurahan_id"`
	Kelurahan    *areadivision.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Id"`
	Telp         *string                 `json:"telp" gorm:"size:20"`
	Status       StatusBL                `json:"status"`
	IsNpwpd      uint64                  `json:"isNpwpd"`
	gormhelper.DateModel
}

type DetailOp struct {
	Id       uint64  `json:"id" gorm:"primaryKey"`
	Npwpd_Id uint64  `json:"npwpd_id"`
	Npwpd    *Npwpd  `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	JenisOp  *string `json:"jenisOp" gorm:"size:200"`
	JumlahOp *string `json:"jumlahOp" gorm:"size:200"`
	TarifOp  *string `json:"tarifOp" gorm:"size:200"`
	UnitOp   *string `json:"unitOp" gorm:"size:50"`
	Notes    *string `json:"notes" gorm:"size:200"`
}

type DetailOpAirTanah struct {
	DetailOp
}

type DetailOpHiburan struct {
	DetailOp
}

type DetailOpHotel struct {
	DetailOp
}

type DetailOpParkir struct {
	DetailOp
}

type DetailOpPpj struct {
	DetailOp
}

type DetailOpReklame struct {
	DetailOp
}

type DetailOpResto struct {
	DetailOp
}

type DetailOpUpdateDto struct {
	JumlahOp *string `json:"jumlahOp" gorm:"size:200"`
	TarifOp  *string `json:"tarifOp" gorm:"size:200"`
	UnitOp   *string `json:"unitOp" gorm:"size:50"`
	Notes    *string `json:"notes" gorm:"size:200"`
}
