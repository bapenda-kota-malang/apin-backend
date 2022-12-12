package regobjekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type RegJpb16 struct {
	nop.NopDetail
	NoBangunan      *int          `json:"noBangunan"`
	KelasBangunan16 KelasBangunan `json:"kelasBangunan16" gorm:"type:char(1)"`
	gh.DateModel
}

type RegJpb16CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan      *int          `json:"noBangunan"`
	KelasBangunan16 KelasBangunan `json:"kelasBangunan16"`
}

type RegJpb16UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan      *int          `json:"noBangunan"`
	KelasBangunan16 KelasBangunan `json:"kelasBangunan16"`
}

type RegJpb16FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}

func (input *RegOpbJpb16CreateDto) GetFasilitasBangunan() *mrfb.CreateDto {
	return input.RegFasilitasBangunans
}

func (input *RegOpbJpb16CreateDto) GetNop() *string {
	return input.Nop
}

func (input *RegOpbJpb16CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *RegOpbJpb16CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *RegOpbJpb16CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *RegOpbJpb16CreateDto) GetObjekPajakBangunan() (CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return CreateDto{}, err
	}
	return data, nil
}
