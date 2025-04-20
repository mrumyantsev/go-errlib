package errlib_test

import (
	"errors"
	"testing"

	"github.com/mrumyantsev/go-errlib"
)

func TestWrap(t *testing.T) {
	err1 := errors.New("socket is in use")
	err2 := errlib.Wrap(err1, "could not start http server")
	err3 := errlib.Wrap(err2, "failed to run application")

	if !(errors.Is(err3, err2) && errors.Is(err3, err1)) {
		t.Fatalf("error '%v' does not contain 2 errors: '%v' and '%v'", err3, err2, err1)
	}
}
