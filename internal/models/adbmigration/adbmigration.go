package adbmigration

import (
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/anggaran"
	adm "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailespthiburan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailespthotel"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptparkir"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptppjnonpln"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptppjpln"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptresto"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/group"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/hargadasarair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jabatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jenisppj"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jurnal"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/klasifikasijalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/menu"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/omset"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/pangkat"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/bapl"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailobjek"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensinarahubung"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensipemilikwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sinkronisasi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthiburan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthotel"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptppjnonpln"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptppjpln"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt/sptnomertracker"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/hargareferensi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/nik"
	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/regnpwpd"
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
		&jenisppj.JenisPPJ{},
		&hargareferensi.HargaReferensi{},
		&nik.Nik{},
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
		&regobjekpajak.RegObjekPajak{},
		&rn.RegNpwpd{},
		&rn.DetailRegObjekPajakHotel{},
		&rn.DetailRegObjekPajakAirTanah{},
		&rn.DetailRegObjekPajakParkir{},
		&rn.DetailRegObjekPajakReklame{},
		&rn.DetailRegObjekPajakPpj{},
		&rn.DetailRegObjekPajakHiburan{},
		&rn.DetailRegObjekPajakResto{},
		&rn.RegNarahubung{},
		&rn.RegPemilikWp{},
	}
	a.AutoMigrate(listModelRegNpwpd...)

	listModelNpwpd := []interface{}{
		&objekpajak.ObjekPajak{},
		&npwpd.Npwpd{},
		&npwpd.PemilikWp{},
		&npwpd.Narahubung{},
		&npwpd.DetailObjekPajakHotel{},
		&npwpd.DetailObjekPajakAirTanah{},
		&npwpd.DetailObjekPajakParkir{},
		&npwpd.DetailObjekPajakReklame{},
		&npwpd.DetailObjekPajakPpj{},
		&npwpd.DetailObjekPajakHiburan{},
		&npwpd.DetailObjekPajakResto{},
	}
	a.AutoMigrate(listModelNpwpd...)

	listModelPendataan := []interface{}{
		&potensiopwp.PotensiOp{},
		&bapl.Bapl{},
		&detailpotensiop.DetailPotensiOp{},
		&potensipemilikwp.PotensiPemilikWp{},
		&potensinarahubung.PotensiNarahubung{},
		&detailobjek.DetailPotensiAirTanah{},
		&detailobjek.DetailPotensiHiburan{},
		&detailobjek.DetailPotensiHotel{},
		&detailobjek.DetailPotensiPPJ{},
		&detailobjek.DetailPotensiParkir{},
		&detailobjek.DetailPotensiReklame{},
		&detailobjek.DetailPotensiResto{},
	}
	a.AutoMigrate(listModelPendataan...)

	listModelPenetapan := []interface{}{
		// esptpd
		&espt.Espt{},
		&detailesptair.DetailEsptAir{},
		&detailespthotel.DetailEsptHotel{},
		&detailespthiburan.DetailEsptHiburan{},
		&detailesptparkir.DetailEsptParkir{},
		&detailesptresto.DetailEsptResto{},
		&detailesptppjnonpln.DetailEsptPpjNonPln{},
		&detailesptppjpln.DetailEsptPpjPln{},

		// sptpd
		&sptnomertracker.SptNomerTracker{},
		&spt.Spt{},
		&detailsptair.DetailSptAir{},
		&detailspthiburan.DetailSptHiburan{},
		&detailspthotel.DetailSptHotel{},
		&detailsptparkir.DetailSptParkir{},
		&detailsptppjnonpln.DetailSptPpjNonPln{},
		&detailsptppjpln.DetailSptPpjPln{},
		&detailsptreklame.DetailSptReklame{},
		&detailsptresto.DetailSptResto{},
	}
	a.AutoMigrate(listModelPenetapan...)

	listModelPembayaran := []interface{}{
		&sspd.Sspd{},
		&sspd.SspdDetail{},
		&sinkronisasi.Sinkronisasi{},
		&sinkronisasi.SinkronisasiDetail{},
	}
	a.AutoMigrate(listModelPembayaran...)
}
