package db

import (
	"database/sql"
	"fmt"
)

// A struct representing an entry in the D
// Note: db auto increments timestamp if not present
type ClipRow struct {
	Content, Timestamp string
}

type DB struct{ *sql.DB }

//Opens up a new CLipDB at the filepath string and inits
func Init(path string) (DB, error) {
	var err error
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		fmt.Print(err)
		return DB{nil}, err
	}
	statement, err := db.Prepare(initDBSQL)
	if err != nil {
		fmt.Println(err)
		return DB{nil}, err
	}

	statement.Exec()
	return DB{db}, nil
}
