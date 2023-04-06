package jaminanbongkar

import (
	"fmt"
	"math"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/apung/go-terbilang"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/jaminanbongkar"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jaminanbongkar/prosesjambong"
	mspt "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mdsrek "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifjambong"
	mtjrek "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifjambongrek"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	excelhelper "github.com/bapenda-kota-malang/apin-backend/pkg/excelhelper"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"github.com/leekchan/accounting"
	"github.com/xuri/excelize/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "jaminanbongkar"

func getTjRek(tx *gorm.DB, tipe string, luas *float64) (data *mtjrek.TarifJambongRek, err error) {
	baseQuery := tx.Model(&mtjrek.TarifJambongRek{}).Where("\"Tipe\" = ?", tipe)
	if luas != nil {
		baseQuery = baseQuery.Where("\"BatasBawah\" <= ?", *luas).Order("\"BatasBawah\" DESC")
	}
	if resStjrek := baseQuery.First(&data); resStjrek.Error != nil {
		err = resStjrek.Error
	} else if resStjrek.RowsAffected == 0 {
		err = fmt.Errorf("data tarif jambong rek not found")
	}
	return
}

func calcLuas(data mdsrek.DetailSptReklame) (luas float64) {
	switch data.JenisDimensi {
	case mdsrek.JenisDimensiPersegiPanjang:
		luas = *data.Panjang * *data.Lebar
	case mdsrek.JenisDimensiPersegiLingkaran:
		r := *data.Diameter / 2
		luas = math.Pi * math.Pow(r, 2)
	}
	luas *= float64(*data.Sisi) * float64(data.Jumlah)
	return
}

func processData(tx *gorm.DB, data *m.JaminanBongkar, dataDetails []m.DetailJambong, typeProcess string, userId uint) error {
	// get data spt with detail and tarif reklame
	var dataSpt mspt.Spt
	result := tx.Preload("DetailSptReklame.TarifReklame").First(&dataSpt, "\"Id\" = ?", data.Spt_Id.String())
	if result.RowsAffected == 0 {
		return fmt.Errorf("data spt not found")
	} else if result.Error != nil {
		return fmt.Errorf("gagal mengambil data spt: %w", result.Error)
	}

	if dataSpt.JenisMasaPajakReklame == nil {
		return fmt.Errorf("tidak ada data spt reklame jenis masa")
	}

	if dataSpt.DetailSptReklame == nil {
		return fmt.Errorf("tidak ada data spt reklame")
	}

	if typeProcess != "update" {
		data.CreateBy_User_Id = userId
		data.Nomor = fmt.Sprintf("JB-0000%s", dataSpt.NomorSpt[len(dataSpt.NomorSpt)-5:])
		// perAk, _ := dataSpt.PeriodeAkhir.Value()
		// // add 7 days
		// perAk = datatypes.Date(perAk.(time.Time).AddDate(0, 0, 7))
		// perAkDate := perAk.(datatypes.Date)
		// data.TanggalBatas = &perAkDate
	}
	data.Nominal = data.BiayaPemutusan

	var dataTjrek *mtjrek.TarifJambongRek
	switch *dataSpt.JenisMasaPajakReklame {
	case mtypes.MasaPajakBulan, mtypes.MasaPajakHari, mtypes.MasaPajakPenyelenggara:
		var err error
		dataTjrek, err = getTjRek(tx, "Insidentil", nil)
		if err != nil {
			return fmt.Errorf("gagal mengambil data tarif jambong rekening: %w", err)
		}
		data.TarifJambongRek_Id = &dataTjrek.Id
	}

	// create map for faster search
	dataDetailsMap := make(map[uint64]struct{})
	for _, v := range dataDetails {
		dataDetailsMap[v.DetailSptReklame_id] = struct{}{}
	}

	detailSptReklameMap := make(map[uint64]mdsrek.DetailSptReklame)
	for _, v := range dataSpt.DetailSptReklame {
		if _, ok := dataDetailsMap[v.Id]; !ok {
			dataDetails = append(dataDetails, m.DetailJambong{DetailSptReklame_id: v.Id})
		}
		detailSptReklameMap[v.Id] = v
	}

	for i := 0; i < len(dataDetails); i++ {
		if _, ok := detailSptReklameMap[dataDetails[i].DetailSptReklame_id]; !ok {
			return fmt.Errorf("invalid id spt reklame")
		}
		sptReklame := detailSptReklameMap[dataDetails[i].DetailSptReklame_id]

		// Insidentil
		switch *dataSpt.JenisMasaPajakReklame {
		case mtypes.MasaPajakBulan, mtypes.MasaPajakHari, mtypes.MasaPajakPenyelenggara:
			data.Nominal += (sptReklame.JumlahRp * (*dataTjrek.Nominal / float64(100)))
			continue
		}

		// get tarif jambong then calculate
		if dataDetails[i].TarifJambong_Id != nil {
			// prevent calc when data given below min id relation :)
			if *dataDetails[i].TarifJambong_Id < 1 {
				continue
			}
			var dataTj tarifjambong.TarifJambong
			res := tx.Model(&tarifjambong.TarifJambong{}).First(&dataTj, dataDetails[i].TarifJambong_Id)
			if res.Error != nil {
				return fmt.Errorf("gagal mengambil data tarif jambong: %w", res.Error)
			} else if res.RowsAffected == 0 {
				return fmt.Errorf("data tarif jambong kosong")
			}
			luas := calcLuas(sptReklame)
			data.Nominal += luas * *dataTj.Nominal
			continue
		}

		// search tarif by jenis reklame
		switch *sptReklame.TarifReklame.JenisReklame {
		// tarif jambong rek
		case "Billboard Disinari", "Billboard, Tembok/Tugu, Shop Panel, Letter, Neon Sign, Prismatek", "Neon Box":
			tipe := "RB"
			if *sptReklame.TarifReklame.JenisReklame == "Neon Box" {
				tipe = "RNB"
			}

			if *sptReklame.TarifReklame.DasarPengenaan != "Luas" {
				return fmt.Errorf("hanya bisa tarif reklame dasar pengenaan luas")
			}

			luas := calcLuas(sptReklame)

			var err error
			dataTjrek, err = getTjRek(tx, tipe, &luas)
			if err != nil {
				return fmt.Errorf("gagal mengambil data tarif jambong rekening: %w", err)
			}
			data.Nominal += luas * *dataTjrek.Nominal
		// tarif jambong
		case "Megatron, TV Media", "Rombong", "Bando Jalan, JPO, Taman Gantung":
			jenisRekStr := *sptReklame.TarifReklame.JenisReklame
			switch jenisRekStr {
			case "Megatron, TV Media":
				jenisRekStr = "Megatron"
			}

			if *sptReklame.TarifReklame.DasarPengenaan != "Luas" {
				return fmt.Errorf("hanya bisa tarif reklame dasar pengenaan luas")
			}

			luas := calcLuas(sptReklame)

			var dataTj tarifjambong.TarifJambong
			res := tx.Model(&tarifjambong.TarifJambong{}).Where("\"JenisReklame\" LIKE ?", fmt.Sprintf("%%%s%%", jenisRekStr)).First(&dataTj)
			if res.Error != nil {
				return fmt.Errorf("gagal mengambil data tarif jambong: %w", res.Error)
			} else if res.RowsAffected == 0 {
				return fmt.Errorf("data tarif jambong kosong")
			}
			tjId := int64(dataTj.Id)
			dataDetails[i].TarifJambong_Id = &tjId
			data.Nominal += luas * *dataTj.Nominal
		// insidentil, tarif jambong rek
		default:
			if dataTjrek == nil || (dataTjrek != nil && *dataTjrek.Tipe != "Insidentil") {
				var err error
				dataTjrek, err = getTjRek(tx, "Insidentil", nil)
				if err != nil {
					return fmt.Errorf("gagal mengambil data tarif jambong: %w", err)
				}
				data.TarifJambongRek_Id = &dataTjrek.Id
			}
			data.Nominal += (sptReklame.JumlahRp * (*dataTjrek.Nominal / float64(100)))
		}
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	err := tx.Transaction(func(tx2 *gorm.DB) error {
		// update
		if typeProcess == "update" {
			if result := tx2.Save(&data); result.Error != nil {
				return result.Error
			}
			if result := tx2.Save(&dataDetails); result.Error != nil {
				return result.Error
			}
			return nil
		}

		// create
		if result := tx2.Create(&data); result.Error != nil {
			return result.Error
		}
		for i := range dataDetails {
			dataDetails[i].JaminanBongkar_Id = data.Id
		}
		if result := tx2.Create(&dataDetails); result.Error != nil {
			return result.Error
		}
		return nil
	})
	data.DetailJambong = dataDetails
	return err
}

func transformUpdate(data *m.JaminanBongkar, detailsJambong []m.DetailJambongCreateDto) (dataDetails []m.DetailJambong, typeData string, err error) {
	if err = sc.CopyWithOption(&dataDetails, &data.DetailJambong, sc.Option{IgnoreEmpty: true}); err != nil {
		return
	}
	detailMap := make(map[uint64]int)
	for i, v := range dataDetails {
		detailMap[v.DetailSptReklame_id] = i
	}

	for _, v := range detailsJambong {
		if idx, ok := detailMap[v.DetailSptReklame_id]; ok {
			dataDetails[idx].TarifJambong_Id = v.TarifJambong_Id
			continue
		}
		dataDetails = append(dataDetails, m.DetailJambong{DetailSptReklame_id: v.DetailSptReklame_id, TarifJambong_Id: v.TarifJambong_Id})
	}
	data.DetailJambong = []m.DetailJambong{}
	typeData = "update"
	return
}

func Create(input m.CreateDto, userId uint, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.JaminanBongkar
	var dataDetails []m.DetailJambong

	typeData := ""
	if tx.Model(&m.JaminanBongkar{}).
		Preload("DetailJambong").
		Where("\"Spt_Id\" = ?", input.Spt_Id.String()).
		First(&data).
		RowsAffected == 0 {
		// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
		if err := sc.Copy(&data, &input); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
		}
		if err := sc.Copy(&dataDetails, &input.DetailsJambong); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
		}
	} else {
		if err := sc.CopyWithOption(&data, &input, sc.Option{IgnoreEmpty: true}); err != nil {
			sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
		}
		var err error
		dataDetails, typeData, err = transformUpdate(&data, input.DetailsJambong)
		if err != nil {
			return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("failed transform data: %s", err), data)
		}
	}

	if err := processData(tx, &data, dataDetails, typeData, userId); err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal menyimpan data: %s", err), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.JaminanBongkar
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.JaminanBongkar{}).
		Preload("Spt.DetailSptReklame").
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

