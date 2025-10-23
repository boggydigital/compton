package compton

import (
	"errors"
)

func ErrUnknownToken(t string) error {
	return errors.New("unknown token: " + t)
}
