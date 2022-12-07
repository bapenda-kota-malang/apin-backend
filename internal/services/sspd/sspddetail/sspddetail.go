package sspd

import (
	"errors"
	"fmt"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
	ms "github.com/bapenda-kota-malang/apin-backend/internal/services/spt"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "detail sspd"

func Create(input m.SspdDetailCreateDto, sspd_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.SspdDetail
	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload sspd detail", data)
	}

	// static add value to field
	data.Sspd_Id = &sspd_id

	// Transaction save to db
	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data %s", source), data)
	}

	// ambil data sspd detail sesuai spt id
	var dataFromDb []m.SspdDetail
	result := tx.Where(m.SspdDetail{Spt_Id: data.Spt_Id}).Find(&dataFromDb)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil data %s", source), data)
	}

	// hitung total nominal bayar
	var tmpNominalBayar float64
	for _, v := range dataFromDb {
		tmpNominalBayar += v.NominalBayar
	}

	// kirim data ke spt untuk merubah status spt
	if err := ms.UpdateStatus(*data.Spt_Id, tmpNominalBayar, tx); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengupdate status data spt", data)
	}
	return rp.OKSimple{Data: data}, nil
}

func Delete(sspd_id uint64, tx *gorm.DB) error {
	var data *m.SspdDetail
	result := tx.Where(m.SspdDetail{Sspd_Id: &sspd_id}).First(&data)
	if result.RowsAffected == 0 {
		return errors.New("data sspd detail tidak dapat ditemukan")
	}

	result = tx.Delete(&data)
	if result.RowsAffected == 0 {
		return errors.New("tidak dapat menghapus data sspd detail")
	}

	return nil
}

func Update(input m.SspdDetailUpdateDto, sspd_id uint64, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}

	var data m.SspdDetail
	fmt.Println("iddhere: ", input.Id)
	//  copy input (payload) ke struct data jika tidak ada akan error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("gagal mengambil data payload %s", source), data)
	}

	var dataFromDb *m.SspdDetail
	result := tx.Where(m.SspdDetail{Id: input.Id}).First(&dataFromDb)
	if result.RowsAffected == 0 {
		return nil, errors.New("data sspd detail tidak dapat ditemukan")
	}

	data.Spt_Id = dataFromDb.Spt_Id
	err := sc.Copy(&dataFromDb, &data)
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload sspd detail", data)
	}

	dataFromDb.Sspd_Id = &sspd_id
	if result := tx.Save(&dataFromDb); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data sspd detail", dataFromDb)
	}

	// ambil data sspd detail sesuai spt id
	var dataFromDbArray []m.SspdDetail
	result = tx.Where(m.SspdDetail{Spt_Id: dataFromDb.Spt_Id}).Find(&dataFromDbArray)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil data %s", source), data)
	}

	// hitung total nominal bayar
	var tmpNominalBayar float64
	for _, v := range dataFromDbArray {
		tmpNominalBayar += v.NominalBayar
	}
	if err := ms.UpdateStatus(*dataFromDb.Spt_Id, tmpNominalBayar, tx); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengupdate status data spt", data)
	}
	return rp.OKSimple{Data: dataFromDb}, nil
}

func GetListSspdDetail(input m.SspdDetailFilterDto) (any, error) {
	var data []m.SspdDetail
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.SspdDetail{}).
		Preload(clause.Associations).
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

func UpdateFromSts(input []uint64, sts_id uint64, tx *gorm.DB) error {
	if tx == nil {
		tx = a.DB
	}

	for _, v := range input {
		var dataFromDb *m.SspdDetail
		result := tx.Where(m.SspdDetail{Sspd_Id: &v}).First(&dataFromDb)
		if result.RowsAffected == 0 {
			return errors.New("data sspd detail tidak dapat ditemukan")
		}
		dataFromDb.Sts_Id = &sts_id
		if result := tx.Save(&dataFromDb); result.Error != nil {
			return errors.New("tidak dapat menyimpan data sspd detail")
		}
	}

	return nil
}
