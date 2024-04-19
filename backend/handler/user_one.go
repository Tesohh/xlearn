package handler

import (
	"github.com/Tesohh/xlearn/backend/db"
	"github.com/labstack/echo/v4"
)

func One[T any](store db.Storer[T], tagpath string) echo.HandlerFunc {
	return func(c echo.Context) error {
		tag := c.Param(tagpath)
		doc, err := store.One(db.Query{"tag": tag})
		if err != nil {
			return err
		}

		return c.JSON(200, doc)
	}
}
