package potensiopwp

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"

	rpth "github.com/bapenda-kota-malang/apin-backend/pkg/reporthelper"
	strh "github.com/bapenda-kota-malang/apin-backend/pkg/stringhelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/base64helper"
	"github.com/bapenda-kota-malang/apin-backend/pkg/excelhelper"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/google/uuid"
	sc "github.com/jinzhu/copier"
	"github.com/lib/pq"
	"github.com/xuri/excelize/v2"

	mpegawai "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "potensiop"

func filePreProcess(b64String, docsname string, userId uint, oldId uuid.UUID) (fileName, path, extFile string, id uuid.UUID, err error) {
	extFile, err = base64helper.GetExtensionBase64(b64String)
	if err != nil {
		return
	}
	path = sh.GetResourcesPath()
	switch extFile {
	case "pdf":
		path = sh.GetPdfPath()
	case "png", "jpeg":
		path = sh.GetImgPath()
	case "xlsx", "xls":
		path = sh.GetExcelPath()
	default:
		err = errors.New("file tidak diketahui")
		return
	}
	if oldId == uuid.Nil {
		id, err = sh.GetUuidv4()
		if err != nil {
			err = errors.New("gagal generate uuid")
			return
		}
	} else {
		id = oldId
	}
	fileName = sh.GenerateFilename(docsname, id, userId, extFile)
	return
}

