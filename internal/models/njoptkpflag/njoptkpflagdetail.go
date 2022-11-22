package njoptkpflag

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type NjoptkpFlagDetail struct {
	Id              uint64       `json:"id" gorm:"primaryKey"`
	NjoptkpFlag_Id  uint64       `json:"njoptkpFlag_id"`
	NjoptkpFlag     *NjoptkpFlag `json:"njoptkpFlag,omitempty" gorm:"foreignKey:NjoptkpFlag_Id"`
	Nop             *string      `json:"nop"`
	LetakObjekPajak *string      `json:"letakObjekPajak"`
	LuasBumi        *float64     `json:"luasBumi" gorm:"type:decimal"`
	LuasBg          *float64     `json:"luasBg" gorm:"type:decimal"`
	Njop            *string      `json:"njop"`
	Bumi            *string      `json:"bumi"`
	Bgn             *string      `json:"bgn"`
	IsNjoptkpFlag   bool         `json:"isNjoptkpFlag"`
	gh.DateModel
}

type NjoptkpFlagDetailCreateDto struct {
	Nop             *string  `json:"nop"`
	LetakObjekPajak *string  `json:"letakObjekPajak"`
	LuasBumi        *float64 `json:"luasBumi"`
	LuasBg          *float64 `json:"luasBg"`
	Njop            *string  `json:"njop"`
	Bumi            *string  `json:"bumi"`
	Bgn             *string  `json:"bgn"`
	IsNjoptkpFlag   bool     `json:"isNjoptkpFlag"`
}

type NjoptkpFlagDetailUpdateDto struct {
	Id              uint64   `json:"id"`
	Nop             *string  `json:"nop"`
	LetakObjekPajak *string  `json:"letakObjekPajak"`
	LuasBumi        *float64 `json:"luasBumi"`
	LuasBg          *float64 `json:"luasBg"`
	Njop            *string  `json:"njop"`
	Bumi            *string  `json:"bumi"`
	Bgn             *string  `json:"bgn"`
	IsNjoptkpFlag   bool     `json:"isNjoptkpFlag"`
}
