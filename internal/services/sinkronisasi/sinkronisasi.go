package sinkronisasi

import (
	"errors"
	"fmt"
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sinkronisasi"
	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "sinkronisasi"

func GetList(input m.FilterDto) (any, error) {
	var data []m.Sinkronisasi
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Sinkronisasi{}).
		Preload(clause.Associations).
		Preload("User").
		Preload("Tbp").
		Preload("DetailSinkronisasi.DetailTbp").
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

func GetDetail(tbp_id int) (any, error) {
	var data *m.Sinkronisasi
	err := a.DB.
		Model(&m.Sinkronisasi{}).
		Preload(clause.Associations).
		Preload("User").
		Preload("Tbp").
		Preload("DetailSinkronisasi.DetailTbp").
		First(&data, tbp_id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return rp.OKSimple{
		Data: data,
	}, err
}

func Create(input m.CreateDto, user_Id uint64) (any, error) {
	var dataSinkronisasi m.Sinkronisasi
	var dataSpt []ms.Spt
	var mapSpt map[string]ms.Spt
	var dataSinkronisasiMerge []m.SinkronisasiMerge
	var resp t.II
	var errChan = make(chan error)
	baseDocsName := "sinkronisasi"

	fileName, path, extFile, err := filePreProcess(input.File, uint(user_Id), baseDocsName+"File")
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}

	go sh.SaveFile(input.File, fileName, path, extFile, errChan)
	if err := <-errChan; err != nil {
		return sh.SetError("request", "create-data", source, "failed", "excel unsupported", input)
	}

	// data sinkronisasi
	if err := sc.Copy(&dataSinkronisasi, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload sinkronisasi", dataSinkronisasi)
	}

	result := a.DB.Where(ms.Spt{TanggalSpt: *input.TanggalSinkronisasi}).Find(&dataSpt)
	if result.Error != nil {
		return nil, result.Error
	}

	dataSinkronisasi.File = fileName
	dataSinkronisasi.User_Id = &user_Id

	for k := range dataSpt {
		mapSpt[*dataSpt[k].VaJatim] = dataSpt[k]
	}

	switch input.JenisPajak {
	case "pdl":
		dataExcel, err := readExcelFilePdl(fileName)
		if err != nil {
			return nil, err
		}
		for k := range dataExcel {
			tmpDataSinkronisasiMerge := m.SinkronisasiMerge{
				Excel_Nominal:      dataExcel[k].Nominal,
				Excel_NominalDenda: dataExcel[k].Nominal,
			}
			if k, exist := mapSpt[*dataExcel[k].NoRekening]; exist {
				tmpDataSinkronisasiMerge.Spt_Id = k.Id
				tmpDataSinkronisasiMerge.Spt_Nominal = &mapSpt[k].JumlahPajak
				tmpDataSinkronisasiMerge.Spt_NominalDenda = mapSpt[k].Denda
			}
			delete(mapSpt, *dataExcel[k].NoRekening)
			dataSinkronisasiMerge = append(dataSinkronisasiMerge, tmpDataSinkronisasiMerge)
		}

		if len(mapSpt) > 0 {
			for _, v := range mapSpt {
				tmpDataSinkronisasiMerge := m.SinkronisasiMerge{
					Spt_Id:           v.Id,
					Spt_Nominal:      &v.JumlahPajak,
					Spt_NominalDenda: v.Denda,
				}
				dataSinkronisasiMerge = append(dataSinkronisasiMerge, tmpDataSinkronisasiMerge)
			}
			fmt.Println("data excel: ", dataExcel)
		}
	case "pbb":
		dataExcel, err := readExcelFilePbbBillAgregat(fileName)
		if err != nil {
			return nil, err
		}
		fmt.Println("data excel: ", dataExcel)
	}

	err = a.DB.Transaction(func(tx *gorm.DB) error {

		// create data sinkronisasi
		err := tx.Create(&dataSinkronisasi).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataSinkronisasi)
	}

	resp = t.II{
		"sinkronisasi": dataSinkronisasiMerge,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func Update(id int, input m.UpdateDto, user_id uint64) (any, error) {
	var dataSinkronisasi *m.Sinkronisasi
	var resp t.II

	result := a.DB.First(&dataSinkronisasi, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tbp tidak dapat ditemukan")
	}

	if err := sc.Copy(&dataSinkronisasi, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload sinkronisasi", input)
	}

	dataSinkronisasi.User_Id = &user_id

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// update tbp
		if result := tx.Save(&dataSinkronisasi); result.Error != nil {
			return errors.New("gagal menyimpan data sinkronisasi")
		}

		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataSinkronisasi)
	}

	resp = t.II{
		"affected":     strconv.Itoa(int(result.RowsAffected)),
		"sinkronisasi": dataSinkronisasi,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}