func Create(input m.CreatePotensiOpDto, userId uint, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.PotensiOp
	if input.Id == uuid.Nil {
		id, err := sh.GetUuidv4()
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		input.Id = id
	} else {
		if err := tx.First(&data, "\"Id\" = ?", input.Id.String()).Error; err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
	}

	if input.FotoKtp != nil {
		var errChan = make(chan error)
		fileName, path, extFile, _, err := filePreProcess(*input.FotoKtp, "FotoKtpPotensiOp", userId, input.Id)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		go sh.SaveFile(*input.FotoKtp, fileName, path, extFile, errChan)
		if err := <-errChan; err != nil {
			return sh.SetError("request", "create-data", source, "failed", "failed save foto", data)
		}
		input.FotoKtp = &fileName
	}

	if input.FormBapl != nil {
		var errChan = make(chan error)
		fileName, path, extFile, _, err := filePreProcess(*input.FormBapl, "FormBaplPotensiOp", userId, input.Id)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		go sh.SaveFile(*input.FormBapl, fileName, path, extFile, errChan)
		if err := <-errChan; err != nil {
			return sh.SetError("request", "create-data", source, "failed", "failed save pdf", data)
		}
		input.FormBapl = &fileName
	}

	if input.DokumenLainnya != nil {
		tmp := pq.StringArray{}
		for i, v := range *input.DokumenLainnya {
			var errChan = make(chan error)
			fileName, path, extFile, _, err := filePreProcess(v, "DokumenLainnya"+(strconv.Itoa(i+1))+"PotensiOp", userId, input.Id)
			if err != nil {
				return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
			}
			go sh.SaveFile(v, fileName, path, extFile, errChan)
			if err := <-errChan; err != nil {
				return sh.SetError("request", "create-data", source, "failed", "failed save file", data)
			}
			tmp = append(tmp, fileName)
		}
		input.DokumenLainnya = &tmp
	}

	if input.FotoObjek != nil {
		tmp := pq.StringArray{}
		for i, v := range *input.FotoObjek {
			var errChan = make(chan error)
			fileName, path, extFile, _, err := filePreProcess(v, "FotoObjek"+(strconv.Itoa(i+1))+"PotensiOp", userId, input.Id)
			if err != nil {
				return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
			}
			go sh.SaveFile(v, fileName, path, extFile, errChan)
			if err := <-errChan; err != nil {
				return sh.SetError("request", "create-data", source, "failed", "failed save pdf", data)
			}
			tmp = append(tmp, fileName)
		}
		input.FotoObjek = &tmp
	}

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.CopyWithOption(&data, &input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// static add value to field
	data.User_Id = uint64(userId)
	data.Status = nt.StatusAktif

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.PotensiOp
	var count int64

	var pagination gh.Pagination
	result := a.DB.Model(&m.PotensiOp{}).
		Joins("Rekening").
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Preload("Rekening").
		Preload("PotensiPemilikWp").
		Preload("DetailPotensiOp.Kecamatan").
		Preload("DetailPotensiOp.Kelurahan").
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

func GetDetail(id uuid.UUID) (any, error) {
	var data *m.PotensiOp

	result := a.DB.
		Preload("DetailPotensiOp.Kecamatan").
		Preload("DetailPotensiOp.Kelurahan").
		Preload("PotensiPemilikWp.Daerah").
		Preload("PotensiPemilikWp.Kelurahan").
		Preload("PotensiNarahubung.Daerah").
		Preload("PotensiNarahubung.Kelurahan").
		Preload("Bapl.Koordinator").
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).First(&data, "\"Id\" = ?", id.String())
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	dataBapl := data.Bapl
	for i := 0; i < len(dataBapl); i++ {
		if dataBapl[i].Petugas_Pegawai_Id == nil {
			continue
		}
		var dataPetugas *[]mpegawai.Pegawai
		petugasArr := []int64(*dataBapl[i].Petugas_Pegawai_Id)
		if res := a.DB.Find(&dataPetugas, petugasArr); res.Error != nil {
			return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
		} else if res.RowsAffected == 0 {
			return nil, nil
		}
		dataBapl[i].Petugas = dataPetugas
	}
	data.Bapl = dataBapl

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id uuid.UUID, input m.UpdatePotensiOpDto, userId uint, tx *gorm.DB) (any, error) {
	var data *m.PotensiOp
	result := a.DB.First(&data, "\"Id\" = ?", id.String())
	if result.RowsAffected == 0 {
		return nil, nil
	}

	if input.FotoKtp != nil {
		var errChan = make(chan error)
		fileName, path, extFile, _, err := filePreProcess(*input.FotoKtp, "FotoKtpPotensiOp", userId, id)
		if err != nil {
			return sh.SetError("request", "update-data", source, "failed", err.Error(), data)
		}
		if data.FotoKtp != nil {
			go sh.ReplaceFile(*data.FotoKtp, *input.FotoKtp, fileName, path, extFile, errChan)
		} else {
			go sh.SaveFile(*input.FotoKtp, fileName, path, extFile, errChan)
		}
		if err := <-errChan; err != nil {
			return sh.SetError("request", "update-data", source, "failed", "failed save foto", data)
		}
		input.FotoKtp = &fileName
	}

	if input.FormBapl != nil {
		var errChan = make(chan error)
		fileName, path, extFile, _, err := filePreProcess(*input.FormBapl, "FormBaplPotensiOp", userId, id)
		if err != nil {
			return sh.SetError("request", "update-data", source, "failed", err.Error(), data)
		}
		if data.FormBapl != nil {
			go sh.ReplaceFile(*data.FormBapl, *input.FormBapl, fileName, path, extFile, errChan)
		} else {
			go sh.SaveFile(*input.FormBapl, fileName, path, extFile, errChan)
		}
		if err := <-errChan; err != nil {
			return sh.SetError("request", "update-data", source, "failed", "failed save pdf", data)
		}
		input.FormBapl = &fileName
	}

	if input.FotoObjek != nil {
		tmp := pq.StringArray{}
		lenData := 0
		if data.FotoObjek != nil {
			tmp = *data.FotoObjek
			re := regexp.MustCompile(`^FotoObjek(\d*)`)
			if match := re.FindStringSubmatch(tmp[len(tmp)-1]); len(match) > 0 {
				lenData, _ = strconv.Atoi(match[1])
			}
		}
		for i, v := range *input.FotoObjek {
			var errChan = make(chan error)
			fileName, path, extFile, _, err := filePreProcess(v, "FotoObjek"+(strconv.Itoa(i+lenData+1))+"PotensiOp", userId, id)
			if err != nil {
				return sh.SetError("request", "update-data", source, "failed", err.Error(), data)
			}
			go sh.SaveFile(v, fileName, path, extFile, errChan)
			if err := <-errChan; err != nil {
				return sh.SetError("request", "update-data", source, "failed", "failed save pdf", data)
			}
			tmp = append(tmp, fileName)
		}
		input.FotoObjek = &tmp
	}

	if input.FotoObjekDeleted != nil && data.FotoObjek != nil {
		deleteMap := make(map[string]struct{})
		for _, v := range *input.FotoObjekDeleted {
			deleteMap[v] = struct{}{}
		}
		newDataArray := pq.StringArray{}
		for _, v := range *input.FotoObjek {
			if _, exist := deleteMap[v]; exist {
				path := sh.GetPathByFilename(v)
				if err := sh.RemoveFile(path, v); err != nil {
					return sh.SetError("request", "update-data", source, "failed", "failed remove file", v)
				}
			} else {
				newDataArray = append(newDataArray, v)
			}
		}
		input.FotoObjek = &newDataArray
		data.FotoObjek = nil
	}

	if input.DokumenLainnya != nil {
		tmp := pq.StringArray{}
		lenData := 0
		if data.DokumenLainnya != nil {
			tmp = *data.DokumenLainnya
			re := regexp.MustCompile(`^DokumenLainnya(\d*)`)
			if match := re.FindStringSubmatch(tmp[len(tmp)-1]); len(match) > 0 {
				lenData, _ = strconv.Atoi(match[1])
			}
		}
		for i, v := range *input.DokumenLainnya {
			var errChan = make(chan error)
			fileName, path, extFile, _, err := filePreProcess(v, "DokumenLainnya"+(strconv.Itoa(i+lenData+1))+"PotensiOp", userId, id)
			if err != nil {
				return sh.SetError("request", "update-data", source, "failed", err.Error(), data)
			}
			go sh.SaveFile(v, fileName, path, extFile, errChan)
			if err := <-errChan; err != nil {
				return sh.SetError("request", "update-data", source, "failed", "failed save file", data)
			}
			tmp = append(tmp, fileName)
		}
		input.DokumenLainnya = &tmp
	}

	if input.DokumenLainnyaDeleted != nil && data.DokumenLainnya != nil {
		deleteMap := make(map[string]struct{})
		for _, v := range *input.DokumenLainnyaDeleted {
			deleteMap[v] = struct{}{}
		}
		newDataArray := pq.StringArray{}
		for _, v := range *input.DokumenLainnya {
			if _, exist := deleteMap[v]; exist {
				path := sh.GetPathByFilename(v)
				if err := sh.RemoveFile(path, v); err != nil {
					return sh.SetError("request", "update-data", source, "failed", "failed remove file", v)
				}
			} else {
				newDataArray = append(newDataArray, v)
			}
		}
		input.DokumenLainnya = &newDataArray
		data.DokumenLainnya = nil
	}

	if err := sc.CopyWithOption(&data, &input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id uuid.UUID) (any, error) {
	var data *m.PotensiOp
	result := a.DB.First(&data, "\"Id\" = ?", id.String())
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	result = a.DB.Where("\"Id\" = ?", id.String()).Delete(&data)
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

func DownloadExcelList(input m.FilterDto) (*excelize.File, error) {
	var data []m.PotensiOp

	result := a.DB.Model(&m.PotensiOp{}).
		Scopes(gh.Filter(input)).
		Preload("Rekening").
		Preload("PotensiPemilikWp").
		Preload("DetailPotensiOp.Kecamatan").
		Preload("DetailPotensiOp.Kelurahan").
		Find(&data)
	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
		return nil, err
	}

	var excelData = func() []interface{} {
		var tmp []interface{}
		tmp = append(tmp, map[string]interface{}{
			"A": "No",
			"B": "Golongan",
			"C": "Nomor",
			"D": "Jenis Usaha",
			"E": "Nama WP",
			"F": "Nama Pemilik",
			"G": "Kecamatan",
			"H": "Kelurahan",
			"I": "Tanggal",
			"J": "Status",
		})
		for i, v := range data {
			n := i + 1
			tmp = append(tmp, map[string]interface{}{
				"A": n,
				"B": func() string {
					if v.Golongan == 1 {
						return "Orang Pribadi"
					} else if v.Golongan == 2 {
						return "Badan"
					} else {
						return ""
					}
				}(),
				"C": v.Id,
				"D": func() string {
					if v.Rekening != nil {
						if v.Rekening.JenisUsaha != nil {
							return *v.Rekening.JenisUsaha
						}
						return ""
					}
					return ""
				}(),
				"E": v.DetailPotensiOp.Nama,
				"F": func() string {
					if len(v.PotensiPemilikWp) != 0 && v.PotensiPemilikWp != nil {
						for _, r := range v.PotensiPemilikWp {
							return r.Nama
						}
					}
					return ""
				}(),
				"G": func() string {
					if v.DetailPotensiOp.Kecamatan != nil {
						return v.DetailPotensiOp.Kecamatan.Nama
					}
					return ""
				}(),
				"H": func() string {
					if v.DetailPotensiOp.Kelurahan != nil {
						return v.DetailPotensiOp.Kelurahan.Nama
					}
					return ""
				}(),
				"I": func() string {
					return v.CreatedAt.Format("2006-01-02")
				}(),
				"J": func() string {
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

type ResponseFile struct {
	ID       uuid.UUID `json:"id"`
	FileName string    `json:"fileName"`
}

func DownloadPdf(id uuid.UUID) (interface{}, error) {

	dataDetail, err := GetDetail(id)
	if err != nil {
		return sh.SetError("request", "download-pdf", source, "failed", "gagal mengambil data ", dataDetail)
	}
	finalresultTmp := dataDetail.(rp.OKSimple).Data.(*m.PotensiOp)

	var KpList = []*rpth.TCItemList{}
	var TrList = []*rpth.TCItemList{}

	DetailPotensiAirTanah := &m.DetailPotensiAirTanah{}
	DetailPotensiHiburans := []m.DetailPotensiHiburan{}
	DetailPotensiHotels := []m.DetailPotensiHotel{}
	DetailPotensiParkirs := []m.DetailPotensiParkir{}
	DetailPotensiPPJNonPLN := &m.DetailPotensiPPJNonPLN{}
	DetailPotensiReklames := []m.DetailPotensiReklame{}
	DetailPotensiResto := &m.DetailPotensiResto{}

	if finalresultTmp.DetailPotensiAirTanah != nil {
		DetailPotensiAirTanah = finalresultTmp.DetailPotensiAirTanah
	}

	DetailPotensiHiburans = append(DetailPotensiHiburans, m.DetailPotensiHiburan{})
	if len(finalresultTmp.DetailPotensiHiburans) != 0 {
		DetailPotensiHiburans = finalresultTmp.DetailPotensiHiburans
		for _, hiburan := range DetailPotensiHiburans {
			KpList = append(KpList, &rpth.TCItemList{
				Title: fmt.Sprintf("%d Orang", hiburan.Kapasitas),
				Value: fmt.Sprintf("Tarif Rp %f Jumlah %d", hiburan.Tarif, hiburan.Jumlah),
				Unit:  "buah",
			})
		}
	}

	DetailPotensiHotels = append(DetailPotensiHotels, m.DetailPotensiHotel{})
	if len(finalresultTmp.DetailPotensiHotels) != 0 {
		DetailPotensiHotels = finalresultTmp.DetailPotensiHotels
		for _, hotel := range DetailPotensiHotels {
			KpList = append(KpList, &rpth.TCItemList{
				Title: fmt.Sprintf("%d Orang", hotel.Kapasitas),
				Value: fmt.Sprintf("Tarif Rp %f Jumlah %d", hotel.TarifFasilitas, hotel.JumlahFasilitas),
				Unit:  "Kamar",
			})
		}
	}
	DetailPotensiParkirs = append(DetailPotensiParkirs, m.DetailPotensiParkir{})
	if len(finalresultTmp.DetailPotensiParkirs) != 0 {
		DetailPotensiParkirs = finalresultTmp.DetailPotensiParkirs
		for _, parkir := range DetailPotensiParkirs {
			var jenis = "Motor"
			if parkir.JenisKendaraan == 2 {
				jenis = "Mobil"
			}
			TrList = append(TrList, &rpth.TCItemList{
				Title: jenis,
				Value: strh.FloatToAny(&parkir.Tarif, "-"),
				Unit:  "Rp. ",
			})
		}
	}

	if finalresultTmp.DetailPotensiPPJNonPLN != nil {
		DetailPotensiPPJNonPLN = finalresultTmp.DetailPotensiPPJNonPLN
	}

	DetailPotensiReklames = append(DetailPotensiReklames, m.DetailPotensiReklame{})
	if len(finalresultTmp.DetailPotensiReklames) != 0 {
		DetailPotensiReklames = finalresultTmp.DetailPotensiReklames
	}

	if finalresultTmp.DetailPotensiResto != nil {
		DetailPotensiResto = finalresultTmp.DetailPotensiResto

		KpList = append(KpList, &rpth.TCItemList{
			Title: "Meja",
			Value: strconv.Itoa(int(DetailPotensiResto.JumlahMeja)),
			Unit:  "buah",
		})

		KpList = append(KpList, &rpth.TCItemList{
			Title: "Kursi",
			Value: strconv.Itoa(int(DetailPotensiResto.JumlahKursi)),
			Unit:  "buah",
		})

		TrList = append(TrList, &rpth.TCItemList{
			Title: "Makanan",
			Value: strh.NullToAny(&DetailPotensiResto.RentangHargaMakanan, "-"),
			Unit:  "Rp. ",
		})

		TrList = append(TrList, &rpth.TCItemList{
			Title: "Minuman",
			Value: strh.NullToAny(&DetailPotensiResto.RentangHargaMinuman, "-"),
			Unit:  "Rp. ",
		})
	}

	tanggalTinjau, hariTinjau, waktuTinjau, _ := th.GetAllFormatTime(finalresultTmp.StartDate)
	finalresult := map[string]interface{}{
		"Data": map[string]string{
			"NamaWP":      strh.NullToAny(&finalresultTmp.PotensiPemilikWp[0].Nama, "-"),
			"AlamatWP":    strh.NullToAny(&finalresultTmp.PotensiPemilikWp[0].Alamat, "-"),
			"NamaUsaha":   strh.NullToAny(&finalresultTmp.DetailPotensiOp.Nama, "-"),
			"AlamatUsaha": strh.NullToAny(&finalresultTmp.DetailPotensiOp.Alamat, "-"),
			"Phone":       strh.NullToAny(finalresultTmp.DetailPotensiOp.Telp, "-"),

			"MulaiBuka":              strh.NullToAny(finalresultTmp.JamBuka, "-"),
			"MulaiTutup":             strh.NullToAny(finalresultTmp.JamTutup, "-"),
			"NamaWPLap":              strh.NullToAny(&finalresultTmp.DetailPotensiOp.Nama, "-"),
			"LuasBangunan":           strh.NullToAny(finalresultTmp.LuasBangunan, "-"),
			"RataRataPengunjung":     strh.FloatToAny(&finalresultTmp.OmsetOp, "-"),
			"PenerapanPajakKonsumen": strh.FloatToAny(finalresultTmp.PajakKonsumenPct, "-"),
			"KebutuhanAir":           strh.BoolToAny(&finalresultTmp.AirTanah, "Air Tanah", "PDAM"),
			"DiameterPipa":           strh.NullToAny(&DetailPotensiAirTanah.DiameterPipa, "-"),
			"JumlahTitik":            strh.NullToAny(&DetailPotensiAirTanah.JumlahSumber, "-"),
			"FasilitasPenerangan":    strh.BoolToAny(&finalresultTmp.Genset, "Genset", "PLN"),
			"JmlGenset":              strh.NullToAny(&DetailPotensiPPJNonPLN.Jumlah, "-"),
			"JenisUsaha":             strh.NullToAny(finalresultTmp.Rekening.JenisUsaha, "-"),
			"KapasitasDaya":          strh.NullToAny(&DetailPotensiPPJNonPLN.KapasitasDaya, "-"),
			"KapasitasMax":           "-",
		},
		"Custom": map[string]string{
			"Hari":    hariTinjau,
			"Tanggal": tanggalTinjau,
			"Pukul":   waktuTinjau,
		},
		"ListItem": map[string][]*rpth.TCItemList{
			"Kapasitas": KpList,
			"Tarif":     TrList,
		},
	}

	_, err = sh.GetUuidv4()
	if err != nil {
		return nil, err
	}
	fileName := sh.GenerateFilename(fmt.Sprintf("potensiopwp_%s", strh.NullToAny(finalresultTmp.Rekening.JenisUsaha, "")), id, 0, "pdf")

	outFile := filepath.Join(sh.GetPdfPath(), fileName)
	if err := GeneratePdf(outFile, finalresult, strh.NullToAny(finalresultTmp.Rekening.Objek, "01")); err != nil {
		return nil, err
	}
	_r := &ResponseFile{
		ID:       id,
		FileName: outFile,
	}
	return rp.OKSimple{Data: _r}, nil
}
