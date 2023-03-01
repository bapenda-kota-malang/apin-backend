package potensiopwp

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/bapl"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailobjek"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiairtanah"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensihotel"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensinarahubung"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensipemilikwp"
)

type CreateDto struct {
	PotensiOp       CreatePotensiOpDto        `json:"potensiOp" validate:"required"`
	DetailPotensiOp detailpotensiop.CreateDto `json:"detailPotensiOp" validate:"required"`
	// DetailPajakDtos    []detailobjek.DetailPajakDto  `json:"detailPajaks"`
	PotensiPemilikWps  []potensipemilikwp.CreateDto  `json:"potensiPemilikWps" validate:"required"`
	PotensiNarahubungs []potensinarahubung.CreateDto `json:"potensiNarahubungs"`
	Bapl               bapl.CreateDto                `json:"bapl" validate:"required"`
}

type UpdateDto struct {
	PotensiOp          UpdatePotensiOpDto            `json:"potensiOp"`
	DetailPotensiOp    detailpotensiop.UpdateDto     `json:"detailPotensiOp"`
	DetailPajakDtos    []detailobjek.UpdateDto       `json:"detailPajaks"`
	PotensiPemilikWps  []potensipemilikwp.UpdateDto  `json:"potensiPemilikWps"`
	PotensiNarahubungs []potensinarahubung.UpdateDto `json:"potensiNarahubungs"`
	Bapl               bapl.UpdateDto                `json:"bapl"`
}

type CreateDtoAirTanah struct {
	CreateDto
	DetailPajakDtos []detailpotensiairtanah.CreateDto `json:"detailPajaks"`
}

type CreateDtoHotel struct {
	CreateDto
	DetailPajakDtos []detailpotensihotel.CreateDto `json:"detailPajaks"`
}
