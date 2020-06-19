package core

// SinglePlayerCompilation
// @playerNum 选手人数
// @roundNum 回合数
// 队伍编号从1开始, 如果选手数为奇数，添加队伍0
func SinglePlayerCompilation(playerNum int, roundNum int) [][]int {
	// 本轮人数是否为奇数
	//var odd bool
	rounds := make([][]int, roundNum)
	// 初始化队员池子
	//playerPoolOne, playerPoolTwo := InitPlayerPool(playerNum)
	//for i := 1; i <= roundNum; i++ {
	//	odd := playerNum % 2
	//	round := make([]int, 0)
	//	rounds = append(rounds, round)
	//}
	return rounds
}

func InitPlayerPool(playerNum int) ([]int, []int) {
	// player 分两部分
	var playerPartOne []int
	for i := 1; i <= playerNum/2; i++ {
		playerPartOne = append(playerPartOne, i)
	}
	var playerPartTwo []int
	for i := playerNum/2 + 1; i <= playerNum; i++ {
		playerPartTwo = append(playerPartTwo, i)
	}
	return playerPartOne, playerPartTwo
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
