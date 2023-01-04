package sksk

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm/clause"

	opbangunan "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	opbumi "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	op "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sksk"
	wp "github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajakpbb"

	// ni "github.com/bapenda-kota-malang/apin-backend/internal/models/"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "sksk"

func Create(input m.CreateDto) (any, error) {
	var data m.SkSk
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	result := a.DB.Create(&data)
	if result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.SkSk
	var count int64
	var pagination gh.Pagination

	result := a.DB.
		Model(&m.SkSk{}).
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

func GetDetail(id int) (interface{}, error) {
	var data *m.SkSk
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

func GetDetailByCode(id int) (interface{}, error) {
	var data *m.SkSk
	result := a.DB.Where("Kode", fmt.Sprint(id)).First(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto) (interface{}, error) {
	var data *m.SkSk

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.Copy(&data, input); err != nil {
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

func Delete(id int) (interface{}, error) {
	var data *m.SkSk
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
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

func Cetak(input m.CetakDto) (interface{}, error) {
	var data *m.SkSk

	if input.NoSK != nil && input.TglSK != nil {
		createData := m.CreateDto{
			PermohonanId:  nil,
			KanwilId:      input.KanwilId,
			KppbbId:       input.KppbbId,
			JnsSK:         input.JnsSK,
			NoSK:          input.NoSK,
			TglSK:         input.TglSK,
			NoBaKantor:    nil,
			NoBaLapangan:  nil,
			TglBaKantor:   nil,
			TglBaLapangan: nil,
			TglCetak:      input.TglCetak,
			NipPenetak:    input.NipPenetak,
		}
		_, err := Create(createData)
		if err != nil {
			return nil, err
		}

	}

	if *input.TypeLaporan == "01" {
		result := a.DB.Find(&data)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
	}

	return rp.OKSimple{Data: data}, nil
}

func generatePenilaianToXLSX(input m.CetakDto) (*excelize.File, error) {
	var (
		start  int = 8
		dataOp *[]op.ObjekPajakPbb
		dataWp *[]wp.WajibPajakPbb
		dataBm *[]opbumi.ObjekPajakBumi
		dataBn *[]opbangunan.ObjekPajakBangunan
	)

	resultOP := a.DB.
		Where("Provinsi_Kode", input.Provinsi_Kode).
		Where("Daerah_Kode", input.Daerah_Kode).
		Where("Kecamatan_Kode", input.Kecamatan_Kode).
		Where("Kelurahan_Kode", input.Kelurahan_Kode).
		Find(&dataOp)

	if resultOP.Error != nil || resultOP.RowsAffected == 0 {
		return nil, resultOP.Error
	}

	resultBm := a.DB.
		Where("Provinsi_Kode", input.Provinsi_Kode).
		Where("Daerah_Kode", input.Daerah_Kode).
		Where("Kecamatan_Kode", input.Kecamatan_Kode).
		Where("Kelurahan_Kode", input.Kelurahan_Kode).
		Find(&dataBm)

	if resultBm.Error != nil || resultBm.RowsAffected == 0 {
		return nil, resultBm.Error
	}

	resultBn := a.DB.
		Where("Provinsi_Kode", input.Provinsi_Kode).
		Where("Daerah_Kode", input.Daerah_Kode).
		Where("Kecamatan_Kode", input.Kecamatan_Kode).
		Where("Kelurahan_Kode", input.Kelurahan_Kode).
		Find(&dataBn)

	if resultBn.Error != nil || resultBn.RowsAffected == 0 {
		return nil, resultBn.Error
	}

	// datas := nil
	xlsx := excelize.NewFile()

	laporanPenilaian := "Laporan Penilaian Individu"
	xlsx.SetSheetName(xlsx.GetSheetName(1), laporanPenilaian)

	xlsx.SetCellValue(laporanPenilaian, "A1", "Daftar PBB Belum Dibayar")
	xlsx.SetCellValue(laporanPenilaian, "A2", "Wilayah Kotamadya Malang Kecamatan "+*input.NamaKecamatan)
	xlsx.SetCellValue(laporanPenilaian, "A3", "Tahun "+*input.Tahun)
	xlsx.MergeCell(laporanPenilaian, "A1", "F1")
	xlsx.MergeCell(laporanPenilaian, "A2", "F2")
	xlsx.MergeCell(laporanPenilaian, "A3", "F3")

	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("A%d", start), "No")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("B%d", start), "Nama Kelurahan")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("C%d", start), "SPPT")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("D%d", start), "Ketetapan")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("E%d", start), "SPPT")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("F%d", start), "Ketetapan")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("G%d", start), "SPPT")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("H%d", start), "Ketetapan")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("I%d", start), "SPPT")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("J%d", start), "Ketetapan")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("K%d", start), "SPPT")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("L%d", start), "Ketetapan")

	for _, itemOP := range *dataOp {

		resultWP := a.DB.
			Preload(clause.Associations).
			Where("WajibPajakPbb_Id", itemOP.Id).
			Find(&dataWp)

		if resultWP.Error != nil || resultWP.RowsAffected == 0 {
			return nil, resultWP.Error
		}

		// err := xlsx.AutoFilter(laporanPenilaian, "A1", "AJ1", "")
		// if err != nil {
		// 	return nil, err
		// }

		// for i, data := range datas {
		// 	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("A%d", i+start+2), data.CreatedAt)
		// 	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("B%d", i+start+2), data.KotaWP_sppt)
		// 	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("C%d", i+start+2), sppt[0])
		// 	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("D%d", i+start+2), ketetapan[0])
		// 	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("E%d", i+start+2), sppt[1])
		// 	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("F%d", i+start+2), ketetapan[1])
		// 	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("G%d", i+start+2), sppt[2])
		// 	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("H%d", i+start+2), ketetapan[2])
		// 	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("I%d", i+start+2), sppt[3])
		// 	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("J%d", i+start+2), ketetapan[3])
		// 	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("K%d", i+start+2), sppt[4])
		// 	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("L%d", i+start+2), ketetapan[4])
		// }

	}
	return xlsx, nil
}
