package spt

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/spt"
	sdsa "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptair"
	sdsh "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailspthotel"
	sdsp "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptparkir"
	sdsrek "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptreklame"
	sdsres "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/detailsptresto"

	// gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func Create(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	fmt.Println("category: ", category)
	switch category {
	case "air":
		var register m.CreateAirDto
		if hh.ValidateStructByIOR(w, r.Body, &register) == false {
			return
		}

		result, err := sdsa.Create(register)
		hh.DataResponse(w, result, err)
	case "hotel":
		var register m.CreateHotelDto
		if hh.ValidateStructByIOR(w, r.Body, &register) == false {
			return
		}

		result, err := sdsh.Create(register)
		hh.DataResponse(w, result, err)
	case "parkir":
		var register m.CreateParkirDto
		if hh.ValidateStructByIOR(w, r.Body, &register) == false {
			return
		}

		result, err := sdsp.Create(register)
		hh.DataResponse(w, result, err)
	case "reklame":
		var register m.CreateReklameDto
		if hh.ValidateStructByIOR(w, r.Body, &register) == false {
			return
		}

		result, err := sdsrek.Create(register)
		hh.DataResponse(w, result, err)
	case "resto":
		var register m.CreateRestoDto
		if hh.ValidateStructByIOR(w, r.Body, &register) == false {
			return
		}

		result, err := sdsres.Create(register)
		hh.DataResponse(w, result, err)
	}
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetList(input)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetail(id)
	hh.DataResponse(w, result, err)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.Delete(id)
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}
	category := chi.URLParam(r, "category")
	fmt.Println("category: ", category)
	switch category {
	case "air":
		var input m.UpdateAirDto
		if hh.ValidateStructByIOR(w, r.Body, &input) == false {
			return
		}

		result, err := sdsa.Update(id, input)
		hh.DataResponse(w, result, err)
		// case "hotel":
		// 	var input m.UpdateHotelDto
		// 	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		// 		return
		// 	}

		// 	result, err := sdsh.Create(input)
		// 	hh.DataResponse(w, result, err)
		// case "parkir":
		// 	var input m.UpdateParkirDto
		// 	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		// 		return
		// 	}

		// 	result, err := sdsp.Update(input)
		// 	hh.DataResponse(w, result, err)
		// case "reklame":
		// 	var input m.UpdateReklameDto
		// 	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		// 		return
		// 	}

		// 	result, err := sdsrek.Update(input)
		// 	hh.DataResponse(w, result, err)
		// case "resto":
		// 	var input m.UpdateRestoDto
		// 	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		// 		return
		// 	}

		// 	result, err := sdsres.Update(input)
		// 	hh.DataResponse(w, result, err)
		// case "jaminan":
		// 	var input m.UpdateJaminanDto
		// 	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		// 		return
		// 	}

		// 	result, err := sjbr.Update(input)
		// 	hh.DataResponse(w, result, err)
	}

}
