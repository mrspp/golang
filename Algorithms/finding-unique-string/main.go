package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Finding unique string")
	findUniqueNumber([]int{2, 45, 5, 1, 46, 3, 54, 13, 4, 5, 6, 23})
}

func findUniqueNumber(arr []int) []int {
	sort.Ints(arr)
	fmt.Println(arr)
	return arr
}
