package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	log.Println("Day day2")

	var filepath = "input.txt"

	strats := readInput(filepath)

	var sumScore int
	for _, s := range strats {
		s.calcScore()
		sumScore += s.Score
	}
	log.Printf("Sum: %v", sumScore)

	sumScore = 0
	for _, s := range strats {
		s.calcScore2()
		sumScore += s.Score
	}
	log.Printf("Sum2: %v", sumScore)
}

type Strategy struct {
	Enemy string
	Self  string
	Score int
}

func readInput(filepath string) []Strategy {
	var strats []Strategy

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var s Strategy
		split := strings.Split(scanner.Text(), " ")
		s.Enemy = split[0]
		s.Self = split[1]
		strats = append(strats, s)
	}

	return strats
}

func (s *Strategy) calcScore() {
	s.Score = 0

	switch s.Self {
	case "X":
		s.Score += 1
		if s.Enemy == "A" {
			s.Score += 3
		} else if s.Enemy == "C" {
			s.Score += 6
		}

	case "Y":
		s.Score += 2
		if s.Enemy == "B" {
			s.Score += 3
		} else if s.Enemy == "A" {
			s.Score += 6
		}

	case "Z":
		s.Score += 3
		if s.Enemy == "C" {
			s.Score += 3
		} else if s.Enemy == "B" {
			s.Score += 6
		}
	}
}

func (s *Strategy) calcScore2() {
	s.Score = 0

	switch s.Self {
	case "X":
		if s.Enemy == "A" {
			s.Score += 3
		} else if s.Enemy == "B" {
			s.Score += 1
		} else if s.Enemy == "C" {
			s.Score += 2
		}

	case "Y":
		s.Score += 3
		if s.Enemy == "A" {
			s.Score += 1
		} else if s.Enemy == "B" {
			s.Score += 2
		} else if s.Enemy == "C" {
			s.Score += 3
		}

	case "Z":
		s.Score += 6
		if s.Enemy == "A" {
			s.Score += 2
		} else if s.Enemy == "B" {
			s.Score += 3
		} else if s.Enemy == "C" {
			s.Score += 1
		}
	}
}
