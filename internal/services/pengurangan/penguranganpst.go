package pengurangan

import (
	"errors"
	"fmt"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pengurangan"
	msksk "github.com/bapenda-kota-malang/apin-backend/internal/models/sksk"
	msppt "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	ssksk "github.com/bapenda-kota-malang/apin-backend/internal/services/sksk"
	ssppt "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const sourcePST = "penguranganPST"

func CreatePST(input m.CreateDtoPST) (any, error) {
	pstDetails, err := checkPstDetail(input.ThnPelayanan, input.BundelPelayanan, input.NoUrutPelayanan)
	if err != nil {
		return nil, err
	}
	if len(pstDetails) != len(input.Data) {
		return nil, errors.New("payload data tidak valid dengan pst detail")
	}
	var data []m.PenguranganPST
	for i := 0; i < len(input.Data); i++ {
		data = append(data, m.PenguranganPST{
			BaseModel: m.BaseModel{
				Kanwil_Kode:            input.Kanwil_Kode,
				Kppbb_Kode:             input.Kppbb_Kode,
				ThnPelayanan:           input.ThnPelayanan,
				BundelPelayanan:        input.BundelPelayanan,
				NoUrutPelayanan:        input.NoUrutPelayanan,
				Provinsi_Kode_Pemohon:  input.Data[i].Provinsi_Kode_Pemohon,
				Daerah_Kode_Pemohon:    input.Data[i].Daerah_Kode_Pemohon,
				Kecamatan_Kode_Pemohon: input.Data[i].Kecamatan_Kode_Pemohon,
				Kelurahan_Kode_Pemohon: input.Data[i].Kelurahan_Kode_Pemohon,
				Blok_Kode_Pemohon:      input.Data[i].Blok_Kode_Pemohon,
				NoUrut_Pemohon:         input.Data[i].NoUrut_Pemohon,
				JenisOp_Pemohon:        input.Data[i].JenisOp_Pemohon,
				JnsSk:                  input.SkSk.JnsSK,
				NoSk:                   input.SkSk.NoSK,
			},
			ThnPengPST:        &input.Data[i].ThnPengPST,
			StatusSK:          &input.Data[i].StatusSK,
			PCTPenguranganPST: &input.Data[i].PCTPenguranganPST,
		})
	}

	skskDto := msksk.CreateDto{
		PermohonanId: pstDetails[0].PermohonanId,
		KanwilId:     &input.Kanwil_Kode,
		KppbbId:      &input.Kppbb_Kode,
		JnsSK:        &input.SkSk.JnsSK,
		NoSK:         &input.SkSk.NoSK,
		TglSK:        input.SkSk.TglSK,
		NoBaKantor:   &input.SkSk.NoLhp,
		TglBaKantor:  input.SkSk.TglLhp,
	}

	err = a.DB.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(data); i++ {
			_, err := ssppt.UpdatePenguranganByNop(msppt.NopDto{
				Propinsi_Id:        &data[i].Provinsi_Kode_Pemohon,
				Dati2_Id:           &data[i].Daerah_Kode_Pemohon,
				Kecamatan_Id:       &data[i].Kecamatan_Kode_Pemohon,
				Keluarahan_Id:      &data[i].Kelurahan_Kode_Pemohon,
				Blok_Id:            &data[i].Blok_Kode_Pemohon,
				NoUrut:             &data[i].NoUrut_Pemohon,
				JenisOP_Id:         &data[i].JenisOp_Pemohon,
				TahunPajakskp_sppt: data[i].ThnPengPST,
			}, *data[i].PCTPenguranganPST, tx)
			if err != nil {
				return err
			}
		}
		if err := tx.Create(&data).Error; err != nil {
			return err
		}
		if _, err := ssksk.Create(skskDto, tx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", sourcePST, "failed", fmt.Sprintf("gagal mengambil menyimpan data: %s", err), data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetListPST(input m.FilterDtoPST) (any, error) {
	var data []m.PenguranganPST
	var count int64
	var pagination gh.Pagination

	result := a.DB.
		Model(&m.PenguranganPST{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", sourcePST, "failed", "gagal mengambil data", data)
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

func GetDetailPST(id int) (interface{}, error) {
	var data *m.PenguranganPST
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", sourcePST, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func UpdatePST(id int, input m.UpdateDtoPST) (interface{}, error) {
	var data *m.PenguranganPST

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.CopyWithOption(&data, input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "update-data", sourcePST, "failed", "gagal mengambil data payload", data)
	}
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Save(&data); result.Error != nil {
			return result.Error
		}
		if input.PCTPenguranganPST != nil {
			_, err := ssppt.UpdatePenguranganByNop(msppt.NopDto{
				Propinsi_Id:        &data.Provinsi_Kode_Pemohon,
				Dati2_Id:           &data.Daerah_Kode_Pemohon,
				Kecamatan_Id:       &data.Kecamatan_Kode_Pemohon,
				Keluarahan_Id:      &data.Kelurahan_Kode_Pemohon,
				Blok_Id:            &data.Blok_Kode_Pemohon,
				NoUrut:             &data.NoUrut_Pemohon,
				JenisOP_Id:         &data.JenisOp_Pemohon,
				TahunPajakskp_sppt: data.ThnPengPST,
			}, *data.PCTPenguranganPST, tx)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", sourceDendaADM, "failed", fmt.Sprintf("gagal mengambil menyimpan data: %s", err), data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func DeletePST(id int) (interface{}, error) {
	var data *m.PenguranganPST
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
