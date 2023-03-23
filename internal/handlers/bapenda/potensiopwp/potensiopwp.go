package potensiopwp

import ( 
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv" 
	"strings"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp"
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
		baseDto := s.DPBaseUsecase{CreateDto: data.CreateDto}
		input = s.NewDPAirTanahUsecase(baseDto, data.DetailPajakDtos)
	case "hiburan":
		var data m.CreateDtoHiburan
		if !hh.ValidateStructByIOR(w, r.Body, &data) {
			return
		}
		baseDto := s.DPBaseUsecase{CreateDto: data.CreateDto}
		input = s.NewDPHiburanUsecase(baseDto, data.DetailPajakDtos)
	case "hotel":
		var data m.CreateDtoHotel
		if !hh.ValidateStructByIOR(w, r.Body, &data) {
			return
		}
		baseDto := s.DPBaseUsecase{CreateDto: data.CreateDto}
		input = s.NewDPHotelUsecase(baseDto, data.DetailPajakDtos)
	case "parkir":
		var data m.CreateDtoParkir
		if !hh.ValidateStructByIOR(w, r.Body, &data) {
			return
		}
		baseDto := s.DPBaseUsecase{CreateDto: data.CreateDto}
		input = s.NewDPParkirUsecase(baseDto, data.DetailPajakDtos)
	case "ppjnonpln":
		var data m.CreateDtoPPJNonPLN
		if !hh.ValidateStructByIOR(w, r.Body, &data) {
			return
		}
		baseDto := s.DPBaseUsecase{CreateDto: data.CreateDto}
		input = s.NewDPPPJNonPLNUsecase(baseDto, data.DetailPajakDtos)
	case "reklame":
		var data m.CreateDtoReklame
		if !hh.ValidateStructByIOR(w, r.Body, &data) {
			return
		}
		baseDto := s.DPBaseUsecase{CreateDto: data.CreateDto}
		input = s.NewDPReklameUsecase(baseDto, data.DetailPajakDtos)
	case "resto":
		var data m.CreateDtoResto
		if !hh.ValidateStructByIOR(w, r.Body, &data) {
			return
		}
		baseDto := s.DPBaseUsecase{CreateDto: data.CreateDto}
		input = s.NewDPRestoUsecase(baseDto, data.DetailPajakDtos)
	default:
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: "category tidak diketahui"}, nil)
		return
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

	result, err := s.CreateTrx(input, authInfo.User_Id)
	hh.DataResponse(w, result, err)
}

// func Update(w http.ResponseWriter, r *http.Request) {
// 	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
// 	if !pass {
// 		return
// 	}

// 	var data m.UpdateDto
// 	if !hh.ValidateStructByIOR(w, r.Body, &data) {
// 		return
// 	}

// 	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

// 	result, err := s.UpdateTrx(id, data, uint(authInfo.User_Id))
// 	hh.DataResponse(w, result, err)
// }

func Delete(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	result, err := s.Delete(id)
	hh.DataResponse(w, result, err)
}

func DownloadExcelList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.DownloadExcelList(input)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=list_daftar_potensi_opwp_baru.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	result.Write(w)
} 

func DownloadPdf(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	result, err := s.DownloadPdf(id)
	if err != nil {
		return
	}

	fName := result.(rp.OKSimple).Data.(*s.ResponseFile).FileName
	if fName == "" {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := os.Open(fName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// membaca konten file
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set header response
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=report-permohonan-%d.pdf", id))
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Length", strconv.Itoa(len(fileContent)))

	// menulis konten file ke response body
	if _, err := w.Write(fileContent); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// selesaikan response
	w.WriteHeader(http.StatusOK)

}  
