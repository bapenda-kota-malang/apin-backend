package jpb13

import (
	jt "github.com/bapenda-kota-malang/apin-backend/internal/models/jpb2/types"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb13 struct {
	nop.NopDetail
	NoBangunan             *int             `json:"noBangunan"`
	KelasBangunan13        jt.KelasBangunan `json:"kelasBangunan13" gorm:"type:char(1)"`
	JumlahApartment        *int             `json:"jumlahApartment"`
	LuasApartAcCentral     *int             `json:"luasApartAcCentral"`
	LuasRuangLainAcCentral *int             `json:"luasRuangLainAcCentral"`
	gh.DateModel
}

type CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan             *int             `json:"noBangunan"`
	KelasBangunan13        jt.KelasBangunan `json:"kelasBangunan13"`
	JumlahApartment        *int             `json:"jumlahApartment"`
	LuasApartAcCentral     *int             `json:"luasApartAcCentral"`
	LuasRuangLainAcCentral *int             `json:"luasRuangLainAcCentral"`
}

type UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan             *int             `json:"noBangunan"`
	KelasBangunan13        jt.KelasBangunan `json:"kelasBangunan13"`
	JumlahApartment        *int             `json:"jumlahApartment"`
	LuasApartAcCentral     *int             `json:"luasApartAcCentral"`
	LuasRuangLainAcCentral *int             `json:"luasRuangLainAcCentral"`
}

type FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan             *int `json:"noBangunan"`
	JumlahApartment        *int `json:"jumlahApartment"`
	LuasApartAcCentral     *int `json:"luasApartAcCentral"`
	LuasRuangLainAcCentral *int `json:"luasRuangLainAcCentral"`
	Page                   int  `json:"page"`
	PageSize               int  `json:"page_size"`
}
