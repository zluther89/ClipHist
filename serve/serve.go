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

type message struct {
	Message string `json:"message"`
}

func main() {
	//notify := make(chan bool)
	db, err := ClipDB.Init("../sqliteDb/ClipHist.db")
	if err != nil {
		fmt.Println(err)
	}
	h := Handlers.Handler{DB: db}
	c := Clip.Init()
	lastClip := ""
	go c.Start(db, &lastClip)
	//go h.Listen(notify)
	connCount := 0

	http.Handle("/", http.FileServer(http.Dir("../ClipHistyFE/public")))
	http.HandleFunc("/content", h.ContentHandler)
	http.Handle("/socket", websocket.Handler(func(ws *websocket.Conn) {
		if connCount >= 1 {
			ws.Close()
			return
		}

		connCount += 1
		var m message

		for {
			go c.Listen(ws, "a change to db")
			// receive a message using the codec
			if err := websocket.JSON.Receive(ws, &m); err != nil {
				log.Println("websocket recieve:", err)
				connCount -= 1
				return
			}
			log.Println("Received message:", m.Message)
		}

	}))

	log.Print("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
