package main

import (
	"ClipHist/Handlers"
	"ClipHist/clipboard"
	"ClipHist/db"
	"fmt"
	"log"
	"net/http"
	"time"
)

var port string = ":3000"

type message struct {
	Message string `json:"message"`
}

func main() {
	db, err := db.Init("../sqliteDb/ClipHist.db")
	if err != nil {
		fmt.Println(err)
	}
	h := Handlers.Handler{DB: db}
	c := clipboard.Init(time.Second)

	go c.StartListener(db)
	http.Handle("/", http.FileServer(http.Dir("../View/public")))
	http.HandleFunc("/content", h.ContentHandler)
	log.Print("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
