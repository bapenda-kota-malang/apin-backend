package registration

import (
	"net/http"
	"strconv"
	"time"

	registration "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

func GetAll(r *http.Request) (interface{}, error) {
	var register []*registration.Registration
	result := apicore.DB.Scopes(gormhelper.Paginate(r)).Find(&register)
	return responses.OK{
		Meta: types.IS{
			"Count": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: register,
	}, nil
}

func RegisterByOperator(r *http.Request, reg registration.RegisterByOperator) (interface{}, error) {
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
		RekeningID: &reg.RekeningID,
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
	err := apicore.DB.Create(&register).Error
	if err != nil {
		return nil, err
	}
	for _, op := range *reg.ObjekPajak {
		op.PendaftaranID = register.ID
		op.Status = registration.StatusBaru
		err := apicore.DB.Create(&op).Error
		if err != nil {
			return nil, err
		}
	}
	for _, p := range *reg.Pemilik {
		p.PendaftaranID = register.ID
		p.Status = registration.StatusPemilikBaru
		err := apicore.DB.Create(&p).Error
		if err != nil {
			return nil, err
		}
	}
	for _, n := range *reg.Narahubung {
		n.PendaftaranID = register.ID
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
