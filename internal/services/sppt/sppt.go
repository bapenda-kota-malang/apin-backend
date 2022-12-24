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
