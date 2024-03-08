package main

import (
	"gameQuiz/migrations"
	"net/http"

	"github.com/labstack/echo/v4"
)

type I map[string]interface{}

func main() {
	migrations.InitDataBase()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, I{
			"Hello": "World!",
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
