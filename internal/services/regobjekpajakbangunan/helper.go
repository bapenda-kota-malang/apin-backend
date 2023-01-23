package regobjekpajakbangunan

import (
	"fmt"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	mopb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbangunan"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func jpbCopier(dto interface{}, resultNop []string, kode string, tx *gorm.DB) (any, error) {
	switch v := dto.(type) {
	case *m.RegOpbJpb2CreateDto:
		var data m.RegJpb2
		if err := sc.Copy(&data, &v.RegJpbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb2", data)
		}

		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb2", data)
		}

		return rp.OKSimple{Data: data}, nil
	case *m.RegOpbJpb3CreateDto:
		var data m.RegJpb3
		if err := sc.Copy(&data, &v.RegJpbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb3", data)
		}
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb3", data)
		}

		return rp.OKSimple{Data: data}, nil
	case *m.RegOpbJpb4CreateDto:
		var data m.RegJpb4
		if err := sc.Copy(&data, &v.RegJpbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb4", data)
		}
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb4", data)
		}

		return rp.OKSimple{Data: data}, nil
	case *m.RegOpbJpb5CreateDto:
		var data m.RegJpb5
		if err := sc.Copy(&data, &v.RegJpbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb5", data)
		}
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb5", data)
		}

		return rp.OKSimple{Data: data}, nil
	case *m.RegOpbJpb6CreateDto:
		var data m.RegJpb6
		if err := sc.Copy(&data, &v.RegJpbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb6", data)
		}
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb6", data)
		}

		return rp.OKSimple{Data: data}, nil
	case *m.RegOpbJpb7CreateDto:
		var data m.RegJpb7
		if err := sc.Copy(&data, &v.RegJpbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb7", data)
		}
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb7", data)
		}

		return rp.OKSimple{Data: data}, nil
	case *m.RegOpbJpb8CreateDto:
		var data m.RegJpb8
		if err := sc.Copy(&data, &v.RegJpbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb8", data)
		}
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb8", data)
		}

		return rp.OKSimple{Data: data}, nil
	case *m.RegOpbJpb9CreateDto:
		var data m.RegJpb9
		if err := sc.Copy(&data, &v.RegJpbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb9", data)
		}
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb9", data)
		}

		return rp.OKSimple{Data: data}, nil
	case *m.RegOpbJpb12CreateDto:
		var data m.RegJpb12
		if err := sc.Copy(&data, &v.RegJpbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb12", data)
		}
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb12", data)
		}

		return rp.OKSimple{Data: data}, nil
	case *m.RegOpbJpb13CreateDto:
		var data m.RegJpb13
		if err := sc.Copy(&data, &v.RegJpbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb13", data)
		}
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb13", data)
		}

		return rp.OKSimple{Data: data}, nil
	case *m.RegOpbJpb14CreateDto:
		var data m.RegJpb14
		if err := sc.Copy(&data, &v.RegJpbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb14", data)
		}
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb14", data)
		}

		return rp.OKSimple{Data: data}, nil
	case *m.RegOpbJpb15CreateDto:
		var data m.RegJpb15
		if err := sc.Copy(&data, &v.RegJpbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb15", data)
		}
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb15", data)
		}

		return rp.OKSimple{Data: data}, nil
	case *m.RegOpbJpb16CreateDto:
		var data m.RegJpb16
		if err := sc.Copy(&data, &v.RegJpbs); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb16", data)
		}
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb16", data)
		}

		return rp.OKSimple{Data: data}, nil
	default:
		return nil, fmt.Errorf("type of dto undefined")

	}
}

