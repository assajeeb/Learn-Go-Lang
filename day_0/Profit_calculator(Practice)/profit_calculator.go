// Build a "Profit Calculator"
// ask for Revenue, Expenses, tax rate
// Calculate earnings before tax (EBT) and earning after tax(profit)
// Calculate ratio (EBT / Profit)
// Output EBT , profit and the ratio

package main

import "fmt"

func main() {

	var revenue, expenses, taxRate float64

	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)

}
