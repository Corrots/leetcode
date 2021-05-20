package main

import (
	"encoding/json"
	"net/http"
)

type Game struct {
	Platform    string `json:"platform"`
	Price       string `json:"price"`
	ReleaseDate string `json:"release_date"`
}

// Test HTTP server
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			game := &Game{
				Platform:    "PS5",
				Price:       "Daemons' Soul",
				ReleaseDate: "2020-11-10",
			}
			b, err := json.Marshal(game)
			if err != nil {
				panic(err)
			}
			w.Write(b)
			w.WriteHeader(http.StatusOK)
		}
	})

	http.ListenAndServe(":8080", nil)
}
