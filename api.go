package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var dao = MoviesDAO{}

func init() {
	dao.Database = "test"
	dao.Server = "localhost"
	dao.Connect()
}

// AllMovies returns a list of all movies
func AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(w, http.StatusOK, movies)
}

// InsertMovie inserts a single movie
func InsertMovie(w http.ResponseWriter, r *http.Request) {
	dao.Insert()
	respondWithJson(w, http.StatusOK, "")
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

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
