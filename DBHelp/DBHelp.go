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
	}
	statement.Exec(s)
	lastContent = s

}

// Inits the table to store paste info in db
func InitTable(d *sql.DB) {
	statement, e := d.Prepare("CREATE TABLE IF NOT EXISTS clip(content string unique, timestamp datetime default current_timestamp)")
	if e != nil {
		fmt.Println(e)
	}
	statement.Exec()
}
