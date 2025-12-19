package errlib

import (
	"errors"
	"fmt"
	"reflect"
)

// Wrap wraps the error with other errors or strings and returns an
// error that is unwrappable by errors.Unwrap() function, and also
// comparable with errors.Is() and errors.As() functions.
//
// The result error string will be like: "c: b: a".
//
// If passed a variable of non-string and non-error type as a wrapper,
// it will raise panic.
func Wrap(err error, wrappers ...interface{}) error {
	length := len(wrappers)

	if length == 0 {
		return err
	}

	for i := 0; i < length; i++ {
		err = wrap(err, wrappers[i])
	}

	return err
}

func wrap(err error, wrapper interface{}) error {
	if wrapper == nil {
		return error(nil)
	}

	if err == nil {
		return toError(wrapper)
	}

	return fmt.Errorf("%s: %w", toString(wrapper), err)
}

func toString(val interface{}) string {
	switch val := val.(type) {
	case error:
		return val.Error()
	case string:
		return val
	default:
		doPanic(val)
	}

	return ""
}

func toError(val interface{}) error {
	switch val := val.(type) {
	case error:
		return val
	case string:
		return errors.New(val)
	default:
		doPanic(val)
	}

	return nil
}

func doPanic(val interface{}) {
	panic("wrap: must be passed error or string type, but passed " +
		reflect.TypeOf(val).Kind().String())
}
