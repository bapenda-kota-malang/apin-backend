package potensiopwp

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/bapl"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensinarahubung"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensipemilikwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Input interface {
	GetPotensiOp() CreatePotensiOpDto
	SetPotensiOpId(id uuid.UUID)
	GetDetailPotensiOp() detailpotensiop.CreateDto
	GetPotensiPemilikWps() []potensipemilikwp.CreateDto
	GetPotensiNarahubungs() []potensinarahubung.CreateDto
	GetBapl() bapl.CreateDto
	ReplacePotensiOpId(id uuid.UUID)
	GetDetailPotensiPajak() interface{}
	SaveDetailPotensiPajak(tx *gorm.DB) error
	CalculateTax(tarifPajak *tarifpajak.TarifPajak) (CreatePotensiOpDto, error)
}

type CreateDto struct {
	PotensiOp          CreatePotensiOpDto            `json:"potensiOp" validate:"required"`
	DetailPotensiOp    detailpotensiop.CreateDto     `json:"detailPotensiOp" validate:"required"`
	PotensiPemilikWps  []potensipemilikwp.CreateDto  `json:"potensiPemilikWps" validate:"required"`
	PotensiNarahubungs []potensinarahubung.CreateDto `json:"potensiNarahubungs"`
	Bapl               bapl.CreateDto                `json:"bapl" validate:"required"`
}

type UpdateDto struct {
	PotensiOp          UpdatePotensiOpDto            `json:"potensiOp"`
	DetailPotensiOp    detailpotensiop.UpdateDto     `json:"detailPotensiOp"`
	PotensiPemilikWps  []potensipemilikwp.UpdateDto  `json:"potensiPemilikWps"`
	PotensiNarahubungs []potensinarahubung.UpdateDto `json:"potensiNarahubungs"`
	Bapl               bapl.UpdateDto                `json:"bapl"`
}
