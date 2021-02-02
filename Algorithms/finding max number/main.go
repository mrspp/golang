package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxNumber(numbers []int) int {
	max := 0
	for _, value := range numbers {
		if value >= max {
			max = value
		}
	}
	return max
}

func main() {
	result := maxNumber([]int{3, 45, 6, 3, 57, 678, 3})
	fmt.Printf("Max number is: %d", result)

}
