package objekpajakbangunan

import (
	mfb "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type Jpb6 struct {
	nop.NopDetail
	NoBangunan   *int          `json:"noBangunan"`
	KelasBanguna KelasBangunan `json:"kelasBangunan6" gorm:"type:char(1)"`
	gh.DateModel
}

type Jpb6CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan     *int          `json:"noBangunan"`
	KelasBangunan6 KelasBangunan `json:"kelasBangunan6"`
}

type Jpb6UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan     *int          `json:"noBangunan"`
	KelasBangunan6 KelasBangunan `json:"kelasBangunan6"`
}

type Jpb6FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan *int `json:"noBangunan"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
}

func (input *OpbJpb6CreateDto) GetFasilitasBangunan() *mfb.CreateDto {
	return input.FasilitasBangunans
}

func (input *OpbJpb6CreateDto) GetNop() *string {
	return input.Nop
}

func (input *OpbJpb6CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *OpbJpb6CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *OpbJpb6CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *OpbJpb6CreateDto) GetObjekPajakBangunan() (*CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return nil, err
	}
	return &data, nil
}
