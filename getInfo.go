package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

func getInfo(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	fmt.Println(params)
	summoners := strings.Split(params["names"], ",")

	results := make(map[string]APIPackage)

	type summonerResults struct {
		name    string
		results APIPackage
	}

	channel := make(chan summonerResults)

	wg := sync.WaitGroup{}

	for _, summoner := range summoners {
		summoner := summoner
		wg.Add(1)
		go func() {
			defer wg.Done()
			summonerinfo := getSummonerStats(summoner)
			channel <- summonerResults{name: summoner, results: summonerinfo}
		}()
	}
	go func() {
		wg.Wait()
		close(channel)
	}()

	for summonerinfo := range channel {
		name := summonerinfo.name
		results[name] = summonerinfo.results
	}

	fmt.Println(results)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
