package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var dao = MoviesDAO{}

// AllMovies returns a list of all movies
func AllMovies(w http.ResponseWriter, r *http.Request) {
	message := GetMovies()
	fmt.Fprintln(w, message)
}

// InsertMovie inserts a single movie
func InsertMovie(w http.ResponseWriter, r *http.Request) {
	dao.Insert()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMovies).Methods("GET")
	r.HandleFunc("/movies", InsertMovie).Methods("POST")
	fmt.Printf("api running at port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
