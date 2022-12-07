// service
package penagihan

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
)

// /// Private funcs start here
const source = "himbauan"

func GetReportHimbauan(re m.RequestDto) (*excelize.File, error) {
	datas := new([]m.Sppt)
	return generateHimbauanToXLSX(*datas)
}

func generateHimbauanToXLSX(datas []m.Sppt) (*excelize.File, error) {
	var (
		start     int = 8
		sppt          = []string{"1000", "2000", "3000", "4000", "5000"}
		ketetapan     = []string{"100.000.000", "200.000.000", "300.000.000", "400.000.000", "500.000.000"}
	)

	xlsx := excelize.NewFile()

	tunggakanReport := "Laporan Tunggakan"
	xlsx.SetSheetName(xlsx.GetSheetName(1), tunggakanReport)

	xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("A%d", start), "No")
	xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("B%d", start), "Nama Kelurahan")
	xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("C%d", start), "SPPT")
	xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("D%d", start), "Ketetapan")
	xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("E%d", start), "SPPT")
	xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("F%d", start), "Ketetapan")
	xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("G%d", start), "SPPT")
	xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("H%d", start), "Ketetapan")
	xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("I%d", start), "SPPT")
	xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("J%d", start), "Ketetapan")
	xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("K%d", start), "SPPT")
	xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("L%d", start), "Ketetapan")

	err := xlsx.AutoFilter(tunggakanReport, "A1", "AJ1", "")
	if err != nil {
		return nil, err
	}

	for i, data := range datas {
		xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("A%d", i+start+2), data.CreatedAt)
		xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("B%d", i+start+2), data.KotaWP_sppt)
		xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("C%d", i+start+2), sppt[0])
		xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("D%d", i+start+2), ketetapan[0])
		xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("E%d", i+start+2), sppt[1])
		xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("F%d", i+start+2), ketetapan[1])
		xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("G%d", i+start+2), sppt[2])
		xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("H%d", i+start+2), ketetapan[2])
		xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("I%d", i+start+2), sppt[3])
		xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("J%d", i+start+2), ketetapan[3])
		xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("K%d", i+start+2), sppt[4])
		xlsx.SetCellValue(tunggakanReport, fmt.Sprintf("L%d", i+start+2), ketetapan[4])
	}

	return xlsx, nil
}
