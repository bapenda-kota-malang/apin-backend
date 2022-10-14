package detailsptresto

import (
	"strconv"
	"time"

	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mdsres "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
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

func Update(id int, input ms.UpdateRestoDto) (any, error) {
	// spt
	var data *ms.Spt
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	data.SptDate = time.Now()
	data.StartDate = th.ParseTime(input.StartDate)
	data.EndDate = th.ParseTime(input.EndDate)
	data.DueDate = th.ParseTime(input.DueDate)
	data.TanggalLunas = th.ParseTime(input.TanggalLunas)
	data.CancelledAt = th.ParseTime(input.CancelledAt)
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	//detailspt
	var dataD mdsres.DetailSptResto
	result = a.DB.Where(mdsres.DetailSptResto{Spt_Id: uint64(id)}).First(&dataD)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&dataD, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataD)
	}
	if result := a.DB.Save(&dataD); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataD)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil

}
