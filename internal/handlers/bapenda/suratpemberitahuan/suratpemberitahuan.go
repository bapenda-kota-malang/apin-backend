package suratpemberitahuan

import (
	"fmt"
	"net/http"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	"github.com/go-chi/chi/v5"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/suratpemberitahuan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/suratpemberitahuan"
)

func CreateSchedule(w http.ResponseWriter, r *http.Request) {
	result, err := s.TrxSchedule(nil)
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}

	result, err := s.GetList(input, nil)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	result, err := s.GetDetail(id, nil)
	hh.DataResponse(w, result, err)
}

func UpdateBulk(w http.ResponseWriter, r *http.Request) {
	var datas m.UpdateBulkDto
	if !hh.ValidateStructByIOR(w, r.Body, &datas) {
		return
	}

	result, err := s.UpdateBulk(datas, nil)
	hh.DataResponse(w, result, err)
}

func UpdateSingle(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	var data m.UpdateDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.Update(id, data, nil)
	hh.DataResponse(w, result, err)
}

func Cetak(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	var data m.UpdateDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.Cetak(id, data)
	if err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	dokumen := result.(rp.OK).Data.(m.SuratPemberitahuan).Dokumen

	http.Redirect(w, r, fmt.Sprintf("http://%s/static/pdf/%s", r.Host, dokumen), http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}
	result, err := s.Delete(id, nil)
	hh.DataResponse(w, result, err)
}
