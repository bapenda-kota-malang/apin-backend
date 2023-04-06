package bapenagihan

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bapenagihan"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/bapenagihan"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	"github.com/go-chi/chi/v5"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var data m.CreateDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.TrxCreate(data, uint(authInfo.User_Id))
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
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	result, err := s.GetDetail(id, nil)
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
	result, err := s.TrxUpdate(id, data, uint(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

func Verify(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}
	var data m.VerifikasiDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.Verify(id, data, uint64(authInfo.User_Id), nil)
	hh.DataResponse(w, result, err)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	result, err := s.Delete(id, nil)
	hh.DataResponse(w, result, err)
}

func DownloadExcelList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}

	result, err := s.DownloadExcelList(input, nil)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=list_berita_acara_penagihan.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	result.Write(w)
}

// Berita acara penagihan / pemeriksaan
func DownloadPDF(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}
	
	result, err := s.DownloadPDF(id)
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
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=berita-acara-penagihan-pemeriksaan-%d.pdf", id))
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