func GetDetail(id int) (any, error) {
	var data *m.JaminanBongkar

	result := a.DB.Preload("Spt.DetailSptReklame").Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
		return tx.Omit("Password")
	}).First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, userId uint, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.JaminanBongkar
	result := tx.Preload("DetailJambong").First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	if err := sc.CopyWithOption(&data, &input, sc.Option{IgnoreEmpty: true}); err != nil {
		sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}
	var dataDetails []m.DetailJambong
	typeData := ""
	var err error
	dataDetails, typeData, err = transformUpdate(data, input.DetailsJambong)
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("failed transform data: %s", err), data)
	}

	if err := processData(tx, data, dataDetails, typeData, userId); err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal menyimpan data: %s", err), data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (any, error) {
	var data *m.JaminanBongkar
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	result = a.DB.Delete(&data, id)
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

func DownloadExcelList() (*excelize.File, error) {
	var (
		data            []m.JaminanBongkar
		getJenisReklame = func(jenisReklame mt.JenisMasa) string {
			switch jenisReklame {
			case mt.MasaPajakTahun:
				return "Masa Pajak Tetap 1 Tahun"
			case mt.MasaPajakBulan:
				return "Masa Pajak Insidentil 1 Bulan"
			case mt.MasaPajakHari:
				return "Masa Pajak Insidentil 1 Hari"
			default:
				return "Masa Pajak Insidentil 1 Kali Penyelenggaraan"
			}
		}
		getStatus = func(status prosesjambong.Status) string {
			if status == prosesjambong.StatusAmbil {
				return "Ambil"
			}
			return "Eksekusi"
		}
	)

	result := a.DB.
		Preload("CreateBy").
		Preload("Spt").
		Preload("ProsesJambong").
		Preload("Spt.ObjekPajak").
		Preload("Spt.Rekening").
		Find(&data)

	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
		return nil, err
	}

	var excelData = func() []interface{} {
		var tmp []interface{}

		tmp = append(tmp, map[string]interface{}{
			"A": "No",
			"B": "No. Jaminan Bongkar",
			"C": "No. SKPD",
			"D": "Nama WP",
			"E": "Tgl",
			"F": "Batas Pengambilan",
			"G": "Jenis Reklame",
			"H": "Nominal",
			"I": "Status",
			"J": "Nama User",
			"K": "Nama Rekening",
		})

		for i, v := range data {
			var (
				n                                                                 = i + 1
				objekPajakName, tanggalBatas, jenisMasaPajakReklame, rekeningName string
			)

			if v.Spt != nil {
				if v.Spt.ObjekPajak != nil {
					if v.Spt.ObjekPajak.Nama != nil {
						objekPajakName = *v.Spt.ObjekPajak.Nama
					}
				}

				if v.Spt.JenisMasaPajakReklame != nil {
					jenisMasaPajakReklame = getJenisReklame(*v.Spt.JenisMasaPajakReklame)
				}

				if v.Spt.Rekening != nil {
					if v.Spt.Rekening.Nama != nil {
						rekeningName = *v.Spt.Rekening.Nama
					}
				}
			}

			if v.TanggalBatas != nil {
				tanggalBatas = timehelper.GetDateFromUTCDatetime(v.TanggalBatas)
			}

			tmp = append(tmp, map[string]interface{}{
				"A": n,
				"B": v.Nomor,
				"C": v.Spt.NomorSpt,
				"D": objekPajakName,
				"E": v.CreatedAt.Format("2006-01-02 15:04:05"),
				"F": tanggalBatas,
				"G": jenisMasaPajakReklame,
				"H": v.Nominal,
				"I": getStatus(v.ProsesJambong.Status),
				"J": v.CreateBy.Name,
				"K": rekeningName,
			})
		}

		return tmp
	}()

	return excelhelper.ExportList(excelData, "List")
}

