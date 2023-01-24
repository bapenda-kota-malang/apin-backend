package wajibpajak

import (
	"net/http"

	bphtbsptpd "github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/bphtb"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/bphtbjenislaporan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/configuration/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/daerah"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kecamatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kelurahan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/nop"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/ppat"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/provinsi"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/espt"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/hargadasarair"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/home"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/jenisppj"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/keberatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/noppbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/npwpd"
	pelayanan "github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/pelayanan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/pengurangan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/profile"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/regobjekpajakbangunan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/regobjekpajakpbb"
	permohonan "github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/regpelayanan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/regwajibpajakpbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/sppt"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/sppt/objekbersama"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/sspd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/static"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/regnpwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/spt"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/tarifpajak"
	rh "github.com/bapenda-kota-malang/apin-backend/pkg/routerhelper"
	"github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetRoutes() http.Handler {
	// Config
	auth.SkipAuhPaths = []string{
		"/auth/login",
		"/auth/logout",
		"/register",
		"/account/register",
		"/account/reset-password",
		"/provinsi",
		"/daerah",
		"/kecamatan",
		"/kelurahan",
		"/rekening",
	}
	auth.Position = 3

	// Start routing
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(auth.GuardMW)

	r.NotFound(er.NotFoundResponse)
	r.MethodNotAllowed(er.MethodNotAllowedResponse)

	r.Get("/", home.Index)
	r.Route("/register", func(r chi.Router) {
		r.Post("/", profile.Create)
		r.Post("/checker/{id}", profile.Checker)
	})

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

	r.Route("/profile", func(r chi.Router) {
		r.Get("/", profile.GetDetail) // EXECUTE
		r.Patch("/", profile.Update)
	})

	r.Route("/provinsi", func(r chi.Router) {
		r.Get("/", provinsi.GetList)
		r.Get("/{id}", provinsi.GetDetail)
	})

	r.Route("/daerah", func(r chi.Router) {
		r.Get("/", daerah.GetList)
		r.Get("/{id}", daerah.GetDetail)
	})

	r.Route("/kecamatan", func(r chi.Router) {
		r.Get("/", kecamatan.GetList)
		r.Get("/{id}", kecamatan.GetDetail)
	})

	r.Route("/kelurahan", func(r chi.Router) {
		r.Get("/", kelurahan.GetList)
		r.Get("/{id}", kelurahan.GetDetail)
	})

	r.Route("/regnpwpd", func(r chi.Router) {
		r.Post("/", regnpwpd.Create)
		r.Get("/", regnpwpd.GetListForWp)
		r.Get("/{id}", regnpwpd.GetDetailForWp)
		r.Patch("/{id}", regnpwpd.UpdateForWp)
		r.Delete("/{id}", regnpwpd.DeleteForWp)
	})

	r.Route("/espt", func(r chi.Router) {
		r.Post("/", espt.Create)
		r.Get("/", espt.GetList)
		r.Get("/{id}", espt.GetDetail)
		r.Patch("/{id}", espt.Update)
		r.Delete("/{id}", espt.Delete)
	})

	rh.RegCrud(r, "/rekening", rekening.Crud{})

	r.Route("/npwpd", func(r chi.Router) {
		r.Get("/", npwpd.GetList)
		r.Get("/{id}", npwpd.GetDetail)
		r.Patch("/{id}", npwpd.Update)
		r.Patch("/{id}/{category}", npwpd.UpdatePhoto)
		r.Delete("/{id}/{category}/{filename}", npwpd.DeletePhoto)
	})

	r.Route("/jenisppj", func(r chi.Router) {
		r.Get("/", jenisppj.GetList)
		r.Get("/{id}", jenisppj.GetDetail)
	})

	r.Route("/tarifpajak", func(r chi.Router) {
		r.Get("/", tarifpajak.GetList)
	})

	r.Route("/hargadasarair", func(r chi.Router) {
		r.Get("/", hargadasarair.GetList)
		r.Get("/peruntukan", hargadasarair.GetPeruntukan)
	})

	r.Route("/sptpd", func(r chi.Router) {
		r.Get("/", spt.GetList)
		r.Get("/{id}", spt.GetDetail)
	})

	r.Route("/skpd", func(r chi.Router) {
		r.Get("/", spt.GetList)
		r.Get("/{id}", spt.GetDetail)
	})

	r.Route("/pengurangan", func(r chi.Router) {
		r.Post("/", pengurangan.Create)
	})

	r.Route("/keberatan", func(r chi.Router) {
		r.Post("/", keberatan.Create)
	})

	fs := http.FileServer(http.Dir(servicehelper.GetResourcesPath()))
	r.Handle("/static/{filename}", http.StripPrefix("/static/", static.AuthFile(static.JoinPrefix(fs))))

	r.Route("/pelayanan", func(r chi.Router) {
		r.Get("/{nop}", pelayanan.GetDetailbyNop)
	})

	r.Route("/regobjekpajakpbb", func(r chi.Router) {
		r.Post("/", regobjekpajakpbb.CreateMw(http.HandlerFunc(regobjekpajakpbb.Create), "wp"))
		r.Get("/{id}", regobjekpajakpbb.GetDetail)
		r.Get("/", regobjekpajakpbb.GetListNop)
	})

	r.Route("/regwajibpajakpbb", func(r chi.Router) {
		r.Post("/", regwajibpajakpbb.Create)
		r.Get("/{id}", regwajibpajakpbb.GetDetail)
	})

	r.Route("/regobjekpajakbangunan", func(r chi.Router) {
		r.Post("/", regobjekpajakbangunan.Create)
	})

	r.Route("/bphtbsptpd", func(r chi.Router) {
		bphtbsptpdCrud := bphtbsptpd.Crud{}
		r.Post("/", bphtbsptpdCrud.CreateMw(http.HandlerFunc(bphtbsptpdCrud.Create), "wp"))
		r.Get("/", bphtbsptpdCrud.GetList)
		r.Get("/{id}", bphtbsptpdCrud.GetDetail)
		r.Get("/jenislaporan", bphtbjenislaporan.Crud{}.GetList)
	})

	r.Route("/ppat", func(r chi.Router) {
		r.Get("/", ppat.GetList)
	})

	rh.RegCrud(r, "/sppt", sppt.Crud{})
	r.Route("/sppt-nop", func(r chi.Router) {
		r.Get("/", sppt.GetListNop)
	})

	r.Route("/njop", func(r chi.Router) {
		r.Get("/", sppt.Crud{}.GetList)
		r.Get("/objekbersama", objekbersama.Crud{}.GetList)
	})

	r.Route("/nop", func(r chi.Router) {
		nopCrud := nop.Crud{}
		r.Get("/", nopCrud.GetList)
		r.Get("/{id}", nopCrud.GetDetail)
	})

	r.Route("/noppbb", func(r chi.Router) {
		r.Get("/", noppbb.GetList)
		r.Get("/{id}", noppbb.GetDetail)
	})

	r.Route("/permohonan", func(r chi.Router) {
		r.Patch("/{id}/status", permohonan.UpdateStatus)
		r.Post("/", permohonan.Create)
		r.Get("/", permohonan.GetList)
		r.Get("/{id}", permohonan.GetDetail)
		r.Patch("/{id}", permohonan.Update)
		r.Delete("/{id}", permohonan.Delete)
	})

	r.Route("/logpayment", func(r chi.Router) {
		r.Get("/{npwpd}", sspd.GetList)
	})

	return r
}
