package main

import (
	"ClipHist/Clip"
	"ClipHist/DBHelp"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Done chan bool

func (d Done) Stop() {
	d <- true
}

var lastClip string

func main() {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	done := Done(make(chan bool))
	defer done.Stop()

	db, error := sql.Open("sqlite3", "./sqliteDb/ClipHist.db")
	if error != nil {
		fmt.Println(error)
	}

	if err := DBHelp.InitTable(db); err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case <-tick.C:
			if clip := Clip.ReadClip(); clip != lastClip && clip != "" {
				lastClip = clip
				DBHelp.WriteHist(db, clip)
			}
		}
	}

}
