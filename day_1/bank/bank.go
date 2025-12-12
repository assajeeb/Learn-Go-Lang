package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var fileName string = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644) /// 0644 is the file permission
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0.0, errors.New("File Not Found")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64) //here 64 mean the float64
	if err != nil {
		return 0.0, errors.New("Faild to Parse The file")
	}
	return balance, nil

}

func main() {
	AccountBalance, error := getBalanceFromFile()
	if error != nil {
		fmt.Print(error)
		panic(error)
	}
	fmt.Println("Welcome to Go Bank!")
	// for i := 1; i < 20; i++ {
	// 	fmt.Println(i)
	// }

	for {
		var Choise int
		fmt.Println("Enter Your Choice")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposite Balance")
		fmt.Println("3. Withdraw Balnce")
		fmt.Println("4. Exit")

		fmt.Scan(&Choise)
		fmt.Println("Your choise is: ", Choise)

		switch Choise {
		case 1:
			fmt.Println("You Balnace is: ", AccountBalance)
		case 2:
			var depositeAmount float64
			fmt.Print("Enter the Deposite Amount: ")
			fmt.Scan(&depositeAmount)

			if depositeAmount <= 0 {
				fmt.Println("Invalid Amount. ")
				continue
			}
			AccountBalance += depositeAmount
			fmt.Println("Balance Updated! Your New Balance is: ", AccountBalance)
			writeBalanceToFile(AccountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter the withdraw Amount: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 || withdrawAmount > AccountBalance {
				fmt.Println("Invalid Withdraw. Try Again")
				continue
			}
			AccountBalance -= withdrawAmount
			fmt.Println("Balance updated! New Balance is: ", AccountBalance)
			writeBalanceToFile(AccountBalance)
		default:
			fmt.Println(" Good Bye! ")
			fmt.Println("Thanks for Using our Go Bank")
			return

		}

		// if Choise == 1 {
		// 	fmt.Println("You Balnace is: ", AccountBalance)
		// } else if Choise == 2 {
		// 	var depositeAmount float64
		// 	fmt.Print("Enter the Deposite Amount: ")
		// 	fmt.Scan(&depositeAmount)

		// 	if depositeAmount <= 0 {
		// 		fmt.Println("Invalid Amount. ")
		// 		continue
		// 	}
		// 	AccountBalance += depositeAmount
		// 	fmt.Println("Balance Updated! Your New Balance is: ", AccountBalance)

		// } else if Choise == 3 {
		// 	var withdrawAmount float64
		// 	fmt.Print("Enter the withdraw Amount: ")
		// 	fmt.Scan(&withdrawAmount)

		// 	if withdrawAmount <= 0 || withdrawAmount > AccountBalance {
		// 		fmt.Println("Invalid Withdraw. Try Again")
		// 		continue
		// 	}
		// 	AccountBalance -= withdrawAmount
		// 	fmt.Println("Balance updated! New Balance is: ", AccountBalance)
		// } else {
		// 	fmt.Println(" Good Bye! ")
		// 	break
		// }

	}

}
