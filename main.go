package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type Response struct {
	Number int `json:"number"`
}

func main() {
	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Number: rand.Intn(100), // generates a random number between 0 and 99
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)
}
