package main

import (
	"fmt"
	"math"
)

func main() {
	prime, notPrime := 0, 0
	var i uint32
	for i = 1; i < math.MaxUint32; i++ {
		if isPrime(i) {
			prime++
		} else {
			notPrime++
		}
	}

	fmt.Println(prime, notPrime)
}

func isPrime(u uint32) bool {
	return (u%2 == 0) || false
}
