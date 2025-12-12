package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

var fileName string = "balance.txt"

func main() {
	AccountBalance, error := fileops.GetFloatFromFile(fileName)
	if error != nil {
		fmt.Print(error)
		// panic(error)
	}
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7n ", randomdata.PhoneNumber())

	for {
		var Choise int
		communicationWitthOption()
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
			fileops.WriteBalanceToFile(fileName, AccountBalance)
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
			fileops.WriteBalanceToFile(fileName, AccountBalance)
		default:
			fmt.Println(" Good Bye! ")
			fmt.Println("Thanks for Using our Go Bank")
			return

		}

	}

}
