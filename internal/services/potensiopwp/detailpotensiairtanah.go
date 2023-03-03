package potensiopwp

import (
	"errors"

	mhda "github.com/bapenda-kota-malang/apin-backend/internal/models/hargadasarair"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	mtp "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"

	shda "github.com/bapenda-kota-malang/apin-backend/internal/services/hargadasarair"

	"github.com/google/uuid"
	"gorm.io/gorm"

	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sc "github.com/jinzhu/copier"
)

type DPAirTanahUsecase struct {
	DPBaseUsecase
	DetailPajakDtos m.CreateDtoDPAirTanah
}

func NewDPAirTanahUsecase(baseDto DPBaseUsecase, detailPajakDto m.CreateDtoDPAirTanah) m.Input {
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
	s.DetailPajakDtos.Potensiop_Id = id
}

func (s *DPAirTanahUsecase) GetDetailPotensiPajak() interface{} {
	return s.DetailPajakDtos
}

func (s *DPAirTanahUsecase) SaveDetailPotensiPajak(tx *gorm.DB) error {
	if tx == nil {
		return errors.New("tx empty, init db connection or use transaction")
	}
	var data m.DetailPotensiAirTanah
	if err := sc.Copy(&data, &s.DetailPajakDtos); err != nil {
		return errors.New("failed copy data to destination model detail potensi air tanah")
	}
	if result := tx.Create(&data); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *DPAirTanahUsecase) CalculateTax(tarifPajak mtp.TarifPajak) error {
	// harga dasar air
	tarifAir := float64(0)
	omsetOpt := "lte"
	switch s.DetailPajakDtos.Peruntukan {
	case mtypes.PeruntukanIndustriAir, mtypes.PeruntukanNiaga, mtypes.PeruntukanNonNiaga, mtypes.PeruntukanPdam:
		strPeruntukan := string(s.DetailPajakDtos.Peruntukan)
		resphdair, err := shda.GetList(mhda.FilterDto{
			Peruntukan:     &strPeruntukan,
			BatasBawah:     s.PotensiOp.OmsetOp,
			BatasBawah_Opt: &omsetOpt,
		})
		if err != nil {
			return err
		}
		hdairs := resphdair.(rp.OK).Data.([]mhda.HargaDasarAir)
		if len(hdairs) == 0 {
			return errors.New("harga dasar air not found")
		}
		hdair := hdairs[len(hdairs)-1]
		s.DetailPajakDtos.HargaDasarAir_Id = hdair.Id
		tarifAir = *hdair.TarifBukanMataAir
		if s.DetailPajakDtos.JenisAbt != mtypes.JenisABTNonMataAir {
			s.DetailPajakDtos.JenisAbt = mtypes.JenisABTMataAir
			tarifAir = *hdair.TarifMataAir
		}
	default:
		return errors.New("unknown peruntukan air")
	}

	// calculate tax
	s.PotensiOp.JumlahPajak = *s.PotensiOp.OmsetOp * (*tarifPajak.TarifPersen / 100) * tarifAir

	return nil
}
