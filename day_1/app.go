package main

import "fmt"

const this_gloablconst = 23

var this_is_globleconst2 = "hello"

func main() {
	hello, year := outputText("hello")
	fmt.Print(hello, year)
}

func outputText(text string) (name string, year int) {

	name, year = text, 0
	return
}
