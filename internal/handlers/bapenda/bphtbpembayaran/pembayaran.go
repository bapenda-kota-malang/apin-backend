package bphtbpembayaran

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/bphtb"
)

type Crud struct{}

type ctxKey string

const ctxKeyFrom ctxKey = "from"

func (c Crud) Create(w http.ResponseWriter, r *http.Request) {
	// reqCtx := r.Context()
	// fromCtx := reqCtx.Value(ctxKeyFrom)
	// from := ""
	// if fromCtx != nil {
	// 	from = fromCtx.(string)
	// }
	// authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	var data m.RequestPembayaranBphtb
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.CreatePembayaran(data, nil)
	hh.DataResponse(w, result, err)
}

func (c Crud) GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}

	result, err := s.GetListPembayaran(input)
	hh.DataResponse(w, result, err)
}

func (c Crud) GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetailPembayaran(id)
	hh.DataResponse(w, result, err)
}

func (c Crud) Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.RequestPembayaranBphtb
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.UpdatePembayaran(id, data, nil)
	hh.DataResponse(w, result, err)
}

func (c Crud) Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.DeletePembayaran(id, nil)
	hh.DataResponse(w, result, err)
}
