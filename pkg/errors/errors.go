package errors

import (
	"fmt"
)

func WrapFail(err error, msg string) error {
	return fmt.Errorf("could not %s: %e", msg, err)
}
