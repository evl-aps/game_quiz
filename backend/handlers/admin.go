package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type I map[string]interface{}

type Question struct {
	ID       int
	Question string
}

type QuestionsList struct {
	List []Question
}

func LoginAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
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

		db, err := sql.Open("sqlite3", "storage.db")
		if err != nil {
			panic(err)
		}

		sql := "INSERT INTO questions(question) VALUES (?)"
		stmt, err := db.Prepare(sql)
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		result, err := stmt.Exec(question)
		if err != nil {
			panic(err)
		}

		return c.JSON(http.StatusOK, I{
			"status":     200,
			"idQuestion": result,
		})
	}
}

func GetQuestionsList() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := sql.Open("sqlite3", "storage.db")
		if err != nil {
			panic(err)
		}
		sql := "SELECT * FROM questions"
		rows, err := db.Query(sql)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		result := QuestionsList{}

		for rows.Next() {
			question := Question{}
			err = rows.Scan(&question.ID, &question.Question)
			if err != nil {
				panic(err)
			}

			result.List = append(result.List, question)
		}

		return c.JSON(http.StatusOK, I{
			"status": 200,
			"list":   result.List,
		})
	}
}
