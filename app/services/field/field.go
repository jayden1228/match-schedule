package field

import (
	"errors"
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

	groups := CombinationSum(fieldCapacityList, playerNum, startDeep, fieldNum)

	return groups, nil
}

// CombinationSum 从数组中选中N个数求和，满足 candidates[m] + ... candidates[n] = target
// @candidates 元素数组
// @target     求和
// @curDep     当前搜索深度
// @tarDep     目标搜索深度
func CombinationSum(candidates []int, target int, curDeep int, tarDeep int) [][]int {
	comb := make([][]int, 0)
	if curDeep <= tarDeep {
		for i := range candidates {
			if candidates[i] == target {
				if curDeep == tarDeep {
					comb = append(comb, []int{candidates[i]})
					break
				}
			} else if candidates[i] < target {
				rtn := CombinationSum(candidates[i:], target-candidates[i], curDeep+1, tarDeep)
				for j := range rtn {
					if len(rtn[j]) == 0 {
						continue
					}
					comb = append(comb, append([]int{candidates[i]}, rtn[j]...))
				}
			} else {
				break
			}
		}
	}

	return comb
}
