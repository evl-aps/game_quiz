package handlers

import (
	"gameQuiz/helpers"
	"gameQuiz/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type G map[string]interface{}

func StartGame() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")

		if name == "" {
			return c.JSON(http.StatusOK, G{
				"status": "error",
				"msg":    "Придумайте ваш ник!",
			})
		}

		users := models.GetName(name).List

		if len(users) > 0 {
			return c.JSON(http.StatusOK, G{
				"status": "error",
				"msg":    "Такой ник уже используется",
			})
		}

		userID, err := models.StartGame(name)
		helpers.PanicHelper(err)

		return c.JSON(http.StatusOK, G{
			"status": 200,
			"userID": userID,
		})
	}
}
