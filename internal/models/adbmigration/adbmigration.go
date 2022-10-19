package adbmigration

import (
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/anggaran"
	adm "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthiburan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthotel"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptparkir"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptppjnonpln"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptppjpln"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptresto"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/group"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/hargadasarair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jabatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jurnal"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/klasifikasijalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/menu"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/omset"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/pangkat"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/registrasinpwpd"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sektor"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/skpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sumberdana"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifjambong"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifjambongrek"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifreklame"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajak"
)

func init() {
	listModelConfigurationReference := []interface{}{
		&skpd.Skpd{},
		&jabatan.Jabatan{},
		&pangkat.Pangkat{},
		&adm.Provinsi{},
		&adm.Daerah{},
		&adm.Kecamatan{},
		&adm.Kelurahan{},
		&sektor.Sektor{},
		&rm.Rekening{},
		&omset.Omset{},
		&anggaran.Anggaran{},
		&sumberdana.SumberDana{},
		&jurnal.Jurnal{},
		&tarifjambongrek.TarifJambongRek{},
		&klasifikasijalan.KlasifikasiJalan{},
		&tarifjambong.TarifJambong{},
		&tarifreklame.TarifReklame{},
		&jalan.Jalan{},
		&hargadasarair.HargaDasarAir{},
		&tarifpajak.TarifPajak{},
	}
	a.AutoMigrate(listModelConfigurationReference...)

	listModelManagementUser := []interface{}{
		&mu.User{},
		&group.Group{},
		&menu.Menu{},
		&pegawai.Pegawai{},
		&ppat.Ppat{},
		&wajibpajak.WajibPajak{},
	}
	a.AutoMigrate(listModelManagementUser...)

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

	listModelPenetapan := []interface{}{
		&espt.Espt{},
		&detailesptair.DetailEsptAir{},
		&detailespthotel.DetailEsptHotel{},
		&detailespthiburan.DetailEsptHiburan{},
		&detailesptparkir.DetailEsptParkir{},
		&detailesptresto.DetailEsptResto{},
		&detailesptppjnonpln.DetailEsptPpjNonPln{},
		&detailesptppjpln.DetailEsptPpjPln{},
	}
	a.AutoMigrate(listModelPenetapan...)
}
