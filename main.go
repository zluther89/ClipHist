package main

import (
	"ClipHist/Clip"
	"ClipHist/ClipDB"
	"ClipHist/Notify"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var lastClip string

func main() {
	done := make(chan bool)
	//channel torecieve notifications
	r := make(chan bool)

	if err := ClipDB.Init("./sqliteDb/ClipHist.db"); err != nil {
		fmt.Println(err)
	}

	go Clip.ChanStart(r)

	go Notify.Recieve(r)

	//keep main alive
	<-done

}
