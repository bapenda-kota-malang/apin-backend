package keberatan

import (
	"errors"
	"fmt"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/keberatan"
	mpst "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	msksk "github.com/bapenda-kota-malang/apin-backend/internal/models/sksk"

	skepkebpbb "github.com/bapenda-kota-malang/apin-backend/internal/services/keberatan/keputusankeberatanpbb"
	spstdetail "github.com/bapenda-kota-malang/apin-backend/internal/services/pelayanan/pstdetail"
	ssksk "github.com/bapenda-kota-malang/apin-backend/internal/services/sksk"
)

const source = "pembetulan keberatan"

func Create(input m.InputPembetulanSkKeberatan) (any, error) {
	var data m.PembetulanKeberatan

	// check nop inside table pst detail
	respPstDetail, err := spstdetail.GetByNoPelayanan(mpst.GetByNoPelayananPstDetailDto{
		TahunPelayanan:  input.TahunPelayanan,
		BundelPelayanan: input.BundelPelayanan,
		NoUrutPelayanan: input.NoUrutPelayanan,
	})
	if err != nil || (respPstDetail == nil && err == nil) {
		return nil, err
	}
	dataPstDetail := respPstDetail.(rp.OKSimple).Data.(mpst.PstDetail)

	// check data inside tabel keputusan keberatan pbb, use nop and no sk lama
	respKepKebPbb, err := skepkebpbb.GetByNopDanSk(m.FilterNopSkDtoKepKebPbb{
		PermohonanProvinsiID:  *dataPstDetail.PermohonanProvinsiID,
		PermohonanKotaID:      *dataPstDetail.PermohonanKotaID,
		PermohonanKecamatanID: *dataPstDetail.PermohonanKecamatanID,
		PermohonanKelurahanID: *dataPstDetail.PermohonanKelurahanID,
		PermohonanBlokID:      *dataPstDetail.PermohonanBlokID,
		NoUrutPemohon:         *dataPstDetail.NoUrutPemohon,
		PemohonJenisOPID:      *dataPstDetail.PemohonJenisOPID,
		JnsSK:                 *input.JenisSKLama,
		NoSK:                  *input.NoSkLama,
	})
	dataKepKebPbb := respKepKebPbb.(rp.OK).Data.(*m.KeputusanKeberatanPbb)
	if err != nil || dataKepKebPbb == nil {
		return nil, err
	}

	dataPemKeb := m.PembetulanKeberatan{
		Kanwil_Kd:                   *dataKepKebPbb.KanwilId,
		Kppbb_Kd:                    *dataKepKebPbb.KppbbId,
		TahunPelayanan:              input.TahunPelayanan,
		BundelPelayanan:             input.BundelPelayanan,
		NoUrutPelayanan:             input.NoUrutPelayanan,
		TahunPelayananKepKeberatan:  dataKepKebPbb.TahunPelayanan,
		BundelPelayananKepKeberatan: dataKepKebPbb.TahunPelayanan,
		NoUrutPelayananKepKeberatan: dataKepKebPbb.NoUrutPelayanan,
		ProvinsiPemohon_Kd:          dataKepKebPbb.PermohonanProvinsiID,
		DaerahPemohon_Kd:            dataKepKebPbb.PermohonanKotaID,
		KecamatanPemohon_Kd:         dataKepKebPbb.PermohonanKecamatanID,
		KelurahanPemohon_Kd:         dataKepKebPbb.PermohonanKelurahanID,
		BlokPemohon_Kd:              dataKepKebPbb.PermohonanBlokID,
		NoUrutPemohonan:             dataKepKebPbb.NoUrutPemohon,
		JenisOpPemohon_Kd:           dataKepKebPbb.PemohonJenisOPID,
		JenisSK:                     input.JenisSKBaru,
		NoSk:                        input.NoSkBaru,
		KelasTanah_Kd:               *dataKepKebPbb.KlsTanah_Kode,
		TahunAwalKelasTanah:         *dataKepKebPbb.KlsTanah_TahunAwal,
		KelasBangunan_Kd:            *dataKepKebPbb.KlsBangunan_Kode,
		TahunAwalKelasBangunan:      *dataKepKebPbb.KlsBangunan_TahunAwal,
		LuasBumiPembetulan:          float64(*dataKepKebPbb.LuasBumi),
		LuasBangunanPembetulan:      float64(*dataKepKebPbb.LuasBangunan),
		NjopBumiPembetulan:          float64(*dataKepKebPbb.NjopBumi),
		NjopBangunanPembetulan:      float64(*dataKepKebPbb.NjopBangunan),
		PbbPembetulan:               float64(*dataKepKebPbb.PBB),
	}

	err = a.DB.Transaction(func(tx *gorm.DB) error {
		// save new data pembetulan keberatan
		if result := tx.Create(&dataPemKeb); result.Error != nil {
			return result.Error
		}
		// update data keputusan keberatan pbb
		dataKepKebPbb.JnsSK = input.JenisSKBaru
		dataKepKebPbb.NoSK = input.NoSkBaru
		dataKepKebPbb.TahunPelayanan = &input.TahunPelayanan
		dataKepKebPbb.BundelPelayanan = &input.BundelPelayanan
		dataKepKebPbb.NoUrutPelayanan = &input.NoUrutPelayanan
		if result := tx.Save(&dataKepKebPbb); result.Error != nil {
			return result.Error
		}

		// create new sksk data
		var skSkCreateDto msksk.CreateDto
		if err := sc.Copy(&skSkCreateDto, input); err != nil {
			return err
		}
		skSkCreateDto.PermohonanId = dataKepKebPbb.PermohonanId
		skSkCreateDto.KanwilId = dataKepKebPbb.KanwilId
		skSkCreateDto.KppbbId = dataKepKebPbb.KppbbId
		skSkCreateDto.JnsSK = input.JenisSKBaru
		skSkCreateDto.NoSK = input.NoSkBaru

		_, err := ssksk.Create(skSkCreateDto, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("error pembetulan sk keberatan: %s", err), data)
	}
	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDtoPemKeb) (any, error) {
	var data []m.PembetulanKeberatan
	var count int64
	var pagination gh.Pagination

	query := a.DB.
		Model(&m.PembetulanKeberatan{}).
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

func GetDetail(id int) (any, error) {
	var data *m.PembetulanKeberatan
	result := a.DB.Model(&m.PembetulanKeberatan{}).
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

func Update(id int, input m.UpdateDtoPemKeb, tx *gorm.DB) (interface{}, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.PembetulanKeberatan

	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.CopyWithOption(&data, input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("gagal memperbarui data: %s", result.Error), data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (interface{}, error) {
	var data *m.PembetulanKeberatan
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
