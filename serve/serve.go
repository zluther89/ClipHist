package main

import (
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
		getContentAndWrite(w)
	case "POST":
		fmt.Println("post")
	}
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
