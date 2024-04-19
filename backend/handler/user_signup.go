package handler

import (
	"net/http"
	"strings"

	"github.com/Tesohh/xlearn/backend/data"
	"github.com/Tesohh/xlearn/backend/db"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrTagTooShort       = echo.NewHTTPError(http.StatusNotAcceptable, "tag is too short")
	ErrTagTooLong        = echo.NewHTTPError(http.StatusNotAcceptable, "tag is too long")
	ErrPasswordTooShort  = echo.NewHTTPError(http.StatusNotAcceptable, "password is too short")
	ErrPasswordTooLong   = echo.NewHTTPError(http.StatusNotAcceptable, "password is too long")
	ErrUserAlreadyExists = echo.NewHTTPError(http.StatusForbidden, "user with that username already exists")
)

type signupBody struct {
	Tag      string `json:"tag"`
	Password string `json:"password"`
}

func UserSignup(c echo.Context) error {
	// boilerplate
	cc := c.(CustomContext)
	body, err := Decode[signupBody](cc.Request().Body)
	if err != nil {
		return err
	}

	// check password
	if len(body.Password) < 8 {
		return ErrPasswordTooShort
	}
	if len([]byte(body.Password)) >= 72 {
		return ErrPasswordTooLong
	}

	// check tag
	body.Tag = strings.ReplaceAll(body.Tag, " ", "-")
	if len(body.Tag) < 3 {
		return ErrTagTooShort
	}
	if len(body.Tag) > 20 {
		return ErrTagTooLong
	}

	// check if user doesn't already exist
	_, err = cc.Stores.Users.One(db.Query{"tag": body.Tag})
	if err == nil {
		return ErrUserAlreadyExists
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		return err
	}

	// add to db
	user := data.User{
		Tag:          body.Tag,
		PasswordHash: string(hash),
		Role:         data.RoleUser,
	}
	err = cc.Stores.Users.Put(user)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}
