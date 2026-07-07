package main

import (
	"CinemaSeatBooking/internal/adapters/redis"
	"CinemaSeatBooking/internal/booking"
	"CinemaSeatBooking/internal/utils"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /movies", listMovies)

	mux.Handle("GET /", http.FileServer(http.Dir("static")))

	store := booking.NewRedisStore(redis.NewClient("localhost:6379"))

	svc := booking.NewService(store)

	bookingHandler := booking.NewHandler(svc)

	mux.HandleFunc("GET /movies/{MovieID}/seats", bookingHandler.ListSeats)
	mux.HandleFunc("POST /movies/{MovieID}/seats/{seatID}/hold", bookingHandler.HoldSeat)

	mux.HandleFunc("PUT /sessions/{sessionID}/confirm", bookingHandler.ConfirmSession)
	mux.HandleFunc("DELETE /sessions{sessionID}", bookingHandler.ReleaseSession)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}

var movies = []movieResponse{
	{ID: "fightClub", Title: "Fight Club", Rows: 6, SeatsPerRow: 8},
	{ID: "dune", Title: "Dune: Part One", Rows: 8, SeatsPerRow: 10},
	{ID: "dune2", Title: "Dune: Part Two", Rows: 10, SeatsPerRow: 10},
}

func listMovies(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, movies)
}

type movieResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Rows        int    `json:"rows"`
	SeatsPerRow int    `json:"seats_per_row"`
}
