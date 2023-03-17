// service
package permohonan

import (
	"errors"
	"fmt"
	mad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	skec "github.com/bapenda-kota-malang/apin-backend/internal/services/kecamatan"
	skel "github.com/bapenda-kota-malang/apin-backend/internal/services/kelurahan"
	rpth "github.com/bapenda-kota-malang/apin-backend/pkg/reporthelper"
	strh "github.com/bapenda-kota-malang/apin-backend/pkg/stringhelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	sc "github.com/jinzhu/copier"
	"github.com/xuri/excelize/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	mopfas "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	mopbng "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	moptnh "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	moppbb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	reg "github.com/bapenda-kota-malang/apin-backend/internal/models/regpelayanan"
	sksk "github.com/bapenda-kota-malang/apin-backend/internal/models/sksk"
	mwppbb "github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajakpbb"
	oppbb "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakpbb"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/excelhelper"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	// oppbb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	// wppbb "github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajakpbb"
)

// /// Private funcs start here
const source = "permohonan"

///// Exported funcs start here

func Create(input m.PermohonanRequestDto) (any, error) {
	var (
		err                   error
		permohonanDetail      *m.PstDetail
		permohonanBaru        *m.PstDataOPBaru
		permohonanDetailInput *m.PstDetailInput
		permohonanPengurangan *m.PstPermohonanPengurangan
		nop                   *m.PermohonanNOP
		pembetulanSpptSKPSTP  *m.PembetulanSpptSKPSTP
		pembatalanSppt        *m.PembatalanSppt
		keputusanKeberatanPbb *m.KeputusanKeberatanPbb
		SPMKP                 *m.SPMKP
		SkSk                  *sksk.SkSk
	)

	noUrut := GetNoUrut(input)
	data := input.SetDataPermohonanRequestDtoTransformer(&noUrut)
	noPelayanan := data.GetDataPermohonanNoPelayanan()
	data.NoPelayanan = &noPelayanan

	err = a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data permohonan ke db satu if karena result dipakai sekali, +error
		if result := tx.Create(&data); result.Error != nil {
			return errors.New("penyimpanan data permohonan gagal")
		}

		permohonanBaru, permohonanDetailInput, permohonanPengurangan, nop = data.SetDataPermohonanTransformer(input)
		permohonanDetail = new(m.PstDetail)
		if err := sc.Copy(&permohonanDetail, &permohonanDetailInput); err != nil {
			return errors.New("set data permohonan detail gagal")
		}
		if result := tx.Create(&permohonanDetail); result.Error != nil {
			return errors.New("penyimpanan data permohonan detail gagal")
		}

		if *data.BundelPelayanan == m.JenisPelayanan[0] {
			if result := tx.Where("PermohonanId", data.Id).Create(&permohonanBaru); result.Error != nil {
				return errors.New("penyimpanan data permohonan baru gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[2] {
			pembetulanSpptSKPSTP = data.SetPembetulanSpptSKPSTP(*nop)
			if result := tx.Where("PermohonanId", data.Id).Create(&pembetulanSpptSKPSTP); result.Error != nil {
				return errors.New("penyimpanan data permohonan pembetulan SPPT/SKP/STP gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[3] {
			pembatalanSppt = data.SetPembatalanSppt(*nop)
			if result := tx.Where("PermohonanId", data.Id).Create(&pembatalanSppt); result.Error != nil {
				return errors.New("penyimpanan data permohonan pembatalan SPPT gagal")
			}
			SkSk = sksk.SetSkSk(data)
			if result := tx.Where("PermohonanId", data.Id).Create(&SkSk); result.Error != nil {
				return errors.New("penyimpanan data permohonan pembatalan SKP gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[8] {
			SPMKP = data.SetSPMKP(*nop)
			if result := tx.Where("PermohonanId", data.Id).Create(&SPMKP); result.Error != nil {
				return errors.New("penyimpanan data permohonan restitusi dan kompensasi gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[5] || *data.BundelPelayanan == m.JenisPelayanan[6] {
			keputusanKeberatanPbb = data.SetKeputusanKeberatanPbb(*nop)
			if result := tx.Where("PermohonanId", data.Id).Create(&keputusanKeberatanPbb); result.Error != nil {
				return errors.New("penyimpanan data permohonan keberatan gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[7] || *data.BundelPelayanan == m.JenisPelayanan[9] {
			if result := tx.Where("PermohonanId", data.Id).Create(&permohonanPengurangan); result.Error != nil {
				return errors.New("penyimpanan data permohonan pengurangan gagal")
			}
		}
		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
	}

	return rp.OKSimple{Data: t.II{
		"permohonan":            data,
		"permohonanBaru":        permohonanBaru,
		"permohonanDetail":      permohonanDetail,
		"permohonanPengurangan": permohonanPengurangan,
		"noPelayanan":           noPelayanan,
	}}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.PstPermohonan
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.PstPermohonan{}).
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

func DownloadExcelList(input m.FilterDto) (*excelize.File, error) {
	var data []m.DownloadListDTO
	result := a.DB.Model(&m.PstPermohonan{}).
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
			"B": "Status Kolektif",
			"C": "Nama",
			"D": "Jenis Pelayanan",
			"E": "NOP",
			"F": "Tanggal",
		})
		for i, v := range data {
			n := i + 1
			tmp = append(tmp, map[string]interface{}{
				"A": n,
				"B": *v.StatusKolektif,
				"C": *v.NamaPemohon,
				"D": *v.BundelPelayanan,
				"E": *v.NOP,
				"F": func() string {
					t, _ := time.Parse(time.RFC3339, *v.TanggalTerima)
					s := t.Format("2006-01-02")
					return s
				}(),
			})
		}
		return tmp
	}()
	return excelhelper.ExportList(excelData, "List Permohonan")
}

func GetNoUrut(input m.PermohonanRequestDto) string {
	var (
		checkdataori  *m.PstPermohonan
		checkdata     *reg.RegPstPermohonan
		tempNoUrutOri int    = 0
		tempNoUrutReg int    = 0
		tempNoUrut    int    = 0
		noUrut        string = "001"
	)

	year := strconv.Itoa(time.Now().Year())

	_ = a.DB.Where("TahunPelayanan", year).
		Where("BundelPelayanan", input.JenisPelayanan).
		Order("\"NoUrutPelayanan\" desc").
		First(&checkdataori)
	if checkdataori.Id != 0 {
		tempNoUrutOri, _ = strconv.Atoi(*checkdata.NoUrutPelayanan)
		tempNoUrutOri = tempNoUrutOri + 1
	}

	_ = a.DB.Where("TahunPelayanan", year).
		Where("BundelPelayanan", input.JenisPelayanan).
		Order("\"NoUrutPelayanan\" desc").
		First(&checkdata)
	if checkdata.Id != 0 {
		tempNoUrutReg, _ = strconv.Atoi(*checkdata.NoUrutPelayanan)
		tempNoUrutReg = tempNoUrut + 1
	}

	if tempNoUrutOri > tempNoUrutReg {
		tempNoUrut = tempNoUrutOri
	} else {
		tempNoUrut = tempNoUrutReg
	}

	noUrut = sh.FixedLengthString(3, fmt.Sprint(tempNoUrut))
	return noUrut
}

func GetStatusNOP(nop *string) (interface{}, error) {
	permohonannop := m.DecodeNOPPermohonan(nop)
	result, err := oppbb.GetDetailbyNop(*permohonannop)

	return result, err
}

func GetDetail(id int) (interface{}, error) {
	var (
		data                  *m.PstPermohonan
		permohonanBaru        *m.PstDataOPBaru
		permohonanDetail      *m.PstDetail
		permohonanPengurangan *m.PstPermohonanPengurangan
		// pembetulanSpptSKPSTP  *m.PembetulanSpptSKPSTP
		// pembatalanSppt        *m.PembatalanSppt
		// keputusanKeberatanPbb *m.KeputusanKeberatanPbb
		// SPMKP                 *m.SPMKP
		// SkSk                  *sksk.SkSk
	)
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", data)
	}
	if result := a.DB.Where("PermohonanId", data.Id).First(&permohonanDetail); result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data NOP Detail", permohonanDetail)
	}

	if *data.BundelPelayanan == m.JenisPelayanan[0] {
		if result := a.DB.Where("PermohonanId", data.Id).First(&permohonanBaru); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data NOP Baru", permohonanBaru)
		}
	} else if *data.BundelPelayanan == m.JenisPelayanan[7] || *data.BundelPelayanan == m.JenisPelayanan[9] {
		if result := a.DB.Where("PermohonanId", data.Id).First(&permohonanPengurangan); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
	}
	// else if *data.BundelPelayanan == m.JenisPelayanan[2] {
	// 	if result := a.DB.Where("PermohonanId", data.Id).First(&pembetulanSpptSKPSTP); result.Error != nil {
	// 		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pembetulan", pembetulanSpptSKPSTP)
	// 	}
	// } else if *data.BundelPelayanan == m.JenisPelayanan[3] {
	// 	if result := a.DB.Where("PermohonanId", data.Id).First(&pembatalanSppt); result.Error != nil {
	// 		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pembatalan SPPT", pembatalanSppt)
	// 	}
	// 	if result := a.DB.Where("PermohonanId", data.Id).First(&SkSk); result.Error != nil {
	// 		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pembatalan SKP", SkSk)
	// 	}
	// } else if *data.BundelPelayanan == m.JenisPelayanan[8] {
	// 	if result := a.DB.Where("PermohonanId", data.Id).First(&SPMKP); result.Error != nil {
	// 		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data restitusi dan kompensasi", SPMKP)
	// 	}
	// } else if *data.BundelPelayanan == m.JenisPelayanan[5] || *data.BundelPelayanan == m.JenisPelayanan[6] {
	// 	if result := a.DB.Where("PermohonanId", data.Id).First(&keputusanKeberatanPbb); result.Error != nil {
	// 		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data keberatan", keputusanKeberatanPbb)
	// 	}
	// }

	finalresult := data.SetPstPermohonanResponse()
	finalresult.PstDataOPBaru = permohonanBaru
	finalresult.PstDetail = permohonanDetail
	finalresult.PstPermohonanPengurangan = permohonanPengurangan

	return rp.OKSimple{Data: finalresult}, nil
}

func transformNOP(nop string) string {
	result := strings.Trim(nop, ".")
	return result
}

func GetDetailbyNop(nop string) (interface{}, error) {
	var (
		data                  *m.PstPermohonan
		permohonanBaru        *m.PstDataOPBaru
		permohonanDetail      *m.PstDetail
		permohonanPengurangan *m.PstPermohonanPengurangan
		oppbb                 *moppbb.ObjekPajakPbb
		wppbb                 *mwppbb.WajibPajakPbb
		opbng                 *mopbng.ObjekPajakBangunan
		optnh                 *moptnh.ObjekPajakBumi
		opfas                 *mopfas.FasilitasBangunan
		jpb                   any
	)
	tempNOP := transformNOP(nop)
	dec_nop := m.DecodeNOPPermohonan(&tempNOP)
	result := a.DB.Where("NOP", tempNOP).First(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", data)
	}
	if result := a.DB.Where("PermohonanId", data.Id).First(&permohonanDetail); result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data NOP Detail", permohonanDetail)
	}

	if *data.BundelPelayanan == m.JenisPelayanan[0] {
		if result := a.DB.Where("PermohonanId", data.Id).First(&permohonanBaru); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data NOP Baru", permohonanBaru)
		}
	} else if *data.BundelPelayanan == m.JenisPelayanan[7] || *data.BundelPelayanan == m.JenisPelayanan[9] {
		if result := a.DB.Where("PermohonanId", data.Id).First(&permohonanPengurangan); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
	}

	if result := a.DB.
		Where("Provinsi_Kode", dec_nop.PermohonanProvinsiID).
		Where("Daerah_Kode", dec_nop.PermohonanKotaID).
		Where("Kecamatan_Kode", dec_nop.PermohonanKecamatanID).
		Where("Kelurahan_Kode", dec_nop.PermohonanKelurahanID).
		Where("Blok_Kode", dec_nop.PermohonanBlokID).
		Where("NoUrut", dec_nop.NoUrutPemohon).
		Where("JenisOp", dec_nop.PemohonJenisOPID).
		First(&wppbb); result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
	}

	if result := a.DB.
		Where("Provinsi_Kode", dec_nop.PermohonanProvinsiID).
		Where("Daerah_Kode", dec_nop.PermohonanKotaID).
		Where("Kecamatan_Kode", dec_nop.PermohonanKecamatanID).
		Where("Kelurahan_Kode", dec_nop.PermohonanKelurahanID).
		Where("Blok_Kode", dec_nop.PermohonanBlokID).
		Where("NoUrut", dec_nop.NoUrutPemohon).
		Where("JenisOp", dec_nop.PemohonJenisOPID).
		Where("WajibPajakPbb_Id", wppbb.Id).
		First(&oppbb); result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
	}

	if result := a.DB.
		Where("Provinsi_Kode", dec_nop.PermohonanProvinsiID).
		Where("Daerah_Kode", dec_nop.PermohonanKotaID).
		Where("Kecamatan_Kode", dec_nop.PermohonanKecamatanID).
		Where("Kelurahan_Kode", dec_nop.PermohonanKelurahanID).
		Where("Blok_Kode", dec_nop.PermohonanBlokID).
		Where("NoUrut", dec_nop.NoUrutPemohon).
		Where("JenisOp", dec_nop.PemohonJenisOPID).
		First(&optnh); result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
	}

	if result := a.DB.
		Where("Provinsi_Kode", dec_nop.PermohonanProvinsiID).
		Where("Daerah_Kode", dec_nop.PermohonanKotaID).
		Where("Kecamatan_Kode", dec_nop.PermohonanKecamatanID).
		Where("Kelurahan_Kode", dec_nop.PermohonanKelurahanID).
		Where("Blok_Kode", dec_nop.PermohonanBlokID).
		Where("NoUrut", dec_nop.NoUrutPemohon).
		Where("JenisOp", dec_nop.PemohonJenisOPID).
		First(&opbng); result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
	}

	if result := a.DB.
		Where("NoBangunan", opbng.NoBangunan).
		First(&opfas); result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
	}

	if opbng.Jpb_Kode == "00" || opbng.Jpb_Kode == "01" || opbng.Jpb_Kode == "10" || opbng.Jpb_Kode == "11" {
	} else if opbng.Jpb_Kode == "02" {
		var jpbres *mopbng.Jpb2
		if result := a.DB.
			Where("NoBangunan", opbng.NoBangunan).
			First(&jpbres); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
		jpb = jpbres
	} else if opbng.Jpb_Kode == "03" {
		var jpbres *mopbng.Jpb3
		if result := a.DB.
			Where("NoBangunan", opbng.NoBangunan).
			First(&jpbres); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
		jpb = jpbres
	} else if opbng.Jpb_Kode == "04" {
		var jpbres *mopbng.Jpb4
		if result := a.DB.
			Where("NoBangunan", opbng.NoBangunan).
			First(&jpbres); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
		jpb = jpbres
	} else if opbng.Jpb_Kode == "05" {
		var jpbres *mopbng.Jpb5
		if result := a.DB.
			Where("NoBangunan", opbng.NoBangunan).
			First(&jpbres); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
		jpb = jpbres
	} else if opbng.Jpb_Kode == "06" {
		var jpbres *mopbng.Jpb6
		if result := a.DB.
			Where("NoBangunan", opbng.NoBangunan).
			First(&jpbres); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
		jpb = jpbres
	} else if opbng.Jpb_Kode == "07" {
		var jpbres *mopbng.Jpb7
		if result := a.DB.
			Where("NoBangunan", opbng.NoBangunan).
			First(&jpbres); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
		jpb = jpbres
	} else if opbng.Jpb_Kode == "08" {
		var jpbres *mopbng.Jpb8
		if result := a.DB.
			Where("NoBangunan", opbng.NoBangunan).
			First(&jpbres); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
		jpb = jpbres
	} else if opbng.Jpb_Kode == "09" {
		var jpbres *mopbng.Jpb9
		if result := a.DB.
			Where("NoBangunan", opbng.NoBangunan).
			First(&jpbres); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
		jpb = jpbres
	} else if opbng.Jpb_Kode == "12" {
		var jpbres *mopbng.Jpb12
		if result := a.DB.
			Where("NoBangunan", opbng.NoBangunan).
			First(&jpbres); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
		jpb = jpbres
	} else if opbng.Jpb_Kode == "13" {
		var jpbres *mopbng.Jpb13
		if result := a.DB.
			Where("NoBangunan", opbng.NoBangunan).
			First(&jpbres); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
		jpb = jpbres
	} else if opbng.Jpb_Kode == "14" {
		var jpbres *mopbng.Jpb14
		if result := a.DB.
			Where("NoBangunan", opbng.NoBangunan).
			First(&jpbres); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
		jpb = jpbres
	} else if opbng.Jpb_Kode == "15" {
		var jpbres *mopbng.Jpb15
		if result := a.DB.
			Where("NoBangunan", opbng.NoBangunan).
			First(&jpbres); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
		jpb = jpbres
	} else if opbng.Jpb_Kode == "16" {
		var jpbres *mopbng.Jpb16
		if result := a.DB.
			Where("NoBangunan", opbng.NoBangunan).
			First(&jpbres); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
		jpb = jpbres
	}

	finalresult := data.SetPstPermohonanResponse()
	finalresult.PstDataOPBaru = permohonanBaru
	finalresult.PstDetail = permohonanDetail
	finalresult.PstPermohonanPengurangan = permohonanPengurangan
	finalresult.PstOpjekPajakPBB = oppbb
	finalresult.PstWajibPajakPBB = wppbb
	finalresult.PstOPBumi = optnh
	finalresult.PstOPBangunan = opbng
	finalresult.PstFasilitasBangunan = opfas
	finalresult.PstJpb = &jpb

	return rp.OKSimple{Data: finalresult}, nil
}

