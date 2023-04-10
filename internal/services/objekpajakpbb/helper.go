package objekpajakpbb

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	msppt "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
)

func nopSearcher(input msppt.Sppt) m.ObjekPajakPbb {
	return m.ObjekPajakPbb{
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

func GetByNopObject(input *m.ObjekPajakPbb) (m.ObjekPajakPbb, error) {
	var data m.ObjekPajakPbb

	kelurahan := *input.Provinsi_Kode + *input.Daerah_Kode + *input.Kecamatan_Kode + *input.Kelurahan_Kode

	filter := m.ObjekPajakPbb{
		NopDetail: nop.NopDetail{
			Provinsi_Kode:  data.Provinsi_Kode,
			Daerah_Kode:    data.Daerah_Kode,
			Kecamatan_Kode: data.Kecamatan_Kode,
			Kelurahan_Kode: &kelurahan,
			Blok_Kode:      data.Blok_Kode,
			NoUrut:         data.NoUrut,
		},
	}
	result := a.DB.Where(filter).First(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func GetPegawaiByNip(input *string) (pegawai.Pegawai, error) {
	var data pegawai.Pegawai

	filter := pegawai.FilterDto{
		Nip: input,
	}

	result := a.DB.Where(filter).First(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}
