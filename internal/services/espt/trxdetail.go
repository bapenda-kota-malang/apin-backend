package espt

import (
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	mdair "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	mdhib "github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthiburan"
	mdhot "github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthotel"
	mdpar "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptparkir"
	mdres "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptresto"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"

	sair "github.com/bapenda-kota-malang/apin-backend/internal/services/detailesptair"
	shib "github.com/bapenda-kota-malang/apin-backend/internal/services/detailespthiburan"
	shot "github.com/bapenda-kota-malang/apin-backend/internal/services/detailespthotel"
	spar "github.com/bapenda-kota-malang/apin-backend/internal/services/detailesptparkir"
	sres "github.com/bapenda-kota-malang/apin-backend/internal/services/detailesptresto"
)

// Service create business flow for esptd via wajib pajak for lapor e-sptd
//
// function flow is:
//
// create for esptd, replace id, create for data details based on data type, assign data details to data espt for respond
func CreateDetail(input m.CreateInput, user_Id uint) (interface{}, error) {
	var data m.Espt
	err := a.DB.Transaction(func(tx *gorm.DB) error {
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
		case []mdair.CreateDto:
			respDetails, err := sair.Create(dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.([]mdair.DetailEsptAir)
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
		case []mdhib.CreateDto:
			respDetails, err := shib.Create(dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.([]mdhib.DetailEsptHiburan)
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
		case []mdres.CreateDto:
			respDetails, err := sres.Create(dataReal, tx)
			if err != nil {
				return err
			}
			if respDetails != nil {
				details := respDetails.(rp.OKSimple).Data.([]mdres.DetailEsptResto)
				data.DetailEsptResto = &details
			}
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
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
			data.DetailEsptAir = &detailList
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
			data.DetailEsptHiburan = &detailList
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
			data.DetailEsptResto = &detailList
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
