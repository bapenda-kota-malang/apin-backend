package objekpajakbangunan

import (
	mfb "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type Jpb13 struct {
	nop.NopDetail
	NoBangunan             *int          `json:"noBangunan"`
	KelasBangunan13        KelasBangunan `json:"kelasBangunan13" gorm:"type:char(1)"`
	JumlahApartment        *int          `json:"jumlahApartment"`
	LuasApartAcCentral     *int          `json:"luasApartAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
	gh.DateModel
}

type Jpb13CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan             *int          `json:"noBangunan"`
	KelasBangunan13        KelasBangunan `json:"kelasBangunan13"`
	JumlahApartment        *int          `json:"jumlahApartment"`
	LuasApartAcCentral     *int          `json:"luasApartAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
}

type Jpb13UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan             *int          `json:"noBangunan"`
	KelasBangunan13        KelasBangunan `json:"kelasBangunan13"`
	JumlahApartment        *int          `json:"jumlahApartment"`
	LuasApartAcCentral     *int          `json:"luasApartAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
}

type Jpb13FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan             *int `json:"noBangunan"`
	JumlahApartment        *int `json:"jumlahApartment"`
	LuasApartAcCentral     *int `json:"luasApartAcCentral"`
	LuasRuangLainAcCentral *int `json:"luasRuangLainAcCentral"`
	Page                   int  `json:"page"`
	PageSize               int  `json:"page_size"`
}

func (input *OpbJpb13CreateDto) GetFasilitasBangunan() *mfb.CreateDto {
	return input.FasilitasBangunans
}

func (input *OpbJpb13CreateDto) GetNop() *string {
	return input.Nop
}

func (input *OpbJpb13CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *OpbJpb13CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *OpbJpb13CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *OpbJpb13CreateDto) GetObjekPajakBangunan() (*CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return nil, err
	}
	return &data, nil
}
