package Handlers

import (
	"ClipHist/Clip"
	"ClipHist/ClipDB"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type Handler struct {
	DB     ClipDB.DB
	change bool
}

func (h *Handler) ContentHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
		if err := h.GetContent(w); err != nil {
			log.Fatal(err)
		}

	case "POST":
		h.HandleContentPost(w, r)
	}
}

func (h *Handler) HandleContentPost(w http.ResponseWriter, r *http.Request) {
	c, e := Clip.DecodeClip(r.Body)
	if e != nil {
		log.Printf("Post error: %v:", e)
		w.WriteHeader(400)
		return
	}
	c.Save()
	w.WriteHeader(200)
}

func (h *Handler) GetContent(w http.ResponseWriter) error {
	encoder := json.NewEncoder(w)
	t, err := h.DB.SelectTop()
	if err != nil {
		return err
	}
	err = encoder.Encode(t)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return nil
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type message struct {
	Message string `json:"message"`
}

// try to use listener to change a boolean on the handler struct
// use the loop in the socket func to 'listen' for changes to the boolean and send messages to client
func (h *Handler) Listen(c <-chan bool) {
	for _ = range c {
		fmt.Println("Test Recieve func")
		fmt.Print(h.change)
		if h.change == true {
			h.change = false
		} else {
			h.change = true
		}
	}
}

func (h *Handler) Socket(ws *websocket.Conn) {
	for {
		// allocate our container struct
		var m message
		// receive a message using the codec
		if err := websocket.JSON.Receive(ws, &m); err != nil {
			log.Println(err)
			break
		}
		log.Println("Received message:", m.Message)
		// send a response
		m2 := message{"Thanks for the message!"}
		if err := websocket.JSON.Send(ws, m2); err != nil {
			log.Println(err)
			break
		}
		m3 := message{"A CHANGEE"}
		if h.change == true {
			if err := websocket.JSON.Send(ws, m3); err != nil {
				log.Println(err)
				break
			}
		}
	}
}

//func handleNotify (alert chan,socket){ send warning on a socket when alert chan is notified}
