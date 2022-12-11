package objekpajakbangunan

import (
	mfb "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type Jpb15 struct {
	nop.NopDetail
	NoBangunan     *int       `json:"noBangunan"`
	LetakTanki     LetakTanki `json:"letakTanki" gorm:"type:char(1)"`
	KapasitasTanki *int       `json:"kapasitasTanki"`
	gh.DateModel
}

type Jpb15CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan     *int       `json:"noBangunan"`
	LetakTanki     LetakTanki `json:"letakTanki"`
	KapasitasTanki *int       `json:"kapasitasTanki"`
}

type Jpb15UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan     *int       `json:"noBangunan"`
	LetakTanki     LetakTanki `json:"letakTanki"`
	KapasitasTanki *int       `json:"kapasitasTanki"`
}

type Jpb15FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan     *int `json:"noBangunan"`
	KapasitasTanki *int `json:"kapasitasTanki"`
	Page           int  `json:"page"`
	PageSize       int  `json:"page_size"`
}

func (input *OpbJpb15CreateDto) GetFasilitasBangunan() *mfb.CreateDto {
	return input.FasilitasBangunans
}

func (input *OpbJpb15CreateDto) GetNop() *string {
	return input.Nop
}

func (input *OpbJpb15CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *OpbJpb15CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *OpbJpb15CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *OpbJpb15CreateDto) GetObjekPajakBangunan() (*CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return nil, err
	}
	return &data, nil
}
