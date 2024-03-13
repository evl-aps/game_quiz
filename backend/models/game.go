package models

import (
	"database/sql"
	"gameQuiz/helpers"
)

type User struct {
	ID   int
	Name string
}

type UserList struct {
	List []User
}

func StartGame(name string) (int64, error) {
	db, err := sql.Open("sqlite3", "storage.db")
	helpers.PanicHelper(err)

	sql := "INSERT INTO users(name) VALUES(?)"
	smtm, err := db.Prepare(sql)
	helpers.PanicHelper(err)
	defer db.Close()

	result, err := smtm.Exec(name)
	helpers.PanicHelper(err)

	return result.LastInsertId()
}

func GetName(name string) UserList {
	db, err := sql.Open("sqlite3", "storage.db")
	helpers.PanicHelper(err)

	sql := "SELECT * FROM users WHERE name = ?"
	rows, err := db.Query(sql, name)
	helpers.PanicHelper(err)

	result := UserList{}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name)
		helpers.PanicHelper(err)
		result.List = append(result.List, user)
	}

	return result
}
