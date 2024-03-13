package models

import (
	"database/sql"
	"gameQuiz/helpers"
)

func GetQuestionsList() QuestionsList {
	db, err := sql.Open("sqlite3", "storage.db")
	helpers.PanicHelper(err)
	sql := "SELECT * FROM questions"
	rows, err := db.Query(sql)
	helpers.PanicHelper(err)
	defer rows.Close()

	result := QuestionsList{}

	for rows.Next() {
		question := Question{}
		err = rows.Scan(&question.ID, &question.Question)
		helpers.PanicHelper(err)

		result.List = append(result.List, question)
	}

	return result
}
