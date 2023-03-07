package bapenda

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	rh "github.com/bapenda-kota-malang/apin-backend/pkg/routerhelper"
	"github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/bapenagihan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/datanir"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/datapetablok"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/datapetaznt"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dataznt"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbfasum"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbfasum/depjpbklsbintang"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbfasum/depminmax"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbfasum/nondep"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jaminanbongkar"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jaminanbongkar/prosesjambong"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/nilaiindividu"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/penetapanmassal"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/profile"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/regnpwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/suratpemberitahuan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/anggaran"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/baplpengajuan"
	bphtbsptpd "github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/bphtb"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/configuration/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/daerah"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbjpb12"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbjpb13"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbjpb14"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbjpb15"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbjpb16"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbjpb2"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbjpb3"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbjpb4"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbjpb5"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbjpb6"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbjpb7"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbjpb8"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbjpb9"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/dbkbmezanin"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/espt"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/geojson"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/group"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/hargadasarair"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/hargareferensi"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/home"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jabatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jenispajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jenisperolehan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jenisppj"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jenisusaha"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jpb"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jurnal"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/keberatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kecamatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kelasbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kelastanah"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kelurahan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/klasifikasijalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/konfigurasipajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kunjungan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/menu"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/nik"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/njoptkpflag"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/nop"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/objekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/objekpajakbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/objekpajakbumi"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/objekpajakpbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/omset"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/pangkat"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/paymentpoint"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/pegawai"
	permohonan "github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/pelayanan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/penagihan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/pengurangan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/potensiopwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/ppat"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/provinsi"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/referensibank"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/regobjekpajakbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/regobjekpajakpbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/reklas"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/satuankerja"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sektor"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sinkronisasi"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sksk"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sppt"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/spt"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sspd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sts"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sumberdana"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/targetrealisasi"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tarifjambong"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tarifjambongrek"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tarifpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tarifreklame"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tempatpembayaran"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/undanganpemeriksaan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/user"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/wajibpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/wajibpajakpbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"

	_ "github.com/bapenda-kota-malang/apin-backend/internal/models/adbmigration"
)

