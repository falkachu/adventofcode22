package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	filepath := "input.txt"
	input := readInput(filepath)
	log.Println(input)
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
