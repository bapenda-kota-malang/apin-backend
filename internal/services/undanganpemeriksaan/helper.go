package undanganpemeriksaan

import (
	"strconv"
	"strings"

	"gorm.io/datatypes"
)

func parseCurrentTime(time string) *datatypes.Time {

	// // set timezone
	// layoutTime := "15:04:05"

	// // init the loc
	// loc, _ := time.LoadLocation("Asia/Jakarta")

	// // set timezone,
	// tmpWaktuRincian := time.Now().In(loc).Format(layoutTime)
	// // t, err := time.Parse("15:04:05", tmpWaktuRincian)

	res1 := strings.Split(time, ":")
	hour, _ := strconv.Atoi(res1[0])
	minute, _ := strconv.Atoi(res1[1])

	t := datatypes.NewTime(hour, minute, 0, 0)
	// if err != nil {
	// 	return nil
	// }
	return &t
}
