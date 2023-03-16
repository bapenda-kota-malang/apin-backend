package adbmigration

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/anggaran"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/anggotaobjekpajak"
	adm "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/bankpersepsi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/banktunggal"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/bapenagihan"
	bapenagihanpetugas "github.com/bapenda-kota-malang/apin-backend/internal/models/bapenagihan/petugas"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/baplpengajuan"
	bphtbjenislaporan "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/jenislaporan"
	bphtb "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/datanir"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/datapetablok"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/datapetaznt"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dataznt"
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
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dokument"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailespthiburan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailespthotel"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptparkir"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptppjnonpln"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptppjpln"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptresto"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/geojson"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/hargadasarair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/indukobjekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jabatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jaminanbongkar"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jaminanbongkar/prosesjambong"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jenispajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jenisppj"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jenisusaha"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jpb"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jurnal"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/kanwil"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/keberatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/kelasbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/kelastanah"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/klasifikasijalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/konfigurasipajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/kppbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/kunjungan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/kunjungankembali"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/menu"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/nilaiindividu"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/njoptkp"
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
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensinarahubung"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensipemilikwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/referensibank"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/referensibuku"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/refumum"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/reganggotaobjekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/regkunjungankembali"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbumi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakpbb"
	regpstpermohonan "github.com/bapenda-kota-malang/apin-backend/internal/models/regpelayanan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/reklas"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sinkronisasi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sksk"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	spptpembayaran "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt/pembayaran"
	sppttandaterima "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt/tandaterima"
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
	"github.com/bapenda-kota-malang/apin-backend/internal/models/targetrealisasi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarif"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tempatpembayaranspptmasal"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/undanganpemeriksaan"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/bidangkerja"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/hargareferensi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jenisperolehan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/nik"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/penetapan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/penilaianpbb"
	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/regnpwpd"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	skpd "github.com/bapenda-kota-malang/apin-backend/internal/models/satuankerja"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sektor"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sumberdana"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifjambong"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifjambongrek"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifreklame"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tempatpembayaran"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	mut "github.com/bapenda-kota-malang/apin-backend/internal/models/usertoken"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajakpbb"
)

