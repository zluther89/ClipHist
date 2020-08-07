package DBHelp

import (
	"database/sql"
	"fmt"
)

var lastContent string

// Writes the most recent content of clipboard to the db, ignores if the content hasn't changed
func WriteHist(d *sql.DB, s string) {
	if lastContent == s {
		return
	}
	statement, e := d.Prepare("INSERT INTO clip(content) VALUES(?)")
	if e != nil {
		fmt.Println(e)
		return
	}
	statement.Exec(s)
	lastContent = s
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

type ClipRow struct {
	Content, Timestamp string
}

// Reads most recent 25
func SelectTopFromDB(d *sql.DB) []ClipRow {
	rows, e := d.Query("SELECT * FROM clip ORDER BY rowid desc LIMIT 25;")
	defer rows.Close()
	if e != nil {
		fmt.Println(e)
	}

	s := []ClipRow{}
	for rows.Next() {
		content, timestamp := "", ""
		err := rows.Scan(&content, &timestamp)
		if err != nil {
			fmt.Println(err)
		}
		cR := ClipRow{content, timestamp}
		s = append(s, cR)
	}
	return s
}

func FindClip(s string, d *sql.DB) (string, error) {
	qStr := fmt.Sprintf("SELECT content FROM clip WHERE timestampe = %v", s)
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
