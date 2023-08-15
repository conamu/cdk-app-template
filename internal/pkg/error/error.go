package errors

import "errors"

func Wrap(msg string, err error) error {
	if err == nil {
		return nil
	}
	msg += ": "
	return errors.Join(errors.New(msg), err)
}
