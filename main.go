package main

import (
	d "ClipHist/DBHelp"
	r "ClipHist/ReadClip"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var lastCopy string

func closeChan(c chan bool) {
	c <- true
}

func main() {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	done := make(chan bool)
	defer closeChan(done)

	db, error := sql.Open("sqlite3", "./sqliteDb/ClipHist.db")
	if error != nil {
		fmt.Println(error)
	}

	error = d.InitTable(db)
	if error != nil {
		fmt.Println(error)
	}

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case <-tick.C:
			clip := r.ReadClip()
			d.WriteHist(db, clip)
		}
	}

}