package sppt

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/apung/go-terbilang"
	mad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	kb "github.com/bapenda-kota-malang/apin-backend/internal/models/kelasbangunan"
	kt "github.com/bapenda-kota-malang/apin-backend/internal/models/kelastanah"
	mopb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	pmh "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"

	"github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	mnop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	m_opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	mp "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt/pembayaran"
	tp "github.com/bapenda-kota-malang/apin-backend/internal/models/tempatpembayaran"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajakpbb"
	saop "github.com/bapenda-kota-malang/apin-backend/internal/services/anggotaobjekpajak"
	sbj "github.com/bapenda-kota-malang/apin-backend/internal/services/bankjatim"
	sd "github.com/bapenda-kota-malang/apin-backend/internal/services/daerah"
	skc "github.com/bapenda-kota-malang/apin-backend/internal/services/kecamatan"
	kbs "github.com/bapenda-kota-malang/apin-backend/internal/services/kelasbangunan"
	skbng "github.com/bapenda-kota-malang/apin-backend/internal/services/kelasbangunan"
	kts "github.com/bapenda-kota-malang/apin-backend/internal/services/kelastanah"
	sktanah "github.com/bapenda-kota-malang/apin-backend/internal/services/kelastanah"
	skl "github.com/bapenda-kota-malang/apin-backend/internal/services/kelurahan"
	sopb "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakbumi"
	sopp "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakpbb"
	srefbuku "github.com/bapenda-kota-malang/apin-backend/internal/services/referensibuku"
	ssob "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt/objekbersama"
	tps "github.com/bapenda-kota-malang/apin-backend/internal/services/tempatpembayaran"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	ibj "github.com/bapenda-kota-malang/apin-backend/pkg/integration/bankjatim"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"github.com/leekchan/accounting"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/datatypes"
	"gorm.io/gorm"
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

