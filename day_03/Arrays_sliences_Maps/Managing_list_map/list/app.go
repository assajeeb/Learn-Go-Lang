package list

import "fmt"

func main() {

	prices := []float64{32.3, 31.33}
	fmt.Println(prices)

	prices[1] = 333

	prices = append(prices, 5.99, 33.33)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{1001.2, 90, 88, 77255}

	prices = append(prices, discountPrices...)
	fmt.Println(prices)

}

// func main() {
// 	var product [3]string = [3]string{"A Book"}
// 	prices := [5]float64{12.0, 22.33, 44.44, 8.88, 33.33}

// 	product[2] = "A carpet"

// 	fmt.Println(prices)
// 	fmt.Println(product)

// 	fmt.Println(prices[1])

// 	featuredPrices := prices[:4] // [a,b]   a is included and b is exclueded
// 	hightlightPrices := featuredPrices[1:3]
// 	hightlightPrices[0] = 1001.00
// 	fmt.Println(featuredPrices)
// 	fmt.Println(hightlightPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(hightlightPrices), cap(hightlightPrices), len(featuredPrices), cap(featuredPrices))
// 	hightlightPrices = hightlightPrices[:4]

// 	fmt.Println(hightlightPrices)

// 	fmt.Println(len(hightlightPrices), cap(hightlightPrices), len(featuredPrices), cap(featuredPrices))
// }
