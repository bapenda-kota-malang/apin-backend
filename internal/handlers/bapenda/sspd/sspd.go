package sspd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/sspd"
	sd "github.com/bapenda-kota-malang/apin-backend/internal/services/sspd/sspddetail"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var input m.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.Create(input, uint64(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var input m.UpdateDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.Update(id, input, uint64(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetList(input, nil)
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

func Cancel(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var input m.CancelDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	result, err := s.Cancel(input, uint64(id))
	hh.DataResponse(w, result, err)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}
	result, err := s.Delete(uint64(id))
	hh.DataResponse(w, result, err)
}

func GetListSspdDetail(w http.ResponseWriter, r *http.Request) {
	var input m.SspdDetailFilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := sd.GetListSspdDetail(input)
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
	w.Header().Set("Content-Disposition", "attachment; filename=list_sspd.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	result.Write(w)
}

// Surat setoran pajak daerah
func DownloadPDF(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
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
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=surat-setoran-pajak-daerah-%d.pdf", id))
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
