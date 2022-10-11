package detailsptparkir

import (
	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mdsp "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
)

const source = "detailSptParkir"

func init() {
	a.AutoMigrate(&ms.Spt{})
	a.AutoMigrate(&mdsp.DetailSptParkir{})
}

func Create(input ms.CreateParkirDto) (any, error) {
	var dataS ms.Spt

	if err := sc.Copy(&dataS, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataS)
	}

	err := a.DB.Create(&dataS).Error
	if err != nil {
		return nil, err
	}

	for _, v := range input.DetailSpt {
		var dataD mdsp.DetailSptParkir
		if err := sc.Copy(&dataD, v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", v)
		}

		dataD.Spt_Id = dataS.Id

		err = a.DB.Create(&dataD).Error
		if err != nil {
			return nil, err
		}

	}

	return rp.OKSimple{Data: t.II{
		"spt":             dataS,
		"detailSptParkir": input.DetailSpt,
	}}, nil
}
