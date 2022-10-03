package registration

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"time"

	registration "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "pendaftaran"

func GetAll(pagination gormhelper.Pagination) (interface{}, error) {
	var (
		register []*registration.Registration
		count    int64
	)

	result := a.DB.Model(&registration.Registration{}).
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
	var register *registration.Registration
	err := a.DB.Model(&registration.Registration{}).
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

func insertDetailOp(objek string, data *[]registration.DetailOp, registerForm *registration.Registration) error {
	var (
		err      error
		mActions map[string]reflect.Type = make(map[string]reflect.Type)
		model    interface{}

		detailOpHotel    registration.DetailOpHotel
		detailOpResto    registration.DetailOpResto
		detailOpHiburan  registration.DetailOpHiburan
		detailOpReklame  registration.DetailOpReklame
		detailOpPpj      registration.DetailOpPpj
		detailOpParkir   registration.DetailOpParkir
		detailOpAirTanah registration.DetailOpAirTanah
	)

	mActions[`detailOpHotel`] = reflect.TypeOf(detailOpHotel)
	mActions[`detailOpResto`] = reflect.TypeOf(detailOpResto)
	mActions[`detailOpHiburan`] = reflect.TypeOf(detailOpHiburan)
	mActions[`detailOpReklame`] = reflect.TypeOf(detailOpReklame)
	mActions[`detailOpPpj`] = reflect.TypeOf(detailOpPpj)
	mActions[`detailOpParkir`] = reflect.TypeOf(detailOpParkir)
	mActions[`detailOpAirTanah`] = reflect.TypeOf(detailOpAirTanah)

	switch objek {
	case "01":
		model = reflect.Zero(mActions["detailOpHotel"]).Interface()
	case "02":
		model = reflect.Zero(mActions["detailOpResto"]).Interface()
	case "03":
		model = reflect.Zero(mActions["detailOpHiburan"]).Interface()
	case "04":
		model = reflect.Zero(mActions["detailOpReklame"]).Interface()
	case "05":
		model = reflect.Zero(mActions["detailOpPpj"]).Interface()
	case "06":
		model = reflect.Zero(mActions["detailOpParkir"]).Interface()
	case "07":
		model = reflect.Zero(mActions["detailOpAirTanah"]).Interface()
	}

	for _, dop := range *data {
		dop.JenisOp = registerForm.Rekening.Nama
		dop.Pendaftaran_Id = registerForm.Id
		dop.Pendaftaran_Npwpd = registerForm.Npwpd

		m := make(map[string]interface{})
		m["Pendaftaran_Id"] = dop.Pendaftaran_Id
		m["Pendaftaran_Npwpd"] = dop.Pendaftaran_Npwpd
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

func RegisterByOperator(r *http.Request, reg registration.RegisterByOperator) (interface{}, error) {
	var rekening *rm.Rekening
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, reg.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	var tmpverify = registration.VerifiyPendaftaranDisetujui
	register := registration.Registration{
		Mode:         registration.ModeOperator,
		Status:       registration.StatusAktif,
		JenisPajak:   reg.JenisPajak,
		Golongan:     reg.Golongan,
		Npwp:         reg.Npwp,
		VerifyStatus: &tmpverify,
		Nomor: func() *string {
			unique := strconv.FormatInt(time.Now().Unix(), 10)
			if reg.IsNomorRegistrasiAuto {
				return &unique
			}
			return reg.NomorRegistrasi
		}(),
		NomorRegistrasi: reg.NomorRegistrasi,
		Npwpd:           reg.Npwpd,
		TanggalPengukuhan: func() *time.Time {
			t, err := time.Parse("2006-01-02", *reg.TanggalPengukuhan)
			if err != nil {
				return nil
			}
			return &t
		}(),
		TanggalNpwpd: func() *time.Time {
			t, err := time.Parse("2006-01-02", *reg.TanggalNpwpd)
			if err != nil {
				return nil
			}
			return &t
		}(),
		Rekening_Id: &reg.Rekening_Id,
		Rekening:    rekening,
		TanggalMulaiUsaha: func() *time.Time {
			t, err := time.Parse("2006-01-02", *reg.TanggalMulaiUsaha)
			if err != nil {
				return nil
			}
			return &t
		}(),
		LuasBangunan:  reg.LuasBangunan,
		JamBukaUsaha:  reg.JamBukaUsaha,
		JamTutupUsaha: reg.JamTutupUsaha,
		Pengunjung:    reg.Pengunjung,
		OmsetOp:       reg.OmsetOp,
		Genset:        &reg.Genset,
		AirTanah:      &reg.AirTanah,
	}
	err = a.DB.Create(&register).Error
	if err != nil {
		return nil, err
	}
	err = insertDetailOp(*rekening.Objek, reg.DetailOp, &register)
	if err != nil {
		return nil, err
	}
	for _, op := range *reg.ObjekPajak {
		op.Pendaftaran_Id = register.Id
		op.Status = registration.StatusBaru
		err := a.DB.Create(&op).Error
		if err != nil {
			return nil, err
		}
	}
	for _, p := range *reg.Pemilik {
		p.Pendaftaran_Id = register.Id
		p.Status = registration.StatusPemilikBaru
		err := a.DB.Create(&p).Error
		if err != nil {
			return nil, err
		}
	}
	for _, n := range *reg.Narahubung {
		n.Pendaftaran_Id = register.Id
		n.Status = registration.StatusNarahubungBaru
		err := a.DB.Create(&n).Error
		if err != nil {
			return nil, err
		}
	}
	return rp.OKSimple{
		Data: register,
	}, nil
}

func Update(id int, input registration.RegisterUpdate) (any, error) {
	//cek rekening objek, dgn mengambil data rekening menggunakan rekening id
	var data *registration.Registration
	// var rekening *rm.Rekening
	// err := apicore.DB.Model(&rm.Rekening{}).First(&rekening, input.RekeningID).Error
	// if err != nil {
	// 	return nil, err
	// }

	// switch *rekening.Objek {
	// case "01":
	// 	// model = reflect.Zero(mActions["detailOpHotel"]).Interface()
	// 	var DataOp *registration.DetailOpHotel
	// 	result := a.DB.Where(registration.DetailOpHotel{registration.DetailOp{Pendaftaran_ID: uint64(id)}}).First(&DataOp)
	// 	if result.RowsAffected == 0 {
	// 		return nil, errors.New("data tidak dapat ditemukan")
	// 	}
	// 	if err := sc.Copy(&DataOp, &input); err != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
	// 	}

	// 	if result := a.DB.Save(&DataOp); result.Error != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
	// 	}
	// case "02":
	// 	// model = reflect.Zero(mActions["detailOpResto"]).Interface()
	// 	var DataOp *registration.DetailOpResto
	// 	result := a.DB.Where(registration.DetailOpResto{registration.DetailOp{Pendaftaran_ID: uint64(id)}}).First(&DataOp)
	// 	if result.RowsAffected == 0 {
	// 		return nil, errors.New("data tidak dapat ditemukan")
	// 	}
	// 	if err := sc.Copy(&DataOp, &input); err != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
	// 	}

	// 	if result := a.DB.Save(&DataOp); result.Error != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
	// 	}
	// case "03":
	// 	// model = reflect.Zero(mActions["detailOpHiburan"]).Interface()
	// 	var DataOp *registration.DetailOpHiburan
	// 	result := a.DB.Where(registration.DetailOpHiburan{registration.DetailOp{Pendaftaran_ID: uint64(id)}}).First(&DataOp)
	// 	if result.RowsAffected == 0 {
	// 		return nil, errors.New("data tidak dapat ditemukan")
	// 	}
	// 	if err := sc.Copy(&DataOp, &input); err != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
	// 	}

	// 	if result := a.DB.Save(&DataOp); result.Error != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
	// 	}
	// case "04":
	// 	// model = reflect.Zero(mActions["detailOpReklame"]).Interface()
	// 	var DataOp *registration.DetailOpReklame
	// 	result := a.DB.Where(registration.DetailOpReklame{registration.DetailOp{Pendaftaran_ID: uint64(id)}}).First(&DataOp)
	// 	if result.RowsAffected == 0 {
	// 		return nil, errors.New("data tidak dapat ditemukan")
	// 	}
	// 	if err := sc.Copy(&DataOp, &input); err != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
	// 	}

	// 	if result := a.DB.Save(&DataOp); result.Error != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
	// 	}
	// case "05":
	// 	// model = reflect.Zero(mActions["detailOpPpj"]).Interface()
	// 	var DataOp *registration.DetailOpPpj
	// 	result := a.DB.Where(registration.DetailOpPpj{registration.DetailOp{Pendaftaran_ID: uint64(id)}}).First(&DataOp)
	// 	if result.RowsAffected == 0 {
	// 		return nil, errors.New("data tidak dapat ditemukan")
	// 	}
	// 	if err := sc.Copy(&DataOp, &input); err != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
	// 	}

	// 	if result := a.DB.Save(&DataOp); result.Error != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
	// 	}
	// case "06":
	// 	// model = reflect.Zero(mActions["detailOpParkir"]).Interface()
	// 	var DataOp *registration.DetailOpParkir
	// 	result := a.DB.Where(registration.DetailOpParkir{registration.DetailOp{Pendaftaran_ID: uint64(id)}}).First(&DataOp)
	// 	if result.RowsAffected == 0 {
	// 		return nil, errors.New("data tidak dapat ditemukan")
	// 	}
	// 	if err := sc.Copy(&DataOp, &input); err != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
	// 	}

	// 	if result := a.DB.Save(&DataOp); result.Error != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
	// 	}
	// case "07":
	// 	// model = reflect.Zero(mActions["detailOpAirTanah"]).Interface()
	// 	var DataOp *registration.DetailOpAirTanah
	// 	result := a.DB.Where(registration.DetailOpAirTanah{registration.DetailOp{Pendaftaran_ID: uint64(id)}}).First(&DataOp)
	// 	if result.RowsAffected == 0 {
	// 		return nil, errors.New("data tidak dapat ditemukan")
	// 	}
	// 	if err := sc.Copy(&DataOp, &input); err != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", DataOp)
	// 	}

	// 	if result := a.DB.Save(&DataOp); result.Error != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", DataOp)
	// 	}
	// }

	// var dataObjekPajak *registration.ObjekPajak
	// var dataPemilikWp *registration.PemilikWp
	// var dataNarahubung *registration.Narahubung

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
	// result = a.DB.Where(registration.ObjekPajak{Pendaftaran_Id: uint64(id)}).First(&dataObjekPajak)
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
	// result = a.DB.Where(registration.PemilikWp{Pendaftaran_Id: uint64(id)}).First(&dataPemilikWp)
	// if result.RowsAffected == 0 {
	// 	return nil, errors.New("data tidak dapat ditemukan")
	// }
	// if err := sc.Copy(&dataPemilikWp, &input); err != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	// }

	// if result := a.DB.Save(&dataPemilikWp); result.Error != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", dataPemilikWp)
	// }

	// //data narahubung
	// result = a.DB.Where(registration.Narahubung{Pendaftaran_Id: uint64(id)}).First(&dataNarahubung)
	// if result.RowsAffected == 0 {
	// 	return nil, errors.New("data tidak dapat ditemukan")
	// }
	// fmt.Println("data: ", input.Narahubung)
	// fmt.Println("data narahubung: ", dataNarahubung)
	// if err := sc.Copy(&dataNarahubung, &input.Narahubung[0]); err != nil {
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

	// // data pemilikwp
	// var dataPemilik *registration.PemilikWp
	// result := a.DB.Where(registration.PemilikWp{Pendaftaran_Id: uint64(id)}).First(&dataPemilik)
	// if result.RowsAffected == 0 {
	// 	return nil, errors.New("data tidak dapat ditemukan")
	// }

	// result = a.DB.Where(registration.PemilikWp{Pendaftaran_Id: uint64(id)}).Delete(&dataPemilik)
	// status := "deleted"
	// if result.RowsAffected == 0 {
	// 	dataPemilik = nil
	// 	status = "no deletion"
	// }

	// // data narahubung
	// var dataNarahubung *registration.Narahubung
	// result = a.DB.Where(registration.Narahubung{Pendaftaran_Id: uint64(id)}).First(&dataNarahubung)
	// if result.RowsAffected == 0 {
	// 	return nil, errors.New("data tidak dapat ditemukan")
	// }

	// result = a.DB.Where(registration.Narahubung{Pendaftaran_Id: uint64(id)}).Delete(&dataNarahubung)
	// status = "deleted"
	// if result.RowsAffected == 0 {
	// 	dataPemilik = nil
	// 	status = "no deletion"
	// }

	// // data objekpajak
	// var dataObjekPajak *registration.ObjekPajak
	// result = a.DB.Where(registration.ObjekPajak{Pendaftaran_Id: uint64(id)}).First(&dataObjekPajak)
	// if result.RowsAffected == 0 {
	// 	return nil, errors.New("data tidak dapat ditemukan")
	// }

	// result = a.DB.Where(registration.ObjekPajak{Pendaftaran_Id: uint64(id)}).Delete(&dataObjekPajak)
	// status = "deleted"
	// if result.RowsAffected == 0 {
	// 	dataPemilik = nil
	// 	status = "no deletion"
	// }

	// // data detailop
	// var dataDetailOp *registration.DetailOpHiburan
	// result = a.DB.Where(registration.DetailOpHiburan{registration.DetailOp{Pendaftaran_Id: uint64(id)}}).First(&dataDetailOp)
	// if result.RowsAffected == 0 {
	// 	return nil, errors.New("data tidak dapat ditemukan")
	// }

	// result = a.DB.Where(registration.DetailOpHiburan{registration.DetailOp{Pendaftaran_Id: uint64(id)}}).Delete(&dataDetailOp)
	// status = "deleted"
	// if result.RowsAffected == 0 {
	// 	dataPemilik = nil
	// 	status = "no deletion"
	// }

	var data *registration.Registration
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	result = a.DB.Delete(&data, id)
	status := "deleted"
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

func VerifyNpwpd(id int, input registration.VerifikasiDto) (any, error) {
	var data *registration.Registration
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	var tmp = time.Now()
	data.VerifyAt = &tmp
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