func verifyJpb(jpb_kode string, nop nop.NopDetail, tx *gorm.DB) (any, error) {
	switch jpb_kode {
	case "02":
		var dataReg m.RegJpb2

		result := tx.Where(dataReg).First(&dataReg)
		if result.RowsAffected == 0 {
			return sh.SetError("request", "create-data", source, "failed", "data jpb2 tidak ditemukan", dataReg)
		}

		var data mopb.Jpb2
		if err := sc.Copy(&data, &dataReg); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb2", dataReg)
		}
		data.Id = 0
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb2", data)

		}
		return rp.OKSimple{Data: data}, nil
	case "03":
		var dataReg m.RegJpb3

		result := tx.Where(dataReg).First(&dataReg)
		if result.RowsAffected == 0 {
			return sh.SetError("request", "create-data", source, "failed", "data jpb3 tidak ditemukan", dataReg)
		}

		var data mopb.Jpb3
		if err := sc.Copy(&data, &dataReg); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb3", dataReg)
		}
		data.Id = 0
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb3", data)

		}
		return rp.OKSimple{Data: data}, nil
	case "04":
		var dataReg m.RegJpb4

		result := tx.Where(dataReg).First(&dataReg)
		if result.RowsAffected == 0 {
			return sh.SetError("request", "create-data", source, "failed", "data jpb4 tidak ditemukan", dataReg)
		}

		var data mopb.Jpb4
		if err := sc.Copy(&data, &dataReg); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb4", dataReg)
		}
		data.Id = 0
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb4", data)

		}
		return rp.OKSimple{Data: data}, nil
	case "05":
		var dataReg m.RegJpb5

		result := tx.Where(dataReg).First(&dataReg)
		if result.RowsAffected == 0 {
			return sh.SetError("request", "create-data", source, "failed", "data jpb15 tidak ditemukan", dataReg)
		}

		var data mopb.Jpb5
		if err := sc.Copy(&data, &dataReg); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb5", dataReg)
		}
		data.Id = 0
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb5", data)

		}
		return rp.OKSimple{Data: data}, nil
	case "06":
		var dataReg m.RegJpb6

		result := tx.Where(dataReg).First(&dataReg)
		if result.RowsAffected == 0 {
			return sh.SetError("request", "create-data", source, "failed", "data jpb6 tidak ditemukan", dataReg)
		}

		var data mopb.Jpb6
		if err := sc.Copy(&data, &dataReg); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb6", dataReg)
		}
		data.Id = 0
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb6", data)

		}
		return rp.OKSimple{Data: data}, nil
	case "07":
		var dataReg m.RegJpb7

		result := tx.Where(dataReg).First(&dataReg)
		if result.RowsAffected == 0 {
			return sh.SetError("request", "create-data", source, "failed", "data jpb7 tidak ditemukan", dataReg)
		}

		var data mopb.Jpb7
		if err := sc.Copy(&data, &dataReg); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb7", dataReg)
		}
		data.Id = 0
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb7", data)

		}
		return rp.OKSimple{Data: data}, nil
	case "08":
		var dataReg m.RegJpb8

		result := tx.Where(dataReg).First(&dataReg)
		if result.RowsAffected == 0 {
			return sh.SetError("request", "create-data", source, "failed", "data jpb8 tidak ditemukan", dataReg)
		}

		var data mopb.Jpb8
		if err := sc.Copy(&data, &dataReg); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb8", dataReg)
		}
		data.Id = 0
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb8", data)

		}
		return rp.OKSimple{Data: data}, nil
	case "09":
		var dataReg m.RegJpb9

		result := tx.Where(dataReg).First(&dataReg)
		if result.RowsAffected == 0 {
			return sh.SetError("request", "create-data", source, "failed", "data jpb9 tidak ditemukan", dataReg)
		}

		var data mopb.Jpb9
		if err := sc.Copy(&data, &dataReg); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb9", dataReg)
		}
		data.Id = 0
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb9", data)

		}
		return rp.OKSimple{Data: data}, nil
	case "12":
		var dataReg m.RegJpb12

		result := tx.Where(dataReg).First(&dataReg)
		if result.RowsAffected == 0 {
			return sh.SetError("request", "create-data", source, "failed", "data jpb12 tidak ditemukan", dataReg)
		}

		var data mopb.Jpb12
		if err := sc.Copy(&data, &dataReg); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb12", dataReg)
		}
		data.Id = 0
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb12", data)

		}
		return rp.OKSimple{Data: data}, nil
	case "13":
		var dataReg m.RegJpb13

		result := tx.Where(dataReg).First(&dataReg)
		if result.RowsAffected == 0 {
			return sh.SetError("request", "create-data", source, "failed", "data jpb13 tidak ditemukan", dataReg)
		}

		var data mopb.Jpb13
		if err := sc.Copy(&data, &dataReg); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb13", dataReg)
		}
		data.Id = 0
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb13", data)

		}
		return rp.OKSimple{Data: data}, nil
	case "14":
		var dataReg m.RegJpb14

		result := tx.Where(dataReg).First(&dataReg)
		if result.RowsAffected == 0 {
			return sh.SetError("request", "create-data", source, "failed", "data jpb14 tidak ditemukan", dataReg)
		}

		var data mopb.Jpb14
		if err := sc.Copy(&data, &dataReg); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb14", dataReg)
		}
		data.Id = 0
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb14", data)

		}
		return rp.OKSimple{Data: data}, nil
	case "15":
		var dataReg m.RegJpb15

		result := tx.Where(dataReg).First(&dataReg)
		if result.RowsAffected == 0 {
			return sh.SetError("request", "create-data", source, "failed", "data jpb15 tidak ditemukan", dataReg)
		}

		var data mopb.Jpb15
		if err := sc.Copy(&data, &dataReg); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb15", dataReg)
		}
		data.Id = 0
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb15", data)

		}
		return rp.OKSimple{Data: data}, nil

	case "16":
		var dataReg m.RegJpb16

		result := tx.Where(dataReg).First(&dataReg)
		if result.RowsAffected == 0 {
			return sh.SetError("request", "create-data", source, "failed", "data jpb16 tidak ditemukan", dataReg)
		}

		var data mopb.Jpb16
		if err := sc.Copy(&data, &dataReg); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload jpb16", dataReg)
		}
		data.Id = 0
		if result := tx.Create(&data); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data jpb16", data)

		}
		return rp.OKSimple{Data: data}, nil

	default:

		return nil, fmt.Errorf("jpb kode %s tidak diketahui", jpb_kode)
	}
}
