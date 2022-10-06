package espt

import (
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	mdair "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	mdhot "github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthotel"
	mdpar "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptparkir"
	mdrek "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptreklame"
	mdres "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptresto"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"

	sair "github.com/bapenda-kota-malang/apin-backend/internal/services/detailesptair"
	shot "github.com/bapenda-kota-malang/apin-backend/internal/services/detailespthotel"
	spar "github.com/bapenda-kota-malang/apin-backend/internal/services/detailesptparkir"
	srek "github.com/bapenda-kota-malang/apin-backend/internal/services/detailesptreklame"
	sres "github.com/bapenda-kota-malang/apin-backend/internal/services/detailesptresto"
)

// Service create business flow for esptd air tanah
func CreateAir(input m.CreateDetailAirDto) (interface{}, error) {
	var data m.Espt
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		respEspt, err := Create(input.Espt, tx)
		if err != nil {
			return err
		}
		data = respEspt.(rp.OKSimple).Data.(m.Espt)

		for i := range input.DataDetails {
			input.DataDetails[i].Espt_Id = data.Id
		}

		respDetail, err := sair.Create(input.DataDetails, tx)
		if err != nil {
			return err
		}
		data.DetailEsptAir = respDetail.(rp.OKSimple).Data.(*[]mdair.DetailEsptAir)
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}
	return rp.OKSimple{Data: data}, nil
}

func CreateHotel(input m.CreateDetailHotelDto) (interface{}, error) {
	var data m.Espt
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		respEspt, err := Create(input.Espt, tx)
		if err != nil {
			return err
		}
		data = respEspt.(rp.OKSimple).Data.(m.Espt)

		for i := range input.DataDetails {
			input.DataDetails[i].Espt_Id = data.Id
		}

		respDetail, err := shot.Create(input.DataDetails, tx)
		if err != nil {
			return err
		}
		data.DetailEsptHotel = respDetail.(rp.OKSimple).Data.(*[]mdhot.DetailEsptHotel)
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}
	return rp.OKSimple{Data: data}, nil
}

func CreateParkir(input m.CreateDetailParkirDto) (interface{}, error) {
	var data m.Espt
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		respEspt, err := Create(input.Espt, tx)
		if err != nil {
			return err
		}
		data = respEspt.(rp.OKSimple).Data.(m.Espt)

		for i := range input.DataDetails {
			input.DataDetails[i].Espt_Id = data.Id
		}

		respDetail, err := spar.Create(input.DataDetails, tx)
		if err != nil {
			return err
		}
		data.DetailEsptParkir = respDetail.(rp.OKSimple).Data.(*[]mdpar.DetailEsptParkir)
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}
	return rp.OKSimple{Data: data}, nil
}
func CreateReklame(input m.CreateDetailReklameDto) (interface{}, error) {
	var data m.Espt
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		respEspt, err := Create(input.Espt, tx)
		if err != nil {
			return err
		}
		data = respEspt.(rp.OKSimple).Data.(m.Espt)

		for i := range input.DataDetails {
			input.DataDetails[i].Espt_Id = data.Id
		}

		respDetail, err := srek.Create(input.DataDetails, tx)
		if err != nil {
			return err
		}
		data.DetailEsptReklame = respDetail.(rp.OKSimple).Data.(*[]mdrek.DetailEsptReklame)
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}
	return rp.OKSimple{Data: data}, nil
}
func CreateResto(input m.CreateDetailRestoDto) (interface{}, error) {
	var data m.Espt
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		respEspt, err := Create(input.Espt, tx)
		if err != nil {
			return err
		}
		data = respEspt.(rp.OKSimple).Data.(m.Espt)

		for i := range input.DataDetails {
			input.DataDetails[i].Espt_Id = data.Id
		}

		respDetail, err := sres.Create(input.DataDetails, tx)
		if err != nil {
			return err
		}
		data.DetailEsptResto = respDetail.(rp.OKSimple).Data.(*[]mdres.DetailEsptResto)
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}
	return rp.OKSimple{Data: data}, nil
}
