package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Sort array")
	arr := []int{2, 4, 1, 7, 6, 78, 34, 235}

	// s := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Ints(arr)
	// fmt.Println(sortArray(arr))
	fmt.Println(arr)
	fmt.Println(sortArray(arr))
}

func sortArray(a []int) []int {
	var temp int = a[1]

	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] < a[j] {
				temp = a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}
