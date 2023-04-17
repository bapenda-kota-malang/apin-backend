package spt

import (
	"context"
	"errors"
	"strconv"

	mrek "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	sbk "github.com/bapenda-kota-malang/apin-backend/internal/services/bidangkerja"
	srek "github.com/bapenda-kota-malang/apin-backend/internal/services/configuration/rekening"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/integration/bankjatim"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/google/uuid"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Verify(ctx context.Context, id uuid.UUID, input m.VerifyDto, userId uint, bidangKerjaKode string) (any, error) {
	var data m.Spt
	// validate data exist and copy input (payload) ke struct data jika tidak ada akan error
	dataRow := a.DB.First(&data, "\"Id\" = ? AND \"Type\" = ?", id.String(), mtypes.JenisPajakOA).RowsAffected
	if dataRow == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if data.StatusPenetapan == mtypes.StatusVerifikasiDisetujuiKabid {
		return sh.SetError("request", "update-data", source, "failed", "data telah disetujui", data)
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	bidangKerjaData, err := sbk.GetByKode(bidangKerjaKode)
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal data bidang kerja: "+err.Error(), data)
	}

	switch bidangKerjaData.Level {
	case mtypes.LevelJabatanKasubidKasubag:
		data.Kasubid_User_Id = &userId
		if input.StatusPenetapan == "disetujui" {
			data.StatusPenetapan = mtypes.StatusVerifikasiDisetujuiKasubid
			break
		}
		data.StatusPenetapan = mtypes.StatusVerifikasiDitolakKasubid
	case mtypes.LevelJabatanKabid:
		data.Kabid_User_Id = &userId
		if input.StatusPenetapan == "disetujui" {
			data.StatusPenetapan = mtypes.StatusVerifikasiDisetujuiKabid
			// get data rekening
			respRek, err := srek.GetDetail(int(data.Rekening_Id))
			if err != nil {
				return sh.SetError("request", "update-data", source, "failed", "rekening tidak diketahui", data)
			}
			dataRekening := respRek.(rp.OKSimple).Data.(*mrek.Rekening)
			data.KodeBilling = generateKodeBilling(dataRekening.KodeBilling, data.NomorSpt[2:])
			va, err := vaManager(ctx, data, bankjatim.ProsesCreate)
			if err != nil {
				sh.SetError("request", "update-data", source, "failed", "gagal membuat va: "+err.Error(), data)
			}
			data.VaJatim = &va
			break
		}
		data.StatusPenetapan = mtypes.StatusVerifikasiDitolakKabid
	default:
		return sh.SetError("request", "update-data", source, "failed", "level jabatan tidak diperbolehkan verifikasi", data)
	}

	switch data.StatusPenetapan {
	case mtypes.StatusVerifikasiBaru,
		mtypes.StatusVerifikasiDisetujuiKasubid,
		mtypes.StatusVerifikasiDisetujuiKabid,
		mtypes.StatusVerifikasiDitolakKasubid,
		mtypes.StatusVerifikasiDitolakKabid:
		// do nothing
	default:
		return sh.SetError("request", "update-data", source, "failed", "status penetapan tidak diketahui", data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}
	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(dataRow)),
		},
		Data: data,
	}, nil
}

func CreateSkpdkbExisting(ctx context.Context, input m.SkpdkbExisting, opts map[string]interface{}) (any, error) {
	var newDataInput m.Input
	var existingData m.Spt
	var createdNewData m.Spt
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// get old data from db
		dataRow := tx.Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).First(&existingData, "\"Id\" = ?", input.Spt_Id.String()).RowsAffected
		if dataRow == 0 {
			return errors.New("data tidak dapat ditemukan")
		}

		// casting to dto
		if existingData.DetailSptAir != nil {
			newDataInput = &m.CreateDetailAirDto{}
		} else if existingData.DetailSptHiburan != nil {
			newDataInput = &m.CreateDetailHiburanDto{}
		} else if existingData.DetailSptHotel != nil {
			newDataInput = &m.CreateDetailHotelDto{}
		} else if existingData.DetailSptParkir != nil {
			newDataInput = &m.CreateDetailParkirDto{}
		} else if existingData.DetailSptNonPln != nil {
			newDataInput = &m.CreateDetailPpjNonPlnDto{}
		} else if existingData.DetailSptPln != nil {
			newDataInput = &m.CreateDetailPpjPlnDto{}
		} else if existingData.DetailSptReklame != nil {
			newDataInput = &m.CreateDetailReklameDto{}
		} else if existingData.DetailSptResto != nil {
			newDataInput = &m.CreateDetailRestoDto{}
		} else {
			newDataInput = &m.CreateDetailBaseDto{}
		}
		// copy data from existing
		if err := newDataInput.SkpdkbDuplicate(&existingData, &input); err != nil {
			return err
		}

		// calculate skpdkb process
		newDataInput.CalculateSkpdkb()
		// create new data
		respNewData, err := CreateDetail(ctx, newDataInput, opts, tx)
		if err != nil {
			return err
		}
		createdNewData = respNewData.(rp.OKSimple).Data.(m.Spt)

		// change existing data value
		// existingData.Ref_Spt_Id = &createdNewData.Id
		// TODO: FLAG SKPDKB
		// update existing data
		if result := tx.Save(&existingData); result.Error != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-skpdkb-existing", source, "failed", "transaction skpd existing: "+err.Error(), createdNewData)
	}
	return rp.OKSimple{Data: createdNewData}, nil
}

func CreateSkpdkbNew(ctx context.Context, input m.Input, opts map[string]interface{}) (any, error) {
	var createdData m.Spt
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// save new data
		respNewData, err := CreateDetail(ctx, input, opts, tx)
		if err != nil {
			return err
		}
		createdData = respNewData.(rp.OKSimple).Data.(m.Spt)
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-skpdkb-new", source, "failed", "transaction skpd new: "+err.Error(), createdData)
	}
	return rp.OKSimple{Data: createdData}, nil
}
