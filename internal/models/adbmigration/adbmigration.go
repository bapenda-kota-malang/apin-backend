package adbmigration

import (
	"os/user"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"

	adm "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthotel"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptparkir"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptreklame"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptresto"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/registrasinpwpd"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/skpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajak"
)

func init() {
	listModelPendaftaran := []interface{}{
		&user.User{},
		&wajibpajak.WajibPajak{},
		&adm.Provinsi{},
		&adm.Daerah{},
		&adm.Kecamatan{},
		&adm.Kelurahan{},
		&rm.Rekening{},
		&skpd.Skpd{},
	}
	a.AutoMigrate(listModelPendaftaran...)

	listModelPenetapan := []interface{}{
		&espt.Espt{},
		&detailesptair.DetailEsptAir{},
		&detailespthotel.DetailEsptHotel{},
		&detailesptparkir.DetailEsptParkir{},
		&detailesptresto.DetailEsptResto{},
		&detailesptreklame.DetailEsptReklame{},
	}
	a.AutoMigrate(listModelPenetapan...)

	listModelPendataan := []interface{}{
		&potensiopwp.PotensiOp{},
		&potensiopwp.PotensiPemilikWp{},
		&potensiopwp.PotensiNarahubung{},
		&potensiopwp.DetailPotensiOp{},
		&potensiopwp.DetailPotensiAirTanah{},
		&potensiopwp.DetailPotensiHiburan{},
		&potensiopwp.DetailPotensiHotel{},
		&potensiopwp.DetailPotensiPPJ{},
		&potensiopwp.DetailPotensiParkir{},
		&potensiopwp.DetailPotensiReklame{},
		&potensiopwp.DetailPotensiResto{},
	}
	a.AutoMigrate(listModelPendataan...)

	listModelRegNpwpd := []interface{}{
		&rn.RegistrasiNpwpd{},
		&rn.DetailRegOpHotel{},
		&rn.DetailRegOpAirTanah{},
		&rn.DetailRegOpParkir{},
		&rn.DetailRegOpReklame{},
		&rn.DetailRegOpPpj{},
		&rn.DetailRegOpHiburan{},
		&rn.DetailRegOpResto{},
		&rn.RegObjekPajak{},
		&rn.RegNarahubung{},
		&rn.RegPemilikWp{},
	}

	a.AutoMigrate(listModelRegNpwpd...)

	listModelNpwpd := []interface{}{
		&npwpd.Npwpd{},
		&npwpd.PemilikWp{},
		&npwpd.ObjekPajak{},
		&npwpd.Narahubung{},
		&npwpd.DetailOpHotel{},
		&npwpd.DetailOpAirTanah{},
		&npwpd.DetailOpParkir{},
		&npwpd.DetailOpReklame{},
		&npwpd.DetailOpPpj{},
		&npwpd.DetailOpHiburan{},
		&npwpd.DetailOpResto{},
	}

	a.AutoMigrate(listModelNpwpd...)
}
