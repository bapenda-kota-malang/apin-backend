package sppt

import (
	"fmt"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	sc "github.com/jinzhu/copier"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	saop "github.com/bapenda-kota-malang/apin-backend/internal/services/anggotaobjekpajak"
	sopb "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakbumi"
	sopp "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakpbb"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"

	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	opbg "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	opb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	pn "github.com/bapenda-kota-malang/apin-backend/internal/models/penetapan"
)

const source = "sppt"

func Create(input m.RequestDto) (any, error) {
	var data m.Sppt

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := a.DB.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.Sppt
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Sppt{}).
		Where("\"TanggalTerbit_sppt\" <= ?", datatypes.Date(time.Now())).
		Order("\"TanggalTerbit_sppt\" ").
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

func GetListNop(user_id int) (any, error) {
	var data *[]m.Sppt

	result := a.DB.Where("User_ID", user_id).Find(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func GetDetail(id int) (any, error) {
	var data *m.Sppt

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func GetByNop(provinsiKode, daerahKode, kecamatanKode, kelurahanKode, blokKode, noUrut, jenisOp string) (any, error) {
	var data m.Sppt
	result := a.DB.Where(&m.Sppt{
		Propinsi_Id:   &provinsiKode,
		Dati2_Id:      &daerahKode,
		Kecamatan_Id:  &kecamatanKode,
		Keluarahan_Id: &kelurahanKode,
		Blok_Id:       &blokKode,
		NoUrut:        &noUrut,
		JenisOP_Id:    &jenisOp,
	}).First(&data)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "get-data-by-nop", source, "failed", "data tidak ada", data)
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-by-nop", source, "failed", "gagal mendapatkan data: "+result.Error.Error(), data)
	}
	return rp.OKSimple{Data: data}, nil
}

func Update(id int, input m.RequestDto) (any, error) {
	var data *m.Sppt
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	if err := sc.Copy(&data, &input); err != nil {
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

func UpdateByNop(input m.RequestDto) (any, error) {
	var data *m.Sppt
	result := a.DB.
		Where("Propinsi_Id", &input.Propinsi_Id).
		Where("Dati2_Id", &input.Dati2_Id).
		Where("Kecamatan_Id", &input.Kecamatan_Id).
		Where("Keluarahan_Id", &input.Keluarahan_Id).
		Where("Blok_Id", &input.Blok_Id).
		Where("NoUrut", &input.NoUrut).
		Where("JenisOP_Id", &input.JenisOP_Id).
		Where("TahunPajakskp_sppt", &input.TahunPajakskp_sppt).
		First(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	if err := sc.Copy(&data, &input); err != nil {
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

func Delete(id int) (any, error) {
	var data *m.Sppt
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

func Penilaian(input m.PenilaianDto) (any, error) {
	var data []m.Sppt
	var respOpp interface{}
	var respAop interface{}
	var respOpb interface{}
	var resp t.II

	condition, _ := nopCondition(input)
	result := a.DB.Where(condition).Find(&data)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "penilaian-data", source, "failed", "gagal mengambil data", data)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {

		resultOpp, err := sopp.PenilaianSppt(data, tx)
		if err != nil {
			return err
		}
		respOpp = resultOpp

		resultAop, err := saop.PenilaianSppt(data, tx)
		if err != nil {
			return err
		}
		respAop = resultAop

		resultOpb, err := sopb.PenilaianSppt(data, tx)
		if err != nil {
			return err
		}
		respOpb = resultOpb

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	resp = t.II{
		"objekPajakPbb":     respOpp.(rp.OKSimple).Data,
		"anggotaObjekPajak": respAop.(rp.OKSimple).Data,
		"objekPajakBumi":    respOpb.(rp.OKSimple).Data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func CetakDaftarTagihan(req m.GetDaftarTagihan) (any, error) {
	var list []m.Sppt
	xlsx := excelize.NewFile()

	transactionReport := "Daftar Tagihan OP"
	xlsx.SetSheetName(xlsx.GetSheetName(1), transactionReport)

	xlsx.SetCellValue(transactionReport, "A1", "Daftar PBB Belum Dibayar")
	xlsx.SetCellValue(transactionReport, "A2", "Wilayah Kotamadya Malang Kecamatan "+*req.NamaKecamatan)
	xlsx.SetCellValue(transactionReport, "A3", "Tahun "+*req.TahunPajakAwal+" s/d "+*req.TahunPajakAkhir+" ketetapan PBB "+*req.KetetapanPBBAwal+" s/d "+*req.KetetapanPBBAkhir)
	xlsx.MergeCell(transactionReport, "A1", "F1")
	xlsx.MergeCell(transactionReport, "A2", "F2")
	xlsx.MergeCell(transactionReport, "A3", "F3")

	result := a.DB.Where("", req.Provinsi_Kode).
		Where("Propinsi_Id", req.Provinsi_Kode).
		Where("Dati2_Id", req.Daerah_Kode).
		Where("Kecamatan_Id", req.Kecamatan_Kode).
		Where("\"TahunPajakskp_sppt\" >= '" + *req.TahunPajakAwal + "'").
		Where("\"TahunPajakskp_sppt\" <= '" + *req.TahunPajakAkhir + "'").
		Where("\"PBBygHarusDibayar_sppt\" >= " + *req.KetetapanPBBAwal).
		Where("\"PBBygHarusDibayar_sppt\" <= " + *req.KetetapanPBBAkhir).
		Where("\"StatusPembayaran_sppt\" == '0'").
		Find(&list).
		Order("\"Propinsi_Id\", \"Dati2_Id\", \"Kecamatan_Id\", \"Keluarahan_Id\"")

	if result.RowsAffected == 0 {
		return sh.SetError("request", "create-data", source, "failed", "gagal menemukan data: "+result.Error.Error(), list)
	}

	tempKel := ""
	isSet := false
	totalPerKel := 0
	skipRow := 4
	tempPBB := 0

	for i, data := range list {
		if &tempKel != data.Keluarahan_Id {
			xlsx.SetCellValue(transactionReport, fmt.Sprintf("A%d", i+skipRow), req.Provinsi_Kode+"."+req.Daerah_Kode+"."+req.Kecamatan_Kode+"."+*data.Keluarahan_Id+" "+*data.Keluarahan_Id)
			xlsx.MergeCell(transactionReport, fmt.Sprintf("A%d", i+skipRow), fmt.Sprintf("F%d", i+skipRow))
			skipRow = skipRow + 1
			xlsx.SetCellValue(transactionReport, fmt.Sprintf("A%d", i+skipRow), "No")
			xlsx.SetCellValue(transactionReport, fmt.Sprintf("B%d", i+skipRow), "NOP")
			xlsx.SetCellValue(transactionReport, fmt.Sprintf("C%d", i+skipRow), "Tahun")
			xlsx.SetCellValue(transactionReport, fmt.Sprintf("D%d", i+skipRow), "Nama WP")
			xlsx.SetCellValue(transactionReport, fmt.Sprintf("E%d", i+skipRow), "Alamat WP/Alamat OP")
			xlsx.SetCellValue(transactionReport, fmt.Sprintf("F%d", i+skipRow), "PBB yang harus dibayar (Rp.)")
			skipRow = skipRow + 1
			tempKel = *data.Keluarahan_Id
			totalPerKel = 0
			isSet = true
		}

		// insert data from transaction stores
		xlsx.SetCellValue(transactionReport, fmt.Sprintf("A%d", i+skipRow), i)
		xlsx.SetCellValue(transactionReport, fmt.Sprintf("B%d", i+skipRow), *data.NoUrut+"."+*data.JenisOP_Id)
		xlsx.SetCellValue(transactionReport, fmt.Sprintf("C%d", i+skipRow), data.TahunPajakskp_sppt)
		xlsx.SetCellValue(transactionReport, fmt.Sprintf("D%d", i+skipRow), data.NamaWP_sppt)
		xlsx.SetCellValue(transactionReport, fmt.Sprintf("E%d", i+skipRow), data.JalanWPskp_sppt)
		tempPBB = *data.PBBygHarusDibayar_sppt - *data.Faktorpengurangan_sppt
		xlsx.SetCellValue(transactionReport, fmt.Sprintf("F%d", i+skipRow), tempPBB)
		totalPerKel = totalPerKel + tempPBB
		if isSet {
			skipRow = skipRow + 1
			xlsx.SetCellValue(transactionReport, fmt.Sprintf("A%d", i+skipRow), "")
			xlsx.SetCellValue(transactionReport, fmt.Sprintf("B%d", i+skipRow), "")
			xlsx.SetCellValue(transactionReport, fmt.Sprintf("C%d", i+skipRow), "")
			xlsx.SetCellValue(transactionReport, fmt.Sprintf("D%d", i+skipRow), "")
			xlsx.SetCellValue(transactionReport, fmt.Sprintf("E%d", i+skipRow), "")
			xlsx.SetCellValue(transactionReport, fmt.Sprintf("F%d", i+skipRow), totalPerKel)
		}
	}

	return rp.OKSimple{
		Data: xlsx,
	}, nil
}

func PenilaianMassal(input m.PenilaianDto) (any, error) {
	var data []m.Sppt

	filter := m.Sppt{
		Propinsi_Id:        input.Provinsi_Kode,
		Dati2_Id:           input.Daerah_Kode,
		Kecamatan_Id:       input.Kecamatan_Kode,
		Keluarahan_Id:      input.Kelurahan_Kode,
		TahunPajakskp_sppt: &input.Tahun,
	}

	result := a.DB.Where(&filter).Order("\"Id\" asc").Find(&data)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "get-data-by-nop", source, "failed", "data tidak ada", data)
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-by-nop", source, "failed", "gagal mendapatkan data: "+result.Error.Error(), data)
	}

	var res []map[string]interface{}
	for _, v := range data {
		var (
			opBumiData opb.ObjekPajakBumi
			opBngData  opbg.ObjekPajakBangunan
		)

		// get ObjekPajakBumi detail
		fBumi := opb.ObjekPajakBumi{
			NopDetail: nop.NopDetail{
				Provinsi_Kode:  v.Propinsi_Id,
				Daerah_Kode:    v.Dati2_Id,
				Kecamatan_Kode: v.Kecamatan_Id,
				Kelurahan_Kode: v.Keluarahan_Id,
				Blok_Kode:      v.Blok_Id,
				NoUrut:         v.NoUrut,
			},
		}
		a.DB.Where(&fBumi).Order("\"Id\" desc").First(&opBumiData)

		type RsltBumi struct {
			PenilaianBumi *float64 `gorm:"column:penilaian_bumi"`
		}
		var rsltBumi RsltBumi
		a.DB.Raw("SELECT * FROM Penilaian_Bumi(?,?,?,?,?,?,?,?,?,?,?)", v.Propinsi_Id, v.Dati2_Id, v.Kecamatan_Id, v.Keluarahan_Id, v.Blok_Id, v.NoUrut, v.JenisOP_Id, opBumiData.NoBumi, opBumiData.KodeZNT, opBumiData.LuasBumi, v.TahunPajakskp_sppt).Scan(&rsltBumi)

		// get ObjekPajakBangunan detail
		fBng := opbg.ObjekPajakBangunan{
			NopDetail: nop.NopDetail{
				Provinsi_Kode:  v.Propinsi_Id,
				Daerah_Kode:    v.Dati2_Id,
				Kecamatan_Kode: v.Kecamatan_Id,
				Kelurahan_Kode: v.Keluarahan_Id,
				Blok_Kode:      v.Blok_Id,
				NoUrut:         v.NoUrut,
			},
		}
		a.DB.Where(&fBng).Order("\"Id\" desc").First(&opBngData)

		type RsltBng struct {
			PenilaianBangunan *float64 `gorm:"column:penilaian_bangunan"`
		}
		var rsltBng RsltBng
		a.DB.Raw("SELECT * FROM Penilaian_Bangunan(?,?,?,?,?,?,?,?,?,?,?,?)", v.Propinsi_Id, v.Dati2_Id, v.Kecamatan_Id, v.Keluarahan_Id, v.Blok_Id, v.NoUrut, v.JenisOP_Id, opBngData.NoBangunan, opBngData.Jpb_Kode, opBngData.LuasBangunan, 1, v.TahunPajakskp_sppt).Scan(&rsltBng)

		njopTkp := m.NjopTkpOld
		thnPjk, _ := strconv.Atoi(input.Tahun)
		if thnPjk >= m.NjopTkpYear {
			njopTkp = m.NjopTkpNew
		}

		nilaiBumi := float64(0)
		luasBumi := float64(0)
		nilaiBng := float64(0)
		luasBng := float64(0)
		if rsltBumi.PenilaianBumi != nil {
			nilaiBumi = *rsltBumi.PenilaianBumi
		}
		if v.LuasBumi_sppt != nil {
			luasBumi = float64(*v.LuasBumi_sppt)
		}
		if rsltBng.PenilaianBangunan != nil {
			nilaiBng = *rsltBng.PenilaianBangunan
		}
		if v.LuasBangunan_sppt != nil {
			luasBng = float64(*v.LuasBangunan_sppt)
		}
		njopBumi := int(nilaiBumi * luasBumi * 1000)
		njopBng := int(nilaiBng * luasBng * 1000)
		njopSppt := njopBumi + njopBng
		dateNow := time.Now()
		yearNow := time.Now().Format("2006")

		newSppt := m.RequestDto{
			Propinsi_Id:            v.Propinsi_Id,
			Dati2_Id:               v.Dati2_Id,
			Kecamatan_Id:           v.Kecamatan_Id,
			Keluarahan_Id:          v.Keluarahan_Id,
			Blok_Id:                v.Blok_Id,
			NoUrut:                 v.NoUrut,
			JenisOP_Id:             v.JenisOP_Id,
			KelasTanah_Id:          v.KelasTanah_Id,
			KelasBangunan_Id:       v.KelasBangunan_Id,
			KelurahanWP_sppt:       v.KelurahanWP_sppt,
			KotaWP_sppt:            v.KotaWP_sppt,
			LuasBumi_sppt:          v.LuasBumi_sppt,
			LuasBangunan_sppt:      v.LuasBangunan_sppt,
			NJOPTKP_sppt:           &njopTkp,
			NJOPBumi_sppt:          &njopBumi,
			NJOPBangunan_sppt:      &njopBng,
			NJOP_sppt:              &njopSppt,
			TanggalTerbit_sppt:     (*datatypes.Date)(&dateNow),
			TahunAwalKelasBangunan: &yearNow,
			TahunAwalKelasTanah:    &yearNow,
			TahunPajakskp_sppt:     &yearNow,
		}

		res = append(res, map[string]interface{}{
			"sppt": newSppt,
			// "penilaian_bumi":     rsltBumi,
			// "penilaian_bangunan": rsltBng,
			// "njop_tkp":           njopTkp,
			// "njop_bumi":          njopBumi,
			// "njop_bng":           njopBng,
			// "njop_sppt":          njopSppt,
		})

		// create new sppt
		Create(newSppt)

	}
	return rp.OKSimple{Data: res}, nil
}

func PenetapanMassal(input m.PenetapanMassalDto) (any, error) {
	var data []m.Sppt

	filter := m.Sppt{
		Propinsi_Id:        input.Provinsi_Kode,
		Dati2_Id:           input.Daerah_Kode,
		Kecamatan_Id:       input.Kecamatan_Kode,
		Keluarahan_Id:      input.Kelurahan_Kode,
		TahunPajakskp_sppt: &input.Tahun,
	}

	result := a.DB.Where(&filter).Order("\"Id\" asc").Find(&data)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "get-data-by-nop", source, "failed", "data tidak ada", data)
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-by-nop", source, "failed", "gagal mendapatkan data: "+result.Error.Error(), data)
	}

	var rfBuku []pn.ReferensiBuku
	var minValue, maxValue []float64
	rBuku := a.DB.Order("\"Id\" asc").Find(&rfBuku)

	if rBuku.RowsAffected == 0 {
		return sh.SetError("request", "get-data-by-nop", source, "failed", "data referensi buku tidak ada", data)
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-by-nop", source, "failed", "gagal mendapatkan data: "+result.Error.Error(), data)
	}

	for _, rb := range rfBuku {
		minValue = append(minValue, *rb.NilaiMin)
		maxValue = append(maxValue, *rb.NilaiMax)
	}
	input.BukuMin = minValue
	input.BukuMax = maxValue

	var res []map[string]interface{}
	for _, v := range data {
		var (
			opBumiData opb.ObjekPajakBumi
			opBngData  opbg.ObjekPajakBangunan
		)

		// get ObjekPajakBumi detail
		fBumi := opb.ObjekPajakBumi{
			NopDetail: nop.NopDetail{
				Provinsi_Kode:  v.Propinsi_Id,
				Daerah_Kode:    v.Dati2_Id,
				Kecamatan_Kode: v.Kecamatan_Id,
				Kelurahan_Kode: v.Keluarahan_Id,
				Blok_Kode:      v.Blok_Id,
				NoUrut:         v.NoUrut,
			},
		}
		a.DB.Where(&fBumi).Order("\"Id\" desc").First(&opBumiData)

		// get ObjekPajakBangunan detail
		fBng := opbg.ObjekPajakBangunan{
			NopDetail: nop.NopDetail{
				Provinsi_Kode:  v.Propinsi_Id,
				Daerah_Kode:    v.Dati2_Id,
				Kecamatan_Kode: v.Kecamatan_Id,
				Kelurahan_Kode: v.Keluarahan_Id,
				Blok_Kode:      v.Blok_Id,
				NoUrut:         v.NoUrut,
			},
		}
		a.DB.Where(&fBng).Order("\"Id\" desc").First(&opBngData)

		type SPRslt struct {
			PenetepanMassal *float64 `gorm:"column:penetapan_massal"`
		}
		var spRslt SPRslt

		// select * ReferensiBuku
		MinB1 := minValue[0]
		MaxB1 := maxValue[0]
		MinB2 := minValue[1]
		MaxB2 := maxValue[1]
		MinB3 := minValue[2]
		MaxB3 := maxValue[2]
		MinB4 := minValue[3]
		MaxB4 := maxValue[3]
		MinB5 := minValue[4]
		MaxB5 := maxValue[4]

		nip := "123456789"
		nipTgl := time.Now().Format("2006-01-02")
		a.DB.Raw("SELECT * FROM Penetapan_Massal(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", v.Propinsi_Id, v.Dati2_Id, v.Kecamatan_Id, v.Keluarahan_Id, v.Blok_Id, v.NoUrut, v.JenisOP_Id, v.NoPersil_sppt, v.NJOPBumi_sppt, v.NJOPBangunan_sppt, v.LuasBumi_sppt, v.LuasBangunan_sppt, 1, v.KanwilBank_Id, v.KPPBBbank_Id, v.BankTunggal_Id, v.BankPersepsi_Id, v.TP_Id, v.NJOPTKP_sppt, MinB1, MaxB1, MinB2, MaxB2, MinB3, MaxB3, MinB4, MaxB4, MinB5, MaxB5, input.Tahun, input.TglJatuhTempo1, input.TglJatuhTempo2, input.TglJatuhTempo3, input.TglJatuhTempo4, input.TglJatuhTempo5, input.TglTerbit1, input.TglTerbit2, input.TglTerbit3, input.TglTerbit4, input.TglTerbit5, nip, nipTgl).Scan(&spRslt)

	}

	return rp.OKSimple{Data: res}, nil
}
