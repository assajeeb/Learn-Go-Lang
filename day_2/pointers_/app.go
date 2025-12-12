package main

import "fmt"

func main() {
	age := 32 // regular variable

	// var agePointer *int
	agePointer := &age

	// fmt.Println("Age: ", *agePointer)
	// adultYears := getAdultYears(&age)
	EditAgeTYoAdultYears(&age)

	fmt.Println(age)

	fmt.Println(*agePointer)
}

func EditAgeTYoAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
