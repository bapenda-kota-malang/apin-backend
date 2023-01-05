package sspd

import (
	"errors"
	"strconv"

	mn "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
	ssd "github.com/bapenda-kota-malang/apin-backend/internal/services/sspd/sspddetail"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/bapenda-kota-malang/apin-backend/pkg/slicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "sspd"

func Create(input m.CreateDto, user_Id uint64) (any, error) {
	var dataSspd m.Sspd
	var dataNpwpd mn.Npwpd
	var dataSspdDetail m.SspdDetailCreateDto
	var respDataSspdDetail interface{}
	var resp t.II

	// data sspd
	if err := sc.Copy(&dataSspd, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload sspd", dataSspd)
	}

	// data sspd detail
	if err := sc.Copy(&dataSspdDetail, &input.SspdDetail); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload sspd detail", dataSspdDetail)
	}

	// data Npwpd
	result := a.DB.Where(mn.Npwpd{Npwpd: input.Npwpd_Npwpd}).First(&dataNpwpd)
	if result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data npwpd", dataNpwpd)
	}

	tmpIsCancelled := false
	// static value
	if input.TanggalBayar == nil {
		dataSspd.TanggalBayar = th.TimeNow()
	} else {
		dataSspd.TanggalBayar = th.ParseTime(input.TanggalBayar)
	}

	dataSspd.NomorTahun = generateNomorTahun(*dataSspd.TanggalBayar)
	dataSspd.NomorUrut = generateNomorUrut()
	dataSspd.NomorOutput = generateNomorOutput(dataSspd.NomorTahun, dataSspd.NomorUrut)
	dataSspd.CreatedBy_User_Id = &user_Id
	dataSspdDetail.WaktuSspdDetail = parseCurrentTime()
	dataSspd.IsCancelled = &tmpIsCancelled
	dataSspd.ObjekPajak_Id = &dataNpwpd.ObjekPajak_Id

	err := a.DB.Transaction(func(tx *gorm.DB) error {

		// create data sspd
		err := tx.Create(&dataSspd).Error
		if err != nil {
			return err
		}

		// create data sspd detail
		resultSspdDetail, err := ssd.Create(dataSspdDetail, dataSspd.Id, tx)
		if err != nil {
			return err
		}
		respDataSspdDetail = resultSspdDetail

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataSspd)
	}

	resp = t.II{
		"sspd":       dataSspd,
		"sspdDetail": respDataSspdDetail.(rp.OKSimple).Data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func GetList(input m.FilterDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.Sspd
	var count int64

	var pagination gh.Pagination
	result := tx.
		Model(&m.Sspd{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("ObjekPajak.Kecamatan").
		Preload("ObjekPajak.Kelurahan").
		Preload("Npwpd.Rekening").
		Preload("SspdDetails.Spt").
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

func GetListWp(userId int, npwpd string, input m.FilterWpDto) (any, error) {
	var data []m.ListLogPayment
	var count int64

	// load spt data
	var sptData []spt.Spt
	var pagination gh.Pagination
	sptBase := a.DB.
		Model(&spt.Spt{}).
		Joins("JOIN \"Npwpd\" ON \"Npwpd\".\"Id\" = \"Spt\".\"Npwpd_Id\" AND \"Npwpd\".\"Npwpd\" = ? AND \"Npwpd\".\"User_Id\" = ?", npwpd, userId)
	if input.PeriodeAwal != nil {
		sptBase = sptBase.Where("\"Spt\".\"PeriodeAwal\" >= ?", input.PeriodeAwal)
	}
	if input.PeriodeAkhir != nil {
		sptBase = sptBase.Where("\"Spt\".\"PeriodeAkhir\" <= ?", input.PeriodeAkhir)
	}
	sptBase = sptBase.Scopes(gh.Paginate(input, &pagination)).Order("\"PeriodeAkhir\" DESC")

	if sptBase.Find(&sptData).Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", sptData)
	}

	if len(sptData) == 0 {
		return rp.OK{
			Meta: t.IS{
				"totalCount":   strconv.Itoa(0),
				"currentCount": strconv.Itoa(0),
				"page":         strconv.Itoa(pagination.Page),
				"pageSize":     strconv.Itoa(pagination.PageSize),
			},
			Data: sptData,
		}, nil
	}

	ids := make(map[string][]string)
	ids["spt"] = []string{}
	for _, v := range sptData {
		ids["spt"] = append(ids["spt"], v.Id.String())
	}

	// load sspd detail bridging data
	var sspdDetailData []m.SspdDetail
	querySspdDetail := a.DB.
		Model(&m.SspdDetail{}).
		Where("\"Spt_Id\" IN ?", ids["spt"])

	if querySspdDetail.Find(&sspdDetailData).Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", sptData)
	}

	if len(sspdDetailData) == 0 {
		return rp.OK{
			Meta: t.IS{
				"totalCount":   strconv.Itoa(0),
				"currentCount": strconv.Itoa(0),
				"page":         strconv.Itoa(pagination.Page),
				"pageSize":     strconv.Itoa(pagination.PageSize),
			},
			Data: sspdDetailData,
		}, nil
	}

	ids["sspdDetail"] = []string{}
	for _, v := range sspdDetailData {
		str := strconv.Itoa(int(*v.Sspd_Id))
		if !slicehelper.StringInSlice(str, ids["sspdDetail"]) {
			ids["sspdDetail"] = append(ids["sspdDetail"], str)
		}
	}

	// load sspd data
	var sspdData []m.Sspd
	result := a.DB.
		Model(&m.Sspd{}).
		Joins("JOIN \"Npwpd\" ON \"Npwpd\".\"Npwpd\" = ? AND \"Npwpd\".\"User_Id\" = ?", npwpd, userId).
		Where("\"Sspd\".\"Id\" IN ?", ids["sspdDetail"]).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&sspdData)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", sspdData)
	}

	// append to list data
	for _, vSpt := range sptData {
		tmp := m.ListLogPayment{}
		for _, vSspdDetail := range sspdDetailData {
			if *vSspdDetail.Spt_Id != vSpt.Id {
				continue
			}
			for _, vSspd := range sspdData {
				if vSspd.Id != *vSspdDetail.Sspd_Id {
					continue
				}
				if err := sc.Copy(&tmp, &vSspd); err != nil {
					return sh.SetError("request", "get-list-log-payment", source, "failed", "gagal menambahkan list data", data)
				}
			}
		}

		tmp.Spt = vSpt
		data = append(data, tmp)
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

func GetDetail(sspd_id int) (any, error) {
	var data *m.Sspd
	err := a.DB.
		Model(&m.Sspd{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("ObjekPajak.Kecamatan").
		Preload("ObjekPajak.Kelurahan").
		Preload("SspdDetails.Spt").
		First(&data, sspd_id).Error
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

func Cancel(input m.CancelDto, sspd_id uint64) (any, error) {
	var data m.Sspd
	result := a.DB.First(&data, sspd_id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data sspd tidak dapat ditemukan")
	}

	if *data.IsCancelled {
		return nil, errors.New("data sudah dicancel")
	}

	// copy flag cancel
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "cancel-sspd", source, "failed", "gagal mengambil data payload sspd", input)
	}

	data.CancelledDate = th.TimeNow()

	if result := a.DB.Save(&data); result.Error != nil {
		return nil, errors.New("gagal menyimpan data sspd: " + result.Error.Error())
	}
	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(sspd_id uint64) (any, error) {
	var status string = "deleted"
	var data *m.Sspd

	// cek data sspd
	result := a.DB.First(&data, sspd_id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data sspd tidak dapat ditemukan")
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// delete sspd detail
		err := ssd.Delete(data.Id, tx)
		if err != nil {
			status = "no deletion"
			return err
		}

		// delete sspd
		result = tx.Delete(&data, sspd_id)
		if result.RowsAffected == 0 {
			return errors.New("tidak dapat menghapus data sspd: " + result.Error.Error())
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
	var dataSspd *m.Sspd
	var dataSspdDetail m.SspdDetailUpdateDto
	var respDataSspdDetail interface{}
	var resp t.II

	result := a.DB.First(&dataSspd, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data sspd tidak dapat ditemukan")
	}
	if err := sc.Copy(&dataSspd, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload sspd", input)
	}

	if err := sc.Copy(&dataSspdDetail, &input.SspdDetail); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload sspd detail", input.SspdDetail)
	}

	if input.TanggalBayar == nil {
		dataSspd.TanggalBayar = th.TimeNow()
	} else {
		dataSspd.TanggalBayar = th.ParseTime(input.TanggalBayar)
	}
	dataSspd.CreatedBy_User_Id = &user_id
	dataSspdDetail.WaktuSspdDetail = parseCurrentTime()
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// update sspd
		if result := tx.Save(&dataSspd); result.Error != nil {
			return errors.New("gagal menyimpan data sspd")
		}

		resultSspdDetail, err := ssd.Update(dataSspdDetail, dataSspd.Id, tx)
		if err != nil {
			return err
		}
		respDataSspdDetail = resultSspdDetail

		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataSspd)
	}

	resp = t.II{
		"affected":   strconv.Itoa(int(result.RowsAffected)),
		"sspdDetail": respDataSspdDetail.(rp.OKSimple).Data,
		"sspd":       dataSspd,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}
