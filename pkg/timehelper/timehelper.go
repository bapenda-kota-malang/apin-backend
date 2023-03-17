package timehelper

import (
	"gorm.io/datatypes"
	"time"
)

func ParseTime(input *string) *time.Time {
	if input == nil {
		return nil
	}
	t, err := time.Parse("2006-01-02", *input)
	if err != nil {
		return nil
	}
	return &t
}

func TimeNow() *time.Time {
	t := time.Now()
	return &t
}

func GetDateFromUTCDatetime(datetime *datatypes.Date) string {
	if datetime != nil {
		gmtDatetime, _ := datetime.Value()
		return gmtDatetime.(time.Time).Format("2006-01-02")
	}
	return "00-00-0000"

}
