package skpd

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

func CreateSimulasi(input m.RequestSimulasiDto) (any, error) {
	var data m.SpptSimulasi

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

func GetListSimulasi(input m.FilterSimulasiDto) (any, error) {
	var data []m.SpptSimulasi
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Sppt{}).
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

func GetDetailSimulasi(id int) (any, error) {
	var data *m.SpptSimulasi

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

func GetSimulasiByNop(nop m.SimulasiNOP, flag string) (any, error) {
	if flag == "sppt" {
		if nop.BlokID != nil {
			var data m.Sppt
			var result *gorm.DB
			result = a.DB.Where(&m.Sppt{
				Propinsi_Id:        &nop.ProvinsiID,
				Dati2_Id:           &nop.KotaID,
				Kecamatan_Id:       &nop.KecamatanID,
				Keluarahan_Id:      &nop.KelurahanID,
				Blok_Id:            nop.BlokID,
				NoUrut:             nop.NoUrut,
				JenisOP_Id:         nop.JenisOpId,
				TahunPajakskp_sppt: nop.TahunPajak,
			}).First(&data)
			if result.RowsAffected == 0 {
				return sh.SetError("request", "get-data-by-nop", source, "failed", "data tidak ada", data)
			} else if result.Error != nil {
				return sh.SetError("request", "get-data-by-nop", source, "failed", "gagal mendapatkan data: "+result.Error.Error(), data)
			}
			return rp.OKSimple{Data: data}, nil
		} else {
			var data []m.Sppt
			var result *gorm.DB
			tempTP, _ := strconv.Atoi(*nop.TahunPajak)
			tempTPs := strconv.Itoa(tempTP - 1)
			result = a.DB.Where(&m.Sppt{
				Propinsi_Id:        &nop.ProvinsiID,
				Dati2_Id:           &nop.KotaID,
				Kecamatan_Id:       &nop.KecamatanID,
				Keluarahan_Id:      &nop.KelurahanID,
				TahunPajakskp_sppt: &tempTPs,
			}).Find(&data)
			if result.RowsAffected == 0 {
				return sh.SetError("request", "get-data-by-nop", source, "failed", "data tidak ada", data)
			} else if result.Error != nil {
				return sh.SetError("request", "get-data-by-nop", source, "failed", "gagal mendapatkan data: "+result.Error.Error(), data)
			}
			return rp.OKSimple{Data: data}, nil
		}
	} else {
		if nop.BlokID != nil {
			var data m.SpptSimulasi
			var result *gorm.DB
			result = a.DB.Where(&m.SpptSimulasi{
				Propinsi_Id:        &nop.ProvinsiID,
				Dati2_Id:           &nop.KotaID,
				Kecamatan_Id:       &nop.KecamatanID,
				Keluarahan_Id:      &nop.KelurahanID,
				Blok_Id:            nop.BlokID,
				NoUrut:             nop.NoUrut,
				JenisOP_Id:         nop.JenisOpId,
				TahunPajakskp_sppt: nop.TahunPajak,
			}).First(&data)
			if result.RowsAffected == 0 {
				return sh.SetError("request", "get-data-by-nop", source, "failed", "data tidak ada", data)
			} else if result.Error != nil {
				return sh.SetError("request", "get-data-by-nop", source, "failed", "gagal mendapatkan data: "+result.Error.Error(), data)
			}
			return rp.OKSimple{Data: data}, nil
		} else {
			var data []m.SpptSimulasi
			var result *gorm.DB
			result = a.DB.Where(&m.SpptSimulasi{
				Propinsi_Id:        &nop.ProvinsiID,
				Dati2_Id:           &nop.KotaID,
				Kecamatan_Id:       &nop.KecamatanID,
				Keluarahan_Id:      &nop.KelurahanID,
				TahunPajakskp_sppt: nop.TahunPajak,
			}).First(&data)
			if result.RowsAffected == 0 {
				return sh.SetError("request", "get-data-by-nop", source, "failed", "data tidak ada", data)
			} else if result.Error != nil {
				return sh.SetError("request", "get-data-by-nop", source, "failed", "gagal mendapatkan data: "+result.Error.Error(), data)
			}
			return rp.OKSimple{Data: data}, nil
		}
	}
}

func UpdateSimulasi(id int, input m.RequestSimulasiDto) (any, error) {
	var data *m.SpptSimulasi
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

func DeleteSimulasi(id int) (any, error) {
	var data *m.SpptSimulasi
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

func DeleteSimulasiByNoP(nop m.SimulasiNOP) (any, error) {
	var data *m.SpptSimulasi
	result := a.DB.Where("Propinsi_Id", nop.ProvinsiID).
		Where("Dati2_Id", nop.KotaID).
		Where("Kecamatan_Id", nop.KecamatanID).
		Where("Keluarahan_Id", nop.KelurahanID).
		Where("NoUrut >= ").
		Find(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	result = a.DB.Where("Propinsi_Id", nop.ProvinsiID).
		Where("Dati2_Id", nop.KotaID).
		Where("Kecamatan_Id", nop.KecamatanID).
		Where("Keluarahan_Id", nop.KelurahanID).
		Delete(&data)
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
