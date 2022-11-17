package potensiopwp

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailobjek"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensinarahubung"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensipemilikwp"
)

type CreateDto struct {
	PotensiOp          CreatePotensiOpDto            `json:"potensiOp" validate:"required"`
	DetailPotensiOp    detailpotensiop.CreateDto     `json:"detailPotensiOp" validate:"required"`
	DetailPajakDtos    []detailobjek.DetailPajakDto  `json:"detailPajaks"`
	PotensiPemilikWps  []potensipemilikwp.CreateDto  `json:"potensiPemilikWps" validate:"required"`
	PotensiNarahubungs []potensinarahubung.CreateDto `json:"potensiNarahubungs"`
}

type UpdateDto struct {
	PotensiOp          UpdatePotensiOpDto            `json:"potensiOp" validate:"required"`
	DetailPotensiOp    detailpotensiop.UpdateDto     `json:"detailPotensiOp" validate:"required"`
	DetailPajakDtos    []detailobjek.DetailPajakDto  `json:"detailPajaks"`
	PotensiPemilikWps  []potensipemilikwp.UpdateDto  `json:"potensiPemilikWps" validate:"required"`
	PotensiNarahubungs []potensinarahubung.UpdateDto `json:"potensiNarahubungs"`
}