func SetRoutes() http.Handler {
	// Config
	auth.SkipAuhPaths = []string{
		"/assets/*",
		"/auth/login",
		"/auth/logout",
		"/account/reset-password",
	}
	auth.Position = 1

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(auth.GuardMW)

	r.NotFound(er.NotFoundResponse)
	r.MethodNotAllowed(er.MethodNotAllowedResponse)

	r.Get("/", home.Index)

	fs := http.FileServer(http.Dir(servicehelper.GetResourcesPath()))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	fsAssets := http.FileServer(http.Dir(servicehelper.GetAssetsPath()))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fsAssets))
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", auth.Login)
	})

	r.Route("/account", func(r chi.Router) {
		// r.Post("/register", account.Create) // replaced withr register
		r.Get("/check", account.Check)
		r.Patch("/change-password", account.ChangePassword)
		r.Post("/reset-password", account.RequestResetPassword)
		r.Get("/reset-password", account.CheckResetPassword)
		r.Patch("/reset-password", account.ResetPassword)
	})

	r.Get("/profile", profile.GetDetail)

	rh.RegCrud(r, "/menu", menu.Crud{})

	rh.RegCrud(r, "/satuankerja", satuankerja.Crud{})

	rh.RegCrud(r, "/jabatan", jabatan.Crud{})

	rh.RegCrud(r, "/pangkat", pangkat.Crud{})

	rh.RegCrud(r, "/target-realisasi", targetrealisasi.Crud{})
	r.Patch("/target-realisasi/schedule", targetrealisasi.UpdateBySchedule)

	rh.RegCrud(r, "/sektor", sektor.Crud{})

	rh.RegCrud(r, "/omset", omset.Crud{})

	rh.RegCrud(r, "/anggaran", anggaran.Crud{})

	rh.RegCrud(r, "/sumberdana", sumberdana.Crud{})

	rh.RegCrud(r, "/jurnal", jurnal.Crud{})

	rh.RegCrud(r, "/jenisppj", jenisppj.Crud{})

	rh.RegCrud(r, "/tarifjambongrek", tarifjambongrek.Crud{})

	rh.RegCrud(r, "/tarifjambong", tarifjambong.Crud{})

	rh.RegCrud(r, "/tarifreklame", tarifreklame.Crud{})

	rh.RegCrud(r, "/klasifikasijalan", klasifikasijalan.Crud{})

	rh.RegCrud(r, "/jalan", jalan.Crud{})

	rh.RegCrud(r, "/hargadasarair", hargadasarair.Crud{})

	r.Get("/hargadasarair/peruntukan", hargadasarair.GetPeruntukan)

	rh.RegCrud(r, "/tarifpajak", tarifpajak.Crud{})

	rh.RegCrud(r, "/rekening", rekening.Crud{})

	rh.RegCrud(r, "/jpb", jpb.Crud{})
	rh.RegCrud(r, "/objekpajak", objekpajak.Crud{})

	rh.RegCrud(r, "/referensibank", referensibank.Crud{})

	rh.RegCrud(r, "/hargareferensi", hargareferensi.Crud{})

	rh.RegCrud(r, "/nik", nik.Crud{})

	rh.RegCrud(r, "/reklas", reklas.Crud{})

	rh.RegCrud(r, "/konfigurasipajak", konfigurasipajak.Crud{})

	rh.RegCrud(r, "/jenisperolehan", jenisperolehan.Crud{})

	rh.RegCrud(r, "/paymentpoint", paymentpoint.Crud{})

	rh.RegCrud(r, "/jenispajak", jenispajak.Crud{})

	rh.RegCrud(r, "/jenisusaha", jenisusaha.Crud{})

	rh.RegCrud(r, "/tempatpembayaran", tempatpembayaran.Crud{})

	r.Patch("/sppt/cetakdaftartagihan", sppt.CetakDaftarTagihan)
	r.Patch("/sppt/penilaian", sppt.Penilaian)
	rh.RegCrud(r, "/sppt", sppt.Crud{})

	r.Post("/sksk/cetak", sksk.Cetak)
	rh.RegCrud(r, "/sksk", sksk.Crud{})

	rh.RegCrud(r, "/bphtbsptpd", bphtbsptpd.Crud{})

	rh.RegCrud(r, "/kelastanah", kelastanah.Crud{})
	r.Get("/kelastanah/download/excel", kelastanah.DownloadExcelList)

	rh.RegCrud(r, "/kelasbangunan", kelasbangunan.Crud{})
	r.Get("/kelasbangunan/download/excel", kelasbangunan.DownloadExcelList)

	r.Get("/nop-bynopstr/{nop}", nop.GetDetailByNopString)
	rh.RegCrud(r, "/nop", nop.Crud{})

	rh.RegCrud(r, "/dbkbjpb2", dbkbjpb2.Crud{})
	r.Get("/dbkbjpb2/download/excel", dbkbjpb2.DownloadExcelList)

	rh.RegCrud(r, "/dbkbjpb3", dbkbjpb3.Crud{})
	r.Get("/dbkbjpb3/download/excel", dbkbjpb3.DownloadExcelList)

	rh.RegCrud(r, "/dbkbjpb4", dbkbjpb4.Crud{})
	r.Get("/dbkbjpb4/download/excel", dbkbjpb4.DownloadExcelList)

	rh.RegCrud(r, "/dbkbjpb5", dbkbjpb5.Crud{})
	r.Get("/dbkbjpb5/download/excel", dbkbjpb5.DownloadExcelList)

	rh.RegCrud(r, "/dbkbjpb6", dbkbjpb6.Crud{})
	r.Get("/dbkbjpb6/download/excel", dbkbjpb6.DownloadExcelList)

	rh.RegCrud(r, "/dbkbjpb7", dbkbjpb7.Crud{})
	r.Get("/dbkbjpb7/download/excel", dbkbjpb7.DownloadExcelList)

	rh.RegCrud(r, "/dbkbjpb8", dbkbjpb8.Crud{})
	r.Get("/dbkbjpb8/download/excel", dbkbjpb8.DownloadExcelList)

	rh.RegCrud(r, "/dbkbjpb9", dbkbjpb9.Crud{})
	r.Get("/dbkbjpb9/download/excel", dbkbjpb9.DownloadExcelList)

	rh.RegCrud(r, "/dbkbjpb12", dbkbjpb12.Crud{})
	r.Get("/dbkbjpb12/download/excel", dbkbjpb12.DownloadExcelList)

	rh.RegCrud(r, "/dbkbjpb13", dbkbjpb13.Crud{})
	r.Get("/dbkbjpb13/download/excel", dbkbjpb13.DownloadExcelList)

	rh.RegCrud(r, "/dbkbjpb14", dbkbjpb14.Crud{})
	r.Get("/dbkbjpb14/download/excel", dbkbjpb14.DownloadExcelList)

	rh.RegCrud(r, "/dbkbjpb15", dbkbjpb15.Crud{})
	r.Get("/dbkbjpb15/download/excel", dbkbjpb15.DownloadExcelList)

	rh.RegCrud(r, "/dbkbjpb16", dbkbjpb16.Crud{})
	r.Get("/dbkbjpb16/download/excel", dbkbjpb16.DownloadExcelList)

	rh.RegCrud(r, "/dbkbmezanin", dbkbmezanin.Crud{})
	r.Get("/dbkbmezanin/download/excel", dbkbmezanin.DownloadExcelList)

	rh.RegCrud(r, "/nilaiindividu", nilaiindividu.Crud{})

	rh.RegCrud(r, "/geojson-data", geojson.Crud{})

	r.Route("/geojson", func(r chi.Router) {
		r.Get("/", geojson.GetListGeoJson)
		r.Get("/{id}", geojson.GetDetailGeoJson)
	})

	r.Route("/bphtbsptpd-approval", func(r chi.Router) {
		r.Get("/{tp}", bphtbsptpd.GetListVerifikasi)
		r.Get("/{tp}/download/excel", bphtbsptpd.DownloadExcelListVerifikasi)
		r.Patch("/{id}/{kd}", bphtbsptpd.Approval)
	})

	r.Route("/kelastanah-kode", func(r chi.Router) {
		r.Get("/{kd}", kelastanah.GetDetailByCode)
	})

	r.Route("/kelasbangunan-kode", func(r chi.Router) {
		r.Get("/{kd}", kelasbangunan.GetDetailByCode)
	})

	r.Route("/spptsimulasi-process", func(r chi.Router) {
		r.Post("/{flag}", sppt.GetSimulasiByNop)
	})

	r.Delete("/kunjungan-detail/{id}", kunjungan.DeleteDetail)
	rh.RegCrud(r, "/kunjungan", kunjungan.Crud{})

	r.Route("/sppt-simulasi", func(r chi.Router) {
		r.Post("/", sppt.CreateSimulasi)
		r.Get("/", sppt.GetListSimulasi)
		r.Get("/{id}", sppt.GetDetailSimulasi)
		r.Patch("/{id}", sppt.UpdateSimulasi)
		r.Delete("/{id}", sppt.DeleteSimulasi)
	})

	r.Route("/njoptkpflag", func(r chi.Router) {
		r.Post("/", njoptkpflag.Create)
		r.Get("/", njoptkpflag.GetList)
		r.Get("/{id}", njoptkpflag.GetDetail)
		r.Patch("/{id}", njoptkpflag.Update)
		r.Delete("/{id}", njoptkpflag.Delete)
	})

	r.Route("/pegawai", func(r chi.Router) {
		r.Post("/", pegawai.Create)
		r.Get("/", pegawai.GetList)
		r.Get("/{id}", pegawai.GetDetail)
		r.Patch("/{id}", pegawai.Update)
		r.Delete("/{id}", pegawai.Delete)
	})

	r.Route("/permohonan", func(r chi.Router) {
		r.Patch("/{id}/status", permohonan.UpdateStatus)
		r.Post("/", permohonan.Create)
		r.Get("/", permohonan.GetList)
		r.Get("/{id}", permohonan.GetDetail)
		r.Patch("/{id}", permohonan.Update)
		r.Delete("/{id}", permohonan.Delete)
		r.Get("/download/excel", permohonan.DownloadExcelList)
		r.Get("/download/pdf/{id}", permohonan.DownloadPdf)
	})

	r.Route("/statnop", func(r chi.Router) {
		r.Get("/{id}", permohonan.GetStatusNOP)
	})

	r.Route("/ppat", func(r chi.Router) {
		r.Post("/", ppat.Create)
		r.Get("/", ppat.GetList)
		r.Get("/{id}", ppat.GetDetail)
		r.Patch("/{id}", ppat.Update)
		r.Delete("/{id}", ppat.Delete)
	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/", user.Create)
		r.Get("/", user.GetList)
		r.Get("/{id}", user.GetDetail)
		r.Patch("/{id}", user.Update)
		r.Delete("/{id}", user.Delete)
		r.Patch("/{id}/verifikasi", user.Verifikasi)
		r.Get("/download/excel", user.DownloadExcelList)
	})

	r.Route("/group", func(r chi.Router) {
		r.Post("/", group.Create)
		r.Get("/", group.GetList)
		r.Get("/{id}", group.GetDetail)
		r.Patch("/{id}", group.Update)
		r.Delete("/{id}", group.Delete)
		r.Get("/download/excel", group.DownloadExcelList)
	})

	r.Get("/npwpd-byno/{no}", npwpd.GetDetailByNoNPWPD)
	r.Route("/npwpd", func(r chi.Router) {
		r.Get("/", npwpd.GetList)
		r.Get("/{id}", npwpd.GetDetail)
		r.Post("/", npwpd.Create)
		r.Patch("/{id}", npwpd.Update)
		r.Delete("/{id}", npwpd.Delete)
		r.Get("/download/excel", npwpd.DownloadExcelList)
		r.Get("/download/excel/{type}", npwpd.DownloadExcelSpec)
	})

	r.Route("/regnpwpd", func(r chi.Router) {
		r.Patch("/{id}/setverifystatus", regnpwpd.VerifyRegistrasiNpwpd)
		r.Get("/", regnpwpd.GetList)
		r.Get("/{id}", regnpwpd.GetDetail)
		r.Delete("/{id}", regnpwpd.Delete)
		r.Patch("/{id}", regnpwpd.Update)
		r.Get("/download/excel", regnpwpd.DownloadExcelList)
	})

	r.Route("/potensiopwp", func(r chi.Router) {
		r.Post("/", potensiopwp.Create)
		r.Get("/", potensiopwp.GetList)
		r.Get("/{id}", potensiopwp.GetDetail)
		r.Patch("/{id}", potensiopwp.Update)
		r.Delete("/{id}", potensiopwp.Delete)
		r.Get("/download/excel", potensiopwp.DownloadExcelList)
	})

	r.Route("/provinsi", func(r chi.Router) {
		r.Get("/{id}/kode", provinsi.GetDetailByCode)
		r.Post("/", provinsi.Create)
		r.Get("/", provinsi.GetList)
		r.Get("/{id}", provinsi.GetDetail)
		r.Patch("/{id}", provinsi.Update)
		r.Delete("/{id}", provinsi.Delete)
	})

	r.Route("/penagihan", func(r chi.Router) {
		r.Get("/himbauan", penagihan.GetReportHimbauan)
		r.Get("/tunggakan", penagihan.GetReportTunggakan)
		r.Get("/himpunan", penagihan.GetReportTunggakanHimpunan)
	})

	r.Route("/daerah", func(r chi.Router) {
		r.Get("/{id}/kode", daerah.GetDetailByCode)
		r.Post("/", daerah.Create)
		r.Get("/", daerah.GetList)
		r.Get("/{id}", daerah.GetDetail)
		r.Patch("/{id}", daerah.Update)
		r.Delete("/{id}", daerah.Delete)
	})

	r.Route("/kecamatan", func(r chi.Router) {
		r.Get("/{id}/kode", kecamatan.GetDetailByCode)
		r.Post("/", kecamatan.Create)
		r.Get("/", kecamatan.GetList)
		r.Get("/{id}", kecamatan.GetDetail)
		r.Patch("/{id}", kecamatan.Update)
		r.Delete("/{id}", kecamatan.Delete)
	})

	r.Route("/kelurahan", func(r chi.Router) {
		r.Get("/{id}/kode", kelurahan.GetDetailByCode)
		r.Post("/", kelurahan.Create)
		r.Get("/", kelurahan.GetList)
		r.Get("/{id}", kelurahan.GetDetail)
		r.Patch("/{id}", kelurahan.Update)
		r.Delete("/{id}", kelurahan.Delete)
	})

	r.Route("/wajibpajak", func(r chi.Router) {
		r.Get("/", wajibpajak.GetList)
		r.Get("/{id}", wajibpajak.GetDetail)
	})

	r.Route("/wajibpajakpbb", func(r chi.Router) {
		r.Get("/", wajibpajakpbb.GetList)
		r.Get("/{id}", wajibpajakpbb.GetDetail)
	})

	r.Route("/sptpd", func(r chi.Router) {
		r.Post("/", spt.Create)
		r.Get("/", spt.GetList)
		r.Get("/{id}", spt.GetDetail)
		r.Patch("/{id}", spt.Update)
		r.Delete("/{id}", spt.Delete)
	})

	r.Route("/skpd", func(r chi.Router) {
		r.Post("/", spt.Create)
		r.Get("/", spt.GetList)
		r.Get("/{id}", spt.GetDetail)
		r.Patch("/{id}", spt.Update)
		r.Patch("/{id}/verify", spt.Verify)
		r.Delete("/{id}", spt.Delete)
	})

	r.Route("/skpdkb", func(r chi.Router) {
		r.Get("/", spt.SkpdkbGetList)
		r.Post("/existing/{type}", spt.SkpdkbExisting)
		r.Post("/new/{type}", spt.SkpdNew)
	})

	// route for espt list data, verify espt, and get detail data for espt before verify
	r.Route("/espt", func(r chi.Router) {
		r.Get("/", espt.GetList)
		r.Get("/{id}", espt.GetDetail)
		r.Patch("/{id}/verify", espt.Verify)
		r.Get("/download/excel", espt.DownloadExcelList)
	})

	r.Route("/sspd", func(r chi.Router) {
		r.Get("/", sspd.GetList)
		r.Get("/{id}", sspd.GetDetail)
		r.Post("/", sspd.Create)
		r.Patch("/{id}", sspd.Update)
		r.Delete("/{id}", sspd.Delete)
		r.Patch("/{id}/cancel", sspd.Cancel)
		r.Get("/sspddetail", sspd.GetListSspdDetail)
		r.Get("/download/excel", sspd.DownloadExcelList)
	})

	r.Route("/sinkronisasi", func(r chi.Router) {
		r.Get("/", sinkronisasi.GetList)
		r.Get("/{id}", sinkronisasi.GetDetail)
		r.Post("/", sinkronisasi.Create)
		r.Post("/detail", sinkronisasi.CreateDetail)
		r.Patch("/detail", sinkronisasi.Update)
	})

	r.Route("/sts", func(r chi.Router) {
		r.Get("/", sts.GetList)
		r.Get("/{id}", sts.GetDetail)
		r.Post("/", sts.Create)
		r.Patch("/{id}", sts.Update)
		r.Delete("/{id}", sts.Delete)
		r.Get("/download/excel", sts.DownloadExcelList)
	})

	r.Route("/undanganpemeriksaan", func(r chi.Router) {
		r.Get("/", undanganpemeriksaan.GetList)
		r.Get("/{id}", undanganpemeriksaan.GetDetail)
		r.Post("/", undanganpemeriksaan.Create)
		r.Patch("/{id}", undanganpemeriksaan.Update)
		r.Patch("/updatestatusterbit", undanganpemeriksaan.UpdateStatusTerbit)
		r.Delete("/{id}", undanganpemeriksaan.Delete)
		r.Get("/download/excel", undanganpemeriksaan.DownloadExcelList)
	})

	r.Route("/suratpemberitahuan", func(r chi.Router) {
		r.Get("/", suratpemberitahuan.GetList)
		r.Get("/{id}", suratpemberitahuan.GetDetail)
		r.Post("/", suratpemberitahuan.CreateSchedule)
		r.Patch("/", suratpemberitahuan.UpdateBulk)
		r.Patch("/{id}", suratpemberitahuan.UpdateSingle)
		r.Delete("/{id}", suratpemberitahuan.Delete)
		r.Get("/download/excel", suratpemberitahuan.DownloadExcelList)
	})

	r.Route("/bapenagihan", func(r chi.Router) {
		r.Get("/", bapenagihan.GetList)
		r.Get("/{id}", bapenagihan.GetDetail)
		r.Post("/", bapenagihan.Create)
		r.Patch("/{id}", bapenagihan.Update)
		r.Patch("/verify/{id}", bapenagihan.Verify)
		r.Delete("/{id}", bapenagihan.Delete)
		r.Get("/download/excel", bapenagihan.DownloadExcelList)
	})

	r.Route("/pengurangan", func(r chi.Router) {
		r.Get("/", pengurangan.GetList)
		r.Get("/{id}", pengurangan.GetDetail)
		r.Post("/", pengurangan.Create)
		r.Patch("/verify/{id}", pengurangan.Verify)
	})

	r.Route("/keberatan", func(r chi.Router) {
		r.Get("/", keberatan.GetList)
		r.Get("/{id}", keberatan.GetDetail)
		r.Post("/", keberatan.Create)
		r.Patch("/verify/{id}", keberatan.Verify)
		r.Get("/download/excel", keberatan.DownloadExcelList)
	})

	r.Route("/baplpengajuan", func(r chi.Router) {
		r.Post("/", baplpengajuan.Create)
		r.Get("/", baplpengajuan.GetList)
		r.Get("/{id}", baplpengajuan.GetDetail)
		r.Delete("/{id}", baplpengajuan.Delete)
	})

	r.Post("/datapetablok/bulk", datapetablok.CreateBulk)
	rh.RegCrud(r, "/datapetablok", datapetablok.Crud{})

	r.Post("/datanir/bulk", datanir.CreateBulk)
	r.Get("/datanir-dokument/{no}", datanir.GetDokByNoDok)
	rh.RegCrud(r, "/datanir", datanir.Crud{})

	r.Post("/datapetaznt/bulk", datapetaznt.CreateBulk)
	rh.RegCrud(r, "/datapetaznt", datapetaznt.Crud{})

	r.Post("/dataznt/bulk", dataznt.CreateBulk)
	rh.RegCrud(r, "/dataznt", dataznt.Crud{})

	rh.RegCrud(r, "/dbkbfasum", dbkbfasum.Crud{})
	rh.RegCrud(r, "/dbkbfasum/nondep", nondep.Crud{})
	rh.RegCrud(r, "/dbkbfasum/depminmax", depminmax.Crud{})
	rh.RegCrud(r, "/dbkbfasum/depjpbklsbintang", depjpbklsbintang.Crud{})

	r.Post("/objekpajakbumi/bulk", objekpajakbumi.UpdateBulk)
	rh.RegCrud(r, "/objekpajakbumi", objekpajakbumi.Crud{})

	r.Route("/objekpajakpbb", func(r chi.Router) {
		r.Post("/", objekpajakpbb.Create)
		r.Get("/", objekpajakpbb.GetList)
		r.Get("/{id}", objekpajakpbb.GetDetail)
		r.Get("/download/excel", objekpajakpbb.DownloadExcelList)
	})

	r.Route("/regobjekpajakpbb", func(r chi.Router) {
		r.Get("/", regobjekpajakpbb.GetList)
		r.Get("/{id}", regobjekpajakpbb.GetDetail)
	})

	r.Route("/objekpajakbangunan", func(r chi.Router) {
		r.Post("/", objekpajakbangunan.Create)
		r.Get("/", objekpajakbangunan.GetList)
		r.Get("/{id}", objekpajakbangunan.GetDetail)
	})

	r.Route("/regobjekpajakbangunan", func(r chi.Router) {
		r.Get("/", regobjekpajakbangunan.GetList)
		r.Get("/{id}", regobjekpajakbangunan.GetDetail)
		r.Patch("/verify/{id}", regobjekpajakbangunan.VerifyLspop)
	})

	rh.RegCrud(r, "/jaminanbongkar", jaminanbongkar.Crud{})

	rh.RegCrud(r, "/prosesjambong", prosesjambong.Crud{})

	rh.RegCrud(r, "/penetapan-massal", penetapanmassal.Crud{})

	return r
}
