package regobjekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type RegJpb14 struct {
	nop.NopDetail
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
	gh.DateModel
}

type RegJpb14CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
}

type RegJpb14UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
}

type RegJpb14FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}

func (input *RegOpbJpb14CreateDto) GetFasilitasBangunan() *mrfb.CreateDto {
	return input.RegFasilitasBangunans
}

func (input *RegOpbJpb14CreateDto) GetNop() *string {
	return input.Nop
}

func (input *RegOpbJpb14CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *RegOpbJpb14CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *RegOpbJpb14CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *RegOpbJpb14CreateDto) GetObjekPajakBangunan() (CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return CreateDto{}, err
	}
	return data, nil
}
