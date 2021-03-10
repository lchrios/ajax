package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func vuejs(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "vue01.html")
}

// Film is the struct containing film data
type Film struct {
	Title    string
	Director string
	Year     int
}

func sendFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var films = [4]Film{
		{"No Manches Frida 2", "Nacho G. Velilla", 2016},
		{"Star Wars 6", "George Lucas", 1976},
		{"Harry Potter 3", "Alfonso Cuarón", 2002},
		{"Elba lazo", "Gillermo del Toro", 2009},
	}

	json.NewEncoder(w).Encode(films)
}

func main() {
	http.HandleFunc("/vue", vuejs)
	http.HandleFunc("/films", sendFilms)
	fmt.Println(time.Now().Format("02-01-2006 15:04:05"))
	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}
