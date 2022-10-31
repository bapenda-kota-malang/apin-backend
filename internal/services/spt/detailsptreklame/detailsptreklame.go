package detailsptreklame

import (
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
)

const source = "detailSptReklame"

func Create(input m.CreateDto) (any, error) {
	var data m.DetailSptReklame

	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	err := a.DB.Create(&data).Error
	if err != nil {
		return nil, err
	}

	return rp.OKSimple{Data: data}, nil
}
