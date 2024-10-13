package errlib

import "fmt"

// Wrap wraps the given error with the given message. Literally it
// makes '2 in 1' error and it could be used to create the chains of
// errors which are comparable by using standard errors.Is function.
func Wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}
