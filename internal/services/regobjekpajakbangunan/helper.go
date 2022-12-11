package regobjekpajakbangunan

import (
	"fmt"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbangunan"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
)

func jpbCopier(dto interface{}, resultNop []string, kode string) (any, error) {
	switch v := dto.(type) {
	case m.OpbJpb2CreateDto:
		var data m.RegJpb2
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg jpb2", data)
		}
		data.NopDetail.Provinsi_Kode = &resultNop[0]
		data.NopDetail.Daerah_Kode = &resultNop[1]
		data.NopDetail.Kecamatan_Kode = &resultNop[2]
		data.NopDetail.Kelurahan_Kode = &resultNop[3]
		data.NopDetail.Blok_Kode = &resultNop[4]
		data.NopDetail.NoUrut = &resultNop[5]
		data.NopDetail.JenisOp = &resultNop[6]
		data.NopDetail.Area_Kode = &kode
		return data, nil
	case m.OpbJpb3CreateDto:
		var data m.RegJpb3
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg jpb3", data)
		}
		data.NopDetail.Provinsi_Kode = &resultNop[0]
		data.NopDetail.Daerah_Kode = &resultNop[1]
		data.NopDetail.Kecamatan_Kode = &resultNop[2]
		data.NopDetail.Kelurahan_Kode = &resultNop[3]
		data.NopDetail.Blok_Kode = &resultNop[4]
		data.NopDetail.NoUrut = &resultNop[5]
		data.NopDetail.JenisOp = &resultNop[6]
		data.NopDetail.Area_Kode = &kode
		return data, nil
	case m.OpbJpb4CreateDto:
		var data m.RegJpb4
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg jpb4", data)
		}
		data.NopDetail.Provinsi_Kode = &resultNop[0]
		data.NopDetail.Daerah_Kode = &resultNop[1]
		data.NopDetail.Kecamatan_Kode = &resultNop[2]
		data.NopDetail.Kelurahan_Kode = &resultNop[3]
		data.NopDetail.Blok_Kode = &resultNop[4]
		data.NopDetail.NoUrut = &resultNop[5]
		data.NopDetail.JenisOp = &resultNop[6]
		data.NopDetail.Area_Kode = &kode
		return data, nil
	case m.OpbJpb5CreateDto:
		var data m.RegJpb5
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg jpb5", data)
		}
		data.NopDetail.Provinsi_Kode = &resultNop[0]
		data.NopDetail.Daerah_Kode = &resultNop[1]
		data.NopDetail.Kecamatan_Kode = &resultNop[2]
		data.NopDetail.Kelurahan_Kode = &resultNop[3]
		data.NopDetail.Blok_Kode = &resultNop[4]
		data.NopDetail.NoUrut = &resultNop[5]
		data.NopDetail.JenisOp = &resultNop[6]
		data.NopDetail.Area_Kode = &kode
		return data, nil
	case m.OpbJpb6CreateDto:
		var data m.RegJpb6
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg jpb6", data)
		}
		data.NopDetail.Provinsi_Kode = &resultNop[0]
		data.NopDetail.Daerah_Kode = &resultNop[1]
		data.NopDetail.Kecamatan_Kode = &resultNop[2]
		data.NopDetail.Kelurahan_Kode = &resultNop[3]
		data.NopDetail.Blok_Kode = &resultNop[4]
		data.NopDetail.NoUrut = &resultNop[5]
		data.NopDetail.JenisOp = &resultNop[6]
		data.NopDetail.Area_Kode = &kode
		return data, nil
	case m.OpbJpb7CreateDto:
		var data m.RegJpb7
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg jpb7", data)
		}
		data.NopDetail.Provinsi_Kode = &resultNop[0]
		data.NopDetail.Daerah_Kode = &resultNop[1]
		data.NopDetail.Kecamatan_Kode = &resultNop[2]
		data.NopDetail.Kelurahan_Kode = &resultNop[3]
		data.NopDetail.Blok_Kode = &resultNop[4]
		data.NopDetail.NoUrut = &resultNop[5]
		data.NopDetail.JenisOp = &resultNop[6]
		data.NopDetail.Area_Kode = &kode
		return data, nil
	case m.OpbJpb8CreateDto:
		var data m.RegJpb8
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg jpb8", data)
		}
		data.NopDetail.Provinsi_Kode = &resultNop[0]
		data.NopDetail.Daerah_Kode = &resultNop[1]
		data.NopDetail.Kecamatan_Kode = &resultNop[2]
		data.NopDetail.Kelurahan_Kode = &resultNop[3]
		data.NopDetail.Blok_Kode = &resultNop[4]
		data.NopDetail.NoUrut = &resultNop[5]
		data.NopDetail.JenisOp = &resultNop[6]
		data.NopDetail.Area_Kode = &kode
		return data, nil
	case m.OpbJpb9CreateDto:
		var data m.RegJpb9
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg jpb9", data)
		}
		data.NopDetail.Provinsi_Kode = &resultNop[0]
		data.NopDetail.Daerah_Kode = &resultNop[1]
		data.NopDetail.Kecamatan_Kode = &resultNop[2]
		data.NopDetail.Kelurahan_Kode = &resultNop[3]
		data.NopDetail.Blok_Kode = &resultNop[4]
		data.NopDetail.NoUrut = &resultNop[5]
		data.NopDetail.JenisOp = &resultNop[6]
		data.NopDetail.Area_Kode = &kode
		return data, nil
	case m.OpbJpb12CreateDto:
		var data m.RegJpb12
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg jpb12", data)
		}
		data.NopDetail.Provinsi_Kode = &resultNop[0]
		data.NopDetail.Daerah_Kode = &resultNop[1]
		data.NopDetail.Kecamatan_Kode = &resultNop[2]
		data.NopDetail.Kelurahan_Kode = &resultNop[3]
		data.NopDetail.Blok_Kode = &resultNop[4]
		data.NopDetail.NoUrut = &resultNop[5]
		data.NopDetail.JenisOp = &resultNop[6]
		data.NopDetail.Area_Kode = &kode
		return data, nil
	case m.OpbJpb13CreateDto:
		var data m.RegJpb13
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg jpb13", data)
		}
		data.NopDetail.Provinsi_Kode = &resultNop[0]
		data.NopDetail.Daerah_Kode = &resultNop[1]
		data.NopDetail.Kecamatan_Kode = &resultNop[2]
		data.NopDetail.Kelurahan_Kode = &resultNop[3]
		data.NopDetail.Blok_Kode = &resultNop[4]
		data.NopDetail.NoUrut = &resultNop[5]
		data.NopDetail.JenisOp = &resultNop[6]
		data.NopDetail.Area_Kode = &kode
		return data, nil
	case m.OpbJpb14CreateDto:
		var data m.RegJpb14
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg jpb14", data)
		}
		data.NopDetail.Provinsi_Kode = &resultNop[0]
		data.NopDetail.Daerah_Kode = &resultNop[1]
		data.NopDetail.Kecamatan_Kode = &resultNop[2]
		data.NopDetail.Kelurahan_Kode = &resultNop[3]
		data.NopDetail.Blok_Kode = &resultNop[4]
		data.NopDetail.NoUrut = &resultNop[5]
		data.NopDetail.JenisOp = &resultNop[6]
		data.NopDetail.Area_Kode = &kode
		return data, nil
	case m.OpbJpb15CreateDto:
		var data m.RegJpb15
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg jpb15", data)
		}
		data.NopDetail.Provinsi_Kode = &resultNop[0]
		data.NopDetail.Daerah_Kode = &resultNop[1]
		data.NopDetail.Kecamatan_Kode = &resultNop[2]
		data.NopDetail.Kelurahan_Kode = &resultNop[3]
		data.NopDetail.Blok_Kode = &resultNop[4]
		data.NopDetail.NoUrut = &resultNop[5]
		data.NopDetail.JenisOp = &resultNop[6]
		data.NopDetail.Area_Kode = &kode
		return data, nil
	case m.OpbJpb16CreateDto:
		var data m.RegJpb16
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg jpb16", data)
		}
		data.NopDetail.Provinsi_Kode = &resultNop[0]
		data.NopDetail.Daerah_Kode = &resultNop[1]
		data.NopDetail.Kecamatan_Kode = &resultNop[2]
		data.NopDetail.Kelurahan_Kode = &resultNop[3]
		data.NopDetail.Blok_Kode = &resultNop[4]
		data.NopDetail.NoUrut = &resultNop[5]
		data.NopDetail.JenisOp = &resultNop[6]
		data.NopDetail.Area_Kode = &kode
		return data, nil
	default:
		return nil, fmt.Errorf("type of dto undefined")

	}
}
