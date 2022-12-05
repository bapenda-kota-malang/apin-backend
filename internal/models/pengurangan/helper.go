package pengurangan

import "time"

func parseTimeNowToPointer() *time.Time {
	t := time.Now()
	return &t
}
