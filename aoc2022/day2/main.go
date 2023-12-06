package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func pt1() {
	score := 0
	for _, line := range strings.Split(input, "\n") {
		opponentStr, youStr, _ := strings.Cut(line, " ")

		score += int(youStr[0] - 'W')

		outcome := int(opponentStr[0]-'A') - int(youStr[0]-'X')

		switch {
		case outcome == 0:
			score += 3
		case outcome == -1 || outcome == 2:
			score += 6
		}

	}
	fmt.Println("p1 answer:", score)
}

func pt2() {
	score := 0
	for _, line := range strings.Split(input, "\n") {
		opponentStr, youStr, _ := strings.Cut(line, " ")

		opponent := int(opponentStr[0] - 'A')
		switch youStr[0] {
		case 'X':
			score += ((opponent + 2) % 3) + 1
		case 'Y':
			score += 3 + (opponent + 1)
		case 'Z':
			score += 6 + ((opponent + 1) % 3) + 1
		}
	}
	fmt.Println("p2 answer:", score)
}

func main() {
	pt1()
	pt2()
}
