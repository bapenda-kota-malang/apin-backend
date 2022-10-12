package npwpd

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "npwpd"

func GetAll(pagination gormhelper.Pagination) (interface{}, error) {
	var (
		register []*npwpd.Npwpd
		count    int64
	)

	result := a.DB.Model(&npwpd.Npwpd{}).
		Preload(clause.Associations).
		//nested preload
		//Preload("PemilikWps.Kelurahan").
		Count(&count).
		// Scopes(gormhelper.Paginate(&pagination)).
		Find(&register)

	return rp.OK{
		Meta: types.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: register,
	}, nil
}

func GetDetail(r *http.Request, regID int) (interface{}, error) {
	var register *npwpd.Npwpd
	err := a.DB.Model(&npwpd.Npwpd{}).
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

func Create(r *http.Request, reg npwpd.CreateDto) (interface{}, error) {
	var rekening *rm.Rekening
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, reg.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	// var tmpverify = npwpd.VerifiyPendaftaranDisetujui
	var nomorNpwpd = func() *uint64 {

		if reg.IsNomorRegistrasiAuto {
			var tmp uint64
			var tmpNpwpd npwpd.Npwpd
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
	var nomorString string
	if nomorNpwpd == nil {
		nomorString = "0001"
	} else {

		nomorString = strconv.Itoa(int(*nomorNpwpd))
	}
	kecamatanIdString := strconv.Itoa(int(*reg.ObjekPajak.Kecamatan_Id))
	kodeJenisUsahaString := *rekening.KodeJenisUsaha
	fmt.Println("data rekening: ", *rekening.Nama)
	fmt.Println("kodejenisusaha: ", kodeJenisUsahaString)
	if kodeJenisUsahaString == "" {
		kodeJenisUsahaString = "xxx"
	}
	npwpdString := nomorString + "." + kecamatanIdString + "." + kodeJenisUsahaString
	register := npwpd.Npwpd{
		JalurRegistrasi: nt.JalurRegistrasiOperator,
		Status:          nt.StatusAktif,
		JenisPajak:      reg.JenisPajak,
		Golongan:        reg.Golongan,
		Npwp:            reg.Npwp,
		// VerifyStatus:   &tmpverify,
		VerifiedAt: func() *time.Time {
			t := time.Now()
			return &t
		}(),
		Nomor:             nomorNpwpd,
		Npwpd:             &npwpdString,
		TanggalPengukuhan: th.ParseTime(*reg.TanggalPengukuhan),
		TanggalNpwpd:      th.ParseTime(*reg.TanggalNpwpd),
		Rekening_Id:       &reg.Rekening_Id,
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
	err = insertDetailOp(*rekening.Objek, reg.DetailOp, &register)
	if err != nil {
		return nil, err
	}
	// objekpajak
	op := *reg.ObjekPajak
	op.Npwpd_Id = register.Id
	op.Status = nt.StatusBaru
	errOp := a.DB.Create(&op).Error
	if errOp != nil {
		return nil, err
	}

	for _, p := range *reg.Pemilik {
		p.Npwpd_Id = register.Id
		// p.Status = npwpd.StatusPemilikBaru
		err := a.DB.Create(&p).Error
		if err != nil {
			return nil, err
		}
	}
	for _, n := range *reg.Narahubung {
		n.Npwpd_Id = register.Id
		// n.Status = npwpd.StatusNarahubungBaru
		err := a.DB.Create(&n).Error
		if err != nil {
			return nil, err
		}
	}
	return rp.OKSimple{
		Data: register,
	}, nil
}

func Update(id int, input npwpd.UpdateDto) (any, error) {
	//cek rekening objek, dgn mengambil data rekening menggunakan rekening id
	var data *npwpd.Npwpd
	var rekening *rm.Rekening
	err := apicore.DB.Model(&rm.Rekening{}).First(&rekening, input.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	fmt.Println("Data rekening: ", *rekening.Objek)
	switch *rekening.Objek {
	case "01":
		// model = reflect.Zero(mActions["detailOpHotel"]).Interface()
		var DataOp *npwpd.DetailOpHotel
		result := a.DB.Where(npwpd.DetailOpHotel{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).First(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		if err := sc.Copy(&DataOp, &input); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
		}

		if result := a.DB.Save(&DataOp); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
		}
	case "02":
		// model = reflect.Zero(mActions["detailOpResto"]).Interface()
		var DataOp *npwpd.DetailOpResto
		result := a.DB.Where(npwpd.DetailOpResto{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).First(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		if err := sc.Copy(&DataOp, &input); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
		}

		if result := a.DB.Save(&DataOp); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
		}
	case "03":
		// model = reflect.Zero(mActions["detailOpHiburan"]).Interface()
		var DataOp *npwpd.DetailOpHiburan
		result := a.DB.Where(npwpd.DetailOpHiburan{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).First(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		if err := sc.Copy(&DataOp, &input); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
		}

		if result := a.DB.Save(&DataOp); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
		}
	case "04":
		// model = reflect.Zero(mActions["detailOpReklame"]).Interface()
		var DataOp *npwpd.DetailOpReklame
		result := a.DB.Where(npwpd.DetailOpReklame{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).First(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		fmt.Println("HERE")
		if err := sc.Copy(&DataOp, &input.DetailOp); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
		}
		fmt.Println("Data OP: ", *input.DetailOp.JumlahOp)
		if result := a.DB.Save(&DataOp); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
		}
	case "05":
		// model = reflect.Zero(mActions["detailOpPpj"]).Interface()
		var DataOp *npwpd.DetailOpPpj
		result := a.DB.Where(npwpd.DetailOpPpj{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).First(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		if err := sc.Copy(&DataOp, &input); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
		}

		if result := a.DB.Save(&DataOp); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
		}
	case "06":
		// model = reflect.Zero(mActions["detailOpParkir"]).Interface()
		var DataOp *npwpd.DetailOpParkir
		result := a.DB.Where(npwpd.DetailOpParkir{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).First(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		if err := sc.Copy(&DataOp, &input); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
		}

		if result := a.DB.Save(&DataOp); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
		}
	case "07":
		// model = reflect.Zero(mActions["detailOpAirTanah"]).Interface()
		var DataOp *npwpd.DetailOpAirTanah
		result := a.DB.Where(npwpd.DetailOpAirTanah{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).First(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		if err := sc.Copy(&DataOp, &input); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
		}

		if result := a.DB.Save(&DataOp); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
		}
	}

	// var dataObjekPajak *npwpd.ObjekPajak

	// var dataNarahubung *npwpd.Narahubung

	//data regis
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", data)
	}

	// //data objekpajak
	// result = a.DB.Where(npwpd.ObjekPajak{Npwpd_Id: uint64(id)}).First(&dataObjekPajak)
	// if result.RowsAffected == 0 {
	// 	return nil, errors.New("data tidak dapat ditemukan")
	// }
	// if err := sc.Copy(&dataObjekPajak, &input); err != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataObjekPajak)
	// }

	// if result := a.DB.Save(&dataObjekPajak); result.Error != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", dataObjekPajak)
	// }

	// //data pemilikwp
	// var dataPemilikWp []*npwpd.PemilikWp
	// result = a.DB.Where(npwpd.PemilikWp{Npwpd_Id: uint64(id)}).Find(&dataPemilikWp)
	// if result.RowsAffected == 0 {
	// 	return nil, errors.New("data tidak dapat ditemukan")
	// }
	// for _, v := range dataPemilikWp {
	// 	var tmp *npwpd.PemilikWp
	// 	if err := sc.Copy(&tmp, &v); err != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	// 	}
	// 	dataPemilikWp = append(dataPemilikWp, tmp)
	// }

	// if result := a.DB.Save(&dataPemilikWp); result.Error != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", dataPemilikWp)
	// }
	// //data narahubung
	// result = a.DB.Where(npwpd.Narahubung{Npwpd_Id: uint64(id)}).First(&dataNarahubung)
	// if result.RowsAffected == 0 {
	// 	return nil, errors.New("data tidak dapat ditemukan")
	// }
	// fmt.Println("data: ", input.Narahubung)
	// fmt.Println("data narahubung: ", dataNarahubung)
	// if err := sc.Copy(&dataNarahubung, &input.Narahubung); err != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataNarahubung)
	// }
	// fmt.Println("data narahubung after: ", dataNarahubung)

	// if result := a.DB.Save(&dataNarahubung); result.Error != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", dataNarahubung)
	// }

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (any, error) {

	// data pemilikwp
	var dataPemilik *npwpd.PemilikWp
	result := a.DB.Where(npwpd.PemilikWp{Npwpd_Id: uint64(id)}).First(&dataPemilik)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	result = a.DB.Where(npwpd.PemilikWp{Npwpd_Id: uint64(id)}).Delete(&dataPemilik)
	status := "deleted"
	if result.RowsAffected == 0 {
		dataPemilik = nil
		status = "no deletion"
	}

	// data narahubung
	var dataNarahubung *npwpd.Narahubung
	result = a.DB.Where(npwpd.Narahubung{Npwpd_Id: uint64(id)}).First(&dataNarahubung)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	result = a.DB.Where(npwpd.Narahubung{Npwpd_Id: uint64(id)}).Delete(&dataNarahubung)
	status = "deleted"
	if result.RowsAffected == 0 {
		dataPemilik = nil
		status = "no deletion"
	}

	// data objekpajak
	var dataObjekPajak *npwpd.ObjekPajak
	result = a.DB.Where(npwpd.ObjekPajak{Npwpd_Id: uint64(id)}).First(&dataObjekPajak)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	result = a.DB.Where(npwpd.ObjekPajak{Npwpd_Id: uint64(id)}).Delete(&dataObjekPajak)
	status = "deleted"
	if result.RowsAffected == 0 {
		dataPemilik = nil
		status = "no deletion"
	}

	// // data detailop

	// var dataDetailOp *npwpd.DetailOpHiburan
	// result = a.DB.Where(npwpd.DetailOpHiburan{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).First(&dataDetailOp)
	// if result.RowsAffected == 0 {
	// 	return nil, errors.New("data tidak dapat ditemukan")
	// }

	// result = a.DB.Where(npwpd.DetailOpHiburan{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Delete(&dataDetailOp)
	// status = "deleted"
	// if result.RowsAffected == 0 {
	// 	dataPemilik = nil
	// 	status = "no deletion"
	// }

	// data regis
	var data *npwpd.Npwpd
	result = a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	// data rekening
	var rekening *rm.Rekening
	err := apicore.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	fmt.Println("data rekening: ", *rekening.Objek)
	// delete detail OP
	switch *rekening.Objek {
	case "01":
		// model = reflect.Zero(mActions["detailOpHotel"]).Interface()
		var DataOp []*npwpd.DetailOpHotel
		result := a.DB.Where(npwpd.DetailOpHotel{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}

		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailOpHotel{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}

	case "02":
		// model = reflect.Zero(mActions["detailOpResto"]).Interface()
		var DataOp []*npwpd.DetailOpResto
		result := a.DB.Where(npwpd.DetailOpResto{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailOpResto{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "03":
		// model = reflect.Zero(mActions["detailOpHiburan"]).Interface()
		var DataOp []*npwpd.DetailOpHiburan
		result := a.DB.Where(npwpd.DetailOpHiburan{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailOpHiburan{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "04":
		// model = reflect.Zero(mActions["detailOpReklame"]).Interface()
		var DataOp []*npwpd.DetailOpReklame
		result := a.DB.Where(npwpd.DetailOpReklame{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailOpReklame{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "05":
		// model = reflect.Zero(mActions["detailOpPpj"]).Interface()
		var DataOp []*npwpd.DetailOpPpj
		result := a.DB.Where(npwpd.DetailOpPpj{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailOpPpj{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "06":
		// model = reflect.Zero(mActions["detailOpParkir"]).Interface()
		var DataOp []*npwpd.DetailOpParkir
		result := a.DB.Where(npwpd.DetailOpParkir{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailOpParkir{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "07":
		// model = reflect.Zero(mActions["detailOpAirTanah"]).Interface()
		var DataOp []*npwpd.DetailOpAirTanah
		result := a.DB.Where(npwpd.DetailOpAirTanah{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailOpAirTanah{npwpd.DetailOp{Npwpd_Id: uint64(id)}}).Delete(&v)
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

func VerifyUser(id int, input mu.VerifikasiDto) (any, error) {
	var data *mu.User
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

// func VerifyNpwpd(id int, input npwpd.VerifikasiDto) (any, error) {
// 	var data *npwpd.Npwpd
// 	result := a.DB.First(&data, id)
// 	if result.RowsAffected == 0 {
// 		return nil, errors.New("data tidak dapat ditemukan")
// 	}

// 	if err := sc.Copy(&data, &input); err != nil {
// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
// 	}

// 	var tmp = time.Now()
// 	data.VerifiedAt = &tmp
// 	if result := a.DB.Save(&data); result.Error != nil {
// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
// 	}
// 	return rp.OK{
// 		Meta: t.IS{
// 			"affected": strconv.Itoa(int(result.RowsAffected)),
// 		},
// 		Data: data,
// 	}, nil
// }
