package espt

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/google/uuid"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mdair "github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptair"
	mdhib "github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailespthiburan"
	mdhot "github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailespthotel"
	mdpar "github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptparkir"
	mdnonpln "github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptppjnonpln"
	mdpln "github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptppjpln"
	mdres "github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptresto"
	mhdair "github.com/bapenda-kota-malang/apin-backend/internal/models/hargadasarair"
	mjppj "github.com/bapenda-kota-malang/apin-backend/internal/models/jenisppj"
	mtp "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"

	sair "github.com/bapenda-kota-malang/apin-backend/internal/services/espt/detailesptair"
	shib "github.com/bapenda-kota-malang/apin-backend/internal/services/espt/detailespthiburan"
	shot "github.com/bapenda-kota-malang/apin-backend/internal/services/espt/detailespthotel"
	spar "github.com/bapenda-kota-malang/apin-backend/internal/services/espt/detailesptparkir"
	snonpln "github.com/bapenda-kota-malang/apin-backend/internal/services/espt/detailesptppjnonpln"
	spln "github.com/bapenda-kota-malang/apin-backend/internal/services/espt/detailesptppjpln"
	sres "github.com/bapenda-kota-malang/apin-backend/internal/services/espt/detailesptresto"
	shda "github.com/bapenda-kota-malang/apin-backend/internal/services/hargadasarair"
	sjppj "github.com/bapenda-kota-malang/apin-backend/internal/services/jenisppj"
	stp "github.com/bapenda-kota-malang/apin-backend/internal/services/tarifpajak"
)

