package tbp

import (
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/tbp"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
)

func generateNomor() int {
	var tmp int
	var tmpTbp mt.Tbp
	nomor := a.DB.Last(&tmpTbp)
	if nomor.Error != nil {
		return 1
	} else {
		tmp = *tmpTbp.TbpNumber
		tmp++
	}
	return tmp
}

// func parseCurrentTime(input string) *time.Time {
// 	t, err := time.Parse("15:04:05", input)
// 	if err != nil {
// 		return nil
// 	}
// 	return &t
// }
