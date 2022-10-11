package timehelper

import "time"

func ParseTime(input string) *time.Time {
	t, err := time.Parse("2006-01-02", input)
	if err != nil {
		return nil
	}
	return &t
}
