package DBHelp

import "database/sql"

var lastContent string

func WriteHist(d *sql.DB, s string) error {
	if lastContent == s {
		return nil
	}
	statement, e := d.Prepare("INSERT INTO clip(content) VALUES(?)")
	if e != nil {
		return e
	}
	statement.Exec(s)
	lastContent = s
	return nil
}

func InitTable(d *sql.DB) error {
	statement, e := d.Prepare("CREATE TABLE IF NOT EXISTS clip(content string unique, timestampe datetime default current_timestamp)")
	if e != nil {
		return e
	}
	statement.Exec()
	return e
}
