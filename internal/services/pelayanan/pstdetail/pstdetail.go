package pstdetail

import (
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "pst detail"

func GetByNoPelayanan(input m.GetByNoPelayananPstDetailDto) (any, error) {
	var data []m.PstDetail

	result := a.DB.
		Model(m.PstDetail{}).
		Where("TahunPelayanan", input.TahunPelayanan).
		Where("BundelPelayanan", input.BundelPelayanan).
		Where("NoUrutPelayanan", input.NoUrutPelayanan).
		Find(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data reklas", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func GetByNop(input m.GetByNop) (any, error) {
	var data []m.GetByNopResponse

	result := a.DB.
		Model(m.PstDetail{}).
		Where("PermohonanProvinsiID", input.PermohonanProvinsiID).
		Where("PermohonanKotaID", input.PermohonanKotaID).
		Where("PermohonanKecamatanID", input.PermohonanKecamatanID).
		Where("PermohonanKelurahanID", input.PermohonanKelurahanID).
		Where("PermohonanBlokID", input.PermohonanBlokID).
		Where("NoUrutPemohon", input.NoUrutPemohon).
		Where("PemohonJenisOPID", input.PemohonJenisOPID).
		Find(&data)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data reklas", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}
