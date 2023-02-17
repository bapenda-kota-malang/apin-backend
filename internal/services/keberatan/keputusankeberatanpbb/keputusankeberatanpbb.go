package keputusankeberatanpbb

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/keberatan"
	msksk "github.com/bapenda-kota-malang/apin-backend/internal/models/sksk"
	msppt "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	ssksk "github.com/bapenda-kota-malang/apin-backend/internal/services/sksk"
	ssppt "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "KeputusanKeberatanPbb"

func updateSpptSp(data m.KeputusanKeberatanPbb, tx *gorm.DB) error {
	if data.JnsKeputusan != m.JnsKeputusanDiterima {
		return nil
	}

	// get data sppt
	var dataSppt msppt.Sppt
	resDataSspt, err := ssppt.GetByNop(
		&data.PermohonanProvinsiID,
		&data.PermohonanKotaID,
		&data.PermohonanKecamatanID,
		&data.PermohonanKelurahanID,
		&data.PermohonanBlokID,
		&data.NoUrutPemohon,
		&data.PemohonJenisOPID,
		data.TahunPelayanan,
	)
	if err != nil {
		return err
	}
	dataSppt = resDataSspt.(rp.OKSimple).Data.(msppt.Sppt)

	_, _, err = ssppt.SpPenilaian(dataSppt, tx)
	if err != nil {
		return err
	}

	// get min max value referensi buku
	minValue, maxValue, err := ssppt.MinMaxValueReferensiBuku()
	if err != nil {
		return err
	}

	// input dto provide tahun, min, max value, tglJatuhTempo, tglTerbit
	dateFormat := time.Now().Format("2006-01-02")
	_, err = ssppt.SpPenetapan(
		msppt.PenetapanMassalDto{
			Tahun:          *dataSppt.TahunPajakskp_sppt,
			TglJatuhTempo1: dateFormat,
			TglJatuhTempo2: dateFormat,
			TglJatuhTempo3: dateFormat,
			TglJatuhTempo4: dateFormat,
			TglJatuhTempo5: dateFormat,
			TglTerbit1:     dateFormat,
			TglTerbit2:     dateFormat,
			TglTerbit3:     dateFormat,
			TglTerbit4:     dateFormat,
			TglTerbit5:     dateFormat,
			BukuMin:        minValue,
			BukuMax:        maxValue,
		},
		dataSppt,
		tx,
	)
	if err != nil {
		return err
	}

	return nil
}

func Create(input m.CreateDtoKepKebPbb, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.KeputusanKeberatanPbb
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	err := tx.Transaction(func(txChild *gorm.DB) error {
		result := txChild.Create(&data)
		if result.Error != nil {
			return result.Error
		}
		var createDtoSksk msksk.CreateDto
		if err := sc.Copy(&createDtoSksk, data); err != nil {
			return err
		}
		if _, err := ssksk.Create(createDtoSksk, txChild); err != nil {
			return err
		}
		if err := updateSpptSp(data, txChild); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data: %s", err), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDtoKepKebPbb) (interface{}, error) {
	var data []m.KeputusanKeberatanPbb
	var count int64
	var pagination gh.Pagination

	query := a.DB.
		Model(&m.KeputusanKeberatanPbb{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination))
	result := query.Find(&data)
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

func GetDetail(id int) (interface{}, error) {
	var data *m.KeputusanKeberatanPbb
	result := a.DB.Model(&m.KeputusanKeberatanPbb{}).
		First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDtoKepKebPbb, tx *gorm.DB) (interface{}, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.KeputusanKeberatanPbb

	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.CopyWithOption(&data, input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}
	err := tx.Transaction(func(txChild *gorm.DB) error {
		if result := txChild.Save(&data); result.Error != nil {
			return result.Error
		}
		if err := updateSpptSp(*data, txChild); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("gagal memperbarui data: %s", err), data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (interface{}, error) {
	var data *m.KeputusanKeberatanPbb
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	result = a.DB.Delete(&data, id)
	status := "deleted"
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
