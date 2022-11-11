package tbp

import (
	"errors"
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/tbp"
	sdt "github.com/bapenda-kota-malang/apin-backend/internal/services/tbp/detailtbp"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "tbp"

func Create(input m.CreateDto, user_Id uint64) (any, error) {
	var dataTbp m.Tbp
	var dataDetailTbp m.DetailTbpCreateDto
	var respDataDetailTbp interface{}
	var resp t.II
	// var dataSpt ms.Spt
	// var dataSptTotal float64
	// var dataSptJumlahPajak float64
	// var dataNominalBayar float64
	// var dataRincianTbpDenda float64

	// data tbp
	if err := sc.Copy(&dataTbp, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload tbp", dataTbp)
	}

	// data rincian tbp
	if err := sc.Copy(&dataDetailTbp, &input.DetailTbp); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload rincian tbp", dataDetailTbp)
	}

	// perhitungan rincian tbp dihitung di fe, user hanya mengirim nominal bayar
	// // cek data spt
	// result := a.DB.Where(ms.Spt{Id: *input.RincianTbp.Spt_Id}).First(&dataSpt)
	// if result.Error != nil {
	// 	return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data spt", dataSpt)
	// }

	// dataSptTotal = *dataSpt.Total
	// dataSptJumlahPajak = *dataSpt.JumlahPajak
	// dataNominalBayar = *input.RincianTbp.NominalBayar
	// dataRincianTbpDenda = *dataRincianTbp.Denda

	// // perhitungan untuk data tbp total
	// if *input.IsKetetapan {
	// 	if dataSptTotal < dataNominalBayar {
	// 		fmt.Println("perlu diisi")
	// 	}
	// 	tmpTotal := dataSptTotal + dataRincianTbpDenda
	// 	dataTbp.Total = &tmpTotal
	// } else {
	// 	if dataSptJumlahPajak < dataNominalBayar {
	// 		fmt.Println("perlu diisi")
	// 	}
	// 	tmpTotal := dataSptJumlahPajak + dataRincianTbpDenda
	// 	dataTbp.Total = &tmpTotal
	// }

	tmpIsCancelled := false
	// static value
	tmpNomor := generateNomor()
	dataTbp.TbpNumber = &tmpNomor
	dataTbp.TanggalBayar = th.TimeNow()
	dataTbp.CreatedBy_User_Id = &user_Id
	dataDetailTbp.Waktu_Detail_Tb = parseCurrentTime()
	dataTbp.IsCancelled = &tmpIsCancelled

	err := a.DB.Transaction(func(tx *gorm.DB) error {

		// create data tbp
		err := tx.Create(&dataTbp).Error
		if err != nil {
			return err
		}

		// create data rincian tbp
		resultDetailTbp, err := sdt.Create(dataDetailTbp, dataTbp.Id, tx)
		if err != nil {
			return err
		}
		respDataDetailTbp = resultDetailTbp

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataTbp)
	}

	resp = t.II{
		"tbp":        dataTbp,
		"rincianTbp": respDataDetailTbp.(rp.OKSimple).Data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.Tbp
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Tbp{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("ObjekPajak.Kecamatan").
		Preload("ObjekPajak.Kelurahan").
		Preload("DetailTbps.Spt").
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

func GetDetail(tbp_id int) (any, error) {
	var data *m.Tbp
	err := a.DB.
		Model(&m.Tbp{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("ObjekPajak.Kecamatan").
		Preload("ObjekPajak.Kelurahan").
		Preload("DetailTbps.Spt").
		First(&data, tbp_id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return rp.OKSimple{
		Data: data,
	}, err
}

func Cancel(input m.CancelDto, tbp_id uint64) (any, error) {
	var data m.Tbp
	result := a.DB.First(&data, tbp_id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tbp tidak dapat ditemukan")
	}

	if *data.IsCancelled {
		return nil, errors.New("data sudah dicancel")
	}

	// copy flag cancel
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "cancel-tbp", source, "failed", "gagal mengambil data payload tbp", input)
	}

	data.CancelledDate = th.TimeNow()

	if result := a.DB.Save(&data); result.Error != nil {
		return nil, errors.New("gagal menyimpan data tbp: " + result.Error.Error())
	}
	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(tbp_id uint64) (any, error) {
	var status string = "deleted"
	var data *m.Tbp

	// cek data tbp
	result := a.DB.First(&data, tbp_id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tbp tidak dapat ditemukan")
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// delete rincian tbp
		err := sdt.Delete(data.Id, tx)
		if err != nil {
			status = "no deletion"
			return err
		}

		// delete tbp
		result = tx.Delete(&data, tbp_id)
		if result.RowsAffected == 0 {
			return errors.New("tidak dapat menghapus data tbp: " + result.Error.Error())
		}
		return nil

	})
	if err != nil {
		return sh.SetError("request", "delete-data", source, "failed", "gagal menghapus data: "+err.Error(), data)
	}

	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, user_id uint64) (any, error) {
	var dataTbp *m.Tbp
	var dataDetailTbp m.DetailTbpUpdateDto
	var respDataDetailTbp interface{}
	var resp t.II

	result := a.DB.First(&dataTbp, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tbp tidak dapat ditemukan")
	}

	if err := sc.Copy(&dataTbp, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload tbp", input)
	}

	if err := sc.Copy(&dataDetailTbp, &input.DetailTbp); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload rincian tbp", input.DetailTbp)
	}

	dataTbp.TanggalBayar = th.TimeNow()
	dataTbp.CreatedBy_User_Id = &user_id

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// update tbp
		if result := tx.Save(&dataTbp); result.Error != nil {
			return errors.New("gagal menyimpan data tbp")
		}

		resultDetailTbp, err := sdt.Update(dataDetailTbp, dataTbp.Id, tx)
		if err != nil {
			return err
		}
		respDataDetailTbp = resultDetailTbp

		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataTbp)
	}

	resp = t.II{
		"affected":   strconv.Itoa(int(result.RowsAffected)),
		"rincianTbp": respDataDetailTbp.(rp.OKSimple).Data,
		"tbp":        dataTbp,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}
