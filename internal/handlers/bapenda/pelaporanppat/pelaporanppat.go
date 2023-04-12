package pelaporanppat

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	msptpd "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pelaporanppat"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/pelaporanppat"
)

type Crud struct{}

func (c Crud) Create(w http.ResponseWriter, r *http.Request) {
	var data m.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	// authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.Create(data, nil)
	hh.DataResponse(w, result, err)
}

func (c Crud) GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetList(input)
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

func GetListLaporanPPAT(w http.ResponseWriter, r *http.Request) {
	var data msptpd.FilterPPATDto
	if !hh.ValidateStructByURL(w, *r.URL, &data) {
		return
	}

	result, err := s.GetListTransaksiPPAT(data)
	hh.DataResponse(w, result, err)
}

func GetDetailLaporanPPAT(w http.ResponseWriter, r *http.Request) {
	var data msptpd.FilterPPATDto
	if !hh.ValidateStructByURL(w, *r.URL, &data) {
		return
	}

	result, err := s.GetDetailTransaksiPPAT(data)
	hh.DataResponse(w, result, err)
}

func GetDetailLapPPAT(w http.ResponseWriter, r *http.Request) {
	ppat := hh.ValidateString(w, r, "ppat")
	if ppat == "" {
		return
	}
	result, err := s.GetDetailbyField("Ppat_Id", ppat)
	hh.DataResponse(w, result, err)
}

func (c Crud) Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}
	// authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.Update(id, data, nil)
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
