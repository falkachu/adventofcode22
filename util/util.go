package util

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(filepath string) []string {
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
