package migrations

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func UserInit() {
	db, err := sql.Open("sqlite3", "storage.db")
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db is nil!")
	}
	migrate(db)
}

func migrate(db *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS users
			(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL);`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
