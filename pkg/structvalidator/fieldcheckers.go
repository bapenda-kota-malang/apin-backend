// built-in field checker
package structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/bapenda-kota-malang/apin-backend/pkg/base64helper"
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
	RegisterFieldChecker("b64size", b64SizeKb)
	RegisterFieldChecker("nik", nikValidator)
	RegisterFieldChecker("nohp", noHpValidator)
	RegisterFieldChecker("notelp", noTelpValidator)
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

// check base64 string validation, ex: base64 or base64=pdf,image,excel or base64=pdf
func base64Validator(val reflect.Value, exptVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	re := regexp.MustCompile(`^(data:)([\w\/\+-]*)(;charset=[\w-]+|;base64){0,1},([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)

	b64RawString := h.ValStringer(val)
	if !re.MatchString(b64RawString) {
		return errors.New("harus memiliki format base64")
	}

	if exptVal != "" {
		ext, err := base64helper.GetExtensionBase64(b64RawString)
		if err != nil {
			return err
		}
		switch ext {
		case "xlsx", "xls":
			ext = "excel"
		case "png", "jpeg":
			ext = "image"
		case "pdf":
			// do nothing change
		}
		supportFilesSlice := strings.Split(exptVal, ",")
		supportFilesMap := make(map[string]struct{})
		for v := range supportFilesSlice {
			supportFilesMap[supportFilesSlice[v]] = struct{}{}
		}
		if _, exist := supportFilesMap[ext]; !exist {
			return errors.New("file not supported")
		}
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
		return errors.New("harus memiliki format nik")
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
	re := regexp.MustCompile(`^(\+62|62|0)8[1-9][0-9]{6,10}$`)

	if !re.MatchString(h.ValStringer(val)) {
		return errors.New("harus memiliki format nomor hp")
	}

	return nil
}

func noTelpValidator(val reflect.Value, exptVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}
	re := regexp.MustCompile(`^(\+62|62|0)[2-9][1-9][0-9]{6,10}$`)

	if !re.MatchString(h.ValStringer(val)) {
		return errors.New("harus memiliki format nomor telp / hp")
	}

	return nil
}

// validate base64 approx size file
//
// using kb for parameter value eg: 1024 means 1024KB or 1MB or 1024000 B max allowed size file
func b64SizeKb(val reflect.Value, exptVal string) error {
	if val.Kind() == reflect.Pointer && val.IsNil() {
		return nil
	}

	exptValInt, err := strconv.Atoi(exptVal)
	if err != nil {
		return errors.New("nilai harus berupa angka/numerik")
	}

	datas := h.ValStringer(val)

	l := len(datas)

	// count how many trailing '=' there are (if any)
	eq := 0
	if l >= 2 {
		if datas[l-1] == '=' {
			eq++
		}
		if datas[l-2] == '=' {
			eq++
		}

		l -= eq
	}

	// basically:
	// eq == 0 :	bits-wasted = 0
	// eq == 1 :	bits-wasted = 2
	// eq == 2 :	bits-wasted = 4

	// so orig length ==  (l*6 - eq*2) / 8

	// if bytes size > max bytes allowed then
	if (l*3-eq)/4 > exptValInt*1000 {
		return fmt.Errorf("ukuran file lebih besar dari %d KB", exptValInt)
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
