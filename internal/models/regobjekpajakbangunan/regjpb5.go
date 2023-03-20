package regobjekpajakbangunan

import (
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type RegJpb5 struct {
	Id                     uint64         `json:"id" gorm:"primarykey;autoIncrement"`
	PstPermohonan_id       *uint64        `json:"pstPermohonan_id"`
	NoBangunan             *int           `json:"noBangunan"`
	LuasKamarAcCentral     *int           `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int           `json:"luasRuangLainAcCentral"`
	KelasBangunan5         *KelasBangunan `json:"kelasBangunan5" gorm:"type:char(1)"`
	gh.DateModel
}

type RegJpb5CreateDto struct {
	PstPermohonan_id       *uint64 `json:"pstPermohonan_id"`
	NoBangunan             *int    `json:"noBangunan"`
	LuasKamarAcCentral     *int    `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int    `json:"luasRuangLainAcCentral"`
}

type RegJpb5UpdateDto struct {
	PstPermohonan_id       *uint64 `json:"pstPermohonan_id"`
	NoBangunan             *int    `json:"noBangunan"`
	LuasKamarAcCentral     *int    `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int    `json:"luasRuangLainAcCentral"`
}

type RegJpb5FilterDto struct {
	PstPermohonan_id       *uint64 `json:"pstPermohonan_id"`
	NoBangunan             *int    `json:"noBangunan"`
	LuasKamarAcCentral     *int    `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int    `json:"luasRuangLainAcCentral"`
	Page                   int     `json:"page"`
	PageSize               int     `json:"page_size"`
}

func (input *RegOpbJpb5CreateDto) GetFasilitasBangunan() *mrfb.CreateDto {
	return input.RegFasilitasBangunans
}

func (input *RegOpbJpb5CreateDto) GetNop() *string {
	return input.Nop
}

func (input *RegOpbJpb5CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *RegOpbJpb5CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *RegOpbJpb5CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *RegOpbJpb5CreateDto) GetObjekPajakBangunan() (CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return CreateDto{}, err
	}
	return data, nil
}
