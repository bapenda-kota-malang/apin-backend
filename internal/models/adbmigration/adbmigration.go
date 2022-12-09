package adbmigration

import (
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/anggaran"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/anggotaobjekpajak"
	adm "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/bapenagihan"
	bapenagihanpetugas "github.com/bapenda-kota-malang/apin-backend/internal/models/bapenagihan/petugas"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/baplpengajuan"
	bphtb "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/datanir"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/datapetablok"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/datapetaznt"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbfasum"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbfasum/depjpbklsbintang"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbfasum/depminmax"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbfasum/nondep"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb12"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb13"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb14"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb15"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb16"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb2"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb3"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb4"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb5"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb6"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb7"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb8"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbjpb9"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbmezanin"
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
	"github.com/bapenda-kota-malang/apin-backend/internal/models/indukobjekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jabatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jenispajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jenisppj"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jenisusaha"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jpb"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jurnal"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/keberatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/kelasbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/kelastanah"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/klasifikasijalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/konfigurasipajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/kunjungankembali"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/menu"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/njoptkpflag"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/omset"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/pangkat"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/paymentpoint"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	pstpermohonan "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/pengurangan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/bapl"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailobjek"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensinarahubung"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensipemilikwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/referensibank"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/reganggotaobjekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/regindukobjekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/regkunjungankembali"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbumi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakpbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/reklas"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sinkronisasi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
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
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sts"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/suratpemberitahuan"
	suratdpemberitahuanetail "github.com/bapenda-kota-malang/apin-backend/internal/models/suratpemberitahuan/detail"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/undanganpemeriksaan"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/hargareferensi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jenisperolehan"
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
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tempatpembayaran"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajakpbb"
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
		&referensibank.ReferensiBank{},
		&hargareferensi.HargaReferensi{},
		&nik.Nik{},
		&konfigurasipajak.KonfigurasiPajak{},
		&jenisperolehan.JenisPerolehan{},
		&paymentpoint.PaymentPoint{},
		&jenispajak.JenisPajak{},
		&jenisusaha.JenisUsaha{},
		&jenisusaha.JenisUsahaDetail{},
		&tempatpembayaran.TempatPembayaran{},
		&nop.Nop{},
		&jpb.Jpb{},
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
		&bapl.PotensiBapl{},
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
		&reklas.Reklas{},
		&njoptkpflag.NjoptkpFlag{},
		&njoptkpflag.NjoptkpFlagDetail{},
		&kelasbangunan.KelasBangunan{},
		&kelastanah.KelasTanah{},
		&dbkbjpb2.DbkbJpb2{},
		&dbkbjpb3.DbkbJpb3{},
		&dbkbjpb4.DbkbJpb4{},
		&dbkbjpb5.DbkbJpb5{},
		&dbkbjpb6.DbkbJpb6{},
		&dbkbjpb7.DbkbJpb7{},
		&dbkbjpb8.DbkbJpb8{},
		&dbkbjpb9.DbkbJpb9{},
		&dbkbjpb12.DbkbJpb12{},
		&dbkbjpb13.DbkbJpb13{},
		&dbkbjpb14.DbkbJpb14{},
		&dbkbjpb15.DbkbJpb15{},
		&dbkbjpb16.DbkbJpb16{},
		&dbkbmezanin.DbkbMezanin{},
		&datapetablok.DataPetaBlok{},
		&datanir.DataNir{},
		&datapetaznt.DataPetaZnt{},
		&dbkbfasum.DbkbFasum{},
		&nondep.DbkbFasumNonDep{},
		&depminmax.DbkbFasumDepMinMax{},
		&depjpbklsbintang.DbkbFasumDepJpbKlsBintang{},
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

	listModelPengajuan := []interface{}{
		&pengurangan.Pengurangan{},
		&keberatan.Keberatan{},
		&baplpengajuan.PengajuanBapl{},
		&keberatan.PembetulanKeberatan{},
	}
	a.AutoMigrate(listModelPengajuan...)

	listModelPembayaran := []interface{}{
		&sts.Sts{},
		&sspd.Sspd{},
		&sspd.SspdDetail{},
		&sinkronisasi.Sinkronisasi{},
		&sinkronisasi.SinkronisasiDetail{},
		&sts.StsDetail{},
		&sts.SumberDanaSts{},
	}
	a.AutoMigrate(listModelPembayaran...)

	listModelSppt := []interface{}{
		&sppt.Sppt{},
		&sppt.SpptObjekBersama{},
		&sppt.SpptSimulasi{},
	}
	a.AutoMigrate(listModelSppt...)

	listModelPelayanan := []interface{}{
		&pstpermohonan.PstPermohonan{},
		&pstpermohonan.PstDetail{},
		&pstpermohonan.PstDataOPBaru{},
		&pstpermohonan.PstPermohonanPengurangan{},
	}
	a.AutoMigrate(listModelPelayanan...)

	listModelWajibPajakPBB := []interface{}{
		&wajibpajakpbb.WajibPajakPbb{},
	}
	a.AutoMigrate(listModelWajibPajakPBB...)

	listModelSpop := []interface{}{
		&objekpajakpbb.ObjekPajakPbb{},
		&objekpajakbumi.ObjekPajakBumi{},
		&indukobjekpajak.IndukObjekPajak{},
		&anggotaobjekpajak.AnggotaObjekPajak{},
		&kunjungankembali.KunjunganKembali{},
	}
	a.AutoMigrate(listModelSpop...)

	listModelRegSpop := []interface{}{
		&regobjekpajakpbb.RegObjekPajakPbb{},
		&regobjekpajakbumi.RegObjekPajakBumi{},
		&regindukobjekpajak.RegIndukObjekPajak{},
		&reganggotaobjekpajak.RegAnggotaObjekPajak{},
		&regkunjungankembali.RegKunjunganKembali{},
	}
	a.AutoMigrate(listModelRegSpop...)

	listModelLspop := []interface{}{
		&objekpajakbangunan.ObjekPajakBangunan{},
		&objekpajakbangunan.Jpb2{},
		&objekpajakbangunan.Jpb3{},
		&objekpajakbangunan.Jpb4{},
		&objekpajakbangunan.Jpb5{},
		&objekpajakbangunan.Jpb6{},
		&objekpajakbangunan.Jpb7{},
		&objekpajakbangunan.Jpb8{},
		&objekpajakbangunan.Jpb9{},
		&objekpajakbangunan.Jpb12{},
		&objekpajakbangunan.Jpb13{},
		&objekpajakbangunan.Jpb14{},
		&objekpajakbangunan.Jpb15{},
		&objekpajakbangunan.Jpb16{},
	}
	a.AutoMigrate(listModelLspop...)

	listModelBphtbSptpd := []interface{}{
		&bphtb.BphtbSptpd{},
		&bphtb.Lampiran{},
	}
	a.AutoMigrate(listModelBphtbSptpd...)

	listModelPenagihan := []interface{}{
		&suratpemberitahuan.SuratPemberitahuan{},
		&suratdpemberitahuanetail.SuratPemberitahuanDetail{},
		&bapenagihan.BaPenagihan{},
		&bapenagihanpetugas.BaPenagihanPetugas{},
		&undanganpemeriksaan.UndanganPemeriksaan{},
	}
	a.AutoMigrate(listModelPenagihan...)
}
