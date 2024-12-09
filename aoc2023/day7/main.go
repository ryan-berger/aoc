package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scoreCard(card string) int {
	counts := make(map[rune]int)
	for _, r := range card {
		counts[r]++
	}

	var uniq []rune
	for k := range counts {
		uniq = append(uniq, k)
	}

	switch len(counts) {
	case 1:
		return 7 // five of a kind
	case 2:
		if min(counts[uniq[0]], counts[uniq[1]]) == 1 {
			return 6 // four of a kind
		}
		return 5 // full house
	case 3:
		if max(max(counts[uniq[0]], counts[uniq[1]]), counts[uniq[2]]) == 3 {
			return 4 // three of a kind
		}
		return 3 // two pair
	case 4:
		return 2
	case 5:
		return 1
	}

	panic("")
}

func part1() {
	type hand struct {
		hand  string
		bid   int
		score int
	}
	var hands []hand

	for _, line := range strings.Split(input, "\n") {
		handStr, bidStr, _ := strings.Cut(line, " ")
		bid, _ := strconv.Atoi(bidStr)

		hands = append(hands, hand{
			hand:  handStr,
			bid:   bid,
			score: scoreCard(handStr),
		})
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].score < hands[j].score && hands[i].hand < hands[j].hand
	})

	sum := 0
	for i := 0; i < len(hands); i++ {
		fmt.Println((i + 1), hands[i].bid, hands[i])
		sum += (i + 1) * hands[i].bid
	}
	fmt.Println(sum)
}

func main() {
	part1()
}
