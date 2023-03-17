package permohonan

import (
	"fmt"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/pelayanan"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type fileResponse struct {
	ID       int    `json:"id"`
	FileName string `json:"fileName"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	var input m.PermohonanRequestDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}
	fmt.Println("reached 1")

	result, err := s.Create(input)
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetList(input)
	hh.DataResponse(w, result, err)
}

func GetStatusNOP(w http.ResponseWriter, r *http.Request) {
	nop := hh.ValidateString(w, r, "id")
	if nop == "" {
		return
	}

	result, err := s.GetStatusNOP(&nop)
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

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.PermohonanRequestDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.Update(id, data)
	hh.DataResponse(w, result, err)
}

func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.PermohonanRequestDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.UpdateStatus(id, data)
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
	w.Header().Set("Content-Disposition", "attachment; filename=list-permohonan.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	result.Write(w)
}
func DownloadPdf(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}
	result, err := s.DownloadPdf(id)
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
	//hh.DataResponse(w, result, err)

	//// Generate HTML content
	//htmlContent := "<html><body><h1>Hello, World!</h1></body></html>"
	//
	//// Write HTML content to a temporary file
	//htmlFile, err := ioutil.TempFile("/tmp", "*.html")
	//if err != nil {
	//	log.Fatalf("Failed to create temporary file: %v", err)
	//}
	//defer os.Remove(htmlFile.Name())
	//if _, err := htmlFile.Write([]byte(htmlContent)); err != nil {
	//	log.Fatalf("Failed to write to temporary file: %v", err)
	//}
	//
	//// Generate PDF from HTML using wkhtmltopdf
	//pdfFile, err := ioutil.TempFile("/tmp", "*.pdf")
	//if err != nil {
	//	log.Fatalf("Failed to create temporary file: %v", err)
	//}
	////defer os.Remove(pdfFile.Name())
	//cmd := exec.Command("wkhtmltopdf", htmlFile.Name(), pdfFile.Name())
	//cmd.Stdin = strings.NewReader(htmlContent)
	//out, err := cmd.Output()
	//
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.Header().Set("Content-Type", "application/pdf")
	//w.Header().Set("Content-Disposition", "attachment; filename=output.pdf")
	//w.Write(out)

	//// Generate HTML content
	//htmlContent := "<html><body><h1>Hello, World!</h1></body></html>"
	//
	//// Write HTML content to a temporary file
	//htmlFile, err := ioutil.TempFile("/tmp", "*.html")
	//if err != nil {
	//	log.Fatalf("Failed to create temporary file: %v", err)
	//}
	//defer os.Remove(htmlFile.Name())
	//if _, err := htmlFile.Write([]byte(htmlContent)); err != nil {
	//	log.Fatalf("Failed to write to temporary file: %v", err)
	//}
	//
	//// Generate PDF from HTML using wkhtmltopdf
	//pdfFile, err := ioutil.TempFile("/tmp", "*.pdf")
	//if err != nil {
	//	log.Fatalf("Failed to create temporary file: %v", err)
	//}
	//defer os.Remove(pdfFile.Name())
	//cmd := exec.Command("wkhtmltopdf", htmlFile.Name(), pdfFile.Name())
	//if err := cmd.Run(); err != nil {
	//	log.Fatalf("Failed to generate PDF: %v", err)
	//}
	//
	//// Parse the PDF file using pdfcpu
	//err = pdfcpu.Read(pdfFile.Name(), nil)
	//if err != nil {
	//	log.Fatalf("Failed to read PDF file: %v", err)
	//}
	//
	//
	//f, err := os.Open("file.pdf")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//defer f.Close()

	//// Parse the PDF file
	//config := api.NewDefaultConfiguration()
	//pdf, err := api.Read(f, config)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//// Set the response headers for a PDF file
	//w.Header().Set("Content-Type", "application/pdf")
	//w.Header().Set("Content-Disposition", "attachment; filename=file.pdf")
	//
	//// Send the PDF file as response
	//err = api.Write(w, pdf)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	//log.Println("PDF file successfully generated and watermarked!")

}
