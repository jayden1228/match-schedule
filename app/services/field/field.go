package field

import (
	"errors"
	"match-schedule/app/services/core"
)

const (
	ErrPlayerNumNotMatchFieldAndRound = "ERR_PLAYER_NUM_NOT_MATCH_FIELD_AND_ROUND"
)

const (
	startDeep = 1
)

// GenFields 分场地
// @playerNum 参赛人员数
// @fieldNum 场地数字
// @maxFieldCapacity 最大场馆容量
// @minFieldCapacity 最小场馆容量
func GenFields(playerNum int, fieldNum int, maxFieldCapacity int, minFieldCapacity int) ([][]int, error) {
	if playerNum <= 0 || fieldNum <= 0 {
		return nil, errors.New(ErrPlayerNumNotMatchFieldAndRound)
	}

	var fieldCapacityList []int
	for i := minFieldCapacity; i <= maxFieldCapacity; i++ {
		fieldCapacityList = append(fieldCapacityList, i)
	}

	groups := core.CombinationSum(fieldCapacityList, playerNum, startDeep, fieldNum)

	return groups, nil
}
