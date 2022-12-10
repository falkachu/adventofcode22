package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	a = iota + 1
	b
	c
	d
	e
	f
	g
	h
	i
	j
	k
	l
	m
	n
	o
	p
	q
	r
	s
	t
	u
	v
	w
	x
	y
	z
	A
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
)

func main() {
	log.Println("Day day3")

	var filepath = "input.txt"

	items := readInput(filepath)
	sitems := splitItems(items)
	matches := findMatch(sitems)
	calcPoints(matches)

	s2items := splitThree(items)
	matches2 := findMatch2(s2items)
	calcPoints(matches2)
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

func findMatch(items [][]string) []string {

	var matches []string

	for _, item := range items {

		match := ""

		for _, cs := range item[0] {
			for _, ce := range item[1] {
				if cs == ce {
					match = fmt.Sprintf("%c", cs)
				}
			}
		}

		matches = append(matches, match)
	}

	return matches
}

func calcPoints(items []string) {
	sum := 0

	for _, it := range items {
		switch it {
		case "a":
			sum += a
		case "b":
			sum += b
		case "c":
			sum += c
		case "d":
			sum += d
		case "e":
			sum += e
		case "f":
			sum += f
		case "g":
			sum += g
		case "h":
			sum += h
		case "i":
			sum += i
		case "j":
			sum += j
		case "k":
			sum += k
		case "l":
			sum += l
		case "m":
			sum += m
		case "n":
			sum += n
		case "o":
			sum += o
		case "p":
			sum += p
		case "q":
			sum += q
		case "r":
			sum += r
		case "s":
			sum += s
		case "t":
			sum += t
		case "u":
			sum += u
		case "v":
			sum += v
		case "w":
			sum += w
		case "x":
			sum += x
		case "y":
			sum += y
		case "z":
			sum += z
		case "A":
			sum += A
		case "B":
			sum += B
		case "C":
			sum += C
		case "D":
			sum += D
		case "E":
			sum += E
		case "F":
			sum += F
		case "G":
			sum += G
		case "H":
			sum += H
		case "I":
			sum += I
		case "J":
			sum += J
		case "K":
			sum += K
		case "L":
			sum += L
		case "M":
			sum += M
		case "N":
			sum += N
		case "O":
			sum += O
		case "P":
			sum += P
		case "Q":
			sum += Q
		case "R":
			sum += R
		case "S":
			sum += S
		case "T":
			sum += T
		case "U":
			sum += U
		case "V":
			sum += V
		case "W":
			sum += W
		case "X":
			sum += X
		case "Y":
			sum += Y
		case "Z":
			sum += Z
		}
	}

	log.Println(sum)
}

func splitThree(items []string) [][]string {
	var splitThree [][]string
	var spli []string

	counter := 1
	pairs := 0

	for _, i := range items {
		if counter < 4 {
			spli = append(spli, i)
			counter++
		}

		if counter == 4 {
			splitThree = append(splitThree, spli)
			log.Printf("PairN: %v | Pair: %v", pairs, spli)
			pairs++
			spli = nil
			counter = 1
		}
	}

	return splitThree
}

func findMatch2(items [][]string) []string {
	var matches []string

	for pair, item := range items {

		match := ""

		for _, cs := range item[0] {
			for _, cm := range item[1] {
				for _, ce := range item[2] {
					if cs == cm && cs == ce {
						match = fmt.Sprintf("%c", cs)
					}
				}
			}
		}

		log.Printf("Pair: %v | Match: %v", pair, match)
		matches = append(matches, match)
	}

	return matches
}
