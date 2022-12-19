package day6

import (
	"adventofcode/util"
	"log"
	"testing"
)

func TestExecDay6(t *testing.T) {
	filepath := "input.txt"
	input := util.ReadLines(filepath)
	aftern := markerAfterN(input[0])
	log.Println(aftern)

	aftern2 := markerAfterN2(input[0], 14)
	log.Println(aftern2)
}
