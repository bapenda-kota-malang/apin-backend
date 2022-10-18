package espt

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthiburan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthotel"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptparkir"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptppjnonpln"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptresto"
)

type CreateDetailAirDto struct {
	Espt        CreateDto                 `json:"espt"`
	DataDetails []detailesptair.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailAirDto) GetEspt() CreateDto {
	return input.Espt
}

func (input CreateDetailAirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailAirDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input CreateDetailAirDto) ReplaceEsptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Espt_Id = id
	}
}

type CreateDetailHotelDto struct {
	Espt        CreateDto                   `json:"espt" validate:"required"`
	DataDetails []detailespthotel.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailHotelDto) GetEspt() CreateDto {
	return input.Espt
}

func (input CreateDetailHotelDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailHotelDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input CreateDetailHotelDto) ReplaceEsptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Espt_Id = &id
	}
}

type CreateDetailHiburanDto struct {
	Espt        CreateDto                     `json:"espt" validate:"required"`
	DataDetails []detailespthiburan.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailHiburanDto) GetEspt() CreateDto {
	return input.Espt
}

func (input CreateDetailHiburanDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailHiburanDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input CreateDetailHiburanDto) ReplaceEsptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Espt_Id = &id
	}
}

type CreateDetailParkirDto struct {
	Espt        CreateDto                    `json:"espt" validate:"required"`
	DataDetails []detailesptparkir.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailParkirDto) GetEspt() CreateDto {
	return input.Espt
}

func (input CreateDetailParkirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailParkirDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input CreateDetailParkirDto) ReplaceEsptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Espt_Id = id
	}
}

type CreateDetailRestoDto struct {
	Espt        CreateDto                   `json:"espt" validate:"required"`
	DataDetails []detailesptresto.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailRestoDto) GetEspt() CreateDto {
	return input.Espt
}

func (input CreateDetailRestoDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailRestoDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input CreateDetailRestoDto) ReplaceEsptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Espt_Id = id
	}
}

type CreateDetailPpjNonPlnDto struct {
	Espt        CreateDto                       `json:"espt" validate:"required"`
	DataDetails []detailesptppjnonpln.CreateDto `json:"dataDetails" validate:"required"`
}

func (input CreateDetailPpjNonPlnDto) GetEspt() CreateDto {
	return input.Espt
}

func (input CreateDetailPpjNonPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input CreateDetailPpjNonPlnDto) LenDetails() int {
	return len(input.DataDetails)
}

func (input CreateDetailPpjNonPlnDto) ReplaceEsptId(id uint) {
	for i := range input.DataDetails {
		input.DataDetails[i].Espt_Id = id
	}
}

type UpdateDetailAirDto struct {
	Espt        UpdateDto                 `json:"espt"`
	DataDetails []detailesptair.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input UpdateDetailAirDto) GetEspt() UpdateDto {
	return input.Espt
}

func (input UpdateDetailAirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input UpdateDetailAirDto) LenDetails() int {
	return len(input.DataDetails)
}

type UpdateDetailHotelDto struct {
	Espt        UpdateDto                   `json:"espt" validate:"required"`
	DataDetails []detailespthotel.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input UpdateDetailHotelDto) GetEspt() UpdateDto {
	return input.Espt
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
	return input.Espt
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

func (input UpdateDetailParkirDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input UpdateDetailParkirDto) LenDetails() int {
	return len(input.DataDetails)
}

type UpdateDetailRestoDto struct {
	Espt        UpdateDto                   `json:"espt" validate:"required"`
	DataDetails []detailesptresto.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input UpdateDetailRestoDto) GetEspt() UpdateDto {
	return input.Espt
}

func (input UpdateDetailRestoDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input UpdateDetailRestoDto) LenDetails() int {
	return len(input.DataDetails)
}

type UpdateDetailPpjNonPlnDto struct {
	Espt        UpdateDto                       `json:"espt" validate:"required"`
	DataDetails []detailesptppjnonpln.UpdateDto `json:"dataDetails" validate:"required"`
}

func (input UpdateDetailPpjNonPlnDto) GetEspt() UpdateDto {
	return input.Espt
}

func (input UpdateDetailPpjNonPlnDto) GetDetails() interface{} {
	return input.DataDetails
}

func (input UpdateDetailPpjNonPlnDto) LenDetails() int {
	return len(input.DataDetails)
}
