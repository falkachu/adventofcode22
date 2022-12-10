package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

	log.Println("Day day1")

	var filepath = "input.txt"

	allElfsInv := getElfInventories(filepath)
	maxelf, maxcalsum := calcMaxInventory(allElfsInv)

	log.Printf("Elf with largest Inventory: %v / Inventory size: %v", maxelf+1, maxcalsum)

	calcMaxThree(allElfsInv)
}

func exOne() {
	var allelfsinv [][]int64
	var elfinv []int64
	var tmpinvsum int64 = 0
	var maxinvsum int64 = 0
	var curelf = 1
	var maxelf = 1

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() != "" {
			// Parse String to Int64
			parseInt, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			// Append Int to Current Elf Inventory
			elfinv = append(elfinv, parseInt)

			// Add Int to Current Elf Inventory Sum
			tmpinvsum += parseInt
		} else {
			log.Printf("Elf: %v / ElfInventory: %v", curelf, elfinv)

			if tmpinvsum > maxinvsum {
				maxinvsum = tmpinvsum
				maxelf = curelf
			}

			allelfsinv = append(allelfsinv, elfinv)
			elfinv = nil
			curelf++
			tmpinvsum = 0
		}
	}

	log.Printf("Max Elf = %v / Inventory Size = %v", maxelf, maxinvsum)
}

func getElfInventories(filepath string) [][]int64 {
	var elfinv []int64
	var allElfsInv [][]int64
	var curelf = 0

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() != "" {
			// Parse String to Int64
			parseInt, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			// Append Int to Current Elf Inventory
			elfinv = append(elfinv, parseInt)
		} else {
			allElfsInv = append(allElfsInv, elfinv)
			elfinv = nil
			curelf++
		}
	}

	return allElfsInv
}

func calcMaxInventory(elfinv [][]int64) (int, int64) {
	var tmpcalsum int64 = 0
	var maxcalsum int64 = 0
	var maxelf int

	for elfnum, elf := range elfinv {
		for _, cal := range elf {
			tmpcalsum += cal
		}

		if tmpcalsum > maxcalsum {
			maxcalsum = tmpcalsum
			maxelf = elfnum
		}

		tmpcalsum = 0
	}

	return maxelf, maxcalsum
}

func calcMaxThree(elfinv [][]int64) {

	sum := 0

	for i := 0; i <= 2; i++ {
		en, cal := calcMaxInventory(elfinv)
		log.Printf("Elf: %v | InvSize: %v", en+1, cal)
		sum += int(cal)
		elfinv[en] = elfinv[len(elfinv)-1]
		elfinv[len(elfinv)-1] = nil
		elfinv = elfinv[:len(elfinv)-1]
	}

	log.Printf("ThreeInvSum: %v", sum)
}
