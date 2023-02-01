package objekpajakbumi

import (
	"strings"

	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	pel "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
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

func transformNOP(inputNop string) string {
	result := strings.Trim(inputNop, ".")
	return result
}

func DecodeNOP(inputNop *string) *pel.PermohonanNOP {
	if inputNop != nil {
		var tempNOP string
		tempNOP = *inputNop
		result := pel.PermohonanNOP{
			PermohonanProvinsiID:  tempNOP[0:2],
			PermohonanKotaID:      tempNOP[2:4],
			PermohonanKecamatanID: tempNOP[4:7],
			PermohonanKelurahanID: tempNOP[7:10],
			PermohonanBlokID:      tempNOP[10:13],
			NoUrutPemohon:         tempNOP[13:17],
			PemohonJenisOPID:      tempNOP[17:18],
		}
		return &result
	}
	return nil
}
