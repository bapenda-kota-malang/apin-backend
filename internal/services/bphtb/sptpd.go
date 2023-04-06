package sptpd

import (
	"fmt"
	"strconv"
	"time"

	"crypto/rand"

	"github.com/google/uuid"
	sc "github.com/jinzhu/copier"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	// "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	maop "github.com/bapenda-kota-malang/apin-backend/internal/models/anggotaobjekpajak"
	area "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	mppat "github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	msspt "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	mwp "github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajak"
	saop "github.com/bapenda-kota-malang/apin-backend/internal/services/anggotaobjekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	sppat "github.com/bapenda-kota-malang/apin-backend/internal/services/ppat"
	ssspt "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt"
	swp "github.com/bapenda-kota-malang/apin-backend/internal/services/wajibpajak"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/excelhelper"
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
			data.Status = "01"
		case "wp":
			respWp, err := swp.GetDetail(authInfo.Ref_Id)
			if err != nil {
				return err
			}
			wpNik := respWp.(rp.OKSimple).Data.(*mwp.WajibPajak).Nik
			data.WajibPajak_NIK = &wpNik
			data.Status = "00"
		default:
			data.Status = "00"
		}

		// process calculation
		respDataSppt, err := ssspt.GetByNop(
			data.PermohonanProvinsiID,
			data.PermohonanKotaID,
			data.PermohonanKecamatanID,
			data.PermohonanKelurahanID,
			data.PermohonanBlokID,
			data.NoUrutPemohon,
			data.PemohonJenisOPID,
			data.TahunPajakSppt,
		)
		if err != nil {
			return fmt.Errorf("get data sppt from nop: %w", err)
		}
		dataSppt := respDataSppt.(rp.OKSimple).Data.(msspt.Sppt)
		data.NjopLuasTanah = float64(*dataSppt.NJOPBumi_sppt)
		data.NjopLuasBangunan = float64(*dataSppt.NJOPBangunan_sppt)
		data.Npoptkp = float64(*dataSppt.NJOPTKP_sppt)
		njopTanah := data.OPLuasTanah * data.NjopLuasTanah
		njopBangunan := data.OPLuasBangunan * data.NjopLuasBangunan

		// data dari anggota objek pajak
		if data.OPLuasTanahBersama > 0 || data.OPLuasBangunanBersama > 0 {
			nopString := fmt.Sprintf("%s.%s.%s.%s.%s.%s.%s", *data.PermohonanProvinsiID,
				*data.PermohonanKotaID,
				*data.PermohonanKecamatanID,
				*data.PermohonanKelurahanID,
				*data.PermohonanBlokID,
				*data.NoUrutPemohon,
				*data.PemohonJenisOPID)
			respAop, err := saop.GetByNop(nopString, tx)
			if err != nil {
				return fmt.Errorf("get data anggota objek pajak from nop: %w", err)
			}
			aopData := respAop.(rp.OKSimple).Data.(*maop.AnggotaObjekPajak)
			data.NjopTanahBersama = float64(*aopData.LuasBumiBeban)
			data.NjopBangunanBersama = float64(*aopData.LuasBangunanBeban)
		}
		njopTanahBersama := data.OPLuasTanahBersama * data.NjopTanahBersama
		njopBangunanBersama := data.OPLuasBangunanBersama * data.NjopBangunanBersama
		data.NjopPbbOp = njopTanah + njopBangunan + njopTanahBersama + njopBangunanBersama

		// set npop value
		if data.NjopPbbOp < data.NilaiPasar {
			data.Npop = data.NilaiPasar
			data.NilaiOp = data.NilaiPasar
		} else {
			data.Npop = data.NjopPbbOp
			data.NilaiOp = data.NjopPbbOp
		}

		// FROM NPOP IF >= NPOTKP -> CALCULATE TO GET NPOPKP THEN CALCULATE NOMINAL SPT OR JUMLAH SETOR
		// TODO: WHY JUMLAH SETOR DAN NOMINAL SPT THERE'S SAME DATA OR DIFFERENT DATA FROM DATABASE EXISTING?
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
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"02", "03", "04", "05", "06", "07", "08", "21"})
	} else if auth == 3 && tp == "ver" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"04", "05", "06", "07", "08", "21"})
	} else if auth == 2 && tp == "ver" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"06", "07", "08", "21"})
	} else if (auth == 0 || auth == 4) && tp == "byr" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"08", "10", "12", "22"})
	} else if (auth == 0 || auth == 4) && tp == "val" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"10", "13", "14", "20"})
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

