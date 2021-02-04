package isprime

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Finding primes")
	arr := []int{1, 2, 34, 5, 6, 344, 11, 17}
	findingPrimes(arr)
}

//IsPrime check if a number is prime
func IsPrime(value int) bool {
	for i := 2; i < int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func findingPrimes(arr []int) {
	for _, val := range arr {
		if IsPrime(val) {
			fmt.Println(val)
		} else {
			continue
		}
	}
}
