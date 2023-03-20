package nilaiindividu

import (
	// nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/nilaiindividu"
	n "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
)

func nopSearcher(input m.NilaiIndividu) m.NilaiIndividu {
	return m.NilaiIndividu{
		NopDetail: n.NopDetail{
			Provinsi_Kode:  input.Provinsi_Kode,
			Daerah_Kode:    input.Daerah_Kode,
			Kecamatan_Kode: input.Kecamatan_Kode,
			Kelurahan_Kode: input.Kelurahan_Kode,
			Blok_Kode:      input.Blok_Kode,
			NoUrut:         input.NoUrut,
			JenisOp:        input.JenisOp,
		},
	}

}
