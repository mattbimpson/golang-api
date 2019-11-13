package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

var movieList []Movie

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movies", InsertMovie).Methods("POST")
	r.Use(loggerMiddleware)
	fmt.Printf("api running at port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	// movies := [2]Movie {
	// 	Movie{ ID: "0", Name: "First movie", Description: "The first one" },
	// 	Movie{ ID: "1", Name: "Second movie", Description: "The second one"},
	// }
	respondWithJSON(w, http.StatusOK, movieList)
}

func InsertMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	movieList = append(movieList, movie)
	// just return the movie with the new id for now
	movie.ID = bson.NewObjectId()
	respondWithJSON(w, http.StatusCreated, movie)
}
