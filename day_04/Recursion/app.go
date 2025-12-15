package main

import "fmt"

func main() {
	fact := facetorial(5)
	fmt.Println(fact)
}

func facetorial(number int) int {
	fmt.Println("this number is now: ", number)
	if number == 0 {
		return 1
	}

	return number * facetorial(number-1)
}
