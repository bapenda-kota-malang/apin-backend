package sppt

import (
	"errors"

	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	opbg "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	opb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	oppbb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	mspptp "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt/pembayaran"
	mtp "github.com/bapenda-kota-malang/apin-backend/internal/models/tempatpembayaran"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
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

func getObjekPajakBumiDetail(data m.Sppt) (*opb.ObjekPajakBumi, error) {
	var opBumiData *opb.ObjekPajakBumi
	fBumi := opb.ObjekPajakBumi{
		NopDetail: nop.NopDetail{
			Provinsi_Kode:  data.Propinsi_Id,
			Daerah_Kode:    data.Dati2_Id,
			Kecamatan_Kode: data.Kecamatan_Id,
			Kelurahan_Kode: data.Keluarahan_Id,
			Blok_Kode:      data.Blok_Id,
			NoUrut:         data.NoUrut,
		},
	}
	if err := a.DB.Where(&fBumi).Order("\"Id\" desc").First(&opBumiData).Error; err != nil {
		return nil, err
	}
	return opBumiData, nil
}

func getObjekPajakBangunanDetail(data m.Sppt) (*opbg.ObjekPajakBangunan, error) {
	var opBngData *opbg.ObjekPajakBangunan

	fBng := opbg.ObjekPajakBangunan{
		NopDetail: nop.NopDetail{
			Provinsi_Kode:  data.Propinsi_Id,
			Daerah_Kode:    data.Dati2_Id,
			Kecamatan_Kode: data.Kecamatan_Id,
			Kelurahan_Kode: data.Keluarahan_Id,
			Blok_Kode:      data.Blok_Id,
			NoUrut:         data.NoUrut,
		},
	}
	if err := a.DB.Where(&fBng).Order("\"Id\" desc").First(&opBngData).Error; err != nil {
		return nil, err
	}
	return opBngData, nil
}

func getObjekPajakPBBDetail(data m.Sppt) (*oppbb.ObjekPajakPbb, error) {
	var result *oppbb.ObjekPajakPbb
	kelurahan := *data.Propinsi_Id + *data.Dati2_Id + *data.Kecamatan_Id + *data.Keluarahan_Id

	filter := oppbb.ObjekPajakPbb{
		NopDetail: nop.NopDetail{
			Provinsi_Kode:  data.Propinsi_Id,
			Daerah_Kode:    data.Dati2_Id,
			Kecamatan_Kode: data.Kecamatan_Id,
			Kelurahan_Kode: &kelurahan,
			Blok_Kode:      data.Blok_Id,
			NoUrut:         data.NoUrut,
		},
	}
	if err := a.DB.Where(&filter).Order("\"Id\" desc").First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func getPembayaranSppt(data m.Sppt) (*mspptp.SpptPembayaran, error) {
	var result *mspptp.SpptPembayaran

	filter := mspptp.SpptPembayaran{
		Provinsi_Kd:    *data.Propinsi_Id,
		Daerah_Kd:      *data.Dati2_Id,
		Kecamatan_Kd:   *data.Kecamatan_Id,
		Kelurahan_Kd:   *data.Keluarahan_Id,
		Blok_Kd:        *data.Blok_Id,
		NoUrut_Kd:      *data.NoUrut,
		JenisOp_Kd:     *data.JenisOP_Id,
		TahunPajakSppt: *data.TahunPajakskp_sppt,
	}
	if err := a.DB.Where(&filter).Order("\"Id\" desc").First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func getTempatPembayaranSppt(data m.Sppt) (*mtp.TempatPembayaran, error) {
	var result *mtp.TempatPembayaran

	filter := mtp.TempatPembayaran{
		BankPersepsi_Id: data.BankPersepsi_Id,
		BankTunggal_Id:  data.BankTunggal_Id,
		Kanwil_Id:       data.KanwilBank_Id,
		Kppbb_Id:        data.KPPBBbank_Id,
	}
	if err := a.DB.Where(&filter).Order("\"Id\" desc").First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
