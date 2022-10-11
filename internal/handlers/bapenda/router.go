package bapenda

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	rh "github.com/bapenda-kota-malang/apin-backend/pkg/routerhelper"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/anggaran"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/auth"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/configuration/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/daerah"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/group"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/hargadasarair"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/home"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jabatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/jurnal"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kecamatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kelurahan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/klasifikasijalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/menu"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/omset"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/pangkat"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/potensiopwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/ppat"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/provinsi"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sektor"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/skpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/spt"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/sumberdana"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tarifjambong"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tarifjambongrek"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tarifpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/tarifreklame"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/user"
	wajibPajak "github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/wajibpajak"
)

func SetRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.NotFound(er.NotFoundResponse)
	r.MethodNotAllowed(er.MethodNotAllowedResponse)

	r.Get("/", home.Index)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", auth.Login)
	})

	r.Route("/account", func(r chi.Router) {
		r.Patch("/reset-password", account.ResetPassword)
		r.Patch("/change-password", account.ChangePassword)
	})

	rh.RegCrud(r, "/menu", menu.Crud{})

	rh.RegCrud(r, "/skpd", skpd.Crud{})

	rh.RegCrud(r, "/jabatan", jabatan.Crud{})

	rh.RegCrud(r, "/pangkat", pangkat.Crud{})

	rh.RegCrud(r, "/sektor", sektor.Crud{})

	rh.RegCrud(r, "/omset", omset.Crud{})

	rh.RegCrud(r, "/anggaran", anggaran.Crud{})

	rh.RegCrud(r, "/sumberdana", sumberdana.Crud{})

	rh.RegCrud(r, "/jurnal", jurnal.Crud{})

	rh.RegCrud(r, "/tarifjambongrek", tarifjambongrek.Crud{})

	rh.RegCrud(r, "/tarifjambong", tarifjambong.Crud{})

	rh.RegCrud(r, "/tarifreklame", tarifreklame.Crud{})

	rh.RegCrud(r, "/klasifikasijalan", klasifikasijalan.Crud{})

	rh.RegCrud(r, "/jalan", jalan.Crud{})

	rh.RegCrud(r, "/hargadasarair", hargadasarair.Crud{})

	rh.RegCrud(r, "/tarifpajak", tarifpajak.Crud{})

	rh.RegCrud(r, "/rekening", rekening.Crud{})

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

	// r.Route("/rekening", func(r chi.Router) {
	// 	r.Get("/", rekening.GetList)
	// })

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
		r.Get("/", wajibPajak.GetList)
		r.Get("/{id}", wajibPajak.GetDetail)
	})

	r.Route("/spt", func(r chi.Router) {
		r.Post("/{category}", spt.Create)
		r.Get("/", spt.GetList)
		r.Get("/{id}", spt.GetDetail)
		r.Patch("/{id}/{category}", spt.Update)
		r.Delete("/{id}", spt.Delete)
	})

	return r
}
