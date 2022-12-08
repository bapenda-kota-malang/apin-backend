package bphtbsptpd

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/bphtb"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
)

type Crud struct{}

func (c Crud) Create(w http.ResponseWriter, r *http.Request) {
	var data m.RequestSptpd
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.Create(data, nil)
	hh.DataResponse(w, result, err)
}

func (c Crud) GetList(w http.ResponseWriter, r *http.Request) {
	var input m.RequestSptpd
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetList(input)
	hh.DataResponse(w, result, err)
}

func GetListVerifikasi(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authInfo, ok := ctx.Value("authInfo").(*auth.AuthInfo)
	if !ok {
		hj.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	tp := hh.ValidateString(w, r, "tp")
	if tp == "" {
		return
	}

	var input m.RequestSptpd
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetListApproval(input, authInfo.Jabatan_Id, tp)
	hh.DataResponse(w, result, err)
}

func (c Crud) GetListPpat(w http.ResponseWriter, r *http.Request) {
	var input m.RequestSptpd
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}

	result, err := s.GetListPpat(input)
	hh.DataResponse(w, result, err)
}

func (c Crud) GetDetailPpat(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetailPpat(id)
	hh.DataResponse(w, result, err)
}

func (c Crud) GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetail(id)
	hh.DataResponse(w, result, err)
}

func Approval(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authInfo, ok := ctx.Value("authInfo").(*auth.AuthInfo)
	if !ok {
		hj.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	kd := hh.ValidateString(w, r, "kd")
	if kd == "" {
		return
	}

	var data m.RequestApprovalSptpd
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.Approval(id, kd, authInfo.Jabatan_Id, data, nil)
	hh.DataResponse(w, result, err)
}

func (c Crud) Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.RequestSptpd
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.Update(id, data, nil)
	hh.DataResponse(w, result, err)
}

func (c Crud) VerifyPpat(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
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
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.Delete(id, nil)
	hh.DataResponse(w, result, err)
}
