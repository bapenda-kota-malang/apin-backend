package sppt

import (
	"errors"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
)

func nopCondition(input m.PenilaianDto) (m.Sppt, error) {
	if input.Kelurahan_Kode != nil {
		tmp := m.Sppt{
			Propinsi_Id:        input.Provinsi_Kode,
			Dati2_Id:           input.Daerah_Kode,
			Kecamatan_Id:       input.Kecamatan_Kode,
			Keluarahan_Id:      input.Kelurahan_Kode,
			TahunPajakskp_sppt: &input.Tahun,
		}

		return tmp, nil
	} else if input.Kelurahan_Kode == nil {
		tmp := m.Sppt{
			Propinsi_Id:        input.Provinsi_Kode,
			Dati2_Id:           input.Daerah_Kode,
			Kecamatan_Id:       input.Kecamatan_Kode,
			TahunPajakskp_sppt: &input.Tahun,
		}

		return tmp, nil
	} else if input.Kecamatan_Kode == nil {
		tmp := m.Sppt{
			Propinsi_Id:        input.Provinsi_Kode,
			Dati2_Id:           input.Daerah_Kode,
			TahunPajakskp_sppt: &input.Tahun,
		}
		return tmp, nil
	} else if input.Daerah_Kode == nil {
		tmp := m.Sppt{
			Propinsi_Id:        input.Provinsi_Kode,
			TahunPajakskp_sppt: &input.Tahun,
		}
		return tmp, nil
	}

	return m.Sppt{}, errors.New("nop tidak ditemukan")
}
