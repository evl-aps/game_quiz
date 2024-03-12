package models

import "database/sql"

type User struct {
	ID   int
	Name string
}

type UserList struct {
	List []User
}

func StartGame(name string) (int64, error) {
	db, err := sql.Open("sqlite3", "storage.db")
	if err != nil {
		panic(err)
	}

	sql := "INSERT INTO users(name) VALUES(?)"
	smtm, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := smtm.Exec(name)
	if err != nil {
		panic(err)
	}

	return result.LastInsertId()
}

func GetName(name string) UserList {
	db, err := sql.Open("sqlite3", "storage.db")
	if err != nil {
		panic(err)
	}

	sql := "SELECT * FROM users WHERE name = ?"
	rows, err := db.Query(sql, name)
	if err != nil {
		panic(err)
	}

	result := UserList{}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err)
		}
		result.List = append(result.List, user)
	}

	return result
}
