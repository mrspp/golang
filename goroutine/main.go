package main

import (
	"fmt"
	"strings"
)

func toUpper(str string, done chan string) {
	val := strings.ToUpper(str)
	done <- val
	fmt.Println(val)
}

func main() {
	fmt.Println("Goroutine Tutorial")

	// notice how we've added the 'go' keyword
	// in front of both our compute function calls

	var input string
	fmt.Scanln(&input)
}
