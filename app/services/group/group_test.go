package group

import (
	"fmt"
	"match-schedule/pkg/constant"
	"testing"
)

func TestGenSinglePlayerGroups(t *testing.T) {
}

func TestGenSinglePlayerGroups1(t *testing.T) {
	r, err := GenGroups(20, 2, 2, constant.SingleMode)
	fmt.Println(err, r)
}
