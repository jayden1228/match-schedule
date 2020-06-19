package field

import (
	"fmt"
	"testing"
)

func Test_GenFields(t *testing.T) {
	minFieldCapacity := 15
	maxFieldCapacity := 18
	playerNum := 50
	fieldNum := 3
	groups, _ := GenFields(playerNum, fieldNum, maxFieldCapacity, minFieldCapacity)
	fmt.Println(groups)
}
