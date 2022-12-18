package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var supplystacks [][]string

	line1 := []string{"Z", "J", "N", "W", "P", "S"}
	line2 := []string{"G", "S", "T"}
	line3 := []string{"V", "Q", "R", "L", "H"}
	line4 := []string{"V", "S", "T", "D"}
	line5 := []string{"Q", "Z", "T", "D", "B", "M", "J"}
	line6 := []string{"M", "W", "T", "J", "D", "C", "Z", "L"}
	line7 := []string{"L", "P", "M", "W", "G", "T", "J"}
	line8 := []string{"N", "G", "M", "T", "B", "F", "Q", "H"}
	line9 := []string{"R", "D", "G", "C", "P", "B", "Q", "W"}

	supplystacks = append(supplystacks, line1, line2, line3, line4, line5, line6, line7, line8, line9)

	log.Println(supplystacks)

	lines := readInput("input.txt")
	instructions := parseLines(lines)

	log.Println(instructions)

	execInstructions(supplystacks, instructions)
}

func readInput(filepath string) []string {
	var items []string

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	return items
}

func parseLines(lines []string) []Instruction {
	var instructions []Instruction

	for _, line := range lines {
		insts := strings.Split(line, " ")
		move, _ := strconv.Atoi(insts[1])
		from, _ := strconv.Atoi(insts[3])
		to, _ := strconv.Atoi(insts[5])
		inst := NewInstruction(move, from, to)
		instructions = append(instructions, inst)
	}

	return instructions
}

type Instruction struct {
	Move int
	From int
	To   int
}

func NewInstruction(move int, from int, to int) Instruction {
	return Instruction{Move: move, From: from, To: to}
}

func execInstructions(supplystacks [][]string, instructions []Instruction) {
	for _, inst := range instructions {
		inst.From -= 1
		inst.To -= 1
		for m := 1; m <= inst.Move; m++ {
			supplystacks[inst.To] = append(supplystacks[inst.To], supplystacks[inst.From][len(supplystacks[inst.From])-1])
			supplystacks[inst.From] = supplystacks[inst.From][:len(supplystacks[inst.From])-1]
		}
	}

	log.Println(supplystacks)
}

func execInstructions2(supplystacks [][]string, instructions []Instruction) {
	for _, inst := range instructions {
		inst.From -= 1
		inst.To -= 1
		mstack := supplystacks[inst.From][(len(supplystacks[inst.From]) - 1 - inst.Move):(len(supplystacks[inst.From]) - 1)]
		supplystacks[inst.To] = append(supplystacks[inst.To], mstack)
		supplystacks[inst.From] = supplystacks[inst.From][:len(supplystacks[inst.From])-1]
	}

	log.Println(supplystacks)
}
