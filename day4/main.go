package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := "input.txt"
	items := readInput(filepath)
	countFullyContain(items)
	countOverlap(items)
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

func countFullyContain(items []string) {
	sum := 0

	for _, content := range items {
		elfs := strings.Split(content, ",")

		elf1range := strings.Split(elfs[0], "-")
		elf2range := strings.Split(elfs[1], "-")

		elf1start, _ := strconv.Atoi(elf1range[0])
		elf1end, _ := strconv.Atoi(elf1range[1])

		elf2start, _ := strconv.Atoi(elf2range[0])
		elf2end, _ := strconv.Atoi(elf2range[1])

		if (elf2start >= elf1start && elf2end <= elf1end) || (elf1start >= elf2start && elf1end <= elf2end) {
			sum++
		}
	}

	log.Println(sum)
}

func countOverlap(items []string) {
	sum := 0

	for _, content := range items {
		elfs := strings.Split(content, ",")

		elf1range := strings.Split(elfs[0], "-")
		elf2range := strings.Split(elfs[1], "-")

		elf1start, _ := strconv.Atoi(elf1range[0])
		elf1end, _ := strconv.Atoi(elf1range[1])

		elf2start, _ := strconv.Atoi(elf2range[0])
		elf2end, _ := strconv.Atoi(elf2range[1])

		if (elf2start > elf1start && elf2start < elf1end) || (elf2end > elf1start && elf2end < elf1end) {
			sum++
		} else if (elf1start > elf2start && elf1start < elf2end) || (elf1end > elf2start && elf1end < elf2end) {
			sum++
		} else if elf1start == elf2start || elf1end == elf2end || elf1start == elf2end || elf2start == elf1end {
			sum++
		}
	}

	log.Println(sum)
}
