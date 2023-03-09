// service
package regpermohonan

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

	ori "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	mropfas "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	mropbng "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbangunan"
	mroptnh "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbumi"
	mroppbb "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakpbb"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/regpelayanan"
	sksk "github.com/bapenda-kota-malang/apin-backend/internal/models/regsksk"
	mrwppbb "github.com/bapenda-kota-malang/apin-backend/internal/models/regwajibpajakpbb"
	oppbb "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakpbb"
)

// /// Private funcs start here
const source = "permohonan"

///// Exported funcs start here

func Create(input m.PermohonanRequestDto) (any, error) {
	var (
		err                   error
		permohonanDetail      *m.RegPstDetail
		permohonanBaru        *m.RegPstDataOPBaru
		permohonanDetailInput *m.RegPstDetailInput
		permohonanPengurangan *m.RegPstPermohonanPengurangan
		nop                   *m.PermohonanNOP
		pembetulanSpptSKPSTP  *m.RegPembetulanSpptSKPSTP
		pembatalanSppt        *m.RegPembatalanSppt
		keputusanKeberatanPbb *m.RegKeputusanKeberatanPbb
		SPMKP                 *m.RegSPMKP
		SkSk                  *sksk.RegSkSk
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
		permohonanDetail = new(m.RegPstDetail)
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
	var data []m.RegPstPermohonan
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.RegPstPermohonan{}).
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
		checkdataori  *ori.PstPermohonan
		checkdata     *m.RegPstPermohonan
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
		tempNoUrutOri, _ = strconv.Atoi(*checkdataori.NoUrutPelayanan)
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
	result, err := oppbb.GetRegDetailbyNop(*permohonannop)

	return result, err
}