type ResponseFile struct {
	ID       int    `json:"id"`
	FileName string `json:"fileName"`
}

// jaminan bongkar reklame
func DownloadPDF(id int) (any, error) {
	var (
		data                                                                              m.JaminanBongkar
		lokasi                                                                            []string
		ketetapan, objekPajakName, objekPajakAlamat, nomorSpt, judulReklmae, jenisReklame string
		periodeAwal, periodeAkhir, tanggalBatas                                           datatypes.Date
		nominal, jaminan                                                                  float64
		now                                                                               time.Time = time.Now()
		ac                                                                                          = accounting.Accounting{Symbol: "", Precision: 0, Thousand: ".", Decimal: ","}
	)

	result := a.DB.
		Preload("CreateBy").
		Preload("Spt").
		Preload("ProsesJambong").
		Preload("Spt.ObjekPajak").
		Preload("Spt.Rekening").
		Preload("Spt.DetailSptReklame.TarifReklame").
		Preload("TarifJambong").
		First(&data, id)

	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
		return nil, err
	}

	for _, v := range data.Spt.DetailSptReklame {
		if v.Lokasi != nil {
			lokasi = append(lokasi, *v.Lokasi)
		}
	}

	if data.TarifJambong != nil {
		if data.TarifJambong.Nominal != nil {
			nominal = *data.TarifJambong.Nominal
		}
	}

	if data.Spt != nil {
		if data.Spt.ObjekPajak != nil {
			if data.Spt.ObjekPajak.Nama != nil {
				objekPajakName = *data.Spt.ObjekPajak.Nama
			}
			if data.Spt.ObjekPajak.Alamat != nil {
				objekPajakAlamat = *data.Spt.ObjekPajak.Alamat
			}
		}

		nomorSpt = data.Spt.NomorSpt
		periodeAwal = data.Spt.PeriodeAwal
		periodeAkhir = data.Spt.PeriodeAkhir

		if data.Spt.JudulReklame != nil {
			judulReklmae = *data.Spt.JudulReklame
		}

		for _, v := range data.Spt.DetailSptReklame {
			// if index != len(data.Spt.DetailSptReklame)-1 {
			var (
				panjang, lebar float64
				sisi           uint64
			)

			if v.Panjang != nil {
				panjang = *v.Panjang
			}
			if v.Lebar != nil {
				lebar = *v.Lebar
			}
			if v.Sisi != nil {
				sisi = *v.Sisi
			}

			if v.TarifReklame != nil {
				if v.TarifReklame.JenisReklame != nil {
					jenisReklame = *v.TarifReklame.JenisReklame
				}
			}

			ketetapan += fmt.Sprintf("%.2f x %v x %v x %v", (panjang * lebar), sisi, v.Jumlah, nominal)
			jaminan = panjang * lebar * float64(sisi) * float64(v.Jumlah) * nominal
			// }
		}
	}

	if data.TanggalBatas != nil {
		tanggalBatas = *data.TanggalBatas
	}

	finalResult := map[string]interface{}{
		"NomorJaminanBongkar": data.Nomor,
		"NamaWajibPajakSPT":   objekPajakName,
		"AlamatWajibPajakSPT": objekPajakAlamat,
		"NomorSKPD":           nomorSpt,
		"MasaPajak": fmt.Sprintf("%v - %v",
			timehelper.ConvertDatatypesDateToTime(&periodeAwal).Format("01/02/2006"),
			timehelper.ConvertDatatypesDateToTime(&periodeAkhir).Format("01/02/2006")),
		"BatasAkhirPengambilan": timehelper.ConvertDatatypesDateToTime(&tanggalBatas).Format("01/02/2006"),
		"LokasiReklame":         strings.Join(lokasi, ", "), // detailsptreklame.lokasi
		"TemaReklame":           judulReklmae,
		"JenisReklame":          jenisReklame,
		"Ketetapan":             ketetapan,
		"Jaminan":               ac.FormatMoney(jaminan),
		"BiayaListrik":          data.BiayaPemutusan,
		"Total":                 ac.FormatMoney(data.Nominal),
		"Terbilang":             terbilang.ToTerbilangRp(int(data.Nominal)),
		"TanggalSekarang":       fmt.Sprintf("%v %v %v", now.Day(), timehelper.GetMonth(&now), now.Year()),
	}

	uuid, err := sh.GetUuidv4()
	if err != nil {
		return nil, err
	}
	fileName := sh.GenerateFilename("jaminan-bongkar-reklame", uuid, 0, "pdf")

	outFile := filepath.Join(sh.GetPdfPath(), fileName)
	if err := GeneratePdf(outFile, finalResult); err != nil {
		return nil, err
	}
	_r := &ResponseFile{
		ID:       id,
		FileName: outFile,
	}
	return rp.OKSimple{Data: _r}, nil
}
