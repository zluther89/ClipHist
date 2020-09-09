package main

import (
	"ClipHist/Clip"
	"ClipHist/ClipDB"
	"ClipHist/Handlers"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var port string = ":3000"

func main() {
	notify := make(chan bool)
	db, err := ClipDB.Init("../sqliteDb/ClipHist.db")
	if err != nil {
		fmt.Println(err)
	}
	h := Handlers.Handler{DB: db}

	go Clip.ChanStart(notify, db)
	go h.Listen(notify)

	http.Handle("/", http.FileServer(http.Dir("../ClipHistyFE/public")))
	http.HandleFunc("/content", h.ContentHandler)
	http.Handle("/socket", websocket.Handler(h.Socket))

	log.Print("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
