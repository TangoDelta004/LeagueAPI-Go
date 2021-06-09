package main

import (
	"fmt"

	"github.com/LeagueAPI-Go/riotModels"
)

type MasteryPackage struct {
	Totalchampmastery int
	Bestchampid       int
	Bestchampmastery  int
	Bestchampion      string
}

func getMastery(summonerid string, apikey string) MasteryPackage {
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-summoner/%s?api_key=%s", summonerid, apikey)
	var championmastery riotModels.ChampionMastery
	err := get(url, &championmastery)
	if err != nil {
		panic(err)
	}
	totalchampmastery := 0
	for _, champion := range championmastery {
		totalchampmastery += champion.ChampionPoints
	}
	bestchampid := championmastery[0].ChampionID
	bestchampmastery := championmastery[0].ChampionPoints

	// use the best champion ID to find the corresponding champion name
	url = "http://ddragon.leagueoflegends.com/cdn/9.18.1/data/en_US/champion.json"
	var ddragon riotModels.Ddragon
	err = get(url, &ddragon)
	if err != nil {
		panic(err)
	}

	var masteryPackage MasteryPackage

	for _, champion := range ddragon.Data {
		if bestchampid == champion.Key {
			masteryPackage.Bestchampion = champion.Name
		}
	}

	masteryPackage = MasteryPackage{Totalchampmastery: totalchampmastery, Bestchampid: bestchampid, Bestchampmastery: bestchampmastery}
	return masteryPackage
}
