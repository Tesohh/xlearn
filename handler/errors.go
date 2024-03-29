package handler

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidPassword               = APIError{errors.New("password invalid"), http.StatusForbidden}
	ErrEmptyBody                     = APIError{errors.New("the request body is empty"), http.StatusBadRequest}
	ErrUsernameTaken                 = APIError{errors.New("username is taken"), http.StatusNotAcceptable}
	ErrTagTaken                      = APIError{errors.New("tag is taken"), http.StatusNotAcceptable}
	ErrMalformedBody                 = APIError{errors.New("malformed body"), http.StatusBadRequest}
	ErrPWTooShort                    = APIError{errors.New("password is too short. must be at least 12 characters"), http.StatusLengthRequired}
	ErrJwtUsernameInexistent         = APIError{errors.New("jwt-username header doesn't exist somehow"), http.StatusBadRequest}
	ErrUnauthorized                  = APIError{errors.New("unauthorized"), http.StatusUnauthorized}
	ErrPathVar                       = APIError{errors.New("invalid path variable"), http.StatusBadRequest}
	ErrOutOfRange                    = APIError{errors.New("index out of range or negative"), http.StatusBadRequest}
	ErrInvalidParentPrefix           = APIError{errors.New("invalid parent prefix"), http.StatusBadRequest}
	ErrInvalidOperation              = APIError{errors.New("invalid operation"), http.StatusBadRequest}
	ErrRequestedItemInexistent       = APIError{errors.New("requested item inexistent"), http.StatusNotFound}
	ErrAlreadyJoinedOrg              = APIError{errors.New("user already joined this org"), http.StatusConflict}
	ErrDatabaseNotEmpty              = APIError{errors.New("database is not empty or was impossible to connect to"), http.StatusConflict}
	ErrLanguageInvalid               = APIError{errors.New("invalid language"), http.StatusNotAcceptable}
	ErrAlreadyCompletedStep          = APIError{errors.New("user already completed this step"), http.StatusConflict}
	ErrNoStepsEverCompleted          = APIError{errors.New("user completed array is empty"), http.StatusNotFound}
	ErrNoStepsCompletedFromAdventure = APIError{errors.New("can't find any completed steps from that adventure"), http.StatusNotFound}
	ErrWrongPin                      = APIError{errors.New("incorrect pin"), http.StatusUnauthorized}
	ErrTooManyAttempts               = APIError{errors.New("too many attempts"), http.StatusUnauthorized}
)
