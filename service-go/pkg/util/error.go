package util

import (
	"fmt"

	"github.com/pkg/errors"
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int, message string, arguments ...interface{}) *Error {
	msg := message
	if len(arguments) > 0 {
		msg = fmt.Sprintf(message, arguments...)
	}

	return &Error{
		Code:    code,
		Message: msg,
	}
}

func IsError(err error) bool {
	return errors.Is(err, &Error{})
}
