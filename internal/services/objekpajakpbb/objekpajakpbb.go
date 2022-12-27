package objekpajakpbb

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	maop "github.com/bapenda-kota-malang/apin-backend/internal/models/anggotaobjekpajak"
	mkk "github.com/bapenda-kota-malang/apin-backend/internal/models/kunjungankembali"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	mopb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	pmh "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	msppt "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	mwp "github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajakpbb"
	saop "github.com/bapenda-kota-malang/apin-backend/internal/services/anggotaobjekpajak"
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
	var dataObjekPajakBumi mopb.CreateDto
	var dataKunjunganKembali mkk.CreateDto
	var dataOpb *mopb.ObjekPajakBumi
	// var respDataWajibPajakPbb interface{}
	var respDataAnggotaObjekPajak interface{}
	var respDataKunjunganKembali interface{}
	var respDataObjekPajakBumi interface{}
	var resp t.II

	resultNop, kode := sh.NopParser(*input.Nop)
	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// check wp pbb exist or not
	var tmpWpPbb mwp.WajibPajakPbb
	result := a.DB.Where(mwp.WajibPajakPbb{Nik: input.WajibPajakPbbs.Nik}).First(&tmpWpPbb)
	if result.RowsAffected == 0 {
		// copy data wajibpajakPbb
		if err := sc.Copy(&dataWajibPajakPbb, input.WajibPajakPbbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload wajibpajakpbb", input.WajibPajakPbbs)
		}
	}

	// copy data objek pajak bumi
	if err := sc.Copy(&dataObjekPajakBumi, input.ObjekPajakBumis); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload objek pajak bumi", input.ObjekPajakBumis)
	}

	// add static field for objek pajak bumi
	dataObjekPajakBumi.Provinsi_Kode = &resultNop[0]
	dataObjekPajakBumi.Daerah_Kode = &resultNop[1]
	dataObjekPajakBumi.Kecamatan_Kode = &resultNop[2]
	dataObjekPajakBumi.Kelurahan_Kode = &resultNop[3]
	dataObjekPajakBumi.Blok_Kode = &resultNop[4]
	dataObjekPajakBumi.NoUrut = &resultNop[5]
	dataObjekPajakBumi.JenisOp = &resultNop[6]
	dataObjekPajakBumi.Area_Kode = &kode

	if input.NopBersama != nil {
		// copy data anggota objek pajak
		if err := sc.Copy(&dataAnggotaObjekPajak, input.AnggotaObjekPajaks); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload anggota objek pajak", input.AnggotaObjekPajaks)
		}

		// data induk objek
		resultNopInduk, kode := sh.NopParser(*input.NopBersama)
		dataAnggotaObjekPajak.Induk_Provinsi_Kode = &resultNopInduk[0]
		dataAnggotaObjekPajak.Induk_Daerah_Kode = &resultNopInduk[1]
		dataAnggotaObjekPajak.Induk_Kecamatan_Kode = &resultNopInduk[2]
		dataAnggotaObjekPajak.Induk_Kelurahan_Kode = &resultNopInduk[3]
		dataAnggotaObjekPajak.Induk_Blok_Kode = &resultNopInduk[4]
		dataAnggotaObjekPajak.Induk_NoUrut = &resultNopInduk[5]
		dataAnggotaObjekPajak.Induk_JenisOp = &resultNopInduk[6]
		dataAnggotaObjekPajak.Induk_Area_Kode = &kode

		dataAnggotaObjekPajak.Provinsi_Kode = &resultNop[0]
		dataAnggotaObjekPajak.Daerah_Kode = &resultNop[1]
		dataAnggotaObjekPajak.Kecamatan_Kode = &resultNop[2]
		dataAnggotaObjekPajak.Kelurahan_Kode = &resultNop[3]
		dataAnggotaObjekPajak.Blok_Kode = &resultNop[4]
		dataAnggotaObjekPajak.NoUrut = &resultNop[5]
		dataAnggotaObjekPajak.JenisOp = &resultNop[6]
		dataAnggotaObjekPajak.Area_Kode = &kode
	}

	if input.KunjunganKembalis != nil {
		// copy data kunjungan kembali
		if err := sc.Copy(&dataKunjunganKembali, input.KunjunganKembalis); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload kunjungan kembali", input.KunjunganKembalis)
		}
		resultNopKK, kode := sh.NopParser(*input.Nop)
		dataKunjunganKembali.NopDetailCreateDto.Provinsi_Kode = &resultNopKK[0]
		dataKunjunganKembali.NopDetailCreateDto.Daerah_Kode = &resultNopKK[1]
		dataKunjunganKembali.NopDetailCreateDto.Kecamatan_Kode = &resultNopKK[2]
		dataKunjunganKembali.NopDetailCreateDto.Kelurahan_Kode = &resultNopKK[3]
		dataKunjunganKembali.NopDetailCreateDto.Blok_Kode = &resultNopKK[4]
		dataKunjunganKembali.NopDetailCreateDto.NoUrut = &resultNopKK[5]
		dataKunjunganKembali.NopDetailCreateDto.JenisOp = &resultNopKK[6]
		dataKunjunganKembali.Area_Kode = &kode
	}

	// add static field for objek pajak pbb

	data.NopDetail.Provinsi_Kode = &resultNop[0]
	data.NopDetail.Daerah_Kode = &resultNop[1]
	data.NopDetail.Kecamatan_Kode = &resultNop[2]
	data.NopDetail.Kelurahan_Kode = &resultNop[3]
	data.NopDetail.Blok_Kode = &resultNop[4]
	data.NopDetail.NoUrut = &resultNop[5]
	data.NopDetail.JenisOp = &resultNop[6]
	data.NopDetail.Area_Kode = &kode

	if input.TanggalPemeriksaan != nil {

		data.TanggalPemeriksaan = th.ParseTime(input.TanggalPemeriksaan)
	}
	if input.TanggalPendataan != nil {

		data.TanggalPendataan = th.ParseTime(input.TanggalPendataan)
	}
	if input.TanggalPerekaman != nil {

		data.TanggalPerekaman = th.ParseTime(input.TanggalPerekaman)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		if tmpWpPbb.Id == 0 {
			// create data wajibpajakpbb
			resultWajibPajakPbb, err := swp.Create(dataWajibPajakPbb, tx)
			if err != nil {
				return err
			}
			// respDataWajibPajakPbb = resultWajibPajakPbb
			resultCastWajibPajakPbb := resultWajibPajakPbb.(rp.OKSimple).Data.(mwp.WajibPajakPbb)
			// add static value
			data.WajibPajakPbb_Id = &resultCastWajibPajakPbb.Id
		} else {
			data.WajibPajakPbb_Id = &tmpWpPbb.Id
		}

		// create data objekpajakpbb
		err := tx.Create(&data).Error
		if err != nil {
			return err
		}

		// cek jika ada inputan nop asal maka cari data objek pajak bumi, jika tidak ada create data objek pajak bumi
		if input.NopAsal != nil {
			result, _ := sh.NopParser(*input.NopAsal)
			condition := mopb.ObjekPajakBumi{
				NopDetail: nop.NopDetail{
					Provinsi_Kode:  &result[0],
					Daerah_Kode:    &result[1],
					Kecamatan_Kode: &result[2],
					Kelurahan_Kode: &result[3],
					Blok_Kode:      &result[4],
					NoUrut:         &result[5],
					JenisOp:        &result[6],
				}}

			// cari data objek pajak bumi

			resultDataOpb := tx.Where(condition).First(&dataOpb)

			// jika ada kurangi luas bumi data exist dgn luas bumi data inputan, jika tidak ada buat data objek pajak bumi baru
			if resultDataOpb.RowsAffected != 0 {
				dataOpb.LuasBumi -= dataObjekPajakBumi.LuasBumi
				if result := a.DB.Save(&dataOpb); result.Error != nil {
					return result.Error
				}
			} else if resultDataOpb.RowsAffected == 0 {
				// create data objek pajak bumi
				resultObjekPajakBumi, err := sopb.Create(dataObjekPajakBumi, tx)
				if err != nil {
					return err
				}
				respDataObjekPajakBumi = resultObjekPajakBumi
			}
		} else {
			// create data objek pajak bumi
			resultObjekPajakBumi, err := sopb.Create(dataObjekPajakBumi, tx)
			if err != nil {
				return err
			}
			respDataObjekPajakBumi = resultObjekPajakBumi
		}

		if input.NopBersama != nil {
			// create data anggota objek pajak
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

	if input.NopBersama != nil && input.KunjunganKembalis != nil && input.NopAsal == nil {
		resp = t.II{
			"objekPajakPbb": data,
			// "wajibPajakPbb":     respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":    respDataObjekPajakBumi.(rp.OKSimple).Data,
			"anggotaObjekPajak": respDataAnggotaObjekPajak.(rp.OKSimple).Data,
			"kunjunganKembali":  respDataKunjunganKembali.(rp.OKSimple).Data,
		}
	} else if input.NopBersama == nil && input.KunjunganKembalis != nil && input.NopAsal == nil {
		resp = t.II{
			"objekPajakPbb": data,
			// "wajibPajakPbb":    respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":   respDataObjekPajakBumi.(rp.OKSimple).Data,
			"kunjunganKembali": respDataKunjunganKembali.(rp.OKSimple).Data,
		}
	} else if input.NopBersama == nil && input.KunjunganKembalis == nil && input.NopAsal == nil {
		resp = t.II{
			"objekPajakPbb": data,
			// "wajibPajakPbb":  respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi": respDataObjekPajakBumi.(rp.OKSimple).Data,
		}
	} else if input.NopBersama != nil && input.KunjunganKembalis != nil && input.NopAsal != nil {
		resp = t.II{
			"objekPajakPbb": data,
			// "wajibPajakPbb":     respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":    dataOpb,
			"anggotaObjekPajak": respDataAnggotaObjekPajak.(rp.OKSimple).Data,
			"kunjunganKembali":  respDataKunjunganKembali.(rp.OKSimple).Data,
		}
	} else if input.NopBersama == nil && input.KunjunganKembalis != nil && input.NopAsal != nil {
		resp = t.II{
			"objekPajakPbb": data,
			// "wajibPajakPbb":    respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":   dataOpb,
			"kunjunganKembali": respDataKunjunganKembali.(rp.OKSimple).Data,
		}
	} else if input.NopBersama == nil && input.KunjunganKembalis == nil && input.NopAsal != nil {
		resp = t.II{
			"objekPajakPbb": data,
			// "wajibPajakPbb":  respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi": dataOpb,
		}
	} else if input.NopBersama != nil && input.KunjunganKembalis == nil && input.NopAsal != nil {
		resp = t.II{
			"objekPajakPbb": data,
			// "wajibPajakPbb":     respDataWajibPajakPbb.(rp.OKSimple).Data,
			"objekPajakBumi":    dataOpb,
			"anggotaObjekPajak": respDataAnggotaObjekPajak.(rp.OKSimple).Data,
		}
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func GetList(input m.FilterDto, refId int) (any, error) {
	var data []m.ObjekPajakPbb
	var count int64

	var pagination gh.Pagination
	baseQuery := a.DB.Model(&m.ObjekPajakPbb{})

	// if refId not 0 then this function called from handler wp that need data to be filtered based from auth login
	if refId != 0 {
		baseQuery = baseQuery.
			Joins("JOIN \"WajibPajakPbb\" ON \"WajibPajakPbb\".\"Id\" = \"ObjekPajakPbb\".\"WajibPajakPbb_Id\"").
			Joins("JOIN \"WajibPajak\" ON \"WajibPajakPbb\".\"Nik\" = \"WajibPajak\".\"Nik\" AND \"WajibPajak\".\"Id\" = ?", refId)
	}

	result := baseQuery.
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

func GetDetail(id int, refId int) (any, error) {
	var data *m.ObjekPajakPbb

	baseQuery := a.DB.Model(&m.ObjekPajakPbb{})

	// if refId not 0 then this function called from handler wp that need data to be filtered based from auth login
	if refId != 0 {
		baseQuery = baseQuery.
			Joins("JOIN \"WajibPajakPbb\" ON \"WajibPajakPbb\".\"Id\" = \"ObjekPajakPbb\".\"WajibPajakPbb_Id\"").
			Joins("JOIN \"WajibPajak\" ON \"WajibPajakPbb\".\"Nik\" = \"WajibPajak\".\"Nik\" AND \"WajibPajak\".\"Id\" = ?", refId)
	}

	result := baseQuery.Preload(clause.Associations).First(&data, id)
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

func GetDetailbyNop(nop pmh.PermohonanNOP) (any, error) {
	var data *m.ObjekPajakPbb

	result := a.DB.
		Where("Provinsi_Kode", nop.PermohonanProvinsiID).
		Where("Daerah_Kode", nop.PermohonanKotaID).
		Where("Kecamatan_Kode", nop.PermohonanKecamatanID).
		Where("Kelurahan_Kode", nop.PermohonanKelurahanID).
		Where("Blok_Kode", nop.PermohonanBlokID).
		Where("NoUrut", nop.NoUrutPemohon).
		Where("JenisOp", nop.PemohonJenisOPID).
		First(&data)
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

func PenilaianSppt(input []msppt.Sppt, tx *gorm.DB) (any, error) {
	var data *m.ObjekPajakPbb
	for k, v := range input {
		condition := nopSearcher(v)
		result := tx.Where(condition).First(&data)
		// if result.Error != nil {
		// 	return sh.SetError("request", "penilaian-data", source, "failed", "gagal mengambil data", data)
		// }
		if result.RowsAffected != 0 {
			data.NjopBangunan = input[k].NJOPBangunan_sppt
			data.NjopBumi = input[k].NJOPBumi_sppt
			if resultSave := tx.Save(&data); resultSave.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data objek pajak pbb", data)
			}
		}

	}
	return rp.OKSimple{Data: data}, nil
}
