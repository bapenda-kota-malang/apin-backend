package sspd

import (
	"strconv"
	"strings"
	"time"

	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	"gorm.io/datatypes"
)

func generateNomor() int {
	var tmp int
	var tmpSspd ms.Sspd
	nomor := a.DB.Last(&tmpSspd)
	if nomor.Error != nil {
		return 1
	} else {
		tmp = *tmpSspd.SspdNumber
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
