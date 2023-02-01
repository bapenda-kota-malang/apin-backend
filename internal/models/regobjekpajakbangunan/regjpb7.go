package regobjekpajakbangunan

import (
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type RegJpb7 struct {
	Id                     uint64        `json:"id" gorm:"primarykey;autoIncrement"`
	PstPermohonan_id       *uint64       `json:"pstPermohonan_id"`
	NoBangunan             *int          `json:"noBangunan"`
	JenisHotel             JenisHotel    `json:"jenisHotel" gorm:"type:char(1)"`
	JumlahBintang          JumlahBintang `json:"jumlahBintang" gorm:"type:char(1)"`
	JumlahKamar            *int          `json:"jumlahKamar"`
	LuasKamarAcCentral     *int          `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
	gh.DateModel
}

type RegJpb7CreateDto struct {
	PstPermohonan_id       *uint64       `json:"pstPermohonan_id"`
	NoBangunan             *int          `json:"noBangunan"`
	JenisHotel             JenisHotel    `json:"jenisHotel"`
	JumlahBintang          JumlahBintang `json:"jumlahBintang"`
	JumlahKamar            *int          `json:"jumlahKamar"`
	LuasKamarAcCentral     *int          `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
}

type RegJpb7UpdateDto struct {
	PstPermohonan_id       *uint64       `json:"pstPermohonan_id"`
	NoBangunan             *int          `json:"noBangunan"`
	JenisHotel             JenisHotel    `json:"jenisHotel"`
	JumlahBintang          JumlahBintang `json:"jumlahBintang"`
	JumlahKamar            *int          `json:"jumlahKamar"`
	LuasKamarAcCentral     *int          `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
}

type RegJpb7FilterDto struct {
	PstPermohonan_id       *uint64 `json:"pstPermohonan_id"`
	NoBangunan             *int    `json:"noBangunan"`
	JumlahKamar            *int    `json:"jumlahKamar"`
	LuasKamarAcCentral     *int    `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int    `json:"luasRuangLainAcCentral"`
	Page                   int     `json:"page"`
	PageSize               int     `json:"page_size"`
}

func (input *RegOpbJpb7CreateDto) GetFasilitasBangunan() *mrfb.CreateDto {
	return input.RegFasilitasBangunans
}

func (input *RegOpbJpb7CreateDto) GetNop() *string {
	return input.Nop
}

func (input *RegOpbJpb7CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *RegOpbJpb7CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *RegOpbJpb7CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *RegOpbJpb7CreateDto) GetObjekPajakBangunan() (CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return CreateDto{}, err
	}
	return data, nil
}
