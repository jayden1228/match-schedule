package group

import (
	"match-schedule/pkg/constant"
	"testing"
)

func TestGenSinglePlayerGroups(t *testing.T) {
	GenSinglePlayerGroups(4, 1, constant.SinglePlayer)
}
