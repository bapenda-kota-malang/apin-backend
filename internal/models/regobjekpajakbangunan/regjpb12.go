package regobjekpajakbangunan

import (
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type RegJpb12 struct {
	Id               uint64       `json:"id" gorm:"primarykey;autoIncrement"`
	PstPermohonan_id *uint64      `json:"pstPermohonan_id"`
	NoBangunan       int          `json:"noBangunan"`
	TipeBangunan     TipeBangunan `json:"tipeBangunan" gorm:"type:char(1)"`
	gh.DateModel
}

type RegJpb12CreateDto struct {
	PstPermohonan_id *uint64      `json:"pstPermohonan_id"`
	NoBangunan       *int         `json:"noBangunan"`
	TipeBangunan     TipeBangunan `json:"tipeBangunan"`
}

type RegJpb12UpdateDto struct {
	PstPermohonan_id *uint64      `json:"pstPermohonan_id"`
	NoBangunan       *int         `json:"noBangunan"`
	TipeBangunan     TipeBangunan `json:"tipeBangunan"`
}

type RegJpb12FilterDto struct {
	PstPermohonan_id *uint64      `json:"pstPermohonan_id"`
	NoBangunan       *int         `json:"noBangunan"`
	TipeBangunan     TipeBangunan `json:"tipeBangunan"`
	Page             int          `json:"page"`
	PageSize         int          `json:"page_size"`
}

func (input *RegOpbJpb12CreateDto) GetFasilitasBangunan() *mrfb.CreateDto {
	return input.RegFasilitasBangunans
}

func (input *RegOpbJpb12CreateDto) GetNop() *string {
	return input.Nop
}

func (input *RegOpbJpb12CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *RegOpbJpb12CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *RegOpbJpb12CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *RegOpbJpb12CreateDto) GetObjekPajakBangunan() (CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return CreateDto{}, err
	}
	return data, nil
}
