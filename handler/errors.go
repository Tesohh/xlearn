package handler

import "errors"

var (
	ErrInvalidPassword       = errors.New("password invalid")
	ErrEmptyBody             = errors.New("the request body is empty")
	ErrUsernameTaken         = errors.New("username is taken")
	ErrMalformedBody         = errors.New("malformed body")
	ErrPWTooShort            = errors.New("password is too short. must be at least 12 characters")
	ErrJwtUsernameInexistent = errors.New("jwt-username header doesn't exist somehow")
	ErrUnauthorized          = errors.New("unauthorized")
)
