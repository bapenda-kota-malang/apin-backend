package espt

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailespthiburan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailespthotel"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptparkir"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptppjnonpln"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptppjpln"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptresto"
)

type CreateDetailAirDto struct {
	Espt        CreateDto               `json:"espt" validate:"required"`
	DataDetails detailesptair.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailAirDto) GetEspt() CreateDto {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input *CreateDetailAirDto) CalculateTax(taxPercentage *float64) {
	input.Espt.JumlahPajak = float32(input.Espt.Omset * (*taxPercentage / 100) * float64(input.DataDetails.Pengenaan))
}

func (input CreateDetailAirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailAirDto) LenDetails() int {
	return 1
}

func (input *CreateDetailAirDto) ChangeDetails(newData interface{}) {
	input.DataDetails.Pengenaan = newData.(detailesptair.CreateDto).Pengenaan
	input.DataDetails.JenisAbt = newData.(detailesptair.CreateDto).JenisAbt
}

func (input *CreateDetailAirDto) ReplaceEsptId(id uint) {
	input.DataDetails.Espt_Id = id
}

func (input *CreateDetailAirDto) ReplaceTarifPajakId(id uint) {
	input.Espt.TarifPajak_Id = id
}

type CreateDetailHotelDto struct {
	Espt        CreateDto                   `json:"espt" validate:"required"`
	DataDetails []detailespthotel.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailHotelDto) GetEspt() CreateDto {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input *CreateDetailHotelDto) CalculateTax(taxPercentage *float64) {
	input.Espt.JumlahPajak = float32(input.Espt.Omset * *taxPercentage / 100)
}

func (input CreateDetailHotelDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailHotelDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input CreateDetailHotelDto) ChangeDetails(newData interface{}) {
	// TODO: ADD PROCESS
}

func (input CreateDetailHotelDto) ReplaceEsptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Espt_Id = id
	}
}

func (input *CreateDetailHotelDto) ReplaceTarifPajakId(id uint) {
	input.Espt.TarifPajak_Id = id
}

type CreateDetailHiburanDto struct {
	Espt        CreateDto                   `json:"espt" validate:"required"`
	DataDetails detailespthiburan.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailHiburanDto) GetEspt() CreateDto {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input *CreateDetailHiburanDto) CalculateTax(taxPercentage *float64) {
	input.Espt.JumlahPajak = float32(input.Espt.Omset * *taxPercentage / 100)
}

func (input CreateDetailHiburanDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailHiburanDto) LenDetails() int {
	return 1
}

func (input CreateDetailHiburanDto) ChangeDetails(newData interface{}) {
	// TODO: ADD PROCESS
}

func (input *CreateDetailHiburanDto) ReplaceEsptId(id uint) {
	input.DataDetails.Espt_Id = id
}

func (input *CreateDetailHiburanDto) ReplaceTarifPajakId(id uint) {
	input.Espt.TarifPajak_Id = id
}

type CreateDetailParkirDto struct {
	Espt        CreateDto                    `json:"espt" validate:"required"`
	DataDetails []detailesptparkir.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailParkirDto) GetEspt() CreateDto {
	return input.Espt
}

func (input *CreateDetailParkirDto) CalculateTax(taxPercentage *float64) {
	input.Espt.JumlahPajak = float32(input.Espt.Omset * *taxPercentage / 100)
}

func (input CreateDetailParkirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailParkirDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input CreateDetailParkirDto) ChangeDetails(newData interface{}) {
	// TODO: ADD PROCESS
}

func (input CreateDetailParkirDto) ReplaceEsptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Espt_Id = id
	}
}

func (input *CreateDetailParkirDto) ReplaceTarifPajakId(id uint) {
	input.Espt.TarifPajak_Id = id
}

type CreateDetailRestoDto struct {
	Espt        CreateDto                 `json:"espt" validate:"required"`
	DataDetails detailesptresto.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailRestoDto) GetEspt() CreateDto {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input *CreateDetailRestoDto) CalculateTax(taxPercentage *float64) {
	input.Espt.JumlahPajak = float32(input.Espt.Omset * *taxPercentage / 100)
}

func (input CreateDetailRestoDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailRestoDto) LenDetails() int {
	return 1
}

func (input CreateDetailRestoDto) ChangeDetails(newData interface{}) {
	// TODO: ADD PROCESS
}

func (input *CreateDetailRestoDto) ReplaceEsptId(id uint) {
	input.DataDetails.Espt_Id = id
}

func (input *CreateDetailRestoDto) ReplaceTarifPajakId(id uint) {
	input.Espt.TarifPajak_Id = id
}

type CreateDetailPpjNonPlnDto struct {
	Espt        CreateDto                     `json:"espt" validate:"required"`
	DataDetails detailesptppjnonpln.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailPpjNonPlnDto) GetEspt() CreateDto {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input *CreateDetailPpjNonPlnDto) CalculateTax(taxPercentage *float64) {
	input.Espt.JumlahPajak = float32(input.Espt.Omset * *taxPercentage / 100)
}

func (input CreateDetailPpjNonPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailPpjNonPlnDto) LenDetails() int {
	return 1
}

func (input CreateDetailPpjNonPlnDto) ChangeDetails(newData interface{}) {
	// TODO: ADD PROCESS
}

func (input *CreateDetailPpjNonPlnDto) ReplaceEsptId(id uint) {
	input.DataDetails.Espt_Id = id
}

