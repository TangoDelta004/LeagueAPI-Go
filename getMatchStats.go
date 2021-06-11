package main

import (
	"fmt"
	"math"

	"github.com/LeagueAPI-Go/riotModels"
)

type MatchStatsPackage struct {
	Kills          float64
	Deaths         float64
	Assists        float64
	Cs             float64
	Vs             float64
	Totaldamage    float64
	Objdamage      float64
	Turretdamage   float64
	Turretkills    float64
	Inhibkills     float64
	Killingspree   float64
	Multikill      float64
	Allyjungle     float64
	Enemyjungle    float64
	Visionwards    float64
	Wardskilled    float64
	Dragonkills    float64
	Baronkills     float64
	Riftkills      float64
	creepspermin10 float64
	creepspermin20 float64
	csdiff10       float64
	csdiff20       float64
	goldpermin10   float64
	goldpermin20   float64
	xppermin10     float64
	xppermin20     float64
	xpdiffpermin10 float64
	xpdiffpermin20 float64
	cpm10count     float64
	cpm20count     float64
	csd10count     float64
	csd20count     float64
	xpd10count     float64
	xpd20count     float64
	xpm10count     float64
	xpm20count     float64
	gpm10count     float64
	gpm20count     float64
}

func getMatchStats(accountid string, apikey string, name string) (MatchStatsPackage, ErrorPackage) {

	var matchStatsPackage MatchStatsPackage

	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/match/v4/matchlists/by-account/%v?api_key=%v", accountid, apikey)
	var matches riotModels.Matches
	err := get(url, &matches)
	if err != nil {
		return MatchStatsPackage{}, ErrorPackage{ErrorCode: "300", ErrorMessage: "Could not Find that summoner by id given"}
	}
	//normally we would loop through ALL matches. but my limited API key prevents this
	for i := 0; i < 3; i++ {
		matchId := matches.Matches[i].GameID
		url = fmt.Sprintf("https://na1.api.riotgames.com/lol/match/v4/matches/%v?api_key=%v", matchId, apikey)
		var matchInfo riotModels.MatchInfo
		err = get(url, &matchInfo)
		if err != nil {
			return MatchStatsPackage{}, ErrorPackage{ErrorCode: "400", ErrorMessage: "Internal Error"}
		}
		// get list of all players identities so we can find the player we are looking for and get his ID
		participantIdentities := matchInfo.ParticipantIdentities
		var participantID int
		for _, participantIdentity := range participantIdentities {
			if participantIdentity.Player.SummonerName == name {
				participantID = participantIdentity.ParticipantID
			}
		}

		//Use his ID to find his match Stats
		participants := matchInfo.Participants
		for _, participant := range participants {
			if participant.ParticipantID == participantID {
				//we have found the correct player. now we find his Stats for this individual game
				if participant.TeamID == 100 {
					matchStatsPackage.Dragonkills = float64(matchInfo.Teams[0].DragonKills)
					matchStatsPackage.Baronkills = float64(matchInfo.Teams[0].BaronKills)
					matchStatsPackage.Riftkills = float64(matchInfo.Teams[0].RiftHeraldKills)
				}
				if participant.TeamID == 200 {
					matchStatsPackage.Dragonkills = float64(matchInfo.Teams[1].DragonKills)
					matchStatsPackage.Baronkills = float64(matchInfo.Teams[1].BaronKills)
					matchStatsPackage.Riftkills = float64(matchInfo.Teams[1].RiftHeraldKills)
				}
				matchStatsPackage.Totaldamage += float64(participant.Stats.TotalDamageDealt)
				matchStatsPackage.Cs += float64(participant.Stats.NeutralMinionsKilled) + float64(participant.Stats.NeutralMinionsKilled)
				matchStatsPackage.Vs += float64(participant.Stats.VisionScore)
				matchStatsPackage.Kills += float64(participant.Stats.Kills)
				matchStatsPackage.Deaths += float64(participant.Stats.Deaths)
				matchStatsPackage.Assists += float64(participant.Stats.Assists)
				matchStatsPackage.Objdamage += float64(participant.Stats.DamageDealtToObjectives)
				matchStatsPackage.Turretdamage += float64(participant.Stats.DamageDealtToTurrets)
				matchStatsPackage.Turretkills += float64(participant.Stats.TurretKills)
				matchStatsPackage.Inhibkills += float64(participant.Stats.InhibitorKills)
				matchStatsPackage.Killingspree += float64(participant.Stats.LargestKillingSpree)
				matchStatsPackage.Multikill += float64(participant.Stats.LargestMultiKill)
				matchStatsPackage.Allyjungle += float64(participant.Stats.NeutralMinionsKilledTeamJungle)
				matchStatsPackage.Enemyjungle += float64(participant.Stats.NeutralMinionsKilledEnemyJungle)
				matchStatsPackage.Visionwards += float64(participant.Stats.VisionWardsBoughtInGame)
				matchStatsPackage.Wardskilled += float64(participant.Stats.WardsKilled)
			}

		}
	}

	matchStatsPackage.Cs = math.Round(matchStatsPackage.Cs / 3)
	matchStatsPackage.Vs = math.Round(matchStatsPackage.Vs / 3)
	matchStatsPackage.Kills = math.Round((matchStatsPackage.Kills / 3))
	matchStatsPackage.Deaths = math.Round(matchStatsPackage.Deaths / 3)
	matchStatsPackage.Assists = math.Round(matchStatsPackage.Assists / 3)
	matchStatsPackage.Objdamage = math.Round(matchStatsPackage.Objdamage / 3)
	matchStatsPackage.Turretdamage = math.Round(matchStatsPackage.Turretdamage / 3)
	matchStatsPackage.Turretkills = math.Round(matchStatsPackage.Turretkills / 3)
	matchStatsPackage.Inhibkills = math.Round(matchStatsPackage.Inhibkills / 3)
	matchStatsPackage.Killingspree = math.Round(matchStatsPackage.Killingspree / 3)
	matchStatsPackage.Multikill = math.Round(matchStatsPackage.Multikill / 3)
	matchStatsPackage.Allyjungle = math.Round(matchStatsPackage.Allyjungle / 3)
	matchStatsPackage.Enemyjungle = math.Round(matchStatsPackage.Enemyjungle / 3)
	matchStatsPackage.Visionwards = math.Round(matchStatsPackage.Visionwards / 3)
	matchStatsPackage.Wardskilled = math.Round(matchStatsPackage.Wardskilled / 3)
	// matchStatsPackage.creepspermin10 = math.Round(matchStatsPackage.creepspermin10 / matchStatsPackage.cpm10count)
	// matchStatsPackage.creepspermin20 = math.Round(matchStatsPackage.creepspermin20 / matchStatsPackage.cpm20count)
	// matchStatsPackage.csdiff10 = math.Round(matchStatsPackage.csdiff10 / matchStatsPackage.csd10count)
	// matchStatsPackage.csdiff20 = math.Round(matchStatsPackage.csdiff20 / matchStatsPackage.csd20count)
	// matchStatsPackage.goldpermin10 = math.Round(matchStatsPackage.goldpermin10 / matchStatsPackage.gpm10count)
	// matchStatsPackage.goldpermin20 = math.Round(matchStatsPackage.goldpermin20 / matchStatsPackage.gpm20count)
	// matchStatsPackage.xppermin10 = math.Round(matchStatsPackage.xppermin10 / matchStatsPackage.xpm10count)
	// matchStatsPackage.xppermin20 = math.Round(matchStatsPackage.xppermin20 / matchStatsPackage.xpm20count)
	// matchStatsPackage.xpdiffpermin10 = math.Round(matchStatsPackage.xpdiffpermin10 / matchStatsPackage.xpd10count)
	// matchStatsPackage.xpdiffpermin20 = math.Round(matchStatsPackage.xpdiffpermin20 / matchStatsPackage.xpd20count)
	matchStatsPackage.Dragonkills = math.Round(matchStatsPackage.Dragonkills / 3)
	matchStatsPackage.Baronkills = math.Round(matchStatsPackage.Baronkills / 3)
	matchStatsPackage.Riftkills = math.Round(matchStatsPackage.Riftkills / 3)
	return matchStatsPackage, ErrorPackage{}
}
