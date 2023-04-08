package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type setTime struct {
	CurrentTime string `json:"current_time"`
}

var (
	port     = 8080
	endpoint = fmt.Sprintf("0.0.0.0:%d", port)
)

func main() {
	// Router definition
	mux := http.NewServeMux()
	mux.HandleFunc("/time", showCurrentTime)

	// Starting server
	fmt.Printf("Server listening on %s\n", endpoint)
	log.Fatal(http.ListenAndServe(endpoint, mux))
}

// showCurrentTime returns the current time on the server
func showCurrentTime(w http.ResponseWriter, r *http.Request) {
	currentTime := []setTime{
		{CurrentTime: http.TimeFormat},
	}

	json.NewEncoder(w).Encode(currentTime)
}
