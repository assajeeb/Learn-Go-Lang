package prices

import (
	"fmt"

	"example.com/practice-tax-calculator/conversition"
	"example.com/practice-tax-calculator/iomanager"
)

type TaxInculedPriceJob struct {
	TaxRate           float64            `json:"tax_rate"`
	InputPrices       []float64          `json:"input_prices"`
	TaxIncluededPrice map[string]string  `json:"tax_included_prices"`
	IoManger          iomanager.IoManger `json:"-"`
}

func (job *TaxInculedPriceJob) LoadData() {
	prices, erro := job.IoManger.LoadDataFromFile()
	if erro != nil {
		fmt.Println(erro)

	}
	inputPrices, err := conversition.StringToFloat(prices)
	if err != nil {
		fmt.Println(err)
		panic(err)

	}
	job.InputPrices = inputPrices
}

func NewTaxIncludedPriceJon(taxrate float64, manageFile iomanager.IoManger) *TaxInculedPriceJob {

	return &TaxInculedPriceJob{
		TaxRate:  taxrate,
		IoManger: manageFile,
	}
}

func (job *TaxInculedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxInculedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%0.2f", taxInculedPrice)
	}
	job.TaxIncluededPrice = result

	job.IoManger.WriteJson(job)

}
