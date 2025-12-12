// Build a "Profit Calculator"
// ask for Revenue, Expenses, tax rate
// Calculate earnings before tax (EBT) and earning after tax(profit)
// Calculate ratio (EBT / Profit)
// Output EBT , profit and the ratio

// use function

//Goals

// 1) Validate user input
// => show error message & exit if invalid input is provided
// -no nefative numbers
// - NOt 0
// 2) store calculated results into file

package main

import (
	"fmt"
	"os"
)

var fileName string = "store.txt"

func storeResult(ebt, profit, ratio float64) {
	resuult := fmt.Sprintf("EBT: %0.1f \n Profit: %0.1f \n ratio: %0.1f \n", ebt, profit, ratio)
	os.WriteFile(fileName, []byte(resuult), 0644)
}
func main() {

	revenue, _ := takeInput("Enter Revene: ")
	expenses, _ := takeInput("Enter Revenue: ")
	taxRate, _ := takeInput("Enter TaxRate: ")
	ebt, profit, ratio := CalculateResult(revenue, expenses, taxRate)

	storeResult(ebt, profit, ratio)

	fmt.Printf("EBT: %0.1f \n", ebt)
	fmt.Printf("Profit: %0.1f \n", profit)
	fmt.Printf("Ratio: %0.1f \n", ratio)

}

func takeInput(whatNeed string) (float64, error) {

	fmt.Print(whatNeed)

	var inputData float64

	_, err := fmt.Scan(&inputData)
	if err != nil || inputData <= 0 {
		panic(" Invaild Input")
	}

	return inputData, err
}

func CalculateResult(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {

	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
