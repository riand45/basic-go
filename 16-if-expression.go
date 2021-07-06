package main

import "fmt"

func main () {
	var name = "John"

	if name == "John" {
		fmt.Println("Hello John")
	} else if name == "Doe" {
		fmt.Println("Hello Doe")
	} else if name == "Lorem" {
		fmt.Println("Hello Lorem")
	} else {
		fmt.Println("Hi, how are you?")
	}

	if length := len(name); length > 5 {
		fmt.Println("The words to long.")
	} else {
		fmt.Println("The name passed.")
	}
}