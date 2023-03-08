package potensiopwp

import (
	"errors"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	"github.com/google/uuid"
	"gorm.io/gorm"

	sc "github.com/jinzhu/copier"
)

type DPHotelUsecase struct {
	DPBaseUsecase
	DetailPajakDtos []m.CreateDtoDPHotel
}

func NewDPHotelUsecase(baseDto DPBaseUsecase, detailPajakDto []m.CreateDtoDPHotel) m.Input {
	return &DPHotelUsecase{baseDto, detailPajakDto}
}

func (s *DPHotelUsecase) ReplacePotensiOpId(id uuid.UUID) {
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

func (s *DPHotelUsecase) GetDetailPotensiPajak() interface{} {
	return s.DetailPajakDtos
}

func (s *DPHotelUsecase) SaveDetailPotensiPajak(tx *gorm.DB) error {
	if tx == nil {
		return errors.New("tx empty, init db connection or use transaction")
	}
	var data []m.DetailPotensiHotel
	if err := sc.Copy(&data, &s.DetailPajakDtos); err != nil {
		return errors.New("failed copy data to destination model detail potensi hotels")
	}
	if result := tx.Create(&data); result.Error != nil {
		return result.Error
	}
	return nil
}
