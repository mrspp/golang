package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Finding primes")
	arr := []int{1, 2, 34, 5, 6, 344, 11, 17}
	findingPrimes(arr)
}

func isPrime(value int) bool {
	for i := 2; i < int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func findingPrimes(arr []int) {
	for _, val := range arr {
		if isPrime(val) {
			fmt.Println(val)
		} else {
			continue
		}
	}
}
