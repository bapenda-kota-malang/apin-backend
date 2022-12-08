package objekpajakpbb

import "strings"

func nopParser(nop string) (result []string) {
	result = strings.Split(nop, ".")
	return result
}
