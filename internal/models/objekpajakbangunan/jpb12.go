package objekpajakbangunan

import (
	mfb "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type Jpb12 struct {
	nop.NopDetail
	NoBangunan   int          `json:"noBangunan"`
	TipeBangunan TipeBangunan `json:"tipeBangunan" gorm:"type:char(1)"`
	gh.DateModel
}

type Jpb12CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan   *int         `json:"noBangunan"`
	TipeBangunan TipeBangunan `json:"tipeBangunan"`
}

type Jpb12UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan   *int         `json:"noBangunan"`
	TipeBangunan TipeBangunan `json:"tipeBangunan"`
}

type Jpb12FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan   *int         `json:"noBangunan"`
	TipeBangunan TipeBangunan `json:"tipeBangunan"`
	Page         int          `json:"page"`
	PageSize     int          `json:"page_size"`
}

func (input *OpbJpb12CreateDto) GetFasilitasBangunan() *mfb.CreateDto {
	return input.FasilitasBangunans
}

func (input *OpbJpb12CreateDto) GetNop() *string {
	return input.Nop
}

func (input *OpbJpb12CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *OpbJpb12CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *OpbJpb12CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *OpbJpb12CreateDto) GetObjekPajakBangunan() (CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return CreateDto{}, err
	}
	return data, nil
}
