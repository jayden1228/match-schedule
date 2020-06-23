package core

import (
	"fmt"
	"testing"
)

func TestDoubleCompetition_PlayerCompilation(t *testing.T) {
	s := DoubleCompetition{
		PlayerNum: 1000,
		RoundNum:  4,
	}

	r, _ := s.PlayerCompilation()
	for _, v := range r {
		fmt.Println("len:", len(v), ",", v)
	}
}
