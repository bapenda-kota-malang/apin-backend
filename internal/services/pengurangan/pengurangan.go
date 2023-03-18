package pengurangan

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	mjabatan "github.com/bapenda-kota-malang/apin-backend/internal/models/jabatan"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pengurangan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajak"
	sjabatan "github.com/bapenda-kota-malang/apin-backend/internal/services/jabatan"
	suser "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/google/uuid"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "pengurangan"

// this function create data pengurangan by staff bapenda
func Create(input m.CreateDto, user_Id uint64) (any, error) {
	var (
		dataPengurangan m.Pengurangan
		errChan         = make(chan error, 6)
		wg              sync.WaitGroup
	)

	// TODO: FILTER USER ID TO EXECUTE THIS FUNCTION
	// resp, err := suser.GetJabatanPegawai(uint(user_Id))
	// if err != nil {
	// 	return sh.SetError("request", "create-data", source, "failed", "gagal data pegawai: "+err.Error(), nil)
	// }
	// if ok := strings.Contains(strings.ToUpper(resp.(string)), "PETUGAS"); !ok {
	// 	return sh.SetError("request", "create-data", source, "failed", "pegawai bukan petugas, sehingga tidak bisa membuat data baru", nil)
	// }

	var dataSpt spt.Spt
	resCheckSpt := a.DB.First(&dataSpt, "\"Id\" = ?", input.Spt_Id.String())
	if resCheckSpt.RowsAffected == 0 {
		return nil, nil
	}

	id, err := sh.GetUuidv4()
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal membuat uuid: "+err.Error(), nil)
	}

	// foto ktp
	fileName, path, extFile, _, err := sh.FilePreProcess(input.FotoKtp, source+"FotoKtp", uint(user_Id), id)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}
	wg.Add(1)
	go sh.BulkSaveFile(&wg, input.FotoKtp, fileName, path, extFile, errChan)
	input.FotoKtp = fileName

	// laporan keuangan
	fileName, path, extFile, _, err = sh.FilePreProcess(input.LaporanKeuangan, source+"LaporanKeuangan", uint(user_Id), id)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}
	wg.Add(1)
	go sh.BulkSaveFile(&wg, input.LaporanKeuangan, fileName, path, extFile, errChan)
	input.LaporanKeuangan = fileName

	// Surat Permohonan
	fileName, path, extFile, _, err = sh.FilePreProcess(input.SuratPermohonan, source+"SuratPermohonan", uint(user_Id), id)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}
	wg.Add(1)
	go sh.BulkSaveFile(&wg, input.SuratPermohonan, fileName, path, extFile, errChan)
	input.SuratPermohonan = fileName

	// dokumen lainnya
	if input.DokumenLainnya != nil {
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.DokumenLainnya, source+"DokumenLainnya", uint(user_Id), id)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
		}
		wg.Add(1)
		go sh.BulkSaveFile(&wg, *input.DokumenLainnya, fileName, path, extFile, errChan)
		input.DokumenLainnya = &fileName
	}

	// data pengurangan
	if err := sc.Copy(&dataPengurangan, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload pengurangan", dataPengurangan)
	}

	wg.Wait()
	close(errChan)
	for v := range errChan {
		if v != nil {
			err = v
			return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan file: "+err.Error(), dataPengurangan)
		}
	}

	// static value
	dataPengurangan.Id = id
	dataPengurangan.Posisi = m.PosisiVerifikasiStaff
	dataPengurangan.Status = m.StatusDiproses
	dataPengurangan.VerifyPetugas_User_Id = &user_Id
	now := time.Now()
	dataPengurangan.TanggalVerifPetugas = &now
	if dataPengurangan.TanggalPengajuan.IsZero() {
		dataPengurangan.TanggalPengajuan = now
	}
	// create data pengurangan
	err = a.DB.Create(&dataPengurangan).Error
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataPengurangan)
	}

	return rp.OKSimple{
		Data: dataPengurangan,
	}, nil
}

