package sppt

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
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

func UpdateVa(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateVaDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.UpdateVa(r.Context(), id, data, uint64(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

// cetak massal
func DownloadPDF(w http.ResponseWriter, r *http.Request) {
	var data m.PenetapanMassalDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.DownloadPDF(data)
	if err != nil {
		return
	}

	fName := result.(rp.OKSimple).Data.(*s.ResponseFile).FileName

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
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=report-cetak-massal.pdf"))
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Length", strconv.Itoa(len(fileContent)))

	// menulis konten file ke response body
	if _, err := w.Write(fileContent); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
