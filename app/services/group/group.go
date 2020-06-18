package group

import (
	"match-schedule/pkg/constant"
)

// GenGroups
// @playerNum
// @roundNum
// @gameType
func GenGroups(playerNum int, roundNum int, gameType int) [][]int {
	switch gameType {
	case constant.DoublePlayer:
		return GenDoublePlayerGroups(playerNum, roundNum)
	case constant.SinglePlayer:
		return GenSinglePlayerGroups(playerNum, roundNum)
	default:
		return nil
	}
}

// GenSinglePlayerGroups
func GenSinglePlayerGroups(playerNum int, roundNum int) [][]int {
	if playerNum*roundNum%2 != 0 || playerNum < roundNum*2 {
		return nil
	}

	return nil
}

// GenSinglePlayerGroups
func GenDoublePlayerGroups(playerNum int, roundNum int) [][]int {
	if playerNum*roundNum%4 != 0 || playerNum < roundNum*4 {
		return nil
	}
	return nil
}
