package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("fibonacci")
	n := 10
	fmt.Print("Fibonacci Loop: ")
	for i := 0; i < n; i++ {
		fmt.Print(strconv.Itoa(fibonacciLoop(i)) + " ")
	}
	fmt.Println()
	fmt.Print("Fibonacci Recursion: ")
	for i := 0; i < n; i++ {
		fmt.Print(strconv.Itoa(fibonacciRecursion(i)) + " ")
	}

}

func fibonacciLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func fibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}
