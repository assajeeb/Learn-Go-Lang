package main

import (
	"fmt"

	"example.com/structs-custom-types/user"
)

func main() {

	firstName := getUserData("Please enter your first naem: ")
	secondName := getUserData("Please enter your second name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(firstName, secondName, birthDate)
	if err != nil {
		fmt.Print(err)
		return
	}
	admin := user.NewAdmin("Admin", "passowrd")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) (value string) {

	fmt.Print(promptText)
	fmt.Scanln(&value)
	return
}
