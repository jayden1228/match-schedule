package core

import (
	"errors"
	"match-schedule/pkg/constant"
	"math"
)

const (
	startDeep = 1
	amplitude = 2
)

type genFieldsOption struct {
	Amplitude int32
}

type GenFieldsOption interface {
	apply(*genFieldsOption)
}

type funcOption struct {
	f func(*genFieldsOption)
}

func (fdo *funcOption) apply(do *genFieldsOption) {
	fdo.f(do)
}

func newFuncOption(f func(*genFieldsOption)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func WithAmplitude(s int32) GenFieldsOption {
	return newFuncOption(func(o *genFieldsOption) {
		o.Amplitude = s
	})
}

func defaultOptions() genFieldsOption {
	return genFieldsOption{
		Amplitude: amplitude,
	}
}

// GenFields 分场地
// @playerNum 选手人数
// @fieldNum 场地数
// @roundNum 回合数
// @opts 可选参数
// 		Amplitude 表示每个场地人数偏离平均值的幅度, 默认 amplitude = 2, 使用withAmplitude修改默认设置
func GenFields(playerNum int32, fieldNum int32, roundNum int32, mode int32, opts ...GenFieldsOption) ([]int32, error) {
	if playerNum <= 0 || fieldNum <= 0 {
		return nil, errors.New(constant.ErrPlayerNumNotMatchFieldAndRound)
	}

	genFieldsOption := defaultOptions()

	for _, opt := range opts {
		opt.apply(&genFieldsOption)
	}

	maxFieldCapacity := playerNum/fieldNum + genFieldsOption.Amplitude
	minFieldCapacity := playerNum/fieldNum - genFieldsOption.Amplitude

	var fieldCapacityList []int32
	for i := minFieldCapacity; i <= maxFieldCapacity; i++ {
		fieldCapacityList = append(fieldCapacityList, i)
	}

	fields := CombinationSum(fieldCapacityList, playerNum, startDeep, fieldNum)

	fields = filterField(fields, roundNum, mode)

	field := OptimalFieldChoice(fields, playerNum, fieldNum)

	return field, nil
}

// filterField 过滤不符合规则的分场地结果
// @fields 场地数
// @roundNum 回合
// @mode 单人/双人
func filterField(fields [][]int32, roundNum int32, mode int32) [][]int32 {
restart:
	odd := mode * 2
	for i, v := range fields {
		if int32(len(v))*roundNum%odd != 0 {
			fields = append(fields[:i], fields[i+1:]...)
			goto restart
		}
	}
	return fields
}

// OptimalFieldChoice 选择最接近平均场次的分配
// 平方差求和最小值为最拟合平均
func OptimalFieldChoice(fields [][]int32, playerNum int32, fieldNum int32) []int32 {
	averageField := playerNum / fieldNum
	minVariance := int32(99999)
	minVarianceIndex := 0
	for i, v := range fields {
		var variance int32
		for _, d := range v {
			variance = variance + int32(math.Pow(float64(d-averageField), 2))
		}
		if variance < minVariance {
			minVariance = variance
			minVarianceIndex = i
		}
	}
	return fields[minVarianceIndex]
}
