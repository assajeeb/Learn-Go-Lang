package main

import (
	"fmt"
	"time"
)

func greeting(Pharse string, done chan bool) {
	fmt.Println("hello, ", Pharse)
	done <- true

}

func slowGreeting(Pharse string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("hi ", Pharse)
	done <- true
	close(done)
}

func main() {

	done := make(chan bool)
	go greeting("Hello world", done)
	go greeting("this is normal Greeting", done)
	go slowGreeting("this is slow greeting", done)
	go greeting("this is normal greeting again", done)

	for range done {

	}
}
