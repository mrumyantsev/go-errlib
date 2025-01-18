package errlib

import "fmt"

// Wrap wraps the given error with the given message. It is a shortcut
// for fmt.Errorf("some error message: %w", err).
func Wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}
