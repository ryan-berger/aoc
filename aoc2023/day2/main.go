package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var day1Max = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func day1() {
	lines := strings.Split(input, "\n")
	sum := 0

Line:
	for _, l := range lines {
		var idStr string

		rest := l[len("Game "):]
		idStr, rest, _ = strings.Cut(rest, ": ")
		id, _ := strconv.Atoi(idStr)

		maxColors := make(map[string]int)
		drawStrings := strings.Split(rest, "; ")
		for _, ds := range drawStrings {
			colorStrs := strings.Split(ds, ", ")
			for _, colorStr := range colorStrs {
				countStr, color, _ := strings.Cut(colorStr, " ")
				count, _ := strconv.Atoi(countStr)
				maxColors[color] = max(maxColors[color], count)
			}
		}

		for color, v := range day1Max {
			if maxColors[color] > v {
				continue Line
			}
		}
		sum += id
	}
	fmt.Println(sum)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func day2() {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, l := range lines {

		rest := l[len("Game "):]
		_, rest, _ = strings.Cut(rest, ": ")

		minColors := make(map[string]int)
		drawStrings := strings.Split(rest, "; ")
		for _, ds := range drawStrings {
			colorStrs := strings.Split(ds, ", ")
			for _, colorStr := range colorStrs {
				countStr, color, _ := strings.Cut(colorStr, " ")
				count, _ := strconv.Atoi(countStr)
				if val, ok := minColors[color]; ok {
					minColors[color] = max(val, count)
				} else {
					minColors[color] = count
				}
			}
		}

		power := 1
		for color := range day1Max {
			power *= minColors[color]
		}
		sum += power
	}
	fmt.Println(sum)
}

func main() {
	day1()
	day2()
}
