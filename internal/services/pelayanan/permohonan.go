// service
package permohonan

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
)

// /// Private funcs start here
const source = "permohonan"

///// Exported funcs start here

func Create(input m.PermohonanRequestDto) (any, error) {
	var (
		err                   error
		permohonanBaru        *m.PstDataOPBaru
		permohonanDetail      *m.PstDetail
		permohonanPengurangan *m.PstPermohonanPengurangan
	)

	noUrut := GetNoUrut(input)
	data := input.SetDataPermohonanRequestDtoTransformer(&noUrut)
	noPelayanan := data.GetDataPermohonanNoPelayanan()
	data.NoPelayanan = &noPelayanan

	err = a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data permohonan ke db satu if karena result dipakai sekali, +error
		if result := tx.Create(&data); result.Error != nil {
			return errors.New("penyimpanan data permohonan gagal")
		}

		permohonanBaru, permohonanDetail, permohonanPengurangan = data.SetDataPermohonanTransformer(input)
		if *data.BundelPelayanan == m.JenisPelayanan[0] {
			if result := tx.Create(&permohonanBaru); result.Error != nil {
				return errors.New("penyimpanan data permohonan baru gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[1] {
			if result := tx.Create(&permohonanDetail); result.Error != nil {
				return errors.New("penyimpanan data permohonan detail gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[2] {
			if result := tx.Create(&permohonanPengurangan); result.Error != nil {
				return errors.New("penyimpanan data permohonan pengurangan gagal")
			}
		}

		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
	}

	return rp.OKSimple{Data: t.II{
		"permohonan":            data,
		"permohonanBaru":        permohonanBaru,
		"permohonanDetail":      permohonanDetail,
		"permohonanPengurangan": permohonanPengurangan,
		"noPelayanan":           noPelayanan,
	}}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.PstPermohonan
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.PstPermohonan{}).
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

func GetNoUrut(input m.PermohonanRequestDto) string {
	var (
		checkdata *m.PstPermohonan
		noUrut    string = "001"
	)

	year := strconv.Itoa(time.Now().Year())
	_ = a.DB.Where("tahunPelayanan", year).
		Where("bundlePelayanan", input.JenisPelayanan).
		Order("noUrut desc").
		First(&checkdata)
	if checkdata.Id != 0 {
		tempNoUrut, _ := strconv.Atoi(*checkdata.NoUrutPelayanan)
		tempNoUrut = tempNoUrut + 1
		noUrut = sh.FixedLengthString(3, fmt.Sprint(tempNoUrut))
	}

	return noUrut
}

func getDataNOP(nop *string, jp *string) (*m.PermohonanNOPResponse, error) {
	var (
		permohonanBaru        *m.PstDataOPBaru
		permohonanDetail      *m.PstDetail
		permohonanPengurangan *m.PstPermohonanPengurangan
	)
	// nop := m.DecodeNOPPermohonan(input.NOP)
	if *jp == m.JenisPelayanan[0] {
		result := a.DB.First(&permohonanBaru, nop)
		if result.RowsAffected == 0 {
			return nil, nil
		} else if result.Error != nil {
			return nil, errors.New("gagal mengambil data")
		}

		return permohonanBaru.GetNOPResponse(), nil
	} else if *jp == m.JenisPelayanan[1] {
		result := a.DB.First(&permohonanDetail, nop)
		if result.RowsAffected == 0 {
			return nil, nil
		} else if result.Error != nil {
			return nil, errors.New("gagal mengambil data")
		}

		return permohonanDetail.GetNOPResponse(), nil
	} else if *jp == m.JenisPelayanan[2] {
		result := a.DB.First(&permohonanPengurangan, nop)
		if result.RowsAffected == 0 {
			return nil, nil
		} else if result.Error != nil {
			return nil, errors.New("gagal mengambil data")
		}

		return permohonanPengurangan.GetNOPResponse(), nil
	}
	return nil, nil
}

func GetNoPelayanan(jt *string) (interface{}, error) {
	y := time.Now().Year()
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(9999-0) + 0
	tempV := sh.FixedLengthString(4, fmt.Sprint(v))
	result := fmt.Sprint(y) + *jt + tempV

	return rp.OKSimple{Data: result}, nil
}

func GetStatusNOP(nop *string, jp *string) (interface{}, error) {
	result, err := getDataNOP(nop, jp)
	if err != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", nop)
	}

	return rp.OKSimple{Data: result}, nil
}

func GetDetailNOP(nop *string, jp *string) (interface{}, error) {
	result, err := getDataNOP(nop, jp)
	if err != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", nop)
	}

	return rp.OKSimple{Data: result}, nil
}

func GetDetail(id int) (interface{}, error) {
	var (
		data                  *m.PstPermohonan
		permohonanBaru        *m.PstDataOPBaru
		permohonanDetail      *m.PstDetail
		permohonanPengurangan *m.PstPermohonanPengurangan
	)
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", data)
	}

	if *data.BundelPelayanan == m.JenisPelayanan[0] {
		if result := a.DB.Where("PermohonanId", data.Id).First(&permohonanBaru); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data NOP Baru", permohonanBaru)
		}
	} else if *data.BundelPelayanan == m.JenisPelayanan[1] {
		if result := a.DB.Where("PermohonanId", data.Id).First(&permohonanDetail); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data NOP Detail", permohonanDetail)
		}
	} else if *data.BundelPelayanan == m.JenisPelayanan[2] {
		if result := a.DB.Where("PermohonanId", data.Id).First(&permohonanPengurangan); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data NOP pengurangan", permohonanPengurangan)
		}
	}

	data.PstDataOPBaru = permohonanBaru
	data.PstDetail = permohonanDetail
	data.PstPermohonanPengurangan = permohonanPengurangan

	return rp.OKSimple{Data: data}, nil
}

