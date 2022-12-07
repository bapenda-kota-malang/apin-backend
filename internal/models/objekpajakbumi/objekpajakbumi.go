package objekpajakbumi

import (
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type ObjekPajakBumi struct {
	opp.NopDetail
	NoBumi          int       `json:"noBumi"`
	KodeZNT         string    `json:"kodeZnt" gorm:"type:char(2)"`
	LuasBumi        int       `json:"luasBumi"`
	JenisBumi       JenisBumi `json:"jenisBumi" gorm:"type:char(1)"`
	NilaiSistemBumi int       `json:"nilaiSistemBumi"`
	gh.DateModel
}
