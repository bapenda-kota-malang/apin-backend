package structvalidator

import (
	"reflect"
	"strings"
)

// parse tag for key - val
func parseTag(tag string) []keyVal {
	kvList := []keyVal{}
	for _, item := range strings.Split(tag, ";") {
		pair := strings.SplitN(strings.TrimSpace(item), "=", 2)
		if len(pair) == 0 {
			continue
		}
		if len(pair) == 1 {
			kvList = append(kvList, keyVal{pair[0], ""})
		}
		if len(pair) == 2 {
			kvList = append(kvList, keyVal{pair[0], pair[1]})
		}
	}
	return kvList
}

// autocast intype, beware as it returns 0 for the default value
func autoCastInt(input int, kind reflect.Value) reflect.Value {
	switch kind.Interface().(type) {
	case int:
		return reflect.ValueOf(input)
	case int8:
		return reflect.ValueOf(int8(input))
	case int16:
		return reflect.ValueOf(int16(input))
	case int32:
		return reflect.ValueOf(int32(input))
	case int64:
		return reflect.ValueOf(int64(input))
	case uint:
		return reflect.ValueOf(uint(input))
	case uint8:
		return reflect.ValueOf(uint8(input))
	case uint16:
		return reflect.ValueOf(uint16(input))
	case uint32:
		return reflect.ValueOf(uint32(input))
	case uint64:
		return reflect.ValueOf(uint64(input))
	case *int:
		x := input
		return reflect.ValueOf(&x)
	case *int8:
		x := int8(input)
		return reflect.ValueOf(&x)
	case *int16:
		x := int16(input)
		return reflect.ValueOf(&x)
	case *int32:
		x := int32(input)
		return reflect.ValueOf(&x)
	case *int64:
		x := int64(input)
		return reflect.ValueOf(&x)
	case *uint:
		x := uint(input)
		return reflect.ValueOf(&x)
	case *uint8:
		x := uint8(input)
		return reflect.ValueOf(&x)
	case *uint16:
		x := uint16(input)
		return reflect.ValueOf(&x)
	case *uint32:
		x := uint32(input)
		return reflect.ValueOf(&x)
	case *uint64:
		x := uint64(input)
		return reflect.ValueOf(&x)
	}
	return reflect.ValueOf(0)
}

// func autoCastIntArr(input any, theType reflect.Type) (reflect.Value, bool) {
// 	failed := false
// 	valX := reflect.ValueOf(input).Interface()
// 	for _, val := range vals {
// 		if valInt, err := strconv.Atoi(val); err != nil {
// 			failed = true
// 			errList[key] = ValidationError{err.Error(), key, val, fieldV.Interface()}
// 		} else {
// 			valX = append(valX, int8(valInt))
// 		}
// 	}
// 	if !failed {
// 		fieldV.Set()
// 	}
// 	return reflect.ValueOf(valX), false
// }
