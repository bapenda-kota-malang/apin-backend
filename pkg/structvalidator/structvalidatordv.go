package structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	h "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator/helper"
)

// register default tag validator
func init() {
	tagValidator = make(map[string]validator)
	RegisterValidator("required", requiredTagValidator)
	RegisterValidator("eq", eqTagValidator)
	RegisterValidator("similar", eqTagValidator)
	RegisterValidator("min", minTagValidator)
	RegisterValidator("max", maxTagValidator)
	RegisterValidator("minLength", minLengthTagValidator)
	RegisterValidator("maxLength", maxLengthTagValidator)
}

// // for the default validator we have: required + comparisons
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

// /// some helper for default validator func
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
	} else if valK >= reflect.Int && valK <= reflect.Uint64 {
		valC = float64(val.Int())
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
