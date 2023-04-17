package sejarah

import (
	msppt "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt/sejarah"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajakpbb"
	ssppt "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"strconv"
)

const source = "sejarah sppt"

func Create(input m.CreateDto) (any, error) {
	var data m.SpptSejarah

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := a.DB.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.SpptSejarah
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.SpptSejarah{}).
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

func Delete(id int) (any, error) {
	var data *m.SpptSejarah
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
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

func GetSejarahSppt(input m.FilterDto) (any, error) {
	var data []m.SpptSejarah
	var count int64

	filter := m.SpptSejarah{
		Propinsi_Id:        input.Propinsi_Id,
		Dati2_Id:           input.Dati2_Id,
		Kecamatan_Id:       input.Kecamatan_Id,
		Keluarahan_Id:      input.Keluarahan_Id,
		Blok_Id:            input.Blok_Id,
		NoUrut:             input.NoUrut,
		JenisOP_Id:         input.JenisOP_Id,
		TahunPajakskp_sppt: input.TahunPajakskp_sppt,
	}

	var pagination gh.Pagination
	result := a.DB.Debug().
		Model(&m.SpptSejarah{}).
		Where(filter).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	if result.RowsAffected == 0 {
		return sh.SetError("request", "get-data-list", source, "failed", "data tidak ditemukan", data)
	}

	sppt := msppt.Sppt{
		Propinsi_Id:        input.Propinsi_Id,
		Dati2_Id:           input.Dati2_Id,
		Kecamatan_Id:       input.Kecamatan_Id,
		Keluarahan_Id:      input.Keluarahan_Id,
		Blok_Id:            input.Blok_Id,
		NoUrut:             input.NoUrut,
		JenisOP_Id:         input.JenisOP_Id,
		TahunPajakskp_sppt: input.TahunPajakskp_sppt,
	}

	//get objek pajak pbb by nop
	objekPajakPbb, err := ssppt.GetObjekPajakPBBDetail(sppt)
	if err != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data objek pajak pbb", data)
	}

	// get wajib pajak pbb by id
	var m_wajibPajakPbb wajibpajakpbb.WajibPajakPbb
	if objekPajakPbb.WajibPajakPbb_Id != nil {
		a.DB.Model(&wajibpajakpbb.WajibPajakPbb{}).Where(`"Id" = ?`, *objekPajakPbb.WajibPajakPbb_Id).First(&m_wajibPajakPbb)
	}

	var dataRecords []m.SejarahSpptResponse
	for _, v := range data {
		dataRecords = append(dataRecords, m.SejarahSpptResponse{
			NamaWP_sppt:            v.NamaWP_sppt,
			JalanWPskp_sppt:        v.JalanWPskp_sppt,
			BlokKavNoWP_sppt:       v.BlokKavNoWP_sppt,
			RtWP_sppt:              v.RtWP_sppt,
			RwWP_sppt:              v.RwWP_sppt,
			KelurahanWP_sppt:       v.KelurahanWP_sppt,
			KotaWP_sppt:            v.KotaWP_sppt,
			PosWPsppt_Id:           v.PosWPsppt_Id,
			Npwp_sppt:              v.Npwp_sppt,
			NoPersil_sppt:          v.NoPersil_sppt,
			TP_Id:                  v.TP_Id,
			KelasBangunan_Id:       v.KelasBangunan_Id,
			KelasTanah_Id:          v.KelasTanah_Id,
			TanggalJatuhTempo_sppt: v.TanggalJatuhTempo_sppt,
			LuasBumi_sppt:          v.LuasBumi_sppt,
			LuasBangunan_sppt:      v.LuasBangunan_sppt,
			NJOPBumi_sppt:          v.NJOPBumi_sppt,
			NJOPBangunan_sppt:      v.NJOPBangunan_sppt,
			NJOP_sppt:              v.NJOP_sppt,
			NJOPTKP_sppt:           v.NJOPTKP_sppt,
			NJKPskp_sppt:           v.NJKPskp_sppt,
			PBBterhutang_sppt:      v.PBBterhutang_sppt,
			Faktorpengurangan_sppt: v.Faktorpengurangan_sppt,
			PBBygHarusDibayar_sppt: v.PBBygHarusDibayar_sppt,
			StatusPembayaran_sppt:  v.StatusPembayaran_sppt,
			TanggalTerbit_sppt:     v.TanggalTerbit_sppt,
			TanggalCetak_sppt:      v.TanggalCetak_sppt,
			NIPPencetakan_sppt:     v.NIPPencetakan_sppt,
		})
	}
	//
	// response
	records := m.ListSejarahSpptResponse{
		NamaWajibPajak:  m_wajibPajakPbb.Nama,
		JalanObjekPajak: objekPajakPbb.Jalan,
		JalanWajibPajak: m_wajibPajakPbb.Jalan,
		List:            dataRecords,
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
