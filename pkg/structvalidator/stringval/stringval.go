package stringval

// some functions that are ready to be registered
// please register yourself all the functions you need :3

import (
	"errors"
	"reflect"
	"regexp"

	rs "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator/refvalstringer"
)

func ValEmailValidator(val reflect.Value, exptVal string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if re.MatchString(rs.ValStringer(val)) == false {
		return errors.New("requires valid email address")
	}

	return nil
}
