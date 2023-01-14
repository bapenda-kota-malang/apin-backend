package jaminanbongkar

import (
	"fmt"
	"math"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/jaminanbongkar"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mdsrek "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifjambong"
	mtjrek "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifjambongrek"

	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "jaminanbongkar"

func getTjRek(tx *gorm.DB, tipe string, luas *float64) (data *mtjrek.TarifJambongRek, err error) {
	baseQuery := tx.Model(&mtjrek.TarifJambongRek{}).Where("\"Tipe\" = ?", tipe)
	if luas != nil {
		baseQuery = baseQuery.Where("\"BatasBawah\" <= ?", *luas).Order("\"BatasBawah\" DESC")
	}
	if resStjrek := baseQuery.First(&data); resStjrek.Error != nil {
		err = resStjrek.Error
	} else if resStjrek.RowsAffected == 0 {
		err = fmt.Errorf("data tarif jambong rek not found")
	}
	return
}

func calcLuas(data mdsrek.DetailSptReklame) (luas float64) {
	switch data.JenisDimensi {
	case mdsrek.JenisDimensiPersegiPanjang:
		luas = *data.Panjang * *data.Lebar
	case mdsrek.JenisDimensiPersegiLingkaran:
		r := *data.Diameter / 2
		luas = math.Pi * math.Pow(r, 2)
	}
	luas *= float64(*data.Sisi) * float64(data.Jumlah)
	return
}

func processData(tx *gorm.DB, data *m.JaminanBongkar, dataDetails []m.DetailJambong, typeProcess string, userId uint) error {
	// get data spt with detail and tarif reklame
	var dataSpt spt.Spt
	result := tx.Preload("DetailSptReklame.TarifReklame").First(&dataSpt, "\"Id\" = ?", data.Spt_Id.String())
	if result.RowsAffected == 0 {
		return fmt.Errorf("data spt not found")
	} else if result.Error != nil {
		return fmt.Errorf("gagal mengambil data spt: %w", result.Error)
	}

	if dataSpt.DetailSptReklame == nil {
		return fmt.Errorf("tidak ada data spt reklame")
	}

	if typeProcess != "update" {
		data.CreateBy_User_Id = userId
		data.Nomor = fmt.Sprintf("JB-0000%s", dataSpt.NomorSpt[len(dataSpt.NomorSpt)-5:])
	}
	data.Nominal = data.BiayaPemutusan

	var dataTjrek *mtjrek.TarifJambongRek
	if data.JenisReklame == m.ReklameInsidentil {
		var err error
		dataTjrek, err = getTjRek(tx, "Insidentil", nil)
		if err != nil {
			return fmt.Errorf("gagal mengambil data tarif jambong rekening: %w", err)
		}
		data.TarifJambongRek_Id = &dataTjrek.Id
	}

	// calc every detail spt reklame and append to detail jambong
	for _, v := range *dataSpt.DetailSptReklame {
		dataDetail := m.DetailJambong{DetailSptReklame_id: v.Id}

		if data.JenisReklame == m.ReklameInsidentil {
			data.Nominal += (v.JumlahRp * (*dataTjrek.Nominal / float64(100)))
		} else {
			switch *v.TarifReklame.JenisReklame {
			// tarif jambong rek
			case "Billboard Disinari", "Billboard, Tembok/Tugu, Shop Panel, Letter, Neon Sign, Prismatek", "Neon Box":
				tipe := "RB"
				if *v.TarifReklame.JenisReklame == "Neon Box" {
					tipe = "RNB"
				}

				if *v.TarifReklame.DasarPengenaan != "Luas" {
					return fmt.Errorf("hanya bisa tarif reklame dasar pengenaan luas")
				}

				luas := calcLuas(v)

				var err error
				dataTjrek, err = getTjRek(tx, tipe, &luas)
				if err != nil {
					return fmt.Errorf("gagal mengambil data tarif jambong rekening: %w", err)
				}
				data.Nominal += luas * *dataTjrek.Nominal
			// tarif jambong
			case "Megatron, TV Media", "Rombong", "Bando Jalan, JPO, Taman Gantung":
				jenisRekStr := *v.TarifReklame.JenisReklame
				switch jenisRekStr {
				case "Megatron, TV Media":
					jenisRekStr = "Megatron"
				}

				if *v.TarifReklame.DasarPengenaan != "Luas" {
					return fmt.Errorf("hanya bisa tarif reklame dasar pengenaan luas")
				}

				luas := calcLuas(v)

				var dataTj tarifjambong.TarifJambong
				res := tx.Model(&tarifjambong.TarifJambong{}).Where("\"JenisReklame\" LIKE ?", fmt.Sprintf("%%%s%%", jenisRekStr)).First(&dataTj)
				if res.Error != nil {
					return fmt.Errorf("gagal mengambil data tarif jambong: %w", res.Error)
				} else if res.RowsAffected == 0 {
					return fmt.Errorf("data tarif jambong kosong")
				}
				dataDetail.TarifJambong_Id = &dataTj.Id
				data.Nominal += luas * *dataTj.Nominal
			// insidentil, tarif jambong rek
			default:
				if dataTjrek == nil || (dataTjrek != nil && *dataTjrek.Tipe != "Insidentil") {
					var err error
					dataTjrek, err = getTjRek(tx, "Insidentil", nil)
					if err != nil {
						return fmt.Errorf("gagal mengambil data tarif jambong: %w", err)
					}
					data.TarifJambongRek_Id = &dataTjrek.Id
				}
				data.Nominal += (v.JumlahRp * (*dataTjrek.Nominal / float64(100)))
			}
		}

		// TODO: change this when update replace slice data details, should be not append
		if typeProcess != "update" {
			dataDetails = append(dataDetails, dataDetail)
		}
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	err := tx.Transaction(func(tx2 *gorm.DB) error {
		if typeProcess == "update" {
			if result := tx2.Save(&data); result.Error != nil {
				return result.Error
			}
			if result := tx2.Save(&dataDetails); result.Error != nil {
				return result.Error
			}
			return nil
		}

		if result := tx2.Create(&data); result.Error != nil {
			return result.Error
		}

		for i := range dataDetails {
			dataDetails[i].JaminanBongkar_Id = data.Id
		}

		if result := tx2.Create(&dataDetails); result.Error != nil {
			return result.Error
		}

		return nil
	})
	return err
}

func Create(input m.CreateDto, userId uint, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.JaminanBongkar
	var dataDetails []m.DetailJambong

	typeData := ""
	if tx.Model(&m.JaminanBongkar{}).
		Preload("DetailJambong").
		Where("\"Spt_Id\" = ?", input.Spt_Id.String()).
		First(&data).
		RowsAffected == 0 {
		// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
		if err := sc.Copy(&data, &input); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
		}
	} else {
		if err := sc.CopyWithOption(&data, &input, sc.Option{IgnoreEmpty: true}); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
		}
		if err := sc.CopyWithOption(&dataDetails, &data.DetailJambong, sc.Option{IgnoreEmpty: true}); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
		}
		data.DetailJambong = []m.DetailJambong{}
		typeData = "update"
	}

	if err := processData(tx, &data, dataDetails, typeData, userId); err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal menyimpan data: %s", err), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.JaminanBongkar
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.JaminanBongkar{}).
		Preload("Spt.DetailSptReklame").
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
	var data *m.JaminanBongkar

	result := a.DB.Preload("Spt.DetailSptReklame").Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
		return tx.Omit("Password")
	}).First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.JaminanBongkar
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

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

func Delete(id int) (any, error) {
	var data *m.JaminanBongkar
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