// Process calculate tax
func taxProcess(rekeningId uint64, omset float64, input m.Input) error {
	omsetOpt := "lte"

	// change detail for detail spt air & ppj pln before calculate tax
	if detail, ok := input.GetDetails().(mdair.CreateDto); ok {
		switch detail.Peruntukan {
		case mtypes.PeruntukanIndustriAir, mtypes.PeruntukanNiaga, mtypes.PeruntukanNonNiaga, mtypes.PeruntukanPdam:
			strPeruntukan := string(detail.Peruntukan)
			resphdair, err := shda.GetList(mhdair.FilterDto{
				Peruntukan:     &strPeruntukan,
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
				detail.Pengenaan = float32(*hdair.TarifBukanMataAir)
			} else {
				detail.JenisAbt = mtypes.JenisABTMataAir
				detail.Pengenaan = float32(*hdair.TarifMataAir)
			}
			input.ChangeDetails(detail)
		default:
			return fmt.Errorf("unknown peruntukan air")
		}
	} else if detail, ok := input.GetDetails().([]mdpln.CreateDto); ok {
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
	yearOrder := "order desc"
	rspTp, err := stp.GetList(mtp.FilterDto{
		Rekening_Id:   &rekeningId,
		Tahun_Opt:     &yearOrder,
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
	tarifPajak := tarifPajaks[0]

	input.ReplaceTarifPajakId(uint(tarifPajak.Id))
	input.CalculateTax(tarifPajak.TarifPersen)
	return nil
}

// Service create business flow for esptd via wajib pajak for lapor e-sptd
//
// function flow is:
//
// create for esptd, replace id, create for data details based on data type, assign data details to data espt for respond
func CreateDetail(input m.Input, user_Id uint) (interface{}, error) {
	var data m.Espt

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		createDto := input.GetEspt().(m.CreateDto)
		err := taxProcess(uint64(createDto.Rekening_Id), createDto.Omset, input)
		if err != nil {
			return err
		}
		createDto = input.GetEspt().(m.CreateDto)
		respEspt, err := Create(createDto, user_Id, tx)
		if err != nil {
			return err
		}
		data = respEspt.(rp.OKSimple).Data.(m.Espt)

		if input.LenDetails() == 0 {
			return nil
		}

		input.ReplaceEsptId(data.Id)

		switch dataReal := input.GetDetails().(type) {
		case mdair.CreateDto:
			respDetails, err := sair.Create(dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdair.DetailEsptAir)
				data.DetailEsptAir = &details
			}
		case mdhib.CreateDto:
			respDetails, err := shib.Create(dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdhib.DetailEsptHiburan)
				data.DetailEsptHiburan = &details
			}
		case mdhot.CreateDto:
			respDetails, err := shot.Create(dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdhot.DetailEsptHotel)
				data.DetailEsptHotel = &details
			}
		case []mdpar.CreateDto:
			respDetails, err := spar.Create(dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.([]mdpar.DetailEsptParkir)
				data.DetailEsptParkir = details
			}
		case mdnonpln.CreateDto:
			respDetails, err := snonpln.Create(dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdnonpln.DetailEsptPpjNonPln)
				data.DetailEsptPpjNonPln = &details
			}
		case []mdpln.CreateDto:
			respDetails, err := spln.Create(dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.([]mdpln.DetailEsptPpjPln)
				data.DetailEsptPpjPln = details
			}
		case mdres.CreateDto:
			respDetails, err := sres.Create(dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdres.DetailEsptResto)
				data.DetailEsptResto = &details
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

// Service business flow for esptd via wajib pajak for update lapor e-sptd
//
// this function flow is:
//
// update data esptd with id, check data details type, loop for update every items, assign data details to data espt for respond
func UpdateDetail(ctx context.Context, id uuid.UUID, input m.Input, user_Id uint) (interface{}, error) {
	var data m.Espt
	affected := "0"
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		updateDto := input.GetEspt().(m.UpdateDto)
		err := taxProcess(uint64(*updateDto.Rekening_Id), *updateDto.Omset, input)
		if err != nil {
			return err
		}
		updateDto = input.GetEspt().(m.UpdateDto)

		respEspt, err := Update(ctx, id, updateDto, user_Id, tx)
		if err != nil {
			return err
		}
		data = respEspt.(rp.OK).Data.(m.Espt)
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
			respDetails, err := sair.Update(id, dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdair.DetailEsptAir)
				data.DetailEsptAir = &details
			}
		case mdhib.UpdateDto:
			id := 0
			if dataReal.Id != 0 {
				id = int(dataReal.Id)
			}
			respDetails, err := shib.Update(id, dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdhib.DetailEsptHiburan)
				data.DetailEsptHiburan = &details
			}
		case mdhot.UpdateDto:
			id := 0
			if dataReal.Id != 0 {
				id = int(dataReal.Id)
			}
			respDetails, err := shot.Update(id, dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdhot.DetailEsptHotel)
				data.DetailEsptHotel = &details
			}
		case []mdpar.UpdateDto:
			var detailList []mdpar.DetailEsptParkir
			for i := range dataReal {
				id := 0
				if dataReal[i].Id != 0 {
					id = int(dataReal[i].Id)
				}
				respDetails, err := spar.Update(id, dataReal[i], tx)
				if err != nil {
					return err
				}
				if respDetails != nil {
					details := respDetails.(rp.OKSimple).Data.(mdpar.DetailEsptParkir)
					detailList = append(detailList, details)
				}
			}
			data.DetailEsptParkir = detailList
		case mdnonpln.UpdateDto:
			id := 0
			if dataReal.Id != 0 {
				id = int(dataReal.Id)
			}
			respDetails, err := snonpln.Update(id, dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdnonpln.DetailEsptPpjNonPln)
				data.DetailEsptPpjNonPln = &details
			}
		case []mdpln.UpdateDto:
			var detailList []mdpln.DetailEsptPpjPln
			for i := range dataReal {
				id := 0
				if dataReal[i].Id != 0 {
					id = int(dataReal[i].Id)
				}
				respDetails, err := spln.Update(id, dataReal[i], tx)
				if err != nil {
					return err
				}
				if respDetails != nil {
					details := respDetails.(rp.OKSimple).Data.(mdpln.DetailEsptPpjPln)
					detailList = append(detailList, details)
				}
			}
			data.DetailEsptPpjPln = detailList
		case mdres.UpdateDto:
			id := 0
			if dataReal.Id != 0 {
				id = int(dataReal.Id)
			}
			respDetails, err := sres.Update(id, dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.(mdres.DetailEsptResto)
				data.DetailEsptResto = &details
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
