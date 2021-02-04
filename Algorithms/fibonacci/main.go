package fibonacci

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
		fmt.Print(strconv.Itoa(Recursion(i)) + " ")
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

//Recursion to find fibonacci sequence
func Recursion(n int) int {
	if n <= 1 {
		return n
	}
	return Recursion(n-1) + Recursion(n-2)
}

//ResultAsString return sequence of fibonacci as string type
func ResultAsString(n int) string {
	// var result []int
	var fiboStr string = ""
	for i := 0; i < n; i++ {
		fiboStr += strconv.Itoa(Recursion(i))
		fiboStr += " "
		// result = append(result, Recursion(i))
	}
	return fiboStr
}

//ResultAsArray return sequence of fibonacci as array type
func ResultAsArray(n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, Recursion(i))
	}
	return result
}
