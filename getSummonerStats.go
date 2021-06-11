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

type ErrorPackage struct {
	ErrorCode    string
	ErrorMessage string
}

func getSummonerStats(name string) (APIPackage, ErrorPackage) {

	//first get the summoners basic information. we will need his ID to get information on the champions he plays
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%v?api_key=%v", name, apikey)
	var summoner riotModels.Summoner
	err := get(url, &summoner)
	if err != nil {
		return APIPackage{}, ErrorPackage{ErrorCode: "400", ErrorMessage: "Could not Find that summoner: " + name}
	}
	summonerid := summoner.ID
	accountid := summoner.AccountID

	// use the summonerID to retreive information on his best champion and his total mastery.
	masteryPackage, masteryErrorPackage := getMastery(summonerid, apikey)
	if masteryErrorPackage.ErrorCode != "200" {
		return APIPackage{}, ErrorPackage{ErrorCode: masteryErrorPackage.ErrorCode, ErrorMessage: masteryErrorPackage.ErrorMessage}
	}
	fmt.Printf("%v's mastery: %v", name, masteryPackage)
	fmt.Println()

	// use Summoner ID to retrieve their rank, tier, and win ratio
	summonerRankPackage, summonerRankErrorPackage := getSummonerRank(summonerid, apikey)
	if summonerRankErrorPackage.ErrorCode != "200" {
		return APIPackage{}, ErrorPackage{ErrorCode: summonerRankErrorPackage.ErrorCode, ErrorMessage: summonerRankErrorPackage.ErrorMessage}
	}
	fmt.Printf("%v's rank: %v", name, summonerRankPackage)
	fmt.Println()

	//use account ID to retrieve an average of their games stats
	matchStatsPackage, matchStatsErrorPackage := getMatchStats(accountid, apikey, name)
	if matchStatsErrorPackage.ErrorCode != "200" {
		return APIPackage{}, ErrorPackage{ErrorCode: matchStatsErrorPackage.ErrorCode, ErrorMessage: matchStatsErrorPackage.ErrorMessage}
	}
	fmt.Printf("%v's matchstats: %v", name, matchStatsPackage)
	fmt.Println()

	apiPackage := APIPackage{Mastery: masteryPackage, SummonerRank: summonerRankPackage, MatchStats: matchStatsPackage}

	return apiPackage, ErrorPackage{}
}
