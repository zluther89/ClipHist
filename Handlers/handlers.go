package Handlers

import (
	"ClipHist/clipboard"
	"ClipHist/db"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	DB db.DB
}

func (h *Handler) ContentHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
		h.GetContent(w)
	case "POST":
		h.HandleContentPost(w, r)
	default:
		w.WriteHeader(400)
	}
}

func (h *Handler) HandleContentPost(w http.ResponseWriter, r *http.Request) {
	c, e := clipboard.Decode(r.Body)
	if e != nil {
		log.Printf("Post error: %v:", e)
		w.WriteHeader(400)
		return
	}
	c.Save()
	w.WriteHeader(200)
}

func (h *Handler) GetContent(w http.ResponseWriter) {
	encoder := json.NewEncoder(w)
	t, err := h.DB.SelectTop()
	if err != nil {
		log.Fatal(err)
	}
	err = encoder.Encode(t)
	if err != nil {
		log.Fatal(err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
