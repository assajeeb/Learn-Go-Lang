// Build a "Profit Calculator"
// ask for Revenue, Expenses, tax rate
// Calculate earnings before tax (EBT) and earning after tax(profit)
// Calculate ratio (EBT / Profit)
// Output EBT , profit and the ratio

// use function

package main

import "fmt"

func main() {

	revenue := takeInput("Enter Revene: ")
	expenses := takeInput("Enter Revenue: ")
	taxRate := takeInput("Enter TaxRate: ")
	ebt, profit, ratio := CalculateResult(revenue, expenses, taxRate)

	fmt.Printf("EBT: %0.1f \n", ebt)
	fmt.Printf("Profit: %0.1f \n", profit)
	fmt.Printf("Ratio: %0.1f \n", ratio)

}

func takeInput(whatNeed string) (output float64) {

	fmt.Print(whatNeed)

	fmt.Scan(&output)

	return
}

func CalculateResult(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {

	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
