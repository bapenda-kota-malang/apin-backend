package sspd

import (
	"strconv"
	"strings"
	"time"

	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	"gorm.io/datatypes"
)

func generateNomorUrut() int {
	var tmp int
	var tmpSspd ms.Sspd
	nomor := a.DB.Last(&tmpSspd)
	if nomor.Error != nil {
		return 1
	} else {
		tmp = tmpSspd.NomorUrut
		tmp++
	}
	return tmp
}

func generateNomorTahun(t time.Time) int {
	tmp := t.Format("2006-01-02")
	tmpString := tmp[2:4]
	result, _ := strconv.Atoi(tmpString)
	return result
}

func generateNomorOutput(tahun, urut int) string {
	tmpTahun := strconv.Itoa(tahun)
	tmpUrut := strconv.Itoa(urut)
	if len(tmpUrut) == 1 {
		tmpUrut = "00000" + tmpUrut
	} else if len(tmpUrut) == 2 {
		tmpUrut = "0000" + tmpUrut
	} else if len(tmpUrut) == 3 {
		tmpUrut = "000" + tmpUrut
	} else if len(tmpUrut) == 4 {
		tmpUrut = "00" + tmpUrut
	} else if len(tmpUrut) == 5 {
		tmpUrut = "0" + tmpUrut
	}
	output := tmpTahun + tmpUrut
	return output
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
