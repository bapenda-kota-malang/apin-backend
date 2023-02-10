package penetapanmassal

import (
	"net/http"
	"strings"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/penetapanmassal"
)

func Copy(w http.ResponseWriter, r *http.Request) {
	var input m.CopySptDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	result, err := s.CopySpt(input)

	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	// get type params
	types := strings.ToUpper(r.URL.Query().Get("type"))

	// get kode jenis usaha
	kodePajak := r.URL.Query().Get("kode_jenis_usaha")

	result, err := s.GetList(input, types, kodePajak)
	hh.DataResponse(w, result, err)
}
