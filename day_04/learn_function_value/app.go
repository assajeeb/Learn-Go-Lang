package functionvalue

import "fmt"

type transformFn func(int) int

func main() {

	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{2, 3, 4, 5}
	double := transformNumbers(&numbers, double)
	triple := transformNumbers(&numbers, triple)
	fmt.Println(double, triple)

	transformNumbers1 := getTransformerFunction(&numbers)
	transformNumbers2 := getTransformerFunction(&moreNumbers)

	transfromedNumbers := transformNumbers(&numbers, transformNumbers1)
	moreTransformedNumber := transformNumbers(&moreNumbers, transformNumbers2)

	fmt.Println(transfromedNumbers)
	fmt.Println(moreTransformedNumber)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {

	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value*2))
	}

	return dNumbers

}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(value int) int {
	return value * 2
}

func triple(value int) int {
	return value * 3
}
