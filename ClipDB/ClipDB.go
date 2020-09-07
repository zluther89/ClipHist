//A Package to init and interact with a SQLite DB for Clip Histy
package ClipDB

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// A struct representing an entry in the D
// Note: db auto increments timestamp if not present
type ClipRow struct {
	Content, Timestamp string
}

var db *sql.DB

//Opens up a new CLipDB at the filepath string and inits
func Init(path string) error {
	var err error
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		fmt.Print(err)
		return err
	}
	statement, e := db.Prepare(initDBSQL)
	if e != nil {
		fmt.Println(e)
		return e
	}

	statement.Exec()
	return nil
}

// Writes the most recent content of clipboard to the db, ignores if the content hasn't changed
func Write(s string) (sql.Result, error) {
	statement, e := db.Prepare(insertSQL) // old "INSERT OR REPLACE INTO clip(content) VALUES(?)"
	if e != nil {
		fmt.Println(e)
		return nil, e
	}
	return statement.Exec(s)
}

// Reads most recent 25
func SelectTop() ([]ClipRow, error) {
	s := []ClipRow{}
	rows, err := db.Query(selectTopSQL) // old "SELECT * FROM clip ORDER BY rowid desc LIMIT 25;"
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		content, timestamp := "", ""
		if err := rows.Scan(&content, &timestamp); err != nil {
			fmt.Println(err)
		}
		cR := ClipRow{Content: content, Timestamp: timestamp}
		cR.formatTime()
		s = append(s, cR)
	}
	return s, err
}

// Finds a clip in the db by timestamp
func FindClip(s string) (string, error) {
	qStr := fmt.Sprintf(findSQL, s)
	row, e := db.Query(qStr)
	defer row.Close()
	if e != nil {
		return "", e
	}
	content := ""
	if e := row.Scan(&content); e != nil {
		return "", e
	}
	return content, nil
}

//formats the timestamp from the db into the format Mon Jan 2 15:04:05 2006
func (c *ClipRow) formatTime() {
	t, e := time.Parse(time.RFC3339, c.Timestamp)
	if e != nil {
		fmt.Println(e)
	}
	c.Timestamp = t.Format("Mon Jan 2 15:04:05 2006")
}
