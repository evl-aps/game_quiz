package models

import "database/sql"

type Question struct {
	ID       int
	Question string
}

type QuestionsList struct {
	List []Question
}

func AddQuestion(question string) (int64, error) {
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

	return result.LastInsertId()
}

func GetQuestionsList() QuestionsList {
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

	return result
}
