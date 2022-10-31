package detailsptresto

import (
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
)

const source = "detailSptResto"

func Create(input m.CreateDto) (any, error) {
	var data m.DetailSptResto

	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	err := a.DB.Create(&data).Error
	if err != nil {
		return nil, err
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(id int, input m.UpdateDto) (any, error) {
	var data *m.DetailSptResto
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	// prevMonth := sh.BeginningOfPreviosMonth()
	// data.SptDate = time.Now()
	// data.StartDate = datatypes.Date(prevMonth)
	// data.EndDate = datatypes.Date(sh.EndOfMonth(prevMonth))
	// data.DueDate = datatypes.Date(sh.EndOfMonth(time.Now()))
	// data.TanggalLunas = th.ParseTime(input.TanggalLunas)
	// data.CancelledAt = th.ParseTime(input.CancelledAt)
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
