package potensiopwp

import (
	"errors"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	mtp "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"
	"github.com/google/uuid"
	"gorm.io/gorm"

	sc "github.com/jinzhu/copier"
)

type DPReklameUsecase struct {
	DPBaseUsecase
	DetailPajakDtos []m.CreateDtoDPReklame
}

func NewDPReklameUsecase(baseDto DPBaseUsecase, detailPajakDto []m.CreateDtoDPReklame) m.Input {
	return &DPReklameUsecase{baseDto, detailPajakDto}
}

func (s *DPReklameUsecase) ReplacePotensiOpId(id uuid.UUID) {
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

func (s *DPReklameUsecase) GetDetailPotensiPajak() interface{} {
	return s.DetailPajakDtos
}

func (s *DPReklameUsecase) SaveDetailPotensiPajak(tx *gorm.DB) error {
	if tx == nil {
		return errors.New("tx empty, init db connection or use transaction")
	}
	var data []m.DetailPotensiReklame
	if err := sc.Copy(&data, &s.DetailPajakDtos); err != nil {
		return errors.New("failed copy data to destination model detail potensi Reklames")
	}
	if result := tx.Create(&data); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *DPReklameUsecase) CalculateTax(tarifPajak mtp.TarifPajak) error {
	// TODO: IMPLEMENT CALC PROCESS
	return nil
}
