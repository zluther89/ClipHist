package db

import (
	"database/sql"
	"fmt"
	"time"
)

// Writes the most recent content of clipboard to the db, ignores if the content hasn't changed
func (db DB) Write(s string) (sql.Result, error) {
	statement, e := db.Prepare(insertSQL) // old "INSERT OR REPLACE INTO clip(content) VALUES(?)"
	if e != nil {
		fmt.Println(e)
		return nil, e
	}
	return statement.Exec(s)
}

//formats the timestamp from the db into the format Mon Jan 2 15:04:05 2006
func (c *ClipRow) formatTime() {
	t, e := time.Parse(time.RFC3339, c.Timestamp)
	if e != nil {
		fmt.Println(e)
	}
	c.Timestamp = t.Format("Mon Jan 2 15:04:05 2006")
}
