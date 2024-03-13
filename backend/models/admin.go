package models

import (
	"database/sql"
	"gameQuiz/helpers"
)

type Question struct {
	ID       int
	Question string
}

type QuestionsList struct {
	List []Question
}

func AddQuestion(question string) (int64, error) {
	db, err := sql.Open("sqlite3", "storage.db")
	helpers.PanicHelper(err)

	sql := "INSERT INTO questions(question) VALUES (?)"
	stmt, err := db.Prepare(sql)
	helpers.PanicHelper(err)
	defer stmt.Close()

	result, err := stmt.Exec(question)
	helpers.PanicHelper(err)

	return result.LastInsertId()
}
