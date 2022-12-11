package objekpajakbangunan

import (
	mfb "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type Jpb16 struct {
	nop.NopDetail
	NoBangunan      *int          `json:"noBangunan"`
	KelasBangunan16 KelasBangunan `json:"kelasBangunan16" gorm:"type:char(1)"`
	gh.DateModel
}

type Jpb16CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan      *int          `json:"noBangunan"`
	KelasBangunan16 KelasBangunan `json:"kelasBangunan16"`
}

type Jpb16UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan      *int          `json:"noBangunan"`
	KelasBangunan16 KelasBangunan `json:"kelasBangunan16"`
}

type Jpb16FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}

func (input *OpbJpb16CreateDto) GetFasilitasBangunan() *mfb.CreateDto {
	return input.FasilitasBangunans
}

func (input *OpbJpb16CreateDto) GetNop() *string {
	return input.Nop
}

func (input *OpbJpb16CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *OpbJpb16CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *OpbJpb16CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *OpbJpb16CreateDto) GetObjekPajakBangunan() (*CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return nil, err
	}
	return &data, nil
}