func GetByNop(provinsiKode, daerahKode, kecamatanKode, kelurahanKode, blokKode, noUrut, jenisOp, tahun *string) (any, error) {
	var data m.Sppt
	baseFilter := a.DB.Where(
		&m.Sppt{
			Propinsi_Id:        provinsiKode,
			Dati2_Id:           daerahKode,
			Kecamatan_Id:       kecamatanKode,
			Keluarahan_Id:      kelurahanKode,
			Blok_Id:            blokKode,
			NoUrut:             noUrut,
			JenisOP_Id:         jenisOp,
			TahunPajakskp_sppt: tahun,
		},
	)
	result := baseFilter.First(&data)
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

func UpdatePenguranganByNop(input m.NopDto, pctPengurangan float64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.Sppt
	result := tx.
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
		return nil, fmt.Errorf("data sppt tidak ditemukan")
	}

	if data.PBBterhutang_sppt == nil {
		return nil, fmt.Errorf("data spp pbb yang terhutang kosong")
	}

	if data.StatusPembayaran_sppt != nil && *data.StatusPembayaran_sppt != "0" {
		return nil, fmt.Errorf("status pembayaran bukan belum bayar (0)")
	}

	faktorPengurangan := int(float64(*data.PBBterhutang_sppt) * pctPengurangan)
	data.Faktorpengurangan_sppt = &faktorPengurangan
	pbbHarusDibayar := *data.PBBterhutang_sppt - faktorPengurangan
	data.PBBygHarusDibayar_sppt = &pbbHarusDibayar

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

func UpdateVa(ctx context.Context, id int, input m.UpdateVaDto, userId uint64) (any, error) {
	var data m.Sppt
	// validate data exist and copy input (payload) ke struct data jika tidak ada akan error
	if dataRow := a.DB.First(&data, id).RowsAffected; dataRow == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if data.VirtualAccoountJatim == nil {
		return nil, errors.New("data sspt ini masih belum ada virtual account")
	}
	data.PBBygHarusDibayar_sppt = input.PBBygHarusDibayar_sppt
	total := *input.PBBygHarusDibayar_sppt
	denda := 0
	if input.Denda != nil {
		total = *data.PBBygHarusDibayar_sppt + *input.Denda
		denda = *input.Denda
	}
	if data.TanggalJatuhTempo_sppt == nil {
		return nil, errors.New("data sspt ini masih belum ada jatuh tempo")
	}
	if data.NamaWP_sppt == nil {
		return nil, errors.New("data sspt ini masih belum ada nama wp")
	}

	tglExp := time.Time(*data.TanggalJatuhTempo_sppt)
	if servicehelper.IsJatuhTempo(data.TanggalJatuhTempo_sppt) {
		tglExp = servicehelper.EndOfMonth(time.Now())
	}
	nop := servicehelper.NopString(
		*data.Propinsi_Id,
		*data.Dati2_Id,
		*data.Kecamatan_Id,
		*data.Keluarahan_Id,
		*data.Blok_Id,
		*data.NoUrut,
		*data.JenisOP_Id,
		"",
	)
	// bank jatim
	payload := ibj.RequestRegistration{
		VirtualAccount: strconv.Itoa(*data.VirtualAccoountJatim),
		Nama:           *data.NamaWP_sppt,
		TotalTagihan:   uint64(total),
		TanggalExp:     tglExp.Format("20060102"),
		Berita1:        fmt.Sprintf("%s %s", nop, time.Now().Format("06")),
		Berita2:        fmt.Sprintf("%d", denda),
		Berita5:        fmt.Sprintf("UPDATE %s", time.Now().Format("02-01-2006")),
		FlagProses:     ibj.ProsesUpdate,
	}
	ctxTo, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	if err := sbj.Registration(ctxTo, payload, userId); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengubah va: "+err.Error(), data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}
	return rp.OK{
		Meta: t.IS{
			"affected": "1",
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
		rsltBumi, rsltBng, err := SpPenilaian(v, a.DB)
		if err != nil {
			return sh.SetError("request", "penilaian-massal-sppt", source, "failed", "gagal proses store procedure: "+err.Error(), data)
		}

		// create new sppt
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
		yearNow := dateNow.Format("2006")

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

	// get data sppt
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

	// get min max value referensi buku
	minValue, maxValue, err := srefbuku.MinMaxValueReferensiBuku()
	if err != nil {
		return sh.SetError("request", "penetapan-massal-sppt", source, "failed", "min max value referensi buku: "+err.Error(), data)
	}
	input.BukuMin = minValue
	input.BukuMax = maxValue

	var res []*m.SpPenetapanMassal
	for _, v := range data {
		resSp, err := SpPenetapan(input, v, a.DB)
		if err != nil {
			return sh.SetError("request", "penetapan-massal-sppt", source, "failed", "gagal proses store procedure: "+err.Error(), data)
		}
		res = append(res, resSp)
	}

	return rp.OKSimple{Data: res}, nil
}

// this function will call SP Penilaian Bumi and SP Penilaian Bangunan
func SpPenilaian(data m.Sppt, tx *gorm.DB) (*m.SpPenilaianBumi, *m.SpPenilaianBangunan, error) {
	if tx == nil {
		tx = a.DB
	}
	var (
		rsltBumi *m.SpPenilaianBumi
		rsltBng  *m.SpPenilaianBangunan
	)

	// get ObjekPajakBumi detail
	opBumiData, err := getObjekPajakBumiDetail(data)
	if err != nil {
		return nil, nil, err
	}

	// execute sp penilaian_bumi
	res := tx.Raw("SELECT * FROM Penilaian_Bumi(?,?,?,?,?,?,?,?,?,?,?)",
		data.Propinsi_Id,
		data.Dati2_Id,
		data.Kecamatan_Id,
		data.Keluarahan_Id,
		data.Blok_Id,
		data.NoUrut,
		data.JenisOP_Id,
		opBumiData.NoBumi,
		opBumiData.KodeZNT,
		opBumiData.LuasBumi,
		data.TahunPajakskp_sppt).Scan(&rsltBumi)
	if res.Error != nil {
		return nil, nil, res.Error
	}

	// get ObjekPajakBangunan detail
	opBngData, err := getObjekPajakBangunanDetail(data)
	if err != nil {
		return nil, nil, err
	}

	// execute sp penilaian_bangunan
	res = tx.Raw("SELECT * FROM Penilaian_Bangunan(?,?,?,?,?,?,?,?,?,?,?,?)",
		data.Propinsi_Id,
		data.Dati2_Id,
		data.Kecamatan_Id,
		data.Keluarahan_Id,
		data.Blok_Id,
		data.NoUrut,
		data.JenisOP_Id,
		opBngData.NoBangunan,
		opBngData.Jpb_Kode,
		opBngData.LuasBangunan,
		1,
		data.TahunPajakskp_sppt).Scan(&rsltBng)
	if res.Error != nil {
		return nil, nil, res.Error
	}

	return rsltBumi, rsltBng, nil
}

// this function will call SP Penetapan Massal
func SpPenetapan(input m.PenetapanMassalDto, data m.Sppt, tx *gorm.DB) (*m.SpPenetapanMassal, error) {
	if tx == nil {
		tx = a.DB
	}
	var (
		rsltSp *m.SpPenetapanMassal
	)

	// get ObjekPajakBumi detail
	// opBumiData, err := getObjekPajakBumiDetail(data)
	// if err != nil {
	// 	return nil, err
	// }

	// get ObjekPajakBangunan detail
	// opBngData, err := getObjekPajakBangunanDetail(data)
	// if err != nil {
	// 	return nil, err
	// }

	// select * ReferensiBuku
	MinB1 := input.BukuMin[0]
	MaxB1 := input.BukuMax[0]
	MinB2 := input.BukuMin[1]
	MaxB2 := input.BukuMax[1]
	MinB3 := input.BukuMin[2]
	MaxB3 := input.BukuMax[2]
	MinB4 := input.BukuMin[3]
	MaxB4 := input.BukuMax[3]
	MinB5 := input.BukuMin[4]
	MaxB5 := input.BukuMax[4]

	nip := "123456789"
	nipTgl := time.Now().Format("2006-01-02")

	// execute sp penilaian_bumi
	res := tx.Raw("SELECT * FROM Penetapan_Massal(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		data.Propinsi_Id, data.Dati2_Id, data.Kecamatan_Id, data.Keluarahan_Id, data.Blok_Id, data.NoUrut, data.JenisOP_Id,
		data.NoPersil_sppt, data.NJOPBumi_sppt, data.NJOPBangunan_sppt, data.LuasBumi_sppt, data.LuasBangunan_sppt, 1,
		data.KanwilBank_Id, data.KPPBBbank_Id, data.BankTunggal_Id, data.BankPersepsi_Id, data.TP_Id, data.NJOPTKP_sppt,
		MinB1, MaxB1, MinB2, MaxB2, MinB3, MaxB3, MinB4, MaxB4, MinB5, MaxB5,
		input.Tahun, input.TglJatuhTempo1, input.TglJatuhTempo2, input.TglJatuhTempo3, input.TglJatuhTempo4, input.TglJatuhTempo5,
		input.TglTerbit1, input.TglTerbit2, input.TglTerbit3, input.TglTerbit4, input.TglTerbit5, nip, nipTgl).Scan(&rsltSp)
	if res.Error != nil {
		return nil, res.Error
	}

	return rsltSp, nil
}

func GetListCatatanSejarahWP(input m.FilterDto) (any, error) {
	var data []m.Sppt
	var count int64

	nop := &m.Sppt{
		Propinsi_Id:   input.Propinsi_Id,
		Dati2_Id:      input.Dati2_Id,
		Kecamatan_Id:  input.Kecamatan_Id,
		Keluarahan_Id: input.Keluarahan_Id,
		Blok_Id:       input.Blok_Id,
		NoUrut:        input.NoUrut,
		JenisOP_Id:    input.JenisOP_Id,
	}

	if nop.Propinsi_Id == nil {
		return sh.SetError("request", "get-data-list", source, "failed", "propinsi_id tidak boleh kosong", data)
	}

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Sppt{}).
		Where(nop).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Order(`"TahunPajakskp_sppt" asc`).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	if result.RowsAffected == 0 {
		return sh.SetError("request", "get-data-list", source, "failed", "data tidak ditemukan", data)
	}

	// get objek pajak pbb by nop
	objekPajakPbb, err := sopp.GetByNop(input.Propinsi_Id, input.Dati2_Id, input.Kecamatan_Id, input.Keluarahan_Id, input.Blok_Id, input.NoUrut, input.JenisOP_Id)
	if err != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	// find kelurahan
	var m_kelurahan areadivision.Kelurahan
	if objekPajakPbb.Kelurahan_Kode != nil {
		a.DB.Model(&areadivision.Kelurahan{}).Where(`"Id" = ?`, *objekPajakPbb.Kecamatan_Kode).First(&m_kelurahan)
	}

	// get wajib pajak pbb by id
	var m_wajibPajakPbb wajibpajakpbb.WajibPajakPbb
	if objekPajakPbb.WajibPajakPbb_Id != nil {
		a.DB.Model(&wajibpajakpbb.WajibPajakPbb{}).Where(`"Id" = ?`, *objekPajakPbb.WajibPajakPbb_Id).First(&m_wajibPajakPbb)
	}

	// responses
	var records m.ListCatatanSejarahWPResponse

	var dataRecords []m.SejarahWPResponse
	for _, v := range data {
		dataRecords = append(dataRecords, m.SejarahWPResponse{
			TahunPajak:        v.TahunPajakskp_sppt,
			TanggalCetak:      v.TanggalCetak_sppt,
			TanggalTerbit:     v.TanggalTerbit_sppt,
			TanggalJatuhTempo: v.TanggalJatuhTempo_sppt,
			NamaWajibPajak:    v.NamaWP_sppt,
			AlamatWajibPajak:  v.JalanWPskp_sppt,
		})
	}

	// response
	var (
		rt, rw string
	)

	if objekPajakPbb.Rt != nil {
		rt = *objekPajakPbb.Rt
	}

	if objekPajakPbb.Rw != nil {
		rw = *objekPajakPbb.Rw
	}

	records = m.ListCatatanSejarahWPResponse{
		AlamatObjekPajak: objekPajakPbb.Jalan,
		Kelurahan:        &m_kelurahan.Nama,
		RT_RW:            rt + "/" + rw,
		LuasTanah:        objekPajakPbb.TotalLuasBumi,
		List:             dataRecords,
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: records,
	}, nil
}

func GetListCatatanPembayaranPbb(input m.FilterDto) (any, error) {
	var data []m.Sppt
	var count int64

	nop := &m.Sppt{
		Propinsi_Id:   input.Propinsi_Id,
		Dati2_Id:      input.Dati2_Id,
		Kecamatan_Id:  input.Kecamatan_Id,
		Keluarahan_Id: input.Keluarahan_Id,
		Blok_Id:       input.Blok_Id,
		NoUrut:        input.NoUrut,
		JenisOP_Id:    input.JenisOP_Id,
	}

	if nop.Propinsi_Id == nil {
		return sh.SetError("request", "get-data-list", source, "failed", "propinsi_id tidak boleh kosong", data)
	}

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Sppt{}).
		Where(nop).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	if result.RowsAffected == 0 {
		return sh.SetError("request", "get-data-list", source, "failed", "data tidak ditemukan", data)
	}

	// get objek pajak pbb by nop
	objekPajakPbb, err := sopp.GetByNop(input.Propinsi_Id, input.Dati2_Id, input.Kecamatan_Id, input.Keluarahan_Id, input.Blok_Id, input.NoUrut, input.JenisOP_Id)
	if err != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	// get wajib pajak pbb by id
	var m_wajibPajakPbb wajibpajakpbb.WajibPajakPbb
	if objekPajakPbb.WajibPajakPbb_Id != nil {
		a.DB.Model(&wajibpajakpbb.WajibPajakPbb{}).Where(`"Id" = ?`, *objekPajakPbb.WajibPajakPbb_Id).First(&m_wajibPajakPbb)
	}

	// responses
	var records m.ListCatatanPembayaranPbbResponse

	var dataRecords []m.CatatanPembayaranPbbResponse
	for _, v := range data {
		// get pembayaran sppt
		m_spptPembayaran, _ := FindSpptPembayaranByNop(v.Propinsi_Id, v.Dati2_Id, v.Kecamatan_Id, v.Keluarahan_Id, v.Blok_Id, v.NoUrut, v.JenisOP_Id, v.TahunPajakskp_sppt)

		dataRecords = append(dataRecords, m.CatatanPembayaranPbbResponse{
			Tahun: v.TahunPajakskp_sppt,
			// TODO : jatuh tempo
			JatuhTempo:   "",
			Pbb:          v.PBBygHarusDibayar_sppt,
			JumlahBayar:  m_spptPembayaran.JumlahSpptYgDibayar,
			Ke:           m_spptPembayaran.PembayaranSpptKe,
			TanggalBayar: m_spptPembayaran.TglPembayaranSppt,
			TanggalRekam: m_spptPembayaran.CreatedAt,
			Perekam:      m_spptPembayaran.NipRekamBayarSppt,
			// TODO : need discussion about relation
			Bank: "",
		})
	}

	// response
	records = m.ListCatatanPembayaranPbbResponse{
		NamaWajibPajak:  m_wajibPajakPbb.Nama,
		JalanObjekPajak: objekPajakPbb.Jalan,
		JalanWajibPajak: m_wajibPajakPbb.Jalan,
		BlokKavNo:       objekPajakPbb.BlokKavNo,
		BlokKavNo2:      objekPajakPbb.BlokKavNo,
		List:            dataRecords,
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: records,
	}, nil
}

// find by nop
func FindSpptPembayaranByNop(provinsiKode, daerahKode, kecamatanKode, kelurahanKode, blokKode, noUrut, jenisOp, tahunPajakSppt *string) (mp.SpptPembayaran, error) {
	var data mp.SpptPembayaran
	condition := &mp.SpptPembayaran{
		Provinsi_Kd:    *provinsiKode,
		Daerah_Kd:      *daerahKode,
		Kecamatan_Kd:   *kecamatanKode,
		Kelurahan_Kd:   *kelurahanKode,
		Blok_Kd:        *blokKode,
		NoUrut_Kd:      *noUrut,
		JenisOp_Kd:     *jenisOp,
		TahunPajakSppt: *tahunPajakSppt,
	}
	result := a.DB.Where(condition).First(&data)
	if result.RowsAffected == 0 {
		return data, nil
	} else if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func Rincian(input m.NopDto) (any, error) {
	var data m.Sppt
	var rslt map[string]interface{}
	filter := m.Sppt{
		Propinsi_Id:        input.Propinsi_Id,
		Dati2_Id:           input.Dati2_Id,
		Kecamatan_Id:       input.Kecamatan_Id,
		Keluarahan_Id:      input.Keluarahan_Id,
		Blok_Id:            input.Blok_Id,
		NoUrut:             input.NoUrut,
		JenisOP_Id:         input.JenisOP_Id,
		TahunPajakskp_sppt: input.TahunPajakskp_sppt,
	}

	a.DB.Where(&filter).First(&data)

	result := a.DB.Where(&filter).First(&data)
	if result.RowsAffected == 0 || result.Error != nil {
		rslt = t.II{
			"sppt":               data,
			"opPBB":              nil,
			"kelasTanah":         nil,
			"kelasBng":           nil,
			"spptBersama":        nil,
			"pembayaranSppt":     nil,
			"tmptPembayaranSppt": nil,
		}
	} else {
		// get kelas tanah detail
		kelasTnh, err := sktanah.GetDetailByCode(*data.KelasTanah_Id)
		if err != nil {
			return nil, err
		}
		// get kelas Bng detail
		kelasBng, err := skbng.GetDetailByCode(*data.KelasBangunan_Id)
		if err != nil {
			return nil, err
		}
		// get ObjekPajakPBB detail
		opPBBData, err := GetObjekPajakPBBDetail(filter)
		if err != nil {
			return nil, err
		}
		// get sppt objek Bersama detail
		spptBersama, err := ssob.GetByNop(*data.Propinsi_Id, *data.Dati2_Id, *data.Kecamatan_Id, *data.Keluarahan_Id, *data.Blok_Id, *data.NoUrut, *data.JenisOP_Id)
		if err != nil {
			return nil, err
		}
		// get sppt pembayaran detail
		pembayaranSppt, err := getPembayaranSppt(data)
		if err != nil {
			return nil, err
		}
		// get tmptPembayaranSppt detail
		tmptPembayaranSppt, err := getTempatPembayaranSppt(data)
		if err != nil {
			return nil, err
		}

		rslt = t.II{
			"sppt":               data,
			"opPBB":              opPBBData,
			"kelasTanah":         kelasTnh.(rp.OKSimple).Data,
			"kelasBng":           kelasBng.(rp.OKSimple).Data,
			"spptBersama":        spptBersama.(rp.OKSimple).Data,
			"pembayaranSppt":     pembayaranSppt,
			"tmptPembayaranSppt": tmptPembayaranSppt,
		}
	}

	return rp.OKSimple{
		Data: rslt,
	}, nil
}

func Salinan(input m.SalinanDto) (any, error) {
	nopRange := make(map[string]interface{})
	var (
		idx int = 0
	)

	for _, v := range input.NOPRange {
		start, _ := strconv.Atoi(*v.Start.NoUrut)
		end, _ := strconv.Atoi(*v.End.NoUrut)

		var temp1 []any
		for noUrut := start; noUrut <= end; noUrut++ {
			length := 4 - int(len(fmt.Sprint(noUrut)))
			temp := fmt.Sprint(strings.Repeat("0", length), noUrut)
			gR, err := Rincian(m.NopDto{
				Propinsi_Id:        input.Propinsi_Id,
				Dati2_Id:           input.Dati2_Id,
				Kecamatan_Id:       input.Kecamatan_Id,
				Keluarahan_Id:      input.Keluarahan_Id,
				Blok_Id:            v.Start.Blok_Id,
				NoUrut:             &temp,
				JenisOP_Id:         v.Start.JenisOP_Id,
				TahunPajakskp_sppt: input.TahunPajakskp_sppt,
			})

			if err == nil {
				temp1 = append(temp1, gR.(rp.OKSimple).Data)
			}
		}
		nopRange[fmt.Sprintf("%d", idx)] = temp1
		idx++
	}

	rslt := t.II{
		"nop_range": nopRange,
	}

	return rp.OKSimple{
		Data: rslt,
	}, nil
}

type ResponseFile struct {
	ID       int    `json:"id"`
	FileName string `json:"fileName"`
}

// cetak massal
func DownloadPDF(input m.PenetapanMassalDto) (any, error) {
	var (
		dataSppt                                                                                                                   m.Sppt
		dataKelurahan                                                                                                              *mad.Kelurahan
		dataKecamatan                                                                                                              *mad.Kecamatan
		dataKota                                                                                                                   *mad.Daerah
		dataKelasTanah                                                                                                             *kt.KelasTanah
		dataKelasBangunan                                                                                                          *kb.KelasBangunan
		dataTempatPembayaran                                                                                                       *tp.TempatPembayaran
		propinsi_Id, dati2_Id, kecamatan_Id, keluarahan_Id, blok_Id, noUrut, jenisOP_Id, tanggalJatuhTempo_sppt                    string
		tahunPajakskp_sppt, namaWP_sppt, jalanWPskp_sppt, blokKavNoWP_sppt, rtWP_sppt, tanggalTerbit_sppt                          string
		rwWP_sppt, kelurahanWP_sppt, kotaWP_sppt, kelasTanah_Id, kelasBangunan_Id, jalan, blokKavNo, rt, rw                        string
		nilaiPerM2Tanah, nilaiPerM2Bangunan, namaTempatPembayaran, nJOP_sppt, nJOPTKP_sppt, pbbterhutang_sppt                      string
		luasBumi_sppt, nJOPBumi_sppt, luasBangunan_sppt, nJOPBangunan_sppt, nJOP_sppt_int, nJOPTKP_sppt_int, pbbterhutang_sppt_int int
		ac                                                                                                                         = accounting.Accounting{Symbol: "", Precision: 0, Thousand: ".", Decimal: ","}
	)

	resultSppt := a.DB.Debug().
		Where(`"Propinsi_Id" = ?`, input.Provinsi_Kode).
		Where(`"Dati2_Id" = ?`, input.Daerah_Kode).
		Where(`"Kecamatan_Id" = ?`, input.Kecamatan_Kode).
		Where(`"Keluarahan_Id" = ?`, input.Kelurahan_Kode).
		Where(`"TahunPajakskp_sppt" = ?`, input.Tahun).
		First(&dataSppt)
	if resultSppt.RowsAffected == 0 {
		return nil, nil
	} else if resultSppt.Error != nil {
		return sh.SetError("request", "download-report-pdf", source, "failed", "gagal menemukan data Sppt: ", dataSppt)
	}

	resultObjekPajakInterface, err := sopp.GetDetailbyNop(pmh.PermohonanNOP{
		PermohonanProvinsiID:  *dataSppt.Propinsi_Id,
		PermohonanKotaID:      *dataSppt.Dati2_Id,
		PermohonanKecamatanID: *dataSppt.Kecamatan_Id,
		PermohonanKelurahanID: *dataSppt.Keluarahan_Id,
		PermohonanBlokID:      *dataSppt.Blok_Id,
		NoUrutPemohon:         *dataSppt.NoUrut,
		PemohonJenisOPID:      *dataSppt.JenisOP_Id})
	if err != nil {
		return nil, err
	}
	if resultObjekPajakInterface == nil {
		return sh.SetError("request", "download-report-pdf", "objekPajakPbb", "failed", "gagal menemukan data ObjekPajakPbb: ", nil)
	}

	dataObjekPajakPbb := resultObjekPajakInterface.(rp.OKSimple).Data.(*mopb.ObjekPajakPbb)

	/** Get detail by kode kota */
	kodeKota, _ := strconv.Atoi(fmt.Sprintf("%v%v", *dataSppt.Propinsi_Id, *dataSppt.Dati2_Id))
	resultKotaInterface, err := sd.GetDetailByCode(kodeKota)
	if err != nil {
		return nil, err
	}
	if resultKotaInterface == nil {
		return sh.SetError("request", "download-report-pdf", "Derah", "failed", "gagal menemukan data Daerah: ", nil)
	} else {
		dataKota = resultKotaInterface.(rp.OKSimple).Data.(*mad.Daerah)
	}

	/** get detail by kode kecamatan */
	kodeKecamatan, _ := strconv.Atoi(fmt.Sprintf("%v%v", dataKota.Kode, *dataSppt.Kecamatan_Id))
	resultKecamatanInterface, err := skc.GetDetailByCode(kodeKecamatan)
	if err != nil {
		return nil, err
	}
	if resultKecamatanInterface == nil {
		return sh.SetError("request", "download-report-pdf", "kecamatan", "failed", "gagal menemukan data Kecamatan: ", nil)
	} else {
		dataKecamatan = resultKecamatanInterface.(rp.OKSimple).Data.(*mad.Kecamatan)
	}

	/** get detail kelurahan by kode kelurahan */
	kodeKelurahan, _ := strconv.Atoi(fmt.Sprintf("%v%v", dataKecamatan.Kode, *dataSppt.Keluarahan_Id))
	resultKelurahanInterface, err := skl.GetDetailByCode(kodeKelurahan)
	if err != nil {
		return nil, err
	}
	if resultKelurahanInterface == nil {
		return sh.SetError("request", "download-report-pdf", "kelurahan", "failed", "gagal menemukan data Kelurahan: ", nil)
	} else {
		dataKelurahan = resultKelurahanInterface.(rp.OKSimple).Data.(*mad.Kelurahan)

	}

	/** get detail kelas tanah */
	if dataSppt.KelasTanah_Id != nil {
		kelasTanahInterface, err := kts.GetDetailByCode(*dataSppt.KelasTanah_Id)
		if err != nil {
			return nil, err
		}
		if kelasTanahInterface == nil {
			return sh.SetError("request", "download-report-pdf", "kelasTanah", "failed", "gagal menemukan data kelas tanah: ", nil)
		} else {
			dataKelasTanah = kelasTanahInterface.(rp.OKSimple).Data.(*kt.KelasTanah)
		}

		if dataKelasTanah.NilaiPerM2Tanah != nil {
			nilaiPerM2Tanah = ac.FormatMoney(*dataKelasTanah.NilaiPerM2Tanah)
		}
	}

	/** get detail kelas banungn */
	if dataSppt.KelasBangunan_Id != nil {
		kelasBangunanInterface, err := kbs.GetDetailByCode(*dataSppt.KelasBangunan_Id)
		if err != nil {
			return nil, err
		}
		if kelasBangunanInterface == nil {
			return sh.SetError("request", "download-report-pdf", "kelasBangunan", "failed", "gagal menemukan data kelas bangunan: ", nil)
		} else {
			dataKelasBangunan = kelasBangunanInterface.(rp.OKSimple).Data.(*kb.KelasBangunan)
		}

		if dataKelasBangunan.NilaiPerM2Bangunan != nil {
			nilaiPerM2Bangunan = ac.FormatMoney(*dataKelasBangunan.NilaiPerM2Bangunan)
		}
	}

	/** get detail tempat pembayaran */
	if dataSppt.TP_Id != nil {
		tempatPembayaranId, _ := strconv.Atoi(*dataSppt.TP_Id)
		tempatPembayaranInterface, err := tps.GetDetail(tempatPembayaranId)
		if err != nil {
			return nil, err
		}
		if tempatPembayaranInterface == nil {
			return sh.SetError("request", "download-report-pdf", "kelasBangunan", "failed", "gagal menemukan data kelas bangunan: ", nil)
		} else {
			dataTempatPembayaran = tempatPembayaranInterface.(rp.OKSimple).Data.(*tp.TempatPembayaran)
		}

		if dataTempatPembayaran.NamaTp != nil {
			namaTempatPembayaran = *dataTempatPembayaran.NamaTp
		}
	}

	if dataSppt.Propinsi_Id != nil {
		propinsi_Id = *dataSppt.Propinsi_Id
	}
	if dataSppt.Dati2_Id != nil {
		dati2_Id = *dataSppt.Dati2_Id
	}
	if dataSppt.Kecamatan_Id != nil {
		kecamatan_Id = *dataSppt.Kecamatan_Id
	}
	if dataSppt.Keluarahan_Id != nil {
		keluarahan_Id = *dataSppt.Keluarahan_Id
	}
	if dataSppt.Blok_Id != nil {
		blok_Id = *dataSppt.Blok_Id
	}
	if dataSppt.NoUrut != nil {
		noUrut = *dataSppt.NoUrut
	}
	if dataSppt.JenisOP_Id != nil {
		jenisOP_Id = *dataSppt.JenisOP_Id
	}
	if dataSppt.TahunPajakskp_sppt != nil {
		tahunPajakskp_sppt = *dataSppt.TahunPajakskp_sppt
	}
	if dataSppt.NamaWP_sppt != nil {
		namaWP_sppt = *dataSppt.NamaWP_sppt
	}
	if dataSppt.JalanWPskp_sppt != nil {
		jalanWPskp_sppt = *dataSppt.JalanWPskp_sppt
	}
	if dataSppt.BlokKavNoWP_sppt != nil {
		blokKavNoWP_sppt = *dataSppt.BlokKavNoWP_sppt
	}
	if dataSppt.RtWP_sppt != nil {
		rtWP_sppt = *dataSppt.RtWP_sppt
	}
	if dataSppt.RwWP_sppt != nil {
		rwWP_sppt = *dataSppt.RwWP_sppt
	}
	if dataSppt.KelurahanWP_sppt != nil {
		kelurahanWP_sppt = *dataSppt.KelurahanWP_sppt
	}
	if dataSppt.KotaWP_sppt != nil {
		kotaWP_sppt = *dataSppt.KotaWP_sppt
	}
	if dataSppt.LuasBumi_sppt != nil {
		luasBumi_sppt = *dataSppt.LuasBumi_sppt
	}
	if dataSppt.KelasTanah_Id != nil {
		kelasTanah_Id = *dataSppt.KelasTanah_Id
	}
	if dataSppt.NJOPBumi_sppt != nil {
		nJOPBumi_sppt = *dataSppt.NJOPBumi_sppt
	}
	if dataSppt.LuasBangunan_sppt != nil {
		luasBangunan_sppt = *dataSppt.LuasBangunan_sppt
	}
	if dataSppt.KelasBangunan_Id != nil {
		kelasBangunan_Id = *dataSppt.KelasBangunan_Id
	}
	if dataSppt.NJOPBangunan_sppt != nil {
		nJOPBangunan_sppt = *dataSppt.NJOPBangunan_sppt
	}
	if dataSppt.NJOP_sppt != nil {
		nJOP_sppt_int = *dataSppt.NJOP_sppt
		nJOP_sppt = ac.FormatMoney(*dataSppt.NJOP_sppt)
	}
	if dataSppt.NJOPTKP_sppt != nil {
		nJOPTKP_sppt_int = *dataSppt.NJOPTKP_sppt
		nJOPTKP_sppt = ac.FormatMoney(*dataSppt.NJOPTKP_sppt)
	}
	if dataSppt.PBBterhutang_sppt != nil {
		pbbterhutang_sppt_int = *dataSppt.PBBterhutang_sppt
		pbbterhutang_sppt = ac.FormatMoney(*dataSppt.PBBterhutang_sppt)
	}
	if dataSppt.TanggalJatuhTempo_sppt != nil {
		tanggalJatuhTempo_sppt = timehelper.ConvertDatatypeDateToTime(dataSppt.TanggalJatuhTempo_sppt).Format("02-Jan-2006")
	}
	if dataSppt.TanggalTerbit_sppt != nil {
		tanggalTerbit_sppt = timehelper.ConvertDatatypeDateToTime(dataSppt.TanggalTerbit_sppt).Format("02-Jan-2006")
	}
	if dataObjekPajakPbb.Jalan != nil {
		jalan = *dataObjekPajakPbb.Jalan
	}
	if dataObjekPajakPbb.BlokKavNo != nil {
		blokKavNo = *dataObjekPajakPbb.BlokKavNo
	}
	if dataObjekPajakPbb.Rt != nil {
		rt = *dataObjekPajakPbb.Rt
	}
	if dataObjekPajakPbb.Rw != nil {
		rw = *dataObjekPajakPbb.Rw
	}

	finalresult := map[string]interface{}{
		"NOP": fmt.Sprintf("%v.%v.%v.%v.%v-%v.%v",
			propinsi_Id,
			dati2_Id,
			kecamatan_Id,
			keluarahan_Id,
			blok_Id,
			noUrut,
			jenisOP_Id),
		"TahunPajakSppt":        tahunPajakskp_sppt,
		"AlamatObjekPajak":      fmt.Sprintf("%v %v", jalan, blokKavNo),
		"NamaWPSppt":            namaWP_sppt,
		"RTRWObjekPajakPbb":     fmt.Sprintf("RT %v/RW %v", rt, rw),
		"AlamatWPPBBNo":         fmt.Sprintf("%v %v", jalanWPskp_sppt, blokKavNoWP_sppt),
		"NamaKelurahan":         dataKelurahan.Nama,
		"NamaKecamatan":         dataKecamatan.Nama,
		"NamaKota":              dataKota.Nama,
		"RtRwSppt":              fmt.Sprintf("RT %v/RW.%v", rtWP_sppt, rwWP_sppt),
		"NamaKelurahanSppt":     kelurahanWP_sppt,
		"NamaKotaSppt":          kotaWP_sppt,
		"LuasBumiSppt":          luasBumi_sppt,
		"KelasTanahSppt":        kelasTanah_Id,
		"NilaiPerM2Tanah":       nilaiPerM2Tanah,
		"NJOPBumi_sppt":         nJOPBumi_sppt,
		"LuasBangunanSppt":      luasBangunan_sppt,
		"KelasBangunanSppt":     kelasBangunan_Id,
		"NilaiPerM2Bangunan":    nilaiPerM2Bangunan,
		"NJOPBangunan_sppt":     nJOPBangunan_sppt,
		"NJOPSppt":              nJOP_sppt,
		"NJOPTKPSppt":           nJOPTKP_sppt,
		"Total":                 ac.FormatMoney(nJOP_sppt_int - nJOPTKP_sppt_int),
		"PresentaseNJOPSppt":    "0.55%",
		"TotalBersih":           pbbterhutang_sppt,
		"Terbilang":             cases.Title(language.Und, cases.NoLower).String(terbilang.ToTerbilangRp(pbbterhutang_sppt_int)),
		"TanggalJatuhTempoSppt": tanggalJatuhTempo_sppt,
		"TanggalTerbitSppt":     tanggalTerbit_sppt,
		"NamaTempatPembayaran":  namaTempatPembayaran,
		"PBBTerhitungSppt":      pbbterhutang_sppt,
	}

	uuid, err := sh.GetUuidv4()
	if err != nil {
		return nil, err
	}
	fileName := sh.GenerateFilename("cetak-massal", uuid, 0, "pdf")

	outFile := filepath.Join(sh.GetPdfPath(), fileName)
	if err := GeneratePdf(outFile, finalresult); err != nil {
		return nil, err
	}
	_r := &ResponseFile{
		// ID:       id,
		FileName: outFile,
	}
	return rp.OKSimple{Data: _r}, nil
}

func GetListCatatanSejarahOp(input m.CatatanSejarahOPDto) (any, error) {
	var (
		dataRecords []m_opp.ObjekPajakPbb
		count       int64
	)

	// pagination
	var pagination gh.Pagination

	// kelurahan kode = provinsi kode + daerah kode + kecamatan kode + kelurahan kode
	kelurahanKode := fmt.Sprintf("%s%s%s%s", *input.Provinsi_Kode, *input.Daerah_Kode, *input.Kecamatan_Kode, *input.Kelurahan_Kode)

	// wilayah
	condition := m_opp.ObjekPajakPbb{
		Provinsi_Kode:  input.Provinsi_Kode,
		Daerah_Kode:    input.Daerah_Kode,
		Kecamatan_Kode: input.Kecamatan_Kode,
		Kelurahan_Kode: &kelurahanKode,
	}

	// get data
	result := a.DB.Model(&m_opp.ObjekPajakPbb{}).
		Debug().
		Where(&condition).
		// daterange by tanggal mutasi on created_at
		Where(`"CreatedAt" BETWEEN ? AND ?`, input.Tanggal_Mutasi_Start, input.Tanggal_Mutasi_End).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&dataRecords)

	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return nil, result.Error
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: dataRecords,
	}, nil
}

