package main

import (
	"ClipHist/Clip"
	"ClipHist/ClipDB"
	"encoding/json"
	"log"
	"net/http"
)

var port string = ":3000"

func main() {
	if err := ClipDB.Init("../sqliteDb/ClipHist.db"); err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.Dir("../ClipHistyFE/public")))
	http.HandleFunc("/content", ContentHandler)

	log.Print("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func ContentHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
		if err := getContent(w); err != nil {
			log.Fatal(err)
		}

	case "POST":
		handleContentPost(w, r)
	}
}

func handleContentPost(w http.ResponseWriter, r *http.Request) {
	c, e := decodeJsonRContent(r)
	if e != nil {
		log.Printf("Post error: %v:", e)
		w.WriteHeader(400)
		return
	}
	c.Save()
	w.WriteHeader(200)
}

func decodeJsonRContent(r *http.Request) (Clip.ClipEntry, error) {
	dec := json.NewDecoder(r.Body)
	var rB Clip.ClipEntry
	if e := dec.Decode(&rB); e != nil {
		return Clip.ClipEntry{}, e
	}

	return rB, nil
}

func getContent(w http.ResponseWriter) error {
	encoder := json.NewEncoder(w)
	t, err := ClipDB.SelectTopFromDB()
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
