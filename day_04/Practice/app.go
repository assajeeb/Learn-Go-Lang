package main

import (
	"example.com/practice-tax-calculator/cmdmanager"
	"example.com/practice-tax-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.3, 0.6, 0.1}

	for _, taxRate := range taxRates {
		// manageFile := mangeFile.NewManageFile("Prices.txt", fmt.Sprintf("result %0.2f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJon(taxRate, &cmdm)

		priceJob.Process()
	}

}
