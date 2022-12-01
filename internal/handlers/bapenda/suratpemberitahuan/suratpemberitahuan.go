package suratpemberitahuan

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

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
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
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
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.Update(id, data, nil)
	hh.DataResponse(w, result, err)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}
	result, err := s.Delete(id, nil)
	hh.DataResponse(w, result, err)
}
