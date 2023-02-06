package anggotaobjekpajak

import (
	"fmt"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/anggotaobjekpajak"
	msppt "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "anggotaobjekpajak"

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.AnggotaObjekPajak

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data anggotaobjekpajak: "+result.Error.Error(), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.AnggotaObjekPajak
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.AnggotaObjekPajak{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data anggotaobjekpajak", data)
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
	var data *m.AnggotaObjekPajak

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data anggotaobjekpajak", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.AnggotaObjekPajak
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data anggotaobjekpajak", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.AnggotaObjekPajak
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

func GetByNop(nop string, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	resultNop, _ := sh.NopParser(nop)
	condition := m.AnggotaObjekPajak{
		Provinsi_Kode:  &resultNop[0],
		Daerah_Kode:    &resultNop[1],
		Kecamatan_Kode: &resultNop[2],
		Kelurahan_Kode: &resultNop[3],
		Blok_Kode:      &resultNop[4],
		NoUrut:         &resultNop[5],
		JenisOp:        &resultNop[6],
	}

	var data *m.AnggotaObjekPajak
	result := tx.Where(condition).First(&data)
	if result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data anggota objek pajak", data)
	}
	if result.RowsAffected == 0 {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("data dengan nop %s tidak ditemukan", nop), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func PenilaianSppt(input []msppt.Sppt, tx *gorm.DB) (any, error) {
	var data *m.AnggotaObjekPajak
	for k, v := range input {
		condition := nopSearcher(v)
		result := tx.Where(condition).First(&data)
		// if result.Error != nil {
		// 	return sh.SetError("request", "penilaian-data", source, "failed", "gagal mengambil data", data)
		// }
		if result.RowsAffected != 0 {
			// TODO: nilai sistem
			data.NjopBangunanBeban = input[k].NJOPBangunan_sppt
			data.NjopBumiBeban = input[k].NJOPBumi_sppt
			if resultSave := tx.Save(&data); resultSave.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data objek pajak pbb", data)
			}
		}

	}
	return rp.OKSimple{Data: data}, nil
}
