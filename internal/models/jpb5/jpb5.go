package jpb5

import (
	jt "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb2/types"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb5 struct {
	nop.NopDetail
	NoBangunan             *int              `json:"noBangunan"`
	KelasBangunan5         *jt.KelasBangunan `json:"kelasBangunan5" gorm:"type:char(1)"`
	LuasKamarAcCentral     *int              `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int              `json:"luasRuangLainAcCentral"`
	gh.DateModel
}

type CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan             *int              `json:"noBangunan"`
	KelasBangunan5         *jt.KelasBangunan `json:"kelasBangunan5"`
	LuasKamarAcCentral     *int              `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int              `json:"luasRuangLainAcCentral"`
}

type UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan             *int              `json:"noBangunan"`
	KelasBangunan5         *jt.KelasBangunan `json:"kelasBangunan5"`
	LuasKamarAcCentral     *int              `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int              `json:"luasRuangLainAcCentral"`
}

type FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan             *int `json:"noBangunan"`
	LuasKamarAcCentral     *int `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int `json:"luasRuangLainAcCentral"`
	Page                   int  `json:"page"`
	PageSize               int  `json:"page_size"`
}
