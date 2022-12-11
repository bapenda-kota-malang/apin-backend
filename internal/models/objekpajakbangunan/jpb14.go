package objekpajakbangunan

import (
	mfb "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type Jpb14 struct {
	nop.NopDetail
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
	gh.DateModel
}

type Jpb14CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
}

type Jpb14UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
}

type Jpb14FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	LuasKanopi *int `json:"luasKanopi"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}

func (input *OpbJpb14CreateDto) GetFasilitasBangunan() *mfb.CreateDto {
	return input.FasilitasBangunans
}

func (input *OpbJpb14CreateDto) GetNop() *string {
	return input.Nop
}

func (input *OpbJpb14CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *OpbJpb14CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *OpbJpb14CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *OpbJpb14CreateDto) GetObjekPajakBangunan() (*CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return nil, err
	}
	return &data, nil
}
