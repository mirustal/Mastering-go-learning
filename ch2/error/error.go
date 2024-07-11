package error

import (
	"errors"
	"fmt"
	"os"
)

func Check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("this is custom error massage")
	}
	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Getuid())
	}
	return nil
}