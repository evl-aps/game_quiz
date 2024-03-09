package main

import (
	"gameQuiz/migrations"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type I map[string]interface{}

func main() {
	migrations.InitDataBase()

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, I{
			"Hello": "World!",
		})
	})

	e.POST("/admin", func(c echo.Context) error {
		name := c.FormValue("name")
		password := c.FormValue("password")

		if name != "root" || password != "toor" {
			return c.JSON(http.StatusOK, I{
				"status": 401,
				"msg":    "Неверное имя или пароль",
			})
		}

		return c.JSON(http.StatusOK, I{
			"status": 200,
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
