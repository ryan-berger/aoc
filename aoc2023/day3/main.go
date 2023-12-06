package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func printAdjacent(grid [][]rune, line, lIndex, rIndex int) {
	var strs []string

	for i := line - 1; i <= line+1; i++ {
		var str string
		for j := lIndex - 1; j < rIndex+1; j++ {
			if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) {
				str += string(grid[i][j])
			}
		}
		if len(str) != 0 {
			strs = append(strs, str)
		}
	}
	fmt.Println(strings.Join(strs, "\n"))
}
func adjacentSymbol(grid [][]rune, line, lIndex, rIndex int) bool {
	for i := line - 1; i <= line+1; i++ {
		for j := lIndex - 1; j < rIndex+1; j++ {
			if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) {
				if r := grid[i][j]; r != '.' && (r < '0' || r > '9') {
					printAdjacent(grid, line, lIndex, rIndex)
					return true
				}
			}
		}
	}
	return false
}

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}

func day1() {
	var grid [][]rune
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}

	sum := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			k := j
			if !isNum(grid[i][k]) {
				continue
			}

			for k < len(grid[i]) && isNum(grid[i][k]) {
				k++
			}

			if j != k && adjacentSymbol(grid, i, j, k) {
				num, _ := strconv.Atoi(string(grid[i][j:k]))
				fmt.Println(string(grid[i][j:k]))
				fmt.Println()
				sum += num
			}
			j = k
		}
	}

	fmt.Println(sum)
}

func day2() {}

func main() {
	day1()
}
