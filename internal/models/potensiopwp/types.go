package potensiopwp

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/bapl"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensinarahubung"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensipemilikwp"
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
}
