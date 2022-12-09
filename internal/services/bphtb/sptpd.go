package sptpd

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	// "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	area "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "bphtbsptpd"

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

func GetListApproval(input m.RequestSptpd, auth int, tp string) (any, error) {
	var data []m.BphtbSptpd
	var count int64
	var result *gorm.DB

	var pagination gh.Pagination
	if (auth == 0 || auth == 4) && tp == "ver" {
		result = a.DB.
			Model(&m.BphtbSptpd{}).
			Where("\"Status\" IN ?", []string{"01", "03", "04", "05", "06", "07", "08", "09"}).
			Scopes(gh.Filter(input)).
			Count(&count).
			Scopes(gh.Paginate(input, &pagination)).
			Find(&data)
	} else if (auth == 0 || auth == 4) && tp == "byr" {
		result = a.DB.
			Model(&m.BphtbSptpd{}).
			Where("\"Status\" IN ?", []string{"09", "10", "11", "12", "13"}).
			Scopes(gh.Filter(input)).
			Count(&count).
			Scopes(gh.Paginate(input, &pagination)).
			Find(&data)
	} else if auth == 3 && tp == "ver" {
		result = a.DB.
			Model(&m.BphtbSptpd{}).
			Where("\"Status\" IN ?", []string{"03", "06", "07", "08", "09"}).
			Scopes(gh.Filter(input)).
			Count(&count).
			Scopes(gh.Paginate(input, &pagination)).
			Find(&data)
	} else if (auth == 0 || auth == 3 || auth == 2) && tp == "val" {
		result = a.DB.
			Model(&m.BphtbSptpd{}).
			Where("\"Status\" IN ?", []string{"11", "14", "15"}).
			Scopes(gh.Filter(input)).
			Count(&count).
			Scopes(gh.Paginate(input, &pagination)).
			Find(&data)
	} else if auth == 2 && tp == "ver" {
		result = a.DB.
			Model(&m.BphtbSptpd{}).
			Where("\"Status\" IN ?", []string{"06", "08", "09"}).
			Scopes(gh.Filter(input)).
			Count(&count).
			Scopes(gh.Paginate(input, &pagination)).
			Find(&data)
	}

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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Approval(id int, kd string, auth int, input m.RequestApprovalSptpd, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.BphtbSptpd
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	data = input.SetDataApproval(data)

	tempVal := []string{"10", "11", "12", "13"}
	if kd == "03" && auth == 4 {
		//verifikasi staff

		tempVal := "0"
		data.Proses = &tempVal
		// data.Dispenda_User_id = auth.user_id
		// data.NamaStaff = ""
	} else if kd == "04" && auth == 4 {
		//penolakan staff

		tempVal := "-1"
		data.Proses = &tempVal
		// data.Dispenda_User_id = auth.user_id
		// data.NamaStaff = ""
	} else if kd == "05" && auth == 3 {
		//verifikasi kabid

		tempVal := "2"
		data.Proses = &tempVal
	} else if kd == "06" && auth == 3 {
		//tolak kabid
		// tempVal := "2"
		// data.Proses = &tempVal
	} else if kd == "07" && auth == 3 {
		//tolak kabid
		// tempVal := "2"
		// data.Proses = &tempVal
	} else if kd == "08" && auth == 2 {
		//tolak kabid
		// tempVal := "2"
		// data.Proses = &tempVal
	} else if kd == "09" && auth == 2 {
		//tolak kabid
		// tempVal := "2"
		// data.Proses = &tempVal
	} else if stringInSlice(kd, tempVal) && auth == 4 {
		//tolak kabid
		// tempVal := "2"
		// data.Proses = &tempVal
	} else if kd == "14" && auth == 3 {
		//tolak kabid
		// tempVal := "2"
		// data.Proses = &tempVal
	} else if kd == "15" && auth == 2 {
		//tolak kabid
		// tempVal := "2"
		// data.Proses = &tempVal
	} else if auth == 0 {
		//tolak kabid
		// tempVal := "2"
		// data.Proses = &tempVal
	} else {
		return sh.SetError("request", "approval-data", source, "failed", "gagal melakukan approval data, user status tidak valid", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "approval-data", source, "failed", "gagal melakukan approval data", data)
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
