package spt

import (
	"errors"
	"strconv"
	"time"

	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mdsa "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptair"
	mdsh "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthotel"
	mdsp "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	mdsrek "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	mdsres "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"
	mdjbr "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/jaminanbongkarreklame"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
)

const source = "spt"

func Create(input m.CreateDto, user_Id uint, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.Spt

	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	data.TanggalSpt = time.Now()
	// prevMonth := sh.BeginningOfPreviosMonth()
	// data.StartDate = datatypes.Date(prevMonth)
	// data.EndDate = datatypes.Date(sh.EndOfMonth(prevMonth))
	// data.DueDate = datatypes.Date(sh.EndOfMonth(time.Now()))

	err := tx.Create(&data).Error
	if err != nil {
		return nil, err
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.Spt
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Spt{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: data,
	}, nil
}

func GetDetail(id int) (any, error) {
	var data *m.Spt

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto) (any, error) {
	var data *m.Spt
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

func Delete(id int) (any, error) {
	//data spt
	var data *m.Spt
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	status := "no deletion"
	// data rekening
	var rekening *rm.Rekening
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}

	switch *rekening.Objek {
	case "01":
		var DataOp []mdsh.DetailSptHotel
		result := a.DB.Where(mdsh.DetailSptHotel{Spt_Id: uint64(id)}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(mdsh.DetailSptHotel{Spt_Id: uint64(id)}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				DataOp = nil
				status = "no deletion"
			}
		}

	case "02":
		var DataOp []mdsres.DetailSptResto
		result := a.DB.Where(mdsres.DetailSptResto{Spt_Id: uint64(id)}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(mdsres.DetailSptResto{Spt_Id: uint64(id)}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				DataOp = nil
				status = "no deletion"
			}
		}
	case "04":
		var DataOp []mdsrek.DetailSptReklame
		result := a.DB.Where(mdsrek.DetailSptReklame{Spt_Id: uint64(id)}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(mdsrek.DetailSptReklame{Spt_Id: uint64(id)}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				DataOp = nil
				status = "no deletion"
			}
		}

		var DataJ []mdjbr.JaminanBongkarReklame
		result = a.DB.Where(mdjbr.JaminanBongkarReklame{Spt_Id: uint64(id)}).Find(&DataJ)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataJ {
			result = a.DB.Where(mdsrek.DetailSptReklame{Spt_Id: uint64(id)}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				DataJ = nil
				status = "no deletion"
			}
		}
	case "07":
		var DataOp []mdsp.DetailSptParkir
		result := a.DB.Where(mdsp.DetailSptParkir{Spt_Id: uint64(id)}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(mdsp.DetailSptParkir{Spt_Id: uint64(id)}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				DataOp = nil
				status = "no deletion"
			}
		}
	case "08":
		var DataOp []mdsa.DetailSptAir
		result := a.DB.Where(mdsa.DetailSptAir{Spt_Id: uint64(id)}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(mdsa.DetailSptAir{Spt_Id: uint64(id)}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				DataOp = nil
				status = "no deletion"
			}
		}
	}

	result = a.DB.Delete(&data, id)
	status = "deleted"
	if result.RowsAffected == 0 {
		data = nil
		status = "no deletion"
	}
	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil

}
