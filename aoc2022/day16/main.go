package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const input = `Valve YK has flow rate=0; tunnels lead to valves GL, FT
Valve QA has flow rate=0; tunnels lead to valves JX, FD
Valve LN has flow rate=0; tunnels lead to valves FD, FG
Valve AU has flow rate=0; tunnels lead to valves BD, PQ
Valve MM has flow rate=0; tunnels lead to valves UL, AA
Valve JX has flow rate=0; tunnels lead to valves QA, NZ
Valve CV has flow rate=0; tunnels lead to valves UP, QW
Valve UZ has flow rate=0; tunnels lead to valves FG, NZ
Valve BP has flow rate=0; tunnels lead to valves TI, DX
Valve NS has flow rate=0; tunnels lead to valves ZL, CW
Valve CO has flow rate=0; tunnels lead to valves BD, AT
Valve RZ has flow rate=0; tunnels lead to valves AA, ZO
Valve PQ has flow rate=0; tunnels lead to valves ML, AU
Valve CW has flow rate=7; tunnels lead to valves UL, PH, OF, NS, GT
Valve FG has flow rate=14; tunnels lead to valves SO, JR, IN, LN, UZ
Valve EZ has flow rate=0; tunnels lead to valves UP, QP
Valve GN has flow rate=0; tunnels lead to valves VQ, CH
Valve QW has flow rate=6; tunnels lead to valves CV, PF, KH, UY, TI
Valve UL has flow rate=0; tunnels lead to valves MM, CW
Valve VQ has flow rate=12; tunnels lead to valves GN, LC
Valve FT has flow rate=0; tunnels lead to valves SG, YK
Valve SG has flow rate=21; tunnels lead to valves FT, LC, NO, QX
Valve BD has flow rate=23; tunnels lead to valves CO, AU, AB
Valve AB has flow rate=0; tunnels lead to valves BD, QX
Valve QP has flow rate=0; tunnels lead to valves AD, EZ
Valve OF has flow rate=0; tunnels lead to valves DX, CW
Valve AA has flow rate=0; tunnels lead to valves QL, RZ, SO, MM, HW
Valve RQ has flow rate=0; tunnels lead to valves GL, QG
Valve ZL has flow rate=0; tunnels lead to valves NS, FD
Valve KH has flow rate=0; tunnels lead to valves GT, QW
Valve JR has flow rate=0; tunnels lead to valves FG, PH
Valve PH has flow rate=0; tunnels lead to valves CW, JR
Valve LC has flow rate=0; tunnels lead to valves VQ, SG
Valve FD has flow rate=20; tunnels lead to valves LN, QA, ZL
Valve NZ has flow rate=15; tunnels lead to valves UZ, JX
Valve ML has flow rate=22; tunnels lead to valves OW, PQ, NO
Valve PF has flow rate=0; tunnels lead to valves QW, CH
Valve UP has flow rate=19; tunnels lead to valves RY, CV, EZ
Valve VM has flow rate=0; tunnels lead to valves RY, CH
Valve DX has flow rate=3; tunnels lead to valves BO, QL, BP, OF, QG
Valve QL has flow rate=0; tunnels lead to valves AA, DX
Valve HW has flow rate=0; tunnels lead to valves UY, AA
Valve GL has flow rate=8; tunnels lead to valves YK, RQ
Valve QG has flow rate=0; tunnels lead to valves DX, RQ
Valve IN has flow rate=0; tunnels lead to valves FG, BO
Valve NO has flow rate=0; tunnels lead to valves SG, ML
Valve SO has flow rate=0; tunnels lead to valves FG, AA
Valve RY has flow rate=0; tunnels lead to valves UP, VM
Valve CH has flow rate=13; tunnels lead to valves GN, VM, PF, ZO
Valve AD has flow rate=17; tunnel leads to valve QP
Valve TI has flow rate=0; tunnels lead to valves BP, QW
Valve UY has flow rate=0; tunnels lead to valves HW, QW
Valve AT has flow rate=24; tunnels lead to valves OW, CO
Valve GT has flow rate=0; tunnels lead to valves CW, KH
Valve ZO has flow rate=0; tunnels lead to valves RZ, CH
Valve QX has flow rate=0; tunnels lead to valves AB, SG
Valve BO has flow rate=0; tunnels lead to valves IN, DX
Valve OW has flow rate=0; tunnels lead to valves AT, ML`

var re = regexp.MustCompile("Valve (\\w+) has flow rate=(\\d+); tunnels? leads? to valves? ([\\w, ]+)")

