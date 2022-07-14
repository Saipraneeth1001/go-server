package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Title string `json:"title"`
}


var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Title: "rrr"})
	movies = append(movies, Movie{ID: "2", Title: "kgf2"})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	fmt.Printf("Starting server on 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
