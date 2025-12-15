package conversition

import (
	"fmt"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	InputPrice := []float64{}
	for _, lins := range strings {
		value, err := strconv.ParseFloat(lins, 64)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		InputPrice = append(InputPrice, value)
	}

	return InputPrice, nil

}
