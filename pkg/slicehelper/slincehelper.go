package slicehelper

func StringInSlice(s string, a []string) bool {
	for _, v := range a {
		if s == v {
			return true
		}
	}
	return false
}
