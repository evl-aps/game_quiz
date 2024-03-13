package handlers

import (
	"encoding/json"
	"gameQuiz/helpers"
	"gameQuiz/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type I map[string]interface{}

type Answers struct {
	Answers []string
}

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
		helpers.PanicHelper(err)

		list := models.GetQuestionsList().List

		return c.JSON(http.StatusOK, I{
			"status":     200,
			"idQuestion": idQuestion,
			"list":       list,
		})
	}
}

func addAnswers() echo.HandlerFunc {
	return func(c echo.Context) error {
		answers := Answers{}
		err := json.Unmarshal([]byte(c.FormValue("answers")), &answers)
		helpers.PanicHelper(err)
		return c.JSON(http.StatusOK, I{})
	}
}
