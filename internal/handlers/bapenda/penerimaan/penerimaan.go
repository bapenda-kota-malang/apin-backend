package penerimaaan

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/penerimaan"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

// GetList func
func GetList(w http.ResponseWriter, r *http.Request) {
	var input sspd.FilterReport
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetList(input)
	hh.DataResponse(w, result, err)
}

// DownloadExcelList func
func DownloadExcelList(w http.ResponseWriter, r *http.Request) {
	var input sspd.FilterReport
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetReportExcelPenerimaan(input)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=rekapitulasi-penerimaan.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	result.Write(w)
}

// DownloadPDFList func
func DownloadPDFList(w http.ResponseWriter, r *http.Request) {
	var input sspd.FilterReport
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetReportPDFPenerimaan(input)
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
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=rekapitulasi-penerimaan.pdf"))
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
