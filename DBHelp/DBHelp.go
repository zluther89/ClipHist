package DBHelp

import (
	"database/sql"
	"fmt"
)

var lastContent string

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

func InitTable(d *sql.DB) error {
	statement, e := d.Prepare("CREATE TABLE IF NOT EXISTS clip(content string unique, timestamp datetime default current_timestamp)")
	if e != nil {
		return e
	}
	statement.Exec()
	return nil
}
