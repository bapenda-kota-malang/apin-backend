// Go Struct Validator
package structvalidator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/maps"
	"gorm.io/datatypes"
)

// error format, exported for casting purpose
type ValidationError struct {
	// Error      error       `json:"error"`
	ErrMessage string      `json:"errMessage"`
	Code       string      `json:"code"`
	ExptdValue string      `json:"exptdValue"`
	GivenValue interface{} `json:"givenValue"`
}

// just key val for the tag
type keyVal struct {
	Key string
	Val string
}

// viladator func interface?
// param: reflect value, string
type validator func(reflect.Value, string) error

// list of validator for the given key from tag
var tagValidator map[string]validator

// tag name to validate
const tagName = "validate"

// Validation of each field based on the registered field checkers
func Validate(input interface{}, nameSpaces ...string) map[string]ValidationError {
	// identiy value and loop if its pointer until reaches non pointer
	inputV := reflect.ValueOf(input)

	// loop until we get what kind lays behind the input
	for inputV.Kind() == reflect.Pointer || inputV.Kind() == reflect.Interface {
		inputV = inputV.Elem()
	}

	// non struct cant be validated
	if inputV.Kind() != reflect.Struct {
		return nil
	}

	// namespace will be available if it is sub validation
	nameSpace := ""
	if len(nameSpaces) > 0 {
		if len(nameSpaces) > 1 && nameSpaces[1] != "" {
			nameSpace += "(embedded:" + nameSpaces[0] + ")."
		} else {
			nameSpace += nameSpaces[0] + "."
		}
	}

	// check each field
	// inputT := reflect.TypeOf(inputV.Interface()) // keep this for now
	inputT := inputV.Type()
	errList := map[string]ValidationError{}
	for i := 0; i < inputV.NumField(); i++ {
		// identify type and value of the field
		fieldT := inputT.Field(i)
		fieldV := inputV.Field(i)
		for fieldV.Kind() == reflect.Ptr {
			if fieldV.IsZero() {
				break
			}
			fieldV = fieldV.Elem()
		}

		// if current field is struct, validate again
		// TODO: find information about this -> || fieldT.Type.Name() == ""
		typeString := fieldT.Type.String()
		if (fieldT.Type.Kind() == reflect.Struct) && typeString != "time.Time" {
			embeddedMode := ""
			if fieldT.Anonymous {
				embeddedMode = "(embedded)"
			}
			tag := fieldT.Tag.Get("json")
			if tag == "" {
				tag = fieldT.Name
			}
			maps.Copy(errList, Validate(fieldV.Interface(), tag, embeddedMode))
			continue
		}

		tag := fieldT.Tag.Get(tagName)
		if tag != "" {
			parsedTag := parseTag(tag)
			key := fieldT.Tag.Get("json")
			if key != "" {
				keys := strings.Split(key, ",")
				if keys[0] != "" {
					key = keys[0]
				} else {
					key = fieldT.Name
				}
			} else {
				key = fieldT.Name
			}
			// based on slice or not
			if fieldV.Kind() == reflect.Slice {
				// special case untuk required
				required := false
				for _, v := range parsedTag {
					if v.Key == "required" {
						required = true
						break
					}
				}
				// empty array
				if fieldV.Len() == 0 {
					if required {
						errList[nameSpace+key] = ValidationError{"nilai/isi dibutuhkan", "required", "", fieldV.Interface()}
					}
					continue
				}
				// loop
				if fieldV.Index(0).Kind() == reflect.Struct {
					for ix := 0; ix < fieldV.Len(); ix++ {
						maps.Copy(errList, Validate(fieldV.Index(ix).Interface(), fmt.Sprintf("%v[%v]", key, ix)))
					}
				} else {
					for ix := 0; ix < fieldV.Len(); ix++ {
						CheckParsedTag(parsedTag, fieldV.Index(ix), errList, fmt.Sprintf("%v[%v]", key, ix))
					}
				}
			} else {
				// non slice
				CheckParsedTag(parsedTag, fieldV, errList, nameSpace+key)
				// 	maps.Copy(errList, Validate(fieldV.Index(ix).Interface(), fmt.Sprintf("%v[%v]", key, ix)))
				// for _, kv := range parsedTag {
				// 	if _, ok := tagValidator[kv.Key]; ok {
				// 		// fmt.Println(fieldT.Name, "is being validated")
				// 		err := tagValidator[kv.Key](fieldV, kv.Val)
				// 		if err != nil {
				// 			// fmt.Println(fieldT.Name, "validation is failed")
				// 			errList[nameSpace+key] = ValidationError{err.Error(), kv.Key, kv.Val, fieldV.Interface()}
				// 			break // 1 err is enough, break from error check of the current field
				// 		}
				// 	}
				// }
			}
		}
	}

	if len(errList) > 0 {
		return errList
	}
	return nil
}

func CheckParsedTag(parsedTag []keyVal, fv reflect.Value, el map[string]ValidationError, key string) {
	for _, kv := range parsedTag {
		if _, ok := tagValidator[kv.Key]; ok {
			// fmt.Println(fieldT.Name, "is being validated")
			err := tagValidator[kv.Key](fv, kv.Val)
			if err != nil {
				// fmt.Println(fieldT.Name, "validation is failed")
				el[key] = ValidationError{err.Error(), kv.Key, kv.Val, fv.Interface()}
				break // 1 err is enough, break from error check of the current field
			}
		}
	}
}

