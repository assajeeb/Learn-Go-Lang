package cmdmanager

import "fmt"

type CmdManager struct {
}

func New() CmdManager {
	return CmdManager{}
}
func (cmdmanager *CmdManager) LoadDataFromFile() ([]string, error) {

	fmt.Println("Please enter your prices. Confirm every price with Enter")
	var prices []string

	for {
		var price string
		fmt.Print("prices: ")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil

}
func (cmdmanager *CmdManager) WriteJson(data interface{}) error {

	fmt.Println(data)
	return nil
}
