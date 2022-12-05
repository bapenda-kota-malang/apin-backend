package keberatan

import (
	"strconv"
	"strings"
	"time"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/keberatan"
	suser "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
)

const source = "keberatan"

func Create(input m.CreateDto, user_Id uint64) (any, error) {
	var dataKeberatan m.Keberatan
	var resp t.II
	var errChan = make(chan error)
	var baseDocsName = "keberatan"

	fileName, path, extFile, err := FilePreProcess(*input.FotoKtp, uint(user_Id), baseDocsName+"FotoKtp")
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}

	go sh.SaveFile(*input.FotoKtp, fileName, path, extFile, errChan)
	if err := <-errChan; err != nil {
		return sh.SetError("request", "create-data", source, "failed", "image unsupported", input)
	}

	// data keberatan
	if err := sc.Copy(&dataKeberatan, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload keberatan", dataKeberatan)
	}

	dataKeberatan.FotoKtp = &fileName
	err = a.DB.Transaction(func(tx *gorm.DB) error {
		// static value
		dataKeberatan.Pemohon_User_Id = &user_Id
		dataKeberatan.TanggalPengajuan = th.TimeNow()

		// add data file
		// slcLaporanKeuangan, err := sh.GetAllTypeFile(*input.LaporanKeuangan, baseDocsName+"LaporanKeuangan", uint(user_Id))
		// if err != nil {
		// 	return err
		// }
		// dataKeberatan.LaporanKeuangan = &slcLaporanKeuangan
		// slcLaporanPengeluaran, err := sh.GetAllTypeFile(*input.LaporanPengeluaran, baseDocsName+"LaporanPengeluaran", uint(user_Id))
		// if err != nil {
		// 	return err
		// }
		// dataKeberatan.LaporanPengeluaran = &slcLaporanPengeluaran
		if input.DokumenLainnya != nil {
			slcDokumenLainnya, err := sh.GetAllTypeFile(*input.DokumenLainnya, baseDocsName+"DokumenLainnya", uint(user_Id))
			if err != nil {
				return err
			}
			dataKeberatan.DokumenLainnya = &slcDokumenLainnya
		}
		if input.SuratPermohonan != nil {
			slcSuratPermohonan, err := sh.GetPdfOrImageFile(*input.SuratPermohonan, baseDocsName+"SuratPermohonan", uint(user_Id))
			if err != nil {
				return err
			}
			dataKeberatan.SuratPermohonan = &slcSuratPermohonan
		}
		if input.SuratPernyataan != nil {
			slcSuratPernyataan, err := sh.GetPdfOrImageFile(*input.SuratPernyataan, baseDocsName+"SuratPernyataan", uint(user_Id))
			if err != nil {
				return err
			}
			dataKeberatan.SuratPernyataan = &slcSuratPernyataan
		}
		// create data keberatan
		err = tx.Create(&dataKeberatan).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataKeberatan)
	}

	resp = t.II{
		"keberatan": dataKeberatan,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.Keberatan
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Keberatan{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data keberatan", data)
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
	var data *m.Keberatan

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data keberatan", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Verify(id int, input m.VerifyDto, userId uint64) (any, error) {
	//ambil data based on id
	var data *m.Keberatan

	result := a.DB.First(&data, id)
	if result.Error != nil {
		return sh.SetError("request", "verify-data", source, "failed", "gagal mengambil data keberatan", data)
	}
	//ambil nama jabatan based on userId
	resp, err := suser.GetJabatanPegawai(uint(userId))
	if err != nil {
		return sh.SetError("request", "verify-data", source, "failed", "gagal data pegawai: "+err.Error(), data)
	}
	jabatan := strings.ToUpper(resp.(string))
	userRole := ""

	//validasi status dan jabatan yg mengolah data
	if petugas := strings.Contains(jabatan, "PETUGAS"); petugas {
		if data.Status != 0 {
			if data.Status == 1 {
				return sh.SetError("request", "verify-data", source, "failed", "data sudah disetujui oleh petugas", data)
			} else if data.Status == 2 {
				return sh.SetError("request", "verify-data", source, "failed", "data sudah ditolak oleh petugas", data)
			}
		}
		data.VerifPetugas_User_Id = &userId
		data.Status = input.Status
		data.TanggalPetugas = time.Now()
		userRole = "petugas"
	} else if kasubid := strings.Contains(jabatan, "KEPALA SUB BIDANG"); kasubid {
		if data.VerifPetugas_User_Id == nil {
			return sh.SetError("request", "verify-data", source, "failed", "data belum diverifikasi oleh petugas", data)
		}
		if data.Status == 2 {
			return sh.SetError("request", "verify-data", source, "failed", "data sudah ditolak oleh petugas", data)
		}
		data.VerifKasubid_User_Id = &userId
		data.Status = input.Status
		data.TanggalKasubid = time.Now()
		userRole = "kasubid"
	} else if kabid := strings.Contains(jabatan, "KEPALA BIDANG"); kabid {
		if data.VerifKasubid_User_Id == nil {
			return sh.SetError("request", "verify-data", source, "failed", "data belum diverifikasi oleh kasubid", data)
		}
		if data.Status == 2 {
			return sh.SetError("request", "verify-data", source, "failed", "data sudah ditolak oleh kasubid", data)
		}
		data.VerifKabid_User_Id = &userId
		data.Status = input.Status
		data.TanggalKabid = time.Now()
		userRole = "kabid"
	} else if kaban := strings.Contains(jabatan, "KEPALA BADAN"); kaban {
		if data.VerifKabid_User_Id == nil {
			return sh.SetError("request", "verify-data", source, "failed", "data belum diverifikasi oleh kabid", data)
		}
		if data.Status == 2 {
			return sh.SetError("request", "verify-data", source, "failed", "data sudah ditolak oleh kabid", data)
		}
		data.VerifKaban_User_Id = &userId
		data.Status = input.Status
		data.TanggalKaban = time.Now()
		userRole = "kaban"
	}

	if userRole == "" {
		return sh.SetError("request", "verify-data", source, "failed", "pegawai bukan kabid, kasubid, kaban maupun petugas", data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "verify-data", source, "failed", "gagal menyimpan data keberatan", data)
	}
	return rp.OKSimple{
		Data: data,
	}, nil
}
