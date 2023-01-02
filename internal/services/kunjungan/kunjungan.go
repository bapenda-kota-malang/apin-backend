package kunjungan

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/kunjungan"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "Kunjungan"

func Create(input m.CreateDto) (any, error) {
	var data m.Kunjungan
	// var peg m.KunjunganDetail

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(&data); result.Error != nil {
			return result.Error
		}

		if len(*input.Pegawais) > 0 {
			for _, item := range *input.Pegawais {
				peg := m.KunjunganDetail{
					Nip:          item.Nip,
					Kunjungan_Id: data.Id,
					Pegawai_Id:   &item.Id,
					Nama:         item.Nama,
					JabatanId:    item.JabatanId,
				}
				if result := tx.Create(&peg); result.Error != nil {
					return result.Error
				}
			}
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}
	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.Kunjungan
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Kunjungan{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data Kunjungan", data)
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
	var data *m.Kunjungan

	result := a.DB.
		First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data Kunjungan", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto) (any, error) {
	var data *m.Kunjungan
	var errSP error

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Save(&data); result.Error != nil {
			return result.Error
		}

		if len(*input.Pegawais) > 0 {
			for _, item := range *input.Pegawais {
				peg := m.KunjunganDetail{
					Nip:          item.Nip,
					Kunjungan_Id: data.Id,
					Pegawai_Id:   &item.Id,
					Nama:         item.Nama,
					JabatanId:    item.JabatanId,
				}

				result := tx.
					Where("Nip", item.Nip).
					Where("Kunjungan_Id", item.Kunjungan_Id).
					First(&peg)

				if result.RowsAffected == 0 {
					errSP = tx.Create(&peg).Error
				} else {
					errSP = tx.Where("Id", peg.Id).Save(&peg).Error
				}

				if errSP != nil {
					return errSP
				}
			}
		}

		return nil
	})

	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal merubah data: "+err.Error(), data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (any, error) {
	var data *m.Kunjungan
	var peg *[]m.KunjunganDetail
	var status string = ""
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		result = tx.Where("Kunjungan_Id", data.Id).Delete(&peg)
		if result.RowsAffected == 0 {
			data = nil
		} else if result.Error != nil {
			status = "no deletion"
			return result.Error
		}

		result = tx.Delete(&data, id)
		status = "deleted"
		if result.RowsAffected == 0 {
			data = nil
			status = "no deletion"
		} else if result.Error != nil {
			status = "no deletion"
			return result.Error
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "delete-data", source, "failed", "gagal menghapus data: "+err.Error(), data)
	}

	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil
}

func DeleteDetail(id int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.KunjunganDetail
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
