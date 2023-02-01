package regobjekpajakbangunan

import (
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type RegJpb15 struct {
	Id               uint64     `json:"id" gorm:"primarykey;autoIncrement"`
	PstPermohonan_id *uint64    `json:"pstPermohonan_id"`
	NoBangunan       *int       `json:"noBangunan"`
	LetakTanki       LetakTanki `json:"letakTanki" gorm:"type:char(1)"`
	KapasitasTanki   *int       `json:"kapasitasTanki"`
	gh.DateModel
}

type RegJpb15CreateDto struct {
	PstPermohonan_id *uint64    `json:"pstPermohonan_id"`
	NoBangunan       *int       `json:"noBangunan"`
	LetakTanki       LetakTanki `json:"letakTanki"`
	KapasitasTanki   *int       `json:"kapasitasTanki"`
}

type RegJpb15UpdateDto struct {
	PstPermohonan_id *uint64    `json:"pstPermohonan_id"`
	NoBangunan       *int       `json:"noBangunan"`
	LetakTanki       LetakTanki `json:"letakTanki"`
	KapasitasTanki   *int       `json:"kapasitasTanki"`
}

type RegJpb15FilterDto struct {
	PstPermohonan_id *uint64 `json:"pstPermohonan_id"`
	NoBangunan       *int    `json:"noBangunan"`
	KapasitasTanki   *int    `json:"kapasitasTanki"`
	Page             int     `json:"page"`
	PageSize         int     `json:"page_size"`
}

func (input *RegOpbJpb15CreateDto) GetFasilitasBangunan() *mrfb.CreateDto {
	return input.RegFasilitasBangunans
}

func (input *RegOpbJpb15CreateDto) GetNop() *string {
	return input.Nop
}

func (input *RegOpbJpb15CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *RegOpbJpb15CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *RegOpbJpb15CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *RegOpbJpb15CreateDto) GetObjekPajakBangunan() (CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return CreateDto{}, err
	}
	return data, nil
}
