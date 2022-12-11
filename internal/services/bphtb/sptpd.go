package sptpd

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	// "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	area "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	mppat "github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	mwp "github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	sppat "github.com/bapenda-kota-malang/apin-backend/internal/services/ppat"
	ssspt "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt"
	swp "github.com/bapenda-kota-malang/apin-backend/internal/services/wajibpajak"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "bphtbsptpd"

func Create(input m.CreateDto, from string, authInfo *auth.AuthInfo) (any, error) {
	var data m.BphtbSptpd
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// bphtb section
		id, err := sh.GetUuidv4()
		if err != nil {
			return err
		}
		var errChan = make(chan error)
		if input.GambarOp != nil {
			fileName, path, extFile, _, err := sh.FilePreProcess(*input.GambarOp, "bphtbGambarOp", uint(authInfo.User_Id), id)
			if err != nil {
				return err
			}
			go sh.SaveFile(*input.GambarOp, fileName, path, extFile, errChan)
			input.GambarOp = &fileName
		}
		if err := sc.Copy(&data, &input); err != nil {
			return err
		}

		// change data based from parameter and add static data
		data.Lampiran = nil
		data.Id = id
		switch from {
		case "ppat":
			respPpat, err := sppat.GetDetail(authInfo.Ref_Id)
			if err != nil {
				return err
			}
			ppatDataId := strconv.Itoa(respPpat.(rp.OKSimple).Data.(*mppat.Ppat).Id)
			data.Ppat_Id = &ppatDataId
		case "wp":
			respWp, err := swp.GetDetail(authInfo.Ref_Id)
			if err != nil {
				return err
			}
			wpNik := respWp.(rp.OKSimple).Data.(*mwp.WajibPajak).Nik
			data.WajibPajak_NIK = &wpNik
		}
		data.Status = "00"

		// process calculation
		dataSppt, err := ssspt.GetFromNop(
			*data.PermohonanProvinsiID,
			*data.PermohonanKotaID,
			*data.PermohonanKecamatanID,
			*data.PermohonanKelurahanID,
			*data.PermohonanBlokID,
			*data.NoUrutPemohon,
			*data.PemohonJenisOPID)
		if err != nil {
			return fmt.Errorf("get data from nop: %w", err)
		}
		data.NjopLuasTanah = float64(*dataSppt.NJOPBumi_sppt)
		data.NjopLuasBangunan = float64(*dataSppt.NJOPBangunan_sppt)
		njopTanah := data.OPLuasTanah * data.NjopLuasTanah
		njopBangunan := data.OPLuasBangunan * data.NjopLuasBangunan

		// data dari anggota objek pajak
		njopTanahBersama := data.OPLuasTanahBersama * data.NjopTanahBersama
		njopBangunanBersama := data.OPLuasBangunanBersama * data.NjopBangunanBersama
		data.NjopPbbOp = njopTanah + njopBangunan + njopTanahBersama + njopBangunanBersama

		npop := data.NjopPbbOp
		if npop < data.NilaiPasar {
			npop = data.NilaiPasar
		}
		data.Npop = npop
		data.NilaiOp = npop

		// FROM NPOP IF >= NPOTKP -> CALCULATE TO GET NPOPKP THEN CALCULATE NOMINAL SPT OR JUMLAH SETOR
		// TODO: WHY JUMLAH SETOR DAN NOMINAL SPT THERE'S SAME DATA OR DIFFERENT DATA FROM DATABASE EXISTING?, HOW TO CALCULATE THAT
		if data.Npop >= data.Npoptkp {
			npopkp := data.Npop - data.Npoptkp
			total := 0.05 * npopkp
			data.JumlahSetor = total
			data.NominalSPT = total
		} else {
			data.JumlahSetor = 0
			data.NominalSPT = 0
		}

		if err := <-errChan; err != nil {
			return err
		}
		if result := tx.Create(&data); result.Error != nil {
			return result.Error
		}

		input.Lampiran.BphtbSptpd_Id = data.Id
		dataLampiran, err := CreateLampiran(input.Lampiran, uint(authInfo.User_Id), tx)
		if err != nil {
			return err
		}
		data.Lampiran = &dataLampiran

		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal membuat data: "+err.Error(), data)
	}
	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto, auth int, tp string) (any, error) {
	var data []m.BphtbSptpd
	var count int64

	queryBase := a.DB.Model(&m.BphtbSptpd{})

	if (auth == 0 || auth == 4) && tp == "ver" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"01", "03", "04", "05", "06", "07", "08", "09"})
	} else if (auth == 0 || auth == 4) && tp == "byr" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"09", "10", "11", "12", "13"})
	} else if auth == 3 && tp == "ver" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"03", "06", "07", "08", "09"})
	} else if (auth == 0 || auth == 3 || auth == 2) && tp == "val" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"11", "14", "15"})
	} else if auth == 2 && tp == "ver" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"06", "08", "09"})
	}

	var pagination gh.Pagination
	result := queryBase.
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

func GetDetail(id uuid.UUID) (any, error) {
	var (
		model *m.BphtbSptpd

		provinsi  *area.Provinsi
		kota      *area.Daerah
		kecamatan *area.Kecamatan
		kelurahan *area.Kelurahan
		// blok      *area.Blok
	)

	result := a.DB.Preload(clause.Associations).First(&model, "\"Id\" = ?", id.String())
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

func Update(id uuid.UUID, input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.BphtbSptpd
	result := tx.First(&data, "\"Id\" = ?", id.String())
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

func Approval(id uuid.UUID, kd string, auth int, input m.RequestApprovalSptpd, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.BphtbSptpd
	result := tx.First(&data, "\"Id\" = ?", id.String())
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

func Delete(id uuid.UUID, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.BphtbSptpd
	result := tx.First(&data, "\"Id\" = ?", id.String())
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
