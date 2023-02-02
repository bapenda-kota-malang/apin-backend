package penetapanmassal

import (
	"net/http"
	"strings"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/penetapanmassal"
)

type Crud struct{}

func (c Crud) Create(w http.ResponseWriter, r *http.Request) {

	hh.DataResponse(w, "test", nil)
}

func (c Crud) GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	// get type params
	types := strings.ToUpper(r.URL.Query().Get("type"))

	kode_jenis_usaha := r.URL.Query().Get("kode_jenis_usaha")

	result, err := s.GetList(input, types, kode_jenis_usaha)
	hh.DataResponse(w, result, err)
}

func (c Crud) GetDetail(w http.ResponseWriter, r *http.Request) {
	hh.DataResponse(w, "test", nil)
}

func (c Crud) Update(w http.ResponseWriter, r *http.Request) {
	hh.DataResponse(w, "test", nil)
}

func (c Crud) Delete(w http.ResponseWriter, r *http.Request) {
	hh.DataResponse(w, "test", nil)
}