func Update(id int, input m.PermohonanRequestDto) (interface{}, error) {
	var (
		data                  *m.PstPermohonan
		permohonanBaru        *m.PstDataOPBaru
		permohonanDetail      *m.PstDetail
		permohonanPengurangan *m.PstPermohonanPengurangan
	)

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "data permohonan tidak dapat ditemukan", input)
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload untuk permohonan", input)
	}

	rowsAffected := 0
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		if result := a.DB.Save(&data); result.Error != nil {
			return errors.New("gagal menyimpan data permohonan")
		}
		if result.RowsAffected > 0 {
			rowsAffected++
		}

		permohonanBaru, permohonanDetail, permohonanPengurangan = data.SetDataPermohonanTransformer(input)
		if *data.BundelPelayanan == m.JenisPelayanan[0] {
			if result := tx.Where("PermohonanId", data.Id).Updates(&permohonanBaru); result.Error != nil {
				return errors.New("penyimpanan data permohonan baru gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[1] {
			if result := tx.Where("PermohonanId", data.Id).Updates(&permohonanDetail); result.Error != nil {
				return errors.New("penyimpanan data permohonan detail gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[2] {
			if result := tx.Where("PermohonanId", data.Id).Updates(&permohonanPengurangan); result.Error != nil {
				return errors.New("penyimpanan data permohonan pengurangan gagal")
			}
		}

		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", err.Error(), input)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(rowsAffected),
		},
		Data: t.II{
			"permohonan":            data,
			"permohonanBaru":        permohonanBaru,
			"permohonanDetail":      permohonanDetail,
			"permohonanPengurangan": permohonanPengurangan,
		},
	}, nil
}

func Delete(id int) (interface{}, error) {
	var (
		data                  *m.PstPermohonan
		permohonanBaru        *m.PstDataOPBaru
		permohonanDetail      *m.PstDetail
		permohonanPengurangan *m.PstPermohonanPengurangan
	)

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	result = a.DB.Delete(&data, id)
	status := "deleted"
	if result.RowsAffected == 0 {
		data = nil
		status = "no deletion"
	} else {
		input := data.GetDataPermohonanRequestDtoTransformer(nil)
		permohonanBaru, permohonanDetail, permohonanPengurangan = data.SetDataPermohonanTransformer(input)
		if *data.BundelPelayanan == m.JenisPelayanan[0] {
			_ = a.DB.Delete(&permohonanBaru, int(*permohonanBaru.PermohonanId) == id)
		} else if *data.BundelPelayanan == m.JenisPelayanan[1] {
			_ = a.DB.Delete(&permohonanDetail, int(*permohonanDetail.PermohonanId) == id)
		} else if *data.BundelPelayanan == m.JenisPelayanan[2] {
			_ = a.DB.Delete(&permohonanPengurangan, int(*permohonanPengurangan.PermohonanId) == id)
		}
	}

	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil
}
