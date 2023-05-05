package objekpajakpbb

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakpbb"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var data m.CreateDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.Create(data)
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}

	result, err := s.GetList(input, 0, nil)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetail(id, 0, nil)
	hh.DataResponse(w, result, err)
}

func UpdateRtRwMassal(w http.ResponseWriter, r *http.Request) {
	var data m.UpdateRtRwMassalDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.UpdateRtRwMassal(data)
	hh.DataResponse(w, result, err)
}

func DownloadExcelList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.DownloadExcelList(input, 0)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=list_spop_lspop.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	result.Write(w)
}

func GetNopTerbesar(w http.ResponseWriter, r *http.Request) {
	var data m.ObjekPajakPbb
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.GetNopTerbesar(data)
	hh.DataResponse(w, result, err)
}

func SejarahOp(w http.ResponseWriter, r *http.Request) {
	var data m.SejarahFilterDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.GetSejarahOp(data)
	hh.DataResponse(w, result, err)
}
