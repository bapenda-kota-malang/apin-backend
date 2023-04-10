package pelaporanppat

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	msptpd "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pelaporanppat"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/pelaporanppat"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
)

type Crud struct{}

func (c Crud) Create(w http.ResponseWriter, r *http.Request) {
	var data m.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	// authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.Create(data, nil)
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

func GetListLaporanPPAT(w http.ResponseWriter, r *http.Request) {
	var data msptpd.FilterPPATDto
	if !hh.ValidateStructByURL(w, *r.URL, &data) {
		return
	}

	result, err := s.GetListTransaksiPPAT(data)
	hh.DataResponse(w, result, err)
}

func GetDetailLaporanPPAT(w http.ResponseWriter, r *http.Request) {
	var data msptpd.FilterPPATDto
	if !hh.ValidateStructByURL(w, *r.URL, &data) {
		return
	}

	result, err := s.GetDetailTransaksiPPAT(data)
	hh.DataResponse(w, result, err)
}

func GetDetailLapPPAT(w http.ResponseWriter, r *http.Request) {
	ppat := hh.ValidateString(w, r, "ppat")
	if ppat == "" {
		return
	}
	result, err := s.GetDetailbyField("Ppat_Id", ppat)
	hh.DataResponse(w, result, err)
}

func (c Crud) Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}
	// authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.Update(id, data, nil)
	hh.DataResponse(w, result, err)
}

func (c Crud) Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.Delete(id, nil)
	hh.DataResponse(w, result, err)
}

// DownloadExcelLaporanPPAT func
func DownloadExcelLaporanPPAT(w http.ResponseWriter, r *http.Request) {
	var data msptpd.FilterPPATDto
	if !hh.ValidateStructByURL(w, *r.URL, &data) {
		return
	}

	result, err := s.GetReportExcelTransaksiPPAT(data)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=list_pelaporan_ppat.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	result.Write(w)
}

// DownloadPDFLaporanPPAT func
func DownloadPDFLaporanPPAT(w http.ResponseWriter, r *http.Request) {
	var data msptpd.FilterPPATDto
	if hh.ValidateStructByURL(w, *r.URL, &data) == false {
		return
	}

	result, err := s.GetReportPDFTransaksiPPAT(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=pelaporan-ppat.pdf"))
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
