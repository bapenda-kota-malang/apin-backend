package revavalstringer

import (
	"fmt"
	"reflect"
	"strconv"
)

func ValStringer(val reflect.Value) string {
	valK := val.Kind()
	var valC string
	if valK == reflect.String {
		valC = val.String()
	} else if valK >= reflect.Int && valK < reflect.Uint64 {
		tmp := 0
		if valK >= reflect.Uint {
			tmp = int(val.Uint())
		} else {
			tmp = int(val.Int())
		}
		valC = strconv.Itoa(tmp)
	} else if valK >= reflect.Float32 && valK < reflect.Float64 {
		valC = fmt.Sprintf("%v", val.Float())
	}
	return valC
}
