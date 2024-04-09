package handler

import "fmt"

type APIError struct {
	Err    error
	Status int
}

func (a APIError) Error() string {
	return a.Err.Error()
}

func (a APIError) Context(s string) APIError {
	return APIError{fmt.Errorf("%w (%s)", a.Err, s), a.Status}
}
