/*
Using library github.com/SebastiaanKlippert/go-wkhtmltopdf to generate pdf from html
*/
package pdf

import (
	"io"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type Pdfg struct {
	*wkhtmltopdf.PDFGenerator
}

// create pdf generate instance, apply basic options, if error return error
//
// PageSize and PageOrientation refer to: https://pkg.go.dev/github.com/SebastiaanKlippert/go-wkhtmltopdf#pkg-constants
//
// if pageSize empty using "A4" size as default size
//
// if pageOrientation empty using "Potrait" as default orientation
//
// grayscale true if you want generated pdf to be grayscale mode
//
//	pdfg, err := fNewPdfg("A4", "Potrait", false)
func NewPdfg(pageSize, pageOrientation string, grayscale bool) (*Pdfg, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	if pageSize == "" {
		pageSize = wkhtmltopdf.PageSizeA4
	}
	if pageOrientation == "" {
		pageOrientation = wkhtmltopdf.OrientationPortrait
	}
	// Set global options
	pdfg.PageSize.Set(pageSize)
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(pageOrientation)
	pdfg.Grayscale.Set(grayscale)
	return &Pdfg{pdfg}, nil
}

// Create pdf and write buffer to outFile file
func (p *Pdfg) CreateAndSave(outFile string) error {
	if err := p.Create(); err != nil {
		return err
	}
	if err := p.WriteFile(outFile); err != nil {
		return err
	}
	return nil
}

// create pdf from inFile file html to outFile
func (p *Pdfg) GeneratePdfFromFile(inFile, outFile string) error {
	page := wkhtmltopdf.NewPage(inFile)
	// Set options for this page
	page.EnableLocalFileAccess.Set(true)
	p.AddPage(page)
	if err := p.CreateAndSave(outFile); err != nil {
		return err
	}
	return nil
}

// create pdf from input reader and save it to outFile
func (p *Pdfg) GeneratePdfFromReader(input io.Reader, outFile string) error {
	page := wkhtmltopdf.NewPageReader(input)
	// Set options for this page
	page.EnableLocalFileAccess.Set(true)
	p.AddPage(page)
	if err := p.CreateAndSave(outFile); err != nil {
		return err
	}
	return nil
}
