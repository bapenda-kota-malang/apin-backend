package keberatan

import (
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/keberatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	sbk "github.com/bapenda-kota-malang/apin-backend/internal/services/bidangkerja"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/excelhelper"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Spt.Npwpd.ObjekPajak").
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

	result := a.DB.Model(&m.Keberatan{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Spt.Npwpd.ObjekPajak").
		First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data keberatan", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Verify(id int, input m.VerifyDto, userId uint64, bidangKerjaKode string) (any, error) {
	//ambil data based on id
	var data *m.Keberatan

	result := a.DB.First(&data, id)
	if result.Error != nil {
		return sh.SetError("request", "verify-data", source, "failed", "gagal mengambil data keberatan", data)
	}

	bidangKerjaData, err := sbk.GetByKode(bidangKerjaKode)
	if err != nil {
		return sh.SetError("request", "verify-data", source, "failed", "gagal data bidang kerja: "+err.Error(), data)
	}

	switch bidangKerjaData.Level {
	case types.LevelJabatanStaff:
		if data.Status != 0 {
			if data.Status == 1 {
				return sh.SetError("request", "verify-data", source, "failed", "data sudah disetujui oleh petugas", data)
			} else if data.Status == 2 {
				return sh.SetError("request", "verify-data", source, "failed", "data sudah ditolak oleh petugas", data)
			}
		}
		data.VerifPetugas_User_Id = &userId
		data.Status = input.Status
		data.TanggalVerifPetugas = parseTimeNowToPointer()
	case types.LevelJabatanKasubidKasubag:
		if data.VerifKasubid_User_Id != nil {
			return sh.SetError("request", "verify-data", source, "failed", "data sudah diverifikasi oleh kasubid", data)
		}
		if data.VerifPetugas_User_Id == nil {
			return sh.SetError("request", "verify-data", source, "failed", "data belum diverifikasi oleh petugas", data)
		}
		if data.Status == 2 {
			return sh.SetError("request", "verify-data", source, "failed", "data sudah ditolak oleh petugas", data)
		}
		data.VerifKasubid_User_Id = &userId
		data.Status = input.Status
		data.TanggalVerifKasubid = parseTimeNowToPointer()
	case types.LevelJabatanKabid:
		if data.VerifKabid_User_Id != nil {
			return sh.SetError("request", "verify-data", source, "failed", "data sudah diverifikasi oleh kabid", data)
		}
		if data.VerifKasubid_User_Id == nil {
			return sh.SetError("request", "verify-data", source, "failed", "data belum diverifikasi oleh kasubid", data)
		}
		if data.Status == 2 {
			return sh.SetError("request", "verify-data", source, "failed", "data sudah ditolak oleh kasubid", data)
		}
		data.VerifKabid_User_Id = &userId
		data.Status = input.Status
		data.TanggalVerifKabid = parseTimeNowToPointer()
	case types.LevelJabatanKaban:
		if data.VerifKaban_User_Id != nil {
			return sh.SetError("request", "verify-data", source, "failed", "data sudah diverifikasi oleh kaban", data)
		}
		if data.VerifKabid_User_Id == nil {
			return sh.SetError("request", "verify-data", source, "failed", "data belum diverifikasi oleh kabid", data)
		}
		if data.Status == 2 {
			return sh.SetError("request", "verify-data", source, "failed", "data sudah ditolak oleh kabid", data)
		}
		data.VerifKaban_User_Id = &userId
		data.Status = input.Status
		data.TanggalVerifKaban = parseTimeNowToPointer()
	default:
		return sh.SetError("request", "verify-data", source, "failed", "level jabatan tidak diperbolehkan verifikasi", data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "verify-data", source, "failed", "gagal menyimpan data keberatan", data)
	}
	return rp.OKSimple{
		Data: data,
	}, nil
}

func DownloadExcelList(input m.FilterDto) (*excelize.File, error) {
	var data []m.Keberatan
	result := a.DB.
		Model(&m.Keberatan{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Spt.Npwpd.ObjekPajak").
		Scopes(gh.Filter(input)).
		Find(&data)
	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data keberatan", data)
		return nil, err
	}

	var excelData = func() []interface{} {
		var tmp []interface{}
		tmp = append(tmp, map[string]interface{}{
			"A": "No",
			"B": "Tanggal",
			"C": "Pemohon",
			"D": "NPWPD",
			"E": "Nama Usaha",
			"F": "Nominal Ketetapan",
			"G": "Persentase Keberatan",
			"H": "Kasubid",
			"I": "Verifikasi Kasubid",
			"J": "Kabid",
			"K": "Verifikasi Kabid",
			"L": "Kaban",
			"M": "Verifikasi Kaban",
			"N": "Petugas",
			"O": "Verifikasi Petugas",
			"P": "Status",
		})
		for i, v := range data {
			n := i + 1
			tmp = append(tmp, map[string]interface{}{
				"A": n,
				"B": func() string {
					if v.TanggalPengajuan != nil {
						return v.TanggalPengajuan.Format("2006-01-02")
					}
					return ""
				}(),
				"C": func() string {
					if v.Pemohon != nil {
						return v.Pemohon.Name
					}
					return ""
				}(),
				"D": func() string {
					if v.Spt != nil && v.Spt.Npwpd != nil && v.Spt.Npwpd.Npwpd != nil {
						return *v.Spt.Npwpd.Npwpd
					}
					return ""
				}(),
				"E": func() string {
					if v.Spt != nil && v.Spt.Npwpd != nil && v.Spt.Npwpd.ObjekPajak != nil {
						return *v.Spt.Npwpd.ObjekPajak.Nama
					}
					return ""
				}(),
				"F": func() float64 {
					if v.NominalKetetapan != nil {
						return *v.NominalKetetapan
					}
					return 0
				}(),
				"G": func() float64 {
					if v.PersentaseKeberatan != nil {
						return *v.PersentaseKeberatan
					}
					return 0
				}(),
				"H": func() string {
					if v.KasubidUser != nil {
						return v.KasubidUser.Name
					}
					return ""
				}(),
				"I": func() string {
					if v.TanggalVerifKasubid != nil {
						return v.TanggalVerifKasubid.Format("2006-01-02")
					}
					return ""
				}(),
				"J": func() string {
					if v.KabidUser != nil {
						return v.KabidUser.Name
					}
					return ""
				}(),
				"K": func() string {
					if v.TanggalVerifKabid != nil {
						return v.TanggalVerifKabid.Format("2006-01-02")
					}
					return ""
				}(),
				"L": func() string {
					if v.KabanUser != nil {
						return v.KabanUser.Name
					}
					return ""
				}(),
				"M": func() string {
					if v.TanggalVerifKaban != nil {
						return v.TanggalVerifKaban.Format("2006-01-02")
					}
					return ""
				}(),
				"N": func() string {
					if v.PetugasUser != nil {
						return v.PetugasUser.Name
					}
					return ""
				}(),
				"O": func() string {
					if v.TanggalVerifPetugas != nil {
						return v.TanggalVerifPetugas.Format("2006-01-02")
					}
					return ""
				}(),
				"P": func() string {
					if v.Status == 0 {
						return "Baru"
					} else if v.Status == 1 || v.Status == 2 {
						return "Disetujui"
					} else if v.Status == 3 || v.Status == 4 {
						return "Ditolak"
					}

					return "Tidak Diketahui"
				}(),
			})
		}
		return tmp
	}()
	return excelhelper.ExportList(excelData, "List")
}
