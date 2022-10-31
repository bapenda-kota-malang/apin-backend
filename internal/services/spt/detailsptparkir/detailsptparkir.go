package detailsptair

import (
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
)

const source = "detailSptHotel"

func Create(input m.CreateDto) (any, error) {
	var data m.DetailSptParkir

	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	err := a.DB.Create(&data).Error
	if err != nil {
		return nil, err
	}

	err = a.DB.Create(&data).Error
	if err != nil {
		return nil, err
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(id int, input m.UpdateDto) (any, error) {
	//detailspt
	var data m.DetailSptParkir
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil

}
