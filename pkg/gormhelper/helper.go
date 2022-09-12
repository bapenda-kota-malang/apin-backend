package gormhelper

// just local functions, avoid import if it is simple

// check if stnig availble in an array
func stringInSlice(s string, a []string) bool {
	for _, v := range a {
		if s == v {
			return true
		}
	}
	return false
}
