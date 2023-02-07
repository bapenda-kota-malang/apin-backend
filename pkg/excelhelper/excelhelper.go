package excelhelper

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ExportList(data []interface{}, sn string) (*excelize.File, error) {
	xlsx := excelize.NewFile()
	xlsx.SetSheetName(xlsx.GetSheetName(0), sn)

	r := 1
	for _, v := range data {
		c := 0
		for col, val := range v.(map[string]interface{}) {
			switch t := val.(type) {
			case string, int:
				xlsx.SetCellValue(sn, fmt.Sprintf("%s%d", string(col), r), t)
			default:
				fmt.Println("wrong type " + string(col))
			}
			c = c + 1
		}
		r = r + 1
	}

	return xlsx, nil
}
