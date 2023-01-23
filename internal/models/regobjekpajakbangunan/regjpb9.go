package regobjekpajakbangunan

import (
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type RegJpb9 struct {
	Id         uint64 `json:"id" gorm:"primarykey;autoIncrement"`
	NoBangunan *int   `json:"noBangunan"`
	gh.DateModel
}

type RegJpb9CreateDto struct {
	NoBangunan *int `json:"noBangunan"`
}

type RegJpb9UpdateDto struct {
	NoBangunan *int `json:"noBangunan"`
}

type RegJpb9FilterDto struct {
	NoBangunan *int `json:"noBangunan"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}

func (input *RegOpbJpb9CreateDto) GetFasilitasBangunan() *mrfb.CreateDto {
	return input.RegFasilitasBangunans
}

func (input *RegOpbJpb9CreateDto) GetNop() *string {
	return input.Nop
}

func (input *RegOpbJpb9CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *RegOpbJpb9CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *RegOpbJpb9CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *RegOpbJpb9CreateDto) GetObjekPajakBangunan() (CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return CreateDto{}, err
	}
	return data, nil
}
