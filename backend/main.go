package main

import (
	"fmt"
	"os"

	"github.com/Tesohh/xlearn/backend/data"
	"github.com/Tesohh/xlearn/backend/db"
	"github.com/Tesohh/xlearn/backend/handler"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	client, err := db.NewMongoClient(os.Getenv("DB_CONNECTION"))
	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	database := client.Database("v2")
	stores := &db.StoreHolder{
		Users: db.MongoStore[data.User]{Client: client, Coll: database.Collection("users")},
	}

	// populate context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := handler.CustomContext{Context: c, Stores: stores}
			return next(cc)
		}
	})

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			c.JSON(he.Code, echo.Map{"error": he.Message})
			e.Logger.Error(he)
			return
		}

		c.JSON(400, echo.Map{"error": "internal server error"})
		e.Logger.Error(err)
	}

	e.GET("/user/one/:usertag", handler.One(stores.Users, "usertag"))
	e.POST("/user/signup", handler.UserSignup)
	e.POST("/user/login", handler.UserLogin)

	fmt.Println("connected to mongo")
	e.Logger.Fatal(e.Start(":8080"))
}
