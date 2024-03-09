package main

import (
	"gameQuiz/admin"
	"gameQuiz/migrations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	migrations.InitDataBase()

	e := echo.New()
	e.Use(middleware.CORS())

	e.POST("/admin", admin.LoginAdmin())
	e.POST("/admin/add-question", admin.AddQuestion())

	e.Logger.Fatal(e.Start(":8000"))
}
