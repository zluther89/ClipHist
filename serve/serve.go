package main

import (
	"ClipHist/ClipDB"
	"ClipHist/Handlers"
	"log"
	"net/http"
)

var port string = ":3000"

func main() {

	db, err := ClipDB.Init("../sqliteDb/ClipHist.db")
	if err != nil {
		log.Fatal(err)
	}
	h := Handlers.Handler{DB: db}

	http.Handle("/", http.FileServer(http.Dir("../ClipHistyFE/public")))
	http.HandleFunc("/content", h.ContentHandler)

	log.Print("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