// Validation from IO Reader
func ValidateIoReader(container interface{}, input io.Reader) map[string]ValidationError {
	decoder := json.NewDecoder(input)
	err := decoder.Decode(&container)
	if err != nil {
		cV := reflect.ValueOf(container)
		for cV.Kind() == reflect.Pointer || cV.Kind() == reflect.Interface {
			cV = cV.Elem()
		}
		structName := cV.Type().Name()
		return map[string]ValidationError{
			"struct": {fmt.Sprintf("gagal parsing payload ke %v, pesan error: %v", structName, err), "struct", fmt.Sprintf("value of %v", structName), ""},
		}
	}

	// same process with normal validation
	return Validate(container)
}

// Validation from url
// caveat: url's structure makes it impossible to do deep parsing
func ValidateURL(container any, url url.URL) map[string]ValidationError {
	cV := reflect.ValueOf(container).Elem()
	for cV.Kind() == reflect.Pointer || cV.Kind() == reflect.Interface {
		cV = cV.Elem()
	}

	cT := cV.Type()
	values := url.Query()
	errList := map[string]ValidationError{}
	for i := 0; i < cV.NumField(); i++ {
		fieldV := cV.Field(i)
		fieldT := cT.Field(i)

		if !fieldV.CanSet() {
			continue
		}

		key := fieldT.Tag.Get("json")
		if key == "" {
			key = fieldT.Name
		}

		vals, ok := values[key]
		if !ok {
			continue
		}

		switch fieldV.Interface().(type) {
		case bool, *bool:
			var v bool
			if strings.ToLower(vals[0]) == "true" || vals[0] == "1" {
				v = true
			} else {
				v = false
			}
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&v))
			} else {
				fieldV.Set(reflect.ValueOf(v))
			}
		case string, *string:
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&vals[0]))
			} else {
				fieldV.Set(reflect.ValueOf(vals[0]))
			}
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64:
			if valInt, err := strconv.Atoi(vals[0]); err != nil {
				errList[key] = ValidationError{err.Error(), key, vals[0], fieldV.Interface()}
			} else {
				v := autoCastInt(valInt, fieldV)
				fieldV.Set(v)
			}
		case float64, *float64:
			strFloat, err := strconv.ParseFloat(vals[0], 64)
			if err != nil {
				errList[key] = ValidationError{err.Error(), key, vals[0], fieldV.Interface()}
			}
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&strFloat))
			} else {
				fieldV.Set(reflect.ValueOf(strFloat))
			}
		case float32, *float32:
			strFloat, err := strconv.ParseFloat(vals[0], 32)
			if err != nil {
				errList[key] = ValidationError{err.Error(), key, vals[0], fieldV.Interface()}
			}
			strFloat32 := float32(strFloat)
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&strFloat32))
			} else {
				fieldV.Set(reflect.ValueOf(strFloat32))
			}
		case []string, *[]string:
			fieldV.Set(reflect.ValueOf(&vals))
		case datatypes.Date, *datatypes.Date:
			time, err := time.Parse("2006-01-02T15:04:05.000Z", vals[0])
			if err != nil {
				errList[key] = ValidationError{err.Error(), key, vals[0], fieldV.Interface()}
			}
			date := datatypes.Date(time)
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&date))
			} else {
				fieldV.Set(reflect.ValueOf(date))
			}
		case time.Time, *time.Time:
			time, err := time.Parse("2006-01-02T15:04:05.000Z", vals[0])
			if err != nil {
				errList[key] = ValidationError{err.Error(), key, vals[0], fieldV.Interface()}
			}
			if fieldV.Kind() == reflect.Ptr {
				fieldV.Set(reflect.ValueOf(&time))
			} else {
				fieldV.Set(reflect.ValueOf(time))
			}
		// TODO: make any *[]int as a function
		case *[]int8:
			failed := false
			valX := []int8{}
			for _, val := range vals {
				if valInt, err := strconv.Atoi(val); err != nil {
					failed = true
					errList[key] = ValidationError{err.Error(), key, val, fieldV.Interface()}
				} else {
					valX = append(valX, int8(valInt))
				}
			}
			if !failed {
				fieldV.Set(reflect.ValueOf(valX))
			}
			// case []int16:
			// 	failed := false
			// 	valX := []int16{}
			// 	for _, val := range vals {
			// 		if valInt, err := strconv.Atoi(val); err != nil {
			// 			failed = true
			// 			errList[key] = ValidationError{err.Error(), key, val, fieldV.Interface()}
			// 		} else {
			// 			valX = append(valX, int16(valInt))
			// 		}
			// 	}
			// 	if !failed {
			// 		fieldV.Set(reflect.ValueOf(valX))
			// 	}
			// case []int32:
			// 	failed := false
			// 	valX := []int32{}
			// 	for _, val := range vals {
			// 		if valInt, err := strconv.Atoi(val); err != nil {
			// 			failed = true
			// 			errList[key] = ValidationError{err.Error(), key, val, fieldV.Interface()}
			// 		} else {
			// 			valX = append(valX, int32(valInt))
			// 		}
			// 	}
			// 	if !failed {
			// 		fieldV.Set(reflect.ValueOf(valX))
			// 	}
			// case []int64:
			// 	failed := false
			// 	valX := []int64{}
			// 	for _, val := range vals {
			// 		if valInt, err := strconv.Atoi(val); err != nil {
			// 			failed = true
			// 			errList[key] = ValidationError{err.Error(), key, val, fieldV.Interface()}
			// 		} else {
			// 			valX = append(valX, int64(valInt))
			// 		}
			// 	}
			// 	if !failed {
			// 		fieldV.Set(reflect.ValueOf(valX))
			// 	}
		}
	}

	if len(errList) > 0 {
		return errList
	}

	return Validate(container)
}

// register a validator
func RegisterFieldChecker(tag string, validatorF validator) {
	tagValidator[tag] = validatorF
}

// unregister a validator
func UnregisterFieldChecker(tag string) {
	delete(tagValidator, tag)
}

///////////// Helper Goes here
