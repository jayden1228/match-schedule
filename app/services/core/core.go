package core

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

const (
	ErrPlayerNumNotMatchRound = "ERR_PLAYER_NUM_NO_MATCH_ROUND"
)

// SinglePlayerCompilation
// @playerNum 选手人数
// @roundNum 回合数
// 队伍编号从1开始, 如果选手数为奇数，添加队伍0
func SinglePlayerCompilation(playerNum int, roundNum int) ([][]int, error) {
RETRY:
	// 选手池
	playerPool := make([]int, playerNum)
	for i := 1; i <= playerNum; i++ {
		playerPool[i-1] = i
	}
	// 对战历史池
	var historyPlayerPairsPool [][2]int
	// 孤儿选手
	var alonePlayerPool []int
	// 轮数存储
	rounds := make([][]int, roundNum)
	// 人数规则
	if playerNum*roundNum%2 != 0 {
		return nil, errors.New("ErrPlayerNumNotMatchRound")
	}
	// 遍历生成多轮赛事
	for i := 1; i <= roundNum; i++ {
		var round []int
		c, p := PickCompetitor(playerPool, historyPlayerPairsPool)
		if p != 0 {
			alonePlayerPool = append(alonePlayerPool, p)
		}
		round = append(round, c...)
		rounds[i-1] = round
	}

	// 孤儿数组是否有重复孤儿重新生成
	if IsRepeatAlonePlayer(alonePlayerPool) {
		log.Println("repeat alone player, try again")
		goto RETRY
	}

	for i := 1; i < len(rounds); i += 2 {
		c, _ := PickCompetitor(alonePlayerPool, historyPlayerPairsPool)
		rounds[i] = append(rounds[i], c...)
	}

	return rounds, nil
}

func PickCompetitor(playerPool []int, historyPool [][2]int) ([]int, int) {
	var result []int
	var alonePlayer int
	pool := make([]int, len(playerPool))
	copy(pool, playerPool)
	count := len(pool) / 2
	for count > 0 {
		var a, b int
		a, pool = GetRandomPlayer(pool)
		b, pool = GetRandomPlayer(pool)
		if !ExistPairs([2]int{a, b}, historyPool) {
			result = append(result, a, b)
			historyPool = append(historyPool, [2]int{a, b})
			count--
		}
	}
	if len(pool) != 0 {
		alonePlayer = pool[0]
	}
	return result, alonePlayer
}

func IsRepeatAlonePlayer(input []int) bool {
	result := false
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				result = true
				return result
			}
		}
	}
	return result
}

func GetRandomPlayer(pool []int) (int, []int) {
	rand.Seed(time.Now().UTC().UnixNano())
	length := len(pool)
	index := rand.Intn(length)
	result := pool[index]
	pool = append(pool[:index], pool[index+1:]...)
	return result, pool
}

func ExistPairs(pair [2]int, historyPool [][2]int) bool {
	for _, v := range historyPool {
		if (v[0] == pair[0] && v[1] == pair[1]) || (v[0] == pair[1] && v[1] == pair[0]) {
			return true
		}
	}
	return false
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
