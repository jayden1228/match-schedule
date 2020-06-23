package core

import (
	"fmt"
	"testing"
)

func TestPlayerCompilation(t *testing.T) {
	s := SingleCompetition{
		PlayerNum: 13,
		RoundNum:  4,
	}
	r, _ := s.PlayerCompilation()
	fmt.Println(r)
}
