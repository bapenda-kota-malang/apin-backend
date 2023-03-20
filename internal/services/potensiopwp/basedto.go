package potensiopwp

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/bapl"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensinarahubung"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensipemilikwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DPBaseUsecase struct {
	potensiopwp.CreateDto
}

func NewDPBaseUsecase(dto potensiopwp.CreateDto) potensiopwp.Input {
	return &DPBaseUsecase{dto}
}

func (s *DPBaseUsecase) GetPotensiOp() potensiopwp.CreatePotensiOpDto {
	return s.PotensiOp
}

func (s *DPBaseUsecase) SetPotensiOpId(id uuid.UUID) {
	s.PotensiOp.Id = id
}

func (s *DPBaseUsecase) GetDetailPotensiOp() detailpotensiop.CreateDto {
	return s.DetailPotensiOp
}

func (s *DPBaseUsecase) GetPotensiPemilikWps() []potensipemilikwp.CreateDto {
	return s.PotensiPemilikWps
}

func (s *DPBaseUsecase) GetPotensiNarahubungs() []potensinarahubung.CreateDto {
	return s.PotensiNarahubungs
}

func (s *DPBaseUsecase) GetBapl() bapl.CreateDto {
	return s.Bapl
}

func (s *DPBaseUsecase) ReplacePotensiOpId(id uuid.UUID) {
	s.DetailPotensiOp.Potensiop_Id = id
	s.Bapl.Potensiop_Id = id
	for v := range s.PotensiPemilikWps {
		s.PotensiPemilikWps[v].Potensiop_Id = id
	}
	for v := range s.PotensiNarahubungs {
		s.PotensiNarahubungs[v].Potensiop_Id = id
	}
}

func (s *DPBaseUsecase) GetDetailPotensiPajak() interface{} {
	return nil
}

func (s *DPBaseUsecase) SaveDetailPotensiPajak(tx *gorm.DB) error {
	return nil
}

func (s *DPBaseUsecase) CalculateTax(tarifPajak *tarifpajak.TarifPajak) (potensiopwp.CreatePotensiOpDto, error) {
	s.PotensiOp.TarifPajak_Id = &tarifPajak.Id
	s.PotensiOp.JumlahPajak = *s.PotensiOp.OmsetOp * (*tarifPajak.TarifPersen / 100)
	return s.PotensiOp, nil
}
