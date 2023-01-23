package regobjekpajakbangunan

import (
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type RegJpb4 struct {
	Id         uint64 `json:"id" gorm:"primarykey;autoIncrement"`
	NoBangunan *int   `json:"noBangunan"`
	gh.DateModel
}

type RegJpb4CreateDto struct {
	NoBangunan *int `json:"noBangunan"`
}

type RegJpb4UpdateDto struct {
	NoBangunan *int `json:"noBangunan"`
}

type RegJpb4FilterDto struct {
	NoBangunan *int `json:"noBangunan"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}

func (input *RegOpbJpb4CreateDto) GetFasilitasBangunan() *mrfb.CreateDto {
	return input.RegFasilitasBangunans
}

func (input *RegOpbJpb4CreateDto) GetNop() *string {
	return input.Nop
}

func (input *RegOpbJpb4CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *RegOpbJpb4CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *RegOpbJpb4CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *RegOpbJpb4CreateDto) GetObjekPajakBangunan() (CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return CreateDto{}, err
	}
	return data, nil
}
