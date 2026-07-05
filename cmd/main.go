package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /movies", listMovies)

	mux.Handle("GET /", http.FileServer(http.Dir("static")))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}

var movies = []movieResponse{
	{ID: "fightClub", Title: "Fight Club", Rows: 6, SeatsPerRow: 8},
	{ID: "dune", Title: "Dune: Part One", Rows: 8, SeatsPerRow: 10},
}

func listMovies(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, movies)
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

type movieResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Rows        int    `json:"rows"`
	SeatsPerRow int    `json:"seats_per_row"`
}
