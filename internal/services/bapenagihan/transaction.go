package bapenagihan

import (
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	"github.com/google/uuid"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bapenagihan"
	mbapenagihanpetugas "github.com/bapenda-kota-malang/apin-backend/internal/models/bapenagihan/petugas"
	mup "github.com/bapenda-kota-malang/apin-backend/internal/models/undanganpemeriksaan"
	sbapenagihanpetugas "github.com/bapenda-kota-malang/apin-backend/internal/services/bapenagihan/petugas"
	sup "github.com/bapenda-kota-malang/apin-backend/internal/services/undanganpemeriksaan"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"gorm.io/gorm"
)

func TrxCreate(input m.CreateDto, userId uint) (any, error) {
	var data m.BaPenagihan
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		respBaPenagihan, err := Create(input, userId, tx)
		if err != nil {
			return err
		}
		data = respBaPenagihan.(rp.OKSimple).Data.(m.BaPenagihan)
		inputDetails := []mbapenagihanpetugas.CreateDto{}
		for _, v := range input.PetugasListDetail {
			inputDetails = append(inputDetails, mbapenagihanpetugas.CreateDto{BaPenagihan_Id: data.Id, Petugas_Id: v})
		}
		respBaPenagihanPetugas, err := sbapenagihanpetugas.Create(inputDetails, tx)
		if err != nil {
			return err
		}
		dataBaPenagihanPetugas := respBaPenagihanPetugas.(rp.OKSimple).Data.([]mbapenagihanpetugas.BaPenagihanPetugas)
		data.BaPenagihanPetugas = &dataBaPenagihanPetugas

		selesai := mup.StatusSelesai
		_, err = sup.Update(int(data.Undangan_Id), uint64(userId), mup.UpdateDto{Status: &selesai}, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal membuat data: "+err.Error(), data)
	}
	return rp.OKSimple{Data: data}, nil
}

func TrxUpdate(id uuid.UUID, input m.UpdateDto, userId uint) (any, error) {
	var data m.BaPenagihan
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		respBaPenagihan, err := Update(id, input, userId, tx)
		if err != nil {
			return err
		}
		data = respBaPenagihan.(rp.OK).Data.(m.BaPenagihan)

		for i := range input.BaPenagihanPetugas {
			input.BaPenagihanPetugas[i].BaPenagihan_Id = data.Id
		}

		respBaPenagihanPetugas, err := sbapenagihanpetugas.Update(input.BaPenagihanPetugas, tx)
		if err != nil {
			return err
		}
		dataBaPenagihanPetugas := respBaPenagihanPetugas.(rp.OK).Data.([]mbapenagihanpetugas.BaPenagihanPetugas)
		data.BaPenagihanPetugas = &dataBaPenagihanPetugas

		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengubah data: "+err.Error(), data)
	}
	return rp.OKSimple{Data: data}, nil
}
