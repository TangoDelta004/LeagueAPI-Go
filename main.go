package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type playerInfo struct {
	summonername      string
	bestchampmastery  int
	bestchamp         string
	totalchampmastery int
	rank              string
	tier              string
	ratio             int
	kills             int
	deaths            int
	assists           int
	cs                int
	vs                int
	totaldamage       int
	objdamage         int
	turretdamage      int
	turretkills       int
	inhibkills        int
	killingspree      int
	multikill         int
	allyjungle        int
	enemyjungle       int
	visionwards       int
	wardskilled       int
	dragonkills       int
	baronkills        int
	riftkills         int
	creepspermin10    int
	creepspermin20    int
	csdiff10          int
	csdiff20          int
	goldpermin10      int
	goldpermin20      int
	xppermin10        int
	xppermin20        int
	xpdiffpermin10    int
	xpdiffpermin20    int
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/getInfo", getInfo).Methods("GET").Queries("names", "{names}")
	log.Fatal(http.ListenAndServe(":8000", r))
}
