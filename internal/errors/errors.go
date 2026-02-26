package errors

import (
	"fmt"
)

func NewMismatchedFormatError(expected, actual string) error {
	return fmt.Errorf("Mismatched format: expected = %s, actual = %s", expected, actual)
}

func NewMissedFieldError(field string) error {
	return fmt.Errorf("Field %s is not presented", field)
}
