package userservice

import (
	"errors"
	"time"
)

func checkExpiration(input time.Time) error {
	now := time.Now()
	if now.Before(input) {
		return nil
	} else {
		return errors.New("token request telah expired")
	}
}
