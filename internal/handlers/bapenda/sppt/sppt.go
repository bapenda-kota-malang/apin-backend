package sppt

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt"
)

type Crud struct{}

func (c Crud) Create(w http.ResponseWriter, r *http.Request) {
	var data m.RequestDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.Create(data)
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

func (c Crud) Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.RequestDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.Update(id, data)
	hh.DataResponse(w, result, err)
}

func (c Crud) Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.Delete(id)
	hh.DataResponse(w, result, err)
}

func Penilaian(w http.ResponseWriter, r *http.Request) {
	var data m.PenilaianDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.Penilaian(data)
	hh.DataResponse(w, result, err)
}

func CetakDaftarTagihan(w http.ResponseWriter, r *http.Request) {
	var data m.GetDaftarTagihan
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.CetakDaftarTagihan(data)
	hh.DataResponse(w, result, err)
}

func PenetapanMassal(w http.ResponseWriter, r *http.Request) {
	var data m.PenetapanMassalDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.PenetapanMassal(data)
	hh.DataResponse(w, result, err)
}

func Rincian(w http.ResponseWriter, r *http.Request) {
	var data m.NopDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.Rincian(data)
	hh.DataResponse(w, result, err)
}

func Salinan(w http.ResponseWriter, r *http.Request) {
	var data m.SalinanDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.Salinan(data)
	hh.DataResponse(w, result, err)
}

func GetRekapitulasi(w http.ResponseWriter, r *http.Request) {
	var data m.RekapitulasiOpRequest
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.GetRekapitulasi(data)
	hh.DataResponse(w, result, err)
}
