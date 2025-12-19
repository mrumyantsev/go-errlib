package errlib_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mrumyantsev/go-errlib"
)

var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
	ErrC = errors.New("c")
	ErrD = errors.New("d")
	ErrE = errors.New("e")
	ErrF = errors.New("f")
)

func TestWrap(t *testing.T) {
	tests := []struct {
		action func(err error) error
		expect string
	}{
		{func(err error) error { return errlib.Wrap(err) }, "<nil>"},
		{func(err error) error { return errlib.Wrap(err, ErrA) }, "a"},
		{func(err error) error { return errlib.Wrap(err, ErrB, ErrC) }, "c: b: a"},
		{func(err error) error { return errlib.Wrap(err, ErrD, ErrE, ErrF) }, "f: e: d: c: b: a"},
		{func(err error) error { return errors.Unwrap(err) }, "e: d: c: b: a"},
		{func(err error) error { return errors.Unwrap(err) }, "d: c: b: a"},
		{func(err error) error { return errors.Unwrap(err) }, "c: b: a"},
		{func(err error) error { return errors.Unwrap(err) }, "b: a"},
		{func(err error) error { return errors.Unwrap(err) }, "a"},
		{func(err error) error { return errors.Unwrap(err) }, "<nil>"},
		{func(err error) error { return errlib.Wrap(err, ErrF, ErrA) }, "a: f"},
		{func(err error) error { return errlib.Wrap(err, ErrC, "foo", "bar", ErrD) }, "d: bar: foo: c: a: f"},
		{func(err error) error { return errors.Unwrap(err) }, "bar: foo: c: a: f"},
		{func(err error) error { return errors.Unwrap(err) }, "foo: c: a: f"},
		{func(err error) error { return errors.Unwrap(err) }, "c: a: f"},
		{func(err error) error { return errors.Unwrap(err) }, "a: f"},
		{func(err error) error { return errors.Unwrap(err) }, "f"},
		{func(err error) error { return errors.Unwrap(err) }, "<nil>"},
	}

	var err error = nil

	for i := range tests {
		err = tests[i].action(err)

		if fmt.Sprintf("%v", err) != tests[i].expect {
			t.Fatalf("got: %v, expected: %s", err, tests[i].expect)
		}
	}
}
