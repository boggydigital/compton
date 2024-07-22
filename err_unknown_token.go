package compton

import "fmt"

func ErrUnknownToken(t string) error {
	return fmt.Errorf("unknown token: %s", t)
}