type graph struct {
	costs map[string]int
	paths map[string][]string
}

func (g graph) cpy() graph {
	ng := graph{
		costs: make(map[string]int),
		paths: make(map[string][]string),
	}

	for k, v := range g.costs {
		ng.costs[k] = v
	}

	for k, v := range g.paths {
		ng.paths[k] = v
	}

	return ng
}

type path struct {
	from, to string
}

var pathCache = make(map[path][]string)

func getPath(from, to string, g graph) []string {
	if p, ok := pathCache[path{from: from, to: to}]; ok {
		return p
	}

	visited := make(map[string]bool)
	visited[from] = true

	visiting := []string{from}
	pred := make(map[string]string)

	for len(visiting) != 0 {
		var next []string
		nextUniq := make(map[string]bool)

		for _, v := range visiting {
			for _, n := range g.paths[v] {
				if visited[n] || nextUniq[n] {
					continue
				}

				nextUniq[n] = true
				visited[n] = true

				next = append(next, n)
				pred[n] = v
				if n == to {
					goto calc
				}
			}
		}

		visiting = next
	}
calc:

	p := []string{to}
	for next, ok := pred[to]; ok; next, ok = pred[next] {
		p = append([]string{next}, p...)
	}

	pathCache[path{from: from, to: to}] = p
	return p
}

func pt1(start string, budget int, g graph) int {
	if budget <= 0 {
		return 0
	}

	var candidates []string
	for k, v := range g.costs {
		if k != start && v > 0 {
			candidates = append(candidates, k)
		}
	}

	score := 0
	for _, candidate := range candidates {
		path := getPath(start, candidate, g)

		ng := g.cpy()
		ng.costs[candidate] = 0

		newScore := g.costs[candidate]*(budget-len(path)) + pt1(candidate, budget-len(path), ng)
		if newScore > score {
			score = newScore
		}
	}

	return score
}

func max(elems ...int) int {
	if len(elems) == 0 {
		panic("derp")
	}

	m := elems[0]
	for i := 1; i < len(elems); i++ {
		if elems[i] > m {
			m = elems[i]
		}
	}
	return m
}

func pt2(start [2]string, budget int, g graph) int {
	if budget <= 0 {
		return 0
	}

	var candidates []string
	for k, v := range g.costs {
		if v > 0 {
			candidates = append(candidates, k)
		}
	}

	score := 0
	for i, c := range candidates {
		me := getPath(start[0], c, g)

		for j, c2 := range candidates {
			elef := getPath(start[1], c2, g)
			bothScore := 0
			bg := g.cpy()

			if i == j {
				bg.costs[c] = 0
				elefScore := g.costs[c]*(budget-len(elef)) + pt2([2]string{start[0], c2}, budget-len(elef), bg)
				meScore := g.costs[c]*(budget-len(me)) + pt2([2]string{c, start[1]}, budget-len(me), bg)

				score = max(score, elefScore, meScore)
				continue
			}

			switch {
			case len(elef) == len(me):
				bg.costs[c] = 0
				bg.costs[c2] = 0
				bothScore = g.costs[c]*(budget-len(me)) + g.costs[c2]*(budget-len(elef)) + pt2([2]string{c, c2}, budget-len(me), bg)
			case len(me) < len(elef):
				bg.costs[c] = 0
				bothScore = g.costs[c]*(budget-len(me)) + pt2([2]string{c, elef[len(me)]}, budget-len(me), bg)
			case len(me) > len(elef):
				bg.costs[c2] = 0
				bothScore = g.costs[c2]*(budget-len(elef)) + pt2([2]string{me[len(elef)], c2}, budget-len(elef), bg)
			}

			score = max(score, bothScore)
		}
	}

	return score
}

func main() {
	split := strings.Split(input, "\n")

	costs := make(map[string]int)
	paths := make(map[string][]string)

	var start = "AA"
	for _, s := range split {
		matches := re.FindStringSubmatch(s)
		matches = matches[1:]

		name, rateStr, tunnels := matches[0], matches[1], matches[2]
		rate, _ := strconv.Atoi(rateStr)

		costs[name] = rate

		for _, t := range strings.Split(tunnels, ", ") {
			paths[name] = append(paths[name], t)
		}
	}

	g := graph{
		costs: costs,
		paths: paths,
	}

	fmt.Println(pt1(start, 30, g))
	fmt.Println(pt2([2]string{start, start}, 26, g))

}
