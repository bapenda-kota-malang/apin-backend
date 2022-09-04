package stringval

// some functions that are ready to be registered
// please register yourself all the functions you need :3

import (
	"net/mail"
	"reflect"

	rs "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator/refvalstringer"
)

func ValEmailValidator(val reflect.Value, exptVal string) error {
	if _, err := mail.ParseAddress(rs.ValStringer(val)); err != nil {
		return err
	}

	return nil
}
