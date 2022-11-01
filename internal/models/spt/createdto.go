package spt

import (
	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mdsa "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptair"
	mdsh "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthotel"
	mdsp "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	mdsrek "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	mdsres "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"
	mdsjbr "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/jaminanbongkarreklame"
	"github.com/jinzhu/copier"
)

type CreateDetailAirDto struct {
	Spt         CreateDto      `json:"spt" validate:"required"`
	DataDetails mdsa.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailAirDto) GetSpt() interface{} {
	input.Spt.LuasLokasi = nil
	return input.Spt
}

func (input CreateDetailAirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailAirDto) LenDetails() int {
	return 1
}

func (input *CreateDetailAirDto) ReplaceSptId(id uint) {
	input.DataDetails.Spt_Id = id
}

func (input *CreateDetailAirDto) ChangeDetails(newData interface{}) {
	input.DataDetails.Pengenaan = newData.(mdsa.CreateDto).Pengenaan
	input.DataDetails.JenisAbt = newData.(mdsa.CreateDto).JenisAbt
}

func (input *CreateDetailAirDto) ReplaceTarifPajakId(id uint) {
	input.Spt.TarifPajak_Id = id
}

func (input *CreateDetailAirDto) CalculateTax(taxPercentage *float64) {
	input.Spt.JumlahPajak = float32(input.Spt.Omset * (*taxPercentage / 100) * float64(input.DataDetails.Pengenaan))
}

func (input *CreateDetailAirDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	if err := copier.Copy(&input.Spt, &esptDetail); err != nil {
		return err
	}
	input.Spt.Lampiran = esptDetail.Attachment
	input.Spt.CreateBy_User_Id = *esptDetail.VerifyBy_User_Id
	if err := copier.Copy(&input.DataDetails, &esptDetail); err != nil {
		return err
	}
	return nil
}

type CreateDetailHotelDto struct {
	Spt         CreateDto        `json:"spt" validate:"required"`
	DataDetails []mdsh.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailHotelDto) GetSpt() interface{} {
	input.Spt.LuasLokasi = nil
	return input.Spt
}

func (input CreateDetailHotelDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailHotelDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input *CreateDetailHotelDto) ReplaceSptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Spt_Id = id
	}
}

func (input *CreateDetailHotelDto) ChangeDetails(newData interface{}) {
	// Do nothing
}

func (input *CreateDetailHotelDto) ReplaceTarifPajakId(id uint) {
	input.Spt.TarifPajak_Id = id
}

func (input *CreateDetailHotelDto) CalculateTax(taxPercentage *float64) {
	input.Spt.JumlahPajak = float32(input.Spt.Omset * (*taxPercentage / 100))
}

func (input *CreateDetailHotelDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	if err := copier.Copy(&input.Spt, &esptDetail); err != nil {
		return err
	}
	input.Spt.Lampiran = esptDetail.Attachment
	input.Spt.CreateBy_User_Id = *esptDetail.VerifyBy_User_Id
	if err := copier.Copy(&input.DataDetails, &esptDetail); err != nil {
		return err
	}
	return nil
}

type CreateDetailParkirDto struct {
	Spt         CreateDto        `json:"spt" validate:"required"`
	DataDetails []mdsp.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailParkirDto) GetSpt() interface{} {
	input.Spt.LuasLokasi = nil
	return input.Spt
}

func (input CreateDetailParkirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailParkirDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input *CreateDetailParkirDto) ReplaceSptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Spt_Id = id
	}
}

func (input *CreateDetailParkirDto) ChangeDetails(newData interface{}) {
	// Do nothing
}

func (input *CreateDetailParkirDto) ReplaceTarifPajakId(id uint) {
	input.Spt.TarifPajak_Id = id
}

func (input *CreateDetailParkirDto) CalculateTax(taxPercentage *float64) {
	input.Spt.JumlahPajak = float32(input.Spt.Omset * (*taxPercentage / 100))
}

func (input *CreateDetailParkirDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	if err := copier.Copy(&input.Spt, &esptDetail); err != nil {
		return err
	}
	input.Spt.Lampiran = esptDetail.Attachment
	input.Spt.CreateBy_User_Id = *esptDetail.VerifyBy_User_Id
	if err := copier.Copy(&input.DataDetails, &esptDetail); err != nil {
		return err
	}
	return nil
}

type CreateDetailReklameDto struct {
	Spt                   CreateDto        `json:"spt" validate:"required"`
	DataDetails           []mdsh.CreateDto `json:"dataDetails" validate:"required"`
	DetailSptReklame      mdsrek.CreateDto `json:"detailSptReklame"`
	JaminanBongkarReklame mdsjbr.CreateDto `json:"jaminanBongkarReklame"`
}

func (input CreateDetailReklameDto) GetSpt() interface{} {
	input.Spt.LuasLokasi = nil
	return input.Spt
}

func (input CreateDetailReklameDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailReklameDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input *CreateDetailReklameDto) ReplaceSptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Spt_Id = id
	}
}

func (input *CreateDetailReklameDto) ChangeDetails(newData interface{}) {
	// Do nothing
}

func (input *CreateDetailReklameDto) ReplaceTarifPajakId(id uint) {
	input.Spt.TarifPajak_Id = id
}

func (input *CreateDetailReklameDto) CalculateTax(taxPercentage *float64) {
	input.Spt.JumlahPajak = float32(input.Spt.Omset * (*taxPercentage / 100))
}

func (input *CreateDetailReklameDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	// do nothing because espt don't have reklame
	return nil
}

type CreateDetailRestoDto struct {
	Spt         CreateDto        `json:"spt" validate:"required"`
	DataDetails mdsres.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailRestoDto) GetSpt() interface{} {
	input.Spt.LuasLokasi = nil
	return input.Spt
}

func (input CreateDetailRestoDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailRestoDto) LenDetails() int {
	return 1
}

func (input *CreateDetailRestoDto) ReplaceSptId(id uint) {
	input.DataDetails.Spt_Id = id
}

func (input *CreateDetailRestoDto) ChangeDetails(newData interface{}) {
	// Do nothing
}

func (input *CreateDetailRestoDto) ReplaceTarifPajakId(id uint) {
	input.Spt.TarifPajak_Id = id
}

func (input *CreateDetailRestoDto) CalculateTax(taxPercentage *float64) {
	input.Spt.JumlahPajak = float32(input.Spt.Omset * (*taxPercentage / 100))
}

func (input *CreateDetailRestoDto) DuplicateEspt(esptDetail *mespt.Espt) error {
	if err := copier.Copy(&input.Spt, &esptDetail); err != nil {
		return err
	}
	input.Spt.Lampiran = esptDetail.Attachment
	input.Spt.CreateBy_User_Id = *esptDetail.VerifyBy_User_Id
	if err := copier.Copy(&input.DataDetails, &esptDetail); err != nil {
		return err
	}
	return nil
}