func GetModelList() (data []interface{}) {
	listModelConfigurationReference := []interface{}{
		&skpd.SatuanKerja{},
		&bidangkerja.BidangKerja{},
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
		&bphtbjenislaporan.BphtbJenisLaporan{},
	}
	data = append(data, listModelConfigurationReference...)

	listModelManagementUser := []interface{}{
		&mu.User{},
		&mut.UserToken{},
		// &group.Group{},
		&menu.Menu{},
		&pegawai.Pegawai{},
		&ppat.Ppat{},
		&wajibpajak.WajibPajak{},
	}
	data = append(data, listModelManagementUser...)

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
	data = append(data, listModelRegNpwpd...)

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
	data = append(data, listModelNpwpd...)

	listModelPendataan := []interface{}{
		// &potensiopwp.PotensiOp{},
		&bapl.PotensiBapl{},
		&detailpotensiop.DetailPotensiOp{},
		&potensipemilikwp.PotensiPemilikWp{},
		&potensinarahubung.PotensiNarahubung{},
		&potensiopwp.DetailPotensiAirTanah{},
		&potensiopwp.DetailPotensiHiburan{},
		&potensiopwp.DetailPotensiHotel{},
		&potensiopwp.DetailPotensiPPJNonPLN{},
		&potensiopwp.DetailPotensiParkir{},
		&potensiopwp.DetailPotensiReklame{},
		&potensiopwp.DetailPotensiResto{},
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
		&dataznt.DataZnt{},
		&dokument.Dokument{},
		&dbkbfasum.DbkbFasum{},
		&nondep.DbkbFasumNonDep{},
		&depminmax.DbkbFasumDepMinMax{},
		&depjpbklsbintang.DbkbFasumDepJpbKlsBintang{},
	}
	data = append(data, listModelPendataan...)

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

	data = append(data, listModelPenetapan...)

	listModelPengajuan := []interface{}{
		&pengurangan.Pengurangan{},
		&pengurangan.PenguranganDendaADM{},
		&pengurangan.PenguranganJPB{},
		&pengurangan.PenguranganPST{},
		&pengurangan.PenguranganPermanen{},
		&keberatan.Keberatan{},
		&baplpengajuan.PengajuanBapl{},
		&keberatan.PembetulanKeberatan{},
		&keberatan.KeputusanKeberatanPbb{},
	}
	data = append(data, listModelPengajuan...)

	listModelPembayaran := []interface{}{
		&sts.Sts{},
		&sspd.Sspd{},
		&sspd.SspdDetail{},
		&sinkronisasi.Sinkronisasi{},
		&sinkronisasi.SinkronisasiDetail{},
		&sts.StsDetail{},
		&sts.SumberDanaSts{},
		&targetrealisasi.TargetRealisasi{},
	}
	data = append(data, listModelPembayaran...)

	listModelSppt := []interface{}{
		&sppt.Sppt{},
		&sppt.SpptObjekBersama{},
		&sppt.SpptSimulasi{},
		&spptpembayaran.SpptPembayaran{},
		&sppttandaterima.SpptTandaTerima{},
	}
	data = append(data, listModelSppt...)

	listModelPelayanan := []interface{}{
		&pstpermohonan.PstPermohonan{},
		&pstpermohonan.PstDetail{},
		&pstpermohonan.PstDataOPBaru{},
		&pstpermohonan.PstPermohonanPengurangan{},
		// &pstpermohonan.KeputusanKeberatanPbb{},
		&pstpermohonan.PstLogApproval{},
		&pstpermohonan.PembatalanSppt{},
		&pstpermohonan.PembetulanSpptSKPSTP{},
		&pstpermohonan.PstLampiran{},
		&pstpermohonan.SPMKP{},
	}
	data = append(data, listModelPelayanan...)

	listModelRegPelayanan := []interface{}{
		&regpstpermohonan.RegPstPermohonan{},
		&regpstpermohonan.RegPstDetail{},
		&regpstpermohonan.RegPstDataOPBaru{},
		&regpstpermohonan.RegPstPermohonanPengurangan{},
		&regpstpermohonan.RegKeputusanKeberatanPbb{},
		&regpstpermohonan.RegPembatalanSppt{},
		&regpstpermohonan.RegPembetulanSpptSKPSTP{},
		&regpstpermohonan.RegPstLampiran{},
		&regpstpermohonan.RegSPMKP{},
	}
	data = append(data, listModelRegPelayanan...)

	listModelWajibPajakPBB := []interface{}{
		&wajibpajakpbb.WajibPajakPbb{},
	}
	data = append(data, listModelWajibPajakPBB...)

	listModelSpop := []interface{}{
		&objekpajakpbb.ObjekPajakPbb{},
		&objekpajakbumi.ObjekPajakBumi{},
		&anggotaobjekpajak.AnggotaObjekPajak{},
		&kunjungankembali.KunjunganKembali{},
	}
	data = append(data, listModelSpop...)

	listModelRegSpop := []interface{}{
		&regobjekpajakpbb.RegObjekPajakPbb{},
		&regobjekpajakbumi.RegObjekPajakBumi{},
		&reganggotaobjekpajak.RegAnggotaObjekPajak{},
		&regkunjungankembali.RegKunjunganKembali{},
	}
	data = append(data, listModelRegSpop...)

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
		&fasilitasbangunan.FasilitasBangunan{},
	}
	data = append(data, listModelLspop...)

	listModelRegLspop := []interface{}{
		&regobjekpajakbangunan.RegObjekPajakBangunan{},
		&regobjekpajakbangunan.RegJpb2{},
		&regobjekpajakbangunan.RegJpb3{},
		&regobjekpajakbangunan.RegJpb4{},
		&regobjekpajakbangunan.RegJpb5{},
		&regobjekpajakbangunan.RegJpb6{},
		&regobjekpajakbangunan.RegJpb7{},
		&regobjekpajakbangunan.RegJpb8{},
		&regobjekpajakbangunan.RegJpb9{},
		&regobjekpajakbangunan.RegJpb12{},
		&regobjekpajakbangunan.RegJpb13{},
		&regobjekpajakbangunan.RegJpb14{},
		&regobjekpajakbangunan.RegJpb15{},
		&regobjekpajakbangunan.RegJpb16{},
		&regfasilitasbangunan.RegFasilitasBangunan{},
	}
	data = append(data, listModelRegLspop...)

	listModelBphtbSptpd := []interface{}{
		&bphtb.BphtbSptpd{},
		&bphtb.Lampiran{},
	}
	data = append(data, listModelBphtbSptpd...)

	listModelPenagihan := []interface{}{
		&suratpemberitahuan.SuratPemberitahuan{},
		&suratdpemberitahuanetail.SuratPemberitahuanDetail{},
		&bapenagihan.BaPenagihan{},
		&bapenagihanpetugas.BaPenagihanPetugas{},
		&undanganpemeriksaan.UndanganPemeriksaan{},
	}
	data = append(data, listModelPenagihan...)

	listModelGeoJson := []interface{}{
		&geojson.GeoJson{},
	}
	data = append(data, listModelGeoJson...)

	listModelKunjungan := []interface{}{
		&kunjungan.Kunjungan{},
		&kunjungan.KunjunganDetail{},
	}
	data = append(data, listModelKunjungan...)

	listModelSkSk := []interface{}{
		&sksk.SkSk{},
	}
	data = append(data, listModelSkSk...)

	listModelNilaiIndividu := []interface{}{
		&nilaiindividu.NilaiIndividu{},
	}
	data = append(data, listModelNilaiIndividu...)

	listModelJaminanBongkar := []interface{}{
		&jaminanbongkar.JaminanBongkar{},
		&jaminanbongkar.DetailJambong{},
		&prosesjambong.ProsesJambong{},
	}
	data = append(data, listModelJaminanBongkar...)

	listModelPenilaian := []interface{}{
		&penilaianpbb.RangePenyusutan{},
		&penilaianpbb.Penyusutan{},
		&penilaianpbb.Fasilitas{},
		&penilaianpbb.FasNonDep{},
		&penilaianpbb.FasDepMinMax{},
		&penilaianpbb.FasDepJpbKlsBintang{},
		&penilaianpbb.DbkbMaterial{},
		&penilaianpbb.BangunanLantai{},
		&penilaianpbb.TipeBangunan{},
		&penilaianpbb.KayuUlin{},
		&penilaianpbb.DbkbStandard{},
		&penilaianpbb.DbkbDayaDukung{},
	}
	data = append(data, listModelPenilaian...)

	listModelPenetapanMassal := []interface{}{
		&penetapan.PenerimaKompensasi{},
		&penetapan.SubjekPajakNJOPTKP{},
		&tarif.Tarif{},
		&penetapan.PenguranganPengenaanJPB{},
		&penetapan.PenguranganPermanen{},
		&penetapan.PenguranganPST{},
		&penetapan.SubjekPajak{},
		&penetapan.DafnomOP{},
		&referensibuku.ReferensiBuku{},
		&njoptkp.NJOPTKP{},
		&penetapan.NJKP{},
		&penetapan.ParameterSPPTSTTSDHKP{},
		&penetapan.RefThnNJKPNJOPTKPTarif{},
	}
	data = append(data, listModelPenetapanMassal...)

	listModeTempatPembayaranSpptMasal := []interface{}{
		&tempatpembayaranspptmasal.TempatPembayaranSPPTMasal{},
		&banktunggal.BankTunggal{},
		&bankpersepsi.BankPersepsi{},
		&kanwil.Kanwil{},
		&kppbb.Kppbb{},
	}
	data = append(data, listModeTempatPembayaranSpptMasal...)

	listModelRefUmum := []interface{}{
		&refumum.RefUmum{},
	}
	data = append(data, listModelRefUmum...)

	listModelIndukObjekPajak := []interface{}{
		&indukobjekpajak.IndukObjekPajak{},
	}
	data = append(data, listModelIndukObjekPajak...)

	return data
}
