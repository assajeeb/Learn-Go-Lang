package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/note/note"
)

func main() {

	title := getUserInput("Enter Title: ")
	contant := getUserInput("Enter contant: ")

	note, err := note.New(title, contant)
	if err != nil {
		return
	}
	note.Display()
	erro := note.Save()
	if erro != nil {
		fmt.Println(erro)
		return
	}
	fmt.Println("Successfully Saved")

}

func getUserInput(promptText string) (value string) {
	fmt.Print(promptText)
	reader := bufio.NewReader(os.Stdin)
	value, _ = reader.ReadString('\n')

	return value

}
