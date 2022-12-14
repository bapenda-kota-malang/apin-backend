package objekpajakbumi

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	msppt "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
)

func nopSearcher(input msppt.Sppt) m.ObjekPajakBumi {
	return m.ObjekPajakBumi{
		NopDetail: nop.NopDetail{
			Provinsi_Kode:  input.Propinsi_Id,
			Daerah_Kode:    input.Dati2_Id,
			Kecamatan_Kode: input.Kecamatan_Id,
			Kelurahan_Kode: input.Keluarahan_Id,
			Blok_Kode:      input.Blok_Id,
			NoUrut:         input.NoUrut,
			JenisOp:        input.JenisOP_Id,
		}}

}