func GetRekapitulasi(input m.RekapitulasiOpRequest) (any, error) {

	filter := m.Sppt{
		Propinsi_Id:        input.Propinsi_Id,
		Dati2_Id:           input.Dati2_Id,
		Kecamatan_Id:       input.Kecamatan_Id,
		TahunPajakskp_sppt: &input.TahunPajakskp_sppt,
	}

	nop := &mnop.NopDetail{
		Provinsi_Kode:  input.Propinsi_Id,
		Daerah_Kode:    input.Dati2_Id,
		Kecamatan_Kode: input.Kecamatan_Id,
	}

	type RekapSppt struct {
		Keluarahan_Id          *string
		TahunPajakskp_sppt     *string `gorm:"column:tahunpajak"`
		JumlahSppt             *int    `gorm:"column:jumlahsppt"`
		PBBterhutang_sppt      *int    `gorm:"column:pbbterhutang"`
		PBBygHarusDibayar_sppt *int    `gorm:"column:pbbygharusdibayar"`
		Lunas                  *int    `gorm:"column:lunas"`
		JatuhTempo             *int    `gorm:"column:jatuhtempo"`
	}

	var dataRekapSppt []RekapSppt
	resultSppt := a.DB.Debug().Model(m.Sppt{}).
		Select("\"Keluarahan_Id\", \"TahunPajakskp_sppt\" as tahunpajak, COUNT(\"Sppt\") AS jumlahsppt, " +
			"SUM(\"PBBterhutang_sppt\") AS pbbterhutang, SUM(\"PBBygHarusDibayar_sppt\") AS pbbygharusdibayar, " +
			"SUM(\"StatusPembayaran_sppt\"::INTEGER) AS lunas, SUM(CASE WHEN \"TanggalJatuhTempo_sppt\" <= CURRENT_DATE THEN 1 ELSE 0 END) AS jatuhtempo").
		Where(&filter).
		Group("\"Keluarahan_Id\"").Group("\"TahunPajakskp_sppt\"").
		Order("\"Keluarahan_Id\"").
		Scan(&dataRekapSppt)
	if resultSppt.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data sppt: "+resultSppt.Error.Error(), dataRekapSppt)
	}

	type RekapOpbb struct {
		JumlahOp          *int `gorm:"column:jumlahop"`
		TotalLuasBumi     *int `gorm:"column:totalluasbumi"`
		TotalLuasBangunan *int `gorm:"column:totalluasbangunan"`
	}

	var dataRekapOpbb []RekapOpbb
	resultOpbb := a.DB.Debug().Model(objekpajakpbb.ObjekPajakPbb{}).
		Select("COUNT(*) AS jumlahop, SUM(\"TotalLuasBumi\") AS totalluasbumi, SUM(\"TotalLuasBangunan\") AS totalluasbangunan").
		Where(&objekpajakpbb.ObjekPajakPbb{NopDetail: *nop}).Group("\"Kelurahan_Kode\"").Order("\"Kelurahan_Kode\"").Scan(&dataRekapOpbb)
	if resultOpbb.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data objek pajak pbb: "+resultOpbb.Error.Error(), dataRekapOpbb)
	}

	type RekapOpBgn struct {
		JumlahBgn   *int `gorm:"column:jumlahbgn"`
		NilaiSistem *int `gorm:"column:nilaisistem"`
	}

	var dataRekapOpBgn []RekapOpBgn
	resultOpBgn := a.DB.Debug().Model(objekpajakbangunan.ObjekPajakBangunan{}).
		Select("COUNT(*) AS jumlahbgn, SUM(\"NilaiSistem\") AS nilaisistem").
		Where(&objekpajakbangunan.ObjekPajakBangunan{NopDetail: *nop}).Group("\"Kelurahan_Kode\"").Order("\"Kelurahan_Kode\"").Scan(&dataRekapOpBgn)
	if resultOpBgn.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data objek pajak bangunan: "+resultOpBgn.Error.Error(), dataRekapOpBgn)
	}

	type RekapOpBumi struct {
		NilaiSistem *int `gorm:"column:nilaisistem"`
	}

	var dataRekapOpBumi []RekapOpBumi
	resultOpBumi := a.DB.Debug().Model(objekpajakbumi.ObjekPajakBumi{}).
		Select("SUM(\"NilaiSistemBumi\") AS nilaisistem").
		Where(&objekpajakbumi.ObjekPajakBumi{NopDetail: *nop}).Group("\"Kelurahan_Kode\"").Order("\"Kelurahan_Kode\"").Scan(&dataRekapOpBumi)
	if resultOpBumi.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data objek pajak bangunan: "+resultOpBumi.Error.Error(), dataRekapOpBumi)
	}

	type RekapSpptPembayaran struct {
		SpptDibayar *int `gorm:"column:spptdibayar"`
	}

	var dataRekapPembayaran []RekapSpptPembayaran
	resultPembayaran := a.DB.Debug().Model(mp.SpptPembayaran{}).
		Select("SUM(\"JumlahSpptYgDibayar\") AS spptdibayar").
		Where(&mp.FilterDto{
			Provinsi_Kd:    input.Propinsi_Id,
			Daerah_Kd:      input.Dati2_Id,
			Kecamatan_Kd:   input.Kecamatan_Id,
			TahunPajakSppt: &input.TahunPajakskp_sppt}).
		Group("\"Kelurahan_Kd\"").Order("\"Kelurahan_Kd\"").Scan(&dataRekapPembayaran)
	if resultPembayaran.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data objek pajak bangunan: "+resultPembayaran.Error.Error(), dataRekapPembayaran)
	}

	dataResponse := []m.RekapitulasiOpResponse{}
	dataResponse = make([]m.RekapitulasiOpResponse, 0)
	var skpop = 0

	for i := 0; i < len(dataRekapSppt); i++ {
		dataResponse = append(dataResponse, m.RekapitulasiOpResponse{
			KelurahanKode:     dataRekapSppt[i].Keluarahan_Id,
			JumlahObjekPajak:  dataRekapOpbb[i].JumlahOp,
			JumlahBangunan:    dataRekapOpBgn[i].JumlahBgn,
			LuasTotalBumi:     dataRekapOpbb[i].TotalLuasBumi,
			LuasTotalBangunan: dataRekapOpbb[i].TotalLuasBangunan,
			NjopBumi:          dataRekapOpBumi[i].NilaiSistem,
			NjopBangunan:      dataRekapOpBgn[i].NilaiSistem,
			TahunPajak:        dataRekapSppt[i].TahunPajakskp_sppt,
			JumlahSppt:        dataRekapSppt[i].JumlahSppt,
			PbbTerhutang:      dataRekapSppt[i].PBBterhutang_sppt,
			PbbHarusDibayar:   dataRekapSppt[i].PBBygHarusDibayar_sppt,
			Lunas:             dataRekapSppt[i].Lunas,
			JatuhTempo:        dataRekapSppt[i].JatuhTempo,
			PembayaranSppt:    dataRekapPembayaran[i].SpptDibayar,
			PembayaranSkpSpop: &skpop,
		})
	}

	return rp.OKSimple{
		Data: dataResponse,
	}, nil
}
