package detailpotensihotel

import (
	"errors"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/bapl"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensihotel"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensinarahubung"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensipemilikwp"
	"github.com/google/uuid"
	"gorm.io/gorm"

	sc "github.com/jinzhu/copier"
)

type DPHotelUsecase struct {
	potensiopwp.CreateDtoHotel
}

func NewDPHotelUsecase(dto potensiopwp.CreateDtoHotel) potensiopwp.Input {
	return &DPHotelUsecase{dto}
}

func (s *DPHotelUsecase) GetPotensiOp() potensiopwp.CreatePotensiOpDto {
	return s.PotensiOp
}

func (s *DPHotelUsecase) SetPotensiOpId(id uuid.UUID) {
	s.PotensiOp.Id = id
}

func (s *DPHotelUsecase) GetDetailPotensiOp() detailpotensiop.CreateDto {
	return s.DetailPotensiOp
}

func (s *DPHotelUsecase) GetPotensiPemilikWps() []potensipemilikwp.CreateDto {
	return s.PotensiPemilikWps
}

func (s *DPHotelUsecase) GetPotensiNarahubungs() []potensinarahubung.CreateDto {
	return s.PotensiNarahubungs
}

func (s *DPHotelUsecase) GetBapl() bapl.CreateDto {
	return s.Bapl
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
