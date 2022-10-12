package registrasinpwpd

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	nm "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/registrasinpwpd"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "registrasiNpwpd"

func init() {
	a.AutoMigrate(&rn.RegistrasiNpwpd{})
}

func insertDetailOp(objek string, data *[]rn.DetailRegOp, registerForm *rn.RegistrasiNpwpd) error {
	var (
		err      error
		mActions map[string]reflect.Type = make(map[string]reflect.Type)
		model    interface{}

		detailRegOpHotel    rn.DetailRegOpHotel
		detailRegOpResto    rn.DetailRegOpResto
		detailRegOpHiburan  rn.DetailRegOpHiburan
		detailRegOpReklame  rn.DetailRegOpReklame
		detailRegOpPpj      rn.DetailRegOpPpj
		detailRegOpParkir   rn.DetailRegOpParkir
		detailRegOpAirTanah rn.DetailRegOpAirTanah
	)

	mActions[`detailRegOpHotel`] = reflect.TypeOf(detailRegOpHotel)
	mActions[`detailRegOpResto`] = reflect.TypeOf(detailRegOpResto)
	mActions[`detailRegOpHiburan`] = reflect.TypeOf(detailRegOpHiburan)
	mActions[`detailRegOpReklame`] = reflect.TypeOf(detailRegOpReklame)
	mActions[`detailRegOpPpj`] = reflect.TypeOf(detailRegOpPpj)
	mActions[`detailRegOpParkir`] = reflect.TypeOf(detailRegOpParkir)
	mActions[`detailRegOpAirTanah`] = reflect.TypeOf(detailRegOpAirTanah)

	switch objek {
	case "01":
		model = reflect.Zero(mActions["detailRegOpHotel"]).Interface()
	case "02":
		model = reflect.Zero(mActions["detailRegOpResto"]).Interface()
	case "03":
		model = reflect.Zero(mActions["detailRegOpHiburan"]).Interface()
	case "04":
		model = reflect.Zero(mActions["detailRegOpReklame"]).Interface()
	case "05":
		model = reflect.Zero(mActions["detailRegOpPpj"]).Interface()
	case "06":
		model = reflect.Zero(mActions["detailRegOpParkir"]).Interface()
	case "07":
		model = reflect.Zero(mActions["detailRegOpAirTanah"]).Interface()
	}

	for _, dop := range *data {
		dop.JenisOp = registerForm.Rekening.Nama
		dop.RegistrasiNpwpd_Id = registerForm.Id
		// dop.Pendaftaran_Npwpd = registerForm.Npwpd

		m := make(map[string]interface{})
		m["RegistrasiNpwpd_Id"] = dop.RegistrasiNpwpd_Id
		// m["Pendaftaran_Npwpd"] = dop.Pendaftaran_Npwpd
		m["JenisOp"] = dop.JenisOp
		m["JumlahOp"] = dop.JumlahOp
		m["TarifOp"] = dop.TarifOp
		m["UnitOp"] = dop.UnitOp
		m["Notes"] = dop.Notes

		err = a.DB.Model(&model).Create(&m).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func Create(reg rn.CreateDto, user_Id uint) (interface{}, error) {
	user_IdConv := uint64(user_Id)
	var rekening *rm.Rekening
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, reg.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	var tmpverify = rn.VerifyStatusBaru
	var nomorRegistrasiNpwpd = func() *uint64 {

		if reg.IsNomorRegistrasiAuto {
			var tmp uint64
			var tmpNpwpd rn.RegistrasiNpwpd
			// row := a.DB.Table("Npwpd").Select("max(Nomor)").Row()
			// row.Scan(&tmp)
			nomor := a.DB.Last(&tmpNpwpd)
			if nomor.Error != nil {
				return nil
			}
			tmp = *tmpNpwpd.Nomor
			fmt.Println("nomor: ", tmp)
			tmp += uint64(1)
			return &tmp
		}
		return reg.Nomor
	}()

	//ubah nomor ke string untuk pembuatan npwpd
	// nomorString := strconv.Itoa(int(*nomorRegistrasiNpwpd))
	// kecamatanIdString := strconv.Itoa(int(*reg.RegObjekPajak.Kecamatan_Id))
	// kodeJenisUsahaString := *rekening.KodeJenisUsaha
	// fmt.Println("data rekening: ", *rekening.Nama)
	// fmt.Println("kodejenisusaha: ", kodeJenisUsahaString)
	// if kodeJenisUsahaString == "" {
	// 	kodeJenisUsahaString = "xxxxx"
	// }
	// npwpdString := nomorString + "." + kecamatanIdString + "." + kodeJenisUsahaString[:3]
	register := rn.RegistrasiNpwpd{
		// ModeRegistrasi: npwpd.ModeOperator,
		Status:       nt.StatusAktif,
		JenisPajak:   reg.JenisPajak,
		User_Id:      &user_IdConv,
		Golongan:     reg.Golongan,
		Npwp:         reg.Npwp,
		VerifyStatus: &tmpverify,
		Nomor:        nomorRegistrasiNpwpd,
		// Npwpd:        &npwpdString,
		// TanggalPengukuhan: func() *time.Time {
		// 	t, err := time.Parse("2006-01-02", *reg.TanggalPengukuhan)
		// 	if err != nil {
		// 		return nil
		// 	}
		// 	return &t
		// }(),
		// TanggalNpwpd: func() *time.Time {
		// 	t, err := time.Parse("2006-01-02", *reg.TanggalNpwpd)
		// 	if err != nil {
		// 		return nil
		// 	}
		// 	return &t
		// }(),
		Rekening_Id:       reg.Rekening_Id,
		Rekening:          rekening,
		TanggalMulaiUsaha: th.ParseTime(*reg.TanggalMulaiUsaha),
		LuasBangunan:      reg.LuasBangunan,
		JamBukaUsaha:      reg.JamBukaUsaha,
		JamTutupUsaha:     reg.JamTutupUsaha,
		Pengunjung:        reg.Pengunjung,
		OmsetOp:           reg.OmsetOp,
		Genset:            &reg.Genset,
		AirTanah:          &reg.AirTanah,
	}
	err = a.DB.Create(&register).Error
	if err != nil {
		return nil, err
	}
	err = insertDetailOp(*rekening.Objek, reg.DetailRegOp, &register)
	if err != nil {
		return nil, err
	}
	// objekpajak
	op := *reg.RegObjekPajak
	op.RegistrasiNpwpd_Id = register.Id
	op.Status = nt.StatusBaru
	errOp := a.DB.Create(&op).Error
	if errOp != nil {
		return nil, err
	}

	for _, p := range *reg.RegPemilik {
		p.RegistrasiNpwpd_Id = register.Id
		p.Status = nt.StatusBaru
		err := a.DB.Create(&p).Error
		if err != nil {
			return nil, err
		}
	}
	for _, n := range *reg.RegNarahubung {
		n.RegistrasiNpwpd_Id = register.Id
		n.Status = nt.StatusBaru
		err := a.DB.Create(&n).Error
		if err != nil {
			return nil, err
		}
	}
	return rp.OKSimple{
		Data: register,
	}, nil
}

func GetAll(pagination gormhelper.Pagination) (interface{}, error) {
	var (
		register []*rn.RegistrasiNpwpd
		count    int64
	)

	result := a.DB.Model(&rn.RegistrasiNpwpd{}).
		Preload(clause.Associations).
		//nested preload
		//Preload("PemilikWps.Kelurahan").
		Count(&count).
		// Scopes(gormhelper.Paginate(&pagination)).
		Find(&register)

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: register,
	}, nil
}

func GetDetail(r *http.Request, regID int) (interface{}, error) {
	var register *rn.RegistrasiNpwpd
	err := a.DB.Model(&rn.RegistrasiNpwpd{}).
		Preload(clause.Associations).First(&register, regID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return rp.OKSimple{
		Data: register,
	}, err
}

func VerifyNpwpd(id int, input rn.VerifikasiDto) (any, error) {
	var data *rn.RegistrasiNpwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	fmt.Println("verify kode: ", *data.VerifyStatus)
	if *data.VerifyStatus != rn.VerifyStatusBaru {
		if *data.VerifyStatus == rn.VerifyStatusDisetujui {
			return nil, errors.New("data sudah disetujui")
		}
		return nil, errors.New("data sudah ditolak")
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	var tmp = time.Now()
	data.VerifiedAt = &tmp
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	var dataNpwpd nm.Npwpd
	if err := sc.Copy(&dataNpwpd, &data); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	//rekening
	var rekening *rm.Rekening
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	// kecamatan_id from regobjekpajak
	var dataRegOp *rn.RegObjekPajak
	err = a.DB.Where(rn.RegObjekPajak{RegistrasiNpwpd_Id: uint64(id)}).First(&dataRegOp).Error
	if err != nil {
		return nil, err
	}

	//creating npwpd
	nomorString := strconv.Itoa(int(*data.Nomor))
	kecamatanIdString := strconv.Itoa(int(*dataRegOp.Kecamatan_Id))
	kodeJenisUsahaString := *rekening.KodeJenisUsaha
	fmt.Println("data rekening: ", *rekening.Nama)
	fmt.Println("kodejenisusaha: ", kodeJenisUsahaString)
	if kodeJenisUsahaString == "" {
		kodeJenisUsahaString = "xxx"
	}
	npwpdString := nomorString + "." + kecamatanIdString + "." + kodeJenisUsahaString
	dataNpwpd.Npwpd = &npwpdString

	//tanggal
	dataNpwpd.TanggalPengukuhan = th.TimeNow()
	dataNpwpd.TanggalNpwpd = th.TimeNow()
	err = a.DB.Create(&dataNpwpd).Error
	if err != nil {
		return nil, err
	}

	//transfer detail from reg

	switch *rekening.Objek {
	case "01":
		var dataReg []*rn.DetailRegOpHotel
		result := a.DB.Where(rn.DetailRegOpHotel{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailOpHotel
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id

			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}

	case "02":
		var dataReg []*rn.DetailRegOpResto
		result := a.DB.Where(rn.DetailRegOpResto{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailOpResto
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id

			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}
	case "03":
		var dataReg []*rn.DetailRegOpHiburan
		result := a.DB.Where(rn.DetailRegOpHiburan{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailOpHiburan
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id

			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}
	case "04":
		var dataReg []*rn.DetailRegOpReklame
		result := a.DB.Where(rn.DetailRegOpReklame{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailOpReklame
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id

			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}
	case "05":
		var dataReg []*rn.DetailRegOpPpj
		result := a.DB.Where(rn.DetailRegOpPpj{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailOpPpj
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id

			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}
	case "06":
		var dataReg []*rn.DetailRegOpParkir
		result := a.DB.Where(rn.DetailRegOpParkir{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailOpParkir
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id

			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}
	case "07":
		var dataReg []*rn.DetailRegOpAirTanah
		result := a.DB.Where(rn.DetailRegOpAirTanah{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailOpAirTanah
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id

			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}
