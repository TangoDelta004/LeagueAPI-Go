package main

import (
	"fmt"

	"github.com/LeagueAPI-Go/riotModels"
)

type APIPackage struct {
	Mastery      MasteryPackage
	SummonerRank SummonerRankPackage
	MatchStats   MatchStatsPackage
}

func getSummonerStats(name string) APIPackage {

	//first get the summoners basic information. we will need his ID to get information on the champions he plays
	apikey := "***REMOVED***"
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%v?api_key=%v", name, apikey)
	var summoner riotModels.Summoner
	err := get(url, &summoner)
	if err != nil {
		panic(err)
	}
	summonerid := summoner.ID
	accountid := summoner.AccountID

	// use the summonerID to retreive information on his best champion and his total mastery.
	masteryPackage := getMastery(summonerid, apikey)
	fmt.Printf("%v's mastery: %v", name, masteryPackage)
	fmt.Println()

	summonerRankPackage := getSummonerRank(summonerid, apikey)
	fmt.Printf("%v's rank: %v", name, summonerRankPackage)
	fmt.Println()

	matchStatsPackage := getMatchStats(accountid, apikey, name)
	fmt.Printf("%v's matchstats: %v", name, matchStatsPackage)
	fmt.Println()

	apiPackage := APIPackage{Mastery: masteryPackage, SummonerRank: summonerRankPackage, MatchStats: matchStatsPackage}

	return apiPackage
}
