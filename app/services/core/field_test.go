package core

import (
	"fmt"
	"testing"
)

func TestOptimalFieldChoice(t *testing.T) {
	input := [][]int{{2, 2, 2}, {1, 3, 2}, {1, 1, 4}, {1, 5, 0}}
	fmt.Println(OptimalFieldChoice(input, 6, 3))
}
