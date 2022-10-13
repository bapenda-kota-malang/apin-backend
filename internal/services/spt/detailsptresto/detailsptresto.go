package detailsptresto

import (
	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mdsres "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
)

const source = "detailSptResto"

func Create(input ms.CreateRestoDto) (any, error) {
	var dataS ms.Spt
	var dataD mdsres.DetailSptResto

	if err := sc.Copy(&dataS, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataS)
	}
	if err := sc.Copy(&dataD, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataD)
	}

	err := a.DB.Create(&dataS).Error
	if err != nil {
		return nil, err
	}

	dataD.Spt_Id = dataS.Id

	err = a.DB.Create(&dataD).Error
	if err != nil {
		return nil, err
	}

	return rp.OKSimple{Data: t.II{
		"spt":            dataS,
		"detailSptResto": dataD,
	}}, nil
}
