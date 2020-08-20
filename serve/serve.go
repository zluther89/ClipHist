package main

import (
	Clip "ClipHist/Clip"
	d "ClipHist/DBHelp"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

var port string = ":3000"

func main() {
	var err error
	db, err = sql.Open("sqlite3", "../sqliteDb/ClipHist.db")
	if err != nil {
		fmt.Println(err)
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
		c := getContent(db)
		w.Write(c)
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

func getContent(db *sql.DB) []byte {
	t := d.SelectTopFromDB(db)
	j, e := json.Marshal(t)
	if e != nil {
		fmt.Println(e)
	}
	return j
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
