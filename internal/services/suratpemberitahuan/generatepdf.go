package suratpemberitahuan

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/suratpemberitahuan"
	"github.com/bapenda-kota-malang/apin-backend/pkg/pdf"
	"github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

type Surat struct {
	Title             string
	Nomor             string
	Sifat             string
	Lampiran          string
	Perihal           string
	Pengelola         string
	Npwpd             string
	Alamat            string
	JenisPajak        string
	DataTable         []m.DataTableSurat
	TotalKetetapan    string
	TotalDenda        string
	TotalTelahDibayar string
	TotalSisaPajak    string
}

func (s *Surat) GeneratePdf(outFile string) error {
	pdfg, err := pdf.NewPdfg(wkhtmltopdf.PageSizeA4, wkhtmltopdf.OrientationPortrait, false)
	if err != nil {
		return err
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	pathSurat := filepath.Join(exPath, "../../", "internal", "models", "suratpemberitahuan", "template", "surat.html")

	tmpl, err := template.ParseFiles(pathSurat)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, s)
	if err != nil {
		return err
	}

	if err := pdfg.GeneratePdfFromReader(buf, outFile); err != nil {
		return err
	}
	return nil
}

func (s *Surat) AppendContent(data m.DataTableSurat) {
	data.No = uint16(len(s.DataTable) + 1)
	s.DataTable = append(s.DataTable, data)
}

func (s *Surat) SetTotal(total map[string]float64) {
	s.TotalKetetapan = servicehelper.FormatCurrency(total["ketetapan"])
	s.TotalDenda = servicehelper.FormatCurrency(total["denda"])
	s.TotalTelahDibayar = servicehelper.FormatCurrency(total["telahDibayar"])
	s.TotalSisaPajak = servicehelper.FormatCurrency(total["sisaPajak"])
}
