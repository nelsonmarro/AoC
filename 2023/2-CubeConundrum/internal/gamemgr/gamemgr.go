package gamemgr

import (
	"strings"

	"github.com/aoc/2023/2/internal/txthelpers"
)

type cubeInfo struct {
	Color string
	Count int
}

type GameManager struct {
	txtGameRecords []string
}

func New(txtGameRecords []string) GameManager {
	return GameManager{
		txtGameRecords: txtGameRecords,
	}
}

func (gm GameManager) ExtractGamesInfo() map[int][]cubeInfo {
	gameRecords := make(map[int][]cubeInfo)
	gameIds := 1
	for _, gameTxtRecord := range gm.txtGameRecords {
		gameInfo := strings.Split(gameTxtRecord, ":")[1]
		gameSubsets := strings.Split(gameInfo, ";")
		gameRecord := make([]cubeInfo, 0, len(gameSubsets))

		for _, set := range gameSubsets {
			counts := txthelpers.ExtractNumbers(set)
			colors := txthelpers.ExtractColors(set)
			for i := 0; i < len(counts); i++ {
				gameRecord = append(gameRecord, cubeInfo{
					Color: colors[i],
					Count: counts[i],
				})
			}
		}
		gameRecords[gameIds] = gameRecord
		gameIds++
	}

	return gameRecords
}

func (gm GameManager) SumValidGameIds(gameRecords map[int][]cubeInfo) int {
	var gameIdsSum int
	for gameId, records := range gameRecords {
		isValidGame := true
		for _, subset := range records {
			if (subset.Count > 14 && subset.Color == "blue") ||
				(subset.Count > 12 && subset.Color == "red") ||
				(subset.Count > 13 && subset.Color == "green") {
				isValidGame = false
				break
			}
		}
		if isValidGame {
			gameIdsSum += gameId
		}
	}
	return gameIdsSum
}

func (gm GameManager) GetPowersSumOfMinCubeSets(gameRecords map[int][]cubeInfo) int {
	var powersSum int

	for _, gameRecord := range gameRecords {
		var maxRed int
		var maxBlue int
		var maxGreen int

		for _, subset := range gameRecord {
			if subset.Count > maxRed && subset.Color == "red" {
				maxRed = subset.Count
			}
			if subset.Count > maxBlue && subset.Color == "blue" {
				maxBlue = subset.Count
			}
			if subset.Count > maxGreen && subset.Color == "green" {
				maxGreen = subset.Count
			}
		}

		powersSum += (maxGreen * maxRed * maxBlue)
	}

	return powersSum
}
