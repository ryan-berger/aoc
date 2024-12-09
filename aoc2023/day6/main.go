package main

import (
	_ "embed"
	"fmt"
)

var races = [][2]int64{
	{59707878, 430121812131276},
}

//var races = [][2]int64{
//	{7, 9},
//	{15, 40},
//	{30, 200},
//}

func simulate(holdFor, distance int64) int64 {
	return holdFor * (distance - holdFor)
}

func part1() {
	var counts []int64
	for _, race := range races {
		count := int64(0)
		for i := int64(0); i < race[0]; i++ {
			if simulate(i, race[0]) > race[1] {
				count++
			}
		}
		counts = append(counts, count)
	}

	mul := int64(1)
	for _, c := range counts {
		mul *= c
	}
	fmt.Println(mul)
}

func part2() {
	var counts []int64
	for _, race := range races {
		count := int64(0)
		for i := int64(0); i < race[0]; i++ {
			if simulate(i, race[0]) > race[1] {
				count++
			}
		}
		counts = append(counts, count)
	}

	mul := int64(1)
	for _, c := range counts {
		mul *= c
	}
	fmt.Println(mul)
}

func main() {
	part1()
	part2()
}