func CreateFromRefPengurangan(input m.CreateFromRefDto, userId uint64, jabatanId int) (any, error) {
	respJabatan, err := sjabatan.GetDetail(jabatanId)
	if err != nil {
		return nil, err
	}
	dataJabatan := respJabatan.(rp.OKSimple).Data.(*mjabatan.Jabatan)
	// TODO: check jabatan case
	if dataJabatan.Nama == nil {
		return nil, nil
	}

	now := time.Now()
	dataPengurangan := m.Pengurangan{
		JenisPengurangan:      *input.JenisPengurangan,
		KeteranganPetugas:     input.KeteranganPetugas,
		VerifyPetugas_User_Id: &userId,
		TanggalVerifPetugas:   &now,
	}
	var dataRefPengurangan m.RefPengurangan
	if result := a.DB.Preload("Pemohon").First(&dataRefPengurangan, "\"Id\" = ?", input.RefPengurangan_Id.String()); result.RowsAffected == 0 {
		return nil, fmt.Errorf("data RefPengurangan tidak ditemukan")
	}
	var dataWp wajibpajak.WajibPajak
	if res := a.DB.Preload(clause.Associations).First(&dataWp, dataRefPengurangan.Pemohon.Ref_Id); res.Error != nil {
		return nil, fmt.Errorf("data wajib pajak tidak ditemukan")
	}

	if err := sc.Copy(&dataPengurangan, &dataRefPengurangan); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload ref pengurangan", dataRefPengurangan)
	}
	dataPengurangan.Id = uuid.Nil
	dataPengurangan.NamaPemohon = dataWp.Nama
	dataPengurangan.AlamatPemohon = fmt.Sprintf(
		"%s, %s, %s,%s,%s,%s",
		dataWp.Alamat,
		dataWp.RtRw,
		dataWp.Kelurahan.Nama,
		dataWp.Kecamatan.Nama,
		dataWp.Kota.Nama,
		dataWp.Provinsi.Nama,
	)
	dataPengurangan.TelpPemohon = &dataWp.Telp
	err = a.DB.Transaction(func(tx *gorm.DB) error {
		dataRefPengurangan.TanggalVerifPetugas = &now
		dataRefPengurangan.VerifyPetugas_User_Id = &userId
		dataRefPengurangan.Status = *input.Status
		switch *input.Status {
		case m.StatusDiproses:
			if err := a.DB.Create(&dataPengurangan).Error; err != nil {
				return err
			}
			dataRefPengurangan.Pengurangan_Id = &dataPengurangan.Id
		case m.StatusDitolak:
			dataRefPengurangan.AlasanPenolakanStaff = input.AlasanDitolakStaff
		}
		if err := a.DB.Save(&dataRefPengurangan).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal pengurangan dari ref pengurangan: "+err.Error(), dataRefPengurangan)
	}
	return rp.OKSimple{
		Data: dataPengurangan,
	}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.Pengurangan
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Pengurangan{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Spt.Npwpd.ObjekPajak").
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data pengurangan", data)
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
	var data *m.Pengurangan

	result := a.DB.Model(&m.Pengurangan{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Spt.Npwpd.ObjekPajak").
		First(&data, "\"Id\" = ?", id.String())
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data pengurangan", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id uuid.UUID, input m.UpdateDto, user_Id uint64) (any, error) {
	var (
		data    m.Pengurangan
		errChan = make(chan error, 6)
	)

	result := a.DB.First(&data, "\"Id\" = ?", id.String())
	if result.RowsAffected == 0 {
		return nil, nil
	}

	// TODO: update file
	if input.FotoKtp != nil {
		fileName, path, extFile, _, err := sh.FilePreProcess(*input.FotoKtp, source+"FotoKtp", uint(user_Id), data.Id)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		go sh.ReplaceFile(data.FotoKtp, *input.FotoKtp, fileName, path, extFile, errChan)
		if err := <-errChan; err != nil {
			return sh.SetError("request", "create-data", source, "failed", "replace image: "+err.Error(), data)
		}
		input.FotoKtp = &fileName
	}

	if input.LaporanKeuangan != nil {
		fileName, path, extFile, _, err := sh.FilePreProcess(*input.LaporanKeuangan, source+"LaporanKeuangan", uint(user_Id), data.Id)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		go sh.ReplaceFile(data.LaporanKeuangan, *input.LaporanKeuangan, fileName, path, extFile, errChan)
		if err := <-errChan; err != nil {
			return sh.SetError("request", "create-data", source, "failed", "replace image: "+err.Error(), data)
		}
		input.LaporanKeuangan = &fileName
	}

	if input.SuratPermohonan != nil {
		fileName, path, extFile, _, err := sh.FilePreProcess(*input.SuratPermohonan, source+"SuratPermohonan", uint(user_Id), data.Id)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		go sh.ReplaceFile(data.SuratPermohonan, *input.SuratPermohonan, fileName, path, extFile, errChan)
		if err := <-errChan; err != nil {
			return sh.SetError("request", "create-data", source, "failed", "replace image: "+err.Error(), data)
		}
		input.SuratPermohonan = &fileName
	}

	if input.DokumenLainnya != nil {
		fileName, path, extFile, _, err := sh.FilePreProcess(*input.DokumenLainnya, source+"DokumenLainnya", uint(user_Id), data.Id)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		if data.DokumenLainnya == nil {
			go sh.SaveFile(*input.DokumenLainnya, fileName, path, extFile, errChan)
		} else {
			go sh.ReplaceFile(*data.DokumenLainnya, *input.DokumenLainnya, fileName, path, extFile, errChan)
		}
		if err := <-errChan; err != nil {
			return sh.SetError("request", "create-data", source, "failed", "replace image: "+err.Error(), data)
		}
		input.DokumenLainnya = &fileName
	}

	if err := sc.CopyWithOption(&data, input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "update-data", sourceJPB, "failed", "gagal mengambil data payload", data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", sourceDendaADM, "failed", fmt.Sprintf("gagal mengambil menyimpan data: %s", result.Error), data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Verify(id uuid.UUID, input m.VerifyDto, userId uint64) (any, error) {
	//ambil data based on id
	var data *m.Pengurangan

	result := a.DB.First(&data, id)
	if result.Error != nil {
		return sh.SetError("request", "verify-data", source, "failed", "gagal mengambil data pengurangan", data)
	}
	//ambil nama jabatan based on userId
	resp, err := suser.GetJabatanPegawai(uint(userId))
	if err != nil {
		return sh.SetError("request", "verify-data", source, "failed", "gagal data pegawai: "+err.Error(), data)
	}
	jabatan := strings.ToUpper(resp.(string))
	userRole := ""
	fmt.Println(jabatan)

	//validasi status dan jabatan yg mengolah data
	// if petugas := strings.Contains(jabatan, "PETUGAS"); petugas {
	// 	if data.Status != 0 {
	// 		if data.Status == 1 {
	// 			return sh.SetError("request", "verify-data", source, "failed", "data sudah disetujui oleh petugas", data)
	// 		} else if data.Status == 2 {
	// 			return sh.SetError("request", "verify-data", source, "failed", "data sudah ditolak oleh petugas", data)
	// 		}
	// 	}
	// 	data.VerifPetugas_User_Id = &userId
	// 	data.Status = input.Status
	// 	data.TanggalVerifPetugas = parseTimeNowToPointer()
	// 	userRole = "petugas"
	// } else if kasubid := strings.Contains(jabatan, "KEPALA SUB BIDANG"); kasubid {
	// 	if data.VerifKasubid_User_Id != nil {
	// 		return sh.SetError("request", "verify-data", source, "failed", "data sudah diverifikasi oleh kasubid", data)
	// 	}
	// 	if data.VerifPetugas_User_Id == nil {
	// 		return sh.SetError("request", "verify-data", source, "failed", "data belum diverifikasi oleh petugas", data)
	// 	}
	// 	if data.Status == 2 {
	// 		return sh.SetError("request", "verify-data", source, "failed", "data sudah ditolak oleh petugas", data)
	// 	}
	// 	data.VerifKasubid_User_Id = &userId
	// 	data.Status = input.Status
	// 	data.TanggalVerifKasubid = parseTimeNowToPointer()
	// 	userRole = "kasubid"
	// } else if kabid := strings.Contains(jabatan, "KEPALA BIDANG"); kabid {
	// 	if data.VerifKabid_User_Id != nil {
	// 		return sh.SetError("request", "verify-data", source, "failed", "data sudah diverifikasi oleh kabid", data)
	// 	}
	// 	if data.VerifKasubid_User_Id == nil {
	// 		return sh.SetError("request", "verify-data", source, "failed", "data belum diverifikasi oleh kasubid", data)
	// 	}
	// 	if data.Status == 2 {
	// 		return sh.SetError("request", "verify-data", source, "failed", "data sudah ditolak oleh kasubid", data)
	// 	}
	// 	data.VerifKabid_User_Id = &userId
	// 	data.Status = input.Status
	// 	data.TanggalVerifKabid = parseTimeNowToPointer()
	// 	userRole = "kabid"
	// } else if kaban := strings.Contains(jabatan, "KEPALA BADAN"); kaban {
	// 	if data.VerifKaban_User_Id != nil {
	// 		return sh.SetError("request", "verify-data", source, "failed", "data sudah diverifikasi oleh kaban", data)
	// 	}
	// 	if data.VerifKabid_User_Id == nil {
	// 		return sh.SetError("request", "verify-data", source, "failed", "data belum diverifikasi oleh kabid", data)
	// 	}
	// 	if data.Status == 2 {
	// 		return sh.SetError("request", "verify-data", source, "failed", "data sudah ditolak oleh kabid", data)
	// 	}
	// 	data.VerifKaban_User_Id = &userId
	// 	data.Status = input.Status
	// 	data.TanggalVerifKaban = parseTimeNowToPointer()
	// 	userRole = "kaban"
	// }

	if userRole == "" {
		return sh.SetError("request", "verify-data", source, "failed", "pegawai bukan kabid, kasubid, kaban maupun petugas", data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "verify-data", source, "failed", "gagal menyimpan data pengurangan", data)
	}
	return rp.OKSimple{
		Data: data,
	}, nil
}

func Delete(id uuid.UUID) (interface{}, error) {
	var data *m.Pengurangan
	result := a.DB.First(&data, "\"Id\" = ?", id.String())
	if result.RowsAffected == 0 {
		return nil, nil
	}

	result = a.DB.Delete(&data, "\"Id\" = ?", id.String())
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
