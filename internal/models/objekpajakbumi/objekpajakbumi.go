package objekpajakbumi

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type ObjekPajakBumi struct {
	nop.NopDetail
	NoBumi          *int      `json:"noBumi"`
	KodeZNT         *string   `json:"kodeZnt" gorm:"type:char(2)"`
	LuasBumi        int       `json:"luasBumi"`
	JenisBumi       JenisBumi `json:"jenisBumi" gorm:"type:char(1)"`
	NilaiSistemBumi *int      `json:"nilaiSistemBumi"`
	gh.DateModel
}

// untuk spop
type CreateDto struct {
	nop.NopDetailCreateDto
	NoBumi          *int      `json:"noBumi"`
	KodeZNT         *string   `json:"kodeZnt"`
	LuasBumi        int       `json:"luasBumi"`
	JenisBumi       JenisBumi `json:"jenisBumi"`
	NilaiSistemBumi *int      `json:"nilaiSistemBumi"`
}

type DataDto struct {
	Id      int     `json:"id" validate:"required"`
	NoUrut  *string `json:"no_urut" validate:"required;minLength=4;maxLength=4"`
	JenisOp *string `json:"jns_op" validate:"required;minLength=1;maxLength=1"`
}

// untuk znt masal
type UpdateBulkDto struct {
	nop.NopDetailCreateDto
	KodeZNT        *string         `json:"znt_kode"`
	KodeZNTBaru    *string         `json:"znt_kode_baru"`
	NomerDokumen   string          `json:"nomerDokumen" validate:"required;minLength=11;maxLength=11"`
	Kpbb_Id        *string         `json:"kpbb_id" validate:"required;minLength=2;maxLength=2"`
	Kanwil_Id      *string         `json:"kanwil_id" validate:"required;minLength=2;maxLength=2"`
	JenisDokumen   *string         `json:"jenisDokumen" validate:"required;minLength=1;maxLength=1"`
	TglPendataan   *datatypes.Date `json:"tglPendataan"`
	NipPendataan   *string         `json:"nipPendataan" validate:"required;minLength=1;maxLength=9"`
	TglPemeriksaan *datatypes.Date `json:"tglPemeriksaan"`
	NipPemeriksaan *string         `json:"nipPemeriksaan" validate:"required;minLength=1;maxLength=9"`
	Datas          []DataDto       `json:"datas" validate:"required"`
}

type UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBumi          *int      `json:"noBumi"`
	KodeZNT         *string   `json:"kodeZnt"`
	LuasBumi        int       `json:"luasBumi"`
	JenisBumi       JenisBumi `json:"jenisBumi"`
	NilaiSistemBumi *int      `json:"nilaiSistemBumi"`
}

type FilterDto struct {
	nop.NopDetailCreateDto
	NoBumi          *int    `json:"noBumi"`
	KodeZNT         *string `json:"kodeZnt"`
	LuasBumi        *int    `json:"luasBumi"`
	NilaiSistemBumi *int    `json:"nilaiSistemBumi"`
	Page            int     `json:"page"`
	PageSize        int     `json:"page_size"`
}