func GetDetail(id int) (interface{}, error) {
	var (
		data                  *m.RegPstPermohonan
		permohonanBaru        *m.RegPstDataOPBaru
		permohonanDetail      *m.RegPstDetail
		permohonanPengurangan *m.RegPstPermohonanPengurangan
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

func GetDetailApproval(id int) (interface{}, error) {
	var (
		data                  *m.RegPstPermohonan
		permohonanBaru        *m.RegPstDataOPBaru
		permohonanDetail      *m.RegPstDetail
		permohonanPengurangan *m.RegPstPermohonanPengurangan
		lampiran              *m.RegPstLampiran
		oppbb                 *mroppbb.RegObjekPajakPbb
		wppbb                 *mrwppbb.RegWajibPajakPbb
		optnh                 *mroptnh.RegObjekPajakBumi
		opbng                 []mropbng.RegObjekPajakBangunan
		opfas                 []mropfas.RegFasilitasBangunan

		tempmroppbb mroppbb.CreateDto

		bangunans []mropbng.CreateDto
		// jpb       any
	)
	// tempNOP := transformNOP(nop)
	// dec_nop := m.DecodeNOPPermohonan(&tempNOP)
	result := a.DB.First(&data, &id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", data)
	}
	if result := a.DB.Where("PermohonanId", data.Id).First(&permohonanDetail); result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data NOP Detail", permohonanDetail)
	}
	_ = a.DB.Where("PermohonanId", data.Id).First(&lampiran)

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
		Where("PstPermohonan_id", data.Id).
		First(&oppbb); result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", oppbb)
	}

	if result := a.DB.
		Where("Id", oppbb.RegWajibPajakPbb_Id).
		First(&wppbb); result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", wppbb)
	}

	if result := a.DB.
		Where("PstPermohonan_id", data.Id).
		First(&optnh); result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", optnh)
	}

	if err := sc.CopyWithOption(&tempmroppbb, &oppbb, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mencopy data oppbb", oppbb)
	}

	if err := sc.CopyWithOption(&tempmroppbb.RegWajibPajakPbbs, &wppbb, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mencopy data wppbb", wppbb)
	}

	if err := sc.CopyWithOption(&tempmroppbb.RegObjekPajakBumis, &optnh, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mencopy data optnh", optnh)
	}

	finalresult := data.SetPstPermohonanResponse()
	finalresult.PstDataOPBaru = permohonanBaru
	finalresult.PstDetail = permohonanDetail
	finalresult.PstPermohonanPengurangan = permohonanPengurangan
	finalresult.PstLampiran = lampiran
	finalresult.PstOpjekPajakPBB = &tempmroppbb

	if resBang := a.DB.
		Where("PstPermohonan_id", data.Id).
		Find(&opbng); resBang.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", opbng)
	} else if resBang.RowsAffected > 0 {
		for _, bang := range opbng {
			var tempBang mropbng.CreateDto
			var fasilitas mropbng.OPBngFasilitasBangunan

			// if err := sc.CopyWithOption(tempBang, bang, sc.Option{IgnoreEmpty: true}); err != nil {
			// 	return sh.SetError("request", "get-data", source, "failed", "gagal mencopy data data", bang)
			// }

			tempBang = bang.RegObjekPajakBangunanTransform()

			if result := a.DB.
				Where("PstPermohonan_id", data.Id).
				Where("NoBangunan", bang.NoBangunan).
				Find(&opfas); result.Error != nil {
				return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data fasilitas", opfas)
			} else {
				for _, fas := range opfas {
					switch *fas.KodeFasilitas {
					case "01":
						fasilitas.FBJumlahACSplit = fas.JumlahSatuan
					case "02":
						fasilitas.FBJumlahACWindow = fas.JumlahSatuan
					case "11":
						fasilitas.FBIsACCentral = fas.JumlahSatuan
					case "12":
						tempint := 12
						fasilitas.FBTipeLapisanKolam = &tempint
						fasilitas.FBLuasKolamRenang = fas.JumlahSatuan
					case "13":
						tempint := 13
						fasilitas.FBTipeLapisanKolam = &tempint
						fasilitas.FBLuasKolamRenang = fas.JumlahSatuan
					case "16":
						fasilitas.FBHalamanBerat = fas.JumlahSatuan
					case "15":
						fasilitas.FBHalamanSendang = fas.JumlahSatuan
					case "14":
						fasilitas.FBHalamanRingan = fas.JumlahSatuan
					case "17":
						fasilitas.FBHalamanLantai = fas.JumlahSatuan
					case "18":
						fasilitas.FBTenisLampuBeton = fas.JumlahSatuan
					case "21":
						fasilitas.FBTenisTanpaLampuBeton = fas.JumlahSatuan
					case "19":
						fasilitas.FBTenisAspal1 = fas.JumlahSatuan
					case "22":
						fasilitas.FBTenisAspal2 = fas.JumlahSatuan
					case "20":
						fasilitas.FBTenisLiatRumput1 = fas.JumlahSatuan
					case "23":
						fasilitas.FBTenisLiatRumput2 = fas.JumlahSatuan
					case "30":
						fasilitas.FBLiftPenumpang = fas.JumlahSatuan
					case "31":
						fasilitas.FBLiftKapsul = fas.JumlahSatuan
					case "32":
						fasilitas.FBLiftBarang = fas.JumlahSatuan
					case "33":
						fasilitas.FBTangga80 = fas.JumlahSatuan
					case "34":
						fasilitas.FBTangga81 = fas.JumlahSatuan
					case "35":
						tempint := 35
						fasilitas.FBPagarBahan = &tempint
						fasilitas.FBPagarPanjang = fas.JumlahSatuan
					case "36":
						tempint := 36
						fasilitas.FBPagarBahan = &tempint
						fasilitas.FBPagarPanjang = fas.JumlahSatuan
					case "37":
						fasilitas.FBPKHydrant = fas.JumlahSatuan
					case "39":
						fasilitas.FBPKSplinkler = fas.JumlahSatuan
					case "38":
						fasilitas.FBPKFireAI = fas.JumlahSatuan
					case "41":
						fasilitas.FBPABX = fas.JumlahSatuan
					case "42":
						fasilitas.FBSumur = fas.JumlahSatuan
					case "07":
						fasilitas.JpbKlinikACCentralKamar = fas.JumlahSatuan
					case "08":
						fasilitas.JpbKlinikACCentralRuang = fas.JumlahSatuan
					case "04":
						fasilitas.JpbHotelACCentralKamar = fas.JumlahSatuan
					case "05":
						fasilitas.JpbHotelACCentralRuang = fas.JumlahSatuan
					case "09":
						fasilitas.JpbApartemenACCentralKamar = fas.JumlahSatuan
					case "10":
						fasilitas.JpbApartemenACCentralLain = fas.JumlahSatuan
					default:
						// return sh.SetError("request", "get-data", source, "failed", "kode fasilitas tidak ditemukan", fas.KodeFasilitas)
					}
				}
			}

			if bang.Jpb_Kode == "00" || bang.Jpb_Kode == "01" || bang.Jpb_Kode == "10" || bang.Jpb_Kode == "11" {
			} else if bang.Jpb_Kode == "02" {
				var jpbres *mropbng.RegJpb2
				if result := a.DB.
					Where("NoBangunan", bang.NoBangunan).
					First(&jpbres); result.Error != nil {
					return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
				}

			} else if bang.Jpb_Kode == "03" {
				var jpbres *mropbng.RegJpb3
				if result := a.DB.
					Where("NoBangunan", bang.NoBangunan).
					First(&jpbres); result.Error != nil {
					return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
				}
				fasilitas.JpbProdTinggi = jpbres.TinggiKolom3
				fasilitas.JpbProdLebar = jpbres.LebarBentang3
				fasilitas.JpbProdDaya = jpbres.DayaDukungLantai
				fasilitas.JpbProdKeliling = jpbres.KelilingDinding3
				fasilitas.JpbProdLuas = jpbres.LuasMezzanine3

			} else if bang.Jpb_Kode == "04" {
				var jpbres *mropbng.RegJpb4
				if result := a.DB.
					Where("NoBangunan", bang.NoBangunan).
					First(&jpbres); result.Error != nil {
					return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
				}

			} else if bang.Jpb_Kode == "05" {
				var jpbres *mropbng.RegJpb5
				if result := a.DB.
					Where("NoBangunan", bang.NoBangunan).
					First(&jpbres); result.Error != nil {
					return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
				}
				fasilitas.JpbKlinikACCentralKamar = jpbres.LuasKamarAcCentral
				fasilitas.JpbKlinikACCentralRuang = jpbres.LuasRuangLainAcCentral

			} else if bang.Jpb_Kode == "06" {
				var jpbres *mropbng.RegJpb6
				if result := a.DB.
					Where("NoBangunan", bang.NoBangunan).
					First(&jpbres); result.Error != nil {
					return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
				}

			} else if bang.Jpb_Kode == "07" {
				var jpbres *mropbng.RegJpb7
				if result := a.DB.
					Where("NoBangunan", bang.NoBangunan).
					First(&jpbres); result.Error != nil {
					return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
				}
				tempStr, _ := strconv.Atoi((string)(jpbres.JenisHotel))
				fasilitas.JpbHotelJenis = &tempStr
				tempStr, _ = strconv.Atoi((string)(jpbres.JumlahBintang))
				fasilitas.JpbHotelBintang = &tempStr
				fasilitas.JpbHotelJmlKamar = jpbres.JumlahKamar
				fasilitas.JpbHotelACCentralKamar = jpbres.LuasKamarAcCentral
				fasilitas.JpbHotelACCentralRuang = jpbres.LuasRuangLainAcCentral
			} else if bang.Jpb_Kode == "08" {
				var jpbres *mropbng.RegJpb8
				if result := a.DB.
					Where("NoBangunan", bang.NoBangunan).
					First(&jpbres); result.Error != nil {
					return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
				}
				fasilitas.JpbProdTinggi = jpbres.TinggiKolom8
				fasilitas.JpbProdLebar = jpbres.LebarBentang8
				fasilitas.JpbProdDaya = jpbres.DayaDukungLantai
				fasilitas.JpbProdKeliling = jpbres.KelilingDinding8
				fasilitas.JpbProdLuas = jpbres.LuasMezzanine8
			} else if bang.Jpb_Kode == "09" {
				var jpbres *mropbng.RegJpb9
				if result := a.DB.
					Where("NoBangunan", bang.NoBangunan).
					First(&jpbres); result.Error != nil {
					return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
				}
			} else if bang.Jpb_Kode == "12" {
				var jpbres *mropbng.RegJpb12
				if result := a.DB.
					Where("NoBangunan", bang.NoBangunan).
					First(&jpbres); result.Error != nil {
					return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
				}
			} else if bang.Jpb_Kode == "13" {
				var jpbres *mropbng.RegJpb13
				if result := a.DB.
					Where("NoBangunan", bang.NoBangunan).
					First(&jpbres); result.Error != nil {
					return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
				}
				fasilitas.JpbApartemenJumlah = jpbres.JumlahApartment
				fasilitas.JpbApartemenACCentralKamar = jpbres.LuasApartAcCentral
				fasilitas.JpbApartemenACCentralLain = jpbres.LuasRuangLainAcCentral

			} else if bang.Jpb_Kode == "14" {
				var jpbres *mropbng.RegJpb14
				if result := a.DB.
					Where("NoBangunan", bang.NoBangunan).
					First(&jpbres); result.Error != nil {
					return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
				}
			} else if bang.Jpb_Kode == "15" {
				var jpbres *mropbng.RegJpb15
				if result := a.DB.
					Where("NoBangunan", bang.NoBangunan).
					First(&jpbres); result.Error != nil {
					return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
				}
				fasilitas.JpbTankiKapasitas = jpbres.KapasitasTanki
				tempStr := (string)(jpbres.LetakTanki)
				fasilitas.JpbTankiLetak = &tempStr
			} else if bang.Jpb_Kode == "16" {
				var jpbres *mropbng.RegJpb16
				if result := a.DB.
					Where("NoBangunan", bang.NoBangunan).
					First(&jpbres); result.Error != nil {
					return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data pengurangan", permohonanPengurangan)
				}
			}

			tempBang.RegFasBangunan = &fasilitas
			bangunans = append(bangunans, tempBang)
		}
	}

	finalresult.PstOpjekPajakPBB.RegObjekPajakBumis.RegObjekPajakBangunans = &bangunans

	return rp.OKSimple{Data: finalresult}, nil
}

