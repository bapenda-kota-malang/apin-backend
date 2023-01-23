package regobjekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type RegJpb13 struct {
	Id                     uint64 `json:"id" gorm:"primarykey;autoIncrement"`
	NoBangunan             *int   `json:"noBangunan"`
	JumlahApartment        *int   `json:"jumlahApartment"`
	LuasApartAcCentral     *int   `json:"luasApartAcCentral"`
	LuasRuangLainAcCentral *int   `json:"luasRuangLainAcCentral"`
	gh.DateModel
}

type RegJpb13CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan             *int `json:"noBangunan"`
	JumlahApartment        *int `json:"jumlahApartment"`
	LuasApartAcCentral     *int `json:"luasApartAcCentral"`
	LuasRuangLainAcCentral *int `json:"luasRuangLainAcCentral"`
}

type RegJpb13UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan             *int `json:"noBangunan"`
	JumlahApartment        *int `json:"jumlahApartment"`
	LuasApartAcCentral     *int `json:"luasApartAcCentral"`
	LuasRuangLainAcCentral *int `json:"luasRuangLainAcCentral"`
}

type RegJpb13FilterDto struct {
	NoBangunan             *int `json:"noBangunan"`
	JumlahApartment        *int `json:"jumlahApartment"`
	LuasApartAcCentral     *int `json:"luasApartAcCentral"`
	LuasRuangLainAcCentral *int `json:"luasRuangLainAcCentral"`
	Page                   int  `json:"page"`
	PageSize               int  `json:"page_size"`
}

func (input *RegOpbJpb13CreateDto) GetFasilitasBangunan() *mrfb.CreateDto {
	return input.RegFasilitasBangunans
}

func (input *RegOpbJpb13CreateDto) GetNop() *string {
	return input.Nop
}

func (input *RegOpbJpb13CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *RegOpbJpb13CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *RegOpbJpb13CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *RegOpbJpb13CreateDto) GetObjekPajakBangunan() (CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return CreateDto{}, err
	}
	return data, nil
}
