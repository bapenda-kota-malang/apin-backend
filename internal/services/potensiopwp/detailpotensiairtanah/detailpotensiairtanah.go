package detailpotensiairtanah

import (
	"errors"

	mpotensiopwp "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiairtanah"
	spotensiopwp "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp"
	"github.com/google/uuid"
	"gorm.io/gorm"

	sc "github.com/jinzhu/copier"
)

type DPAirTanahUsecase struct {
	spotensiopwp.DPBaseUsecase
	DetailPajakDtos []m.CreateDto
}

func NewDPAirTanahUsecase(baseDto spotensiopwp.DPBaseUsecase, detailPajakDto []m.CreateDto) mpotensiopwp.Input {
	return &DPAirTanahUsecase{baseDto, detailPajakDto}
}

func (s *DPAirTanahUsecase) ReplacePotensiOpId(id uuid.UUID) {
	s.DetailPotensiOp.Potensiop_Id = id
	s.Bapl.Potensiop_Id = id
	for v := range s.PotensiPemilikWps {
		s.PotensiPemilikWps[v].Potensiop_Id = id
	}
	for v := range s.PotensiNarahubungs {
		s.PotensiNarahubungs[v].Potensiop_Id = id
	}
	for v := range s.DetailPajakDtos {
		s.DetailPajakDtos[v].Potensiop_Id = id
	}
}

func (s *DPAirTanahUsecase) GetDetailPotensiPajak() interface{} {
	return s.DetailPajakDtos
}

func (s *DPAirTanahUsecase) SaveDetailPotensiPajak(tx *gorm.DB) error {
	if tx == nil {
		return errors.New("tx empty, init db connection or use transaction")
	}
	var data []m.DetailPotensiAirTanah
	if err := sc.Copy(&data, &s.DetailPajakDtos); err != nil {
		return errors.New("failed copy data to destination model detail potensi air tanah")
	}
	if result := tx.Create(&data); result.Error != nil {
		return result.Error
	}
	return nil
}
