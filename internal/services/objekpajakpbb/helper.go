package objekpajakpbb

import "strings"

func nopParser(nop string) (result []string, area_kode string) {
	result = strings.Split(nop, ".")
	area_kode = result[0] + result[1] + result[2] + result[3]
	return result, area_kode
}
