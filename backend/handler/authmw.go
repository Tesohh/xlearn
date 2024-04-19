package handler

import "github.com/labstack/echo/v4"

func AuthMW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(CustomContext)
		return next(cc)
	}
}
