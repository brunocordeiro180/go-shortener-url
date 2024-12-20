package db

import (
	"database/sql"
)

const file = "./urls.db"

func MigrateTables(db *sql.DB) error {
	const create string = `
		create table if not exists urls (
		    id integer primary key autoincrement,
		    url varchar(256) unique not null ,
		    key varchar(100) unique not null  
		)
	`

	if _, err := db.Exec(create); err != nil {
		return err
	}

	return nil
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	return db, err
}
