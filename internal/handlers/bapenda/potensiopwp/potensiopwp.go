package potensiopwp

import (
	"net/http"
	"strings"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/detailpotensiairtanah"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/detailpotensihotel"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	"github.com/go-chi/chi/v5"
)

// GetList Data Potensi Objek Pajak with pagination
func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}
	result, err := s.GetList(input)
	hh.DataResponse(w, result, err)
}

// Return data detail Objek Pajak
func GetDetail(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	result, err := s.GetDetail(id)
	hh.DataResponse(w, result, err)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var input m.Input

	switch strings.ToLower(r.URL.Query().Get("category")) {
	case "airtanah":
		var data m.CreateDtoAirTanah
		if !hh.ValidateStructByIOR(w, r.Body, &data) {
			return
		}
		input = detailpotensiairtanah.NewDPAirTanahUsecase(s.DPBaseUsecase{CreateDto: data.CreateDto}, data.DetailPajakDtos)
	case "hotel":
		var data m.CreateDtoHotel
		if !hh.ValidateStructByIOR(w, r.Body, &data) {
			return
		}
		input = detailpotensihotel.NewDPHotelUsecase(data)
	default:
		var data m.CreateDto
		if !hh.ValidateStructByIOR(w, r.Body, &data) {
			return
		}
		input = potensiopwp.NewDPBaseUsecase(data)
	}

	switch input.GetPotensiOp().Golongan {
	case nt.GolonganBadan:
		break
	case nt.GolonganOrangPribadi:
		break
	default:
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: "golongan tidak diketahui"}, nil)
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	result, err := s.CreateTrx(input, uint(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	var data m.UpdateDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	result, err := s.UpdateTrx(id, data, uint(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	result, err := s.Delete(id)
	hh.DataResponse(w, result, err)
}
