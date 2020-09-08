package Handlers

import (
	"ClipHist/Clip"
	"ClipHist/ClipDB"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	DB      ClipDB.DB
	Encoder *json.Encoder
}

func (h Handler) ContentHandler(w http.ResponseWriter, r *http.Request) {
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

func (h Handler) HandleContentPost(w http.ResponseWriter, r *http.Request) {
	c, e := Clip.DecodeClip(r.Body)
	if e != nil {
		log.Printf("Post error: %v:", e)
		w.WriteHeader(400)
		return
	}
	c.Save()
	w.WriteHeader(200)
}

func (h Handler) GetContent(w http.ResponseWriter) error {
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

//func handleNotify (alert chan,socket){ send warning on a socket when alert chan is notified}