func DownloadExcelListVerifikasi(input m.FilterDto, auth int, tp string) (*excelize.File, error) {
	var data []m.BphtbSptpd

	queryBase := a.DB.Model(&m.BphtbSptpd{})

	if (auth == 0 || auth == 6) && tp == "ver" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"02", "03", "04", "05", "06", "07", "08", "21"})
	} else if auth == 4 && tp == "ver" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"04", "05", "06", "07", "08", "21"})
	} else if auth == 3 && tp == "ver" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"06", "07", "08", "21"})
	} else if (auth == 0 || auth == 6) && tp == "byr" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"08", "10", "12", "22"})
	} else if (auth == 0 || auth == 6) && tp == "val" {
		queryBase = queryBase.Where("\"Status\" IN ?", []string{"10", "13", "14", "20"})
	}

	result := queryBase.
		Scopes(gh.Filter(input)).
		Find(&data)
	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
		return nil, err
	}

	var excelData []interface{}

	if tp == "byr" {
		excelData = func() []interface{} {
			var tmp []interface{}
			tmp = append(tmp, map[string]interface{}{
				"A": "No",
				"B": "No Pelayanan",
				"C": "No SSPD",
				"D": "Nama WP",
				"E": "Alamat OP",
				"F": "Jumlah Setor",
				"G": "Tanggal",
			})
			for i, v := range data {
				n := i + 1
				tmp = append(tmp, map[string]interface{}{
					"A": n,
					"B": func() string {
						if v.NoPelayanan != nil {
							return *v.NoPelayanan
						}
						return ""
					}(),
					"C": func() string {
						if v.NoDokumen != nil {
							return *v.NoDokumen
						}
						return ""
					}(),
					"D": func() string {
						if v.NamaWp != nil {
							return *v.NamaWp
						}
						return ""
					}(),
					"E": func() string {
						if v.NOPAlamat != nil {
							return *v.NOPAlamat
						}
						return ""
					}(),
					"F": func() string {
						return strconv.FormatFloat(v.JumlahSetor, 'f', 0, 64)
					}(),
					"G": "-",
				})
			}
			return tmp
		}()
	} else {
		excelData = func() []interface{} {
			var tmp []interface{}
			tmp = append(tmp, map[string]interface{}{
				"A": "No",
				"B": "Tanggal Pengajuan",
				"C": "No Pelayanan",
				"D": "Nama Wajib Pajak",
				"E": "NOP Alamat OP",
				"F": "Jumlah Setor",
				"G": "Status",
			})
			for i, v := range data {
				n := i + 1
				tmp = append(tmp, map[string]interface{}{
					"A": n,
					"B": func() string {
						t, _ := v.Tanggal.Value()
						return t.(time.Time).Format("2006-01-02")
					}(),
					"C": func() string {
						if v.NoPelayanan != nil {
							return *v.NoPelayanan
						}
						return ""
					}(),
					"D": func() string {
						if v.NamaWp != nil {
							return *v.NamaWp
						}
						return ""
					}(),
					"E": func() string {
						if v.Alamat != nil {
							return *v.Alamat
						}
						return ""
					}(),
					"F": func() string {
						return strconv.FormatFloat(v.JumlahSetor, 'f', 0, 64)
					}(),
					"G": v.Status,
				})
			}
			return tmp
		}()
	}
	return excelhelper.ExportList(excelData, "List")
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

	if data.Lampiran == nil {
		tempLampiran := new(m.Lampiran)
		data.Lampiran = tempLampiran
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func GetDetailbyField(field string, value string) (any, error) {
	var data *[]m.BphtbSptpd

	result := a.DB.Where(field, value).Find(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func GetListTransaksiPPAT(input m.FilterPPATDto) (any, error) {
	var data []m.TransaksiPPAT
	var count int64

	queryBase := a.DB.Model(&m.BphtbSptpd{})
	queryBase = queryBase.
		Select("DISTINCT ON (\"Ppat_Id\") \"Ppat_Id\"", "(\"Ppat\".\"Nama\") \"Ppat_Name\"", "count(\"Sptpd_Id\") \"Sptpd_Id\"", "count(\"JumlahSetor\") \"CountJumlahSetor\"", "sum(\"NilaiOp\") \"NilaiOp\"", "sum(\"JumlahSetor\") \"JumlahSetor\"").
		Joins("LEFT JOIN \"Ppat\" ON \"Ppat\".\"Id\" = CAST(\"Ppat_Id\" AS INTEGER) ")

	if input.Ppat_Id != nil {
		queryBase = queryBase.Where("Ppat_Id", input.Ppat_Id)
	}

	if input.Bulan != nil {
		queryBase = queryBase.Where("EXTRACT('month' from \"VerifikasiPpatAt\") = ?", input.Bulan)
	}

	if input.Tahun != nil {
		queryBase = queryBase.Where("EXTRACT('year' from \"VerifikasiPpatAt\") = ?", input.Tahun)
	}

	result := queryBase.
		Order("\"Ppat_Id\"").
		Group("\"Ppat_Id\", \"Ppat\".\"Nama\"").
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func GetDetailTransaksiPPAT(input m.FilterPPATDto) (any, error) {
	var data []m.BphtbSptpd
	var dataPPAT mppat.Ppat
	var count int64
	var dateString string
	var namePPAT string

	queryBase := a.DB.Model(&m.BphtbSptpd{})
	queryBase = queryBase.
		Joins("LEFT JOIN \"Ppat\" ON \"Ppat\".\"Id\" = CAST(\"Ppat_Id\" AS INTEGER) ")

	if input.Ppat_Id != nil {
		queryBase = queryBase.Where("Ppat_Id", input.Ppat_Id)
	}

	if input.Bulan != nil {
		queryBase = queryBase.Where("EXTRACT('month' from \"VerifikasiPpatAt\") = ?", input.Bulan)
	}

	if input.Tahun != nil {
		queryBase = queryBase.Where("EXTRACT('year' from \"VerifikasiPpatAt\") = ?", input.Tahun)
	}

	queryBase = queryBase.Order("\"TahunPajakSppt\"")
	result := queryBase.
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	tempPPATID, _ := strconv.Atoi(*input.Ppat_Id)
	resPPAT := a.DB.First(dataPPAT, tempPPATID)
	if resPPAT.Error != nil {
		namePPAT = "PPAT" + *input.Ppat_Id
	} else {
		namePPAT = dataPPAT.Nama
	}

	if input.Bulan != nil && input.Tahun != nil {
		dateString = strconv.Itoa(*input.Tahun) + "-" + strconv.Itoa(*input.Bulan) + "-01"
	} else if input.Bulan != nil {
		dateString = "2022-" + strconv.Itoa(*input.Bulan) + "-01"
	} else if input.Tahun != nil {
		dateString = strconv.Itoa(*input.Tahun) + "-12-01"
	} else {
		if len(data) > 0 && data[0].TahunPajakSppt != nil {
			dateString = *data[0].TahunPajakSppt + "-12-01"
		} else {
			dateString = "2022-12-01"
		}
	}
	resT, errPer := time.Parse("2006-01-02", dateString)
	if errPer != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "periode tidak ditemukan", data)
	}

	startPeriode := time.Date(resT.Year(), resT.Month(), 1, 0, 0, 0, 0, time.UTC)
	endPeriode := startPeriode.AddDate(0, 1, 0).Add(-time.Nanosecond)
	if input.Bulan == nil {
		startPeriode = time.Date(resT.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"start":        startPeriode.String(),
			"end":          endPeriode.String(),
			"name":         namePPAT,
		},
		Data: data,
	}, nil
}

func LastDayOfMonth(t time.Time) int {
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, 0).Add(-time.Nanosecond)
	return lastDay.Day()
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

func LogApproval(id uuid.UUID, input m.RequestApprovalSptpd, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.LogApproval

	data = m.LogApproval{
		Sptpd_Id: id,
		User_id:  *input.User_id,
		Status:   *input.Status,
	}

	var err error
	if data.Id == 0 {
		err = tx.Create(&data).Error
	} else {
		err = tx.Save(&data).Error
	}
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data datanir: "+err.Error(), data)
	}

	return rp.OKSimple{Data: data}, nil
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
	_, errLog := LogApproval(id, input, tx)
	if errLog != nil {
		return sh.SetError("request", "approval-data", source, "failed", "gagal loging approval data", errLog)
	}

	if kd == "03" && auth == 6 {
		//Dikembalikan Staf
		tempVal := "-1"
		data.Proses = &tempVal
	} else if kd == "04" && auth == 6 {
		//Kasubid
		tempVal := "0"
		data.Proses = &tempVal
	} else if kd == "05" && auth == 4 {
		//Dikembalikan Kasubid
		tempVal := "-1"
		data.Proses = &tempVal
	} else if kd == "06" && auth == 4 {
		//Kabid
		tempVal := "2"
		data.Proses = &tempVal
	} else if kd == "07" && auth == 3 {
		//Dikembalikan Kabid
	} else if kd == "08" && auth == 3 {
		//Cetak
		p, _ := rand.Prime(rand.Reader, 32)
		tempNumber := p.String()
		data.IdBilling = &tempNumber
	} else if kd == "10" && auth == 6 {
		//Lunas
	} else if kd == "12" && auth == 6 {
		//Kurang bayar
	} else if kd == "14" && auth == 6 {
		//Validasi
	} else if kd == "20" || kd == "21" || kd == "22" {
		//Batal
	} else if auth == 0 {
		//admin
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
