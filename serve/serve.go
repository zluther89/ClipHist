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

func main() {
	db, _ = sql.Open("sqlite3", "../sqliteDb/ClipHist.db")

	http.Handle("/", http.FileServer(http.Dir("../ClipHistyFE/public")))
	http.HandleFunc("/hello", HelloServer)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
