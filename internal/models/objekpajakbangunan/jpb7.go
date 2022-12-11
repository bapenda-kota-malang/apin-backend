package objekpajakbangunan

import (
	mfb "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type Jpb7 struct {
	nop.NopDetail
	NoBangunan             *int          `json:"noBangunan"`
	JenisHotel             JenisHotel    `json:"jenisHotel" gorm:"type:char(1)"`
	JumlahBintang          JumlahBintang `json:"jumlahBintang" gorm:"type:char(1)"`
	JumlahKamar            *int          `json:"jumlahKamar"`
	LuasKamarAcCentral     *int          `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
	gh.DateModel
}

type Jpb7CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan             *int          `json:"noBangunan"`
	JenisHotel             JenisHotel    `json:"jenisHotel"`
	JumlahBintang          JumlahBintang `json:"jumlahBintang"`
	JumlahKamar            *int          `json:"jumlahKamar"`
	LuasKamarAcCentral     *int          `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
}

type Jpb7UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan             *int          `json:"noBangunan"`
	JenisHotel             JenisHotel    `json:"jenisHotel"`
	JumlahBintang          JumlahBintang `json:"jumlahBintang"`
	JumlahKamar            *int          `json:"jumlahKamar"`
	LuasKamarAcCentral     *int          `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int          `json:"luasRuangLainAcCentral"`
}

type Jpb7FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan             *int `json:"noBangunan"`
	JumlahKamar            *int `json:"jumlahKamar"`
	LuasKamarAcCentral     *int `json:"luasKamarAcCentral"`
	LuasRuangLainAcCentral *int `json:"luasRuangLainAcCentral"`
	Page                   int  `json:"page"`
	PageSize               int  `json:"page_size"`
}

func (input *OpbJpb7CreateDto) GetFasilitasBangunan() *mfb.CreateDto {
	return input.FasilitasBangunans
}

func (input *OpbJpb7CreateDto) GetNop() *string {
	return input.Nop
}

func (input *OpbJpb7CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *OpbJpb7CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *OpbJpb7CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *OpbJpb7CreateDto) GetObjekPajakBangunan() (*CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return nil, err
	}
	return &data, nil
}
