package baplpengajuan

import (
	"fmt"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/baplpengajuan"
	mp "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "bapl pengajuan"

func Create(input m.CreateDto, user_Id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	if input.PetugasLapangan_Pegawai_Id != nil {
		for _, v := range input.PetugasLapangan_Pegawai_Id {
			var data *mp.Pegawai

			result := a.DB.First(&data, &v)
			if result.RowsAffected == 0 {
				return sh.SetError("request", "get-data-detail", source, "failed", fmt.Sprintf("pegawai dengan id %d tidak terdaftar", *v), data)
			} else if result.Error != nil {
				return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data pegawa: "+result.Error.Error(), data)
			}
		}
	}

	var data m.PengajuanBapl
	var baseDocsName = "pengurangan"
	var errChan = make(chan error)

	fileName, path, extFile, err := filePreProcess(input.Dokumentasi, uint(user_Id), baseDocsName+"dokumentasi")
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}

	go sh.SaveFile(input.Dokumentasi, fileName, path, extFile, errChan)
	if err := <-errChan; err != nil {
		return sh.SetError("request", "create-data", source, "failed", "image unsupported", input)
	}

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	if input.Keberatan_Id == nil && input.Pengurangan_Id == nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload, keberatan_id atau pengurangan_id harus diisi salah satu", data)
	}

	if input.Keberatan_Id != nil {
		data.JenisTransaksi = 1
	} else {
		data.JenisTransaksi = 2
	}

	data.Dokumentasi = fileName
	data.EntryBy_User_Id = &user_Id
	data.TanggalKunjungan = th.ParseTime(input.TanggalKunjungan)
	// add data file
	if input.DokumenLainnya != nil {
		slcDokumenLainnya, err := sh.GetAllTypeFile(*input.DokumenLainnya, baseDocsName+"DokumenLainnya", uint(user_Id))
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data dokumen lainnya: "+err.Error(), data)
		}
		data.DokumenLainnya = &slcDokumenLainnya
	}
	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data bapl pengajuan", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.PengajuanBapl
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.PengajuanBapl{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Npwpd.ObjekPajak").
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data bapl pengajuan", data)
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
	var data *m.PengajuanBapl

	result := a.DB.Model(&m.PengajuanBapl{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Npwpd.ObjekPajak").
		First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data bapl pengajuan", data)
	}

	if data.PetugasLapangan_Pegawai_Id != nil {
		for _, v := range *data.PetugasLapangan_Pegawai_Id {
			var dataPegawai mp.Pegawai

			result := a.DB.First(&dataPegawai, &v)
			if result.RowsAffected == 0 {
				return nil, nil
			} else if result.Error != nil {
				return sh.SetError("request", "get-data-detail", source, "failed", fmt.Sprintf("gagal mengambil data pegawai dengan id %d", &v), dataPegawai)
			}

			data.Pegawai = append(data.Pegawai, dataPegawai)
		}

	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.PengajuanBapl
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data bapl pengajuan", data)
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
	var data *m.PengajuanBapl
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
