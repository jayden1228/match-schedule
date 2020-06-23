package group

import (
	"fmt"
	"match-schedule/app/services/core"
	"match-schedule/pkg/constant"
	"testing"
)

func TestGenSinglePlayerGroups(t *testing.T) {
	r, err := GenGroups(20, 2, 2, constant.SingleMode, core.WithAmplitude(2))
	fmt.Println(err, r)
}
