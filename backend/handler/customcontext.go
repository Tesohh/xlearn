package handler

import (
	"github.com/Tesohh/xlearn/backend/db"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Stores *db.StoreHolder
}
