package potensiopwp

import (
	"errors"
	"math"
	"strings"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	mtp "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"
	mtr "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifreklame"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	str "github.com/bapenda-kota-malang/apin-backend/internal/services/tarifreklame"
	"github.com/google/uuid"
	"gorm.io/gorm"

	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
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

func (s *DPReklameUsecase) CalculateTax(tarifPajak *mtp.TarifPajak) (potensiopwp.CreatePotensiOpDto, error) {
	tax := float64(0)
	for i := 0; i < len(s.DetailPajakDtos); i++ {
		tahun := float64(s.DetailPajakDtos[i].JumlahTahun)
		bulan := float64(s.DetailPajakDtos[i].JumlahBulan)
		minggu := float64(s.DetailPajakDtos[i].JumlahMinggu)
		hari := float64(s.DetailPajakDtos[i].JumlahHari)

		resp, err := str.GetDetail(int(s.DetailPajakDtos[i].TarifReklame_Id))
		if err != nil {
			return m.CreatePotensiOpDto{}, err
		}
		tarifReklame := resp.(rp.OKSimple).Data.(*mtr.TarifReklame)

		// calculate base on jenis masa pajak
		switch s.DetailPajakDtos[i].JenisMasa {
		case mtypes.MasaPajakTahun:
			tahun *= *tarifReklame.Tarif
			bulan *= (*tarifReklame.Tarif / 12)
		case mtypes.MasaPajakBulan:
			bulan *= *tarifReklame.Tarif
		case mtypes.MasaPajakHari:
			resp, err = str.GetList(mtr.FilterDto{
				JenisMasa:    int16(mtypes.MasaPajakHari),
				JenisReklame: tarifReklame.JenisReklame,
			})
			if err != nil {
				return m.CreatePotensiOpDto{}, err
			}
			for _, vTR := range resp.(rp.OK).Data.([]mtr.TarifReklame) {
				if vTR.MasaPajak == nil {
					continue
				}
				switch strings.ToLower(*vTR.MasaPajak) {
				case "bulan":
					bulan *= *vTR.Tarif
				case "minggu":
					minggu *= *vTR.Tarif
				case "hari":
					hari *= *vTR.Tarif
				}
			}
		case mtypes.MasaPajakPenyelenggara:
			hari *= *tarifReklame.Tarif
		}
		// basic calculate tax
		basicTax := (hari + minggu + bulan + tahun) * float64(s.DetailPajakDtos[i].JumlahReklame)

		// if dasar pengenaan luas then add to calculation
		if *tarifReklame.DasarPengenaan == "Luas" {
			sisi := s.DetailPajakDtos[i].JumlahSisi
			if sisi < 1 {
				sisi = 1
				s.DetailPajakDtos[i].JumlahSisi = uint64(sisi)
			}

			luas := float64(0)
			if s.DetailPajakDtos[i].Diameter != nil {
				r := float64(*s.DetailPajakDtos[i].Diameter) / 2
				luas = math.Pi * math.Pow(r, 2)
			} else if s.DetailPajakDtos[i].Panjang != nil && s.DetailPajakDtos[i].Lebar != nil {
				luas = *s.DetailPajakDtos[i].Panjang * *s.DetailPajakDtos[i].Lebar
			} else {
				return m.CreatePotensiOpDto{}, errors.New("dasar pengenaan luas tetapi diameter / panjang dan lebar null")
			}

			basicTax *= (luas * float64(sisi))
		}
		s.DetailPajakDtos[i].JumlahRp = basicTax
		tax += basicTax
	}
	if tarifPajak != nil {
		s.PotensiOp.TarifPajak_Id = &tarifPajak.Id
	}
	s.PotensiOp.JumlahPajak = tax
	return s.PotensiOp, nil
}
