package nilaiindividu

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm/clause"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/nilaiindividu"
	n "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajakpbb"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "nilaiindividu"

func Create(input m.CreateDto) (any, error) {
	var data m.NilaiIndividu
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	result := a.DB.Create(&data)
	if result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.NilaiIndividu
	var count int64
	var pagination gh.Pagination

	if input.BlokNOP != nil {
		var blokNOP = *input.BlokNOP
		// split 1-2.3 to 1 2 3
		var blokKode, noUrut, jenisOp string

		parts := strings.Split(blokNOP, "-")
		blokKode = parts[0]

		for _, part := range parts {
			subparts := strings.Split(part, ".")
			noUrut = subparts[0]
			for _, subpart := range subparts {
				jenisOp = subpart
			}
		}

		input.Blok_Kode = &blokKode
		input.NoUrut = &noUrut
		input.JenisOp = &jenisOp

		input.BlokNOP = nil
	}

	result := a.DB.
		Model(&m.NilaiIndividu{}).
		Scopes(gh.Filter(input)).
		Preload(clause.Associations).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	var records []m.ListResponse
	for _, item := range data {
		byNop := n.NopDetail{
			Provinsi_Kode:  input.Provinsi_Kode,
			Daerah_Kode:    input.Daerah_Kode,
			Kecamatan_Kode: input.Kecamatan_Kode,
			Kelurahan_Kode: input.Kelurahan_Kode,
			Blok_Kode:      input.Blok_Kode,
			NoUrut:         input.NoUrut,
			JenisOp:        input.JenisOp,
		}

		var temp m.ListResponse
		var tempBlok string
		if item.Blok_Kode != nil {
			tempBlok = *item.Blok_Kode
		}
		var tempNoUrut string
		if item.NoUrut != nil {
			tempNoUrut = *item.NoUrut
		}
		var tempJenisOp string
		if item.JenisOp != nil {
			tempJenisOp = *item.JenisOp
		}
		temp.BlokNOP = fmt.Sprintf("%s-%s.%s", tempBlok, tempNoUrut, tempJenisOp)
		temp.NilaiIndividu = item.NilaiIndividu

		// get objek pajak by nop
		var objekPajak objekpajakpbb.ObjekPajakPbb
		condition := nopSearcher(item)
		result := a.DB.Model(&objekpajakpbb.ObjekPajakPbb{}).Where(condition).First(&objekPajak)
		if result.RowsAffected != 0 {
			temp.LetakObjekPajak = objekPajak.Jalan

			// get wajib pajak pbb by id
			var m_wajibpajakpbb wajibpajakpbb.WajibPajakPbb
			result = a.DB.Model(&wajibpajakpbb.WajibPajakPbb{}).Where("id", objekPajak.WajibPajakPbb_Id).First(&m_wajibpajakpbb)
			if result.RowsAffected != 0 {
				temp.NamaWP = m_wajibpajakpbb.Nama
			}
		}

		// get objek pajak bangunan by nop
		var m_objekpajakbangunan objekpajakbangunan.ObjekPajakBangunan
		result = a.DB.Model(&objekpajakbangunan.ObjekPajakBangunan{}).Where(byNop).First(&m_objekpajakbangunan)
		if result.RowsAffected != 0 {
			temp.NoBng = m_objekpajakbangunan.NoBangunan
			temp.NilaiSistem = m_objekpajakbangunan.NilaiSistem
		}

		records = append(records, temp)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: records,
	}, nil
}

func GetDetail(id int) (interface{}, error) {
	var data *m.NilaiIndividu
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

func GetDetailByCode(id int) (interface{}, error) {
	var data *m.NilaiIndividu
	result := a.DB.Where("Kode", fmt.Sprint(id)).First(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto) (interface{}, error) {
	var data *m.NilaiIndividu

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.Copy(&data, input); err != nil {
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

func Delete(id int) (interface{}, error) {
	var data *m.NilaiIndividu
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
