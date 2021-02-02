package main

import "fmt"

var (
	sum = func(a int, b int) int {
		return a + b
	}
)

func main() {
	fmt.Println(sum(3, 5))
}
