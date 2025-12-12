package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteBalanceToFile(filename string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644) /// 0644 is the file permission
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0.0, errors.New("file Not Found")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64) //here 64 mean the float64
	if err != nil {
		return 0.0, errors.New("faild to Parse The file")
	}
	return value, nil

}
