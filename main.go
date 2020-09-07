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
	r := make(chan bool)

	err := ClipDB.Init("./sqliteDb/ClipHist.db")
	if err != nil {
		fmt.Println(err)
	}

	go Clip.ChanStart(r)
	go Notify.Recieve(r)

	<-done

}
