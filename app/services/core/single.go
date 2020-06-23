package core

import (
	"errors"
	"log"
	"match-schedule/pkg/constant"
	"math/rand"
	"time"
)

type SingleCompetition struct {
	PlayerNum int
	RoundNum  int
}

// PlayerCompilation
// @playerNum 选手人数
// @roundNum 回合数
func (s *SingleCompetition) PlayerCompilation() ([][]int, error) {
RETRY:
	// 选手池
	playerPool := make([]int, s.PlayerNum)
	for i := 1; i <= s.PlayerNum; i++ {
		playerPool[i-1] = i
	}
	// 对战历史池
	var historyPlayerPairsPool [][2]int
	// 孤儿选手
	var alonePlayerPool []int
	// 轮数存储
	rounds := make([][]int, s.RoundNum)
	// 人数规则
	if s.PlayerNum*s.RoundNum%2 != 0 {
		return nil, errors.New(constant.ErrPlayerNumNotMatchRound)
	}
	// 遍历生成多轮赛事
	for i := 1; i <= s.RoundNum; i++ {
		var round []int
		c, p := s.playerRoundCompilation(playerPool, historyPlayerPairsPool)
		if p != 0 {
			alonePlayerPool = append(alonePlayerPool, p)
		}
		round = append(round, c...)
		rounds[i-1] = round
	}
	// 孤儿数组是否有重复孤儿重新生成
	if s.isRepeatAlonePlayer(alonePlayerPool) {
		log.Println("repeat alone player, try again")
		goto RETRY
	}
	// 孤儿选手添加到已有的队列
	if len(alonePlayerPool) > 0 {
		for i := 1; i < len(rounds); i += 2 {
			c, _ := s.pairPlayer(alonePlayerPool, historyPlayerPairsPool)
			rounds[i] = append(rounds[i], c...)
		}
	}

	return rounds, nil
}

// playerRoundCompilation 单场单轮对战
// @playerPool     选手池
// @competitorPool 对手池, 已分配过的对手，添加到该池子
func (s *SingleCompetition) playerRoundCompilation(playerPool []int, competitorPool [][2]int) ([]int, int) {
	var result []int
	var alonePlayer int
	pool := make([]int, len(playerPool))
	copy(pool, playerPool)
	count := len(pool) / 2
	for count > 0 {
		var c []int
		c, pool = s.pairPlayer(pool, competitorPool)
		result = append(result, c...)
		count--
	}
	if len(pool) != 0 {
		alonePlayer = pool[0]
	}
	return result, alonePlayer
}

// pairPlayer 单场地单轮一组选手
// @playerPool     选手池
// @competitorPool 对手池, 已分配过的对手，添加到该池子
func (s *SingleCompetition) pairPlayer(pool []int, competitorPool [][2]int) ([]int, []int) {
	var result []int
	for len(result) == 0 {
		var a, b int
		a, pool = s.getRandomPlayer(pool)
		b, pool = s.getRandomPlayer(pool)
		if !s.existPairs([2]int{a, b}, competitorPool) {
			result = append(result, a, b)
			competitorPool = append(competitorPool, [2]int{a, b})
		}
	}
	return result, pool
}

// isRepeatAlonePlayer 是否有重复孤儿选手
// @alonePlayerPool 孤儿选手池
func (s *SingleCompetition) isRepeatAlonePlayer(alonePlayerPool []int) bool {
	result := false
	for i := 0; i < len(alonePlayerPool); i++ {
		for j := i + 1; j < len(alonePlayerPool); j++ {
			if alonePlayerPool[i] == alonePlayerPool[j] {
				result = true
				return result
			}
		}
	}
	return result
}

// getRandomPlayer 随机获取一个选手
// @playerPool 选手池
func (s *SingleCompetition) getRandomPlayer(playerPool []int) (int, []int) {
	rand.Seed(time.Now().UTC().UnixNano())
	length := len(playerPool)
	index := rand.Intn(length)
	result := playerPool[index]
	playerPool = append(playerPool[:index], playerPool[index+1:]...)
	return result, playerPool
}

// existPairs 判断是否存在相同的匹配队友
// @pair 一组选手
// @competitorPool 对手池
func (s *SingleCompetition) existPairs(pair [2]int, competitorPool [][2]int) bool {
	for _, v := range competitorPool {
		if (v[0] == pair[0] && v[1] == pair[1]) || (v[0] == pair[1] && v[1] == pair[0]) {
			return true
		}
	}
	return false
}
