package main

import (
	"fmt"

	"github.com/aoc/2023/2/internal/constants"
	"github.com/aoc/2023/2/internal/gamemgr"
	"github.com/nelsonmarro/gopher-toolbox/filemanager"
)

func main() {
	fm := filemanager.New(constants.FinalInputPath)

	lines, err := fm.ReadLines()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	gm := gamemgr.New(lines)
	gameRecords := gm.ExtractGamesInfo()
	gameIdsSum := gm.SumValidGameIds(gameRecords)
	powersSum := gm.GetPowersSumOfMinCubeSets(gameRecords)

	fmt.Println(gameIdsSum)
	fmt.Println(powersSum)
}
