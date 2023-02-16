// service
package permohonan

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	sc "github.com/jinzhu/copier"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	mopfas "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	mkepkebpbb "github.com/bapenda-kota-malang/apin-backend/internal/models/keberatan"
	mopbng "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	moptnh "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	moppbb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	mwppbb "github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajakpbb"

	rmopfas "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	rmopbng "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbangunan"
	rmoptnh "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbumi"
	rmoppbb "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakpbb"
	rmwppbb "github.com/bapenda-kota-malang/apin-backend/internal/models/regwajibpajakpbb"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	reg "github.com/bapenda-kota-malang/apin-backend/internal/models/regpelayanan"
	sksk "github.com/bapenda-kota-malang/apin-backend/internal/models/sksk"
	skepkebpbb "github.com/bapenda-kota-malang/apin-backend/internal/services/keberatan/keputusankeberatanpbb"
	oppbb "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakpbb"
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
		kepKebPbbDto          *mkepkebpbb.CreateDtoKepKebPbb
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
			if result := tx.Create(&permohonanBaru); result.Error != nil {
				return errors.New("penyimpanan data permohonan baru gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[2] {
			pembetulanSpptSKPSTP = data.SetPembetulanSpptSKPSTP(*nop)
			if result := tx.Create(&pembetulanSpptSKPSTP); result.Error != nil {
				return errors.New("penyimpanan data permohonan pembetulan SPPT/SKP/STP gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[3] {
			pembatalanSppt = data.SetPembatalanSppt(*nop)
			if result := tx.Create(&pembatalanSppt); result.Error != nil {
				return errors.New("penyimpanan data permohonan pembatalan SPPT gagal")
			}
			SkSk = sksk.SetSkSk(data)
			if result := tx.Create(&SkSk); result.Error != nil {
				return errors.New("penyimpanan data permohonan pembatalan SKP gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[8] {
			SPMKP = data.SetSPMKP(*nop)
			if result := tx.Create(&SPMKP); result.Error != nil {
				return errors.New("penyimpanan data permohonan restitusi dan kompensasi gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[5] || *data.BundelPelayanan == m.JenisPelayanan[6] {
			if err := sc.Copy(&kepKebPbbDto, &data); err != nil {
				return errors.New("set data keputusan keberatan pbb  gagal")
			}
			if _, err := skepkebpbb.Create(*kepKebPbbDto, tx); err != nil {
				return errors.New("penyimpanan data keputusan keberatan pbb gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[7] || *data.BundelPelayanan == m.JenisPelayanan[9] {
			if result := tx.Create(&permohonanPengurangan); result.Error != nil {
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
		lampiran              *m.PstLampiran
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
	if result := a.DB.Where("PermohonanId", data.Id).First(&lampiran); result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data Lampiran Baru", lampiran)
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
	finalresult.PstLampiran = lampiran
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
		kepKebPbbDto          *mkepkebpbb.UpdateDtoKepKebPbb
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
			// copy to dto
			if err := sc.CopyWithOption(&kepKebPbbDto, &data, sc.Option{IgnoreEmpty: true}); err != nil {
				return errors.New("set data keputusan keberatan pbb gagal")
			}
			// search data
			var dataKepKebPbb mkepkebpbb.KeputusanKeberatanPbb
			if res := tx.Select("Id").Where("PermohonanId", data.Id).First(&dataKepKebPbb); res.Error != nil {
				return res.Error
			}
			// update
			if _, err := skepkebpbb.Update(int(dataKepKebPbb.Id), *kepKebPbbDto, tx); err != nil {
				return errors.New("penyimpanan data keputusan keberatan pbb gagal")
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
		keputusanKeberatanPbb *mkepkebpbb.KeputusanKeberatanPbb
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

func LogApproval(id uint64, input m.RequestPSTLogApproval, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.PstLogApproval

	data = m.PstLogApproval{
		Permohonan_Id: id,
		User_id:       input.User_id,
		Catatan:       input.Catatan,
		Status:        input.Status,
	}

	var err error
	if data.Id == 0 {
		err = tx.Create(&data).Error
	} else {
		err = tx.Save(&data).Error
	}
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data data: "+err.Error(), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func CopyDataReg(id uint64, nopInput string, jenisPengurangan string, kelasbangunan string, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var (
		err error

		data                  m.PstPermohonan
		dataReg               reg.RegPstPermohonan
		nop                   *m.PermohonanNOP
		permohonanDetail      *m.PstDetail
		permohonanBaru        *m.PstDataOPBaru
		permohonanPengurangan *m.PstPermohonanPengurangan
		pembetulanSpptSKPSTP  *m.PembetulanSpptSKPSTP
		keputusanKeberatanPbb *m.KeputusanKeberatanPbb

		regPermohonanDetail      *reg.RegPstDetail
		regPermohonanBaru        *reg.RegPstDataOPBaru
		regPermohonanPengurangan *reg.RegPstPermohonanPengurangan
		regPembetulanSpptSKPSTP  *reg.RegPembetulanSpptSKPSTP
		regKeputusanKeberatanPbb *reg.RegKeputusanKeberatanPbb

		oppbb *moppbb.ObjekPajakPbb
		wppbb *mwppbb.WajibPajakPbb
		optnh *moptnh.ObjekPajakBumi
		opbng *mopbng.ObjekPajakBangunan
		opfas *mopfas.FasilitasBangunan

		roppbb *rmoppbb.RegObjekPajakPbb
		rwppbb *rmwppbb.RegWajibPajakPbb
		roptnh *rmoptnh.RegObjekPajakBumi
		ropbng []rmopbng.RegObjekPajakBangunan
		ropfas []rmopfas.RegFasilitasBangunan
	)

	nop = m.DecodeNOPPermohonan(&nopInput)
	areaCode := nop.PermohonanProvinsiID + nop.PermohonanKotaID + nop.PermohonanKecamatanID + nop.PermohonanKelurahanID

	result := a.DB.First(&dataReg, id)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data: "+err.Error(), data)
	}

	idPermohonan := dataReg.Id

	if err := sc.Copy(&data, &dataReg); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	err = a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data permohonan ke db satu if karena result dipakai sekali, +error
		if result := tx.Create(&data); result.Error != nil {
			return errors.New("penyimpanan data permohonan gagal")
		}

		_ = a.DB.Where("PermohonanId", idPermohonan).First(&regPermohonanDetail)
		if err := sc.Copy(&permohonanDetail, &regPermohonanDetail); err != nil {
			return errors.New("set data permohonan detail gagal")
		}
		if err := sc.CopyWithOption(&permohonanDetail, &nop, sc.Option{IgnoreEmpty: true}); err != nil {
			return errors.New("set data permohonan detail gagal")
		}
		permohonanDetail.PermohonanId = &data.Id

		if result := tx.Create(&permohonanDetail); result.Error != nil {
			return errors.New("penyimpanan data permohonan detail gagal")
		}

		if *data.BundelPelayanan == m.JenisPelayanan[0] {
			_ = a.DB.Where("PermohonanId", idPermohonan).First(&regPermohonanBaru)
			if err := sc.Copy(&permohonanBaru, &regPermohonanBaru); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			if err := sc.CopyWithOption(&permohonanBaru, &nop, sc.Option{IgnoreEmpty: true}); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			regPermohonanBaru.PermohonanId = &data.Id

			if result := tx.Create(&permohonanBaru); result.Error != nil {
				return errors.New("penyimpanan data permohonan baru gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[2] {
			_ = a.DB.Where("PermohonanId", idPermohonan).First(&regPembetulanSpptSKPSTP)
			if err := sc.Copy(&pembetulanSpptSKPSTP, &regPembetulanSpptSKPSTP); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			if err := sc.CopyWithOption(&pembetulanSpptSKPSTP, &nop, sc.Option{IgnoreEmpty: true}); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			pembetulanSpptSKPSTP.PermohonanId = &data.Id

			if result := tx.Create(&pembetulanSpptSKPSTP); result.Error != nil {
				return errors.New("penyimpanan data permohonan pembetulan SPPT/SKP/STP gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[5] || *data.BundelPelayanan == m.JenisPelayanan[6] {
			_ = a.DB.Where("PermohonanId", idPermohonan).First(&regKeputusanKeberatanPbb)
			if err := sc.Copy(&keputusanKeberatanPbb, &regKeputusanKeberatanPbb); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			if err := sc.CopyWithOption(&keputusanKeberatanPbb, &nop, sc.Option{IgnoreEmpty: true}); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			keputusanKeberatanPbb.PermohonanId = &data.Id

			if result := tx.Create(&keputusanKeberatanPbb); result.Error != nil {
				return errors.New("penyimpanan data permohonan keberatan gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[7] || *data.BundelPelayanan == m.JenisPelayanan[9] {
			_ = a.DB.Where("PermohonanId", idPermohonan).First(&regPermohonanPengurangan)
			if err := sc.Copy(&permohonanPengurangan, &regPermohonanPengurangan); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			if err := sc.CopyWithOption(&permohonanPengurangan, &nop, sc.Option{IgnoreEmpty: true}); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			permohonanPengurangan.PermohonanId = &data.Id
			permohonanPengurangan.JenisPengurangan = &jenisPengurangan

			if result := tx.Create(&permohonanPengurangan); result.Error != nil {
				return errors.New("penyimpanan data permohonan pengurangan gagal")
			}
		}

		// roppbb                 *roppbb.RegObjekPajakPbb
		_ = tx.Where("PstPermohonan_id", idPermohonan).First(&roppbb)
		if err := sc.Copy(&oppbb, &roppbb); err != nil {
			return errors.New("set data permohonan detail gagal")
		}
		oppbb.Provinsi_Kode = &nop.PermohonanProvinsiID
		oppbb.Daerah_Kode = &nop.PermohonanKotaID
		oppbb.Kecamatan_Kode = &nop.PermohonanKecamatanID
		oppbb.Kelurahan_Kode = &nop.PermohonanKelurahanID
		oppbb.Blok_Kode = &nop.PermohonanBlokID
		oppbb.NoUrut = &nop.NoUrutPemohon
		oppbb.JenisOp = &nop.PemohonJenisOPID
		oppbb.Area_Kode = &areaCode
		// oppbb.PstPermohonan_id = &data.Id

		if result := tx.Create(&oppbb); result.Error != nil {
			return errors.New("penyimpanan data permohonan pengurangan gagal")
		}

		// rwppbb                 *rwppbb.RegWajibPajakPbb
		_ = tx.Where("Id", roppbb.RegWajibPajakPbb_Id).First(&rwppbb)
		if err := sc.Copy(&wppbb, &rwppbb); err != nil {
			return errors.New("set data permohonan detail gagal")
		}
		if result := tx.Create(&wppbb); result.Error != nil {
			return errors.New("penyimpanan data permohonan pengurangan gagal")
		}

		// roptnh                 *roptnh.RegObjekPajakBumi
		_ = tx.Where("PstPermohonan_id", idPermohonan).First(&roptnh)
		if err := sc.Copy(&optnh, &roptnh); err != nil {
			return errors.New("set data permohonan detail gagal")
		}
		// optnh.PstPermohonan_id = &data.Id
		if result := tx.Create(&optnh); result.Error != nil {
			return errors.New("penyimpanan data permohonan pengurangan gagal")
		}

		// ropbng                 *ropbng.RegObjekPajakBangunan
		resBang := a.DB.Where("PstPermohonan_id", idPermohonan).Find(&ropbng)
		if resBang.RowsAffected > 0 {
			for _, val := range ropbng {
				if err := sc.Copy(&opbng, &val); err != nil {
					return errors.New("set data permohonan detail gagal")
				}
				// opbng.PstPermohonan_id = &data.Id
				if result := tx.Create(&opbng); result.Error != nil {
					return errors.New("penyimpanan data permohonan pengurangan gagal")
				}
			}

			// ropfas                 *ropfas.RegFasilitasBangunan
			resFas := a.DB.Where("PstPermohonan_id", idPermohonan).Find(&ropfas)
			if resFas.RowsAffected > 0 {
				for _, val := range ropfas {
					if err := sc.Copy(&opfas, &val); err != nil {
						return errors.New("set data permohonan detail gagal")
					}
					opfas.Provinsi_Kode = &nop.PermohonanProvinsiID
					opfas.Daerah_Kode = &nop.PermohonanKotaID
					opfas.Kecamatan_Kode = &nop.PermohonanKecamatanID
					opfas.Kelurahan_Kode = &nop.PermohonanKelurahanID
					opfas.Blok_Kode = &nop.PermohonanBlokID
					opfas.NoUrut = &nop.NoUrutPemohon
					opfas.JenisOp = &nop.PemohonJenisOPID
					opfas.Area_Kode = &areaCode
					// opfas.PstPermohonan_id = &data.Id
					if result := tx.Create(&opfas); result.Error != nil {
						return errors.New("penyimpanan data permohonan pengurangan gagal")
					}
				}
			}

			if opbng.Jpb_Kode == "00" || opbng.Jpb_Kode == "01" || opbng.Jpb_Kode == "10" || opbng.Jpb_Kode == "11" {
			} else if opbng.Jpb_Kode == "02" {
				var jpbres *mopbng.Jpb2
				var jpbreg *rmopbng.RegJpb2
				if result := a.DB.
					Where("NoBangunan", opbng.NoBangunan).
					First(&jpbreg); result.Error != nil {
					return result.Error
				}
				if err := sc.Copy(&jpbres, &jpbreg); err != nil {
					return errors.New("set data gagal")
				}

				jpbres.Provinsi_Kode = &nop.PermohonanProvinsiID
				jpbres.Daerah_Kode = &nop.PermohonanKotaID
				jpbres.Kecamatan_Kode = &nop.PermohonanKecamatanID
				jpbres.Kelurahan_Kode = &nop.PermohonanKelurahanID
				jpbres.Blok_Kode = &nop.PermohonanBlokID
				jpbres.NoUrut = &nop.NoUrutPemohon
				jpbres.JenisOp = &nop.PemohonJenisOPID
				jpbres.Area_Kode = &areaCode
				jpbres.KelasBangunan2 = mopbng.KelasBangunan(kelasbangunan)

				if result := tx.Create(&jpbres); result.Error != nil {
					return errors.New("penyimpanan data gagal")
				}
			} else if opbng.Jpb_Kode == "03" {
				var jpbres *mopbng.Jpb3
				var jpbreg *rmopbng.RegJpb3
				if result := a.DB.
					Where("NoBangunan", opbng.NoBangunan).
					First(&jpbreg); result.Error != nil {
					return result.Error
				}
				if err := sc.Copy(&jpbres, &jpbreg); err != nil {
					return errors.New("set data gagal")
				}

				jpbres.Provinsi_Kode = &nop.PermohonanProvinsiID
				jpbres.Daerah_Kode = &nop.PermohonanKotaID
				jpbres.Kecamatan_Kode = &nop.PermohonanKecamatanID
				jpbres.Kelurahan_Kode = &nop.PermohonanKelurahanID
				jpbres.Blok_Kode = &nop.PermohonanBlokID
				jpbres.NoUrut = &nop.NoUrutPemohon
				jpbres.JenisOp = &nop.PemohonJenisOPID
				jpbres.Area_Kode = &areaCode

				if result := tx.Create(&jpbres); result.Error != nil {
					return errors.New("penyimpanan data gagal")
				}
			} else if opbng.Jpb_Kode == "04" {
				var jpbres *mopbng.Jpb4
				var jpbreg *rmopbng.RegJpb4
				if result := a.DB.
					Where("NoBangunan", opbng.NoBangunan).
					First(&jpbreg); result.Error != nil {
					return result.Error
				}
				if err := sc.Copy(&jpbres, &jpbreg); err != nil {
					return errors.New("set data gagal")
				}

				jpbres.Provinsi_Kode = &nop.PermohonanProvinsiID
				jpbres.Daerah_Kode = &nop.PermohonanKotaID
				jpbres.Kecamatan_Kode = &nop.PermohonanKecamatanID
				jpbres.Kelurahan_Kode = &nop.PermohonanKelurahanID
				jpbres.Blok_Kode = &nop.PermohonanBlokID
				jpbres.NoUrut = &nop.NoUrutPemohon
				jpbres.JenisOp = &nop.PemohonJenisOPID
				jpbres.Area_Kode = &areaCode
				jpbres.KelasBangunan4 = mopbng.KelasBangunan(kelasbangunan)

				if result := tx.Create(&jpbres); result.Error != nil {
					return errors.New("penyimpanan data gagal")
				}
			} else if opbng.Jpb_Kode == "05" {
				var jpbres *mopbng.Jpb5
				var jpbreg *rmopbng.RegJpb5
				if result := a.DB.
					Where("NoBangunan", opbng.NoBangunan).
					First(&jpbreg); result.Error != nil {
					return result.Error
				}
				if err := sc.Copy(&jpbres, &jpbreg); err != nil {
					return errors.New("set data gagal")
				}

				jpbres.Provinsi_Kode = &nop.PermohonanProvinsiID
				jpbres.Daerah_Kode = &nop.PermohonanKotaID
				jpbres.Kecamatan_Kode = &nop.PermohonanKecamatanID
				jpbres.Kelurahan_Kode = &nop.PermohonanKelurahanID
				jpbres.Blok_Kode = &nop.PermohonanBlokID
				jpbres.NoUrut = &nop.NoUrutPemohon
				jpbres.JenisOp = &nop.PemohonJenisOPID
				jpbres.Area_Kode = &areaCode
				jpbres.KelasBangunan5 = mopbng.KelasBangunan(kelasbangunan)

				if result := tx.Create(&jpbres); result.Error != nil {
					return errors.New("penyimpanan data gagal")
				}
			} else if opbng.Jpb_Kode == "06" {
				var jpbres *mopbng.Jpb6
				var jpbreg *rmopbng.RegJpb6
				if result := a.DB.
					Where("NoBangunan", opbng.NoBangunan).
					First(&jpbreg); result.Error != nil {
					return result.Error
				}
				if err := sc.Copy(&jpbres, &jpbreg); err != nil {
					return errors.New("set data gagal")
				}

				jpbres.Provinsi_Kode = &nop.PermohonanProvinsiID
				jpbres.Daerah_Kode = &nop.PermohonanKotaID
				jpbres.Kecamatan_Kode = &nop.PermohonanKecamatanID
				jpbres.Kelurahan_Kode = &nop.PermohonanKelurahanID
				jpbres.Blok_Kode = &nop.PermohonanBlokID
				jpbres.NoUrut = &nop.NoUrutPemohon
				jpbres.JenisOp = &nop.PemohonJenisOPID
				jpbres.Area_Kode = &areaCode
				jpbres.KelasBangunan6 = mopbng.KelasBangunan(kelasbangunan)

				if result := tx.Create(&jpbres); result.Error != nil {
					return errors.New("penyimpanan data gagal")
				}
			} else if opbng.Jpb_Kode == "07" {
				var jpbres *mopbng.Jpb7
				var jpbreg *rmopbng.RegJpb7
				if result := a.DB.
					Where("NoBangunan", opbng.NoBangunan).
					First(&jpbreg); result.Error != nil {
					return result.Error
				}
				if err := sc.Copy(&jpbres, &jpbreg); err != nil {
					return errors.New("set data gagal")
				}

				jpbres.Provinsi_Kode = &nop.PermohonanProvinsiID
				jpbres.Daerah_Kode = &nop.PermohonanKotaID
				jpbres.Kecamatan_Kode = &nop.PermohonanKecamatanID
				jpbres.Kelurahan_Kode = &nop.PermohonanKelurahanID
				jpbres.Blok_Kode = &nop.PermohonanBlokID
				jpbres.NoUrut = &nop.NoUrutPemohon
				jpbres.JenisOp = &nop.PemohonJenisOPID
				jpbres.Area_Kode = &areaCode

				if result := tx.Create(&jpbres); result.Error != nil {
					return errors.New("penyimpanan data gagal")
				}
			} else if opbng.Jpb_Kode == "08" {
				var jpbres *mopbng.Jpb8
				var jpbreg *rmopbng.RegJpb8
				if result := a.DB.
					Where("NoBangunan", opbng.NoBangunan).
					First(&jpbreg); result.Error != nil {
					return result.Error
				}
				if err := sc.Copy(&jpbres, &jpbreg); err != nil {
					return errors.New("set data gagal")
				}

				jpbres.Provinsi_Kode = &nop.PermohonanProvinsiID
				jpbres.Daerah_Kode = &nop.PermohonanKotaID
				jpbres.Kecamatan_Kode = &nop.PermohonanKecamatanID
				jpbres.Kelurahan_Kode = &nop.PermohonanKelurahanID
				jpbres.Blok_Kode = &nop.PermohonanBlokID
				jpbres.NoUrut = &nop.NoUrutPemohon
				jpbres.JenisOp = &nop.PemohonJenisOPID
				jpbres.Area_Kode = &areaCode

				if result := tx.Create(&jpbres); result.Error != nil {
					return errors.New("penyimpanan data gagal")
				}
			} else if opbng.Jpb_Kode == "09" {
				var jpbres *mopbng.Jpb9
				var jpbreg *rmopbng.RegJpb9
				if result := a.DB.
					Where("NoBangunan", opbng.NoBangunan).
					First(&jpbreg); result.Error != nil {
					return result.Error
				}
				if err := sc.Copy(&jpbres, &jpbreg); err != nil {
					return errors.New("set data gagal")
				}

				jpbres.Provinsi_Kode = &nop.PermohonanProvinsiID
				jpbres.Daerah_Kode = &nop.PermohonanKotaID
				jpbres.Kecamatan_Kode = &nop.PermohonanKecamatanID
				jpbres.Kelurahan_Kode = &nop.PermohonanKelurahanID
				jpbres.Blok_Kode = &nop.PermohonanBlokID
				jpbres.NoUrut = &nop.NoUrutPemohon
				jpbres.JenisOp = &nop.PemohonJenisOPID
				jpbres.Area_Kode = &areaCode
				jpbres.KelasBangunan9 = mopbng.KelasBangunan(kelasbangunan)

				if result := tx.Create(&jpbres); result.Error != nil {
					return errors.New("penyimpanan data gagal")
				}
			} else if opbng.Jpb_Kode == "12" {
				var jpbres *mopbng.Jpb12
				var jpbreg *rmopbng.RegJpb12
				if result := a.DB.
					Where("NoBangunan", opbng.NoBangunan).
					First(&jpbreg); result.Error != nil {
					return result.Error
				}
				if err := sc.Copy(&jpbres, &jpbreg); err != nil {
					return errors.New("set data gagal")
				}

				jpbres.Provinsi_Kode = &nop.PermohonanProvinsiID
				jpbres.Daerah_Kode = &nop.PermohonanKotaID
				jpbres.Kecamatan_Kode = &nop.PermohonanKecamatanID
				jpbres.Kelurahan_Kode = &nop.PermohonanKelurahanID
				jpbres.Blok_Kode = &nop.PermohonanBlokID
				jpbres.NoUrut = &nop.NoUrutPemohon
				jpbres.JenisOp = &nop.PemohonJenisOPID
				jpbres.Area_Kode = &areaCode

				if result := tx.Create(&jpbres); result.Error != nil {
					return errors.New("penyimpanan data gagal")
				}
			} else if opbng.Jpb_Kode == "13" {
				var jpbres *mopbng.Jpb13
				var jpbreg *rmopbng.RegJpb13
				if result := a.DB.
					Where("NoBangunan", opbng.NoBangunan).
					First(&jpbreg); result.Error != nil {
					return result.Error
				}
				if err := sc.Copy(&jpbres, &jpbreg); err != nil {
					return errors.New("set data gagal")
				}

				jpbres.Provinsi_Kode = &nop.PermohonanProvinsiID
				jpbres.Daerah_Kode = &nop.PermohonanKotaID
				jpbres.Kecamatan_Kode = &nop.PermohonanKecamatanID
				jpbres.Kelurahan_Kode = &nop.PermohonanKelurahanID
				jpbres.Blok_Kode = &nop.PermohonanBlokID
				jpbres.NoUrut = &nop.NoUrutPemohon
				jpbres.JenisOp = &nop.PemohonJenisOPID
				jpbres.Area_Kode = &areaCode
				jpbres.KelasBangunan13 = mopbng.KelasBangunan(kelasbangunan)

				if result := tx.Create(&jpbres); result.Error != nil {
					return errors.New("penyimpanan data gagal")
				}
			} else if opbng.Jpb_Kode == "14" {
				var jpbres *mopbng.Jpb14
				var jpbreg *rmopbng.RegJpb14
				if result := a.DB.
					Where("NoBangunan", opbng.NoBangunan).
					First(&jpbreg); result.Error != nil {
					return result.Error
				}
				if err := sc.Copy(&jpbres, &jpbreg); err != nil {
					return errors.New("set data gagal")
				}

				jpbres.Provinsi_Kode = &nop.PermohonanProvinsiID
				jpbres.Daerah_Kode = &nop.PermohonanKotaID
				jpbres.Kecamatan_Kode = &nop.PermohonanKecamatanID
				jpbres.Kelurahan_Kode = &nop.PermohonanKelurahanID
				jpbres.Blok_Kode = &nop.PermohonanBlokID
				jpbres.NoUrut = &nop.NoUrutPemohon
				jpbres.JenisOp = &nop.PemohonJenisOPID
				jpbres.Area_Kode = &areaCode

				if result := tx.Create(&jpbres); result.Error != nil {
					return errors.New("penyimpanan data gagal")
				}
			} else if opbng.Jpb_Kode == "15" {
				var jpbres *mopbng.Jpb15
				var jpbreg *rmopbng.RegJpb15
				if result := a.DB.
					Where("NoBangunan", opbng.NoBangunan).
					First(&jpbreg); result.Error != nil {
					return result.Error
				}
				if err := sc.Copy(&jpbres, &jpbreg); err != nil {
					return errors.New("set data gagal")
				}

				jpbres.Provinsi_Kode = &nop.PermohonanProvinsiID
				jpbres.Daerah_Kode = &nop.PermohonanKotaID
				jpbres.Kecamatan_Kode = &nop.PermohonanKecamatanID
				jpbres.Kelurahan_Kode = &nop.PermohonanKelurahanID
				jpbres.Blok_Kode = &nop.PermohonanBlokID
				jpbres.NoUrut = &nop.NoUrutPemohon
				jpbres.JenisOp = &nop.PemohonJenisOPID
				jpbres.Area_Kode = &areaCode

				if result := tx.Create(&jpbres); result.Error != nil {
					return errors.New("penyimpanan data gagal")
				}
			} else if opbng.Jpb_Kode == "16" {
				var jpbres *mopbng.Jpb16
				var jpbreg *rmopbng.RegJpb16
				if result := a.DB.
					Where("NoBangunan", opbng.NoBangunan).
					First(&jpbreg); result.Error != nil {
					return result.Error
				}
				if err := sc.Copy(&jpbres, &jpbreg); err != nil {
					return errors.New("set data gagal")
				}

				jpbres.Provinsi_Kode = &nop.PermohonanProvinsiID
				jpbres.Daerah_Kode = &nop.PermohonanKotaID
				jpbres.Kecamatan_Kode = &nop.PermohonanKecamatanID
				jpbres.Kelurahan_Kode = &nop.PermohonanKelurahanID
				jpbres.Blok_Kode = &nop.PermohonanBlokID
				jpbres.NoUrut = &nop.NoUrutPemohon
				jpbres.JenisOp = &nop.PemohonJenisOPID
				jpbres.Area_Kode = &areaCode
				jpbres.KelasBangunan16 = mopbng.KelasBangunan(kelasbangunan)

				if result := tx.Create(&jpbres); result.Error != nil {
					return errors.New("penyimpanan data gagal")
				}
			}
		}

		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data data: "+err.Error(), data)
	}

	return rp.OKSimple{Data: data}, nil
}
