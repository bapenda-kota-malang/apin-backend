package timehelper

import (
	"time"

	"gorm.io/datatypes"
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

func GetAllFormatTime(datetime *datatypes.Date) (string, string, string, string) {
	if datetime == nil {
		return "00-00-0000", "-", "00:00", "-"
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

	// Memformat menjadi nama bulan dalam bahasa indonesia
	bulan := [...]string{"januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	monthOfYear := bulan[t.Month()]

	// Mengembalikan tiga nilai keluaran
	return formattedDate, dayOfWeek, formattedTime, monthOfYear
}

func ConvertDatatypesTimeToTime(datetime *datatypes.Time) time.Time {
	if datetime != nil {
		gmtDatetime, _ := datetime.Value()
		return gmtDatetime.(time.Time)
	}
	return time.Now()
}

func ConvertDatatypesTimeToString(datetime *datatypes.Time) string {
	if datetime != nil {
		gmtDatetime, _ := datetime.Value()
		return gmtDatetime.(string)
	}
	return time.Now().String()
}

// Day name in Indonesia
func GetDay(datetime *time.Time) string {
	if datetime == nil {
		now := time.Now()
		datetime = &now
	}

	switch int(datetime.Weekday()) {
	case 1:
		return "Senin"
	case 2:
		return "Selasa"
	case 3:
		return "Rabu"
	case 4:
		return "Kamis"
	case 5:
		return "Jum'at"
	case 6:
		return "Sabtu"
	default:
		return "Minggu"
	}
}

// Month name in Indonesia
func GetMonth(datetime *time.Time) string {
	if datetime == nil {
		now := time.Now()
		datetime = &now
	}

	switch int(datetime.Month()) {
	case 1:
		return "Januari"
	case 2:
		return "Februari"
	case 3:
		return "Maret"
	case 4:
		return "April"
	case 5:
		return "Mei"
	case 6:
		return "Juni"
	case 7:
		return "Juli"
	case 8:
		return "Agustus"
	case 9:
		return "September"
	case 10:
		return "Oktober"
	case 11:
		return "November"
	default:
		return "Desember"
	}
}

func ConvertDatatypesDateToTime(datetime *datatypes.Date) time.Time {
	if datetime != nil {
		gmtDatetime, _ := datetime.Value()
		return gmtDatetime.(time.Time)
	}
	return time.Now()
}

// // Month name in Indonesia
// func GetMonth(datetime *time.Time) string {
// 	if datetime == nil {
// 		now := time.Now()
// 		datetime = &now
// 	}

// 	switch int(datetime.Month()) {
// 	case 1:
// 		return "Januari"
// 	case 2:
// 		return "Februari"
// 	case 3:
// 		return "Maret"
// 	case 4:
// 		return "April"
// 	case 5:
// 		return "Mei"
// 	case 6:
// 		return "Juni"
// 	case 7:
// 		return "Juli"
// 	case 8:
// 		return "Agustus"
// 	case 9:
// 		return "September"
// 	case 10:
// 		return "Oktober"
// 	case 11:
// 		return "November"
// 	default:
// 		return "Desember"
// 	}
// }

// func ConvertDatatypesDateToTime(datetime *datatypes.Date) time.Time {
// 	if datetime != nil {
// 		gmtDatetime, _ := datetime.Value()
// 		return gmtDatetime.(time.Time)
// 	}
// 	return time.Now()
// }

// // Month name in Indonesia
// func GetMonth(datetime *time.Time) string {
// 	if datetime == nil {
// 		now := time.Now()
// 		datetime = &now
// 	}

// 	switch int(datetime.Month()) {
// 	case 1:
// 		return "Januari"
// 	case 2:
// 		return "Februari"
// 	case 3:
// 		return "Maret"
// 	case 4:
// 		return "April"
// 	case 5:
// 		return "Mei"
// 	case 6:
// 		return "Juni"
// 	case 7:
// 		return "Juli"
// 	case 8:
// 		return "Agustus"
// 	case 9:
// 		return "September"
// 	case 10:
// 		return "Oktober"
// 	case 11:
// 		return "November"
// 	default:
// 		return "Desember"
// 	}
// }
