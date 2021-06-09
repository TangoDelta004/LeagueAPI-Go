package main

import (
	"fmt"

	"github.com/LeagueAPI-Go/riotModels"
)

type SummonerRankPackage struct {
	Rank  string
	Tier  string
	Ratio float64
}

func getSummonerRank(summonerid string, apikey string) SummonerRankPackage {
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/league/v4/entries/by-summoner/%v?api_key=%v", summonerid, apikey)
	var summonerRanks riotModels.SummonerRanks
	err := get(url, &summonerRanks)
	if err != nil {
		panic(err)
	}

	var rank string
	var tier string
	var ratio float64
	for _, queueRank := range summonerRanks {
		if queueRank.QueueType == "RANKED_SOLO_5x5" {
			rank = queueRank.Rank
			tier = queueRank.Tier
			ratio = (float64(queueRank.Wins) / float64((queueRank.Wins + queueRank.Losses)) * 100)

		}
	}
	summonerRankPackage := SummonerRankPackage{Rank: rank, Tier: tier, Ratio: ratio}
	return summonerRankPackage
}
