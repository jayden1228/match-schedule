package core

import (
	"errors"
	"log"
	"match-schedule/pkg/constant"
	"math/rand"
	"time"
)

type DoubleCompetition struct {
	PlayerNum int32
	RoundNum  int32
}

// PlayerCompilation 单场地多轮对战
// @playerNum 选手人数
// @roundNum 回合数
func (s *DoubleCompetition) PlayerCompilation() ([][]int32, error) {
RETRY:
	// 选手池
	playerPool := make([]int32, s.PlayerNum)
	for i := int32(1); i <= s.PlayerNum; i++ {
		playerPool[i-1] = i
	}
	// 队友历史池
	var teammatePool [][2]int32
	// 对战历史池
	var competitorPool [][2]int32
	// 孤儿选手
	var alonePlayerPool []int32
	// 轮数存储
	rounds := make([][]int32, s.RoundNum)
	// 人数规则
	if s.PlayerNum*s.RoundNum%4 != 0 {
		return nil, errors.New(constant.ErrPlayerNumNotMatchRound)
	}
	// 遍历生成多轮赛事
	for i := int32(1); i <= s.RoundNum; i++ {
		var round []int32
		c, p := s.playerRoundCompilation(playerPool, teammatePool, competitorPool)
		if len(p) != 0 {
			alonePlayerPool = append(alonePlayerPool, p...)
		}
		round = append(round, c...)
		rounds[i-1] = round
	}
	// 孤儿数组是否有重复孤儿重新生成
	if s.isRepeatAlonePlayer(alonePlayerPool) {
		log.Println("try again...")
		goto RETRY
	}
	// 判断孤儿数组是否满足匹配规则
	alonePlayerPoolLength := len(alonePlayerPool)
	if alonePlayerPoolLength%2 != 0 {
		return nil, errors.New(constant.ErrAlonePlayerNotMatchRule)
	}

	if alonePlayerPoolLength > 0 {
		step := len(rounds) / (alonePlayerPoolLength / 4)
		// 孤儿选手添加到已有的队列
		for i := 0; i <= alonePlayerPoolLength/4; i += step {
			var c []int32
			c, alonePlayerPool = s.pairPlayer(alonePlayerPool, teammatePool, competitorPool)
			rounds[i] = append(rounds[i], c...)
		}
	}

	return rounds, nil
}

// playerRoundCompilation 单场单轮对战
// @playerPool     选手池
// @teammatePool   队友池, 已分配过的队友，添加到该池子
// @competitorPool 对手池, 已分配过的对手，添加到该池子
func (s *DoubleCompetition) playerRoundCompilation(playerPool []int32, teammatePool [][2]int32, competitorPool [][2]int32) ([]int32, []int32) {
	var result []int32
	var alonePlayers []int32
	pool := make([]int32, len(playerPool))
	copy(pool, playerPool)
	count := len(pool) / 4
	for count > 0 {
		var c []int32
		c, pool = s.pairPlayer(pool, teammatePool, competitorPool)
		result = append(result, c...)
		count--
	}

	for i := 0; i < len(pool); i++ {
		alonePlayers = append(alonePlayers, pool[i])
	}

	return result, alonePlayers
}

// pairPlayer 单场地单轮一组选手
// @playerPool     选手池
// @teammatePool   队友池, 已分配过的队友，添加到该池子
// @competitorPool 对手池, 已分配过的对手，添加到该池子
func (s *DoubleCompetition) pairPlayer(pool []int32, teammatePool [][2]int32, competitorPool [][2]int32) ([]int32, []int32) {
	var result []int32
	for len(result) == 0 {
		var a, b, c, d int32
		a, pool = s.getRandomPlayer(pool)
		b, pool = s.getRandomPlayer(pool)
		c, pool = s.getRandomPlayer(pool)
		d, pool = s.getRandomPlayer(pool)

		if !s.existPairs([4]int32{a, b, c, d}, teammatePool, competitorPool) {
			result = append(result, a, b, c, d)
			teammatePool = append(teammatePool, [2]int32{a, b})
			competitorPool = append(competitorPool, [2]int32{a, c}, [2]int32{a, d}, [2]int32{b, c}, [2]int32{b, d})
		}
	}
	return result, pool
}

// isRepeatAlonePlayer 是否有重复孤儿选手
// @alonePlayerPool 孤儿选手池
func (s *DoubleCompetition) isRepeatAlonePlayer(alonePlayerPool []int32) bool {
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
func (s *DoubleCompetition) getRandomPlayer(playerPool []int32) (int32, []int32) {
	rand.Seed(time.Now().UTC().UnixNano())
	length := len(playerPool)
	index := rand.Intn(length)
	result := playerPool[index]
	playerPool = append(playerPool[:index], playerPool[index+1:]...)
	return result, playerPool
}

// existPairs 判断是否存在相同的匹配队友
// @pair 一组选手
// @teammatePool   队友池
// @competitorPool 对手池
func (s *DoubleCompetition) existPairs(pair [4]int32, teammatePool [][2]int32, competitorPool [][2]int32) bool {
	for _, v := range teammatePool {
		if (v[0] == pair[0] && v[1] == pair[1]) || (v[0] == pair[1] && v[1] == pair[0]) {
			return true
		}
		if (v[0] == pair[2] && v[1] == pair[3]) || (v[0] == pair[3] && v[1] == pair[2]) {
			return true
		}
	}
	// TODO 待业务确认是否可以对战部分成员重复
	//for _, v := range competitorPool {
	//	if (v[0] == pair[0] && v[1] == pair[2]) || (v[0] == pair[2] && v[1] == pair[0]) {
	//		return true
	//	}
	//	if (v[0] == pair[1] && v[1] == pair[3]) || (v[0] == pair[3] && v[1] == pair[1]) {
	//		return true
	//	}
	//}
	return false
}
