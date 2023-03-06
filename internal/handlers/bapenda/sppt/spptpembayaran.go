package sppt

import (
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt/pembayaran"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt/pembayaran"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

type CrudSpptPembayaran struct{}

func (c CrudSpptPembayaran) Create(w http.ResponseWriter, r *http.Request) {
	var data m.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	result, err := s.Create(data, authInfo.User_Id)
	hh.DataResponse(w, result, err)
}

func (c CrudSpptPembayaran) GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetList(input)
	hh.DataResponse(w, result, err)
}

func (c CrudSpptPembayaran) GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetail(id)
	hh.DataResponse(w, result, err)
}

func (c CrudSpptPembayaran) Update(w http.ResponseWriter, r *http.Request) {
	hj.WriteJSON(w, http.StatusBadGateway, t.II{
		"message": "not yet implemented",
	}, nil)
	return
}

func (c CrudSpptPembayaran) Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.Delete(id)
	hh.DataResponse(w, result, err)
}
