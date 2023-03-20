package permohonan

import (
	"bytes"
	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/bapenda-kota-malang/apin-backend/pkg/pdf"
	"html/template"
	"os"
	"path/filepath"
)

func GeneratePdf(outFile string, p interface{}) error {
	pdfg, err := pdf.NewPdfg(wkhtmltopdf.PageSizeA4, wkhtmltopdf.OrientationPortrait, false)
	if err != nil {
		return err
	}

	wdPath, err := os.Getwd()
	if err != nil {
		return err
	}

	path := filepath.Join(wdPath, "../../", "internal", "models", "pelayanan", "templates", "permohonan.html")

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, p)
	if err != nil {
		return err
	}

	if err := pdfg.GeneratePdfFromReader(buf, outFile); err != nil {
		return err
	}
	return nil
}
