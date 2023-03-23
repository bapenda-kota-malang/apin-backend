package potensiopwp

import (
	"bytes"
	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/bapenda-kota-malang/apin-backend/pkg/pdf"
	"html/template"
	"os"
	"path/filepath"
)

func GeneratePdf(outFile string, p interface{}, kodeObjek string) error {
	pdfg, err := pdf.NewPdfg(wkhtmltopdf.PageSizeA4, wkhtmltopdf.OrientationPortrait, false)
	if err != nil {
		return err
	}

	wdPath, err := os.Getwd()
	if err != nil {
		return err
	}
	var templateFile = "objekpajak1.html"
	switch kodeObjek {
	case "01": //kost hotel
		templateFile = "objekpajak4.html"
	case "02": // RM
		templateFile = "objekpajak1.html"
	case "03": // karaoke
		templateFile = "objekpajak3.2.html"
	case "04": // reklame
		templateFile = "objekpajak1.html"
	case "05": //no pln
		templateFile = "objekpajak2.html"
	case "06": //none
		templateFile = "objekpajak1.html"
	case "07": //parkir
		templateFile = "objekpajak7.html"
	case "08": // airtanah
		templateFile = "objekpajak6.html"
	}
	path := filepath.Join(wdPath, "../../", "internal", "models", "potensiopwp", "templates", templateFile)

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
