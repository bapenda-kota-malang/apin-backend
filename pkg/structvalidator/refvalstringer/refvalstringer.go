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
		valC = strconv.Itoa(int(val.Int()))
	} else if valK >= reflect.Float32 && valK < reflect.Float64 {
		valC = fmt.Sprintf("%v", val.Float())
	}
	return valC
}
