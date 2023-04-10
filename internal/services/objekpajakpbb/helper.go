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

func GetByNopObject(input m.ObjekPajakPbb) (*m.ObjekPajakPbb, error) {
	var data *m.ObjekPajakPbb

	kelurahan := *input.NopDetail.Provinsi_Kode + *input.NopDetail.Daerah_Kode + *input.NopDetail.Kecamatan_Kode + *input.NopDetail.Kelurahan_Kode

	filter := m.ObjekPajakPbb{
		NopDetail: nop.NopDetail{
			Provinsi_Kode:  input.NopDetail.Provinsi_Kode,
			Daerah_Kode:    input.NopDetail.Daerah_Kode,
			Kecamatan_Kode: input.NopDetail.Kecamatan_Kode,
			Kelurahan_Kode: &kelurahan,
			Blok_Kode:      input.NopDetail.Blok_Kode,
			NoUrut:         input.NopDetail.NoUrut,
		},
	}
	result := a.DB.Debug().Where(&filter).First(&data)
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
