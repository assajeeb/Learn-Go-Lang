package main

import "fmt"

type floatMap map[string]float64

func main() {

	userName := make([]string, 2, 5)

	userName[0] = "Julie"
	userName = append(userName, "max")
	userName = append(userName, "Hello")

	fmt.Println(userName)

	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 222.2
	courseRatings["go1"] = 222.1
	courseRatings["go2"] = 222.2
	courseRatings.output()

	for index, value := range userName {
		fmt.Println(index, value)
	}

	for key, value := range courseRatings {
		fmt.Println(key, value)
	}

}

func (m floatMap) output() {
	fmt.Print(m)
}
