package sptpd

import (
	"strconv"
	"time"

	area "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// this function return list data bphtb with specific status either 00, 01, 02
func GetListPpat(input m.RequestSptpd) (any, error) {
	var data []m.BphtbSptpd
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.BphtbSptpd{}).
		Where("\"Status\" IN ?", []string{"00", "01", "02"}).
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

// this function return list data bphtb with specific status either 00, 01, 02
func GetDetailPpat(id int) (any, error) {
	var (
		model    *m.BphtbSptpd
		lampiran *m.Lampiran

		provinsi  *area.Provinsi
		kota      *area.Daerah
		kecamatan *area.Kecamatan
		kelurahan *area.Kelurahan
		// blok      *area.Blok
	)

	result := a.DB.Where("\"Status\" IN ?", []string{"00", "01", "02"}).First(&model, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", model)
	}

	data := new(m.ResponseSptpd)
	if err := sc.Copy(&data, &model); err != nil {
		return sh.SetError("request", "copy-data-detail", source, "failed", "gagal menyalin data", model)
	}

	// set OP Data
	if data.PermohonanProvinsiID != nil {
		a.DB.Where("Kode", *data.PermohonanProvinsiID).First(&provinsi)
		a.DB.Where("Kode", *data.PermohonanProvinsiID+*data.PermohonanKotaID).First(&kota)
		a.DB.Where("Kode", *data.PermohonanProvinsiID+*data.PermohonanKotaID+*data.PermohonanKecamatanID).First(&kecamatan)
		a.DB.Where("Kode", *data.PermohonanProvinsiID+*data.PermohonanKotaID+*data.PermohonanKecamatanID+*data.PermohonanKelurahanID).First(&kelurahan)
		data.OPProvinsi = &provinsi.Nama
		data.OPKota = &kota.Nama
		data.OPKecamatan = &kecamatan.Nama
		data.OPKelurahan = &kelurahan.Nama
	}

	// Set WP Data
	if data.Provinsi_Id != nil {
		a.DB.Where("Kode", *data.Provinsi_Id).First(&provinsi)
		a.DB.Where("Kode", *data.Provinsi_Id+*data.Kabupaten_id).First(&kota)
		a.DB.Where("Kode", *data.Provinsi_Id+*data.Kabupaten_id+*data.Kecamatan_Id).First(&kecamatan)
		a.DB.Where("Kode", *data.Provinsi_Id+*data.Kabupaten_id+*data.Kecamatan_Id+*data.Kelurahan_Id).First(&kelurahan)
		data.Provinsi_wp = &provinsi.Nama
		data.Kabupaten_wp = &kota.Nama
		data.Kecamatan_wp = &kecamatan.Nama
		data.Kelurahan_wp = &kelurahan.Nama
	}

	if data.NoDokumen != nil {
		a.DB.Where("NoSspd", model.NoDokumen).First(&lampiran)
		data.DataLampiran = lampiran
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

// this function change status to either 01 or 02 based from input, update ppat_id and verifikasiPpatAt
func VerifyPpat(id int, input m.VerifyPpatDto, userId int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.BphtbSptpd

	result := a.DB.Where("\"Status\" IN ?", []string{"00", "01", "02"}).First(&data, id)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "data tidak ditemukan", nil)
	} else if result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mencari data existing: "+result.Error.Error(), nil)
	}

	// prepare data
	statusString := ""
	switch input.Status {
	case 1:
		statusString = "01"
	case 2:
		statusString = "02"
	}
	now := time.Now()

	// change field
	data.Status = &statusString
	data.VerifikasiPpatAt = &now

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengubah data status oleh ppat", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}
