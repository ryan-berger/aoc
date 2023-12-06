package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func part1() {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		first, last := len(line), -1
		for i, c := range line {
			if c >= '0' && c <= '9' {
				first = min(i, first)
				last = max(i, last)
			}
		}
		sum += int(line[first]-'0')*10 + int(line[last]-'0')
	}
	fmt.Println(sum)
}

var strToNum = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func part2() {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		// forward
	Forward:
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c >= '0' && c <= '9' {
				sum += int(c-'0') * 10
				break Forward
			}

			for numStr, v := range strToNum {
				if strings.HasPrefix(line[i:], numStr) {
					sum += v * 10
					i += len(numStr) - 1
					break Forward
				}
			}
		}
	Backward:
		for i := len(line) - 1; i >= 0; i-- {
			c := line[i]
			if c >= '0' && c <= '9' {
				sum += int(c - '0')
				break Backward
			}

			for numStr, v := range strToNum {
				if strings.HasPrefix(line[i:], numStr) {
					sum += v
					i += len(numStr) - 1
					break Backward
				}
			}
		}

	}
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
