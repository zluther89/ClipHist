package main

import (
	"ClipHist/Clip"
	"ClipHist/ClipDB"
	"ClipHist/Notify"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Done chan bool

func (d Done) Stop() {
	d <- true
}

var lastClip string

func main() {
	done := Done(make(chan bool))
	r := make(chan bool)
	//defer done.Stop()

	//Notify.Init()

	err := ClipDB.Init("./sqliteDb/ClipHist.db")
	if err != nil {
		fmt.Println(err)
	}

	go Clip.ChanStart(r)
	go Notify.Recieve(r)

	// for {
	// 	select {
	// 	case <-r:
	// 		fmt.Println("testing chan")
	// 	}
	// }
	<-done

}
