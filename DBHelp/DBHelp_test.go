package DBHelp

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSelectTopFromDB(t *testing.T) {

	db, err := sql.Open("sqlite3", "../sqliteDb/ClipHist.db")
	if err != nil {
		t.Error("Error Opening DB", err)
	}

	i := SelectTopFromDB(db)
	if len(i) > 25 {
		t.Errorf("Error, expected max 25 responses but got %v", len(i))
	}

}

//Testing init of a new db if one does not exist
func TestInitTable(t *testing.T) {
	db, err := sql.Open("sqlite3", "../TestDB/ClipHist.db")
	if err != nil {
		t.Error("Error Opening DB", err)
	}

	if e := InitTable(db); e != nil {
		t.Errorf("Error creating new db:%v", e)
	}

}
