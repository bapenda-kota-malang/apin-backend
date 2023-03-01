package potensiopwp

import (
	"errors"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	"github.com/google/uuid"
	"gorm.io/gorm"

	sc "github.com/jinzhu/copier"
)

type DPHiburanUsecase struct {
	DPBaseUsecase
	DetailPajakDtos []m.CreateDtoDPHiburan
}

func NewDPHiburanUsecase(baseDto DPBaseUsecase, detailPajakDto []m.CreateDtoDPHiburan) m.Input {
	return &DPHiburanUsecase{baseDto, detailPajakDto}
}

func (s *DPHiburanUsecase) ReplacePotensiOpId(id uuid.UUID) {
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

func (s *DPHiburanUsecase) GetDetailPotensiPajak() interface{} {
	return s.DetailPajakDtos
}

func (s *DPHiburanUsecase) SaveDetailPotensiPajak(tx *gorm.DB) error {
	if tx == nil {
		return errors.New("tx empty, init db connection or use transaction")
	}
	var data []m.DetailPotensiHiburan
	if err := sc.Copy(&data, &s.DetailPajakDtos); err != nil {
		return errors.New("failed copy data to destination model detail potensi hiburans")
	}
	if result := tx.Create(&data); result.Error != nil {
		return result.Error
	}
	return nil
}
