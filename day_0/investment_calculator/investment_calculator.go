package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmmount float64
	expectedReturnRate := 5.5
	var year float64
	fmt.Print("Enter investment Amount: ")
	fmt.Scan(&investmentAmmount)

	fmt.Print("Enter ExpectedReturn Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter investment year: ")
	fmt.Scan(&year)

	var futureValue = investmentAmmount * math.Pow((1+expectedReturnRate/100), year)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, year)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
