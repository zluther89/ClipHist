//A Package to init and interact with a SQLite DB for Clip Histy
package db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Reads most recent 25
func (db DB) SelectTop() ([]ClipRow, error) {
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
func (db DB) FindClip(s string) (string, error) {
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
