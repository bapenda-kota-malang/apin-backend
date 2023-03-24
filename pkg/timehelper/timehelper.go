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

func GetAllFormatTime(datetime *datatypes.Date) (string, string, string) {
	if datetime == nil {
		return "00-00-0000", "-", "00:00"
	}
	gmtDatetime, _ := datetime.Value()
	t, _ := gmtDatetime.(time.Time)
	// Memformat menjadi dd-MM-YYYY
	dateFormat := "02-01-2006" // 02 -> day, 01 -> month, 2006 -> year
	formattedDate := t.Format(dateFormat)

	// Memformat menjadi nama hari dalam bahasa Indonesia
	hari := [...]string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
	dayOfWeek := hari[t.Weekday()]

	// Memformat menjadi HH:mm
	timeFormat := "15:04" // 15 -> hour, 04 -> minute
	formattedTime := t.Format(timeFormat)

	// Mengembalikan tiga nilai keluaran
	return formattedDate, dayOfWeek, formattedTime
}
