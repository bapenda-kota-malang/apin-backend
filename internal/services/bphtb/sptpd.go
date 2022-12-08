package sptpd

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	area "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "objekpajak"

func Create(input m.RequestSptpd, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.BphtbSptpd

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.RequestSptpd) (any, error) {
	var data []m.BphtbSptpd
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.BphtbSptpd{}).
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
	var (
		model    *m.BphtbSptpd
		lampiran *m.Lampiran

		provinsi  *area.Provinsi
		kota      *area.Daerah
		kecamatan *area.Kecamatan
		kelurahan *area.Kelurahan
		// blok      *area.Blok
	)

	result := a.DB.First(&model, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", model)
	}

	data := new(m.ResponseSptpd)
	if err := sc.Copy(&data, &model); err != nil {
		return sh.SetError("request", "copy-data-detail", source, "failed", "gagal menyalin data", model)
	}

	_ = a.DB.Where("Kode", data.PermohonanProvinsiID).First(&provinsi)
	_ = a.DB.Where("Kode", *data.PermohonanProvinsiID+*data.PermohonanKotaID).First(&kota)
	_ = a.DB.Where("Kode", *data.PermohonanProvinsiID+*data.PermohonanKotaID+*data.PermohonanKecamatanID).First(&kecamatan)
	_ = a.DB.Where("Kode", *data.PermohonanProvinsiID+*data.PermohonanKotaID+*data.PermohonanKecamatanID+*data.PermohonanKelurahanID).First(&kelurahan)

	data.OPProvinsi = &provinsi.Nama
	data.OPKota = &kota.Nama
	data.OPKecamatan = &kecamatan.Nama
	data.OPKelurahan = &kelurahan.Nama

	_ = a.DB.Where("Kode", data.Provinsi_Id).First(&provinsi)
	_ = a.DB.Where("Kode", *data.Provinsi_Id+*data.Kabupaten_id).First(&kota)
	_ = a.DB.Where("Kode", *data.Provinsi_Id+*data.Kabupaten_id+*data.Kecamatan_Id).First(&kecamatan)
	_ = a.DB.Where("Kode", *data.Provinsi_Id+*data.Kabupaten_id+*data.Kecamatan_Id+*data.Kelurahan_Id).First(&kelurahan)

	data.Provinsi_wp = &provinsi.Nama
	data.Kabupaten_wp = &kota.Nama
	data.Kecamatan_wp = &kecamatan.Nama
	data.Kelurahan_wp = &kelurahan.Nama

	_ = a.DB.Where("NoSspd", model.NoDokumen).First(&lampiran)

	data.DataLampiran = lampiran

	return rp.OKSimple{
		Data: data,
	}, nil
}

func GetDetailbyField(field string, value string) (any, error) {
	var data *m.BphtbSptpd

	result := a.DB.Where(field, value).First(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.RequestSptpd, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.BphtbSptpd
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.BphtbSptpd
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	result = tx.Delete(&data, id)
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
