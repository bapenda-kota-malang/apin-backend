package sejarah

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	opb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type SejarahObjekPajakBumi struct {
	nop.NopDetail
	ObjekPajakBumi_Id *uint64       `json:"objekPajakBumi_Id"`
	NoBumi            *int          `json:"noBumi"`
	KodeZNT           *string       `json:"kodeZnt" gorm:"type:char(2)"`
	LuasBumi          int           `json:"luasBumi"`
	JenisBumi         opb.JenisBumi `json:"jenisBumi" gorm:"type:char(1)"`
	NilaiSistemBumi   *int          `json:"nilaiSistemBumi"`
	gh.DateModel
}

type ListSejarahOpBumi struct {
	KodeZNT         *string       `json:"kodeZnt"`
	LuasBumi        int           `json:"luasBumi"`
	JenisBumi       opb.JenisBumi `json:"jenisBumi"`
	NilaiSistemBumi *int          `json:"nilaiSistemBumi"`
}
