package spt

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mdair "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptair"

	// mdhib "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthiburan"
	// mdhot "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthotel"
	// mdpar "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	// mdnonpln "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptppjnonpln"
	// mdpln "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptppjpln"
	// mdres "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"
	mhdair "github.com/bapenda-kota-malang/apin-backend/internal/models/hargadasarair"
	mtp "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"

	sair "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptair"
	// shib "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailspthiburan"
	// shot "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailspthotel"
	// spar "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptparkir"
	// snonpln "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptppjnonpln"
	// spln "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptppjpln"
	// sres "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptresto"
	shda "github.com/bapenda-kota-malang/apin-backend/internal/services/hargadasarair"
	stp "github.com/bapenda-kota-malang/apin-backend/internal/services/tarifpajak"
)

// Process calculate tax
func taxProcess(rekeningId uint64, omset float64, input m.Input) error {
	yearNow := uint64(time.Now().Year())
	omsetOpt := "lte"

	// change detail for detail spt air & ppj pln before calculate tax
	if detail, ok := input.GetDetails().(mdair.CreateDto); ok {
		switch detail.Peruntukan {
		case mtypes.PeruntukanIndustriAir, mtypes.PeruntukanNiaga, mtypes.PeruntukanNonNiaga, mtypes.PeruntukanPdam:
			resphdair, err := shda.GetList(mhdair.FilterDto{
				Peruntukan:     &detail.Peruntukan,
				BatasBawah:     &omset,
				BatasBawah_Opt: &omsetOpt,
			})
			if err != nil {
				return err
			}
			hdairs := resphdair.(rp.OK).Data.([]mhdair.HargaDasarAir)
			if len(hdairs) == 0 {
				return fmt.Errorf("harga dasar air not found")
			}
			hdair := hdairs[len(hdairs)-1]
			if detail.JenisAbt == mtypes.JenisABTNonMataAir {
				detail.Pengenaan = *hdair.TarifBukanMataAir
			} else {
				detail.JenisAbt = mtypes.JenisABTMataAir
				detail.Pengenaan = *hdair.TarifMataAir
			}
			input.ChangeDetails(detail)
		default:
			return fmt.Errorf("unknown peruntukan air")
		}
	}

	// if detail, ok := input.GetDetails().([]mdpln.CreateDto); ok {
	// 	for v := range detail {
	// 		resp, err := sjppj.GetDetail(int(detail[v].JenisPPJ_Id))
	// 		if err != nil {
	// 			return err
	// 		}
	// 		detail[v].JenisPPJ = resp.(rp.OKSimple).Data.(*mjppj.JenisPPJ)
	// 	}
	// 	input.ChangeDetails(detail)
	// }

	// get tarif pajak data for calculate tax
	rspTp, err := stp.GetList(mtp.FilterDto{
		Rekening_Id:   &rekeningId,
		Tahun:         &yearNow,
		OmsetAwal:     &omset,
		OmsetAwal_Opt: &omsetOpt,
	})
	if err != nil {
		return err
	}

	tarifPajaks := rspTp.(rp.OK).Data.([]mtp.TarifPajak)
	if len(tarifPajaks) == 0 {
		return fmt.Errorf("tarif pajak not found")
	}
	tarifPajak := tarifPajaks[len(tarifPajaks)-1]

	input.ReplaceTarifPajakId(uint(tarifPajak.Id))
	input.CalculateTax(tarifPajak.TarifPersen)
	return nil
}

// Transform espt to spt data
func TransformEspt(esptDetail *mespt.Espt) (input m.Input, err error) {
	if esptDetail.DetailEsptAir != nil {
		input = &m.CreateDetailAirDto{}
	}
	if esptDetail.DetailEsptHiburan != nil {
		// TODO: ADD HIBURAN
		// input = &m.{}
	}
	if esptDetail.DetailEsptHotel != nil {
		input = &m.CreateDetailHotelDto{}
	}
	if esptDetail.DetailEsptParkir != nil {
		input = &m.CreateDetailParkirDto{}
	}
	if esptDetail.DetailEsptPpjNonPln != nil {
		// TODO: ADD MODEL
		// input = &m.CreateDetailAirDto{}
	}
	if esptDetail.DetailEsptPpjPln != nil {
		// TODO: ADD MODEL
		// input = &m.CreateDetailAirDto{}
	}
	if esptDetail.DetailEsptResto != nil {
		input = &m.CreateDetailRestoDto{}
	}

	if input == nil {
		err = fmt.Errorf("uknown espt data for espt")
	}

	input.DuplicateEspt(esptDetail)

	return
}

// Service create business flow for sptd via wajib pajak for lapor spt
//
// function flow is:
//
// create for sptd, replace id, create for data details based on data type, assign data details to data spt for respond
func CreateDetail(input m.Input, user_Id uint, newFile bool, tx *gorm.DB) (interface{}, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.Spt
	createDto := input.GetSpt().(m.CreateDto)

	err := tx.Transaction(func(tx *gorm.DB) error {
		if createDto.JumlahPajak == 0 {
			err := taxProcess(*createDto.Rekening_Id, createDto.Omset, input)
			if err != nil {
				return err
			}
			createDto = input.GetSpt().(m.CreateDto)
		}

		respSpt, err := Create(createDto, user_Id, newFile, tx)
		if err != nil {
			return err
		}
		data = respSpt.(rp.OKSimple).Data.(m.Spt)

		input.ReplaceSptId(uint(data.Id))

		if input.LenDetails() == 0 {
			return nil
		}

		switch dataDetails := input.GetDetails().(type) {
		case mdair.CreateDto:
			respDetails, err := sair.Create(dataDetails, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdair.DetailSptAir)
				data.DetailSptAir = &details
			}
		default:
			return fmt.Errorf("data details unknown")
		}

		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data: %s", err), data)
	}
	return rp.OKSimple{Data: data}, nil
}
