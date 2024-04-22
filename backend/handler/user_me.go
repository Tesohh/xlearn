package handler

import "github.com/labstack/echo/v4"

func UserMe(c echo.Context) error {
	cc := c.(CustomContext)
	return c.JSON(200, cc.User)
}
