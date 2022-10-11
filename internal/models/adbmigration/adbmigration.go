package adbmigration

import (
	"os/user"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"

	adm "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
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

}
