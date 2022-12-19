package day6

import (
	"log"
)

func markerAfterN(input string) int {
	aftern := 4

	for i := 0; i < len(input); i++ {
		if input[i] != input[i+1] && input[i] != input[i+2] && input[i] != input[i+3] && input[i+1] != input[i+2] && input[i+1] != input[i+3] && input[i+2] != input[i+3] {
			log.Println(input[i : i+4])
			break
		} else {
			aftern++
		}
	}

	return aftern
}

func markerAfterN2(input string, seqsize int) int {
	aftern := seqsize

	for i := 0; i < len(input); i++ {
		if areAllCharsUnique(input[i : i+seqsize]) {
			log.Println(input[i : i+seqsize])
			break
		} else {
			aftern++
		}
	}

	return aftern
}

// areAllCharsUnique returns true if all numbers in the input slice are unique, false otherwise
func areAllCharsUnique(seq string) bool {
	// Create a map to store the chars that have been seen
	seenChars := make(map[int32]bool)

	// Iterate over the chars in the slice
	for _, char := range seq {
		// Check if the number has already been added to the map
		if seenChars[char] {
			// Return false if the number has already been seen
			return false
		}

		// Add the number to the map
		seenChars[char] = true
	}

	// Return true if all chars are unique
	return true
}
