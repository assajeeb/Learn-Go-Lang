package main

import (
	"fmt"

	mangeFile "example.com/practice-tax-calculator/manageFiles"
	"example.com/practice-tax-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.3, 0.6, 0.1}
	doneChaneList := make([]chan bool, len(taxRates))
	errorChanList := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		manageFile := mangeFile.NewManageFile("Prices.txt", fmt.Sprintf("result %0.2f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJon(taxRate, &manageFile)
		doneChaneList[index] = make(chan bool)
		errorChanList[index] = make(chan error)

		go priceJob.Process(doneChaneList[index], errorChanList[index])
	}

	for index, _ := range taxRates {
		select {
		case err := <-errorChanList[index]:
			fmt.Println(err)
		case <-doneChaneList[index]:
			fmt.Println("done!")

		}
	}

}
