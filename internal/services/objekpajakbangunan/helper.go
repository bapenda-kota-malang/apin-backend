package objekpajakbangunan

import (
	"fmt"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
)

func jpbCopier(dto interface{}, resultNop []string, kode string) (any, error) {
	switch v := dto.(type) {
	case m.OpbJpb2CreateDto:
		var data m.Jpb2
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb2", data)
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
		var data m.Jpb3
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb3", data)
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
		var data m.Jpb4
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb4", data)
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
		var data m.Jpb5
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb5", data)
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
		var data m.Jpb6
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb6", data)
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
		var data m.Jpb7
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb7", data)
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
		var data m.Jpb8
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb8", data)
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
		var data m.Jpb9
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb9", data)
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
		var data m.Jpb12
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb12", data)
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
		var data m.Jpb13
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb13", data)
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
		var data m.Jpb14
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb14", data)
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
		var data m.Jpb15
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb15", data)
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
		var data m.Jpb16
		if err := sc.Copy(&data, &v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb16", data)
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
