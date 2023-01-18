package TargetRealisasi

import (
	"fmt"
	"strconv"
	"time"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	mjp "github.com/bapenda-kota-malang/apin-backend/internal/models/jenispajak"
	mnpwpd "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	mrek "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	msspd "github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/targetrealisasi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "TargetRealisasi"

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.TargetRealisasi

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data TargetRealisasi", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func CreateArray(input []m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.TargetRealisasi

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data TargetRealisasi", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.TargetRealisasi
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.TargetRealisasi{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data TargetRealisasi", data)
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
	var data *m.TargetRealisasi

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data TargetRealisasi", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.TargetRealisasi
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data TargetRealisasi", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func nestedSearchRek(dataRek []mrek.Rekening) []mrek.Rekening {
	for _, vRek := range dataRek {
		var nestedDataRek []mrek.Rekening
		if result := a.DB.Where("\"Parent_Id\" = ?", strconv.Itoa(int(vRek.Id))).Find(&nestedDataRek); result.RowsAffected == 0 {
			continue
		}

		dataRek = append(dataRek, nestedSearchRek(nestedDataRek)...)
	}
	return dataRek
}

func UpdateBySchedule() (any, error) {
	var data []m.TargetRealisasi
	result := a.DB.Where("\"Tahun\" = ?", strconv.Itoa(time.Now().Year())).Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("gagal mengambil data target realisasi tahun ini: %s", result.Error), data)
	} else if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "tidak ada data target realisasi tahun ini", data)
	}

	var dataJenisPajak []mjp.JenisPajak
	if result := a.DB.Find(&dataJenisPajak); result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "tidak ada data jenis pajak", nil)
	}

	rekeningMap := make(map[string][]uint64)
	for _, vJp := range dataJenisPajak {
		var dataRek []mrek.Rekening
		rekId := strconv.Itoa(int(*vJp.Rekening_Id))
		if result := a.DB.Where("\"Id\" = ?", rekId).Find(&dataRek); result.RowsAffected == 0 {
			continue
		}

		dataRek = nestedSearchRek(dataRek)
		dataIdRek := []uint64{}
		for _, v := range dataRek {
			dataIdRek = append(dataIdRek, v.Id)
		}
		rekeningMap[*vJp.Kode] = dataIdRek
	}

	statusNpwpd := [2]types.Status{types.StatusAktif, types.StatusTutupSementara}
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range data {
			var cWp int64
			if res := tx.
				Model(&mnpwpd.Npwpd{}).
				Where("\"Rekening_Id\" IN ?", rekeningMap[v.JenisPajak_Kode]).
				Where("\"Status\" IN ?", statusNpwpd).
				Count(&cWp); res.Error != nil {
				return res.Error
			}

			var sumReal *float64
			if res := tx.
				Model(&msspd.Sspd{}).
				Select("SUM(\"Total\")").
				Where("\"RekeningBendahara_Rekening_Id\" IN ?", rekeningMap[v.JenisPajak_Kode]).
				Scan(&sumReal); res.Error != nil {
				return res.Error
			}

			v.JumlahWp = uint64(cWp)
			if sumReal == nil {
				v.Realisasi = 0
			} else {
				v.Realisasi = *sumReal
			}
			if res := tx.Save(&v); res.Error != nil {
				return res.Error
			}
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("gagal update schedule data: %s", err), nil)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Delete(id int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.TargetRealisasi
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	result = tx.Delete(&data, id)
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
