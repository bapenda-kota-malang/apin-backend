package sksk

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	ni "github.com/bapenda-kota-malang/apin-backend/internal/models/nilaiindividu"
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

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.SkSk
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	result := tx.Create(&data)
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
	// var data *m.SkSk

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
		_, err := Create(createData, nil)
		if err != nil {
			return nil, err
		}

	}

	if *input.TypeLaporan == "01" {
		resultL, err := generatePenilaianToXLSX(input)
		if err != nil {
			return nil, err
		}

		return rp.OKSimple{Data: resultL}, nil
	}

	return nil, nil
}

func generatePenilaianToXLSX(input m.CetakDto) (any, error) {
	var (
		start  int = 11
		dataOp *[]op.ObjekPajakPbb
		dataWp *[]wp.WajibPajakPbb
		dataBm *[]opbumi.ObjekPajakBumi
		dataBn *[]opbangunan.ObjekPajakBangunan
		dataNi *[]ni.NilaiIndividu
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

	resultNi := a.DB.
		Where("Provinsi_Kode", input.Provinsi_Kode).
		Where("Daerah_Kode", input.Daerah_Kode).
		Where("Kecamatan_Kode", input.Kecamatan_Kode).
		Where("Kelurahan_Kode", input.Kelurahan_Kode).
		Find(&dataNi)

	if resultNi.Error != nil || resultNi.RowsAffected == 0 {
		return nil, resultNi.Error
	}

	// datas := nil
	xlsx := excelize.NewFile()

	halaman := 1

	laporanPenilaian := "Halaman " + fmt.Sprint(halaman)
	xlsx.SetSheetName(xlsx.GetSheetName(1), laporanPenilaian)

	//header - halaman 1

	if halaman == 1 {
		tempTgl := (*time.Time)(input.TglSK).String()

		xlsx.SetCellValue(laporanPenilaian, "M1", "Lampiran V")
		xlsx.SetCellValue(laporanPenilaian, "M2", "Keputusan Menteri Keuangan")
		xlsx.SetCellValue(laporanPenilaian, "M3", "Nomor : "+*input.NoSK)
		xlsx.SetCellValue(laporanPenilaian, "M4", "Tanggal : "+tempTgl)

		xlsx.MergeCell(laporanPenilaian, "M1", "O1")
		xlsx.MergeCell(laporanPenilaian, "M2", "O2")
		xlsx.MergeCell(laporanPenilaian, "M3", "O3")
		xlsx.MergeCell(laporanPenilaian, "M4", "O4")
	}

	//per halaman

	xlsx.SetCellValue(laporanPenilaian, "A5", "KLASIFIKASI DAN BESARNYA NJOP BUMI DAN BANGUNAN")
	xlsx.SetCellValue(laporanPenilaian, "O5", "HALAMAN : "+fmt.Sprint(halaman))
	xlsx.SetCellValue(laporanPenilaian, "A6", "DENGAN NILAI INDIVIDU "+*input.Tahun)

	xlsx.MergeCell(laporanPenilaian, "A5", "N5")
	xlsx.MergeCell(laporanPenilaian, "A6", "N6")

	xlsx.SetCellValue(laporanPenilaian, "A8", "PROPINSI  : "+input.Provinsi_Kode+" - "+*input.NamaPropinsi)
	xlsx.SetCellValue(laporanPenilaian, "A9", "KAB/KOTA  : "+input.Daerah_Kode+" - "+*input.NamaDati2)
	xlsx.SetCellValue(laporanPenilaian, "M8", "KECAMATAN : "+input.Kecamatan_Kode+" - "+*input.NamaKecamatan)
	xlsx.SetCellValue(laporanPenilaian, "M9", "KEL/DESA  : "+input.Kelurahan_Kode+" - "+*input.NamaKelurahan)

	xlsx.MergeCell(laporanPenilaian, "A8", "B8")
	xlsx.MergeCell(laporanPenilaian, "A9", "B9")
	xlsx.MergeCell(laporanPenilaian, "M8", "O8")
	xlsx.MergeCell(laporanPenilaian, "M9", "O9")

	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("A%d", start), "No")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("B%d", start), "Nomor Object Pajak")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("C%d", start), "Nama & Alamat Wajib Pajak")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("D%d", start), "Rt/Rw")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("E%d", start), "Jml Bng")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("F%d", start), "Luas")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("H%d", start), "Kode ZNT")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("I%d", start), "Kelas")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("K%d", start), "NJOP(Rp.)/m2")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("M%d", start), "NJOP (Rp. 000,-)")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("O%d", start), "JUMLAH NJOP  ")

	xlsx.MergeCell(laporanPenilaian, fmt.Sprintf("A%d", start), fmt.Sprintf("A%d", start+1))
	xlsx.MergeCell(laporanPenilaian, fmt.Sprintf("B%d", start), fmt.Sprintf("B%d", start+1))
	xlsx.MergeCell(laporanPenilaian, fmt.Sprintf("C%d", start), fmt.Sprintf("C%d", start+1))
	xlsx.MergeCell(laporanPenilaian, fmt.Sprintf("D%d", start), fmt.Sprintf("D%d", start+1))
	xlsx.MergeCell(laporanPenilaian, fmt.Sprintf("E%d", start), fmt.Sprintf("E%d", start+1))
	xlsx.MergeCell(laporanPenilaian, fmt.Sprintf("H%d", start), fmt.Sprintf("H%d", start+1))
	xlsx.MergeCell(laporanPenilaian, fmt.Sprintf("O%d", start), fmt.Sprintf("O%d", start+1))

	xlsx.MergeCell(laporanPenilaian, fmt.Sprintf("F%d", start), fmt.Sprintf("G%d", start))
	xlsx.MergeCell(laporanPenilaian, fmt.Sprintf("I%d", start), fmt.Sprintf("J%d", start))
	xlsx.MergeCell(laporanPenilaian, fmt.Sprintf("K%d", start), fmt.Sprintf("L%d", start))
	xlsx.MergeCell(laporanPenilaian, fmt.Sprintf("M%d", start), fmt.Sprintf("N%d", start))

	start = start + 1

	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("F%d", start), "TANAH/TNH-BERS")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("G%d", start), "BNG/BNG-BERS")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("I%d", start), "TANAH")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("J%d", start), "BNG")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("K%d", start), "TANAH/TNH-BERS")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("L%d", start), "BNG/BNG-BERS")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("M%d", start), "TANAH/TNH-BERS")
	xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("N%d", start), "BNG/BNG-BERS")

	start = start + 1

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

		for i, itemWp := range *dataWp {

			tempOP := *itemWp.ObjekPajakPbbs // main dari nilai individu
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("A%d", i+start), i)
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("B%d", i+start), tempOP[0].NoPersil)                         // NOP DARI NILAI INDIVIDU ??? what the ??? [0] => harus nya di total but how with the number...
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("C%d", i+start), itemWp.Nama)                                // +objek pajak pbb
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("D%d", i+start), fmt.Sprintf("%d/%d", itemWp.Rt, itemWp.Rw)) //ambil dari op pbb
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("E%d", i+start), len(tempOP))                                // count op bangunan yang nop nya sama
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("F%d", i+start), tempOP[0].TotalLuasBumi)                    // dari bumi
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("G%d", i+start), tempOP[0].TotalLuasBangunan)                // dari bangunan
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("H%d", i+start), tempOP[0].Area_Kode)                        // bumi znt ??? harusnya kode znt tapi ntah dapat nya darimana ???
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("I%d", i+start), "kelas tanah")
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("J%d", i+start), "kelas bangunan")
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("K%d", i+start), tempOP[0].NjopBumi)
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("L%d", i+start), tempOP[0].NjopBangunan)
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("M%d", i+start), *tempOP[0].NjopBumi**tempOP[0].TotalLuasBumi)
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("N%d", i+start), *tempOP[0].NjopBangunan**tempOP[0].TotalLuasBangunan)
			xlsx.SetCellValue(laporanPenilaian, fmt.Sprintf("N%d", i+start), (*tempOP[0].NjopBumi**tempOP[0].TotalLuasBumi)+(*tempOP[0].NjopBangunan**tempOP[0].TotalLuasBangunan))
		}

	}
	return xlsx, nil
}
