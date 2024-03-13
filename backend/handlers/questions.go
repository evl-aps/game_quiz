package handlers

import (
	"gameQuiz/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetQuestionsList() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, I{
			"status": 200,
			"list":   models.GetQuestionsList().List,
		})
	}
}
