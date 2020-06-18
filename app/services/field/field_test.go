package field

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

func Test_GenFields(t *testing.T) {
	minFieldCapacity := 15
	maxFieldCapacity := 18
	playerNum := 50
	fieldNum := 3
	groups, _ := GenFields(playerNum, fieldNum, maxFieldCapacity, minFieldCapacity)
	fmt.Println(groups)
}
