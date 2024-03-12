package main

import (
	"gameQuiz/handlers"
	"gameQuiz/migration"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	migration.InitDataBase()

	e := echo.New()
	e.Use(middleware.CORS())

	e.POST("/start-game", handlers.StartGame())

	e.POST("/admin", handlers.LoginAdmin())
	e.POST("/admin/add-question", handlers.AddQuestion())
	e.GET("/get-question", handlers.GetQuestionsList())

	e.Logger.Fatal(e.Start(":8000"))
}
