package bphtbsptpd

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	"github.com/go-chi/chi/v5"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/bphtb"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
)

type Crud struct{}

type ctxKey string

const ctxKeyFrom ctxKey = "from"

// middleware specific for create handler, this middleware will add from value to context request with key from
func (c Crud) CreateMw(next http.Handler, from string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ctxKeyFrom, from)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (c Crud) Create(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()
	fromCtx := reqCtx.Value(ctxKeyFrom)
	from := ""
	if fromCtx != nil {
		from = fromCtx.(string)
	}
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	var data m.CreateDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.Create(data, from, authInfo)
	hh.DataResponse(w, result, err)
}

func (c Crud) GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}

	result, err := s.GetList(input, -1, "")
	hh.DataResponse(w, result, err)
}

func GetListVerifikasi(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}

	authInfo, ok := r.Context().Value("authInfo").(*auth.AuthInfo)
	if !ok {
		hj.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	tp := hh.ValidateString(w, r, "tp")
	if tp == "" {
		return
	}

	result, err := s.GetList(input, authInfo.Jabatan_Id, tp)
	hh.DataResponse(w, result, err)
}

func DownloadExcelListVerifikasi(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}

	authInfo, ok := r.Context().Value("authInfo").(*auth.AuthInfo)
	if !ok {
		hj.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	tp := hh.ValidateString(w, r, "tp")
	if tp == "" {
		return
	}

	result, err := s.DownloadExcelListVerifikasi(input, authInfo.Jabatan_Id, tp)
	if err != nil {
		return
	}

	name := "verifikasi_bphtb"
	if tp == "val" {
		name = "validasi_bphtb"
	} else if tp == "byr" {
		name = "laporan_sspd_bphtb"
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=list_%s.xlsx", name))
	w.Header().Set("Content-Transfer-Encoding", "binary")
	result.Write(w)
}

func (c Crud) GetDetail(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	result, err := s.GetDetail(id)
	hh.DataResponse(w, result, err)
}

func GetDetailPembayaran(w http.ResponseWriter, r *http.Request) {
	no := hh.ValidateString(w, r, "no")
	if no == "" {
		return
	}
	no = strings.Replace(no, "_", ".", -1)

	result, err := s.GetDetailbyField("NoDokumen", no)
	hh.DataResponse(w, result, err)
}

func GetListTransaksiPPAT(w http.ResponseWriter, r *http.Request) {
	var data m.FilterPPATDto
	if !hh.ValidateStructByURL(w, *r.URL, &data) {
		return
	}

	result, err := s.GetListTransaksiPPAT(data)
	hh.DataResponse(w, result, err)
}

func GetDetailTransaksiPPAT(w http.ResponseWriter, r *http.Request) {
	ppat := hh.ValidateString(w, r, "ppat")
	if ppat == "" {
		return
	}

	result, err := s.GetDetailbyField("Ppat_Id", ppat)
	hh.DataResponse(w, result, err)
}

func Approval(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authInfo, ok := ctx.Value("authInfo").(*auth.AuthInfo)
	if !ok {
		hj.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	kd := hh.ValidateString(w, r, "kd")
	if kd == "" {
		return
	}

	var data m.RequestApprovalSptpd
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.Approval(id, kd, authInfo.Jabatan_Id, data, nil)
	hh.DataResponse(w, result, err)
}

func (c Crud) Update(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	var data m.CreateDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.Update(id, data, nil)
	hh.DataResponse(w, result, err)
}

func (c Crud) VerifyPpat(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	var data m.VerifyPpatDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	result, err := s.VerifyPpat(id, data, authInfo.User_Id, nil)
	hh.DataResponse(w, result, err)
}

func (c Crud) Delete(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	result, err := s.Delete(id, nil)
	hh.DataResponse(w, result, err)
}
