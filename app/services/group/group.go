package group

import (
	"errors"
	"match-schedule/app/services/core"
	"match-schedule/pkg/constant"
)

// GenGroups
// @playerNum
// @roundNum
// @fieldNum
// @mode 单人/双人
func GenGroups(playerNum int, fieldNum int, roundNum int, mode int) ([][][]int, error) {
	if err := validateInput(playerNum, roundNum, mode); err != nil {
		return nil, err
	}
	// 分场地
	fields, err := core.GenFields(playerNum, fieldNum, roundNum, core.WithAmplitude(3), core.WithMode(mode))
	if err != nil {
		return nil, err
	}
	// 生成各场地分组
	var groups [][][]int
	for _, v := range fields {
		s := core.SingleCompetition{
			PlayerNum: v,
			RoundNum:  roundNum,
		}
		group, err := s.PlayerCompilation()
		if err != nil {
			return nil, nil
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func validateInput(playerNum int, roundNum int, mode int) error {
	switch mode {
	case constant.SingleMode:
		if playerNum*roundNum%2 != 0 || playerNum < roundNum*2 {
			return errors.New(constant.ErrPlayerNumNotMatchFieldAndRound)
		}
		return nil
	case constant.DoubleMode:
		if playerNum*roundNum%4 != 0 || playerNum < roundNum*4 {
			return errors.New(constant.ErrPlayerNumNotMatchFieldAndRound)
		}
		return nil
	}
	return errors.New(constant.ErrModeWrong)
}
