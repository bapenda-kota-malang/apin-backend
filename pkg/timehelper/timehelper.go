package timehelper

import "time"

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
