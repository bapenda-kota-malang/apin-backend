package potensiopwp

import (
	"errors"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	"github.com/google/uuid"
	"gorm.io/gorm"

	sc "github.com/jinzhu/copier"
)

type DPPPJNonPLNUsecase struct {
	DPBaseUsecase
	DetailPajakDtos m.CreateDtoDPPPJNonPLN
}

func NewDPPPJNonPLNUsecase(baseDto DPBaseUsecase, detailPajakDto m.CreateDtoDPPPJNonPLN) m.Input {
	return &DPPPJNonPLNUsecase{baseDto, detailPajakDto}
}

func (s *DPPPJNonPLNUsecase) ReplacePotensiOpId(id uuid.UUID) {
	s.DetailPotensiOp.Potensiop_Id = id
	s.Bapl.Potensiop_Id = id
	for v := range s.PotensiPemilikWps {
		s.PotensiPemilikWps[v].Potensiop_Id = id
	}
	for v := range s.PotensiNarahubungs {
		s.PotensiNarahubungs[v].Potensiop_Id = id
	}
	s.DetailPajakDtos.Potensiop_Id = id
}

func (s *DPPPJNonPLNUsecase) GetDetailPotensiPajak() interface{} {
	return s.DetailPajakDtos
}

func (s *DPPPJNonPLNUsecase) SaveDetailPotensiPajak(tx *gorm.DB) error {
	if tx == nil {
		return errors.New("tx empty, init db connection or use transaction")
	}
	var data []m.DetailPotensiPPJNonPLN
	if err := sc.Copy(&data, &s.DetailPajakDtos); err != nil {
		return errors.New("failed copy data to destination model detail potensi PPJNonPLNs")
	}
	if result := tx.Create(&data); result.Error != nil {
		return result.Error
	}
	return nil
}
