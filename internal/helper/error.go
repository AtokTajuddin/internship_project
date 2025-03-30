package helper

import "errors"

func NewError(message string) error {
	return errors.New(message)
}

func WrapError(err error, message string) error {
	return errors.New(message + ": " + err.Error())
}
