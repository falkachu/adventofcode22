package main

import (
	"bufio"
	"log"
	"os"
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

func splitItems(items []string) [][]string {
	var splitems [][]string

	for _, i := range items {
		var spli []string
		strstart := i[:len(i)/2]
		strend := i[len(i)/2:]
		spli = append(spli, strstart)
		spli = append(spli, strend)

		splitems = append(splitems, spli)
	}

	return splitems
}
