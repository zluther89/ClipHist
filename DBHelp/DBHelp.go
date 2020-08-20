package DBHelp

import (
	"database/sql"
	"fmt"
	"time"
)

// A struct representing an entry in the D
// Note: db auto increments timestamp if not present
type ClipRow struct {
	Content, Timestamp string
}

// Writes the most recent content of clipboard to the db, ignores if the content hasn't changed
func WriteHist(d *sql.DB, s string) {
	statement, e := d.Prepare("INSERT OR REPLACE INTO clip(content) VALUES(?)")
	if e != nil {
		fmt.Println(e)
		return
	}
	statement.Exec(s)
}

// Inits the table to store paste info in db
func InitTable(d *sql.DB) {
	statement, e := d.Prepare("CREATE TABLE IF NOT EXISTS clip(content string unique, timestamp datetime default current_timestamp)")
	if e != nil {
		fmt.Println(e)
		return
	}
	statement.Exec()
}

// Reads most recent 25
func SelectTopFromDB(d *sql.DB) []ClipRow {
	s := []ClipRow{}
	rows, e := d.Query("SELECT * FROM clip ORDER BY rowid desc LIMIT 25;")
	defer rows.Close()
	if e != nil {
		fmt.Println(e)
		return s
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
	return s
}

// Finds a clip in the db by timestamp
func FindClip(s string, d *sql.DB) (string, error) {
	qStr := fmt.Sprintf("SELECT content FROM clip WHERE timestamp = %v", s)
	row, e := d.Query(qStr)
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
