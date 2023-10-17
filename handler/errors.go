package handler

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidPassword         = APIError{errors.New("password invalid"), http.StatusForbidden}
	ErrEmptyBody               = APIError{errors.New("the request body is empty"), http.StatusBadRequest}
	ErrUsernameTaken           = errors.New("username is taken")
	ErrTagTaken                = errors.New("tag is taken")
	ErrMalformedBody           = APIError{errors.New("malformed body"), http.StatusBadRequest}
	ErrPWTooShort              = APIError{errors.New("password is too short. must be at least 12 characters"), http.StatusLengthRequired}
	ErrJwtUsernameInexistent   = errors.New("jwt-username header doesn't exist somehow")
	ErrUnauthorized            = APIError{errors.New("unauthorized"), http.StatusUnauthorized}
	ErrPathVar                 = errors.New("invalid path variable")
	ErrOutOfRange              = errors.New("index out of range or negative")
	ErrInvalidParentPrefix     = errors.New("invalid parent prefix")
	ErrInvalidOperation        = errors.New("invalid operation")
	ErrRequestedItemInexistent = errors.New("requested item inexistent")
)
