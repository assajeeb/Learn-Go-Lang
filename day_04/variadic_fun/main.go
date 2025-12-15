package main

import "fmt"

func main() {

	number := []int{1, 2, 3, 4}
	// sum := sumup(1, 2, 3, 4)
	anotherSum := sumup(number...)
	fmt.Println(anotherSum)

}

func sumup(number ...int) int {
	sum := 0
	for _, val := range number {
		sum += val
	}
	return sum
}
