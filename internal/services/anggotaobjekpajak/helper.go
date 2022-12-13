package anggotaobjekpajak

import (
	// nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/anggotaobjekpajak"
	msppt "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
)

func nopSearcher(input msppt.Sppt) m.AnggotaObjekPajak {
	return m.AnggotaObjekPajak{
		Provinsi_Kode:  input.Propinsi_Id,
		Daerah_Kode:    input.Dati2_Id,
		Kecamatan_Kode: input.Kecamatan_Id,
		Kelurahan_Kode: input.Keluarahan_Id,
		Blok_Kode:      input.Blok_Id,
		NoUrut:         input.NoUrut,
		JenisOp:        input.JenisOP_Id,
	}

}