func (input *CreateDetailPpjNonPlnDto) ReplaceTarifPajakId(id uint) {
	input.Espt.TarifPajak_Id = id
}

type CreateDetailPpjPlnDto struct {
	Espt        CreateDto                    `json:"espt" validate:"required"`
	DataDetails []detailesptppjpln.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailPpjPlnDto) GetEspt() CreateDto {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input *CreateDetailPpjPlnDto) CalculateTax(taxPercentage *float64) {
	tax := float32(0)
	for v := range input.DataDetails {
		if input.DataDetails[v].JenisPPJ.Jenis == "LAIN LAIN" {
			continue
		}
		jumlah := float32(input.DataDetails[v].JumlahRekening)
		jumlah = jumlah * (input.DataDetails[v].JenisPPJ.TarifPersen / 100)
		tax += jumlah
	}
	input.Espt.JumlahPajak = tax + float32(input.Espt.Omset)
}

func (input CreateDetailPpjPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailPpjPlnDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input *CreateDetailPpjPlnDto) ChangeDetails(newData interface{}) {
	input.DataDetails = newData.([]detailesptppjpln.CreateDto)
}

func (input CreateDetailPpjPlnDto) ReplaceEsptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Espt_Id = id
	}
}

func (input *CreateDetailPpjPlnDto) ReplaceTarifPajakId(id uint) {
	input.Espt.TarifPajak_Id = id
}

type UpdateDetailAirDto struct {
	Espt        UpdateDto               `json:"espt"`
	DataDetails detailesptair.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input UpdateDetailAirDto) GetEspt() UpdateDto {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input UpdateDetailAirDto) CalculateTax(taxPercentage *float64) {
	calc := float32(*input.Espt.Omset * *taxPercentage / 100)
	input.Espt.JumlahPajak = &calc
}

func (input UpdateDetailAirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input UpdateDetailAirDto) LenDetails() int {
	return 1
}

type UpdateDetailHotelDto struct {
	Espt        UpdateDto                   `json:"espt" validate:"required"`
	DataDetails []detailespthotel.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input UpdateDetailHotelDto) GetEspt() UpdateDto {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input UpdateDetailHotelDto) CalculateTax(taxPercentage *float64) {
	calc := float32(*input.Espt.Omset * *taxPercentage / 100)
	input.Espt.JumlahPajak = &calc
}

func (input UpdateDetailHotelDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input UpdateDetailHotelDto) LenDetails() int {
	return len(input.DataDetails)
}

type UpdateDetailHiburanDto struct {
	Espt        UpdateDto                     `json:"espt" validate:"required"`
	DataDetails []detailespthiburan.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input UpdateDetailHiburanDto) GetEspt() UpdateDto {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input UpdateDetailHiburanDto) CalculateTax(taxPercentage *float64) {
	calc := float32(*input.Espt.Omset * *taxPercentage / 100)
	input.Espt.JumlahPajak = &calc
}

func (input UpdateDetailHiburanDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input UpdateDetailHiburanDto) LenDetails() int {
	return len(input.DataDetails)
}

type UpdateDetailParkirDto struct {
	Espt        UpdateDto                    `json:"espt" validate:"required"`
	DataDetails []detailesptparkir.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input UpdateDetailParkirDto) GetEspt() UpdateDto {
	return input.Espt
}

func (input UpdateDetailParkirDto) CalculateTax(taxPercentage *float64) {
	calc := float32(*input.Espt.Omset * *taxPercentage / 100)
	input.Espt.JumlahPajak = &calc
}

func (input UpdateDetailParkirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input UpdateDetailParkirDto) LenDetails() int {
	return len(input.DataDetails)
}

type UpdateDetailRestoDto struct {
	Espt        UpdateDto                 `json:"espt" validate:"required"`
	DataDetails detailesptresto.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input UpdateDetailRestoDto) GetEspt() UpdateDto {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input UpdateDetailRestoDto) CalculateTax(taxPercentage *float64) {
	calc := float32(*input.Espt.Omset * *taxPercentage / 100)
	input.Espt.JumlahPajak = &calc
}

func (input UpdateDetailRestoDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input UpdateDetailRestoDto) LenDetails() int {
	return 1
}

type UpdateDetailPpjNonPlnDto struct {
	Espt        UpdateDto                     `json:"espt" validate:"required"`
	DataDetails detailesptppjnonpln.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input UpdateDetailPpjNonPlnDto) GetEspt() UpdateDto {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input UpdateDetailPpjNonPlnDto) CalculateTax(taxPercentage *float64) {
	calc := float32(*input.Espt.Omset * *taxPercentage / 100)
	input.Espt.JumlahPajak = &calc
}

func (input UpdateDetailPpjNonPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input UpdateDetailPpjNonPlnDto) LenDetails() int {
	return 1
}

type UpdateDetailPpjPlnDto struct {
	Espt        UpdateDto                    `json:"espt" validate:"required"`
	DataDetails []detailesptppjpln.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input UpdateDetailPpjPlnDto) GetEspt() UpdateDto {
	input.Espt.LuasLokasi = nil
	return input.Espt
}

func (input UpdateDetailPpjPlnDto) CalculateTax(taxPercentage *float64) {
	calc := float32(*input.Espt.Omset * *taxPercentage / 100)
	input.Espt.JumlahPajak = &calc
}

func (input UpdateDetailPpjPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input UpdateDetailPpjPlnDto) LenDetails() int {
	return len(input.DataDetails)
}
