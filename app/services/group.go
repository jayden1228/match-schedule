package group

import (
	"errors"
)

const (
	ErrPlayerNumNotMatchFieldAndRound = "ERR_PLAYER_NUM_NOT_MATCH_FIELD_AND_ROUND"
)

const (
	startDeep = 1
)

// GenGroups 分组服务
// @playerNum 参赛人员数
// @roundNum 回合数
// @fieldNum 场地数字
func GenGroups(playerNum int, roundNum int, fieldNum int, maxFieldCapacity int, minFieldCapacity int) ([][]int, error) {
	if err := Validate(playerNum, roundNum, fieldNum); err != nil {
		return nil, err
	}

	var fieldCapacityList []int
	for i := minFieldCapacity; i <= maxFieldCapacity; i++ {
		fieldCapacityList = append(fieldCapacityList, i)
	}

	groups := CombinationSum(fieldCapacityList, playerNum, startDeep, fieldNum)

	return groups, nil
}

// Validate 校验入参
// @playerNum 参赛人员数
// @roundNum 回合数
// @fieldNum 场地数字
func Validate(playerNum int, roundNum int, fieldNum int) error {
	if playerNum <= 0 || roundNum <= 0 || fieldNum <= 0 {
		return errors.New(ErrPlayerNumNotMatchFieldAndRound)
	}
	if playerNum*roundNum%2 != 0 {
		return errors.New(ErrPlayerNumNotMatchFieldAndRound)
	}
	if playerNum < roundNum*2 {
		return errors.New(ErrPlayerNumNotMatchFieldAndRound)
	}
	return nil
}

// CombinationSum 从数组中选中N个数求和，满足 candidates[m] + ... candidates[n] = target
// @candidates 元素数组
// @target     求和
// @curDep     当前搜索深度
// @tarDep     目标搜索深度
func CombinationSum(candidates []int, target int, curDeep int, tarDeep int) [][]int {
	comb := make([][]int, 0)
	for i := range candidates {
		if candidates[i] == target {
			if curDeep == tarDeep {
				comb = append(comb, []int{candidates[i]})
			}
		} else if candidates[i] < target {
			rtn := CombinationSum(candidates[i:], target-candidates[i], curDeep+1, tarDeep)
			for j := range rtn {
				if len(rtn[j]) == 0 {
					continue
				}
				comb = append(comb, append([]int{candidates[i]}, rtn[j]...))
			}
		}
	}
	return comb
}
