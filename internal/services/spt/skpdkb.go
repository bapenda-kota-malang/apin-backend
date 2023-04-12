package spt

import (
	"errors"
	"strconv"
	"strings"

	mrek "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	srek "github.com/bapenda-kota-malang/apin-backend/internal/services/configuration/rekening"
	suser "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/google/uuid"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Verify(id uuid.UUID, input m.VerifyDto, userId uint) (any, error) {
	var data *m.Spt
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

	resp, err := suser.GetJabatanPegawai(userId)
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal data pegawai: "+err.Error(), data)
	}
	jabatan := strings.ToUpper(resp.(string))
	userRole := ""
	if kasubid := strings.Contains(jabatan, "KEPALA SUB BIDANG"); kasubid {
		data.Kasubid_User_Id = &userId
		userRole = "kasubid"
	} else if kabid := strings.Contains(jabatan, "KEPALA BIDANG"); kabid {
		data.Kabid_User_Id = &userId
		userRole = "kabid"
	}
	if userRole == "" {
		return sh.SetError("request", "update-data", source, "failed", "pegawai bukan kabid atau kasubid", data)
	}

	switch input.StatusPenetapan {
	case "disetujui":
		if userRole == "kasubid" {
			data.StatusPenetapan = mtypes.StatusVerifikasiDisetujuiKasubid
		} else if userRole == "kabid" {
			data.StatusPenetapan = mtypes.StatusVerifikasiDisetujuiKabid
		}
	case "ditolak":
		if userRole == "kasubid" {
			data.StatusPenetapan = mtypes.StatusVerifikasiDitolakKasubid
		} else if userRole == "kabid" {
			data.StatusPenetapan = mtypes.StatusVerifikasiDitolakKabid
		}
	default:
		return sh.SetError("request", "update-data", source, "failed", "status tidak diketahui", data)
	}

	switch data.StatusPenetapan {
	case mtypes.StatusVerifikasiDisetujuiKabid:
		// get data rekening
		respRek, err := srek.GetDetail(int(data.Rekening_Id))
		if err != nil {
			return sh.SetError("request", "update-data", source, "failed", "rekening tidak diketahui", data)
		}
		dataRekening := respRek.(rp.OKSimple).Data.(*mrek.Rekening)
		data.KodeBilling = generateKodeBilling(dataRekening.KodeBilling, data.NomorSpt[2:])
	case mtypes.StatusVerifikasiBaru,
		mtypes.StatusVerifikasiDisetujuiKasubid,
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

func CreateSkpdkbExisting(input m.SkpdkbExisting, opts map[string]interface{}) (any, error) {
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
		respNewData, err := CreateDetail(newDataInput, opts, tx)
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

func CreateSkpdkbNew(input m.Input, opts map[string]interface{}) (any, error) {
	var createdData m.Spt
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// save new data
		respNewData, err := CreateDetail(input, opts, tx)
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
