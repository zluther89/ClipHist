package main

import (
	d "ClipHist/DBHelp"
	ReadClip "ClipHist/ReadClip"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type reqBody struct {
	Content, Timestamp string
}

var db *sql.DB

var port string = ":3000"

func main() {
	db, _ = sql.Open("sqlite3", "../sqliteDb/ClipHist.db")

	http.Handle("/", http.FileServer(http.Dir("../ClipHistyFE/public")))
	http.HandleFunc("/content", contentRouter)

	log.Print("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func contentRouter(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
		s := r.URL.Query().Get("search")
		fmt.Println(s == "")
		getContentAndWrite(w)
	case "POST":
		handleContentPost(w, r)
	}
}

func handleContentPost(w http.ResponseWriter, r *http.Request) {
	if e := r.ParseForm(); e != nil {
		fmt.Printf("ParseForm() err: %v", e)
		return
	}
	content := decodeJsonRContent(r)
	if e := ReadClip.WriteToClip(content); e != nil {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
	}
}

func decodeJsonRContent(r *http.Request) string {
	dec := json.NewDecoder(r.Body)
	var rB reqBody
	if e := dec.Decode(&rB); e != nil {
		fmt.Println(e)
		return ""
	}
	return rB.Content

}

func getContentAndWrite(w http.ResponseWriter) {
	t := d.SelectTopFromDB(db)
	j, e := json.Marshal(t)
	if e != nil {
		fmt.Println(e)
	}
	w.Write(j)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
