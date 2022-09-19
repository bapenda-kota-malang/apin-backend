package registration

import (
	"net/http"
	"reflect"
	"strconv"
	"time"

	registration "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAll(pagination gormhelper.Pagination) (interface{}, error) {
	var (
		register []*registration.Registration
		count    int64
	)

	result := apicore.DB.Model(&registration.Registration{}).
		Preload(clause.Associations).
		//nested preload
		//Preload("PemilikWps.Kelurahan").
		Count(&count).
		Scopes(gormhelper.Paginate(&pagination)).
		Find(&register)

	return responses.OK{
		Meta: types.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: register,
	}, nil
}

func GetDetail(r *http.Request, regID string) (interface{}, error) {
	var register *registration.Registration
	err := apicore.DB.Model(&registration.Registration{}).
		Preload(clause.Associations).
		First(&register, regID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return responses.OKSimple{
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
		dop.Pendaftaran_ID = registerForm.ID

		m := make(map[string]interface{})
		m["Pendaftaran_ID"] = dop.Pendaftaran_ID
		m["JenisOp"] = dop.JenisOp
		m["JumlahOp"] = dop.JumlahOp
		m["TarifOp"] = dop.TarifOp
		m["UnitOp"] = dop.UnitOp
		m["Notes"] = dop.Notes

		err = apicore.DB.Model(&model).Create(&m).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func RegisterByOperator(r *http.Request, reg registration.RegisterByOperator) (interface{}, error) {
	var rekening *rm.Rekening
	err := apicore.DB.Model(&rm.Rekening{}).First(&rekening, reg.RekeningID).Error
	if err != nil {
		return nil, err
	}
	register := registration.Registration{
		Mode:       registration.ModeOperator,
		Status:     registration.StatusAktif,
		JenisPajak: reg.JenisPajak,
		Golongan:   reg.Golongan,
		Npwp:       reg.Npwp,
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
		Rekening_ID: &reg.RekeningID,
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
	err = apicore.DB.Create(&register).Error
	if err != nil {
		return nil, err
	}
	err = insertDetailOp(*rekening.Objek, reg.DetailOp, &register)
	if err != nil {
		return nil, err
	}
	for _, op := range *reg.ObjekPajak {
		op.Pendaftaran_ID = register.ID
		op.Status = registration.StatusBaru
		err := apicore.DB.Create(&op).Error
		if err != nil {
			return nil, err
		}
	}
	for _, p := range *reg.Pemilik {
		p.Pendaftaran_ID = register.ID
		p.Status = registration.StatusPemilikBaru
		err := apicore.DB.Create(&p).Error
		if err != nil {
			return nil, err
		}
	}
	for _, n := range *reg.Narahubung {
		n.Pendaftaran_ID = register.ID
		n.Status = registration.StatusNarahubungBaru
		err := apicore.DB.Create(&n).Error
		if err != nil {
			return nil, err
		}
	}
	return responses.OKSimple{
		Data: register,
	}, nil
}
