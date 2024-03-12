package handlers

import (
	"gameQuiz/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type I map[string]interface{}

func LoginAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		password := c.FormValue("password")

		if name != "root" || password != "toor" {
			return c.JSON(http.StatusOK, I{
				"status": "error",
				"msg":    "Неверное имя или пароль",
			})
		}

		return c.JSON(http.StatusOK, I{
			"status": 200,
		})
	}
}

func AddQuestion() echo.HandlerFunc {
	return func(c echo.Context) error {
		question := c.FormValue("question")

		if question == "" {
			return c.JSON(http.StatusOK, I{
				"status": "error",
				"msg":    "Заполните поле 'Вопрос'",
			})
		}

		idQuestion, err := models.AddQuestion(question)
		if err != nil {
			panic(err)
		}

		list := models.GetQuestionsList().List

		return c.JSON(http.StatusOK, I{
			"status":     200,
			"idQuestion": idQuestion,
			"list":       list,
		})
	}
}

func GetQuestionsList() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, I{
			"status": 200,
			"list":   models.GetQuestionsList().List,
		})
	}
}
