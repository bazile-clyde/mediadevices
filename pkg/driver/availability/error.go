package availability

import (
	"errors"
)

var (
	ErrUnimplemented = errors.New("not implemented")
	ErrBusy          = errors.New("device or resource busy")
	ErrNoDevice      = errors.New("no such device")
)

type errorString struct {
	s string
}

func NewError(text string) error {
	return &errorString{text}
}

func (e *errorString) Error() string {
	return e.s
}
