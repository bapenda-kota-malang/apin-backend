package espt

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	mdair "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	mdhib "github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthiburan"
	mdhot "github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthotel"
	mdpar "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptparkir"
	mdnonpln "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptppjnonpln"
	mdpln "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptppjpln"
	mdres "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptresto"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	mjppj "github.com/bapenda-kota-malang/apin-backend/internal/models/jenisppj"
	mtp "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"

	sair "github.com/bapenda-kota-malang/apin-backend/internal/services/detailesptair"
	shib "github.com/bapenda-kota-malang/apin-backend/internal/services/detailespthiburan"
	shot "github.com/bapenda-kota-malang/apin-backend/internal/services/detailespthotel"
	spar "github.com/bapenda-kota-malang/apin-backend/internal/services/detailesptparkir"
	snonpln "github.com/bapenda-kota-malang/apin-backend/internal/services/detailesptppjnonpln"
	spln "github.com/bapenda-kota-malang/apin-backend/internal/services/detailesptppjpln"
	sres "github.com/bapenda-kota-malang/apin-backend/internal/services/detailesptresto"
	shda "github.com/bapenda-kota-malang/apin-backend/internal/services/hargadasarair"
	sjppj "github.com/bapenda-kota-malang/apin-backend/internal/services/jenisppj"
	stp "github.com/bapenda-kota-malang/apin-backend/internal/services/tarifpajak"
)

// Service create business flow for esptd via wajib pajak for lapor e-sptd
//
// function flow is:
//
// create for esptd, replace id, create for data details based on data type, assign data details to data espt for respond
func CreateDetail(input m.CreateInput, user_Id uint) (interface{}, error) {
	var data m.Espt
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		rekeningId := uint64(input.GetEspt().Rekening_Id)
		yearNow := uint64(time.Now().Year())
		tarifPajak, err := stp.GetTarif(&rekeningId, &yearNow)
		if err != nil {
			return err
		}

		if detail, ok := input.GetDetails().(mdair.CreateDto); ok {
			switch detail.Peruntukan {
			case mdair.PeruntukanIndustriAir:
			case mdair.PeruntukanNiaga:
			case mdair.PeruntukanNonNiaga:
			case mdair.PeruntukanPdam:
			default:
				return fmt.Errorf("unknown peruntukan air")
			}
			hargaDasarAir, err := shda.GetTarif(string(detail.Peruntukan), fmt.Sprintf("%f", input.GetEspt().Omset))
			if err != nil {
				return err
			}
			if detail.JenisAbt == mdair.JenisABTNonMataAir {
				detail.Pengenaan = float32(*hargaDasarAir.TarifBukanMataAir)
			} else {
				detail.JenisAbt = mdair.JenisABTMataAir
				detail.Pengenaan = float32(*hargaDasarAir.TarifMataAir)
			}
			input.ChangeDetails(detail)
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

		input.ReplaceTarifPajakId(uint(tarifPajak.Id))
		input.CalculateTax(tarifPajak.TarifPersen)

		respEspt, err := Create(input.GetEspt(), user_Id, tx)
		if err != nil {
			return err
		}
		data = respEspt.(rp.OKSimple).Data.(m.Espt)

		input.ReplaceEsptId(data.Id)

		if input.LenDetails() == 0 {
			return nil
		}

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
		case []mdhot.CreateDto:
			respDetails, err := shot.Create(dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.([]mdhot.DetailEsptHotel)
				data.DetailEsptHotel = &details
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
		case []mdpar.CreateDto:
			respDetails, err := spar.Create(dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.([]mdpar.DetailEsptParkir)
				data.DetailEsptParkir = &details
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
				data.DetailEsptPpjPln = &details
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
func UpdateDetail(id int, input m.UpdateInput, user_Id uint) (interface{}, error) {
	var data m.Espt
	affected := "0"
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		rekeningId := uint64(*input.GetEspt().Rekening_Id)
		yearNow := uint64(time.Now().Year())
		tarifPajak, err := stp.GetList(mtp.FilterDto{Rekening_Id: &rekeningId, Tahun: &yearNow})
		if err != nil {
			return err
		}

		input.CalculateTax(tarifPajak.(rp.OK).Data.([]mtp.TarifPajak)[0].TarifPersen)

		respEspt, err := Update(id, input.GetEspt(), user_Id, tx)
		if err != nil {
			return err
		}
		data = respEspt.(rp.OK).Data.(m.Espt)
		affected = respEspt.(rp.OK).Meta["affected"]

		if input.LenDetails() == 0 {
			return nil
		}

		switch dataReal := input.GetDetails().(type) {
		case []mdair.UpdateDto:
			var detailList []mdair.DetailEsptAir
			for i := range dataReal {
				id := 0
				if dataReal[i].Id != 0 {
					id = int(dataReal[i].Id)
				}
				respDetails, err := sair.Update(id, dataReal[i], tx)
				if err != nil {
					return err
				}
				if respDetails != nil {
					details := respDetails.(rp.OKSimple).Data.(mdair.DetailEsptAir)
					detailList = append(detailList, details)
				}
			}
			// data.DetailEsptAir = &detailList
		case []mdhot.UpdateDto:
			var detailList []mdhot.DetailEsptHotel
			for i := range dataReal {
				id := 0
				if dataReal[i].Id != 0 {
					id = int(dataReal[i].Id)
				}
				respDetails, err := shot.Update(id, dataReal[i], tx)
				if err != nil {
					return err
				}
				if respDetails != nil {
					details := respDetails.(rp.OKSimple).Data.(mdhot.DetailEsptHotel)
					detailList = append(detailList, details)
				}
			}
			data.DetailEsptHotel = &detailList
		case []mdhib.UpdateDto:
			var detailList []mdhib.DetailEsptHiburan
			for i := range dataReal {
				id := 0
				if dataReal[i].Id != 0 {
					id = int(dataReal[i].Id)
				}
				respDetails, err := shib.Update(id, dataReal[i], tx)
				if err != nil {
					return err
				}
				if respDetails != nil {
					details := respDetails.(rp.OKSimple).Data.(mdhib.DetailEsptHiburan)
					detailList = append(detailList, details)
				}
			}
			// data.DetailEsptHiburan = &detailList
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
			data.DetailEsptParkir = &detailList
		case []mdres.UpdateDto:
			var detailList []mdres.DetailEsptResto
			for i := range dataReal {
				id := 0
				if dataReal[i].Id != 0 {
					id = int(dataReal[i].Id)
				}
				respDetails, err := sres.Update(id, dataReal[i], tx)
				if err != nil {
					return err
				}
				if respDetails != nil {
					details := respDetails.(rp.OKSimple).Data.(mdres.DetailEsptResto)
					detailList = append(detailList, details)
				}
			}
			// data.DetailEsptResto = &detailList
		case []mdnonpln.UpdateDto:
			var detailList []mdnonpln.DetailEsptPpjNonPln
			for i := range dataReal {
				id := 0
				if dataReal[i].Id != 0 {
					id = int(dataReal[i].Id)
				}
				respDetails, err := snonpln.Update(id, dataReal[i], tx)
				if err != nil {
					return err
				}
				if respDetails != nil {
					details := respDetails.(rp.OKSimple).Data.(mdnonpln.DetailEsptPpjNonPln)
					detailList = append(detailList, details)
				}
			}
			// data.DetailEsptPpjNonPln = &detailList
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
			data.DetailEsptPpjPln = &detailList
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
