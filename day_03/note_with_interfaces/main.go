package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

// type outputTable interface{
// 	saver
// 	displayer
// }

// also can be done by

type outputTable interface {
	Save() error
	Display()
}

func main() {

	todoText := getUserInput("Enter ToDo: ")
	title := getUserInput("Enter Title: ")
	contant := getUserInput("Enter contant: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Print(err)
		return
	}
	note, err := note.New(title, contant)
	if err != nil {
		return
	}
	error := outPutAndSave(todo)
	if error != nil {
		fmt.Println(error)
		return
	}

	erro := outPutAndSave(note)

	if erro != nil {
		fmt.Println(error)
		return
	}
	printSomeThing(1)
	printSomeThing(1.5)
	printSomeThing(todo)

	fmt.Println("Successfully Saved")

}

func printSomeThing(value interface{}) { // this will take any kind of value

	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", intVal)
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String: ", stringVal)
	}

	float64Val, ok := value.(float64)

	if ok {
		fmt.Println("float64: ", float64Val)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Interger", value)
	// case float64:
	// 	fmt.Println("Float", value)
	// case string:
	// 	fmt.Print("go todo")

	// }

	// fmt.Println(value)
}
func outPutAndSave(data outputTable) error {
	data.Display()
	return data.Save()
}

func getUserInput(promptText string) (value string) {
	fmt.Print(promptText)
	reader := bufio.NewReader(os.Stdin)
	value, _ = reader.ReadString('\n')

	return value

}
