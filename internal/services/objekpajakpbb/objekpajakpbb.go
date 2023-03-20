package objekpajakpbb

import (
	"fmt"
	"strconv"

	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/kecamatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/kelurahan"
	sc "github.com/jinzhu/copier"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	excelhelper "github.com/bapenda-kota-malang/apin-backend/pkg/excelhelper"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	maop "github.com/bapenda-kota-malang/apin-backend/internal/models/anggotaobjekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	mkk "github.com/bapenda-kota-malang/apin-backend/internal/models/kunjungankembali"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	mopb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	pmh "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	regpmh "github.com/bapenda-kota-malang/apin-backend/internal/models/regpelayanan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
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

func getAreaData(data m.ObjekPajakPbb) (*areadivision.Kecamatan, *areadivision.Kelurahan, error) {
	kecamatanKode := fmt.Sprintf("%s%s%s", *data.Provinsi_Kode, *data.Daerah_Kode, *data.Kecamatan_Kode)
	kecamatanKodeInt, _ := strconv.Atoi(kecamatanKode)
	resKec, err := kecamatan.GetDetailByCode(kecamatanKodeInt)
	if resKec == nil || err != nil {
		return nil, nil, err
	}

	// data kelurahan
	kelurahanKode := fmt.Sprintf("%s%s", kecamatanKode, *data.Kelurahan_Kode)
	kelurahanKodeInt, _ := strconv.Atoi(kelurahanKode)
	resKel, err := kelurahan.GetDetailByCode(kelurahanKodeInt)
	if resKel == nil || err != nil {
		return nil, nil, err
	}
	return resKec.(rp.OKSimple).Data.(*areadivision.Kecamatan), resKel.(rp.OKSimple).Data.(*areadivision.Kelurahan), nil
}

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
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload anggota objek pajak: "+err.Error(), input.AnggotaObjekPajaks)
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

func GetList(input m.FilterDto, refId int, user_id *int) (any, error) {
	var data []m.ObjekPajakPbb
	var dataWp []mwp.WajibPajakPbb
	var count int64

	var pagination gh.Pagination
	baseQuery := a.DB.Model(&m.ObjekPajakPbb{})

	// if refId not 0 then this function called from handler wp that need data to be filtered based from auth login

	if user_id != nil {
		baseQuery = baseQuery.
			Joins("JOIN \"WajibPajakPbb\" ON \"WajibPajakPbb\".\"Id\" = \"ObjekPajakPbb\".\"WajibPajakPbb_Id\" AND \"WajibPajakPbb\".\"User_ID\" = ?", user_id)
	} else if refId != 0 {
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

	for i := 0; i < len(data); i++ {
		kecamatan, kelurahan, err := getAreaData(data[i])
		if err != nil {
			return sh.SetError("request", "get-data-list", source, "failed", fmt.Sprintf("gagal mengambil data: %s", err), data[i])
		}
		data[i].Kecamatan = kecamatan
		data[i].Kelurahan = kelurahan

		resWpPbb, err := swp.GetDetail(int(*data[i].WajibPajakPbb_Id))
		if err != nil {
			return sh.SetError("request", "get-data-list", source, "failed", fmt.Sprintf("gagal mengambil data: %s", err), data[i])
		}
		if resWpPbb == nil {
			continue
		}
		dataWpTmp := resWpPbb.(rp.OKSimple).Data.(*mwp.WajibPajakPbb)
		dataWpTmp.ObjekPajakPbbs = &[]m.ObjekPajakPbb{data[i]}
		dataWp = append(dataWp, *dataWpTmp)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: dataWp,
	}, nil
}

func GetDetail(id int, refId int, user_id *int) (any, error) {
	var (
		data   *m.ObjekPajakPbb
		datawp *mwp.WajibPajakPbb
	)

	baseQuery := a.DB.Model(&m.ObjekPajakPbb{})

	// if refId not 0 then this function called from handler wp that need data to be filtered based from auth login
	if user_id != nil {
		baseQuery = baseQuery.
			Joins("JOIN \"WajibPajakPbb\" ON \"WajibPajakPbb\".\"Id\" = \"ObjekPajakPbb\".\"WajibPajakPbb_Id\" AND \"WajibPajakPbb\".\"User_ID\" = ?", user_id)
	} else if refId != 0 {
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

	kecamatan, kelurahan, err := getAreaData(*data)
	if err != nil {
		return sh.SetError("request", "get-data-list", source, "failed", fmt.Sprintf("gagal mengambil data: %s", err), data)
	}
	data.Kecamatan = kecamatan
	data.Kelurahan = kelurahan

	resWpPbb, err := swp.GetDetail(int(*data.WajibPajakPbb_Id))
	if resWpPbb == nil || err != nil {
		return sh.SetError("request", "get-data-list", source, "failed", fmt.Sprintf("gagal mengambil data: %s", err), data)
	}
	datawp = resWpPbb.(rp.OKSimple).Data.(*mwp.WajibPajakPbb)
	datawp.ObjekPajakPbbs = &[]m.ObjekPajakPbb{*data}

	return rp.OKSimple{
		Data: datawp,
	}, nil
	// if user_id != nil {
	// 	// var tempData []m.ObjekPajakPbb
	// 	// tempData = append(tempData, *data)
	// 	// datawp.ObjekPajakPbbs = &tempData
	// } else {
	// 	return rp.OKSimple{
	// 		Data: data,
	// 	}, nil
	// }
}

func GetDetailbyField(field string, value string) (any, error) {
	var data *m.ObjekPajakPbb

	result := a.DB.Where(field, value).First(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	kecamatan, kelurahan, err := getAreaData(*data)
	if err != nil {
		return sh.SetError("request", "get-data-list", source, "failed", fmt.Sprintf("gagal mengambil data: %s", err), data)
	}
	data.Kecamatan = kecamatan
	data.Kelurahan = kelurahan

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

	kecamatan, kelurahan, err := getAreaData(*data)
	if err != nil {
		return sh.SetError("request", "get-data-list", source, "failed", fmt.Sprintf("gagal mengambil data: %s", err), data)
	}
	data.Kecamatan = kecamatan
	data.Kelurahan = kelurahan

	return rp.OKSimple{
		Data: data,
	}, nil
}

func GetRegDetailbyNop(nop regpmh.PermohonanNOP) (any, error) {
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

	kecamatan, kelurahan, err := getAreaData(*data)
	if err != nil {
		return sh.SetError("request", "get-data-list", source, "failed", fmt.Sprintf("gagal mengambil data: %s", err), data)
	}
	data.Kecamatan = kecamatan
	data.Kelurahan = kelurahan

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

func UpdatePelayananWP(id int, input m.UpdateDto, from string, authInfo *auth.AuthInfo, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	var data *m.ObjekPajakPbb
	var datawp mwp.WajibPajakPbb
	var dataWajibPajakPbb mwp.UpdateDto
	var dataObjekPajakBumi mopb.UpdateDto
	// var dataOpb *mopb.ObjekPajakBumi
	var respDataWajibPajakPbb interface{}
	var respDataObjekPajakBumi interface{}
	var resp t.II
	var kode string

	if input.Nop != nil {
		var resultNop []string
		resultNop, kode = sh.NopParser(*input.Nop)
		if input.Provinsi_Kode == nil {
			input.Provinsi_Kode = &resultNop[0]
			input.Daerah_Kode = &resultNop[1]
			input.Kecamatan_Kode = &resultNop[2]
			input.Kelurahan_Kode = &resultNop[3]
			input.Blok_Kode = &resultNop[4]
			input.NoUrut = &resultNop[5]
			input.JenisOp = &resultNop[6]
		}
	} else if input.Provinsi_Kode != nil {
		kode = fmt.Sprintf("%s%s%s%s", *input.Provinsi_Kode, *input.Daerah_Kode, *input.Kecamatan_Kode, *input.Kelurahan_Kode)
	} else {
		return sh.SetError("request", "update-data", source, "failed", "data nop tidak ditemukan", input)
	}

	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.CopyWithOption(&data, &input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	// copy data wajibpajakPbb
	resultWp := tx.First(&datawp, data.WajibPajakPbb_Id)
	if resultWp.RowsAffected > 0 {
		if err := sc.Copy(&dataWajibPajakPbb, datawp); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg wajibpajakpbb", input.WajibPajakPbbs)
		}
	}
	if err := sc.CopyWithOption(&dataWajibPajakPbb, input.WajibPajakPbbs, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg wajibpajakpbb", input.WajibPajakPbbs)
	}
	tempUID := uint64(authInfo.User_Id)
	dataWajibPajakPbb.User_ID = &tempUID

	// copy data objek pajak bumi
	if input.ObjekPajakBumis != nil {
		if err := sc.Copy(&dataObjekPajakBumi, input.ObjekPajakBumis); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg objek pajak bumi", input.ObjekPajakBumis)
		}
		dataObjekPajakBumi.Provinsi_Kode = input.Provinsi_Kode
		dataObjekPajakBumi.Daerah_Kode = input.Daerah_Kode
		dataObjekPajakBumi.Kecamatan_Kode = input.Kecamatan_Kode
		dataObjekPajakBumi.Kelurahan_Kode = input.Kelurahan_Kode
		dataObjekPajakBumi.Blok_Kode = input.Blok_Kode
		dataObjekPajakBumi.NoUrut = input.NoUrut
		dataObjekPajakBumi.JenisOp = input.JenisOp
		dataObjekPajakBumi.Area_Kode = &kode
	}

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
		// create data wajibpajakpbb
		resultWajibPajakPbb, err := swp.Update(int(*data.WajibPajakPbb_Id), dataWajibPajakPbb, tx)
		if err != nil {
			return err
		}
		respDataWajibPajakPbb = resultWajibPajakPbb

		// create data objekpajakpbb
		if result := tx.Save(&data); result.Error != nil {
			return result.Error
		}

		if input.ObjekPajakBumis != nil {
			// create data objek pajak bumi
			resultObjekPajakBumi, err := sopb.Update(0, dataObjekPajakBumi, tx)
			if err != nil {
				return err
			}
			respDataObjekPajakBumi = resultObjekPajakBumi
		}

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	if input.ObjekPajakBumis != nil {
		resp = t.II{
			"regObjekPajakPbb":  data,
			"regWajibPajakPbb":  respDataWajibPajakPbb.(rp.OK).Data,
			"regObjekPajakBumi": respDataObjekPajakBumi.(rp.OK).Data,
		}
	} else {
		resp = t.II{
			"regObjekPajakPbb":  data,
			"regWajibPajakPbb":  respDataWajibPajakPbb.(rp.OK).Data,
			"regObjekPajakBumi": nil,
		}
	}
	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: resp,
	}, nil
}

func UpdateRtRwMassal(input m.UpdateRtRwMassalDto) (any, error) {
	var data []m.ObjekPajakPbb
	result := a.DB.
		Where("Provinsi_Kode", input.Provinsi_Kode).
		Where("Daerah_Kode", input.Daerah_Kode).
		Where("Kecamatan_Kode", input.Kecamatan_Kode).
		Where("Kelurahan_Kode", input.Kelurahan_Kode).
		Where("\"Blok_Kode\" BETWEEN ? AND ?", input.AwalBlok_Kode, input.AkhirBlok_Kode).
		Where("NOT (\"Blok_Kode\" = ? AND \"NoUrut\" < ?)", input.AwalBlok_Kode, input.AwalNoUrut).
		Where("NOT (\"Blok_Kode\" = ? AND \"NoUrut\" > ?)", input.AkhirBlok_Kode, input.AkhirNoUrut).
		Where("\"JenisOp\" BETWEEN ? AND ?", input.AwalJenisOp, input.AkhirJenisOp).
		Find(&data)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "update rt rw massal error: data tidak ditemukan", data)
	}

	for i := 0; i < len(data); i++ {
		data[i].Rt = input.Rt
		data[i].Rw = input.Rw
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

func GetByNop(provinsiKode, daerahKode, kecamatanKode, kelurahanKode, blokKode, noUrut, jenisOp *string) (m.ObjekPajakPbb, error) {
	var data m.ObjekPajakPbb
	condition := nopSearcher(sppt.Sppt{
		Propinsi_Id:   provinsiKode,
		Dati2_Id:      daerahKode,
		Kecamatan_Id:  kecamatanKode,
		Keluarahan_Id: kelurahanKode,
		Blok_Id:       blokKode,
		NoUrut:        noUrut,
		JenisOP_Id:    jenisOp,
	})
	result := a.DB.Where(condition).First(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func DownloadExcelList(input m.FilterDto, refId int) (*excelize.File, error) {
	var data []m.ObjekPajakPbb
	baseQuery := a.DB.Model(&m.ObjekPajakPbb{})

	// if refId not 0 then this function called from handler wp that need data to be filtered based from auth login
	if refId != 0 {
		baseQuery = baseQuery.
			Joins("JOIN \"WajibPajakPbb\" ON \"WajibPajakPbb\".\"Id\" = \"ObjekPajakPbb\".\"WajibPajakPbb_Id\"").
			Joins("JOIN \"WajibPajak\" ON \"WajibPajakPbb\".\"Nik\" = \"WajibPajak\".\"Nik\" AND \"WajibPajak\".\"Id\" = ?", refId)
	}

	result := baseQuery.
		Scopes(gh.Filter(input)).
		Find(&data)
	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
		return nil, err
	}

	var excelData = func() []interface{} {
		var tmp []interface{}
		tmp = append(tmp, map[string]interface{}{
			"A": "No",
			"B": "NOP",
			"C": "Nama WP",
			"D": "Lokasi",
		})
		for i, v := range data {
			n := i + 1
			tmp = append(tmp, map[string]interface{}{
				"A": n,
				"B": fmt.Sprintf("%s%s%s%s%s%s%s", *v.Provinsi_Kode, *v.Daerah_Kode, *v.Kecamatan_Kode, *v.Kelurahan_Kode, *v.Blok_Kode, *v.NoUrut, *v.JenisOp),
				"C": "",
				"D": "",
			})
		}
		return tmp
	}()
	return excelhelper.ExportList(excelData, "List")
}
