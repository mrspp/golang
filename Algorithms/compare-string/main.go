package main

import "fmt"

func main() {
	fmt.Println("Finding longest string")
	fmt.Println(findingLongerString("asafsdgs", "asfsdghsyher"))
}

func findingLongerString(str1, str2 string) string {
	if len(str1) > len(str2) {
		return str1
	} else {
		return str2
	}
}
