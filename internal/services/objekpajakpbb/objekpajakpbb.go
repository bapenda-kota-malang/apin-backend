package objekpajakpbb

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	maop "github.com/bapenda-kota-malang/apin-backend/internal/models/anggotaobjekpajak"
	miop "github.com/bapenda-kota-malang/apin-backend/internal/models/indukobjekpajak"
	mkk "github.com/bapenda-kota-malang/apin-backend/internal/models/kunjungankembali"
	mopb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	mwp "github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajakpbb"
	saop "github.com/bapenda-kota-malang/apin-backend/internal/services/anggotaobjekpajak"
	siop "github.com/bapenda-kota-malang/apin-backend/internal/services/indukobjekpajak"
	skk "github.com/bapenda-kota-malang/apin-backend/internal/services/kunjungankembali"
	sopb "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakbumi"
	swp "github.com/bapenda-kota-malang/apin-backend/internal/services/wajibpajakpbb"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
)

const source = "objekpajakpbb"

func Create(input m.CreateDto) (any, error) {
	var data m.ObjekPajakPbb
	var dataWajibPajakPbb mwp.CreateDto
	var dataAnggotaObjekPajak maop.CreateDto
	var dataIndukObjekPajak miop.CreateDto
	var dataObjekPajakBumi mopb.CreateDto
	var dataKunjunganKembali mkk.CreateDto
	var respDataWajibPajakPbb interface{}
	var respDataAnggotaObjekPajak interface{}
	var respDataIndukObjekPajak interface{}
	var respDataObjekPajakBumi interface{}
	var respDataKunjunganKembali interface{}
	var resp t.II

	resultNop := nopParser(*input.Nop)
	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// copy data wajibpajakPbb
	if err := sc.Copy(&dataWajibPajakPbb, input.WajibPajakPbbs); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload wajibpajakpbb", input.WajibPajakPbbs)
	}

	// copy data objek pajak bumi
	if err := sc.Copy(&dataObjekPajakBumi, input.ObjekPajakBumis); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload objek pajak bumi", input.ObjekPajakBumis)
	}

	if input.IndukObjekPajaks != nil {
		// copy data induk objek pajak
		if err := sc.Copy(&dataIndukObjekPajak, input.IndukObjekPajaks); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload induk objek pajak", input.IndukObjekPajaks)
		}

		// copy data anggota objek pajak
		if err := sc.Copy(&dataAnggotaObjekPajak, input.AnggotaObjekPajaks); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload anggota objek pajak", input.AnggotaObjekPajaks)
		}
		resultNopInduk := nopParser(*input.NopInduk)
		dataIndukObjekPajak.NopDetailCreateDto.Provinsi_Kode = &resultNopInduk[0]
		dataIndukObjekPajak.NopDetailCreateDto.Kota_Kode = &resultNopInduk[1]
		dataIndukObjekPajak.NopDetailCreateDto.Kecamatan_Kode = &resultNopInduk[2]
		dataIndukObjekPajak.NopDetailCreateDto.Kelurahan_Kode = &resultNopInduk[3]
		dataIndukObjekPajak.NopDetailCreateDto.Blok_Kode = &resultNopInduk[4]
		dataIndukObjekPajak.NopDetailCreateDto.NoUrut = &resultNopInduk[5]
		dataIndukObjekPajak.NopDetailCreateDto.JenisOp = &resultNopInduk[6]

		resultNopAnggota := nopParser(*input.NopAnggota)
		dataAnggotaObjekPajak.Provinsi_Kode = &resultNopAnggota[0]
		dataAnggotaObjekPajak.Kota_Kode = &resultNopAnggota[1]
		dataAnggotaObjekPajak.Kecamatan_Kode = &resultNopAnggota[2]
		dataAnggotaObjekPajak.Kelurahan_Kode = &resultNopAnggota[3]
		dataAnggotaObjekPajak.Blok_Kode = &resultNopAnggota[4]
		dataAnggotaObjekPajak.NoUrut = &resultNopAnggota[5]
		dataAnggotaObjekPajak.JenisOp = &resultNopAnggota[6]
	}

	if input.KunjunganKembalis != nil {
		// copy data kunjungan kembali
		if err := sc.Copy(&dataKunjunganKembali, input.KunjunganKembalis); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload kunjungan kembali", input.KunjunganKembalis)
		}
	}

	data.NopDetail.Provinsi_Kode = &resultNop[0]
	data.NopDetail.Kota_Kode = &resultNop[1]
	data.NopDetail.Kecamatan_Kode = &resultNop[2]
	data.NopDetail.Kelurahan_Kode = &resultNop[3]
	data.NopDetail.Blok_Kode = &resultNop[4]
	data.NopDetail.NoUrut = &resultNop[5]
	data.NopDetail.JenisOp = &resultNop[6]

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// create data wajibpajakpbb
		resultWajibPajakPbb, err := swp.Create(dataWajibPajakPbb, tx)
		if err != nil {
			return err
		}
		respDataWajibPajakPbb = resultWajibPajakPbb
		resultCastWajibPajakPbb := resultWajibPajakPbb.(rp.OKSimple).Data.(mwp.WajibPajakPbb)
		// add static value
		data.WajibPajakPbb_Id = &resultCastWajibPajakPbb.Id
		if input.TanggalPemeriksaan != nil {

			data.TanggalPemeriksaan = th.ParseTime(input.TanggalPemeriksaan)
		}
		if input.TanggalPendataan != nil {

			data.TanggalPendataan = th.ParseTime(input.TanggalPendataan)
		}
		if input.TanggalPerekaman != nil {

			data.TanggalPerekaman = th.ParseTime(input.TanggalPerekaman)
		}

		// create data objekpajakpbb
		err = tx.Create(&data).Error
		if err != nil {
			return err
		}

		// create data objek pajak bumi
		resultObjekPajakBumi, err := sopb.Create(dataObjekPajakBumi, tx)
		if err != nil {
			return err
		}
		respDataObjekPajakBumi = resultObjekPajakBumi

		if input.IndukObjekPajaks != nil {
			// create data induk objek pajak
			resultIndukObjekPajak, err := siop.Create(dataIndukObjekPajak, tx)
			if err != nil {
				return err
			}
			respDataIndukObjekPajak = resultIndukObjekPajak
			resultCastIndukObjekPajak := resultIndukObjekPajak.(rp.OKSimple).Data.(miop.IndukObjekPajak)

			// create data anggota objek pajak
			dataAnggotaObjekPajak.IndukObjekPajak_Id = &resultCastIndukObjekPajak.Id
			resultAnggotaObjekPajak, err := saop.Create(dataAnggotaObjekPajak, tx)
			if err != nil {
				return err
			}
			respDataAnggotaObjekPajak = resultAnggotaObjekPajak
		}

		if input.KunjunganKembalis != nil {
			// create data kunjungan kembali
			resultKunjunganKembali, err := skk.Create(dataKunjunganKembali, tx)
			if err != nil {
				return err
			}
			respDataKunjunganKembali = resultKunjunganKembali
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	if input.IndukObjekPajaks != nil && input.KunjunganKembalis != nil {
		resp = t.II{
			"objekPajakPbb":     data,
			"wajibPajakPbb":     respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":    respDataObjekPajakBumi.(rp.OKSimple).Data,
			"indukObjekPajak":   respDataIndukObjekPajak.(rp.OKSimple).Data,
			"anggotaObjekPajak": respDataAnggotaObjekPajak.(rp.OKSimple).Data,
			"kunjunganKembali":  respDataKunjunganKembali.(rp.OKSimple).Data,
		}
	} else if input.IndukObjekPajaks != nil && input.KunjunganKembalis == nil {
		resp = t.II{
			"objekPajakPbb":     data,
			"wajibPajakPbb":     respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":    respDataObjekPajakBumi.(rp.OKSimple).Data,
			"indukObjekPajak":   respDataIndukObjekPajak.(rp.OKSimple).Data,
			"anggotaObjekPajak": respDataAnggotaObjekPajak.(rp.OKSimple).Data,
		}
	} else if input.IndukObjekPajaks == nil && input.KunjunganKembalis != nil {
		resp = t.II{
			"objekPajakPbb":    data,
			"wajibPajakPbb":    respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":   respDataObjekPajakBumi.(rp.OKSimple).Data,
			"kunjunganKembali": respDataKunjunganKembali.(rp.OKSimple).Data,
		}
	} else if input.IndukObjekPajaks == nil && input.KunjunganKembalis == nil {
		resp = t.II{
			"objekPajakPbb":  data,
			"wajibPajakPbb":  respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi": respDataObjekPajakBumi.(rp.OKSimple).Data,
		}
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.ObjekPajakPbb
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.ObjekPajakPbb{}).
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
	var data *m.ObjekPajakPbb

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

func GetDetailbyField(field string, value string) (any, error) {
	var data *m.ObjekPajakPbb

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

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.ObjekPajakPbb
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
	var data *m.ObjekPajakPbb
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
