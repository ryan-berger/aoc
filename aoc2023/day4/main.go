package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type card struct {
	num     int
	winning map[int]struct{}
	ours    map[int]struct{}
}

func parseCard(idx int, cardStr string) card {
	collectNums := func(s string) map[int]struct{} {
		nums := make(map[int]struct{})
		for _, numStr := range strings.Split(s, " ") {
			if numStr == "" {
				continue
			}
			num, _ := strconv.Atoi(numStr)
			nums[num] = struct{}{}
		}

		return nums
	}

	winning, ours, _ := strings.Cut(cardStr, " | ")

	return card{
		num:     idx + 1,
		winning: collectNums(winning),
		ours:    collectNums(ours),
	}
}

func parseCards() []card {
	var cards []card
	for i, line := range strings.Split(input, "\n") {
		_, rest, _ := strings.Cut(line, ":")
		cards = append(cards, parseCard(i, rest))
	}
	return cards
}

func part1(cards []card) {
	sum := 0
	for _, c := range cards {
		count := 0

		for winning := range c.winning {
			if _, ok := c.ours[winning]; ok {
				count++
			}
		}

		if count != 0 {
			sum += 1 << (count - 1)
		}
	}
	fmt.Println(sum)
}

func part2(cards []card) {
	count := make(map[int]int)
	for i := range cards {
		count[i] = 1
	}

	for i, c := range cards {
		winCount := 0
		for winning := range c.winning {
			if _, ok := c.ours[winning]; ok {
				winCount++
			}
		}
		for j := 1; j <= winCount; j++ {
			count[i+j] += count[i]
		}
	}

	sum := 0
	for _, v := range count {
		sum += v
	}
	fmt.Println(sum)

}

func main() {
	cards := parseCards()
	part1(cards)
	part2(cards)
}
