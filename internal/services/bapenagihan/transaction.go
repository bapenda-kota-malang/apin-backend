package bapenagihan

import (
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	"github.com/google/uuid"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bapenagihan"
	mbapenagihandetail "github.com/bapenda-kota-malang/apin-backend/internal/models/bapenagihan/detail"
	sbapenagihandetail "github.com/bapenda-kota-malang/apin-backend/internal/services/bapenagihan/detail"
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
		inputDetails := []mbapenagihandetail.CreateDto{}
		for _, v := range input.PetugasListDetail {
			inputDetails = append(inputDetails, mbapenagihandetail.CreateDto{BaPenagihan_Id: data.Id, Petugas_User_Id: v})
		}
		respBaPenagihanDetail, err := sbapenagihandetail.Create(inputDetails, tx)
		if err != nil {
			return err
		}
		dataBaPenagihanDetail := respBaPenagihanDetail.(rp.OKSimple).Data.([]mbapenagihandetail.BaPenagihanDetail)
		data.BaPenagihanDetail = &dataBaPenagihanDetail

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

		for i := range input.BaPenagihanDetails {
			input.BaPenagihanDetails[i].BaPenagihan_Id = data.Id
		}

		respBaPenagihanDetail, err := sbapenagihandetail.Update(input.BaPenagihanDetails, tx)
		if err != nil {
			return err
		}
		dataBaPenagihanDetail := respBaPenagihanDetail.(rp.OK).Data.([]mbapenagihandetail.BaPenagihanDetail)
		data.BaPenagihanDetail = &dataBaPenagihanDetail

		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengubah data: "+err.Error(), data)
	}
	return rp.OKSimple{Data: data}, nil
}
