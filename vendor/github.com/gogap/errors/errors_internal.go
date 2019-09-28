package errors

import (
	"fmt"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
