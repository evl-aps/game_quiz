package migration

import (
	"database/sql"
	"gameQuiz/helpers"

	_ "github.com/mattn/go-sqlite3"
)

func InitDataBase() {
	db, err := sql.Open("sqlite3", "storage.db")
	helpers.PanicHelper(err)
	if db == nil {
		panic("db is nil")
	}

	migrateUsersTable(db)
	migrateQuestionTable(db)
	migrateAnswerTable(db)
}

func migrateUsersTable(db *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS users
			(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL);`
	_, err := db.Exec(sql)
	helpers.PanicHelper(err)
}

func migrateQuestionTable(db *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS questions
			(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			question VARCHAR NOT NULL)`
	_, err := db.Exec(sql)
	helpers.PanicHelper(err)
}

func migrateAnswerTable(db *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS asnwers
			(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			id_question INTEGER NOT NULL,
			answer_text VARCHAR NOT NULL)`
	_, err := db.Exec(sql)
	helpers.PanicHelper(err)
}
