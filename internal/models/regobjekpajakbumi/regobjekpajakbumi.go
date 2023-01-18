package regobjekpajakbumi

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	mopbng "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegObjekPajakBumi struct {
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

	RegObjekPajakBangunans *[]mopbng.CreateDto `json:"regObjekPajakBng"`
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
