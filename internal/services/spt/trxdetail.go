package spt

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/google/uuid"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mdair "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptair"
	mdhiburan "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthiburan"
	mdhotel "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthotel"
	mdparkir "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	mdnonpln "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptppjnonpln"
	mdpln "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptppjpln"
	mdreklame "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	mdresto "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"

	mespt "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mhdair "github.com/bapenda-kota-malang/apin-backend/internal/models/hargadasarair"
	mjppj "github.com/bapenda-kota-malang/apin-backend/internal/models/jenisppj"
	mtp "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"

	sdair "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptair"
	sdhiburan "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailspthiburan"
	sdhotel "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailspthotel"
	sdparkir "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptparkir"
	sdnonpln "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptppjnonpln"
	sdpln "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptppjpln"
	sdreklame "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptreklame"
	sdresto "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptresto"

	shda "github.com/bapenda-kota-malang/apin-backend/internal/services/hargadasarair"
	sjppj "github.com/bapenda-kota-malang/apin-backend/internal/services/jenisppj"
	stp "github.com/bapenda-kota-malang/apin-backend/internal/services/tarifpajak"
)

// Process calculate tax
func taxProcess(rekeningId *uint64, omset *float64, input m.Input) error {
	yearNow := uint64(time.Now().Year())
	omsetOpt := "lte"

	// change detail for detail spt air & ppj pln before calculate tax
	if detail, ok := input.GetDetails().(mdair.CreateDto); ok {
		switch detail.Peruntukan {
		case mtypes.PeruntukanIndustriAir, mtypes.PeruntukanNiaga, mtypes.PeruntukanNonNiaga, mtypes.PeruntukanPdam:
			strPeruntukan := string(detail.Peruntukan)
			resphdair, err := shda.GetList(mhdair.FilterDto{
				Peruntukan:     &strPeruntukan,
				BatasBawah:     omset,
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

	if detail, ok := input.GetDetails().([]mdpln.CreateDto); ok {
		for v := range detail {
			resp, err := sjppj.GetDetail(int(detail[v].JenisPPJ_Id))
			if err != nil {
				return err
			}
			detail[v].JenisPPJ = resp.(rp.OKSimple).Data.(*mjppj.JenisPPJ)
		}
		input.ChangeDetails(detail)
	}

	// get tarif pajak data for calculate tax
	rspTp, err := stp.GetList(mtp.FilterDto{
		Rekening_Id:   rekeningId,
		Tahun:         &yearNow,
		OmsetAwal:     omset,
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
	} else if esptDetail.DetailEsptHiburan != nil {
		input = &m.CreateDetailHiburanDto{}
	} else if esptDetail.DetailEsptParkir != nil {
		input = &m.CreateDetailParkirDto{}
	} else if esptDetail.DetailEsptHotel != nil {
		input = &m.CreateDetailHotelDto{}
	} else if esptDetail.DetailEsptPpjNonPln != nil {
		input = &m.CreateDetailPpjNonPlnDto{}
	} else if esptDetail.DetailEsptPpjPln != nil {
		input = &m.CreateDetailPpjPlnDto{}
	} else if esptDetail.DetailEsptResto != nil {
		input = &m.CreateDetailRestoDto{}
	} else {
		input = &m.CreateDetailBaseDto{}
	}

	input.DuplicateEspt(esptDetail)

	return
}

// Service create business flow for sptd via wajib pajak for lapor spt
//
// function flow is:
//
// create for sptd, replace id, create for data details based on data type, assign data details to data spt for respond
func CreateDetail(input m.Input, opts map[string]interface{}, tx *gorm.DB) (interface{}, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.Spt

	err := tx.Transaction(func(tx *gorm.DB) error {
		createDto := input.GetSpt(opts["baseUri"].(string)).(m.CreateDto)
		if createDto.JumlahPajak == 0 {
			err := taxProcess(createDto.Rekening_Id, &createDto.Omset, input)
			if err != nil {
				return err
			}
			if opts["baseUri"].(string) == "skpdkb" {
				input.CalculateSkpdkb()
			}
			createDto = input.GetSpt(opts["baseUri"].(string)).(m.CreateDto)
		}

		respSpt, err := Create(createDto, opts, tx)
		if err != nil {
			return err
		}
		data = respSpt.(rp.OKSimple).Data.(m.Spt)

		if input.LenDetails() == 0 {
			return nil
		}

		input.ReplaceSptId(data.Id)

		switch dataDetails := input.GetDetails().(type) {
		case mdair.CreateDto:
			respDetails, err := sdair.Create(dataDetails, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdair.DetailSptAir)
				data.DetailSptAir = &details
			}
		case mdhiburan.CreateDto:
			respDetails, err := sdhiburan.Create(dataDetails, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdhiburan.DetailSptHiburan)
				data.DetailSptHiburan = &details
			}
		case []mdhotel.CreateDto:
			respDetails, err := sdhotel.Create(dataDetails, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.([]mdhotel.DetailSptHotel)
				data.DetailSptHotel = &details
			}
		case []mdparkir.CreateDto:
			respDetails, err := sdparkir.Create(dataDetails, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.([]mdparkir.DetailSptParkir)
				data.DetailSptParkir = &details
			}
		case mdnonpln.CreateDto:
			respDetails, err := sdnonpln.Create(dataDetails, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdnonpln.DetailSptPpjNonPln)
				data.DetailSptNonPln = &details
			}
		case []mdpln.CreateDto:
			respDetails, err := sdpln.Create(dataDetails, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.([]mdpln.DetailSptPpjPln)
				data.DetailSptPln = &details
			}
		case []mdreklame.CreateDto:
			// TODO: REKLAME PROCESS
			respDetails, err := sdreklame.Create(dataDetails, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.([]mdreklame.DetailSptReklame)
				data.DetailSptReklame = &details
			}
		case mdresto.CreateDto:
			respDetails, err := sdresto.Create(dataDetails, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdresto.DetailSptResto)
				data.DetailSptResto = &details
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

func UpdateDetail(id uuid.UUID, input m.Input, opts map[string]interface{}) (interface{}, error) {
	var data m.Spt
	affected := "0"
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		updateDto := input.GetSpt(opts["baseUri"].(string)).(m.UpdateDto)
		err := taxProcess(updateDto.Rekening_Id, updateDto.Omset, input)
		if err != nil {
			return err
		}
		updateDto = input.GetSpt(opts["baseUri"].(string)).(m.UpdateDto)

		respEspt, err := Update(id, updateDto, opts, tx)
		if err != nil {
			return err
		}
		data = respEspt.(rp.OK).Data.(m.Spt)
		affected = respEspt.(rp.OK).Meta["affected"]

		if input.LenDetails() == 0 {
			return nil
		}

		switch dataReal := input.GetDetails().(type) {
		case mdair.UpdateDto:
			id := 0
			if dataReal.Id != 0 {
				id = int(dataReal.Id)
			}
			respDetails, err := sdair.Update(id, dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdair.DetailSptAir)
				data.DetailSptAir = &details
			}
		case mdhiburan.UpdateDto:
			id := 0
			if dataReal.Id != 0 {
				id = int(dataReal.Id)
			}
			respDetails, err := sdhiburan.Update(id, dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdhiburan.DetailSptHiburan)
				data.DetailSptHiburan = &details
			}
		case []mdhotel.UpdateDto:
			var detailList []mdhotel.DetailSptHotel
			for i := range dataReal {
				id := 0
				if dataReal[i].Id != 0 {
					id = int(dataReal[i].Id)
				}
				respDetails, err := sdhotel.Update(id, dataReal[i], tx)
				if err != nil {
					return err
				}
				if respDetails != nil {
					details := respDetails.(rp.OKSimple).Data.(mdhotel.DetailSptHotel)
					detailList = append(detailList, details)
				}
			}
			data.DetailSptHotel = &detailList
		case []mdparkir.UpdateDto:
			var detailList []mdparkir.DetailSptParkir
			for i := range dataReal {
				id := 0
				if dataReal[i].Id != 0 {
					id = int(dataReal[i].Id)
				}
				respDetails, err := sdparkir.Update(id, dataReal[i], tx)
				if err != nil {
					return err
				}
				if respDetails != nil {
					details := respDetails.(rp.OKSimple).Data.(mdparkir.DetailSptParkir)
					detailList = append(detailList, details)
				}
			}
			data.DetailSptParkir = &detailList
		case mdnonpln.UpdateDto:
			id := 0
			if dataReal.Id != 0 {
				id = int(dataReal.Id)
			}
			respDetails, err := sdnonpln.Update(id, dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdnonpln.DetailSptPpjNonPln)
				data.DetailSptNonPln = &details
			}
		case []mdpln.UpdateDto:
			var detailList []mdpln.DetailSptPpjPln
			for i := range dataReal {
				id := 0
				if dataReal[i].Id != 0 {
					id = int(dataReal[i].Id)
				}
				respDetails, err := sdpln.Update(id, dataReal[i], tx)
				if err != nil {
					return err
				}
				if respDetails != nil {
					details := respDetails.(rp.OKSimple).Data.(mdpln.DetailSptPpjPln)
					detailList = append(detailList, details)
				}
			}
			data.DetailSptPln = &detailList
		case mdresto.UpdateDto:
			id := 0
			if dataReal.Id != 0 {
				id = int(dataReal.Id)
			}
			respDetails, err := sdresto.Update(id, dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdresto.DetailSptResto)
				data.DetailSptResto = &details
			}
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", err.Error(), data)
	}
	return rp.OK{
		Meta: t.IS{
			"affected": affected,
		},
		Data: data,
	}, nil
}
