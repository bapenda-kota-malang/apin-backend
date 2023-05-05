package undanganpemeriksaan

import (
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/undanganpemeriksaan"
	muppemeriksa "github.com/bapenda-kota-malang/apin-backend/internal/models/undanganpemeriksaan/pemeriksa"
	suppemeriksa "github.com/bapenda-kota-malang/apin-backend/internal/services/undanganpemeriksaan/pemeriksa"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"gorm.io/gorm"
)

func TrxCreate(input m.CreateDto, userId uint64) (any, error) {
	var data m.UndanganPemeriksaan
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		respUndanganPemeriksaan, err := Create(input, userId, tx)
		if err != nil {
			return err
		}
		data = respUndanganPemeriksaan.(rp.OKSimple).Data.(m.UndanganPemeriksaan)
		inputDetails := []muppemeriksa.CreateDto{}
		for _, v := range input.PemeriksaDetail {
			inputDetails = append(inputDetails, muppemeriksa.CreateDto{UndanganPemeriksaan_Id: data.Id, Petugas_Id: v})
		}
		respUndanganPemeriksaanPemeriksa, err := suppemeriksa.Create(inputDetails, tx)
		if err != nil {
			return err
		}
		dataUndanganPemeriksaanPemeriksa := respUndanganPemeriksaanPemeriksa.(rp.OKSimple).Data.([]muppemeriksa.UndanganPemeriksaanPemeriksa)
		data.Pemeriksa = &dataUndanganPemeriksaanPemeriksa

		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal membuat data: "+err.Error(), data)
	}
	return rp.OKSimple{Data: data}, nil
}

func TrxUpdate(id int, input m.UpdateDto, userId uint64) (any, error) {
	var data m.UndanganPemeriksaan
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		respUndanganPemeriksaan, err := Update(id, userId, input, tx)
		if err != nil {
			return err
		}
		data = respUndanganPemeriksaan.(rp.OK).Data.(m.UndanganPemeriksaan)

		for i := range input.PemeriksaDetail {
			input.PemeriksaDetail[i].UndanganPemeriksaan_Id = data.Id
		}

		respPemeriksaDetail, err := suppemeriksa.Update(input.PemeriksaDetail, tx)
		if err != nil {
			return err
		}
		dataPemeriksaDetail := respPemeriksaDetail.(rp.OK).Data.([]muppemeriksa.UndanganPemeriksaanPemeriksa)
		data.Pemeriksa = &dataPemeriksaDetail

		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengubah data: "+err.Error(), data)
	}
	return rp.OKSimple{Data: data}, nil
}