func Update(id int, input m.PermohonanRequestDto) (interface{}, error) {
	var (
		data                  *m.RegPstPermohonan
		permohonanBaru        *m.RegPstDataOPBaru
		permohonanDetail      *m.RegPstDetail
		permohonanDetailInput *m.RegPstDetailInput
		permohonanPengurangan *m.RegPstPermohonanPengurangan
		nop                   *m.PermohonanNOP
		pembetulanSpptSKPSTP  *m.RegPembetulanSpptSKPSTP
		pembatalanSppt        *m.RegPembatalanSppt
		keputusanKeberatanPbb *m.RegKeputusanKeberatanPbb
		SPMKP                 *m.RegSPMKP
		SkSk                  *sksk.RegSkSk
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
		data             *m.RegPstPermohonan
		permohonanDetail *m.RegPstDetail
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
		data                  *m.RegPstPermohonan
		permohonanBaru        *m.RegPstDataOPBaru
		permohonanDetail      *m.RegPstDetail
		permohonanPengurangan *m.RegPstPermohonanPengurangan
		pembetulanSpptSKPSTP  *m.RegPembetulanSpptSKPSTP
		pembatalanSppt        *m.RegPembatalanSppt
		keputusanKeberatanPbb *m.RegKeputusanKeberatanPbb
		SPMKP                 *m.RegSPMKP
		SkSk                  *sksk.RegSkSk
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

func Approval(kd string, auth int, input m.RequestApprovalPermohonan, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.RegPstPermohonan
	result := tx.First(&data, "\"Id\" = ?", input.Id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	data.Status = input.Status

	if kd == "00" && auth == 4 {
		//baru
	} else if kd == "01" && auth == 4 {
		//diterima
	} else if kd == "02" && auth == 4 {
		//ditolak
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

func UpdateApproval(id int, auth int, input m.PermohonanApprovalRequestDto) (interface{}, error) {
	var (
		err error

		data m.RegPstPermohonan
		nop  *m.PermohonanNOP

		regPermohonanDetail      *m.RegPstDetail
		regPermohonanBaru        *m.RegPstDataOPBaru
		regPermohonanPengurangan *m.RegPstPermohonanPengurangan
		regPembetulanSpptSKPSTP  *m.RegPembetulanSpptSKPSTP
		regKeputusanKeberatanPbb *m.RegKeputusanKeberatanPbb

		regPermohonanDetailInput      *m.RegPstDetail
		regPermohonanBaruInput        *m.RegPstDataOPBaru
		regPermohonanPenguranganInput *m.RegPstPermohonanPengurangan
		regPembetulanSpptSKPSTPInput  *m.RegPembetulanSpptSKPSTP
		regKeputusanKeberatanPbbInput *m.RegKeputusanKeberatanPbb

		roppbb *mroppbb.RegObjekPajakPbb
		rwppbb *mrwppbb.RegWajibPajakPbb
		roptnh *mroptnh.RegObjekPajakBumi
		// ropbng []mropbng.RegObjekPajakBangunan
		ropfas []mropfas.RegFasilitasBangunan

		opbng mropbng.RegObjekPajakBangunan
		// opfas []mropfas.RegFasilitasBangunan

		// areaCode    string
		tempKodeFas string
		errTemp     error
	)

	if input.NOP != nil {
		nop = m.DecodeNOPPermohonan(input.NOP)
		// areaCode = nop.PermohonanProvinsiID + nop.PermohonanKotaID + nop.PermohonanKecamatanID + nop.PermohonanKelurahanID
	}

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data: "+err.Error(), data)
	}
	data.TahunPelayanan = input.TahunPajak

	idPermohonan := data.Id

	err = a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data permohonan ke db satu if karena result dipakai sekali, +error
		if result := tx.Save(&data); result.Error != nil {
			return errors.New("penyimpanan data permohonan gagal")
		}

		regPermohonanBaruInput = input.PstDataOPBaru
		regPermohonanDetailInput = input.PstDetail
		regPermohonanPenguranganInput = input.PstPermohonanPengurangan
		// nop = m.DecodeNOPPermohonan(input.NOP)

		_ = a.DB.Where("PermohonanId", idPermohonan).First(&regPermohonanDetail)
		if err := sc.CopyWithOption(&regPermohonanDetail, &regPermohonanDetailInput, sc.Option{IgnoreEmpty: true}); err != nil {
			return errors.New("set data permohonan detail gagal")
		}
		if err := sc.CopyWithOption(&regPermohonanDetail, &nop, sc.Option{IgnoreEmpty: true}); err != nil {
			return errors.New("set data permohonan detail gagal")
		}
		regPermohonanDetail.PermohonanId = &data.Id

		if result := tx.Save(&regPermohonanDetail); result.Error != nil {
			return errors.New("penyimpanan data permohonan detail gagal")
		}

		if *data.BundelPelayanan == m.JenisPelayanan[0] {
			_ = a.DB.Where("PermohonanId", idPermohonan).First(&regPermohonanBaru)
			if err := sc.CopyWithOption(&regPermohonanBaru, &regPermohonanBaruInput, sc.Option{IgnoreEmpty: true}); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			if err := sc.CopyWithOption(&regPermohonanBaru, &nop, sc.Option{IgnoreEmpty: true}); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			regPermohonanBaru.PermohonanId = &data.Id

			if result := tx.Save(&regPermohonanBaru); result.Error != nil {
				return errors.New("penyimpanan data permohonan baru gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[2] {
			regPembetulanSpptSKPSTPInput = data.SetPembetulanSpptSKPSTP(*nop)
			_ = a.DB.Where("PermohonanId", idPermohonan).First(&regPembetulanSpptSKPSTP)
			if err := sc.CopyWithOption(&regPembetulanSpptSKPSTP, &regPembetulanSpptSKPSTPInput, sc.Option{IgnoreEmpty: true}); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			if err := sc.CopyWithOption(&regPembetulanSpptSKPSTP, &nop, sc.Option{IgnoreEmpty: true}); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			regPembetulanSpptSKPSTP.PermohonanId = &data.Id

			if result := tx.Save(&regPembetulanSpptSKPSTP); result.Error != nil {
				return errors.New("penyimpanan data permohonan pembetulan SPPT/SKP/STP gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[5] || *data.BundelPelayanan == m.JenisPelayanan[6] {
			regKeputusanKeberatanPbbInput = data.SetKeputusanKeberatanPbb(*nop)
			_ = a.DB.Where("PermohonanId", idPermohonan).First(&regKeputusanKeberatanPbb)
			if err := sc.CopyWithOption(&regKeputusanKeberatanPbb, &regKeputusanKeberatanPbbInput, sc.Option{IgnoreEmpty: true}); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			if err := sc.CopyWithOption(&regKeputusanKeberatanPbb, &nop, sc.Option{IgnoreEmpty: true}); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			regKeputusanKeberatanPbb.PermohonanId = &data.Id

			if result := tx.Save(&regKeputusanKeberatanPbb); result.Error != nil {
				return errors.New("penyimpanan data permohonan keberatan gagal")
			}
		} else if *data.BundelPelayanan == m.JenisPelayanan[7] || *data.BundelPelayanan == m.JenisPelayanan[9] {
			_ = a.DB.Where("PermohonanId", idPermohonan).First(&regPermohonanPengurangan)
			if err := sc.CopyWithOption(&regPermohonanPengurangan, &regPermohonanPenguranganInput, sc.Option{IgnoreEmpty: true}); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			if err := sc.CopyWithOption(&regPermohonanPengurangan, &nop, sc.Option{IgnoreEmpty: true}); err != nil {
				return errors.New("set data permohonan detail gagal")
			}
			regPermohonanPengurangan.PermohonanId = &data.Id
			regPermohonanPengurangan.JenisPengurangan = input.JenisPengurangan
			tempPP, _ := strconv.ParseFloat(*input.PersentasePengurangan, 64)
			regPermohonanPengurangan.PersentasePengurangan = &tempPP

			if result := tx.Save(&regPermohonanPengurangan); result.Error != nil {
				return errors.New("penyimpanan data permohonan pengurangan gagal")
			}
		}

		// roppbb                 *roppbb.RegObjekPajakPbb
		_ = tx.Where("PstPermohonan_id", idPermohonan).First(&roppbb)
		if err := sc.CopyWithOption(&roppbb, &input.PstOpjekPajakPBB, sc.Option{IgnoreEmpty: true}); err != nil {
			return errors.New("set data permohonan detail gagal")
		}
		if result := tx.Save(&roppbb); result.Error != nil {
			return errors.New("penyimpanan data permohonan pengurangan gagal")
		}

		// rwppbb                 *rwppbb.RegWajibPajakPbb
		_ = tx.Where("Id", roppbb.RegWajibPajakPbb_Id).First(&rwppbb)
		if err := sc.CopyWithOption(&rwppbb, &input.PstOpjekPajakPBB.RegWajibPajakPbbs, sc.Option{IgnoreEmpty: true}); err != nil {
			return errors.New("set data permohonan detail gagal")
		}
		if result := tx.Save(&rwppbb); result.Error != nil {
			return errors.New("penyimpanan data permohonan pengurangan gagal")
		}

		// roptnh                 *roptnh.RegObjekPajakBumi
		_ = tx.Where("PstPermohonan_id", idPermohonan).First(&roptnh)
		if err := sc.CopyWithOption(&roptnh, &input.PstOpjekPajakPBB.RegObjekPajakBumis, sc.Option{IgnoreEmpty: true}); err != nil {
			return errors.New("set data permohonan detail gagal")
		}
		if result := tx.Save(&roptnh); result.Error != nil {
			return errors.New("penyimpanan data permohonan pengurangan gagal")
		}

		// ropbng                 *ropbng.RegObjekPajakBangunan
		tempBangs := input.PstOpjekPajakPBB.RegObjekPajakBumis.RegObjekPajakBangunans
		if tempBangs != nil {
			notNil := *tempBangs
			for _, val := range notNil {
				_ = a.DB.Where("PstPermohonan_id", idPermohonan).Where("NoBangunan", val.NoBangunan).First(&opbng)
				if err := sc.CopyWithOption(&opbng, &val, sc.Option{IgnoreEmpty: true}); err != nil {
					return errors.New("set data permohonan detail gagal")
				}
				if result := tx.Save(&opbng); result.Error != nil {
					return errors.New("penyimpanan data permohonan pengurangan gagal")
				}

				// ropfas                 *ropfas.RegFasilitasBangunan
				resFas := tx.
					Where("PstPermohonan_id", idPermohonan).
					Where("NoBangunan", val.NoBangunan).
					Find(&ropfas)
				if resFas.RowsAffected > 0 {
					for _, dataFasBangunan := range ropfas {
						dataFasBangunan.Provinsi_Kode = &nop.PermohonanProvinsiID
						dataFasBangunan.Daerah_Kode = &nop.PermohonanKotaID
						dataFasBangunan.Kecamatan_Kode = &nop.PermohonanKecamatanID
						dataFasBangunan.Kelurahan_Kode = &nop.PermohonanKelurahanID
						dataFasBangunan.NoUrut = &nop.NoUrutPemohon
						dataFasBangunan.JenisOp = &nop.PemohonJenisOPID
						dataFasBangunan.NoBangunan = val.NoBangunan
						dataFasBangunan.PstPermohonan_id = &idPermohonan

						//create data fasilitas/jpb berdasar fasilitas/jpb di isi
						if val.RegFasBangunan.FBJumlahACSplit != nil {
							tempKodeFas = "01"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBJumlahACSplit
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBJumlahACWindow != nil {
							tempKodeFas = "02"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBJumlahACWindow
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBIsACCentral != nil {
							tempKodeFas = "11"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBIsACCentral
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBTipeLapisanKolam != nil {
							tempKodeFas = strconv.Itoa(*val.RegFasBangunan.FBTipeLapisanKolam)
							if val.RegFasBangunan.FBLuasKolamRenang != nil {
								dataFasBangunan.KodeFasilitas = &tempKodeFas
								dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBLuasKolamRenang
							}
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBHalamanBerat != nil {
							tempKodeFas = "16"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBHalamanBerat
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBHalamanSendang != nil {
							tempKodeFas = "15"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBHalamanSendang
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBHalamanRingan != nil {
							tempKodeFas = "14"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBJumlahACSplit
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBHalamanLantai != nil {
							tempKodeFas = "17"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBHalamanLantai
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBTenisLampuBeton != nil {
							tempKodeFas = "18"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBTenisLampuBeton
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBTenisTanpaLampuBeton != nil {
							tempKodeFas = "21"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBTenisTanpaLampuBeton
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBTenisAspal1 != nil {
							tempKodeFas = "19"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBTenisAspal1
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBTenisAspal2 != nil {
							tempKodeFas = "22"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBTenisAspal2
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBTenisLiatRumput1 != nil {
							tempKodeFas = "20"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBTenisLiatRumput1
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBTenisLiatRumput2 != nil {
							tempKodeFas = "23"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBTenisLiatRumput2
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBLiftPenumpang != nil {
							tempKodeFas = "30"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBLiftPenumpang
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBLiftKapsul != nil {
							tempKodeFas = "31"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBLiftKapsul
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBLiftBarang != nil {
							tempKodeFas = "32"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBLiftBarang
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBTangga80 != nil {
							tempKodeFas = "33"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBTangga80
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBTangga81 != nil {
							tempKodeFas = "34"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBTangga81
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBPagarBahan != nil {
							tempKodeFas = strconv.Itoa(*val.RegFasBangunan.FBPagarBahan)
							if val.RegFasBangunan.FBPagarPanjang != nil {
								dataFasBangunan.KodeFasilitas = &tempKodeFas
								dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBPagarPanjang
							}
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBPKHydrant != nil {
							tempKodeFas = "37"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBPKHydrant
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBPKSplinkler != nil {
							tempKodeFas = "39"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBPKSplinkler
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBPKFireAI != nil {
							tempKodeFas = "38"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBPKFireAI
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBPABX != nil {
							tempKodeFas = "41"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBPABX
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
						if val.RegFasBangunan.FBSumur != nil {
							tempKodeFas = "42"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.RegFasBangunan.FBSumur
							errTemp = tx.Save(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1
						}
					}
				}

				if val.Jpb_Kode == "00" || val.Jpb_Kode == "01" || val.Jpb_Kode == "10" || val.Jpb_Kode == "11" {
				} else if val.Jpb_Kode == "02" {
					var jpbres *mropbng.RegJpb2
					if result := tx.
						Where("PstPermohonan_id", idPermohonan).
						Where("NoBangunan", val.NoBangunan).
						First(&jpbres); result.Error != nil {
						return result.Error
					}

					temp := mropbng.KelasBangunan(*val.RegFasBangunan.KelasBangunan)
					jpbres.KelasBangunan2 = &temp

					if result := tx.Save(&jpbres); result.Error != nil {
						return errors.New("penyimpanan data gagal")
					}
				} else if val.Jpb_Kode == "03" {
					var jpbres *mropbng.RegJpb3
					if result := tx.
						Where("PstPermohonan_id", idPermohonan).
						Where("NoBangunan", val.NoBangunan).
						First(&jpbres); result.Error != nil {
						return result.Error
					}

					jpbres.TinggiKolom3 = val.RegFasBangunan.JpbProdTinggi
					jpbres.LebarBentang3 = val.RegFasBangunan.JpbProdLebar
					jpbres.DayaDukungLantai = val.RegFasBangunan.JpbProdDaya
					jpbres.KelilingDinding3 = val.RegFasBangunan.JpbProdKeliling
					jpbres.LuasMezzanine3 = val.RegFasBangunan.JpbProdLuas

					if result := tx.Save(&jpbres); result.Error != nil {
						return errors.New("penyimpanan data gagal")
					}
				} else if val.Jpb_Kode == "04" {
					var jpbres *mropbng.RegJpb4
					if result := tx.
						Where("PstPermohonan_id", idPermohonan).
						Where("NoBangunan", val.NoBangunan).
						First(&jpbres); result.Error != nil {
						return result.Error
					}

					temp := mropbng.KelasBangunan(*val.RegFasBangunan.KelasBangunan)
					jpbres.KelasBangunan4 = &temp

					if result := tx.Save(&jpbres); result.Error != nil {
						return errors.New("penyimpanan data gagal")
					}
				} else if val.Jpb_Kode == "05" {
					var jpbres *mropbng.RegJpb5
					if result := tx.
						Where("PstPermohonan_id", idPermohonan).
						Where("NoBangunan", val.NoBangunan).
						First(&jpbres); result.Error != nil {
						return result.Error
					}

					jpbres.LuasKamarAcCentral = val.RegFasBangunan.JpbKlinikACCentralKamar
					jpbres.LuasRuangLainAcCentral = val.RegFasBangunan.JpbKlinikACCentralRuang

					temp := mropbng.KelasBangunan(*val.RegFasBangunan.KelasBangunan)
					jpbres.KelasBangunan5 = &temp

					if result := tx.Save(&jpbres); result.Error != nil {
						return errors.New("penyimpanan data gagal")
					}
				} else if val.Jpb_Kode == "06" {
					var jpbres *mropbng.RegJpb6
					if result := tx.
						Where("PstPermohonan_id", idPermohonan).
						Where("NoBangunan", val.NoBangunan).
						First(&jpbres); result.Error != nil {
						return result.Error
					}

					temp := mropbng.KelasBangunan(*val.RegFasBangunan.KelasBangunan)
					jpbres.KelasBangunan6 = &temp

					if result := tx.Save(&jpbres); result.Error != nil {
						return errors.New("penyimpanan data gagal")
					}
				} else if val.Jpb_Kode == "07" {
					var jpbres *mropbng.RegJpb7
					if result := tx.
						Where("PstPermohonan_id", idPermohonan).
						Where("NoBangunan", val.NoBangunan).
						First(&jpbres); result.Error != nil {
						return result.Error
					}

					jpbres.JenisHotel = (mropbng.JenisHotel)(strconv.Itoa(*val.RegFasBangunan.JpbHotelJenis))
					jpbres.JumlahBintang = (mropbng.JumlahBintang)(strconv.Itoa(*val.RegFasBangunan.JpbHotelBintang))
					jpbres.JumlahKamar = val.RegFasBangunan.JpbHotelJmlKamar
					jpbres.LuasKamarAcCentral = val.RegFasBangunan.JpbHotelACCentralKamar
					jpbres.LuasRuangLainAcCentral = val.RegFasBangunan.JpbHotelACCentralRuang

					if result := tx.Save(&jpbres); result.Error != nil {
						return errors.New("penyimpanan data gagal")
					}
				} else if val.Jpb_Kode == "08" {
					var jpbres *mropbng.RegJpb8
					if result := tx.
						Where("PstPermohonan_id", idPermohonan).
						Where("NoBangunan", val.NoBangunan).
						First(&jpbres); result.Error != nil {
						return result.Error
					}

					jpbres.TinggiKolom8 = val.RegFasBangunan.JpbProdTinggi
					jpbres.LebarBentang8 = val.RegFasBangunan.JpbProdLebar
					jpbres.DayaDukungLantai = val.RegFasBangunan.JpbProdDaya
					jpbres.KelilingDinding8 = val.RegFasBangunan.JpbProdKeliling
					jpbres.LuasMezzanine8 = val.RegFasBangunan.JpbProdLuas

					if result := tx.Save(&jpbres); result.Error != nil {
						return errors.New("penyimpanan data gagal")
					}
				} else if val.Jpb_Kode == "09" {
					var jpbres *mropbng.RegJpb9
					if result := tx.
						Where("PstPermohonan_id", idPermohonan).
						Where("NoBangunan", val.NoBangunan).
						First(&jpbres); result.Error != nil {
						return result.Error
					}

					temp := mropbng.KelasBangunan(*val.RegFasBangunan.KelasBangunan)
					jpbres.KelasBangunan9 = &temp

					if result := tx.Save(&jpbres); result.Error != nil {
						return errors.New("penyimpanan data gagal")
					}
				} else if val.Jpb_Kode == "12" {
					var jpbres *mropbng.RegJpb12
					if result := tx.
						Where("PstPermohonan_id", idPermohonan).
						Where("NoBangunan", val.NoBangunan).
						First(&jpbres); result.Error != nil {
						return result.Error
					}

					if result := tx.Save(&jpbres); result.Error != nil {
						return errors.New("penyimpanan data gagal")
					}
				} else if val.Jpb_Kode == "13" {
					var jpbres *mropbng.RegJpb13
					if result := tx.
						Where("PstPermohonan_id", idPermohonan).
						Where("NoBangunan", val.NoBangunan).
						First(&jpbres); result.Error != nil {
						return result.Error
					}

					jpbres.JumlahApartment = val.RegFasBangunan.JpbApartemenJumlah
					jpbres.LuasApartAcCentral = val.RegFasBangunan.JpbApartemenACCentralKamar
					jpbres.LuasRuangLainAcCentral = val.RegFasBangunan.JpbApartemenACCentralLain

					temp := mropbng.KelasBangunan(*val.RegFasBangunan.KelasBangunan)
					jpbres.KelasBangunan13 = &temp

					if result := tx.Save(&jpbres); result.Error != nil {
						return errors.New("penyimpanan data gagal")
					}
				} else if val.Jpb_Kode == "14" {
					var jpbres *mropbng.RegJpb14
					if result := tx.
						Where("PstPermohonan_id", idPermohonan).
						Where("NoBangunan", val.NoBangunan).
						First(&jpbres); result.Error != nil {
						return result.Error
					}

					if result := tx.Save(&jpbres); result.Error != nil {
						return errors.New("penyimpanan data gagal")
					}
				} else if val.Jpb_Kode == "15" {
					var jpbres *mropbng.RegJpb15
					if result := tx.
						Where("PstPermohonan_id", idPermohonan).
						Where("NoBangunan", val.NoBangunan).
						First(&jpbres); result.Error != nil {
						return result.Error
					}

					jpbres.KapasitasTanki = val.RegFasBangunan.JpbTankiKapasitas
					jpbres.LetakTanki = (mropbng.LetakTanki)(*val.RegFasBangunan.JpbTankiLetak)

					if result := tx.Save(&jpbres); result.Error != nil {
						return errors.New("penyimpanan data gagal")
					}
				} else if val.Jpb_Kode == "16" {
					var jpbres *mropbng.RegJpb16
					if result := tx.
						Where("PstPermohonan_id", idPermohonan).
						Where("NoBangunan", val.NoBangunan).
						First(&jpbres); result.Error != nil {
						return result.Error
					}
					// jpbres.KelasBangunan16 = mropbng.KelasBangunan(*val.RegFasBangunan.KelasBangunan)

					if result := tx.Save(&jpbres); result.Error != nil {
						return errors.New("penyimpanan data gagal")
					}
				}
			}
		}

		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data data: "+err.Error(), data)
	}

	tempApproval := m.RequestApprovalPermohonan{
		NoPelayanan: data.NoPelayanan,
		NOP:         input.NOP,
		NIP:         input.NIP,
		Status:      input.Status,
		User_ID:     input.User_ID,
	}

	resultApproval, errApproval := Approval(*input.Status, auth, tempApproval, a.DB)
	if errApproval != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data data: "+errApproval.Error(), data)
	}

	// return rp.OKSimple{Data: data}, nil
	return rp.OK{
		// Meta: t.IS{
		// 	"affected": strconv.Itoa(rowsAffected),
		// },
		Data: t.II{
			"permohonan":     data,
			"resultApproval": resultApproval,
		},
	}, nil
}
