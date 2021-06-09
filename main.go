package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var apikey string

func main() {

	key, ok := os.LookupEnv("RIOT_API_KEY")
	if !ok {
		panic("missing environment variable for API key")
	}
	apikey = key

	r := mux.NewRouter()
	r.HandleFunc("/api/getInfo", getInfo).Methods("GET").Queries("names", "{names}")
	log.Fatal(http.ListenAndServe(":8000", r))
}
