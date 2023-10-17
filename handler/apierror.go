package handler

type APIError struct {
	Err    error
	Status int
}

func (a APIError) Error() string {
	return a.Err.Error()
}
