package registrasinpwpd

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	nm "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/registrasinpwpd"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "registrasiNpwpd"

func Create(reg rn.CreateDto, user_Id uint) (interface{}, error) {
	user_IdConv := uint64(user_Id)
	var rekening *rm.Rekening
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, reg.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	var tmpverify = rn.VerifyStatusBaru
	// var tmpNomor = func() string {

	// 	if reg.IsNomorRegistrasiAuto {
	// 		var tmp string
	// 		var tmpNpwpd rn.RegistrasiNpwpd
	// 		nomor := a.DB.Last(&tmpNpwpd)
	// 		if nomor.Error != nil {
	// 			return "1000"
	// 		} else {
	// 			intconv, _ := strconv.Atoi(tmpNpwpd.Nomor)
	// 			intconv++
	// 			tmp = strconv.Itoa(intconv)
	// 		}
	// 		return tmp
	// 	}
	// 	return reg.Nomor
	// }()
	// foto ktp
	var imgNameChan = make(chan string)
	var errChan = make(chan error)

	go sh.SaveImage(reg.FotoKtp, imgNameChan, errChan)
	if err := <-errChan; err != nil {
		return sh.SetError("request", "create-data", source, "failed", "image unsupported", reg)
	}
	var tmp string = <-imgNameChan

	register := rn.RegistrasiNpwpd{
		// ModeRegistrasi: npwpd.ModeOperator,
		Status: nt.StatusAktif,
		// JenisPajak:        reg.JenisPajak,
		User_Id:      user_IdConv,
		Golongan:     reg.Golongan,
		Npwp:         reg.Npwp,
		VerifyStatus: &tmpverify,
		// Nomor:             tmpNomor,
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
		FotoKtp:           tmp,
		LainLain:          sh.GetArrayPhoto(reg.LainLain),
		SuratIzinUsaha:    sh.GetArrayPhoto(reg.SuratIzinUsaha),
		FotoObjek:         sh.GetArrayPhoto(reg.FotoObjek),
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

func GetList(input rn.FilterDto) (interface{}, error) {
	var data []rn.RegistrasiNpwpd
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&rn.RegistrasiNpwpd{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	// result := a.DB.Model(&rn.RegistrasiNpwpd{}).
	// 	Preload(clause.Associations).
	// 	//nested preload
	// 	//Preload("PemilikWps.Kelurahan").
	// 	Count(&count).
	// 	// Scopes(gormhelper.Paginate(&pagination)).
	// 	Find(&register)

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: data,
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
	if *data.VerifyStatus != rn.VerifyStatusBaru {
		if *data.VerifyStatus == rn.VerifyStatusDisetujui {
			return nil, errors.New("data sudah disetujui")
		}
		return nil, errors.New("data sudah ditolak")
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	data.VerifiedAt = th.TimeNow()
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
	// nomorString := strconv.Itoa(int(*data.Nomor))
	kecamatanIdString := strconv.Itoa(int(*dataRegOp.Kecamatan_Id))
	kodeJenisUsahaString := *rekening.KodeJenisUsaha
	fmt.Println("data rekening: ", *rekening.Nama)
	fmt.Println("kodejenisusaha: ", kodeJenisUsahaString)
	if kodeJenisUsahaString == "" {
		kodeJenisUsahaString = "xxx"
	}
	npwpdString := data.Nomor + "." + kecamatanIdString + "." + kodeJenisUsahaString
	dataNpwpd.Npwpd = &npwpdString

	//tanggal
	dataNpwpd.TanggalPengukuhan = th.TimeNow()
	dataNpwpd.TanggalNpwpd = th.TimeNow()
	dataNpwpd.Id = 0
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
			dataOp.Id = 0
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
			dataOp.Id = 0
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
			dataOp.Id = 0
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
			dataOp.Id = 0
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
			dataOp.Id = 0
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
			dataOp.Id = 0
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
			dataOp.Id = 0
			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}
	}

	var dataRP []*rn.RegPemilikWp
	result = a.DB.Where(rn.RegPemilikWp{RegistrasiNpwpd_Id: uint64(id)}).Find(&dataRP)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	for _, v := range dataRP {
		var dataP nm.PemilikWp
		if err := sc.Copy(&dataP, v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", v)
		}

		dataP.Npwpd_Id = dataNpwpd.Id
		dataP.Id = 0
		err = a.DB.Create(&dataP).Error
		if err != nil {
			return nil, err
		}
	}

	var dataRN []*rn.RegNarahubung
	result = a.DB.Where(rn.RegNarahubung{RegistrasiNpwpd_Id: uint64(id)}).Find(&dataRN)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	for _, v := range dataRN {
		var dataN nm.Narahubung
		if err := sc.Copy(&dataN, v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", v)
		}

		dataN.Npwpd_Id = dataNpwpd.Id
		dataN.Id = 0
		err = a.DB.Create(&dataN).Error
		if err != nil {
			return nil, err
		}
	}

	var dataROP *rn.RegObjekPajak
	result = a.DB.Where(rn.RegObjekPajak{RegistrasiNpwpd_Id: uint64(id)}).First(&dataROP)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	var dataObjekPajak nm.ObjekPajak
	if err := sc.Copy(&dataObjekPajak, dataROP); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataROP)
	}

	dataObjekPajak.Npwpd_Id = dataNpwpd.Id
	dataObjekPajak.Id = 0
	err = a.DB.Create(&dataObjekPajak).Error
	if err != nil {
		return nil, err
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (any, error) {

	var status string
	// data pemilikwp
	var dataPemilik []*rn.RegPemilikWp
	result := a.DB.Where(rn.RegPemilikWp{RegistrasiNpwpd_Id: uint64(id)}).Find(&dataPemilik)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	for _, v := range dataPemilik {

		result = a.DB.Where(rn.RegPemilikWp{RegistrasiNpwpd_Id: uint64(id)}).Delete(&v)
		status = "deleted"
		if result.RowsAffected == 0 {
			dataPemilik = nil
			status = "no deletion"
		}
	}

	// data narahubung
	var dataNarahubung []*rn.RegNarahubung
	result = a.DB.Where(rn.RegNarahubung{RegistrasiNpwpd_Id: uint64(id)}).Find(&dataNarahubung)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	for _, v := range dataNarahubung {
		result = a.DB.Where(rn.RegNarahubung{RegistrasiNpwpd_Id: uint64(id)}).Delete(&v)
		status = "deleted"
		if result.RowsAffected == 0 {
			dataPemilik = nil
			status = "no deletion"
		}
	}

	// data objekpajak
	var dataObjekPajak []*rn.RegObjekPajak
	result = a.DB.Where(rn.RegObjekPajak{RegistrasiNpwpd_Id: uint64(id)}).Find(&dataObjekPajak)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	for _, v := range dataObjekPajak {
		result = a.DB.Where(rn.RegObjekPajak{RegistrasiNpwpd_Id: uint64(id)}).Delete(&v)
		status = "deleted"
		if result.RowsAffected == 0 {
			dataPemilik = nil
			status = "no deletion"
		}
	}

	// data regis
	var data *rn.RegistrasiNpwpd
	result = a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	// data rekening
	var rekening *rm.Rekening
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	fmt.Println("data rekening: ", *rekening.Objek)
	// delete detail OP
	switch *rekening.Objek {
	case "01":
		var DataOp []*rn.DetailRegOpHotel
		result := a.DB.Where(rn.DetailRegOpHotel{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}

		for _, v := range DataOp {
			result = a.DB.Where(rn.DetailRegOpHotel{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}

	case "02":
		var DataOp []*rn.DetailRegOpResto
		result := a.DB.Where(rn.DetailRegOpResto{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(rn.DetailRegOpResto{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "03":
		var DataOp []*rn.DetailRegOpHiburan
		result := a.DB.Where(rn.DetailRegOpHiburan{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(rn.DetailRegOpHiburan{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "04":
		var DataOp []*rn.DetailRegOpReklame
		result := a.DB.Where(rn.DetailRegOpReklame{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(rn.DetailRegOpReklame{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "05":
		var DataOp []*rn.DetailRegOpPpj
		result := a.DB.Where(rn.DetailRegOpPpj{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(rn.DetailRegOpPpj{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "06":
		var DataOp []*rn.DetailRegOpParkir
		result := a.DB.Where(rn.DetailRegOpParkir{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(rn.DetailRegOpParkir{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "07":
		var DataOp []*rn.DetailRegOpAirTanah
		result := a.DB.Where(rn.DetailRegOpAirTanah{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(rn.DetailRegOpAirTanah{rn.DetailRegOp{RegistrasiNpwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	}

	result = a.DB.Delete(&data, id)
	status = "deleted"
	if result.RowsAffected == 0 {
		data = nil
		status = "no deletion"
	}
	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil

}

func Update(id int, input rn.UpdateDto, user_Id uint) (any, error) {
	var data *rn.RegistrasiNpwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	userIdConv := uint64(user_Id)
	if data.User_Id != userIdConv {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	for _, v := range input.RegPemilik {
		var dataP *rn.RegPemilikWp
		result := a.DB.First(&dataP, v.Id)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		if err := sc.Copy(&dataP, &v); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataP)
		}

		dataP.RegistrasiNpwpd_Id = data.Id
		if result := a.DB.Save(&dataP); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataP)
		}
	}

	for _, v := range input.RegNarahubung {
		var dataN *rn.RegNarahubung
		result := a.DB.First(&dataN, v.Id)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		if err := sc.Copy(&dataN, &v); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataN)
		}
		dataN.RegistrasiNpwpd_Id = data.Id
		if result := a.DB.Save(&dataN); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataN)
		}
	}

	var dataOP *rn.RegObjekPajak
	result = a.DB.First(&dataOP, input.RegObjekPajak.Id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.Copy(&dataOP, &input.RegObjekPajak); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataOP)
	}
	dataOP.RegistrasiNpwpd_Id = data.Id
	if result := a.DB.Save(&dataOP); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataOP)
	}

	//cek rekening objek, dgn mengambil data rekening menggunakan rekening id
	var rekening *rm.Rekening
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	fmt.Println("Data rekening: ", *rekening.Objek)
	switch *rekening.Objek {
	case "01":
		for _, v := range input.DetailRegOp {
			var dataDetail *rn.DetailRegOpHotel
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.RegistrasiNpwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
			}
		}
	case "02":
		for _, v := range input.DetailRegOp {
			var dataDetail *rn.DetailRegOpResto
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.RegistrasiNpwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
			}
		}
	case "03":
		for _, v := range input.DetailRegOp {
			var dataDetail *rn.DetailRegOpHiburan
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.RegistrasiNpwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
			}
		}
	case "04":
		for _, v := range input.DetailRegOp {
			var dataDetail *rn.DetailRegOpReklame
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.RegistrasiNpwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
			}
		}
	case "05":
		for _, v := range input.DetailRegOp {
			var dataDetail *rn.DetailRegOpPpj
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.RegistrasiNpwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
			}
		}
	case "06":
		for _, v := range input.DetailRegOp {
			var dataDetail *rn.DetailRegOpParkir
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.RegistrasiNpwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
			}
		}
	case "07":
		for _, v := range input.DetailRegOp {
			var dataDetail *rn.DetailRegOpAirTanah
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.RegistrasiNpwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
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
