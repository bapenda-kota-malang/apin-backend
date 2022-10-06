package bapenda

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	rh "github.com/bapenda-kota-malang/apin-backend/pkg/routerhelper"
	"github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/auth"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/configuration/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/daerah"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/esptd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/group"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/home"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kecamatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kelurahan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/menu"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/pendaftaran"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/potensiopwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/ppat"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/provinsi"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/user"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/wajibpajak"
)

func SetRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.NotFound(er.NotFoundResponse)
	r.MethodNotAllowed(er.MethodNotAllowedResponse)

	r.Get("/", home.Index)

	fs := http.FileServer(http.Dir(servicehelper.GetResourcesPath()))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", auth.Login)
	})

	r.Route("/account", func(r chi.Router) {
		r.Patch("/reset-password", account.ResetPassword)
		r.Patch("/change-password", account.ChangePassword)
	})

	rh.RegCrud(r, "/menu", menu.Crud{})

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

	r.Route("/registration", func(r chi.Router) {
		r.Get("/", pendaftaran.GetList)
		r.Get("/{id}", pendaftaran.GetDetail)
		r.Post("/operator", pendaftaran.RegisterByOperator)
		r.Patch("/{id}", pendaftaran.Update)
		r.Patch("/npwpd/{id}/verifikasi", pendaftaran.VerifyNpwpd)
		r.Delete("/{id}", pendaftaran.Delete)
	})

	r.Route("/rekening", func(r chi.Router) {
		r.Get("/", rekening.GetList)
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

	r.Route("/esptd", func(r chi.Router) {
		r.Get("/", esptd.GetList)
		r.Get("/{id}", esptd.GetDetail)
		r.Delete("/{id}", esptd.Delete)
	})

	return r
}
