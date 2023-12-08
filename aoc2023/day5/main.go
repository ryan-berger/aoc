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

type mapEntry struct {
	from [2]int
	to   [2]int
}

type rangeMap []mapEntry

func mapTo[T, U any](ts []T, f func(T) U) []U {
	res := make([]U, len(ts))
	for i, t := range ts {
		res[i] = f(t)
	}
	return res
}

func toInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func parseMaps() ([]int, []rangeMap) {
	lines := strings.Split(input, "\n")
	seeds := mapTo(strings.Split(lines[0][len("seeds: "):], " "), toInt)

	var maps []rangeMap

	for i := 2; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		if r := lines[i][0]; r < '0' || r > '9' {
			continue
		}

		var entries []mapEntry
		for ; i < len(lines) && lines[i] != ""; i++ {
			rangeDef := mapTo(strings.Split(lines[i], " "), toInt)
			dest, source, length := rangeDef[0], rangeDef[1], rangeDef[2]
			entry := mapEntry{
				from: [2]int{source, source + length},
				to:   [2]int{dest, dest + length},
			}
			entries = append(entries, entry)
		}

		maps = append(maps, entries)
	}

	return seeds, maps
}

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

func minFrom(l []int) int {
	m := l[0]
	for _, e := range l {
		if e < m {
			m = e
		}
	}
	return m
}

func part1(seeds []int, maps []rangeMap) {
	var mapped []int
	for _, seed := range seeds {
		curSeed := seed

		for _, m := range maps {
			i, found := sort.Find(len(m), func(idx int) int {
				fromStart, fromEnd := m[idx].from[0], m[idx].from[1]
				start := curSeed
				switch {
				case start < fromStart:
					return -1
				case start > fromEnd:
					return 1
				case start >= fromStart && start <= fromEnd:
					return 0
				default:
					panic("oops!")
				}
			})

			if !found {
				continue
			}

			curSeed += m[i].to[0] - m[i].from[0]
		}
		mapped = append(mapped, curSeed)
	}
	fmt.Println(minFrom(mapped))
}

func bfsRanges(ranges [][2]int, maps []rangeMap) [][2]int {
	curRanges := ranges
	for _, m := range maps {
		var nextRanges [][2]int
		for _, r := range curRanges {
			i, found := sort.Find(len(m), func(idx int) int {
				fromStart, fromEnd := m[idx].from[0], m[idx].from[1]
				start := r[0]
				switch {
				case start < fromStart:
					return -1
				case start > fromEnd:
					return 1
				case start >= fromStart && start <= fromEnd:
					return 0
				default:
					panic("oops!")
				}
			})

			if !found {
				nextRanges = append(nextRanges, r)
				continue
			}

			for j := i; j < len(m); j++ {
				cmpRange := m[j].from
				diff := m[j].to[0] - cmpRange[0]

				smallestRange := [2]int{max(cmpRange[0], r[0]) + diff, min(cmpRange[1], r[1]) + diff}
				nextRanges = append(nextRanges, smallestRange)
				if cmpRange[1] > r[1] {
					break
				}
			}
		}
		curRanges = nextRanges
	}

	return curRanges
}

func minFromRange(ranges [][2]int) int {
	m := ranges[0][0]
	for _, r := range ranges {
		if r[0] < m {
			m = r[0]
		}
	}
	return m
}

func part2(seeds []int, maps []rangeMap) {
	var ranges [][2]int
	for i := 0; i < len(seeds); i += 2 {
		ranges = append(ranges, [2]int{seeds[i], seeds[i] + seeds[i+1] - 1})
	}

	var results []int
	for _, r := range ranges {
		results = append(results, minFromRange(bfsRanges([][2]int{r}, maps)))
	}
	fmt.Println(minFrom(results))
}

func main() {
	seeds, maps := parseMaps()
	for i := range maps {
		m := maps[i]

		sort.Slice(m, func(i, j int) bool {
			return m[i].from[0] < m[j].from[0]
		})
	}

	part1(seeds, maps)
	part2(seeds, maps)
}