func Update(id int, input m.PermohonanRequestDto) (interface{}, error) {
	var (
		data                  *m.PstPermohonan
		permohonanBaru        *m.PstDataOPBaru
		permohonanDetail      *m.PstDetail
		permohonanDetailInput *m.PstDetailInput
		permohonanPengurangan *m.PstPermohonanPengurangan
		nop                   *m.PermohonanNOP
		pembetulanSpptSKPSTP  *m.PembetulanSpptSKPSTP
		pembatalanSppt        *m.PembatalanSppt
		keputusanKeberatanPbb *m.KeputusanKeberatanPbb
		SPMKP                 *m.SPMKP
		SkSk                  *sksk.SkSk
	)

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "data permohonan tidak dapat ditemukan", input)
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload untuk permohonan", input)
	}

	rowsAffected := 0
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		if result := a.DB.Save(&data); result.Error != nil {
			return errors.New("gagal menyimpan data permohonan")
		}
		if result.RowsAffected > 0 {
			rowsAffected++
		}
		permohonanBaru, permohonanDetailInput, permohonanPengurangan, nop = data.SetDataPermohonanTransformer(input)
		result := a.DB.Where("PermohonanId", data.Id).First(&permohonanDetail)
		if result.RowsAffected == 0 {
			return errors.New("data permohonan detail tidak ditemukan")
		}
		if err := sc.Copy(&permohonanDetail, &permohonanDetailInput); err != nil {
			return errors.New("set data permohonan detail gagal")
		}
		if result := tx.Where("PermohonanId", data.Id).Updates(&permohonanDetail); result.Error != nil {
			return errors.New("penyimpanan data permohonan detail gagal")
		}

		if *data.BundelPelayanan == m.JenisPelayanan[0] {
			if result := tx.Where("PermohonanId", data.Id).Updates(&permohonanBaru); result.Error != nil {
				return errors.New("penyimpanan data permohonan baru gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[2] {
			pembetulanSpptSKPSTP = data.SetPembetulanSpptSKPSTP(*nop)
			if result := tx.Where("PermohonanId", data.Id).Updates(&pembetulanSpptSKPSTP); result.Error != nil {
				return errors.New("penyimpanan data permohonan pembetulan SPPT/SKP/STP gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[3] {
			pembatalanSppt = data.SetPembatalanSppt(*nop)
			if result := tx.Where("PermohonanId", data.Id).Updates(&pembatalanSppt); result.Error != nil {
				return errors.New("penyimpanan data permohonan pembatalan SPPT gagal")
			}
			SkSk = sksk.SetSkSk(*data)
			if result := tx.Where("PermohonanId", data.Id).Updates(&SkSk); result.Error != nil {
				return errors.New("penyimpanan data permohonan pembatalan SKP gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[8] {
			SPMKP = data.SetSPMKP(*nop)
			if result := tx.Where("PermohonanId", data.Id).Updates(&SPMKP); result.Error != nil {
				return errors.New("penyimpanan data permohonan restitusi dan kompensasi gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[5] || *data.BundelPelayanan == m.JenisPelayanan[6] {
			keputusanKeberatanPbb = data.SetKeputusanKeberatanPbb(*nop)
			if result := tx.Where("PermohonanId", data.Id).Updates(&keputusanKeberatanPbb); result.Error != nil {
				return errors.New("penyimpanan data permohonan keberatan gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[7] || *data.BundelPelayanan == m.JenisPelayanan[9] {
			if result := tx.Where("PermohonanId", data.Id).Updates(&permohonanPengurangan); result.Error != nil {
				return errors.New("penyimpanan data permohonan pengurangan gagal")
			}
		}

		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", err.Error(), input)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(rowsAffected),
		},
		Data: t.II{
			"permohonan":            data,
			"permohonanBaru":        permohonanBaru,
			"permohonanDetail":      permohonanDetail,
			"permohonanPengurangan": permohonanPengurangan,
		},
	}, nil
}

func UpdateStatus(id int, input m.PermohonanRequestDto) (interface{}, error) {
	var (
		data             *m.PstPermohonan
		permohonanDetail *m.PstDetail
	)

	_ = a.DB.First(&data, id)

	result := a.DB.Where("PermohonanId", data.Id).First(&permohonanDetail)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "data permohonan tidak dapat ditemukan", input)
	}
	if err := sc.Copy(&permohonanDetail, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload untuk permohonan", input)
	}
	permohonanDetail.NIP = input.NIP
	permohonanDetail.Notes = input.CatatanPenyerahan
	if input.TanggalPenyerahan != nil {
		tempTglPenyerahan := (*datatypes.Date)(gormhelper.StringToDate(*input.TanggalPenyerahan))
		permohonanDetail.TanggalPenyerahan = tempTglPenyerahan
	}

	rowsAffected := 0
	if result := a.DB.Where("PermohonanId", data.Id).Updates(&permohonanDetail); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal melakukan update status detil permohonan", input)
	}
	if result.RowsAffected > 0 {
		rowsAffected++
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(rowsAffected),
		},
		Data: t.II{
			"permohonan":       data,
			"permohonanDetail": permohonanDetail,
		},
	}, nil
}

func Delete(id int) (interface{}, error) {
	var (
		data                  *m.PstPermohonan
		permohonanBaru        *m.PstDataOPBaru
		permohonanDetail      *m.PstDetail
		permohonanPengurangan *m.PstPermohonanPengurangan
		pembetulanSpptSKPSTP  *m.PembetulanSpptSKPSTP
		pembatalanSppt        *m.PembatalanSppt
		keputusanKeberatanPbb *m.KeputusanKeberatanPbb
		SPMKP                 *m.SPMKP
		SkSk                  *sksk.SkSk
	)

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	result = a.DB.Delete(&data, id)
	status := "deleted"
	if result.RowsAffected == 0 {
		data = nil
		status = "no deletion"
	} else {
		_ = a.DB.Delete(&permohonanDetail, int(*permohonanDetail.PermohonanId) == id)
		if *data.BundelPelayanan == m.JenisPelayanan[0] {
			_ = a.DB.Delete(&permohonanBaru, int(*permohonanBaru.PermohonanId) == id)
		} else if *data.BundelPelayanan == m.JenisPelayanan[2] {
			_ = a.DB.Delete(&pembetulanSpptSKPSTP, int(*pembetulanSpptSKPSTP.PermohonanId) == id)
		} else if *data.BundelPelayanan == m.JenisPelayanan[3] {
			_ = a.DB.Delete(&pembatalanSppt, int(*pembatalanSppt.PermohonanId) == id)
			_ = a.DB.Delete(&SkSk, int(*SkSk.PermohonanId) == id)
		} else if *data.BundelPelayanan == m.JenisPelayanan[8] {
			_ = a.DB.Delete(&SPMKP, int(*SPMKP.PermohonanId) == id)
		} else if *data.BundelPelayanan == m.JenisPelayanan[5] || *data.BundelPelayanan == m.JenisPelayanan[6] {
			_ = a.DB.Delete(&keputusanKeberatanPbb, int(*keputusanKeberatanPbb.PermohonanId) == id)
		} else if *data.BundelPelayanan == m.JenisPelayanan[7] || *data.BundelPelayanan == m.JenisPelayanan[9] {
			_ = a.DB.Delete(&permohonanPengurangan, int(*permohonanPengurangan.PermohonanId) == id)
		}
	}

	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil
}

type ResponseFile struct {
	ID       int    `json:"id"`
	FileName string `json:"fileName"`
}

func DownloadPdf(id int) (interface{}, error) {
	var (
		data                  *m.PstPermohonan
		permohonanBaru        *m.PstDataOPBaru
		permohonanDetail      *m.PstDetail
		permohonanPengurangan *m.PstPermohonanPengurangan
		// pembetulanSpptSKPSTP  *m.PembetulanSpptSKPSTP
		// pembatalanSppt        *m.PembatalanSppt
		// keputusanKeberatanPbb *m.KeputusanKeberatanPbb
		// SPMKP                 *m.SPMKP
		// SkSk                  *sksk.SkSk
	)
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", data)
	}

	if result := a.DB.Where("PermohonanId", data.Id).First(&permohonanDetail); result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data NOP Detail", permohonanDetail)
	}

	if *data.BundelPelayanan == m.JenisPelayanan[0] {
		if result := a.DB.Where("PermohonanId", data.Id).First(&permohonanBaru); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data NOP Baru", permohonanBaru)
		}
	} else if *data.BundelPelayanan == m.JenisPelayanan[7] || *data.BundelPelayanan == m.JenisPelayanan[9] {
		if result := a.DB.Where("PermohonanId", data.Id).First(&permohonanPengurangan); result.Error != nil {
			return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
		}
	}

	dataDetail, err := GetDetail(id)
	finalresultTmp := dataDetail.(rp.OKSimple).Data.(m.PstPermohonanResponse)

	idKel, err1 := strconv.Atoi(*finalresultTmp.PstDetail.PermohonanKelurahanID)
	idKec, err2 := strconv.Atoi(*finalresultTmp.PstDetail.PermohonanKecamatanID)
	var (
		dataKelurahan = &mad.Kelurahan{}
		dataKecamatan = &mad.Kecamatan{}
	)
	if err1 == nil {
		tmpKel, _ := skel.GetDetail(idKel)
		dataKelurahan = tmpKel.(rp.OKSimple).Data.(*mad.Kelurahan)
	}
	if err2 == nil {
		tmpKec, _ := skec.GetDetail(idKec)
		dataKecamatan = tmpKec.(rp.OKSimple).Data.(*mad.Kecamatan)
	}
	var rawNop = strh.NullToAny(finalresultTmp.NOP, "-")
	var rawJP = strh.NullToAny(finalresultTmp.BundelPelayanan, "-")
	var dataJP = m.FindJenisPelayananByID(rawJP)
	var rawBerkas = strh.NullToAny(finalresultTmp.PenerimaanBerkas, "-")
	var listBerkas = strings.Split(rawBerkas, ",")

	var jpList = []*rpth.TCheckList{}
	for id, jp := range m.JenisBerkasPermohonan {
		var newVal = &rpth.TCheckList{
			Nomer:  strconv.Itoa(id + 1),
			Text:   jp,
			Status: "",
		}
		jpList = append(jpList, newVal)
	}
	print(jpList)
	for _, berkas := range listBerkas {
		idxBerkas, error := strconv.Atoi(strh.IsNumeric(berkas))

		if error == nil {
			jpList[idxBerkas].Status = "x"
		}
	}
	var nopDot = m.DecodeNOPPermohonan(&rawNop).GetNopDotFormat()
	var finalresult = map[string]interface{}{
		"Data": map[string]string{
			"noPelayanan":       strh.NullToAny(finalresultTmp.NoPelayanan, "-"),
			"nop":               rawNop,
			"bundlePelayanan":   strh.NullToAny(finalresultTmp.BundelPelayanan, "-"),
			"noUrutPelayanan":   strh.NullToAny(finalresultTmp.NoUrutPelayanan, "-"),
			"noSuratPermohonan": strh.NullToAny(finalresultTmp.NoSuratPermohonan, "-"),
			"namaWP":            strh.NullToAny(finalresultTmp.NamaPemohon, "-"),
			"letakOP":           strh.NullToAny(finalresultTmp.AlamatPemohon, "-"),
			"keterangan":        strh.NullToAny(finalresultTmp.Keterangan, "-"),
			"catatan":           strh.NullToAny(finalresultTmp.CatatanPermohonan, "-"),
			"nip":               strh.NullToAny(finalresultTmp.NIP, "-"),
		},
		"Custom": map[string]string{
			"NOP":            nopDot,
			"Kelurahan":      strh.NullToAny(&dataKelurahan.Nama, "-"),
			"Kecamatan":      strh.NullToAny(&dataKecamatan.Nama, "-"),
			"JenisPelayanan": strh.NullToAny(&dataJP.Name, "-"),
			"TglSelesai":     th.GetDateFromUTCDatetime(finalresultTmp.PstDetail.TanggalSelesai),
			"TglPelayanan":   th.GetDateFromUTCDatetime(finalresultTmp.PstDetail.TanggalPenyerahan),
		},
		"Checklist": map[string][]*rpth.TCheckList{
			"JenisBerkas": jpList,
		},
	}

	uuid, err := sh.GetUuidv4()
	if err != nil {
		return nil, err
	}
	fileName := sh.GenerateFilename("permohonan", uuid, 0, "pdf")

	outFile := filepath.Join(sh.GetPdfPath(), fileName)
	if err := GeneratePdf(outFile, finalresult); err != nil {
		return nil, err
	}
	_r := &ResponseFile{
		ID:       id,
		FileName: outFile,
	}
	return rp.OKSimple{Data: _r}, nil
}
