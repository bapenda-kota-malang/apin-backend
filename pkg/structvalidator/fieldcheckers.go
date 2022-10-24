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
	RegisterFieldChecker("validemail", emailValidator)
	RegisterFieldChecker("base64", base64Validator)
	RegisterFieldChecker("nik", nikValidator)
	RegisterFieldChecker("nohp", noHpValidator)
}

///// Field checkers

func requiredTagValidator(val reflect.Value, exptVal string) error {
	if (val.Kind() == reflect.String && val.String() == "") || (val.Kind() == reflect.Ptr && val.IsNil()) {
		return errors.New("nilai/isi dibutuhkan")
	}
	return nil
}

func eqTagValidator(val reflect.Value, exptVal string) error {
	return nil
}

func minTagValidator(val reflect.Value, exptVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	if err := valLimiter(val, exptVal, "<"); err != nil {
		return err
	}
	return nil
}

func maxTagValidator(val reflect.Value, exptVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	if err := valLimiter(val, exptVal, ">"); err != nil {
		return err
	}
	return nil
}

func minLengthTagValidator(val reflect.Value, exptVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
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
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
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
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !re.MatchString(h.ValStringer(val)) {
		return errors.New("harus memiliki format email yang benar")
	}

	return nil
}

func base64Validator(val reflect.Value, exptVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	re := regexp.MustCompile(`^(data:)([\w\/\+-]*)(;charset=[\w-]+|;base64){0,1},([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)

	if !re.MatchString(h.ValStringer(val)) {
		return errors.New("harus memiliki format base64")
	}

	return nil
}

/*
Artinya
validasi 6 karakter pertama, 2 angka pertama sesuai dengan format provinsi di Indonesia, 2 karakter berikutnya angka, dan 2 karakter selanjutnya juga angka.
6 karakter berikutnya mengikuti format tanggal lahir 01..31 (untuk laki-laki) atau 41..71 (untuk perempuan), bulan lahir 01-12, dan tahun lahir 2 angka
4 karakter terakhir adalah angka sequence.
*/
func nikValidator(val reflect.Value, exptVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	re := regexp.MustCompile(`^(1[1-9]|21|[37][1-6]|5[1-3]|6[1-5]|[89][12])\d{2}\d{2}([04][1-9]|[1256][0-9]|[37][01])(0[1-9]|1[0-2])\d{2}\d{4}$`)

	if !re.MatchString(h.ValStringer(val)) {
		return errors.New("harus memiliki format base64")
	}

	return nil
}

// validate no hp indonesia
//
// dianggap valid untuk: +62897123456, 0897123456.
func noHpValidator(val reflect.Value, exptVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	re := regexp.MustCompile(`^(\+62|62|0)8[1-9][0-9]{6,9}$`)

	if !re.MatchString(h.ValStringer(val)) {
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
