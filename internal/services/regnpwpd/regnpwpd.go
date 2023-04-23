package regnpwpd

import (
	"errors"
	"net/http"
	"strconv"

	nm "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	op "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajak"
	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/regnpwpd"
	rop "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajak"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	sn "github.com/bapenda-kota-malang/apin-backend/internal/services/npwpd"
	rsn "github.com/bapenda-kota-malang/apin-backend/internal/services/regnpwpd/regnarahubung"
	rsop "github.com/bapenda-kota-malang/apin-backend/internal/services/regnpwpd/regobjekpajak"
	rsp "github.com/bapenda-kota-malang/apin-backend/internal/services/regnpwpd/regpemilik"
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

const source = "registrasi npwpd"

func Create(input rn.CreateDto, user_Id uint) (interface{}, error) {
	user_IdConv := uint64(user_Id)
	var rekening *rm.Rekening
	var dataOp rop.RegObjekPajakCreateDto
	var register rn.RegNpwpd
	var dataPemilik []rn.RegPemilikWpCreateDto
	var dataNarahubung []rn.RegNarahubungCreateDto
	var errChan = make(chan error)
	var respDataOp interface{}
	var respDataPemilik interface{}
	var respDataNarahubung interface{}
	var resp t.II
	baseDocsName := "regNpwpd"

	fileName, path, extFile, err := filePreProcess(input.FotoKtp, user_Id, baseDocsName+"FotoKtp")
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}

	go sh.SaveFile(input.FotoKtp, fileName, path, extFile, errChan)
	if err := <-errChan; err != nil {
		return sh.SetError("request", "create-data", source, "failed", "image unsupported", input)
	}

	// get data rekening
	err = a.DB.Model(&rm.Rekening{}).First(&rekening, input.Rekening_Id).Error
	if err != nil {
		return nil, err
	}

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error

	// data registrasi objek pajak
	if err := sc.Copy(&dataOp, &input.RegObjekPajak); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload regObjekPajak", dataOp)
	}

	// data pemilik wajib pajak
	if err := sc.Copy(&dataPemilik, &input.RegPemilik); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload regPemilik", dataPemilik)
	}

	// data narahubung
	if err := sc.Copy(&dataNarahubung, &input.RegNarahubung); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload regNarahubung", dataNarahubung)
	}

	// data registrasi npwpd
	if err := sc.Copy(&register, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload registrasi npwpd", register)
	}
	register.FotoKtp = fileName
	err = a.DB.Transaction(func(tx *gorm.DB) error {

		// objekpajak
		resultRegObjekPajak, err := rsop.Create(dataOp, tx)
		if err != nil {
			return err
		}
		resultCastRegObjekPajak := resultRegObjekPajak.(rp.OKSimple).Data.(rop.RegObjekPajak)
		respDataOp = resultRegObjekPajak

		// add static field
		register.JenisPajak = mt.JenisPajakSA
		register.Status = mt.StatusAktif
		register.User_Id = user_IdConv
		register.RegObjekPajak_Id = resultCastRegObjekPajak.Id
		register.VerifyStatus = rn.VerifyStatusBaru
		register.Rekening = rekening
		register.TanggalMulaiUsaha = th.ParseTime(input.TanggalMulaiUsaha)
		slcLainLain, err := sh.GetArrayFile(input.LainLain, baseDocsName+"LainLain", user_Id)
		if err != nil {
			return err
		}
		register.LainLain = slcLainLain
		slcIzinUsaha, err := sh.GetArrayPdfAndImage(input.SuratIzinUsaha, baseDocsName+"SuratIzinUsaha", user_Id)
		if err != nil {
			return err
		}
		register.SuratIzinUsaha = slcIzinUsaha
		slcFotoObjek, err := sh.GetArrayPhoto(input.FotoObjek, baseDocsName+"FotoObjek", user_Id)
		if err != nil {
			return err
		}
		register.FotoObjek = slcFotoObjek

		err = tx.Create(&register).Error
		if err != nil {
			return err
		}

		if input.DetailRegOp != nil {
			err = insertDetailOp(*rekening.Objek, input.DetailRegOp, &register, tx)
			if err != nil {
				return err
			}
		}

		// set directur value to null if golongan orang pribadi
		if input.Golongan == 1 {
			for k := range dataPemilik {
				dataPemilik[k].Direktur_Nama = nil
				dataPemilik[k].Direktur_Alamat = nil
				dataPemilik[k].Direktur_Daerah_Id = nil
				dataPemilik[k].Direktur_Kelurahan_Id = nil
				dataPemilik[k].Direktur_Nik = nil
				dataPemilik[k].Direktur_Telp = nil
			}
		}

		resultRegPemilik, err := rsp.Create(dataPemilik, register.Id, tx)
		if err != nil {
			return err
		}
		respDataPemilik = resultRegPemilik

		resultRegNarahubung, err := rsn.Create(dataNarahubung, register.Id, tx)
		if err != nil {
			return err
		}
		respDataNarahubung = resultRegNarahubung

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), register)
	}

	resp = t.II{
		"objekPajak": respDataOp.(rp.OKSimple).Data,
		"regNpwpd":   register,
		"pemilikWp":  respDataPemilik.(rp.OKSimple).Data,
		"narahubung": respDataNarahubung.(rp.OKSimple).Data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func GetList(input rn.FilterDto) (interface{}, error) {
	var data []rn.RegNpwpd
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&rn.RegNpwpd{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Joins("RegObjekPajak").
		Preload("RegObjekPajak.Kecamatan").
		Preload("RegObjekPajak.Kelurahan").
		Preload("RegPemilikWps.Daerah").
		Preload("RegPemilikWps.Kelurahan").
		Preload("RegPemilikWps.Direktur_Daerah").
		Preload("RegPemilikWps.Direktur_Kelurahan").
		Preload("RegNarahubungs.Daerah").
		Preload("RegNarahubungs.Kelurahan").
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

func DownloadExcelList(input rn.FilterDto) (*excelize.File, error) {
	var data []rn.RegNpwpd

	result := a.DB.
		Model(&rn.RegNpwpd{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("RegObjekPajak.Kecamatan").
		Preload("RegObjekPajak.Kelurahan").
		Preload("RegPemilikWps.Daerah").
		Preload("RegPemilikWps.Kelurahan").
		Preload("RegPemilikWps.Direktur_Daerah").
		Preload("RegPemilikWps.Direktur_Kelurahan").
		Preload("RegNarahubungs.Daerah").
		Preload("RegNarahubungs.Kelurahan").
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
			"B": "ID",
			"C": "Jenis",
			"D": "Golongan",
			"E": "Jenis Usaha",
			"F": "Nama WP",
			"G": "Nama Pemilik",
			"H": "Kecamatan",
			"I": "Kelurahan",
			"J": "Tgl Daftar",
			"K": "Status",
		})
		for i, v := range data {
			n := i + 1
			tmp = append(tmp, map[string]interface{}{
				"A": n,
				"B": v.Id,
				"C": func() string {
					if v.JenisPajak == mt.JenisPajakOA {
						return "OA"
					} else if v.JenisPajak == mt.JenisPajakSA {
						return "SA"
					} else {
						return ""
					}
				}(),
				"D": func() string {
					if v.Golongan == 1 {
						return "Orang Pribadi"
					} else if v.Golongan == 2 {
						return "Badan"
					} else {
						return ""
					}
				}(),
				"E": *v.Rekening.JenisUsaha,
				"F": *v.RegObjekPajak.Nama,
				"G": *v.RegPemilikWps[0].Nama,
				"H": v.RegObjekPajak.Kecamatan.Nama,
				"I": v.RegObjekPajak.Kelurahan.Nama,
				"J": func() string {
					return v.CreatedAt.Format("2006-01-02")
				}(),
				"K": func() string {
					if v.Status == 0 {
						return "Baru"
					} else if v.Status == 1 || v.Status == 2 {
						return "Disetujui"
					} else if v.Status == 3 || v.Status == 4 {
						return "Ditolak"
					} else {
						return ""
					}
				}(),
			})
		}
		return tmp
	}()
	return excelhelper.ExportList(excelData, "List")
}

func GetDetail(r *http.Request, regID int) (interface{}, error) {
	var register *rn.RegNpwpd
	err := a.DB.Model(&rn.RegNpwpd{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("RegObjekPajak.Kecamatan").
		Preload("RegObjekPajak.Kelurahan").
		Preload("RegPemilikWps.Daerah").
		Preload("RegPemilikWps.Kelurahan").
		Preload("RegPemilikWps.Direktur_Daerah").
		Preload("RegPemilikWps.Direktur_Kelurahan").
		Preload("RegNarahubungs.Daerah").
		Preload("RegNarahubungs.Kelurahan").
		First(&register, regID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return rp.OKSimple{
		Data: register,
	}, err
}

func VerifyNpwpd(id int, input rn.VerifikasiDto) (any, error) {
	if input.VerifyStatus > 2 {
		return nil, errors.New("status tidak diketahui")
	}

	var data *rn.RegNpwpd
	var dataNpwpd nm.Npwpd
	var rekening *rm.Rekening

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data registrasi npwpd tidak dapat ditemukan")
	}
	if data.VerifyStatus != rn.VerifyStatusBaru {
		if data.VerifyStatus == rn.VerifyStatusDisetujui {
			return nil, errors.New("data sudah disetujui")
		} else if data.VerifyStatus == rn.VerifyStatusDitolak {
			return nil, errors.New("data sudah ditolak")
		}
	}

	if input.VerifyStatus == 2 {
		data.VerifyStatus = rn.VerifyStatusDitolak
		if result := a.DB.Save(&data); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data registrasi npwpd", data)
		}
		return rp.OK{
			Meta: t.IS{
				"affected": strconv.Itoa(int(result.RowsAffected)),
			},
			Data: data,
		}, nil
	}

	// copy verifikasi status
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload registrasi npwpd", data)
	}

	data.VerifiedAt = th.TimeNow()

	// copy data reg npwpd ke npwpd
	if err := sc.Copy(&dataNpwpd, &data); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload npwpd", data)
	}

	// rekening
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}

	err = a.DB.Transaction(func(tx *gorm.DB) error {
		resultObjekPajak, err := rsop.Verify(data.RegObjekPajak_Id, tx)
		if err != nil {
			return err
		}

		resultCastObjekPajak := resultObjekPajak.(rp.OKSimple).Data.(op.ObjekPajak)

		if result := tx.Save(&data); result.Error != nil {
			return errors.New("gagal menyimpan data reg npwpd")
		}

		// kecamatan_id from regobjekpajak
		var dataRegOp *rop.RegObjekPajak
		err = tx.Where(rop.RegObjekPajak{Id: data.RegObjekPajak_Id}).First(&dataRegOp).Error
		if err != nil {
			return errors.New("tidak dapat menemukan data reg objek pajak")
		}

		// creating npwpd
		var tmpNomor = generateNomor()
		tmpNpwpd := sn.GenerateNpwpd(tmpNomor, *dataRegOp.Kecamatan_Id, *rekening.KodeJenisUsaha)

		dataNpwpd.Nomor = tmpNomor
		dataNpwpd.Npwpd = &tmpNpwpd

		// tanggal
		dataNpwpd.TanggalPengukuhan = th.TimeNow()
		dataNpwpd.TanggalNpwpd = th.TimeNow()
		dataNpwpd.Id = 0
		dataNpwpd.ObjekPajak_Id = resultCastObjekPajak.Id
		err = tx.Create(&dataNpwpd).Error
		if err != nil {
			return err
		}

		// cek data detail objek pajak ada/tidak
		err = checkDataDetailObjekPajak(uint64(id), *rekening.Objek, tx)
		if err == nil {
			// transfer detail from reg
			err = verifyDetailRegObjekPajak(uint64(id), dataNpwpd.Id, *rekening.Objek, tx)
			if err != nil {
				return err
			}
		}

		// verifikasi data pemilik
		err = rsp.Verify(uint64(id), dataNpwpd.Id, tx)
		if err != nil {
			return err
		}

		// verifikasi data narahubung
		err = rsn.Verify(uint64(id), dataNpwpd.Id, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data verifikasi transaction: "+err.Error(), data)
	}
	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: dataNpwpd,
	}, nil
}

func Delete(id int) (any, error) {
	var status string = "deleted"
	var data *rn.RegNpwpd
	var rekening *rm.Rekening

	// cek data regNpwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data registrasi npwpd tidak dapat ditemukan")
	}

	// untuk mengambil data regobjekpajak
	tmpObjekPajakId := data.RegObjekPajak_Id

	// cek data rekening untuk mengambil data objek untuk menghapus detail objek pajak
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	err = a.DB.Transaction(func(tx *gorm.DB) error {

		// cek dan hapus data pemilik
		err := rsp.Delete(data.Id, tx)
		if err != nil {
			status = "no deletion"
			return err
		}

		// cek dan hapus data narahubung
		err = rsn.Delete(data.Id, tx)
		if err != nil {
			status = "no deletion"
			return err
		}

		// cek data detail objek pajak ada/tidak
		err = checkDataDetailObjekPajak(data.Id, *rekening.Objek, tx)
		if err == nil {
			// delete reg detail objek pajak
			err = deleteDetailObjekPajak(data.Id, *rekening.Objek, tx)
			if err != nil {
				status = "no deletion"
				return err
			}
		}

		// delete data regNpwpd
		result = tx.Delete(&data, id)
		if result.RowsAffected == 0 {
			return errors.New("tidak dapat menghapus data registrasi npwpd")
		}

		// delete reg objek pajak
		err = rsop.Delete(tmpObjekPajakId, tx)
		if err != nil {
			status = "no deletion"
			return err
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "delete-data", source, "failed", "gagal menghapus data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil

}

func Update(id int, input rn.UpdateDto) (any, error) {
	var data *rn.RegNpwpd
	var dataNarahubung []rn.RegNarahubungUpdateDto
	var dataPemilik []rn.RegPemilikWpUpdateDto
	var dataObjekPajak rop.RegObjekPajakUpdateDto
	var dataDetailRegObjekPajak []rn.DetailRegObjekPajak
	var respDataObjekPajak interface{}
	var respDataPemilik interface{}
	var respDataNarahubung interface{}
	var resp t.II
	var rekening *rm.Rekening

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data registrasi npwpd tidak dapat ditemukan")
	}

	//cek rekening objek, dgn mengambil data rekening menggunakan rekening id
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload registrasi npwpd", data)
	}

	if err := sc.Copy(&dataNarahubung, &input.RegNarahubung); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload reg narahubung", data)
	}

	if err := sc.Copy(&dataPemilik, &input.RegPemilik); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload reg pemilik", data)
	}

	if err := sc.Copy(&dataObjekPajak, &input.RegObjekPajak); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload reg objek pajak", data)
	}

	if input.DetailRegObjekPajak != nil {
		if err := sc.Copy(&dataDetailRegObjekPajak, &input.DetailRegObjekPajak); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload detail reg objek pajak", data)
		}
	}
	err = a.DB.Transaction(func(tx *gorm.DB) error {

		if result := tx.Save(&data); result.Error != nil {
			return errors.New("gagal menyimpan data registrasi npwpd")
		}

		resultRegPemilik, err := rsp.Update(dataPemilik, data.Id, data.Golongan, tx)
		if err != nil {
			return err
		}
		respDataPemilik = resultRegPemilik

		resultRegNarahubung, err := rsn.Update(dataNarahubung, data.Id, tx)
		if err != nil {
			return err
		}
		respDataNarahubung = resultRegNarahubung

		resultRegObjekPajak, err := rsop.Update(dataObjekPajak, data.RegObjekPajak_Id, tx)
		if err != nil {
			return err
		}
		respDataObjekPajak = resultRegObjekPajak

		if input.DetailRegObjekPajak != nil {
			err = updateDetailObjekPajak(dataDetailRegObjekPajak, data.Id, *rekening.Objek, *rekening.Nama, tx)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	resp = t.II{
		"affected":   strconv.Itoa(int(result.RowsAffected)),
		"objekPajak": respDataObjekPajak.(rp.OKSimple).Data,
		"regNpwpd":   data,
		"pemilikWp":  respDataPemilik.(rp.OKSimple).Data,
		"narahubung": respDataNarahubung.(rp.OKSimple).Data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil

}

func GetListForWp(input rn.FilterDto) (any, error) {
	var data []rn.RegNpwpd
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&rn.RegNpwpd{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("RegObjekPajak.Kecamatan").
		Preload("RegObjekPajak.Kelurahan").
		Preload("RegPemilikWps.Daerah").
		Preload("RegPemilikWps.Kelurahan").
		Preload("RegPemilikWps.Direktur_Daerah").
		Preload("RegPemilikWps.Direktur_Kelurahan").
		Preload("RegNarahubungs.Daerah").
		Preload("RegNarahubungs.Kelurahan").
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

func GetDetailForWp(id int, user_Id uint64) (interface{}, error) {
	var data *rn.RegNpwpd
	err := a.DB.Model(&rn.RegNpwpd{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("RegObjekPajak.Kecamatan").
		Preload("RegObjekPajak.Kelurahan").
		Preload("RegPemilikWps.Daerah").
		Preload("RegPemilikWps.Kelurahan").
		Preload("RegPemilikWps.Direktur_Daerah").
		Preload("RegPemilikWps.Direktur_Kelurahan").
		Preload("RegNarahubungs.Daerah").
		Preload("RegNarahubungs.Kelurahan").
		First(&data, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	if data.User_Id != user_Id {
		return nil, errors.New("tidak dapat melihat data yang bukan milik anda")
	}
	return rp.OKSimple{
		Data: data,
	}, err
}

func UpdateForWp(id int, input rn.UpdateDto, user_Id uint) (any, error) {
	var dataCheck *rn.RegNpwpd
	result := a.DB.First(&dataCheck, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data registrasi npwpd tidak dapat ditemukan")
	}

	userIdConv := uint64(user_Id)
	if dataCheck.User_Id != userIdConv {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	resultUpdate, err := Update(id, input)
	return resultUpdate, err
}

func DeleteForWp(id int, user_Id uint) (any, error) {
	var dataCheck *rn.RegNpwpd
	result := a.DB.First(&dataCheck, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data registrasi npwpd tidak dapat ditemukan")
	}

	userIdConv := uint64(user_Id)
	if dataCheck.User_Id != userIdConv {
		return nil, errors.New("tidak dapat menghapus data yang bukan milik anda")
	}

	resultDelete, err := Delete(id)
	return resultDelete, err
}
