// built-in field checker
package structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	h "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator/helper"
)

// register the field checkers
func init() {
	tagValidator = make(map[string]validator)
	RegisterFieldChecker("required", requiredTagValidator)
	RegisterFieldChecker("eq", eqTagValidator)
	RegisterFieldChecker("similar", eqTagValidator)
	RegisterFieldChecker("min", minTagValidator)
	RegisterFieldChecker("max", maxTagValidator)
	RegisterFieldChecker("minLength", minLengthTagValidator)
	RegisterFieldChecker("maxLength", maxLengthTagValidator)
	RegisterFieldChecker("validemail", maxLengthTagValidator)
	RegisterFieldChecker("base64", base64Validator)
}

///// Field checkers

func requiredTagValidator(val reflect.Value, exptVal string) error {
	if (val.Kind() == reflect.String && val.String() == "") || (val.Kind() == reflect.Ptr && val.IsNil()) {
		val.Interface()
		return errors.New("nilai/isi dibutuhkan")
	}
	return nil
}

func eqTagValidator(val reflect.Value, exptVal string) error {
	return nil
}

func minTagValidator(val reflect.Value, exptVal string) error {
	if err := valLimiter(val, exptVal, "<"); err != nil {
		return err
	}
	return nil
}

func maxTagValidator(val reflect.Value, exptVal string) error {
	if err := valLimiter(val, exptVal, ">"); err != nil {
		return err
	}
	return nil
}

func minLengthTagValidator(val reflect.Value, exptVal string) error {
	exptValInt, err := strconv.Atoi(exptVal)
	if err != nil {
		return errors.New("nilai harus berupa angka/numerik")
	}

	valC := h.ValStringer(val) // value converted
	if len(valC) < exptValInt {
		return fmt.Errorf("panjang minimum adalah %v", exptVal)
	}
	return nil
}

func maxLengthTagValidator(val reflect.Value, exptVal string) error {
	exptValInt, err := strconv.Atoi(exptVal)
	if err != nil {
		return errors.New("nilai harus berupa angka/numerik")
	}

	valC := h.ValStringer(val) // value converted
	if len(valC) > exptValInt {
		return fmt.Errorf("panjang maximum adalah %v", exptVal)
	}
	return nil
}

func emailValidator(val reflect.Value, exptVal string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if re.MatchString(h.ValStringer(val)) == false {
		return errors.New("harus memiliki format email yang benar")
	}

	return nil
}

func base64Validator(val reflect.Value, exptVal string) error {
	re := regexp.MustCompile("^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$")

	if re.MatchString(h.ValStringer(val)) == false {
		return errors.New("harus memiliki format base64")
	}

	return nil
}

///// some helper for the default field checker

func valLimiter(val reflect.Value, exptVal string, mode string) error {
	exptValFloat, err := strconv.ParseFloat(exptVal, 64)
	if err != nil {
		return err
	}

	valC := 0.0 // converted value
	valK := val.Kind()
	if valK == reflect.String {
		valCT, err := strconv.ParseFloat(val.String(), 64)
		if err != nil {
			return errors.New("nilai harus berupa angka/numerik")
		}
		valC = valCT
	} else if valK >= reflect.Int && valK <= reflect.Int64 {
		valC = float64(val.Int())
	} else if valK >= reflect.Uint && valK <= reflect.Uint64 {
		valC = float64(val.Uint())
	} else if valK <= reflect.Float32 && valK <= reflect.Float64 {
		valC = val.Float()
	}

	if mode == "<" {
		if exptValFloat > valC {
			return fmt.Errorf("nilai minimum adalah %v", exptVal)
		}
	} else {
		if exptValFloat < valC {
			return fmt.Errorf(fmt.Sprintf("nilai maximum adalah %v", exptVal))
		}
	}
	return nil
}
