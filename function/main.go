package main

import "fmt"

func result(a int, b int) (sum int, mul int) {
	sum = a + b
	mul = a * b
	return // Return statement without specify variable name
}

func main() {
	var sum, mul int
	sum, mul = result(2, 5)
	fmt.Println(sum)
	fmt.Println(mul)
}
