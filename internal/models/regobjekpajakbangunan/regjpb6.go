package regobjekpajakbangunan

import (
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type RegJpb6 struct {
	Id               uint64  `json:"id" gorm:"primarykey;autoIncrement"`
	PstPermohonan_id *uint64 `json:"pstPermohonan_id"`
	NoBangunan       *int    `json:"noBangunan"`
	gh.DateModel
}

type RegJpb6CreateDto struct {
	PstPermohonan_id *uint64 `json:"pstPermohonan_id"`
	NoBangunan       *int    `json:"noBangunan"`
}

type RegJpb6UpdateDto struct {
	PstPermohonan_id *uint64 `json:"pstPermohonan_id"`
	NoBangunan       *int    `json:"noBangunan"`
}

type RegJpb6FilterDto struct {
	PstPermohonan_id *uint64 `json:"pstPermohonan_id"`
	NoBangunan       *int    `json:"noBangunan"`
	Page             int     `json:"page"`
	PageSize         int     `json:"page_size"`
}

func (input *RegOpbJpb6CreateDto) GetFasilitasBangunan() *mrfb.CreateDto {
	return input.RegFasilitasBangunans
}

func (input *RegOpbJpb6CreateDto) GetNop() *string {
	return input.Nop
}

func (input *RegOpbJpb6CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *RegOpbJpb6CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *RegOpbJpb6CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *RegOpbJpb6CreateDto) GetObjekPajakBangunan() (CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return CreateDto{}, err
	}
	return data, nil
}
