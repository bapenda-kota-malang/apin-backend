package bapenda

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	rh "github.com/bapenda-kota-malang/apin-backend/pkg/routerhelper"
	"github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/bapenagihan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/regnpwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/suratpemberitahuan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/anggaran"
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
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jurnal"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kecamatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kelasbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kelastanah"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kelurahan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/klasifikasijalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/konfigurasipajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/menu"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/nik"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/njoptkpflag"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/nop"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/objekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/omset"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/pangkat"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/paymentpoint"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/potensiopwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/ppat"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/provinsi"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/referensibank"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/reklas"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sektor"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sinkronisasi"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/skpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/spt"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sspd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sts"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sumberdana"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tarifjambong"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tarifjambongrek"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tarifpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tarifreklame"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tempatpembayaran"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/undanganpemeriksaan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/user"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/wajibpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"

	_ "github.com/bapenda-kota-malang/apin-backend/internal/models/adbmigration"
)

func SetRoutes() http.Handler {
	// Config
	auth.SkipAuhPaths = []string{
		"/auth/login",
		"/auth/logout",
		"/account/reset-password",
		"/account/change-password",
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

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", auth.Login)
	})

	r.Route("/account", func(r chi.Router) {
		// r.Post("/register", account.Create) // replaced withr register
		r.Get("/check", account.Check)
		r.Patch("/reset-password", account.ResetPassword)
		r.Patch("/change-password", account.ChangePassword)
	})

	rh.RegCrud(r, "/menu", menu.Crud{})

	rh.RegCrud(r, "/satuankerja", skpd.Crud{})

	rh.RegCrud(r, "/jabatan", jabatan.Crud{})

	rh.RegCrud(r, "/pangkat", pangkat.Crud{})

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

	rh.RegCrud(r, "/kelastanah", kelastanah.Crud{})

	rh.RegCrud(r, "/kelasbangunan", kelasbangunan.Crud{})

	rh.RegCrud(r, "/nop", nop.Crud{})

	rh.RegCrud(r, "/dbkbjpb2", dbkbjpb2.Crud{})

	rh.RegCrud(r, "/dbkbjpb3", dbkbjpb3.Crud{})

	rh.RegCrud(r, "/dbkbjpb4", dbkbjpb4.Crud{})

	rh.RegCrud(r, "/dbkbjpb5", dbkbjpb5.Crud{})

	rh.RegCrud(r, "/dbkbjpb6", dbkbjpb6.Crud{})

	rh.RegCrud(r, "/dbkbjpb7", dbkbjpb7.Crud{})

	rh.RegCrud(r, "/dbkbjpb8", dbkbjpb8.Crud{})

	rh.RegCrud(r, "/dbkbjpb9", dbkbjpb9.Crud{})

	rh.RegCrud(r, "/dbkbjpb12", dbkbjpb12.Crud{})

	rh.RegCrud(r, "/dbkbjpb13", dbkbjpb13.Crud{})

	rh.RegCrud(r, "/dbkbjpb14", dbkbjpb14.Crud{})

	rh.RegCrud(r, "/dbkbjpb15", dbkbjpb15.Crud{})

	rh.RegCrud(r, "/dbkbjpb16", dbkbjpb16.Crud{})

	rh.RegCrud(r, "/dbkbmezanin", dbkbmezanin.Crud{})

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
	})

	r.Route("/group", func(r chi.Router) {
		r.Post("/", group.Create)
		r.Get("/", group.GetList)
		r.Get("/{id}", group.GetDetail)
		r.Patch("/{id}", group.Update)
		r.Delete("/{id}", group.Delete)
	})

	r.Route("/npwpd", func(r chi.Router) {
		r.Get("/", npwpd.GetList)
		r.Get("/{id}", npwpd.GetDetail)
		r.Post("/", npwpd.Create)
		r.Patch("/{id}", npwpd.Update)
		r.Delete("/{id}", npwpd.Delete)

	})

	r.Route("/regnpwpd", func(r chi.Router) {
		r.Patch("/{id}/setverifystatus", regnpwpd.VerifyRegistrasiNpwpd)
		r.Get("/", regnpwpd.GetList)
		r.Get("/{id}", regnpwpd.GetDetail)
		r.Delete("/{id}", regnpwpd.Delete)
		r.Patch("/{id}", regnpwpd.Update)
	})

	r.Route("/potensiopwp", func(r chi.Router) {
		r.Post("/", potensiopwp.Create)
		r.Get("/", potensiopwp.GetList)
		r.Get("/{id}", potensiopwp.GetDetail)
		r.Patch("/{id}", potensiopwp.Update)
		r.Delete("/{id}", potensiopwp.Delete)
	})

	r.Route("/provinsi", func(r chi.Router) {
		r.Post("/", provinsi.Create)
		r.Get("/", provinsi.GetList)
		r.Get("/{id}", provinsi.GetDetail)
		r.Patch("/{id}", provinsi.Update)
		r.Delete("/{id}", provinsi.Delete)
	})

	r.Route("/daerah", func(r chi.Router) {
		r.Post("/", daerah.Create)
		r.Get("/", daerah.GetList)
		r.Get("/{id}", daerah.GetDetail)
		r.Patch("/{id}", daerah.Update)
		r.Delete("/{id}", daerah.Delete)
	})

	r.Route("/kecamatan", func(r chi.Router) {
		r.Post("/", kecamatan.Create)
		r.Get("/", kecamatan.GetList)
		r.Get("/{id}", kecamatan.GetDetail)
		r.Patch("/{id}", kecamatan.Update)
		r.Delete("/{id}", kecamatan.Delete)
	})

	r.Route("/kelurahan", func(r chi.Router) {
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
		r.Post("/existing/{type}", spt.SkpdkbExisting)
		r.Post("/new/{type}", spt.SkpdNew)
	})

	// route for espt list data, verify espt, and get detail data for espt before verify
	r.Route("/espt", func(r chi.Router) {
		r.Get("/", espt.GetList)
		r.Get("/{id}", espt.GetDetail)
		r.Patch("/{id}/verify", espt.Verify)
	})

	r.Route("/sspd", func(r chi.Router) {
		r.Get("/", sspd.GetList)
		r.Get("/{id}", sspd.GetDetail)
		r.Post("/", sspd.Create)
		r.Patch("/{id}", sspd.Update)
		r.Delete("/{id}", sspd.Delete)
		r.Patch("/{id}/cancel", sspd.Cancel)
		r.Get("/sspddetail", sspd.GetListSspdDetail)
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
	})

	r.Route("/undanganpemeriksaan", func(r chi.Router) {
		r.Get("/", undanganpemeriksaan.GetList)
		r.Get("/{id}", undanganpemeriksaan.GetDetail)
		r.Post("/", undanganpemeriksaan.Create)
		r.Patch("/{id}", undanganpemeriksaan.Update)
		r.Patch("/updatestatusterbit", undanganpemeriksaan.UpdateStatusTerbit)
		r.Delete("/{id}", undanganpemeriksaan.Delete)
	})

	r.Route("/suratpemberitahuan", func(r chi.Router) {
		r.Get("/", suratpemberitahuan.GetList)
		r.Get("/{id}", suratpemberitahuan.GetDetail)
		r.Post("/", suratpemberitahuan.CreateSchedule)
		r.Patch("/", suratpemberitahuan.UpdateBulk)
		r.Patch("/{id}", suratpemberitahuan.UpdateSingle)
		r.Delete("/{id}", suratpemberitahuan.Delete)
	})

	r.Route("/bapenagihan", func(r chi.Router) {
		r.Get("/", bapenagihan.GetList)
		r.Get("/{id}", bapenagihan.GetDetail)
		r.Post("/", bapenagihan.Create)
		r.Patch("/{id}", bapenagihan.Update)
		r.Patch("/verify/{id}", bapenagihan.Verify)
		r.Delete("/{id}", bapenagihan.Delete)
	})
	return r
}
