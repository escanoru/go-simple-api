package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type setTime struct {
	CurrentTime string `json:"current_time"`
	Joke        string `json:"current_joke"`
}

// Joke struct that defines the fields that will hold the json return data
type Joke struct {
	IconUrl string
	Id      string
	URL     string
	Value   string
}

var (
	port     = 8081
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
	holdData := &[]setTime{
		{CurrentTime: http.TimeFormat},
		{Joke: callJoker()},
	}

	json.NewEncoder(w).Encode(*holdData)
}

// callJoker calls jokeFetcher()
func callJoker() string {
	jokeText, err := jokeFetcher()
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("%+v\n\n", jokeText)
	return jokeText
}

// jokeFetcher fetches a random joke and returns it along with an error.
func jokeFetcher() (string, error) {
	log.Println("Fetching Joke")
	jokeInstance := &Joke{}
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return "", err
	}

	err = json.NewDecoder(resp.Body).Decode(jokeInstance)
	return jokeInstance.Value, err
}
