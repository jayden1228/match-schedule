package group

import (
	"fmt"
	"testing"
)

func Test_CombinationSum(t *testing.T) {
	target := 50
	candidate := []int{15, 16, 17, 18, 19, 20}
	tarDeep := 3
	result := CombinationSum(candidate, target, 1, tarDeep)
	fmt.Println(result)
}

func Test_GenGroups(t *testing.T) {
	maxFieldCapacity := 20
	minFieldCapacity := 10
	playerNum := 50
	fieldNum := 3
	groups, _ := GenGroups(playerNum, 2, fieldNum, maxFieldCapacity, minFieldCapacity)
	fmt.Println(groups)
}
