package main

import "fmt"

func main() {
	fmt.Print(add(1, 2))
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
