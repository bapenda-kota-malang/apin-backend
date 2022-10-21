package detailsptair

import (
	"strconv"
	"time"

	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mdsa "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptair"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/datatypes"
)

const source = "detailSptAir"

func Create(input ms.CreateAirDto) (any, error) {
	var dataS ms.Spt
	var dataD mdsa.DetailSptAir

	if err := sc.Copy(&dataS, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataS)
	}
	if err := sc.Copy(&dataD, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataD)
	}
	prevMonth := sh.BeginningOfPreviosMonth()
	dataS.SptDate = time.Now()
	dataS.StartDate = datatypes.Date(prevMonth)
	dataS.EndDate = datatypes.Date(sh.EndOfMonth(prevMonth))
	dataS.DueDate = datatypes.Date(sh.EndOfMonth(time.Now()))
	dataS.TanggalLunas = th.ParseTime(input.TanggalLunas)
	dataS.CancelledAt = th.ParseTime(input.CancelledAt)

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
		"spt":          dataS,
		"detailSptAir": dataD,
	}}, nil
}

func Update(id int, input ms.UpdateAirDto) (any, error) {
	// spt
	var data *ms.Spt
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	prevMonth := sh.BeginningOfPreviosMonth()
	data.SptDate = time.Now()
	data.StartDate = datatypes.Date(prevMonth)
	data.EndDate = datatypes.Date(sh.EndOfMonth(prevMonth))
	data.DueDate = datatypes.Date(sh.EndOfMonth(time.Now()))
	data.TanggalLunas = th.ParseTime(input.TanggalLunas)
	data.CancelledAt = th.ParseTime(input.CancelledAt)
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	//detailspt
	var dataD mdsa.DetailSptAir
	result = a.DB.Where(mdsa.DetailSptAir{Spt_Id: uint64(id)}).First(&dataD)
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
