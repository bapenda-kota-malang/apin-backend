package tbp

import (
	"strconv"
	"strings"
	"time"

	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/tbp"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	"gorm.io/datatypes"
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

func parseCurrentTime() *datatypes.Time {

	// set timezone
	layoutTime := "15:04:05"

	// init the loc
	loc, _ := time.LoadLocation("Asia/Jakarta")

	// set timezone,
	tmpWaktuRincian := time.Now().In(loc).Format(layoutTime)
	// t, err := time.Parse("15:04:05", tmpWaktuRincian)

	res1 := strings.Split(tmpWaktuRincian, ":")
	hour, _ := strconv.Atoi(res1[0])
	minute, _ := strconv.Atoi(res1[1])
	second, _ := strconv.Atoi(res1[2])

	t := datatypes.NewTime(hour, minute, second, 0)
	// if err != nil {
	// 	return nil
	// }
	return &t
}
