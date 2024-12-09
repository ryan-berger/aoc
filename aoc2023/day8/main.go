package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func parseGraph(graphStr string) (string, map[string][2]string) {
	graph := make(map[string][2]string)

	lines := strings.Split(graphStr, "\n")
	instructions := lines[0]

	for _, line := range lines[2:] {
		source := line[:3]

		startLeft := len("AAA = (")

		left := line[startLeft : startLeft+3]
		line = line[startLeft+3+2:]
		right := line[:3]

		graph[source] = [2]string{left, right}
	}

	return instructions, graph
}

func followInstructions(insts, start, end string, graph map[string][2]string) int {
	mod := len(insts)
	counter := 0

	cur := start
	for cur[3-len(end):] != end {

		direction := insts[counter%mod]
		if direction == 'L' {
			cur = graph[cur][0]
		} else {
			cur = graph[cur][1]
		}

		counter++
	}

	return counter
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func main() {
	instructions, graph := parseGraph(input)

	var starts []string
	for k := range graph {
		if k[2] == 'A' {
			starts = append(starts, k)
		}
	}

	var cycles []int
	for _, start := range starts {
		cycles = append(cycles, followInstructions(instructions, start, "Z", graph))
	}

	fmt.Println(lcm(cycles[0], cycles[1], cycles[2:]...))
}
